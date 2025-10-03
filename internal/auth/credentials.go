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
// Package auth 提供 LWA (Login with Amazon) 认证功能。
//
// 此包实现了 Amazon SP-API 所需的 OAuth 2.0 认证流程，
// 包括访问令牌的获取、缓存和刷新。
//
// 参考文档:
//   - https://developer-docs.amazon.com/sp-api/docs/connecting-to-the-selling-partner-api
package auth

import (
	"errors"
	"fmt"
)

// Credentials 表示 LWA 认证凭据。
//
// 这些凭据用于从 LWA 服务器获取访问令牌。
//
// 支持两种认证模式:
//   - Regular operations: 需要 RefreshToken（grant_type=refresh_token）
//   - Grantless operations: 需要 Scopes（grant_type=client_credentials）
type Credentials struct {
	// ClientID 是您的 LWA 应用程序的客户端 ID。
	// 必需。
	ClientID string

	// ClientSecret 是您的 LWA 应用程序的客户端密钥。
	// 必需。
	ClientSecret string

	// RefreshToken 是用于获取访问令牌的刷新令牌。
	// 用于常规操作（需要卖家授权）。
	// 如果设置了 Scopes，则此字段可选。
	RefreshToken string

	// Scopes 是 grantless operations 的授权范围。
	// 用于不需要卖家授权的操作，如 Notifications API。
	// 如果设置了 RefreshToken，则此字段可选。
	//
	// 可用的 scopes:
	//   - "sellingpartnerapi::notifications" - Notifications API
	//   - "sellingpartnerapi::client_credential:rotation" - Application Management API
	Scopes []string

	// Endpoint 是 LWA 授权服务器的端点 URL。
	// 必需。
	// 例如: "https://api.amazon.com/auth/o2/token"
	Endpoint string
}

// 预定义的 LWA 端点。
const (
	// EndpointNA 是北美区域的 LWA 端点。
	EndpointNA = "https://api.amazon.com/auth/o2/token"

	// EndpointEU 是欧洲区域的 LWA 端点。
	EndpointEU = "https://api.amazon.co.uk/auth/o2/token"

	// EndpointFE 是远东区域的 LWA 端点。
	EndpointFE = "https://api.amazon.co.jp/auth/o2/token"
)

// 预定义的 Scopes（用于 grantless operations）。
const (
	// ScopeNotifications 用于 Notifications API。
	ScopeNotifications = "sellingpartnerapi::notifications"

	// ScopeCredentialRotation 用于 Application Management API。
	ScopeCredentialRotation = "sellingpartnerapi::client_credential:rotation"
)

// 错误定义。
var (
	// ErrInvalidCredentials 表示凭据无效。
	ErrInvalidCredentials = errors.New("invalid credentials")

	// ErrMissingClientID 表示缺少客户端 ID。
	ErrMissingClientID = errors.New("client ID is required")

	// ErrMissingClientSecret 表示缺少客户端密钥。
	ErrMissingClientSecret = errors.New("client secret is required")

	// ErrMissingRefreshToken 表示缺少刷新令牌。
	ErrMissingRefreshToken = errors.New("refresh token is required for regular operations")

	// ErrMissingScopes 表示缺少 scopes。
	ErrMissingScopes = errors.New("scopes are required for grantless operations")

	// ErrMissingEndpoint 表示缺少端点 URL。
	ErrMissingEndpoint = errors.New("endpoint URL is required")

	// ErrBothRefreshTokenAndScopes 表示同时设置了 RefreshToken 和 Scopes。
	ErrBothRefreshTokenAndScopes = errors.New("cannot specify both refresh token and scopes")
)

