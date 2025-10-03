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

// Package signer 提供请求签名功能。
//
// 此包实现了 Amazon SP-API 所需的各种签名机制：
// - LWA (Login with Amazon) 授权签名
// - RDT (Restricted Data Token) 签名
// - AWS Signature Version 4 签名
//
// 基于官方 SP-API 文档:
//   - https://developer-docs.amazon.com/sp-api/docs/connecting-to-the-selling-partner-api
package signer

import (
	"context"
	"net/http"
)

// Signer 定义签名器接口。
//
// 签名器负责为 HTTP 请求添加必要的认证信息。
// 所有签名器实现都应该是并发安全的。
type Signer interface {
	// Sign 为请求签名。
	//
	// 此方法会修改传入的请求，添加必要的认证头或参数。
	//
	// 参数:
	//   - ctx: 请求上下文
	//   - req: 需要签名的 HTTP 请求
	//
	// 返回值:
	//   - error: 如果签名失败，返回错误
	Sign(ctx context.Context, req *http.Request) error
}

// ChainSigner 是签名器链，按顺序执行多个签名器。
//
// 这允许组合多个签名机制，例如先添加 LWA 令牌，
// 然后再添加 AWS Signature V4。
type ChainSigner struct {
	signers []Signer
}

// NewChainSigner 创建新的签名器链。
//
// 参数:
//   - signers: 签名器列表，按顺序执行
//
// 返回值:
//   - *ChainSigner: 签名器链实例
//
// 示例:
//   chain := signer.NewChainSigner(
//       lwaSigner,
//       rdtSigner,
//   )
func NewChainSigner(signers ...Signer) *ChainSigner {
	return &ChainSigner{
		signers: signers,
	}
}

// Sign 按顺序执行所有签名器。
//
// 如果任何签名器返回错误，立即停止并返回该错误。
func (c *ChainSigner) Sign(ctx context.Context, req *http.Request) error {
	for _, signer := range c.signers {
		if err := signer.Sign(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// Add 添加签名器到链中。
func (c *ChainSigner) Add(signer Signer) {
	c.signers = append(c.signers, signer)
}
