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

package signer

import (
	"context"
	"fmt"
	"net/http"

	"github.com/vanling1111/amazon-sp-api-go-sdk/internal/auth"
)

// LWASigner 实现 LWA (Login with Amazon) 授权签名。
//
// 此签名器从 LWA 客户端获取访问令牌，
// 并将其添加到请求的 Authorization 头中。
type LWASigner struct {
	lwaClient *auth.Client
}

// NewLWASigner 创建新的 LWA 签名器。
//
// 参数:
//   - lwaClient: LWA 认证客户端
//
// 返回值:
//   - *LWASigner: LWA 签名器实例
//
// 示例:
//   creds, _ := auth.NewCredentials(...)
//   lwaClient := auth.NewClient(creds)
//   signer := signer.NewLWASigner(lwaClient)
func NewLWASigner(lwaClient *auth.Client) *LWASigner {
	return &LWASigner{
		lwaClient: lwaClient,
	}
}

// Sign 为请求添加 LWA 授权头。
//
// 此方法会获取 LWA 访问令牌，并将其添加到
// x-amz-access-token 头中。
//
// 根据官方 SP-API 文档:
//   - 常规操作使用 LWA 访问令牌
//   - 受限操作使用 RDT（由 RDT Signer 处理）
//
// 参数:
//   - ctx: 请求上下文
//   - req: 需要签名的 HTTP 请求
//
// 返回值:
//   - error: 如果获取令牌或添加头失败，返回错误
func (s *LWASigner) Sign(ctx context.Context, req *http.Request) error {
	// 获取访问令牌
	accessToken, err := s.lwaClient.GetAccessToken(ctx)
	if err != nil {
		return fmt.Errorf("get LWA access token: %w", err)
	}

	// 添加 x-amz-access-token 头
	req.Header.Set("x-amz-access-token", accessToken)

	return nil
}

// SetLWAClient 设置 LWA 客户端。
//
// 这允许在运行时更换 LWA 客户端。
func (s *LWASigner) SetLWAClient(client *auth.Client) {
	s.lwaClient = client
}
