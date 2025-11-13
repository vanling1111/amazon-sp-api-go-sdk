// Copyright 2025 Amazon SP-API Go SDK Authors.
//
// This file is part of Amazon SP-API Go SDK.
//
// Amazon SP-API Go SDK is dual-licensed:
//
// 1. GNU Affero General Public License v3.0 (AGPL-3.0) for open source use
//   - Free for personal, educational, and open source projects
//   - Your project must also be open sourced under AGPL-3.0
//   - See: https://www.gnu.org/licenses/agpl-3.0.html
//
// 2. Commercial License for proprietary/commercial use
//   - Required for any commercial, enterprise, or proprietary use
//   - Allows closed source distribution
//   - Contact: vanling1111@gmail.com
//
// Unless you have obtained a commercial license, this file is licensed
// under AGPL-3.0. By using this software, you agree to comply with the
// terms of the applicable license. All rights reserved.

package spapi

import (
	"context"
	"net/http"
	"time"
)

// Middleware 定义中间件类型。
//
// 中间件可以在请求前后执行自定义逻辑，例如日志记录、指标收集、错误处理等。
//
// 示例:
//
//	func LoggingMiddleware(logger Logger) Middleware {
//	    return func(next Handler) Handler {
//	        return func(ctx context.Context, req *http.Request) (*http.Response, error) {
//	            start := time.Now()
//	            logger.Info("request started", Field{"path", req.URL.Path})
//	            resp, err := next(ctx, req)
//	            logger.Info("request completed", Field{"duration", time.Since(start)})
//	            return resp, err
//	        }
//	    }
//	}
type Middleware func(next Handler) Handler

// Handler 定义请求处理器类型。
//
// Handler接收context和HTTP请求，返回HTTP响应和错误。
type Handler func(ctx context.Context, req *http.Request) (*http.Response, error)

// LoggingMiddleware 创建日志中间件。
//
// 记录每个请求的开始和结束，包括耗时、状态码等信息。
//
// 参数:
//   - logger: 日志器实现
//
// 示例:
//
//	client := spapi.NewClient(
//	    spapi.WithRegion(spapi.RegionNA),
//	    spapi.WithCredentials(...),
//	    spapi.WithMiddleware(
//	        spapi.LoggingMiddleware(logger),
//	    ),
//	)
func LoggingMiddleware(logger Logger) Middleware {
	return func(next Handler) Handler {
		return func(ctx context.Context, req *http.Request) (*http.Response, error) {
			start := time.Now()
			
			logger.Info("request started",
				Field{"method", req.Method},
				Field{"url", req.URL.String()},
			)
			
			resp, err := next(ctx, req)
			
			duration := time.Since(start)
			
			if err != nil {
				logger.Error("request failed",
					Field{"method", req.Method},
					Field{"url", req.URL.String()},
					Field{"duration", duration},
					Field{"error", err.Error()},
				)
			} else {
				logger.Info("request completed",
					Field{"method", req.Method},
					Field{"url", req.URL.String()},
					Field{"duration", duration},
					Field{"status", resp.StatusCode},
				)
			}
			
			return resp, err
		}
	}
}

// MetricsMiddleware 创建指标收集中间件。
//
// 记录每个请求的指标，包括请求数、错误数、耗时等。
//
// 参数:
//   - metrics: 指标收集器实现
//
// 示例:
//
//	client := spapi.NewClient(
//	    spapi.WithRegion(spapi.RegionNA),
//	    spapi.WithCredentials(...),
//	    spapi.WithMiddleware(
//	        spapi.MetricsMiddleware(metrics),
//	    ),
//	)
func MetricsMiddleware(metrics MetricsCollector) Middleware {
	return func(next Handler) Handler {
		return func(ctx context.Context, req *http.Request) (*http.Response, error) {
			start := time.Now()
			
			resp, err := next(ctx, req)
			
			duration := time.Since(start)
			
			if err != nil {
				metrics.RecordError(req.URL.Path, "request_error")
			} else {
				metrics.RecordRequest(req.URL.Path, req.Method, duration, resp.StatusCode)
			}
			
			return resp, err
		}
	}
}

// TracingMiddleware 创建分布式追踪中间件。
//
// 为每个请求创建span，记录追踪信息。
//
// 参数:
//   - tracer: 追踪器实现
//
// 示例:
//
//	client := spapi.NewClient(
//	    spapi.WithRegion(spapi.RegionNA),
//	    spapi.WithCredentials(...),
//	    spapi.WithMiddleware(
//	        spapi.TracingMiddleware(tracer),
//	    ),
//	)
func TracingMiddleware(tracer Tracer) Middleware {
	return func(next Handler) Handler {
		return func(ctx context.Context, req *http.Request) (*http.Response, error) {
			ctx, span := tracer.StartSpan(ctx, "sp-api.request")
			defer span.End()
			
			span.SetAttribute("http.method", req.Method)
			span.SetAttribute("http.url", req.URL.String())
			
			resp, err := next(ctx, req)
			
			if err != nil {
				span.RecordError(err)
			} else {
				span.SetAttribute("http.status_code", resp.StatusCode)
			}
			
			return resp, err
		}
	}
}

// ChainMiddlewares 链接多个中间件。
//
// 按照提供的顺序执行中间件。
//
// 参数:
//   - middlewares: 中间件列表
//
// 返回值:
//   - Middleware: 链接后的中间件
func ChainMiddlewares(middlewares ...Middleware) Middleware {
	return func(next Handler) Handler {
		// 从后向前应用中间件
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}
		return next
	}
}
