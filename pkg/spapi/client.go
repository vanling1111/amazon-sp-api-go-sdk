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

package spapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/vanling1111/amazon-sp-api-go-sdk/internal/auth"
	"github.com/vanling1111/amazon-sp-api-go-sdk/internal/ratelimit"
	"github.com/vanling1111/amazon-sp-api-go-sdk/internal/signer"
	"github.com/vanling1111/amazon-sp-api-go-sdk/internal/transport"
)

// Client 是 Amazon SP-API 的主客户端。
//
// Client 封装了所有必要的组件（认证、传输、签名、速率限制），
// 提供简洁、易用的接口来访问各个 SP-API。
//
// Client 是并发安全的，可以在多个 goroutine 中复用。
//
// 基于官方 SP-API 文档:
//   - https://developer-docs.amazon.com/sp-api/docs/
//
// 示例:
//
//	client, err := spapi.NewClient(
//	    spapi.WithRegion(models.RegionNorthAmerica),
//	    spapi.WithCredentials("client-id", "client-secret", "refresh-token"),
//	)
//	if err != nil {
//	    log.Fatal(err)
//	}
type Client struct {
	// config 是客户端配置
	config *Config

	// lwaClient 是 LWA 认证客户端
	lwaClient *auth.Client

	// httpClient 是 HTTP 传输客户端
	httpClient *transport.Client

	// signer 是请求签名器
	signer signer.Signer

	// rateLimitManager 是速率限制管理器
	rateLimitManager *ratelimit.Manager
}

// NewClient 创建新的 SP-API 客户端。
//
// 参数:
//   - opts: 客户端配置选项（使用 Functional Options 模式）
//
// 返回值:
//   - *Client: SP-API 客户端实例
//   - error: 如果配置无效或初始化失败，返回错误
//
// 示例:
//
//	// Regular 操作（使用 refresh token）
//	client, err := spapi.NewClient(
//	    spapi.WithRegion(models.RegionNorthAmerica),
//	    spapi.WithCredentials("client-id", "client-secret", "refresh-token"),
//	)
//
//	// Grantless 操作
//	client, err := spapi.NewClient(
//	    spapi.WithRegion(models.RegionEurope),
//	    spapi.WithGrantlessCredentials("client-id", "client-secret", []string{
//	        "sellingpartnerapi::notifications",
//	    }),
//	)
//
//	// 自定义配置
//	client, err := spapi.NewClient(
//	    spapi.WithRegion(models.RegionNorthAmerica),
//	    spapi.WithCredentials("client-id", "client-secret", "refresh-token"),
//	    spapi.WithHTTPTimeout(60 * time.Second),
//	    spapi.WithMaxRetries(5),
//	    spapi.WithDebug(),
//	)
//
// 官方文档:
//   - https://developer-docs.amazon.com/sp-api/docs/connecting-to-the-selling-partner-api
func NewClient(opts ...ClientOption) (*Client, error) {
	// 1. 创建默认配置
	config := DefaultConfig()

	// 2. 应用用户提供的选项
	for _, opt := range opts {
		opt(config)
	}

	// 3. 验证配置
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	// 4. 创建 LWA 认证客户端
	var lwaCredentials *auth.Credentials
	var err error

	if len(config.Scopes) > 0 {
		// Grantless 操作
		lwaCredentials, err = auth.NewGrantlessCredentials(
			config.ClientID,
			config.ClientSecret,
			config.Scopes,
			config.Region.LWAEndpoint,
		)
	} else {
		// Regular 操作
		lwaCredentials, err = auth.NewCredentials(
			config.ClientID,
			config.ClientSecret,
			config.RefreshToken,
			config.Region.LWAEndpoint,
		)
	}

	if err != nil {
		return nil, fmt.Errorf("create LWA credentials: %w", err)
	}

	lwaClient := auth.NewClient(lwaCredentials)

	// 5. 创建 HTTP 传输客户端
	transportConfig := &transport.Config{
		Timeout:             config.HTTPTimeout,
		MaxIdleConns:        200, // 生产级连接池配置
		MaxIdleConnsPerHost: 20,
		MaxConnsPerHost:     50,
		IdleConnTimeout:     90 * config.HTTPTimeout,
		UserAgent:           "amazon-sp-api-go-sdk/1.0.0",
		Debug:               config.Debug,
	}

	httpClient := transport.NewClient(config.Region.Endpoint, transportConfig)

	// 6. 设置 Metrics 记录器（如果提供）
	httpClient.SetMetrics(config.MetricsRecorder)

	// 7. 添加标准中间件
	httpClient.Use(transport.UserAgentMiddleware(transportConfig.UserAgent))
	httpClient.Use(transport.DateMiddleware()) // 添加 x-amz-date 头部（官方要求）
	httpClient.Use(transport.RequestIDMiddleware())

	// 8. 添加重试中间件（官方建议的 back-off strategy）
	if config.MaxRetries > 0 {
		retryConfig := &transport.RetryConfig{
			MaxRetries:      config.MaxRetries,
			InitialInterval: 100,   // 100ms
			MaxInterval:     30000, // 30s
			Multiplier:      2.0,
			ShouldRetry:     nil, // 使用默认的重试判断函数
		}
		httpClient.Use(transport.RetryMiddleware(retryConfig))
	}

	// 9. 创建签名器（LWA 签名器）
	lwaSigner := signer.NewLWASigner(lwaClient)

	// 10. 创建速率限制管理器
	// 官方文档建议：读取 x-amzn-RateLimit-Limit 头部，不要硬编码
	rateLimitManager := ratelimit.NewManager(
		ratelimit.WithDefaultRate(1.0, 5), // 保守的默认值
	)

	// 11. 构建客户端
	client := &Client{
		config:           config,
		lwaClient:        lwaClient,
		httpClient:       httpClient,
		signer:           lwaSigner,
		rateLimitManager: rateLimitManager,
	}

	return client, nil
}

