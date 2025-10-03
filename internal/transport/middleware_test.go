// Copyright 2025 Amazon SP-API Go SDK Authors.
//
// This file is part of Amazon SP-API Go SDK.
//
// Amazon SP-API Go SDK is dual-licensed:
//
// 1. GNU Affero General Public License v3.0 (AGPL-3.0) for open source use
//    - Free for personal, educational, and open source projects
//    - Your project must also be open sourced under AGPL-3.0
//    - See: https://www.gnu.org/licenses/agpl-3.0.html
//
// 2. Commercial License for proprietary/commercial use
//    - Required for any commercial, enterprise, or proprietary use
//    - Allows closed source distribution
//    - Contact: vanling1111@gmail.com
//
// Unless you have obtained a commercial license, this file is licensed
// under AGPL-3.0. By using this software, you agree to comply with the
// terms of the applicable license. All rights reserved.
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
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestUserAgentMiddleware(t *testing.T) {
	var receivedUserAgent string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedUserAgent = r.Header.Get("User-Agent")
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// 使用空的 UserAgent 配置，避免 client.Do 中的默认设置覆盖
	config := DefaultConfig()
	config.UserAgent = ""
	client := NewClient(server.URL, config)

	userAgent := "test-agent/1.0.0"
	client.Use(UserAgentMiddleware(userAgent))

	req, _ := http.NewRequest(http.MethodGet, server.URL+"/test", nil)
	ctx := context.Background()
	_, err := client.Do(ctx, req)
	if err != nil {
		t.Fatalf("Do() error = %v", err)
	}

	if receivedUserAgent != userAgent {
		t.Errorf("User-Agent = %v, want %v", receivedUserAgent, userAgent)
	}
}

func TestHeaderMiddleware(t *testing.T) {
	var receivedHeader string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedHeader = r.Header.Get("X-Custom-Header")
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewClient(server.URL, nil)
	headers := map[string]string{
		"X-Custom-Header": "custom-value",
	}
	client.Use(HeaderMiddleware(headers))

	req, _ := http.NewRequest(http.MethodGet, server.URL+"/test", nil)
	ctx := context.Background()
	_, err := client.Do(ctx, req)
	if err != nil {
		t.Fatalf("Do() error = %v", err)
	}

	if receivedHeader != "custom-value" {
		t.Errorf("X-Custom-Header = %v, want %v", receivedHeader, "custom-value")
	}
}

