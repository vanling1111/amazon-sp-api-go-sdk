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
	"time"
)

// noOpLogger 是Logger接口的no-op实现。
//
// 默认情况下使用此实现，不输出任何日志。
type noOpLogger struct{}

// NewNoOpLogger 创建一个no-op logger。
func NewNoOpLogger() Logger {
	return &noOpLogger{}
}

func (n *noOpLogger) Debug(msg string, fields ...Field) {}
func (n *noOpLogger) Info(msg string, fields ...Field)  {}
func (n *noOpLogger) Warn(msg string, fields ...Field)  {}
func (n *noOpLogger) Error(msg string, fields ...Field) {}
func (n *noOpLogger) With(fields ...Field) Logger       { return n }

// noOpMetrics 是MetricsCollector接口的no-op实现。
//
// 默认情况下使用此实现，不收集任何指标。
type noOpMetrics struct{}

// NewNoOpMetrics 创建一个no-op metrics collector。
func NewNoOpMetrics() MetricsCollector {
	return &noOpMetrics{}
}

func (n *noOpMetrics) RecordRequest(api, method string, duration time.Duration, statusCode int) {}
func (n *noOpMetrics) RecordError(api, errorType string)                                        {}
func (n *noOpMetrics) RecordRateLimitWait(api string, duration time.Duration)                   {}

// noOpTracer 是Tracer接口的no-op实现。
//
// 默认情况下使用此实现，不进行任何追踪。
type noOpTracer struct{}

// NewNoOpTracer 创建一个no-op tracer。
func NewNoOpTracer() Tracer {
	return &noOpTracer{}
}

func (n *noOpTracer) StartSpan(ctx context.Context, name string) (context.Context, Span) {
	return ctx, &noOpSpan{}
}

// noOpSpan 是Span接口的no-op实现。
type noOpSpan struct{}

func (n *noOpSpan) End()                                       {}
func (n *noOpSpan) SetAttribute(key string, value interface{}) {}
func (n *noOpSpan) RecordError(err error)                      {}
