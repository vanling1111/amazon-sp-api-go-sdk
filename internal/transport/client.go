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

// Package transport 提供 HTTP 传输层功能。
//
// 此包封装了 HTTP 请求的发送和处理，支持中间件、重试和连接池管理。
//
// 基于官方 SP-API 文档:
//   - https://developer-docs.amazon.com/sp-api/docs/connecting-to-the-selling-partner-api
package transport

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/vanling1111/amazon-sp-api-go-sdk/internal/metrics"
)

// Client 是 HTTP 传输客户端。
//
// Client 封装了 HTTP 请求的发送和处理逻辑，
// 支持中间件链、重试机制和自定义配置。
// Client 是并发安全的。
type Client struct {
	baseURL     string
	httpClient  *http.Client
	middlewares []Middleware
	config      *Config
	metrics     metrics.Recorder // 可选的指标记录器
}

// Config 定义传输客户端的配置。
type Config struct {
	// Timeout 是 HTTP 请求的超时时间。
	Timeout time.Duration

	// MaxIdleConns 是最大空闲连接数。
	MaxIdleConns int

	// MaxIdleConnsPerHost 是每个主机的最大空闲连接数。
	MaxIdleConnsPerHost int

	// MaxConnsPerHost 是每个主机的最大连接数（包括活跃和空闲）。
	MaxConnsPerHost int

	// IdleConnTimeout 是空闲连接的超时时间。
	IdleConnTimeout time.Duration

	// UserAgent 是请求的 User-Agent 头。
	UserAgent string

	// Debug 启用调试模式。
	Debug bool
}

// DefaultConfig 返回默认配置。
//
// 默认配置针对生产环境优化：
//   - MaxIdleConns: 200 (支持更高并发)
//   - MaxIdleConnsPerHost: 20 (每个 SP-API 端点)
//   - IdleConnTimeout: 90s (保持连接复用)
//   - MaxConnsPerHost: 50 (限制单主机连接数)
func DefaultConfig() *Config {
	return &Config{
		Timeout:             30 * time.Second,
		MaxIdleConns:        200, // 增加以支持更高并发
		MaxIdleConnsPerHost: 20,  // 每个端点增加连接数
		IdleConnTimeout:     90 * time.Second,
		MaxConnsPerHost:     50, // 新增：限制单主机最大连接数
		UserAgent:           "amazon-sp-api-go-sdk/1.0.0",
		Debug:               false,
	}
}

// Middleware 定义中间件函数签名。
//
// 中间件可以在请求发送前后执行自定义逻辑，
// 例如添加认证头、记录日志、重试等。
type Middleware func(next Handler) Handler

// Handler 定义请求处理函数签名。
type Handler func(ctx context.Context, req *http.Request) (*http.Response, error)

// NewClient 创建新的 HTTP 传输客户端。
//
// 参数:
//   - baseURL: API 的基础 URL
//   - config: 客户端配置（如果为 nil，使用默认配置）
//
// 返回值:
//   - *Client: HTTP 传输客户端实例
//
// 示例:
//   client := transport.NewClient(
//       "https://sellingpartnerapi-na.amazon.com",
//       nil, // 使用默认配置
//   )
func NewClient(baseURL string, config *Config) *Client {
	if config == nil {
		config = DefaultConfig()
	}

	// 创建自定义 HTTP 传输
	tr := &http.Transport{
		MaxIdleConns:        config.MaxIdleConns,
		MaxIdleConnsPerHost: config.MaxIdleConnsPerHost,
		MaxConnsPerHost:     config.MaxConnsPerHost,
		IdleConnTimeout:     config.IdleConnTimeout,
		// 启用 HTTP/2 支持
		ForceAttemptHTTP2:     true,
		// 禁用压缩以提高性能（SP-API 响应通常已压缩）
		DisableCompression:    false,
		// 启用 Keep-Alive
		DisableKeepAlives:     false,
		// 连接超时
		ResponseHeaderTimeout: 10 * time.Second,
		// TLS 握手超时
		TLSHandshakeTimeout:   10 * time.Second,
	}

	// 创建 HTTP 客户端
	httpClient := &http.Client{
		Timeout:   config.Timeout,
		Transport: tr,
	}

	return &Client{
		baseURL:     baseURL,
		httpClient:  httpClient,
		middlewares: []Middleware{},
		config:      config,
		metrics:     metrics.DefaultRecorder, // 默认使用 NoOp 记录器
	}
}

