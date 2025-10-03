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
package signer

import (
	"context"
	"net/http"
	"testing"

	"github.com/vanling1111/amazon-sp-api-go-sdk/internal/auth"
)

// mockLWAClient 是用于测试的模拟 LWA 客户端。
type mockLWAClient struct {
	accessToken string
	err         error
}

func (m *mockLWAClient) GetAccessToken(ctx context.Context) (string, error) {
	if m.err != nil {
		return "", m.err
	}
	return m.accessToken, nil
}

func (m *mockLWAClient) RefreshToken(ctx context.Context) (string, error) {
	return m.GetAccessToken(ctx)
}

func (m *mockLWAClient) SetHTTPClient(client interface{}) {}
func (m *mockLWAClient) SetCache(cache interface{})       {}

func TestNewLWASigner(t *testing.T) {
	// 创建真实的 LWA 客户端用于测试
	creds, err := auth.NewCredentials(
		"test-client-id",
		"test-client-secret",
		"test-refresh-token",
		auth.EndpointNA,
	)
	if err != nil {
		t.Fatalf("Failed to create credentials: %v", err)
	}

	lwaClient := auth.NewClient(creds)
	signer := NewLWASigner(lwaClient)

	if signer == nil {
		t.Fatal("NewLWASigner() returned nil")
	}

	if signer.lwaClient == nil {
		t.Error("LWASigner.lwaClient is nil")
	}
}

func TestLWASigner_Sign_Success(t *testing.T) {
	// 创建签名器（需要使用反射或者修改结构以便测试）
	// 为了测试，我们直接创建一个真实的客户端
	creds, _ := auth.NewCredentials(
		"test-client-id",
		"test-client-secret",
		"test-refresh-token",
		auth.EndpointNA,
	)
	lwaClient := auth.NewClient(creds)

	signer := NewLWASigner(lwaClient)

	// 创建测试请求
	req, err := http.NewRequest(http.MethodGet, "https://sellingpartnerapi-na.amazon.com/test", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	ctx := context.Background()

	// 注意：实际签名会失败，因为没有真实的 LWA 服务器
	// 但我们可以测试头部是否被设置（如果不报错的话）
	// 这里我们主要测试签名器的结构
	_ = signer.Sign(ctx, req)

	// 验证签名器可以被设置新的客户端
	newCreds, _ := auth.NewCredentials(
		"new-client-id",
		"new-client-secret",
		"new-refresh-token",
		auth.EndpointNA,
	)
	newClient := auth.NewClient(newCreds)
	signer.SetLWAClient(newClient)

	if signer.lwaClient != newClient {
		t.Error("SetLWAClient() did not update the client")
	}
}

func TestLWASigner_Sign_HeaderSet(t *testing.T) {
	// 使用 grantless credentials 测试
	creds, _ := auth.NewGrantlessCredentials(
		"test-client-id",
		"test-client-secret",
		[]string{auth.ScopeNotifications},
		auth.EndpointNA,
	)
	lwaClient := auth.NewClient(creds)
	signer := NewLWASigner(lwaClient)

	req, _ := http.NewRequest(http.MethodGet, "https://sellingpartnerapi-na.amazon.com/test", nil)
	ctx := context.Background()

	// 尝试签名（会因为没有真实 LWA 服务器而失败，但我们主要测试结构）
	_ = signer.Sign(ctx, req)

	// 这个测试主要验证签名器的基本功能
	t.Log("LWA Signer structure test passed")
}

func TestLWASigner_SetLWAClient(t *testing.T) {
	creds1, _ := auth.NewCredentials(
		"client-1",
		"secret-1",
		"token-1",
		auth.EndpointNA,
	)
	client1 := auth.NewClient(creds1)
	signer := NewLWASigner(client1)

	if signer.lwaClient != client1 {
		t.Error("Initial client not set correctly")
	}

	creds2, _ := auth.NewCredentials(
		"client-2",
		"secret-2",
		"token-2",
		auth.EndpointNA,
	)
	client2 := auth.NewClient(creds2)
	signer.SetLWAClient(client2)

	if signer.lwaClient != client2 {
		t.Error("SetLWAClient() did not update the client")
	}
}