func TestTimeoutMiddleware(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 模拟慢响应
		time.Sleep(200 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewClient(server.URL, nil)
	client.Use(TimeoutMiddleware(100 * time.Millisecond))

	req, _ := http.NewRequest(http.MethodGet, server.URL+"/test", nil)
	ctx := context.Background()
	_, err := client.Do(ctx, req)

	// 应该超时
	if err == nil {
		t.Error("Expected timeout error, got nil")
	}
}

func TestRequestIDMiddleware(t *testing.T) {
	var receivedRequestID string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedRequestID = r.Header.Get("X-Request-ID")
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewClient(server.URL, nil)
	client.Use(RequestIDMiddleware())

	req, _ := http.NewRequest(http.MethodGet, server.URL+"/test", nil)
	ctx := context.Background()
	_, err := client.Do(ctx, req)
	if err != nil {
		t.Fatalf("Do() error = %v", err)
	}

	if receivedRequestID == "" {
		t.Error("X-Request-ID should not be empty")
	}

	// 检查 Request ID 格式
	if len(receivedRequestID) < 5 || receivedRequestID[:4] != "req-" {
		t.Errorf("Request ID format incorrect: %s", receivedRequestID)
	}
}

func TestLoggingMiddleware(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewClient(server.URL, nil)
	client.Use(LoggingMiddleware())

	req, _ := http.NewRequest(http.MethodGet, server.URL+"/test", nil)
	ctx := context.Background()
	_, err := client.Do(ctx, req)
	if err != nil {
		t.Fatalf("Do() error = %v", err)
	}

	// 日志中间件应该不影响请求的执行
	// 主要是确保不会 panic 或返回错误
}

func TestMultipleMiddlewares(t *testing.T) {
	var requestIDSet bool
	var customHeaderSet bool

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestIDSet = r.Header.Get("X-Request-ID") != ""
		customHeaderSet = r.Header.Get("X-Custom") == "value"
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewClient(server.URL, nil)

	// 添加多个中间件
	client.Use(RequestIDMiddleware())
	client.Use(HeaderMiddleware(map[string]string{
		"X-Custom": "value",
	}))

	req, _ := http.NewRequest(http.MethodGet, server.URL+"/test", nil)
	ctx := context.Background()
	_, err := client.Do(ctx, req)
	if err != nil {
		t.Fatalf("Do() error = %v", err)
	}

	if !requestIDSet {
		t.Error("Request ID was not set")
	}

	if !customHeaderSet {
		t.Error("Custom header was not set")
	}
}

func TestGenerateRequestID(t *testing.T) {
	id1 := generateRequestID()

	// 稍微等待，确保纳秒时间戳不同
	time.Sleep(1 * time.Millisecond)

	id2 := generateRequestID()

	if id1 == "" {
		t.Error("Generated request ID should not be empty")
	}

	if id1 == id2 {
		t.Error("Request IDs should be unique")
	}
}

func TestDateMiddleware(t *testing.T) {
	// 创建测试处理器
	handler := func(ctx context.Context, req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString("OK")),
		}, nil
	}

	// 创建中间件
	middleware := DateMiddleware()
	wrappedHandler := middleware(handler)

	// 创建测试请求
	req, err := http.NewRequest(http.MethodGet, "https://example.com/test", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// 执行请求
	ctx := context.Background()
	resp, err := wrappedHandler(ctx, req)
	if err != nil {
		t.Fatalf("Handler returned error: %v", err)
	}
	defer resp.Body.Close()

	// 验证 x-amz-date 头部
	dateHeader := req.Header.Get("x-amz-date")
	if dateHeader == "" {
		t.Error("x-amz-date header not set")
	}

	// 验证日期格式：20190430T123600Z
	// 格式应该是：YYYYMMDDTHHmmssZ (16 个字符)
	if len(dateHeader) != 16 {
		t.Errorf("x-amz-date format incorrect, got length %d, want 16", len(dateHeader))
	}

	// 验证格式：应该以 T 分隔，以 Z 结尾
	if dateHeader[8] != 'T' {
		t.Errorf("x-amz-date should have 'T' at position 8, got %c", dateHeader[8])
	}
	if dateHeader[len(dateHeader)-1] != 'Z' {
		t.Errorf("x-amz-date should end with 'Z', got %c", dateHeader[len(dateHeader)-1])
	}

	t.Logf("x-amz-date: %s", dateHeader)
}

func TestDateMiddleware_DoesNotOverride(t *testing.T) {
	// 创建测试处理器
	handler := func(ctx context.Context, req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString("OK")),
		}, nil
	}

	// 创建中间件
	middleware := DateMiddleware()
	wrappedHandler := middleware(handler)

	// 创建测试请求，预设 x-amz-date
	req, err := http.NewRequest(http.MethodGet, "https://example.com/test", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	customDate := "20190430T123600Z"
	req.Header.Set("x-amz-date", customDate)

	// 执行请求
	ctx := context.Background()
	resp, err := wrappedHandler(ctx, req)
	if err != nil {
		t.Fatalf("Handler returned error: %v", err)
	}
	defer resp.Body.Close()

	// 验证 x-amz-date 头部未被覆盖
	dateHeader := req.Header.Get("x-amz-date")
	if dateHeader != customDate {
		t.Errorf("x-amz-date was overridden, got %s, want %s", dateHeader, customDate)
	}
}