// Use 添加中间件。
//
// 中间件按添加顺序执行。
//
// 参数:
//   - middleware: 要添加的中间件
//
// 示例:
//   client.Use(loggingMiddleware)
//   client.Use(retryMiddleware)
func (c *Client) Use(middleware Middleware) {
	c.middlewares = append(c.middlewares, middleware)
}

// SetMetrics 设置指标记录器。
//
// 参数:
//   - recorder: 指标记录器实现
//
// 示例:
//   client.SetMetrics(myPrometheusRecorder)
func (c *Client) SetMetrics(recorder metrics.Recorder) {
	if recorder != nil {
		c.metrics = recorder
	}
}

// Do 发送 HTTP 请求。
//
// 此方法会依次应用所有注册的中间件，然后发送请求。
//
// 参数:
//   - ctx: 请求上下文
//   - req: HTTP 请求
//
// 返回值:
//   - *http.Response: HTTP 响应
//   - error: 如果请求失败，返回错误
//
// 示例:
//   req, _ := http.NewRequest("GET", "/orders/v0/orders", nil)
//   resp, err := client.Do(ctx, req)
//   if err != nil {
//       log.Fatal(err)
//   }
//   defer resp.Body.Close()
func (c *Client) Do(ctx context.Context, req *http.Request) (*http.Response, error) {
	// 设置 User-Agent
	if req.Header.Get("User-Agent") == "" {
		req.Header.Set("User-Agent", c.config.UserAgent)
	}

	// 构建处理链
	handler := c.buildHandler()

	// 执行请求
	return handler(ctx, req)
}

// buildHandler 构建处理链。
func (c *Client) buildHandler() Handler {
	// 最终处理器：发送 HTTP 请求
	handler := c.doRequest

	// 反向应用中间件（后添加的先执行）
	for i := len(c.middlewares) - 1; i >= 0; i-- {
		handler = c.middlewares[i](handler)
	}

	return handler
}

// doRequest 是最终的请求处理器。
func (c *Client) doRequest(ctx context.Context, req *http.Request) (*http.Response, error) {
	// 确保请求包含 context
	req = req.WithContext(ctx)

	// 记录请求开始时间
	startTime := time.Now()

	// 发送请求
	resp, err := c.httpClient.Do(req)
	
	// 记录请求延迟
	duration := time.Since(startTime)
	c.metrics.RecordTiming(metrics.MetricRequestDuration, duration, map[string]string{
		metrics.LabelOperation: req.URL.Path,
	})

	// 记录请求总数
	c.metrics.RecordCounter(metrics.MetricRequestTotal, 1, map[string]string{
		metrics.LabelOperation: req.URL.Path,
	})

	if err != nil {
		// 记录错误
		c.metrics.RecordCounter(metrics.MetricRequestErrors, 1, map[string]string{
			metrics.LabelOperation: req.URL.Path,
			metrics.LabelErrorType: "network",
		})
		return nil, fmt.Errorf("send HTTP request: %w", err)
	}

	// 记录响应状态码
	c.metrics.RecordCounter(metrics.MetricRequestTotal, 1, map[string]string{
		metrics.LabelOperation: req.URL.Path,
		metrics.LabelStatusCode: strconv.Itoa(resp.StatusCode),
	})

	return resp, nil
}

// BaseURL 返回基础 URL。
func (c *Client) BaseURL() string {
	return c.baseURL
}

// HTTPClient 返回底层的 HTTP 客户端。
//
// 这允许高级用户直接访问 HTTP 客户端进行自定义配置。
func (c *Client) HTTPClient() *http.Client {
	return c.httpClient
}

