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

package signer

import (
	"context"
	"errors"
	"net/http"
	"testing"
)

// mockSigner 是用于测试的模拟签名器。
type mockSigner struct {
	signFunc  func(ctx context.Context, req *http.Request) error
	callCount int
}

func (m *mockSigner) Sign(ctx context.Context, req *http.Request) error {
	m.callCount++
	if m.signFunc != nil {
		return m.signFunc(ctx, req)
	}
	return nil
}

func TestNewChainSigner(t *testing.T) {
	signer1 := &mockSigner{}
	signer2 := &mockSigner{}

	chain := NewChainSigner(signer1, signer2)

	if chain == nil {
		t.Fatal("NewChainSigner() returned nil")
	}

	if len(chain.signers) != 2 {
		t.Errorf("Expected 2 signers, got %d", len(chain.signers))
	}
}

func TestChainSigner_Sign(t *testing.T) {
	signer1 := &mockSigner{}
	signer2 := &mockSigner{}

	chain := NewChainSigner(signer1, signer2)

	req, _ := http.NewRequest(http.MethodGet, "https://api.example.com/test", nil)
	ctx := context.Background()

	err := chain.Sign(ctx, req)
	if err != nil {
		t.Fatalf("Sign() error = %v", err)
	}

	if signer1.callCount != 1 {
		t.Errorf("Signer1 callCount = %d, want 1", signer1.callCount)
	}

	if signer2.callCount != 1 {
		t.Errorf("Signer2 callCount = %d, want 1", signer2.callCount)
	}
}

func TestChainSigner_SignError(t *testing.T) {
	expectedError := errors.New("signing error")

	signer1 := &mockSigner{
		signFunc: func(ctx context.Context, req *http.Request) error {
			return expectedError
		},
	}
	signer2 := &mockSigner{}

	chain := NewChainSigner(signer1, signer2)

	req, _ := http.NewRequest(http.MethodGet, "https://api.example.com/test", nil)
	ctx := context.Background()

	err := chain.Sign(ctx, req)
	if err == nil {
		t.Fatal("Sign() error = nil, want error")
	}

	if !errors.Is(err, expectedError) {
		t.Errorf("Sign() error = %v, want %v", err, expectedError)
	}

	// 第二个签名器不应该被调用
	if signer2.callCount != 0 {
		t.Errorf("Signer2 callCount = %d, want 0 (should stop on error)", signer2.callCount)
	}
}

func TestChainSigner_Add(t *testing.T) {
	chain := NewChainSigner()

	if len(chain.signers) != 0 {
		t.Errorf("Initial signers count = %d, want 0", len(chain.signers))
	}

	signer1 := &mockSigner{}
	chain.Add(signer1)

	if len(chain.signers) != 1 {
		t.Errorf("After Add(), signers count = %d, want 1", len(chain.signers))
	}

	signer2 := &mockSigner{}
	chain.Add(signer2)

	if len(chain.signers) != 2 {
		t.Errorf("After second Add(), signers count = %d, want 2", len(chain.signers))
	}
}

func TestChainSigner_EmptyChain(t *testing.T) {
	chain := NewChainSigner()

	req, _ := http.NewRequest(http.MethodGet, "https://api.example.com/test", nil)
	ctx := context.Background()

	err := chain.Sign(ctx, req)
	if err != nil {
		t.Fatalf("Sign() with empty chain should not error, got %v", err)
	}
}
