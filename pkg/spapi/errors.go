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

package spapi

import (
	"errors"
	"fmt"
)

// 配置错误。
var (
	// ErrInvalidRegion 表示无效的区域配置。
	ErrInvalidRegion = errors.New("invalid region")

	// ErrMissingClientID 表示缺少 LWA 客户端 ID。
	ErrMissingClientID = errors.New("missing LWA client ID")

	// ErrMissingClientSecret 表示缺少 LWA 客户端密钥。
	ErrMissingClientSecret = errors.New("missing LWA client secret")

	// ErrMissingCredentials 表示缺少认证凭证（RefreshToken 或 Scopes）。
	ErrMissingCredentials = errors.New("missing credentials: either refresh token or scopes required")

	// ErrInvalidTimeout 表示无效的超时配置。
	ErrInvalidTimeout = errors.New("invalid timeout: must be greater than zero")

	// ErrInvalidMaxRetries 表示无效的重试次数配置。
	ErrInvalidMaxRetries = errors.New("invalid max retries: must be non-negative")

	// ErrInvalidRateLimitBuffer 表示无效的速率限制缓冲配置。
	ErrInvalidRateLimitBuffer = errors.New("invalid rate limit buffer: must be between 0.0 and 1.0")
)

// 客户端错误。
var (
	// ErrClientNotInitialized 表示客户端未初始化。
	ErrClientNotInitialized = errors.New("client not initialized")

	// ErrInvalidMarketplace 表示无效的市场配置。
	ErrInvalidMarketplace = errors.New("invalid marketplace")

	// ErrAuthenticationFailed 表示认证失败。
	ErrAuthenticationFailed = errors.New("authentication failed")

	// ErrRateLimitExceeded 表示超过速率限制。
	ErrRateLimitExceeded = errors.New("rate limit exceeded")

	// ErrRequestTimeout 表示请求超时。
	ErrRequestTimeout = errors.New("request timeout")

	// ErrContextCanceled 表示上下文被取消。
	ErrContextCanceled = errors.New("context canceled")
)

// API 请求错误。
var (
	// ErrInvalidRequest 表示无效的请求参数。
	ErrInvalidRequest = errors.New("invalid request")

	// ErrResourceNotFound 表示资源未找到（HTTP 404）。
	ErrResourceNotFound = errors.New("resource not found")

	// ErrUnauthorized 表示未授权（HTTP 401）。
	ErrUnauthorized = errors.New("unauthorized")

	// ErrForbidden 表示禁止访问（HTTP 403）。
	ErrForbidden = errors.New("forbidden")

	// ErrServerError 表示服务器端错误（HTTP 5xx）。
	ErrServerError = errors.New("server error")

	// ErrServiceUnavailable 表示服务不可用（HTTP 503）。
	ErrServiceUnavailable = errors.New("service unavailable")
)

// APIError 表示 SP-API 返回的错误。
//
// 此类型封装了 SP-API 的错误响应，包含：
//   - HTTP 状态码
//   - 错误代码（如 "InvalidInput"）
//   - 错误消息
//   - 额外的错误详情
type APIError struct {
	// StatusCode 是 HTTP 状态码
	StatusCode int

	// Code 是 SP-API 错误代码
	Code string

	// Message 是错误消息
	Message string

	// Details 是额外的错误详情
	Details string
}

// Error 实现 error 接口。
func (e *APIError) Error() string {
	if e.Code != "" {
		return fmt.Sprintf("SP-API error (HTTP %d): %s - %s", e.StatusCode, e.Code, e.Message)
	}
	return fmt.Sprintf("SP-API error (HTTP %d): %s", e.StatusCode, e.Message)
}

// IsRetryable 判断错误是否可以重试。
//
// 返回值:
//   - bool: 如果错误可以重试，返回 true
func (e *APIError) IsRetryable() bool {
	// 5xx 错误和 429 (Rate Limit) 错误可以重试
	return e.StatusCode >= 500 || e.StatusCode == 429
}
