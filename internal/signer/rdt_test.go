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
package signer

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"testing"
)

// mockRDTProvider 是用于测试的模拟 RDT 提供者。
func mockRDTProvider(rdt string, err error) RDTProvider {
	return func(ctx context.Context, resourcePath string, dataElements []string) (string, error) {
		if err != nil {
			return "", err
		}
		return rdt, nil
	}
}

func TestNewRDTSigner(t *testing.T) {
	provider := mockRDTProvider("test-rdt-token", nil)
	signer := NewRDTSigner(provider)

	if signer == nil {
		t.Fatal("NewRDTSigner() returned nil")
	}

	if signer.rdtProvider == nil {
		t.Error("RDTSigner.rdtProvider is nil")
	}
}

func TestRDTSigner_Sign_Success(t *testing.T) {
	testRDT := "test-rdt-token-123"
	provider := mockRDTProvider(testRDT, nil)
	signer := NewRDTSigner(provider)

	req, err := http.NewRequest(http.MethodGet, "https://sellingpartnerapi-na.amazon.com/orders/v0/orders/123", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	ctx := context.Background()
	err = signer.Sign(ctx, req)
	if err != nil {
		t.Fatalf("Sign() error = %v", err)
	}

	// 验证 x-amz-access-token 头部被设置
	token := req.Header.Get("x-amz-access-token")
	if token != testRDT {
		t.Errorf("x-amz-access-token = %s, want %s", token, testRDT)
	}
}

func TestRDTSigner_Sign_ProviderError(t *testing.T) {
	expectedError := errors.New("failed to get RDT")
	provider := mockRDTProvider("", expectedError)
	signer := NewRDTSigner(provider)

	req, _ := http.NewRequest(http.MethodGet, "https://sellingpartnerapi-na.amazon.com/orders/v0/orders/123", nil)
	ctx := context.Background()

	err := signer.Sign(ctx, req)
	if err == nil {
		t.Fatal("Sign() error = nil, want error")
	}

	if !strings.Contains(err.Error(), "get RDT") {
		t.Errorf("Sign() error = %v, should contain 'get RDT'", err)
	}
}

func TestRDTSigner_Sign_WithDataElements(t *testing.T) {
	testRDT := "test-rdt-with-data-elements"

	provider := func(ctx context.Context, resourcePath string, dataElements []string) (string, error) {
		return testRDT, nil
	}

	signer := NewRDTSigner(provider)

	req, _ := http.NewRequest(http.MethodGet, "https://sellingpartnerapi-na.amazon.com/orders/v0/orders/123", nil)

	// 模拟设置数据元素（实际使用中可能通过 context 或其他方式传递）
	ctx := context.Background()

	err := signer.Sign(ctx, req)
	if err != nil {
		t.Fatalf("Sign() error = %v", err)
	}

	// 验证令牌被设置
	token := req.Header.Get("x-amz-access-token")
	if token != testRDT {
		t.Errorf("x-amz-access-token = %s, want %s", token, testRDT)
	}
}

func TestRDTSigner_SetRDTProvider(t *testing.T) {
	provider1 := mockRDTProvider("token1", nil)
	signer := NewRDTSigner(provider1)

	// 使用第一个 provider
	req1, _ := http.NewRequest(http.MethodGet, "https://sellingpartnerapi-na.amazon.com/orders/v0/orders/123", nil)
	ctx := context.Background()
	_ = signer.Sign(ctx, req1)

	token1 := req1.Header.Get("x-amz-access-token")
	if token1 != "token1" {
		t.Errorf("First provider token = %s, want token1", token1)
	}

	// 更换 provider
	provider2 := mockRDTProvider("token2", nil)
	signer.SetRDTProvider(provider2)

	// 使用第二个 provider
	req2, _ := http.NewRequest(http.MethodGet, "https://sellingpartnerapi-na.amazon.com/orders/v0/orders/456", nil)
	_ = signer.Sign(ctx, req2)

	token2 := req2.Header.Get("x-amz-access-token")
	if token2 != "token2" {
		t.Errorf("Second provider token = %s, want token2", token2)
	}
}

func TestRDTSigner_RequiresRDT_Header(t *testing.T) {
	provider := mockRDTProvider("test-rdt", nil)
	signer := NewRDTSigner(provider)

	// 测试：设置了 x-amzn-RDT-Required 头
	req, _ := http.NewRequest(http.MethodGet, "https://sellingpartnerapi-na.amazon.com/catalog/items/B123", nil)
	req.Header.Set("x-amzn-RDT-Required", "true")

	ctx := context.Background()
	err := signer.Sign(ctx, req)
	if err != nil {
		t.Fatalf("Sign() error = %v", err)
	}

	token := req.Header.Get("x-amz-access-token")
	if token != "test-rdt" {
		t.Errorf("RDT not applied when header set, got token = %s", token)
	}
}

func TestRDTSigner_RestrictedPaths(t *testing.T) {
	tests := []struct {
		name string
		path string
	}{
		{"orders API", "https://sellingpartnerapi-na.amazon.com/orders/v0/orders/123"},
		{"order address", "https://sellingpartnerapi-na.amazon.com/orders/v0/orders/123/address"},
		{"order buyerInfo", "https://sellingpartnerapi-na.amazon.com/orders/v0/orders/123/buyerInfo"},
		{"MFN API", "https://sellingpartnerapi-na.amazon.com/mfn/v0/shipments/123"},
		{"messaging API", "https://sellingpartnerapi-na.amazon.com/messaging/v1/orders/123"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			provider := mockRDTProvider("test-rdt-"+tt.name, nil)
			signer := NewRDTSigner(provider)

			req, _ := http.NewRequest(http.MethodGet, tt.path, nil)
			ctx := context.Background()

			err := signer.Sign(ctx, req)
			if err != nil {
				t.Fatalf("Sign() error = %v", err)
			}

			token := req.Header.Get("x-amz-access-token")
			if token == "" {
				t.Error("RDT should be applied for restricted path")
			}
		})
	}
}

