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
package tracing

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
)

// TestMiddleware tests tracing middleware
func TestMiddleware(t *testing.T) {
	// 创建 span 记录器
	exporter := tracetest.NewInMemoryExporter()
	tp := trace.NewTracerProvider(
		trace.WithSyncer(exporter),
	)
	defer tp.Shutdown(context.Background())

	otel.SetTracerProvider(tp)
	tracer := tp.Tracer("test")

	// 创建中间件
	middleware := Middleware(tracer)

	// 创建测试请求
	req := httptest.NewRequest("GET", "https://sellingpartnerapi-na.amazon.com/orders/v0/orders", nil)
	ctx := context.Background()

	// 执行请求
	handler := middleware(func(ctx context.Context, req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
		}, nil
	})

	resp, err := handler(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	// 验证 span 被创建
	spans := exporter.GetSpans()
	assert.Len(t, spans, 1)

	span := spans[0]
	assert.Equal(t, "HTTP GET /orders/v0/orders", span.Name)
	assert.Equal(t, "GET", span.Attributes[0].Value.AsString())
}

// TestMiddleware_Error tests error recording
func TestMiddleware_Error(t *testing.T) {
	exporter := tracetest.NewInMemoryExporter()
	tp := trace.NewTracerProvider(trace.WithSyncer(exporter))
	defer tp.Shutdown(context.Background())

	tracer := tp.Tracer("test")
	middleware := Middleware(tracer)

	req := httptest.NewRequest("GET", "https://test.com/api", nil)
	testErr := errors.New("test error")

	handler := middleware(func(ctx context.Context, req *http.Request) (*http.Response, error) {
		return nil, testErr
	})

	_, err := handler(context.Background(), req)
	assert.Error(t, err)

	// 验证错误被记录
	spans := exporter.GetSpans()
	assert.Len(t, spans, 1)
	assert.Len(t, spans[0].Events, 1) // Error event
}

// TestStartSpan tests creating spans
func TestStartSpan(t *testing.T) {
	exporter := tracetest.NewInMemoryExporter()
	tp := trace.NewTracerProvider(trace.WithSyncer(exporter))
	defer tp.Shutdown(context.Background())

	otel.SetTracerProvider(tp)

	ctx := context.Background()
	ctx, span := StartSpan(ctx, "TestOperation",
		attribute.String("test.key", "test.value"),
	)
	span.End()

	spans := exporter.GetSpans()
	assert.Len(t, spans, 1)
	assert.Equal(t, "TestOperation", spans[0].Name)
}

// TestRecordError tests error recording
func TestRecordError(t *testing.T) {
	exporter := tracetest.NewInMemoryExporter()
	tp := trace.NewTracerProvider(trace.WithSyncer(exporter))
	defer tp.Shutdown(context.Background())

	otel.SetTracerProvider(tp)

	ctx, span := StartSpan(context.Background(), "TestError")
	RecordError(ctx, errors.New("test error"))
	span.End()

	spans := exporter.GetSpans()
	assert.Len(t, spans, 1)
	assert.Len(t, spans[0].Events, 1)
}

// TestSetAttributes tests setting attributes
func TestSetAttributes(t *testing.T) {
	exporter := tracetest.NewInMemoryExporter()
	tp := trace.NewTracerProvider(trace.WithSyncer(exporter))
	defer tp.Shutdown(context.Background())

	otel.SetTracerProvider(tp)

	ctx, span := StartSpan(context.Background(), "TestAttrs")
	SetAttributes(ctx,
		attribute.String("key1", "value1"),
		attribute.Int("key2", 42),
	)
	span.End()

	spans := exporter.GetSpans()
	assert.Len(t, spans, 1)
	// 应该有初始属性 + 我们添加的
	assert.Greater(t, len(spans[0].Attributes), 1)
}