// Config 返回客户端的配置副本。
//
// 返回值:
//   - *Config: 配置副本
func (c *Client) Config() *Config {
	// 返回副本以防止外部修改
	configCopy := *c.config
	return &configCopy
}

// Close 关闭客户端并释放资源。
//
// 注意：调用 Close 后，客户端将不可用。
//
// 示例:
//
//	client, err := spapi.NewClient(...)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer client.Close()
func (c *Client) Close() error {
	// 目前无需特殊清理
	// 未来如果添加后台任务，在此处停止
	return nil
}

// GetAccessToken 获取当前的 LWA 访问令牌。
//
// 此方法主要用于调试和测试。通常情况下，SDK 会自动处理令牌管理，
// 用户不需要直接调用此方法。
//
// 参数:
//   - ctx: 请求上下文
//
// 返回值:
//   - string: LWA 访问令牌
//   - error: 如果获取失败，返回错误
//
// 示例:
//
//	token, err := client.GetAccessToken(context.Background())
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println("Access Token:", token)
func (c *Client) GetAccessToken(ctx context.Context) (string, error) {
	token, err := c.lwaClient.GetAccessToken(ctx)
	if err != nil {
		return "", fmt.Errorf("get access token: %w", err)
	}
	return token, nil
}

// RateLimitManager 返回速率限制管理器。
//
// 此方法允许高级用户直接访问速率限制管理器，
// 用于自定义速率限制策略或监控速率限制状态。
//
// 返回值:
//   - *ratelimit.Manager: 速率限制管理器
//
// 示例:
//
//	manager := client.RateLimitManager()
//	count := manager.Count()
//	fmt.Printf("Active limiters: %d\n", count)
func (c *Client) RateLimitManager() *ratelimit.Manager {
	return c.rateLimitManager
}

// HTTPClient 返回底层的 HTTP 传输客户端。
//
// 此方法允许高级用户直接访问 HTTP 客户端，
// 用于添加自定义中间件或修改配置。
//
// 返回值:
//   - *transport.Client: HTTP 传输客户端
//
// 示例:
//
//	httpClient := client.HTTPClient()
//	httpClient.Use(myCustomMiddleware)
func (c *Client) HTTPClient() *transport.Client {
	return c.httpClient
}

