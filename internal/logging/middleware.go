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
package logging

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"time"
)

// Options 日志中间件选项
type Options struct {
	// LogHeaders 是否记录 HTTP 头
	LogHeaders bool

	// LogBody 是否记录请求/响应体
	LogBody bool

	// MaxBodySize 记录的最大 body 大小（字节）
	MaxBodySize int

	// RedactHeaders 需要脱敏的 HTTP 头（如 token）
	RedactHeaders []string

	// RedactFields 需要脱敏的 JSON 字段
	RedactFields []string
}

// DefaultOptions 返回默认日志选项
func DefaultOptions() *Options {
	return &Options{
		LogHeaders:  true,
		LogBody:     false, // 默认不记录 body（可能很大）
		MaxBodySize: 4096,  // 4KB
		RedactHeaders: []string{
			"x-amz-access-token",
			"authorization",
		},
		RedactFields: []string{
			"refreshToken",
			"clientSecret",
		},
	}
}

// LoggingMiddleware 日志中间件
type LoggingMiddleware struct {
	logger  Logger
	options *Options
}

// NewMiddleware 创建日志中间件
func NewMiddleware(logger Logger, options *Options) *LoggingMiddleware {
	if options == nil {
		options = DefaultOptions()
	}

	return &LoggingMiddleware{
		logger:  logger,
		options: options,
	}
}

// Wrap 包装 HTTP 客户端
func (m *LoggingMiddleware) Wrap(next func(context.Context, *http.Request) (*http.Response, error)) func(context.Context, *http.Request) (*http.Response, error) {
	return func(ctx context.Context, req *http.Request) (*http.Response, error) {
		start := time.Now()

		// 记录请求
		m.logRequest(req)

		// 复制请求体（如果需要记录）
		var bodyBytes []byte
		if m.options.LogBody && req.Body != nil {
			bodyBytes, _ = io.ReadAll(req.Body)
			req.Body.Close()
			req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		}

		// 执行请求
		resp, err := next(ctx, req)

		duration := time.Since(start)

		// 记录响应
		m.logResponse(req, resp, err, duration)

		return resp, err
	}
}

// logRequest 记录请求
func (m *LoggingMiddleware) logRequest(req *http.Request) {
	fields := []Field{
		String("method", req.Method),
		String("url", req.URL.String()),
		String("host", req.Host),
	}

	if m.options.LogHeaders {
		headers := m.redactHeaders(req.Header)
		fields = append(fields, Field{Key: "headers", Value: headers})
	}

	m.logger.Debug("HTTP Request", fields...)
}

// logResponse 记录响应
func (m *LoggingMiddleware) logResponse(req *http.Request, resp *http.Response, err error, duration time.Duration) {
	fields := []Field{
		String("method", req.Method),
		String("url", req.URL.String()),
		Duration("duration", duration),
	}

	if err != nil {
		// 请求失败
		fields = append(fields, Error(err))
		m.logger.Error("HTTP Request Failed", fields...)
		return
	}

	// 请求成功
	fields = append(fields, Int("status", resp.StatusCode))

	if m.options.LogHeaders && resp != nil {
		headers := m.redactHeaders(resp.Header)
		fields = append(fields, Field{Key: "response_headers", Value: headers})
	}

	// 根据状态码选择日志级别
	if resp.StatusCode >= 500 {
		m.logger.Error("HTTP Server Error", fields...)
	} else if resp.StatusCode >= 400 {
		m.logger.Warn("HTTP Client Error", fields...)
	} else {
		m.logger.Info("HTTP Request Success", fields...)
	}
}

// redactHeaders 脱敏 HTTP 头
func (m *LoggingMiddleware) redactHeaders(headers http.Header) map[string]string {
	result := make(map[string]string)

	for key, values := range headers {
		value := values[0]

		// 检查是否需要脱敏
		redact := false
		for _, redactKey := range m.options.RedactHeaders {
			if key == redactKey || key == http.CanonicalHeaderKey(redactKey) {
				redact = true
				break
			}
		}

		if redact {
			result[key] = "***REDACTED***"
		} else {
			result[key] = value
		}
	}

	return result
}
