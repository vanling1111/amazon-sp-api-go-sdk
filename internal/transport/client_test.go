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
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	baseURL := "https://api.example.com"

	client := NewClient(baseURL, nil)

	if client == nil {
		t.Fatal("NewClient() returned nil")
	}

	if client.BaseURL() != baseURL {
		t.Errorf("BaseURL() = %v, want %v", client.BaseURL(), baseURL)
	}

	if client.httpClient == nil {
		t.Error("HTTP client not initialized")
	}

	if client.config == nil {
		t.Error("Config not initialized")
	}
}

func TestNewClient_WithCustomConfig(t *testing.T) {
	baseURL := "https://api.example.com"
	config := &Config{
		Timeout:   10 * time.Second,
		UserAgent: "custom-agent/1.0.0",
	}

	client := NewClient(baseURL, config)

	if client == nil {
		t.Fatal("NewClient() returned nil")
	}

	if client.config.Timeout != config.Timeout {
		t.Errorf("Timeout = %v, want %v", client.config.Timeout, config.Timeout)
	}

	if client.config.UserAgent != config.UserAgent {
		t.Errorf("UserAgent = %v, want %v", client.config.UserAgent, config.UserAgent)
	}
}

func TestClient_Do(t *testing.T) {
	// 创建测试服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}))
	defer server.Close()

	client := NewClient(server.URL, nil)

	req, err := http.NewRequest(http.MethodGet, server.URL+"/test", nil)
	if err != nil {
		t.Fatalf("NewRequest() error = %v", err)
	}

	ctx := context.Background()
	resp, err := client.Do(ctx, req)
	if err != nil {
		t.Fatalf("Do() error = %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func TestClient_UseMiddleware(t *testing.T) {
	// 创建测试服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewClient(server.URL, nil)

	// 添加中间件
	middlewareCalled := false
	client.Use(func(next Handler) Handler {
		return func(ctx context.Context, req *http.Request) (*http.Response, error) {
			middlewareCalled = true
			return next(ctx, req)
		}
	})

	req, _ := http.NewRequest(http.MethodGet, server.URL+"/test", nil)
	ctx := context.Background()
	_, err := client.Do(ctx, req)
	if err != nil {
		t.Fatalf("Do() error = %v", err)
	}

	if !middlewareCalled {
		t.Error("Middleware was not called")
	}
}

func TestClient_UserAgent(t *testing.T) {
	// 创建测试服务器
	var receivedUserAgent string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedUserAgent = r.Header.Get("User-Agent")
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	config := &Config{
		UserAgent: "test-user-agent/1.0.0",
	}
	client := NewClient(server.URL, config)

	req, _ := http.NewRequest(http.MethodGet, server.URL+"/test", nil)
	ctx := context.Background()
	_, err := client.Do(ctx, req)
	if err != nil {
		t.Fatalf("Do() error = %v", err)
	}

	if receivedUserAgent != config.UserAgent {
		t.Errorf("User-Agent = %v, want %v", receivedUserAgent, config.UserAgent)
	}
}

func TestDefaultConfig(t *testing.T) {
	config := DefaultConfig()

	if config == nil {
		t.Fatal("DefaultConfig() returned nil")
	}

	if config.Timeout <= 0 {
		t.Error("Default timeout should be > 0")
	}

	if config.UserAgent == "" {
		t.Error("Default user agent should not be empty")
	}

	if config.MaxIdleConns <= 0 {
		t.Error("Default MaxIdleConns should be > 0")
	}
}

