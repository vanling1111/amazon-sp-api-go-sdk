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

package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Client 是 LWA 认证客户端。
//
// 此客户端负责与 LWA 服务器通信，获取和缓存访问令牌。
// Client 是并发安全的，可以在多个 goroutine 中使用。
type Client struct {
	credentials *Credentials
	httpClient  *http.Client
	cache       TokenCache
}

// TokenCache 定义令牌缓存接口。
//
// 实现此接口可以自定义令牌缓存策略。
type TokenCache interface {
	// Get 获取缓存的令牌。
	// 如果令牌不存在或已过期，返回 nil 和 false。
	Get(key string) (*Token, bool)

	// Set 设置令牌到缓存。
	Set(key string, token *Token)

	// Delete 删除缓存的令牌。
	Delete(key string)
}

// Token 表示 LWA 访问令牌。
type Token struct {
	// AccessToken 是访问令牌。
	AccessToken string `json:"access_token"`

	// TokenType 是令牌类型（通常是 "bearer"）。
	TokenType string `json:"token_type"`

	// ExpiresIn 是令牌的有效期（秒）。
	ExpiresIn int `json:"expires_in"`

	// ExpiresAt 是令牌的过期时间。
	ExpiresAt time.Time `json:"-"`
}

// IsExpired 检查令牌是否已过期。
//
// 为了安全起见，会提前 60 秒判定为过期。
func (t *Token) IsExpired() bool {
	if t == nil {
		return true
	}
	return time.Now().Add(60 * time.Second).After(t.ExpiresAt)
}

// lwaResponse 表示 LWA 服务器的响应。
type lwaResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Error       string `json:"error,omitempty"`
	ErrorDesc   string `json:"error_description,omitempty"`
}

// 错误定义。
var (
	// ErrAuthFailed 表示认证失败。
	ErrAuthFailed = errors.New("authentication failed")

	// ErrInvalidResponse 表示 LWA 服务器返回了无效的响应。
	ErrInvalidResponse = errors.New("invalid LWA response")
)

// NewClient 创建新的 LWA 客户端。
//
// 参数:
//   - credentials: LWA 认证凭据
//
// 返回值:
//   - *Client: LWA 客户端实例
//
// 示例:
//
//	creds, _ := auth.NewCredentials(...)
//	client := auth.NewClient(creds)
func NewClient(credentials *Credentials) *Client {
	return &Client{
		credentials: credentials,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		cache: NewMemoryCache(),
	}
}

// SetHTTPClient 设置自定义 HTTP 客户端。
//
// 这允许您自定义超时、代理等设置。
func (c *Client) SetHTTPClient(httpClient *http.Client) {
	c.httpClient = httpClient
}

// SetCache 设置自定义令牌缓存。
//
// 这允许您使用自己的缓存实现（如 Redis）。
func (c *Client) SetCache(cache TokenCache) {
	c.cache = cache
}

// GetAccessToken 获取访问令牌。
//
// 此方法首先检查缓存，如果缓存中没有有效的令牌，
// 则从 LWA 服务器获取新令牌并缓存。
//
// 参数:
//   - ctx: 请求上下文
//
// 返回值:
//   - string: 访问令牌
//   - error: 如果获取失败，返回错误
//
// 示例:
//
//	ctx := context.Background()
//	token, err := client.GetAccessToken(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println("Access Token:", token)
func (c *Client) GetAccessToken(ctx context.Context) (string, error) {
	// 生成缓存键
	cacheKey := c.getCacheKey()

	// 检查缓存
	if cachedToken, ok := c.cache.Get(cacheKey); ok {
		if !cachedToken.IsExpired() {
			return cachedToken.AccessToken, nil
		}
		// 令牌已过期，删除缓存
		c.cache.Delete(cacheKey)
	}

	// 从 LWA 服务器获取新令牌
	token, err := c.fetchToken(ctx)
	if err != nil {
		return "", err
	}

	// 缓存令牌
	c.cache.Set(cacheKey, token)

	return token.AccessToken, nil
}

// fetchToken 从 LWA 服务器获取新令牌。
func (c *Client) fetchToken(ctx context.Context) (*Token, error) {
	// 构建请求参数
	data := c.buildTokenRequest()

	// 创建 HTTP 请求
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.credentials.Endpoint,
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 发送请求
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	// 解析响应
	var lwaResp lwaResponse
	if err := json.Unmarshal(body, &lwaResp); err != nil {
		return nil, fmt.Errorf("parse response: %w", err)
	}

	// 检查错误
	if lwaResp.Error != "" {
		return nil, fmt.Errorf("%w: %s - %s", ErrAuthFailed, lwaResp.Error, lwaResp.ErrorDesc)
	}

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: HTTP %d: %s", ErrAuthFailed, resp.StatusCode, string(body))
	}

	// 验证响应
	if lwaResp.AccessToken == "" {
		return nil, fmt.Errorf("%w: empty access token", ErrInvalidResponse)
	}

	// 构建令牌对象
	token := &Token{
		AccessToken: lwaResp.AccessToken,
		TokenType:   lwaResp.TokenType,
		ExpiresIn:   lwaResp.ExpiresIn,
		ExpiresAt:   time.Now().Add(time.Duration(lwaResp.ExpiresIn) * time.Second),
	}

	return token, nil
}

// buildTokenRequest 构建 LWA 令牌请求参数。
//
// 根据凭据类型（regular 或 grantless）选择不同的请求参数。
func (c *Client) buildTokenRequest() url.Values {
	data := url.Values{
		"client_id":     {c.credentials.ClientID},
		"client_secret": {c.credentials.ClientSecret},
	}

	if c.credentials.IsGrantless() {
		// Grantless operation: 使用 client_credentials grant type
		data.Set("grant_type", "client_credentials")
		// 将 scopes 用空格分隔连接
		data.Set("scope", strings.Join(c.credentials.Scopes, " "))
	} else {
		// Regular operation: 使用 refresh_token grant type
		data.Set("grant_type", "refresh_token")
		data.Set("refresh_token", c.credentials.RefreshToken)
	}

	return data
}

// getCacheKey 生成缓存键。
//
// 对于 regular operations，使用 client_id 和 refresh_token 的组合。
// 对于 grantless operations，使用 client_id 和 scopes 的组合。
func (c *Client) getCacheKey() string {
	if c.credentials.IsGrantless() {
		// Grantless: 使用 client_id 和 scopes 作为缓存键
		scopesKey := strings.Join(c.credentials.Scopes, ",")
		return fmt.Sprintf("lwa_token:%s:grantless:%s", c.credentials.ClientID, scopesKey)
	}

	// Regular: 使用 client_id 和 refresh_token 作为缓存键
	return fmt.Sprintf("lwa_token:%s:regular:%s", c.credentials.ClientID, c.credentials.RefreshToken)
}

// RefreshToken 强制刷新令牌。
//
// 此方法会忽略缓存，直接从 LWA 服务器获取新令牌。
//
// 参数:
//   - ctx: 请求上下文
//
// 返回值:
//   - string: 新的访问令牌
//   - error: 如果刷新失败，返回错误
func (c *Client) RefreshToken(ctx context.Context) (string, error) {
	// 删除缓存
	cacheKey := c.getCacheKey()
	c.cache.Delete(cacheKey)

	// 获取新令牌
	return c.GetAccessToken(ctx)
}
