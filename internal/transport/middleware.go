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

package transport

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware 创建日志记录中间件。
//
// 此中间件记录每个请求的方法、URL 和执行时间。
//
// 返回值:
//   - Middleware: 日志中间件
//
// 示例:
//   client.Use(transport.LoggingMiddleware())
func LoggingMiddleware() Middleware {
	return func(next Handler) Handler {
		return func(ctx context.Context, req *http.Request) (*http.Response, error) {
			start := time.Now()

			// 记录请求
			log.Printf("[HTTP] --> %s %s", req.Method, req.URL.String())

			// 执行请求
			resp, err := next(ctx, req)

			// 记录响应
			duration := time.Since(start)
			if err != nil {
				log.Printf("[HTTP] <-- %s %s - Error: %v (took %v)",
					req.Method, req.URL.String(), err, duration)
			} else {
				log.Printf("[HTTP] <-- %s %s - %d (took %v)",
					req.Method, req.URL.String(), resp.StatusCode, duration)
			}

			return resp, err
		}
	}
}

// UserAgentMiddleware 创建 User-Agent 中间件。
//
// 此中间件为请求添加自定义 User-Agent 头。
//
// 参数:
//   - userAgent: User-Agent 字符串
//
// 返回值:
//   - Middleware: User-Agent 中间件
//
// 示例:
//   client.Use(transport.UserAgentMiddleware("my-app/1.0.0"))
func UserAgentMiddleware(userAgent string) Middleware {
	return func(next Handler) Handler {
		return func(ctx context.Context, req *http.Request) (*http.Response, error) {
			if req.Header.Get("User-Agent") == "" {
				req.Header.Set("User-Agent", userAgent)
			}
			return next(ctx, req)
		}
	}
}

// HeaderMiddleware 创建添加自定义头的中间件。
//
// 此中间件为每个请求添加指定的 HTTP 头。
//
// 参数:
//   - headers: 要添加的 HTTP 头
//
// 返回值:
//   - Middleware: 头部中间件
//
// 示例:
//   headers := map[string]string{
//       "X-Custom-Header": "value",
//   }
//   client.Use(transport.HeaderMiddleware(headers))
func HeaderMiddleware(headers map[string]string) Middleware {
	return func(next Handler) Handler {
		return func(ctx context.Context, req *http.Request) (*http.Response, error) {
			for key, value := range headers {
				if req.Header.Get(key) == "" {
					req.Header.Set(key, value)
				}
			}
			return next(ctx, req)
		}
	}
}

// TimeoutMiddleware 创建超时中间件。
//
// 此中间件为请求设置超时时间。
//
// 参数:
//   - timeout: 超时时间
//
// 返回值:
//   - Middleware: 超时中间件
//
// 示例:
//   client.Use(transport.TimeoutMiddleware(30 * time.Second))
func TimeoutMiddleware(timeout time.Duration) Middleware {
	return func(next Handler) Handler {
		return func(ctx context.Context, req *http.Request) (*http.Response, error) {
			ctx, cancel := context.WithTimeout(ctx, timeout)
			defer cancel()

			return next(ctx, req)
		}
	}
}

// RequestIDMiddleware 创建请求 ID 中间件。
//
// 此中间件为每个请求生成唯一的请求 ID，
// 并将其添加到请求头中。
//
// 返回值:
//   - Middleware: 请求 ID 中间件
//
// 示例:
//   client.Use(transport.RequestIDMiddleware())
func RequestIDMiddleware() Middleware {
	return func(next Handler) Handler {
		return func(ctx context.Context, req *http.Request) (*http.Response, error) {
			// 生成请求 ID
			requestID := generateRequestID()

			// 添加到请求头
			req.Header.Set("X-Request-ID", requestID)

			// 执行请求
			return next(ctx, req)
		}
	}
}

// DateMiddleware 创建日期时间中间件。
//
// 此中间件为每个请求自动添加 x-amz-date 头部。
// 格式符合 AWS/Amazon 标准：ISO 8601 基本格式（YYYYMMDDTHHmmssZ）。
//
// 根据官方 SP-API 文档要求：
//   - 头名称：x-amz-date
//   - 格式：20190430T123600Z
//   - 说明：请求的日期和时间
//
// 返回值:
//   - Middleware: 日期时间中间件
//
// 示例:
//   client.Use(transport.DateMiddleware())
//
// 官方文档:
//   https://developer-docs.amazon.com/sp-api/docs/connecting-to-the-selling-partner-api#step-3-add-headers-to-the-uri
func DateMiddleware() Middleware {
	return func(next Handler) Handler {
		return func(ctx context.Context, req *http.Request) (*http.Response, error) {
			// 只在头部不存在时才添加
			if req.Header.Get("x-amz-date") == "" {
				// ISO 8601 基本格式：20060102T150405Z
				dateStr := time.Now().UTC().Format("20060102T150405Z")
				req.Header.Set("x-amz-date", dateStr)
			}
			return next(ctx, req)
		}
	}
}

// generateRequestID 生成唯一的请求 ID。
func generateRequestID() string {
	// 简单实现：使用时间戳
	// 实际项目中可以使用 UUID 或其他更复杂的方案
	return fmt.Sprintf("req-%d", time.Now().UnixNano())
}