// NewCredentials 创建新的认证凭据。
//
// 所有参数都是必需的。如果任何参数为空，将返回错误。
//
// 参数:
//   - clientID: LWA 应用程序的客户端 ID
//   - clientSecret: LWA 应用程序的客户端密钥
//   - refreshToken: 刷新令牌
//   - endpoint: LWA 端点 URL（使用预定义常量或自定义 URL）
//
// 返回值:
//   - *Credentials: 认证凭据实例
//   - error: 如果参数无效，返回错误
//
// 示例:
//
//	creds, err := auth.NewCredentials(
//	    "your-client-id",
//	    "your-client-secret",
//	    "your-refresh-token",
//	    auth.EndpointNA,
//	)
//	if err != nil {
//	    log.Fatal(err)
//	}
func NewCredentials(clientID, clientSecret, refreshToken, endpoint string) (*Credentials, error) {
	creds := &Credentials{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RefreshToken: refreshToken,
		Endpoint:     endpoint,
	}

	if err := creds.Validate(); err != nil {
		return nil, err
	}

	return creds, nil
}

// NewGrantlessCredentials 创建用于 grantless operations 的认证凭据（grant_type=client_credentials）。
//
// 此函数用于创建不需要卖家授权的操作的凭据，如 Notifications API。
// 所有参数都是必需的。如果任何参数为空，将返回错误。
//
// 参数:
//   - clientID: LWA 应用程序的客户端 ID
//   - clientSecret: LWA 应用程序的客户端密钥
//   - scopes: 授权范围（使用预定义的 Scope* 常量）
//   - endpoint: LWA 端点 URL（使用预定义常量或自定义 URL）
//
// 返回值:
//   - *Credentials: 认证凭据实例
//   - error: 如果参数无效，返回错误
//
// 示例:
//
//	// Notifications API
//	creds, err := auth.NewGrantlessCredentials(
//	    "your-client-id",
//	    "your-client-secret",
//	    []string{auth.ScopeNotifications},
//	    auth.EndpointNA,
//	)
//	if err != nil {
//	    log.Fatal(err)
//	}
func NewGrantlessCredentials(clientID, clientSecret string, scopes []string, endpoint string) (*Credentials, error) {
	creds := &Credentials{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       scopes,
		Endpoint:     endpoint,
	}

	if err := creds.Validate(); err != nil {
		return nil, err
	}

	return creds, nil
}

// IsGrantless 返回凭据是否用于 grantless operations。
//
// 返回值:
//   - bool: 如果是 grantless 凭据返回 true，否则返回 false
func (c *Credentials) IsGrantless() bool {
	return len(c.Scopes) > 0
}

// Validate 验证凭据是否有效。
//
// 检查所有必需字段是否已设置。
//
// 验证规则:
//   - ClientID 和 ClientSecret 始终必需
//   - Endpoint 始终必需
//   - RefreshToken 或 Scopes 二选一（不能同时为空，也不能同时存在）
//
// 返回值:
//   - error: 如果凭据无效，返回具体的错误信息
func (c *Credentials) Validate() error {
	if c.ClientID == "" {
		return ErrMissingClientID
	}

	if c.ClientSecret == "" {
		return ErrMissingClientSecret
	}

	if c.Endpoint == "" {
		return ErrMissingEndpoint
	}

	// RefreshToken 和 Scopes 必须且只能有一个
	hasRefreshToken := c.RefreshToken != ""
	hasScopes := len(c.Scopes) > 0

	if !hasRefreshToken && !hasScopes {
		return ErrInvalidCredentials
	}

	if hasRefreshToken && hasScopes {
		return ErrBothRefreshTokenAndScopes
	}

	return nil
}

// String 返回凭据的字符串表示（隐藏敏感信息）。
//
// 此方法用于日志记录和调试，不会暴露敏感信息。
func (c *Credentials) String() string {
	return fmt.Sprintf("Credentials{ClientID: %s, Endpoint: %s}",
		maskString(c.ClientID), c.Endpoint)
}

// maskString 遮盖敏感字符串，只显示前后几个字符。
func maskString(s string) string {
	if len(s) <= 8 {
		return "****"
	}
	return s[:4] + "****" + s[len(s)-4:]
}
