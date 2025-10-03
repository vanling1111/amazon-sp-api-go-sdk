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

// Package tracing 提供分布式追踪功能。
//
// 此包集成 OpenTelemetry，支持 Jaeger、Zipkin 等追踪系统。
package tracing

import (
	"context"
	"fmt"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

const (
	// instrumentationName 是追踪的 instrumentation 名称
	instrumentationName = "github.com/vanling1111/amazon-sp-api-go-sdk"
)

// Middleware 追踪中间件。
//
// 此中间件为每个 HTTP 请求创建 span，记录请求详情。
//
// 参数:
//   - tracer: OpenTelemetry tracer（可选，nil 则使用全局 tracer）
//
// 返回值:
//   - func: 中间件函数
//
// 示例:
//
//	tracer := otel.Tracer("sp-api")
//	middleware := tracing.Middleware(tracer)
func Middleware(tracer trace.Tracer) func(next func(context.Context, *http.Request) (*http.Response, error)) func(context.Context, *http.Request) (*http.Response, error) {
	if tracer == nil {
		tracer = otel.Tracer(instrumentationName)
	}

	return func(next func(context.Context, *http.Request) (*http.Response, error)) func(context.Context, *http.Request) (*http.Response, error) {
		return func(ctx context.Context, req *http.Request) (*http.Response, error) {
			// 创建 span
			spanName := fmt.Sprintf("HTTP %s %s", req.Method, req.URL.Path)
			ctx, span := tracer.Start(ctx, spanName,
				trace.WithSpanKind(trace.SpanKindClient),
			)
			defer span.End()

			// 添加基础属性
			span.SetAttributes(
				attribute.String("http.method", req.Method),
				attribute.String("http.url", req.URL.String()),
				attribute.String("http.host", req.Host),
				attribute.String("http.scheme", req.URL.Scheme),
			)

			// 执行请求
			resp, err := next(ctx, req)

			// 记录结果
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
				return resp, err
			}

			// 添加响应属性
			if resp != nil {
				span.SetAttributes(
					attribute.Int("http.status_code", resp.StatusCode),
				)

				if resp.StatusCode >= 400 {
					span.SetStatus(codes.Error, fmt.Sprintf("HTTP %d", resp.StatusCode))
				} else {
					span.SetStatus(codes.Ok, "")
				}
			}

			return resp, nil
		}
	}
}

// StartSpan 开始一个新的 span。
//
// 便捷函数，用于在业务代码中创建 span。
//
// 参数:
//   - ctx: 上下文
//   - name: span 名称
//   - attrs: 可选属性
//
// 返回值:
//   - context.Context: 包含 span 的新上下文
//   - trace.Span: span 实例
//
// 示例:
//
//	ctx, span := tracing.StartSpan(ctx, "ProcessOrder",
//	    attribute.String("order.id", orderID),
//	)
//	defer span.End()
func StartSpan(ctx context.Context, name string, attrs ...attribute.KeyValue) (context.Context, trace.Span) {
	tracer := otel.Tracer(instrumentationName)
	return tracer.Start(ctx, name, trace.WithAttributes(attrs...))
}

// RecordError 记录错误到当前 span。
//
// 参数:
//   - ctx: 包含 span 的上下文
//   - err: 错误
func RecordError(ctx context.Context, err error) {
	span := trace.SpanFromContext(ctx)
	if span != nil && span.IsRecording() {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}
}

// SetAttributes 设置当前 span 的属性。
//
// 参数:
//   - ctx: 包含 span 的上下文
//   - attrs: 属性列表
func SetAttributes(ctx context.Context, attrs ...attribute.KeyValue) {
	span := trace.SpanFromContext(ctx)
	if span != nil && span.IsRecording() {
		span.SetAttributes(attrs...)
	}
}

