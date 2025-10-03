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
	"fmt"
	"net/http"
	"strings"
)

// RDTSigner 实现 RDT (Restricted Data Token) 签名。
//
// RDT 用于访问包含个人身份信息 (PII) 的受限数据。
// 某些 SP-API 端点要求使用 RDT 进行授权。
type RDTSigner struct {
	// rdtProvider 是 RDT 提供者函数。
	// 它根据请求的资源路径和操作返回相应的 RDT。
	rdtProvider RDTProvider
}

// RDTProvider 定义 RDT 提供者函数签名。
//
// 参数:
//   - ctx: 请求上下文
//   - resourcePath: 资源路径（例如: "/orders/v0/orders/{orderId}"）
//   - dataElements: 需要访问的数据元素列表
//
// 返回值:
//   - string: RDT 令牌
//   - error: 如果获取 RDT 失败，返回错误
type RDTProvider func(ctx context.Context, resourcePath string, dataElements []string) (string, error)

// NewRDTSigner 创建新的 RDT 签名器。
//
// 参数:
//   - provider: RDT 提供者函数
//
// 返回值:
//   - *RDTSigner: RDT 签名器实例
//
// 示例:
//
//	provider := func(ctx context.Context, path string, elements []string) (string, error) {
//	    // 调用 Tokens API 获取 RDT
//	    return tokensAPI.CreateRestrictedDataToken(ctx, path, elements)
//	}
//	signer := signer.NewRDTSigner(provider)
func NewRDTSigner(provider RDTProvider) *RDTSigner {
	return &RDTSigner{
		rdtProvider: provider,
	}
}

// Sign 为请求添加 RDT 授权。
//
// 此方法检查请求是否需要 RDT，如果需要，
// 则获取 RDT 并将其添加到 x-amz-access-token 头中。
//
// 参数:
//   - ctx: 请求上下文
//   - req: 需要签名的 HTTP 请求
//
// 返回值:
//   - error: 如果获取 RDT 或添加头失败，返回错误
func (s *RDTSigner) Sign(ctx context.Context, req *http.Request) error {
	// 如果没有配置 RDT 提供者，跳过签名
	if s.rdtProvider == nil {
		return nil
	}

	// 检查请求是否需要 RDT
	if !s.requiresRDT(req) {
		return nil
	}

	// 提取资源路径和数据元素
	resourcePath := req.URL.Path
	dataElements := s.extractDataElements(req)

	// 获取 RDT
	rdt, err := s.rdtProvider(ctx, resourcePath, dataElements)
	if err != nil {
		return fmt.Errorf("get RDT: %w", err)
	}

	// 使用 RDT 替换现有的访问令牌
	req.Header.Set("x-amz-access-token", rdt)

	return nil
}

// requiresRDT 检查请求是否需要 RDT。
//
// 以下情况需要 RDT：
// 1. 请求头中包含 x-amzn-RDT-Required 标记
// 2. 请求路径匹配受限 API 端点
func (s *RDTSigner) requiresRDT(req *http.Request) bool {
	// 检查是否有 RDT 要求标记
	if req.Header.Get("x-amzn-RDT-Required") == "true" {
		return true
	}

	// 检查路径是否属于受限端点
	path := req.URL.Path

	// 受限端点列表（根据官方文档）
	restrictedPaths := []string{
		"/orders/v0/orders/",                    // Get Order - 包含 PII
		"/orders/v0/orders/{orderId}/address",   // Get Order Address
		"/orders/v0/orders/{orderId}/buyerInfo", // Get Order Buyer Info
		"/mfn/v0/shipments/",                    // 部分 MFN API
		"/messaging/v1/orders/",                 // Messaging API
	}

	for _, restrictedPath := range restrictedPaths {
		if strings.Contains(path, restrictedPath) {
			return true
		}
	}

	return false
}

// extractDataElements 从请求中提取需要访问的数据元素。
//
// 这通常从请求头、查询参数或其他元数据中提取。
func (s *RDTSigner) extractDataElements(req *http.Request) []string {
	// 从请求头中提取数据元素
	elementsHeader := req.Header.Get("x-amzn-RDT-DataElements")
	if elementsHeader != "" {
		return strings.Split(elementsHeader, ",")
	}

	// 根据路径推断数据元素
	path := req.URL.Path

	if strings.Contains(path, "/orders/") {
		if strings.Contains(path, "/address") {
			return []string{"buyerInfo", "shippingAddress"}
		}
		if strings.Contains(path, "/buyerInfo") {
			return []string{"buyerInfo"}
		}
		// 默认订单数据元素
		return []string{"buyerInfo", "shippingAddress"}
	}

	// 默认返回空列表
	return []string{}
}

// SetRDTProvider 设置 RDT 提供者。
//
// 这允许在运行时更换 RDT 提供者。
func (s *RDTSigner) SetRDTProvider(provider RDTProvider) {
	s.rdtProvider = provider
}
