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
// Package prometheus 提供 Prometheus 指标导出功能。
package prometheus

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Metrics Prometheus 指标收集器
type Metrics struct {
	// 请求总数
	requestsTotal *prometheus.CounterVec

	// 请求延迟
	requestDuration *prometheus.HistogramVec

	// 错误总数
	errorsTotal *prometheus.CounterVec

	// 速率限制等待时间
	rateLimitWait *prometheus.HistogramVec
}

// NewMetrics 创建 Prometheus 指标收集器。
//
// 参数:
//   - namespace: 指标命名空间（如 "spapi"）
//
// 返回值:
//   - *Metrics: 指标收集器实例
//
// 示例:
//
//	metrics := prometheus.NewMetrics("spapi")
func NewMetrics(namespace string) *Metrics {
	return &Metrics{
		requestsTotal: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Name:      "requests_total",
				Help:      "Total number of API requests",
			},
			[]string{"api", "method", "status"},
		),

		requestDuration: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Name:      "request_duration_seconds",
				Help:      "API request duration in seconds",
				Buckets:   prometheus.DefBuckets,
			},
			[]string{"api", "method"},
		),

		errorsTotal: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Name:      "errors_total",
				Help:      "Total number of errors",
			},
			[]string{"api", "error_type"},
		),

		rateLimitWait: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Name:      "rate_limit_wait_seconds",
				Help:      "Time spent waiting for rate limiter",
				Buckets:   []float64{0.001, 0.01, 0.1, 0.5, 1, 5, 10},
			},
			[]string{"api"},
		),
	}
}

// RecordRequest 记录 API 请求。
//
// 参数:
//   - api: API 名称（如 "orders"）
//   - method: HTTP 方法
//   - status: HTTP 状态码
//   - duration: 请求持续时间（秒）
func (m *Metrics) RecordRequest(api, method string, status int, duration float64) {
	statusStr := http.StatusText(status)
	if statusStr == "" {
		statusStr = fmt.Sprintf("%d", status)
	}

	m.requestsTotal.WithLabelValues(api, method, statusStr).Inc()
	m.requestDuration.WithLabelValues(api, method).Observe(duration)
}

// RecordError 记录错误。
//
// 参数:
//   - api: API 名称
//   - errorType: 错误类型（如 "rate_limit", "network", "server"）
func (m *Metrics) RecordError(api, errorType string) {
	m.errorsTotal.WithLabelValues(api, errorType).Inc()
}

// RecordRateLimitWait 记录速率限制等待时间。
//
// 参数:
//   - api: API 名称
//   - waitTime: 等待时间（秒）
func (m *Metrics) RecordRateLimitWait(api string, waitTime float64) {
	m.rateLimitWait.WithLabelValues(api).Observe(waitTime)
}
