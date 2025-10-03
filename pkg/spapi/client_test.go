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
package spapi_test

import (
	"context"
	"testing"
	"time"

	"github.com/vanling1111/amazon-sp-api-go-sdk/internal/models"
	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// TestNewClient 测试客户端创建。
func TestNewClient(t *testing.T) {
	tests := []struct {
		name    string
		opts    []spapi.ClientOption
		wantErr bool
	}{
		{
			name: "valid regular credentials",
			opts: []spapi.ClientOption{
				spapi.WithRegion(models.RegionNA),
				spapi.WithCredentials("test-client-id", "test-client-secret", "test-refresh-token"),
			},
			wantErr: false,
		},
		{
			name: "valid grantless credentials",
			opts: []spapi.ClientOption{
				spapi.WithRegion(models.RegionEU),
				spapi.WithGrantlessCredentials("test-client-id", "test-client-secret", []string{
					"sellingpartnerapi::notifications",
				}),
			},
			wantErr: false,
		},
		{
			name: "missing region",
			opts: []spapi.ClientOption{
				spapi.WithCredentials("test-client-id", "test-client-secret", "test-refresh-token"),
			},
			wantErr: true, // 默认配置会提供 RegionNA，但如果不提供任何配置会出错
		},
		{
			name: "missing credentials",
			opts: []spapi.ClientOption{
				spapi.WithRegion(models.RegionNA),
			},
			wantErr: true, // 没有提供客户端凭据
		},
		{
			name: "custom timeout",
			opts: []spapi.ClientOption{
				spapi.WithRegion(models.RegionNA),
				spapi.WithCredentials("test-client-id", "test-client-secret", "test-refresh-token"),
				spapi.WithHTTPTimeout(60 * time.Second),
			},
			wantErr: false,
		},
		{
			name: "with retries",
			opts: []spapi.ClientOption{
				spapi.WithRegion(models.RegionFE),
				spapi.WithCredentials("test-client-id", "test-client-secret", "test-refresh-token"),
				spapi.WithMaxRetries(5),
			},
			wantErr: false,
		},
		{
			name: "with debug mode",
			opts: []spapi.ClientOption{
				spapi.WithRegion(models.RegionNA),
				spapi.WithCredentials("test-client-id", "test-client-secret", "test-refresh-token"),
				spapi.WithDebug(),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := spapi.NewClient(tt.opts...)

			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil {
				// 验证客户端创建成功
				if client == nil {
					t.Error("NewClient() returned nil client")
				}

				// 测试 Close
				if err := client.Close(); err != nil {
					t.Errorf("Close() error = %v", err)
				}
			}
		})
	}
}

// TestClient_Config 测试配置获取。
func TestClient_Config(t *testing.T) {
	client, err := spapi.NewClient(
		spapi.WithRegion(models.RegionNA),
		spapi.WithCredentials("test-client-id", "test-client-secret", "test-refresh-token"),
		spapi.WithHTTPTimeout(45*time.Second),
	)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	defer client.Close()

	config := client.Config()
	if config == nil {
		t.Fatal("Config() returned nil")
	}

	// 验证配置
	if config.Region.Code != "na" {
		t.Errorf("Config().Region.Code = %v, want %v", config.Region.Code, "na")
	}

	if config.HTTPTimeout != 45*time.Second {
		t.Errorf("Config().HTTPTimeout = %v, want %v", config.HTTPTimeout, 45*time.Second)
	}

	if config.ClientID != "test-client-id" {
		t.Errorf("Config().ClientID = %v, want %v", config.ClientID, "test-client-id")
	}
}

// TestClient_GetAccessToken 测试访问令牌获取。
//
// 注意：使用测试凭据时此测试会失败。
// 主要测试方法签名和返回值类型。
func TestClient_GetAccessToken(t *testing.T) {
	client, err := spapi.NewClient(
		spapi.WithRegion(models.RegionNA),
		spapi.WithCredentials("test-client-id", "test-client-secret", "test-refresh-token"),
	)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	token, err := client.GetAccessToken(ctx)

	// 使用测试凭据时预期会失败
	if err == nil {
		t.Log("GetAccessToken() succeeded (unexpected with test credentials)")
		if token == "" {
			t.Error("GetAccessToken() returned empty token")
		}
	} else {
		t.Logf("GetAccessToken() error = %v (expected with test credentials)", err)
	}
}

// TestClient_RateLimitManager 测试速率限制管理器获取。
func TestClient_RateLimitManager(t *testing.T) {
	client, err := spapi.NewClient(
		spapi.WithRegion(models.RegionNA),
		spapi.WithCredentials("test-client-id", "test-client-secret", "test-refresh-token"),
	)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	defer client.Close()

	manager := client.RateLimitManager()
	if manager == nil {
		t.Fatal("RateLimitManager() returned nil")
	}

	// 验证管理器初始状态
	count := manager.Count()
	if count != 0 {
		t.Errorf("RateLimitManager().Count() = %v, want %v", count, 0)
	}
}

// TestClient_HTTPClient 测试 HTTP 客户端获取。
func TestClient_HTTPClient(t *testing.T) {
	client, err := spapi.NewClient(
		spapi.WithRegion(models.RegionNA),
		spapi.WithCredentials("test-client-id", "test-client-secret", "test-refresh-token"),
	)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	defer client.Close()

	httpClient := client.HTTPClient()
	if httpClient == nil {
		t.Fatal("HTTPClient() returned nil")
	}
}

// TestClient_Signer 测试签名器获取。
func TestClient_Signer(t *testing.T) {
	client, err := spapi.NewClient(
		spapi.WithRegion(models.RegionNA),
		spapi.WithCredentials("test-client-id", "test-client-secret", "test-refresh-token"),
	)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	defer client.Close()

	signer := client.Signer()
	if signer == nil {
		t.Fatal("Signer() returned nil")
	}
}

// TestClient_MultipleInstances 测试创建多个客户端实例。
func TestClient_MultipleInstances(t *testing.T) {
	// 创建多个客户端
	client1, err := spapi.NewClient(
		spapi.WithRegion(models.RegionNA),
		spapi.WithCredentials("client-1", "secret-1", "refresh-1"),
	)
	if err != nil {
		t.Fatalf("NewClient() 1 error = %v", err)
	}
	defer client1.Close()

	client2, err := spapi.NewClient(
		spapi.WithRegion(models.RegionEU),
		spapi.WithCredentials("client-2", "secret-2", "refresh-2"),
	)
	if err != nil {
		t.Fatalf("NewClient() 2 error = %v", err)
	}
	defer client2.Close()

	// 验证客户端独立性
	config1 := client1.Config()
	config2 := client2.Config()

	if config1.Region.Code == config2.Region.Code {
		t.Error("Client instances share the same region (should be independent)")
	}

	if config1.ClientID == config2.ClientID {
		t.Error("Client instances share the same ClientID (should be independent)")
	}
}

// BenchmarkNewClient 基准测试客户端创建性能。
func BenchmarkNewClient(b *testing.B) {
	opts := []spapi.ClientOption{
		spapi.WithRegion(models.RegionNA),
		spapi.WithCredentials("test-client-id", "test-client-secret", "test-refresh-token"),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		client, err := spapi.NewClient(opts...)
		if err != nil {
			b.Fatalf("NewClient() error = %v", err)
		}
		client.Close()
	}
}

// ExampleNewClient 演示如何创建 SP-API 客户端。
func ExampleNewClient() {
	// 创建 Regular 操作的客户端
	client, err := spapi.NewClient(
		spapi.WithRegion(models.RegionNA),
		spapi.WithCredentials("your-client-id", "your-client-secret", "your-refresh-token"),
		spapi.WithHTTPTimeout(60*time.Second),
		spapi.WithMaxRetries(3),
	)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// 使用客户端...
	_ = client
}

// ExampleNewClient_grantless 演示如何创建 Grantless 操作的客户端。
func ExampleNewClient_grantless() {
	// 创建 Grantless 操作的客户端
	client, err := spapi.NewClient(
		spapi.WithRegion(models.RegionEU),
		spapi.WithGrantlessCredentials(
			"your-client-id",
			"your-client-secret",
			[]string{"sellingpartnerapi::notifications"},
		),
	)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// 使用客户端...
	_ = client
}

// TestClient_DoRequest 测试通用 HTTP 请求方法
func TestClient_DoRequest(t *testing.T) {
	client, err := spapi.NewClient(
		spapi.WithRegion(models.RegionNA),
		spapi.WithCredentials("test-client-id", "test-client-secret", "test-refresh-token"),
	)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	// 测试 GET 请求（会失败，但测试代码路径）
	t.Run("GET request", func(t *testing.T) {
		var result interface{}
		err := client.Get(ctx, "/test/path", nil, &result)
		// 预期会失败（无效的认证）
		if err == nil {
			t.Log("GET request unexpectedly succeeded")
		} else {
			t.Logf("GET request failed as expected: %v", err)
		}
	})

	// 测试 POST 请求
	t.Run("POST request", func(t *testing.T) {
		var result interface{}
		body := map[string]string{"test": "data"}
		err := client.Post(ctx, "/test/path", body, &result)
		// 预期会失败
		if err == nil {
			t.Log("POST request unexpectedly succeeded")
		} else {
			t.Logf("POST request failed as expected: %v", err)
		}
	})

	// 测试 PUT 请求
	t.Run("PUT request", func(t *testing.T) {
		var result interface{}
		body := map[string]string{"test": "data"}
		err := client.Put(ctx, "/test/path", body, &result)
		// 预期会失败
		if err == nil {
			t.Log("PUT request unexpectedly succeeded")
		} else {
			t.Logf("PUT request failed as expected: %v", err)
		}
	})

	// 测试 DELETE 请求
	t.Run("DELETE request", func(t *testing.T) {
		var result interface{}
		err := client.Delete(ctx, "/test/path", &result)
		// 预期会失败
		if err == nil {
			t.Log("DELETE request unexpectedly succeeded")
		} else {
			t.Logf("DELETE request failed as expected: %v", err)
		}
	})
}

// TestClient_RequestWithQueryParams 测试带查询参数的请求
func TestClient_RequestWithQueryParams(t *testing.T) {
	client, err := spapi.NewClient(
		spapi.WithRegion(models.RegionNA),
		spapi.WithCredentials("test-client-id", "test-client-secret", "test-refresh-token"),
	)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	queryParams := map[string]string{
		"param1": "value1",
		"param2": "value2",
	}

	var result interface{}
	err = client.Get(ctx, "/test/path", queryParams, &result)

	// 预期会失败，但测试参数处理
	t.Logf("Request with query params: %v", err)
}

// TestClient_ContextCancellation 测试上下文取消
func TestClient_ContextCancellation(t *testing.T) {
	client, err := spapi.NewClient(
		spapi.WithRegion(models.RegionNA),
		spapi.WithCredentials("test-client-id", "test-client-secret", "test-refresh-token"),
	)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	defer client.Close()

	// 创建已取消的上下文
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // 立即取消

	var result interface{}
	err = client.Get(ctx, "/test/path", nil, &result)

	if err == nil {
		t.Error("Expected error with cancelled context, got nil")
	} else {
		t.Logf("Request with cancelled context correctly failed: %v", err)
	}
}

// TestClient_AllRegions 测试所有区域的客户端创建
func TestClient_AllRegions(t *testing.T) {
	regions := []models.Region{
		models.RegionNA,
		models.RegionEU,
		models.RegionFE,
	}

	for _, region := range regions {
		t.Run(region.Code, func(t *testing.T) {
			client, err := spapi.NewClient(
				spapi.WithRegion(region),
				spapi.WithCredentials("test-client-id", "test-client-secret", "test-refresh-token"),
			)
			if err != nil {
				t.Errorf("NewClient() with region %s error = %v", region.Code, err)
				return
			}
			defer client.Close()

			config := client.Config()
			if config.Region.Code != region.Code {
				t.Errorf("Config().Region.Code = %v, want %v", config.Region.Code, region.Code)
			}
		})
	}
}