func TestRDTSigner_ExtractDataElements(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		header   string
		expected int
	}{
		{
			name:     "from header",
			url:      "https://sellingpartnerapi-na.amazon.com/orders/v0/orders/999",
			header:   "buyerInfo,shippingAddress",
			expected: 2,
		},
		{
			name:     "order address path",
			url:      "https://sellingpartnerapi-na.amazon.com/orders/v0/orders/123/address",
			header:   "",
			expected: 2, // buyerInfo, shippingAddress
		},
		{
			name:     "order buyerInfo path",
			url:      "https://sellingpartnerapi-na.amazon.com/orders/v0/orders/123/buyerInfo",
			header:   "",
			expected: 1, // buyerInfo
		},
		{
			name:     "default orders path",
			url:      "https://sellingpartnerapi-na.amazon.com/orders/v0/orders/123",
			header:   "",
			expected: 2, // buyerInfo, shippingAddress
		},
		{
			name:     "non-restricted path",
			url:      "https://sellingpartnerapi-na.amazon.com/catalog/items/B123",
			header:   "",
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var capturedElements []string
			provider := func(ctx context.Context, resourcePath string, dataElements []string) (string, error) {
				capturedElements = dataElements
				return "test-rdt", nil
			}

			signer := NewRDTSigner(provider)
			req, _ := http.NewRequest(http.MethodGet, tt.url, nil)

			if tt.header != "" {
				req.Header.Set("x-amzn-RDT-DataElements", tt.header)
			}

			ctx := context.Background()
			_ = signer.Sign(ctx, req)

			if len(capturedElements) != tt.expected {
				t.Errorf("extracted %d elements, want %d. Elements: %v",
					len(capturedElements), tt.expected, capturedElements)
			}
		})
	}
}