// Signer 返回请求签名器。
//
// 此方法主要用于测试和调试。
//
// 返回值:
//   - signer.Signer: 请求签名器
func (c *Client) Signer() signer.Signer {
	return c.signer
}

// ==================== 通用 HTTP 请求方法 ====================
//
// 以下方法提供了通用的 HTTP 请求能力，供各个 API 客户端使用。
// 每个 API 客户端只需要轻量级包装这些方法，大大减少重复代码。

// DoRequest 执行一个通用的 HTTP 请求。
//
// 此方法是所有 API 请求的基础，提供：
//  - 自动 LWA 认证
//  - 速率限制检查
//  - 请求签名
//  - 错误处理
//  - 响应解析
//
// 参数:
//   - ctx: 请求上下文
//   - method: HTTP 方法（GET、POST、PUT、DELETE等）
//   - path: API 路径（相对于区域端点，如 "/orders/v0/orders"）
//   - query: 查询参数（可选，传 nil 表示无查询参数）
//   - body: 请求体（可选，传 nil 表示无请求体）
//   - result: 响应结果的指针（将被 JSON 解码填充，传 nil 表示不解析响应）
//
// 返回值:
//   - error: 如果请求失败，返回错误
//
// 示例:
//
//	var response orders.GetOrdersResponse
//	err := client.DoRequest(ctx, "GET", "/orders/v0/orders", map[string]string{
//	    "MarketplaceIds": "ATVPDKIKX0DER",
//	    "CreatedAfter": "2023-01-01T00:00:00Z",
//	}, nil, &response)
func (c *Client) DoRequest(ctx context.Context, method, path string, query map[string]string, body, result interface{}) error {
	// 1. 获取 access token
	accessToken, err := c.lwaClient.GetAccessToken(ctx)
	if err != nil {
		return fmt.Errorf("failed to get access token: %w", err)
	}

	// 2. 速率限制检查
	if c.rateLimitManager != nil {
		sellerID := c.extractSellerID()
		appID := c.config.ClientID
		marketplace := c.extractMarketplaceID(query)
		operation := c.extractOperationName(method, path)

		allowed := c.rateLimitManager.Allow(sellerID, appID, marketplace, operation)
		if !allowed {
			return ErrRateLimitExceeded
		}
	}

	// 3. 构建请求
	req, err := c.buildRequest(ctx, method, path, query, body, accessToken)
	if err != nil {
		return fmt.Errorf("failed to build request: %w", err)
	}

	// 4. 签名请求
	if err := c.signer.Sign(ctx, req); err != nil {
		return fmt.Errorf("failed to sign request: %w", err)
	}

	// 5. 发送请求
	resp, err := c.httpClient.Do(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// 6. 处理响应
	if err := c.handleResponse(resp, result); err != nil {
		return err
	}

	// 7. 更新速率限制（从响应头提取）
	if c.rateLimitManager != nil {
		sellerID := c.extractSellerID()
		appID := c.config.ClientID
		marketplace := c.extractMarketplaceID(query)
		operation := c.extractOperationName(method, path)
		c.updateRateLimitFromResponse(resp, sellerID, appID, marketplace, operation)
	}

	return nil
}

// buildRequest 构建 HTTP 请求。
func (c *Client) buildRequest(ctx context.Context, method, path string, query map[string]string, body interface{}, accessToken string) (*http.Request, error) {
	// 构建完整 URL
	fullURL := c.config.Region.Endpoint + path
	if len(query) > 0 {
		queryParams := url.Values{}
		for key, value := range query {
			queryParams.Add(key, value)
		}
		fullURL += "?" + queryParams.Encode()
	}

	// 编码请求体
	var bodyReader io.Reader
	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(bodyBytes)
	}

	// 创建请求
	req, err := http.NewRequestWithContext(ctx, method, fullURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// 设置必需的 headers
	req.Header.Set("x-amz-access-token", accessToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	return req, nil
}

// handleResponse 处理 HTTP 响应。
func (c *Client) handleResponse(resp *http.Response, result interface{}) error {
	// 读取响应体
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// 检查 HTTP 状态码
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return c.handleErrorResponse(resp.StatusCode, bodyBytes)
	}

	// 如果 result 为 nil，不解析响应体
	if result == nil {
		return nil
	}

	// 解析响应体
	if err := json.Unmarshal(bodyBytes, result); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return nil
}

// handleErrorResponse 处理错误响应。
func (c *Client) handleErrorResponse(statusCode int, body []byte) error {
	// 尝试解析为标准错误格式
	var apiError struct {
		Errors []struct {
			Code    string `json:"code"`
			Message string `json:"message"`
			Details string `json:"details"`
		} `json:"errors"`
	}

	if err := json.Unmarshal(body, &apiError); err == nil && len(apiError.Errors) > 0 {
		// 返回第一个错误
		return &APIError{
			StatusCode: statusCode,
			Code:       apiError.Errors[0].Code,
			Message:    apiError.Errors[0].Message,
			Details:    apiError.Errors[0].Details,
		}
	}

	// 如果无法解析为标准错误格式，返回通用错误
	return &APIError{
		StatusCode: statusCode,
		Message:    string(body),
	}
}

// updateRateLimitFromResponse 从响应头更新速率限制。
func (c *Client) updateRateLimitFromResponse(resp *http.Response, sellerID, appID, marketplace, operation string) {
	// 提取 x-amzn-RateLimit-Limit header
	rateLimitHeader := resp.Header.Get("x-amzn-RateLimit-Limit")
	if rateLimitHeader == "" {
		return
	}

	// 解析速率限制值
	rate, err := strconv.ParseFloat(rateLimitHeader, 64)
	if err != nil {
		return
	}

	// 计算 burst（通常是 rate 的 20 倍，这是 SP-API 的默认行为）
	burst := int(rate * 20)
	if burst < 1 {
		burst = 1
	}

	// 更新速率限制
	_ = c.rateLimitManager.UpdateRate(sellerID, appID, marketplace, operation, rate, burst)
}

// extractSellerID 提取 Seller ID。
//
// Seller ID 可以从以下来源获取（优先级从高到低）：
//  1. 配置中的 SellerID
//  2. 从 refresh token 中提取（如果可能）
//  3. 使用 "default" 作为回退值
//
// 返回值:
//   - string: Seller ID
func (c *Client) extractSellerID() string {
	// 优先使用配置中的 SellerID
	if c.config.SellerID != "" {
		return c.config.SellerID
	}

	// 使用 client ID 作为唯一标识
	// 为每个应用提供独立的速率限制跟踪
	return c.config.ClientID
}

// extractMarketplaceID 从查询参数中提取 Marketplace ID。
//
// Marketplace ID 可能出现在以下查询参数中：
//  - MarketplaceIds (大多数 API)
//  - MarketplaceId (部分 API)
//  - marketplace_ids (部分 API)
//
// 如果查询参数中包含多个 Marketplace ID（逗号分隔），
// 则返回第一个。
//
// 参数:
//   - query: 查询参数
//
// 返回值:
//   - string: Marketplace ID，如果未找到则返回 "global"
func (c *Client) extractMarketplaceID(query map[string]string) string {
	// 尝试标准参数名
	if ids := query["MarketplaceIds"]; ids != "" {
		// 如果包含多个 ID（逗号分隔），返回第一个
		if idx := len(ids); idx > 0 {
			for i := 0; i < len(ids); i++ {
				if ids[i] == ',' {
					return ids[:i]
				}
			}
			return ids
		}
	}

	// 尝试单数形式
	if id := query["MarketplaceId"]; id != "" {
		return id
	}

	// 尝试小写形式
	if ids := query["marketplace_ids"]; ids != "" {
		if idx := len(ids); idx > 0 {
			for i := 0; i < len(ids); i++ {
				if ids[i] == ',' {
					return ids[:i]
				}
			}
			return ids
		}
	}

	// 回退值：使用 "global" 表示跨市场的速率限制
	return "global"
}

// extractOperationName 从 HTTP 方法和路径提取操作名称。
//
// 操作名称用于速率限制的细粒度控制。
// 格式：{API名称}:{操作名称}
//
// 例如：
//  - GET /orders/v0/orders -> orders:getOrders
//  - GET /orders/v0/orders/{orderId} -> orders:getOrder
//  - POST /feeds/2021-06-30/feeds -> feeds:createFeed
//
// 参数:
//   - method: HTTP 方法
//   - path: API 路径
//
// 返回值:
//   - string: 标准化的操作名称
func (c *Client) extractOperationName(method, path string) string {
	// 移除开头的 "/"
	if len(path) > 0 && path[0] == '/' {
		path = path[1:]
	}

	// 提取 API 名称（第一个路径段）
	apiName := ""
	endIdx := 0
	for i := 0; i < len(path); i++ {
		if path[i] == '/' {
			apiName = path[:i]
			endIdx = i + 1
			break
		}
	}

	if apiName == "" {
		// 如果没有找到 "/"，整个路径就是 API 名称
		return method + ":" + path
	}

	// 跳过版本号（如 "v0", "2021-06-30"）
	versionEndIdx := endIdx
	for i := endIdx; i < len(path); i++ {
		if path[i] == '/' {
			versionEndIdx = i + 1
			break
		}
	}

	// 提取资源名称（用于构建操作名）
	resourcePath := ""
	if versionEndIdx < len(path) {
		resourcePath = path[versionEndIdx:]

		// 移除路径参数（如 {orderId}）
		cleanPath := ""
		inBrace := false
		for i := 0; i < len(resourcePath); i++ {
			if resourcePath[i] == '{' {
				inBrace = true
			} else if resourcePath[i] == '}' {
				inBrace = false
			} else if !inBrace {
				cleanPath += string(resourcePath[i])
			}
		}
		resourcePath = cleanPath
	}

	// 构建标准化的操作名称
	operationName := apiName
	if resourcePath != "" {
		operationName += ":" + method + ":" + resourcePath
	} else {
		operationName += ":" + method
	}

	return operationName
}

// Get 执行 GET 请求。
//
// 参数:
//   - ctx: 请求上下文
//   - path: API 路径
//   - query: 查询参数（可选）
//   - result: 响应结果的指针
//
// 返回值:
//   - error: 如果请求失败，返回错误
//
// 示例:
//
//	var response orders.GetOrderResponse
//	err := client.Get(ctx, "/orders/v0/orders/123-4567890-1234567", nil, &response)
func (c *Client) Get(ctx context.Context, path string, query map[string]string, result interface{}) error {
	return c.DoRequest(ctx, "GET", path, query, nil, result)
}

// Post 执行 POST 请求。
//
// 参数:
//   - ctx: 请求上下文
//   - path: API 路径
//   - body: 请求体
//   - result: 响应结果的指针
//
// 返回值:
//   - error: 如果请求失败，返回错误
//
// 示例:
//
//	request := feeds.CreateFeedSpecification{
//	    FeedType: "POST_PRODUCT_DATA",
//	}
//	var response feeds.CreateFeedResponse
//	err := client.Post(ctx, "/feeds/2021-06-30/feeds", request, &response)
func (c *Client) Post(ctx context.Context, path string, body, result interface{}) error {
	return c.DoRequest(ctx, "POST", path, nil, body, result)
}

// Put 执行 PUT 请求。
//
// 参数:
//   - ctx: 请求上下文
//   - path: API 路径
//   - body: 请求体
//   - result: 响应结果的指针
//
// 返回值:
//   - error: 如果请求失败，返回错误
func (c *Client) Put(ctx context.Context, path string, body, result interface{}) error {
	return c.DoRequest(ctx, "PUT", path, nil, body, result)
}

// Delete 执行 DELETE 请求。
//
// 参数:
//   - ctx: 请求上下文
//   - path: API 路径
//   - result: 响应结果的指针（可选）
//
// 返回值:
//   - error: 如果请求失败，返回错误
func (c *Client) Delete(ctx context.Context, path string, result interface{}) error {
	return c.DoRequest(ctx, "DELETE", path, nil, nil, result)
}
