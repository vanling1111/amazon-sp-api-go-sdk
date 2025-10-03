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

package transport

import (
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"time"
)

// RetryConfig 定义重试配置。
type RetryConfig struct {
	// MaxRetries 是最大重试次数。
	MaxRetries int

	// InitialInterval 是初始重试间隔。
	InitialInterval time.Duration

	// MaxInterval 是最大重试间隔。
	MaxInterval time.Duration

	// Multiplier 是退避乘数。
	Multiplier float64

	// ShouldRetry 是判断是否应该重试的函数。
	// 如果为 nil，使用默认实现。
	ShouldRetry func(resp *http.Response, err error) bool
}

// DefaultRetryConfig 返回默认重试配置。
func DefaultRetryConfig() *RetryConfig {
	return &RetryConfig{
		MaxRetries:      3,
		InitialInterval: 1 * time.Second,
		MaxInterval:     30 * time.Second,
		Multiplier:      2.0,
		ShouldRetry:     defaultShouldRetry,
	}
}

// RetryMiddleware 创建重试中间件。
//
// 此中间件实现指数退避重试策略。
//
// 参数:
//   - config: 重试配置（如果为 nil，使用默认配置）
//
// 返回值:
//   - Middleware: 重试中间件
//
// 示例:
//   config := &transport.RetryConfig{
//       MaxRetries: 3,
//       InitialInterval: 1 * time.Second,
//   }
//   client.Use(transport.RetryMiddleware(config))
func RetryMiddleware(config *RetryConfig) Middleware {
	if config == nil {
		config = DefaultRetryConfig()
	}

	if config.ShouldRetry == nil {
		config.ShouldRetry = defaultShouldRetry
	}

	return func(next Handler) Handler {
		return func(ctx context.Context, req *http.Request) (*http.Response, error) {
			var resp *http.Response
			var err error

			// 保存请求体（如果有）以便重试
			var bodyBytes []byte
			if req.Body != nil {
				bodyBytes, err = io.ReadAll(req.Body)
				if err != nil {
					return nil, fmt.Errorf("read request body: %w", err)
				}
				req.Body.Close()
			}

			// 执行请求和重试
			for attempt := 0; attempt <= config.MaxRetries; attempt++ {
				// 重新设置请求体
				if bodyBytes != nil {
					req.Body = io.NopCloser(byteReader(bodyBytes))
				}

				// 执行请求
				resp, err = next(ctx, req)

				// 检查是否应该重试
				if !config.ShouldRetry(resp, err) {
					// 不需要重试，直接返回
					return resp, err
				}

				// 如果还有重试机会，等待后重试
				if attempt < config.MaxRetries {
					// 计算退避时间
					backoff := calculateBackoff(
						attempt,
						config.InitialInterval,
						config.MaxInterval,
						config.Multiplier,
					)

					// 等待
					select {
					case <-time.After(backoff):
						// 继续重试
					case <-ctx.Done():
						// 上下文被取消
						return nil, ctx.Err()
					}
				}
			}

			// 所有重试都失败
			if err != nil {
				return nil, fmt.Errorf("max retries exceeded: %w", err)
			}

			return resp, nil
		}
	}
}

// defaultShouldRetry 是默认的重试判断函数。
//
// 以下情况会重试：
// - 网络错误
// - HTTP 5xx 错误
// - HTTP 429 Too Many Requests
func defaultShouldRetry(resp *http.Response, err error) bool {
	// 网络错误，重试
	if err != nil {
		return true
	}

	// HTTP 5xx 错误，重试
	if resp.StatusCode >= 500 {
		return true
	}

	// HTTP 429 Too Many Requests，重试
	if resp.StatusCode == http.StatusTooManyRequests {
		return true
	}

	// 其他情况不重试
	return false
}

// calculateBackoff 计算退避时间。
//
// 使用指数退避算法：interval = initial * (multiplier ^ attempt)
func calculateBackoff(attempt int, initial, max time.Duration, multiplier float64) time.Duration {
	backoff := float64(initial) * math.Pow(multiplier, float64(attempt))

	if backoff > float64(max) {
		return max
	}

	return time.Duration(backoff)
}

// byteReader 从字节切片创建 io.Reader。
type byteReader []byte

func (b byteReader) Read(p []byte) (n int, err error) {
	if len(b) == 0 {
		return 0, io.EOF
	}
	n = copy(p, b)
	return n, nil
}
