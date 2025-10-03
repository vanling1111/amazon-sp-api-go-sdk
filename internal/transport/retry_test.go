// Copyright 2025 Amazon SP-API Go SDK Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package transport

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

func TestRetryMiddleware_SuccessFirstAttempt(t *testing.T) {
	callCount := int32(0)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&callCount, 1)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewClient(server.URL, nil)
	client.Use(RetryMiddleware(nil))

	req, _ := http.NewRequest(http.MethodGet, server.URL+"/test", nil)
	ctx := context.Background()
	resp, err := client.Do(ctx, req)
	if err != nil {
		t.Fatalf("Do() error = %v", err)
	}
	defer resp.Body.Close()

	if atomic.LoadInt32(&callCount) != 1 {
		t.Errorf("Expected 1 call, got %d", callCount)
	}
}

func TestRetryMiddleware_RetryOn500(t *testing.T) {
	callCount := int32(0)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		count := atomic.AddInt32(&callCount, 1)
		if count < 3 {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}))
	defer server.Close()

	config := &RetryConfig{
		MaxRetries:      3,
		InitialInterval: 10 * time.Millisecond,
		Multiplier:      2.0,
	}

	client := NewClient(server.URL, nil)
	client.Use(RetryMiddleware(config))

	req, _ := http.NewRequest(http.MethodGet, server.URL+"/test", nil)
	ctx := context.Background()
	resp, err := client.Do(ctx, req)
	if err != nil {
		t.Fatalf("Do() error = %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if atomic.LoadInt32(&callCount) != 3 {
		t.Errorf("Expected 3 calls, got %d", callCount)
	}
}

func TestRetryMiddleware_MaxRetriesExceeded(t *testing.T) {
	callCount := int32(0)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&callCount, 1)
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	config := &RetryConfig{
		MaxRetries:      2,
		InitialInterval: 10 * time.Millisecond,
		Multiplier:      2.0,
	}

	client := NewClient(server.URL, nil)
	client.Use(RetryMiddleware(config))

	req, _ := http.NewRequest(http.MethodGet, server.URL+"/test", nil)
	ctx := context.Background()
	resp, err := client.Do(ctx, req)
	if err != nil {
		t.Fatalf("Do() error = %v", err)
	}
	defer resp.Body.Close()

	// 应该执行 1 次初始请求 + 2 次重试 = 3 次
	expectedCalls := int32(3)
	if atomic.LoadInt32(&callCount) != expectedCalls {
		t.Errorf("Expected %d calls, got %d", expectedCalls, callCount)
	}

	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %v, want %v", resp.StatusCode, http.StatusInternalServerError)
	}
}

func TestDefaultShouldRetry(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		err        error
		want       bool
	}{
		{
			name:       "network error",
			statusCode: 0,
			err:        errors.New("network error"),
			want:       true,
		},
		{
			name:       "500 error",
			statusCode: 500,
			err:        nil,
			want:       true,
		},
		{
			name:       "503 error",
			statusCode: 503,
			err:        nil,
			want:       true,
		},
		{
			name:       "429 error",
			statusCode: 429,
			err:        nil,
			want:       true,
		},
		{
			name:       "200 success",
			statusCode: 200,
			err:        nil,
			want:       false,
		},
		{
			name:       "400 error",
			statusCode: 400,
			err:        nil,
			want:       false,
		},
		{
			name:       "404 error",
			statusCode: 404,
			err:        nil,
			want:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var resp *http.Response
			if tt.statusCode != 0 {
				resp = &http.Response{
					StatusCode: tt.statusCode,
				}
			}

			got := defaultShouldRetry(resp, tt.err)
			if got != tt.want {
				t.Errorf("defaultShouldRetry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateBackoff(t *testing.T) {
	tests := []struct {
		name       string
		attempt    int
		initial    time.Duration
		max        time.Duration
		multiplier float64
		wantMin    time.Duration
		wantMax    time.Duration
	}{
		{
			name:       "first attempt",
			attempt:    0,
			initial:    1 * time.Second,
			max:        30 * time.Second,
			multiplier: 2.0,
			wantMin:    1 * time.Second,
			wantMax:    1 * time.Second,
		},
		{
			name:       "second attempt",
			attempt:    1,
			initial:    1 * time.Second,
			max:        30 * time.Second,
			multiplier: 2.0,
			wantMin:    2 * time.Second,
			wantMax:    2 * time.Second,
		},
		{
			name:       "third attempt",
			attempt:    2,
			initial:    1 * time.Second,
			max:        30 * time.Second,
			multiplier: 2.0,
			wantMin:    4 * time.Second,
			wantMax:    4 * time.Second,
		},
		{
			name:       "exceeds max",
			attempt:    10,
			initial:    1 * time.Second,
			max:        30 * time.Second,
			multiplier: 2.0,
			wantMin:    30 * time.Second,
			wantMax:    30 * time.Second,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateBackoff(tt.attempt, tt.initial, tt.max, tt.multiplier)

			if got < tt.wantMin || got > tt.wantMax {
				t.Errorf("calculateBackoff() = %v, want between %v and %v",
					got, tt.wantMin, tt.wantMax)
			}
		})
	}
}

func TestDefaultRetryConfig(t *testing.T) {
	config := DefaultRetryConfig()

	if config == nil {
		t.Fatal("DefaultRetryConfig() returned nil")
	}

	if config.MaxRetries <= 0 {
		t.Error("MaxRetries should be > 0")
	}

	if config.InitialInterval <= 0 {
		t.Error("InitialInterval should be > 0")
	}

	if config.Multiplier <= 1.0 {
		t.Error("Multiplier should be > 1.0")
	}

	if config.ShouldRetry == nil {
		t.Error("ShouldRetry should not be nil")
	}
}
