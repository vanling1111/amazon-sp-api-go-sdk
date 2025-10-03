// Copyright 2025 Amazon SP-API Go SDK Authors.
//
// This file is part of Amazon SP-API Go SDK.
//
// Amazon SP-API Go SDK is dual-licensed:
//
// 1. GNU Affero General Public License v3.0 (AGPL-3.0) for open source use
//   - Free for personal, educational, and open source projects
//   - Your project must also be open sourced under AGPL-3.0
//   - See: https://www.gnu.org/licenses/agpl-3.0.html
//
// 2. Commercial License for proprietary/commercial use
//   - Required for any commercial, enterprise, or proprietary use
//   - Allows closed source distribution
//   - Contact: vanling1111@gmail.com
//
// Unless you have obtained a commercial license, this file is licensed
// under AGPL-3.0. By using this software, you agree to comply with the
// terms of the applicable license. All rights reserved.
package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	creds, err := NewCredentials(
		"test-client-id",
		"test-client-secret",
		"test-refresh-token",
		EndpointNA,
	)
	if err != nil {
		t.Fatalf("NewCredentials() error = %v", err)
	}

	client := NewClient(creds)
	if client == nil {
		t.Fatal("NewClient() returned nil")
	}

	if client.credentials != creds {
		t.Error("Client credentials not set correctly")
	}

	if client.httpClient == nil {
		t.Error("Client httpClient not initialized")
	}

	if client.cache == nil {
		t.Error("Client cache not initialized")
	}
}

func TestClient_GetAccessToken(t *testing.T) {
	// 创建测试服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 验证请求方法
		if r.Method != http.MethodPost {
			t.Errorf("Request method = %v, want POST", r.Method)
		}

		// 验证 Content-Type
		contentType := r.Header.Get("Content-Type")
		if contentType != "application/x-www-form-urlencoded" {
			t.Errorf("Content-Type = %v, want application/x-www-form-urlencoded", contentType)
		}

		// 解析请求体
		if err := r.ParseForm(); err != nil {
			t.Errorf("ParseForm() error = %v", err)
		}

		// 验证参数
		if r.Form.Get("grant_type") != "refresh_token" {
			t.Errorf("grant_type = %v, want refresh_token", r.Form.Get("grant_type"))
		}

		// 返回成功响应
		resp := lwaResponse{
			AccessToken: "test-access-token",
			TokenType:   "bearer",
			ExpiresIn:   3600,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// 创建客户端
	creds, _ := NewCredentials(
		"test-client-id",
		"test-client-secret",
		"test-refresh-token",
		server.URL,
	)
	client := NewClient(creds)

	// 获取访问令牌
	ctx := context.Background()
	token, err := client.GetAccessToken(ctx)
	if err != nil {
		t.Fatalf("GetAccessToken() error = %v", err)
	}

	if token != "test-access-token" {
		t.Errorf("GetAccessToken() = %v, want test-access-token", token)
	}

	// 第二次调用应该从缓存获取
	token2, err := client.GetAccessToken(ctx)
	if err != nil {
		t.Fatalf("GetAccessToken() (cached) error = %v", err)
	}

	if token2 != token {
		t.Error("Cached token should be the same")
	}
}

func TestClient_GetAccessToken_Error(t *testing.T) {
	// 创建返回错误的测试服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := lwaResponse{
			Error:     "invalid_grant",
			ErrorDesc: "The provided authorization grant is invalid",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// 创建客户端
	creds, _ := NewCredentials(
		"test-client-id",
		"test-client-secret",
		"invalid-refresh-token",
		server.URL,
	)
	client := NewClient(creds)

	// 获取访问令牌应该失败
	ctx := context.Background()
	_, err := client.GetAccessToken(ctx)
	if err == nil {
		t.Error("GetAccessToken() error = nil, want error")
	}
}

func TestClient_RefreshToken(t *testing.T) {
	callCount := 0

	// 创建测试服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++

		resp := lwaResponse{
			AccessToken: "test-access-token-" + string(rune(callCount)),
			TokenType:   "bearer",
			ExpiresIn:   3600,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// 创建客户端
	creds, _ := NewCredentials(
		"test-client-id",
		"test-client-secret",
		"test-refresh-token",
		server.URL,
	)
	client := NewClient(creds)

	ctx := context.Background()

	// 第一次获取令牌
	token1, err := client.GetAccessToken(ctx)
	if err != nil {
		t.Fatalf("GetAccessToken() error = %v", err)
	}

	if callCount != 1 {
		t.Errorf("callCount = %v after first call, want 1", callCount)
	}

	// 第二次调用应该从缓存获取
	_, err = client.GetAccessToken(ctx)
	if err != nil {
		t.Fatalf("GetAccessToken() (cached) error = %v", err)
	}

	if callCount != 1 {
		t.Errorf("callCount = %v after cached call, want 1", callCount)
	}

	// 强制刷新
	token2, err := client.RefreshToken(ctx)
	if err != nil {
		t.Fatalf("RefreshToken() error = %v", err)
	}

	if callCount != 2 {
		t.Errorf("callCount = %v after RefreshToken(), want 2", callCount)
	}

	if token1 == token2 {
		t.Error("Refreshed token should be different")
	}
}

func TestClient_SetHTTPClient(t *testing.T) {
	creds, _ := NewCredentials(
		"test-client-id",
		"test-client-secret",
		"test-refresh-token",
		EndpointNA,
	)
	client := NewClient(creds)

	customHTTPClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	client.SetHTTPClient(customHTTPClient)

	if client.httpClient != customHTTPClient {
		t.Error("SetHTTPClient() did not set the HTTP client correctly")
	}
}

func TestClient_SetCache(t *testing.T) {
	creds, _ := NewCredentials(
		"test-client-id",
		"test-client-secret",
		"test-refresh-token",
		EndpointNA,
	)
	client := NewClient(creds)

	customCache := NewMemoryCache()
	client.SetCache(customCache)

	if client.cache != customCache {
		t.Error("SetCache() did not set the cache correctly")
	}
}
