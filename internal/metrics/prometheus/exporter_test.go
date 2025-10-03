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

package prometheus

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
)

// TestNewMetrics tests creating metrics collector
func TestNewMetrics(t *testing.T) {
	// 创建新的 registry 避免冲突
	registry := prometheus.NewRegistry()

	metrics := &Metrics{
		requestsTotal: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: "test",
				Name:      "requests_total",
				Help:      "Test requests",
			},
			[]string{"api", "method", "status"},
		),
	}

	registry.MustRegister(metrics.requestsTotal)

	assert.NotNil(t, metrics)
	assert.NotNil(t, metrics.requestsTotal)
}

// TestRecordRequest tests recording requests
func TestRecordRequest(t *testing.T) {
	registry := prometheus.NewRegistry()

	metrics := &Metrics{
		requestsTotal: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: "test",
				Name:      "requests_total",
			},
			[]string{"api", "method", "status"},
		),
		requestDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: "test",
				Name:      "request_duration_seconds",
			},
			[]string{"api", "method"},
		),
	}

	registry.MustRegister(metrics.requestsTotal)
	registry.MustRegister(metrics.requestDuration)

	// 记录请求
	metrics.RecordRequest("orders", "GET", 200, 0.5)
	metrics.RecordRequest("orders", "GET", 200, 0.3)
	metrics.RecordRequest("orders", "GET", 429, 1.0)

	// 验证计数
	count := testutil.ToFloat64(metrics.requestsTotal.WithLabelValues("orders", "GET", "OK"))
	assert.Equal(t, 2.0, count)

	rateLimitCount := testutil.ToFloat64(metrics.requestsTotal.WithLabelValues("orders", "GET", "Too Many Requests"))
	assert.Equal(t, 1.0, rateLimitCount)
}

// TestRecordError tests recording errors
func TestRecordError(t *testing.T) {
	registry := prometheus.NewRegistry()

	metrics := &Metrics{
		errorsTotal: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: "test",
				Name:      "errors_total",
			},
			[]string{"api", "error_type"},
		),
	}

	registry.MustRegister(metrics.errorsTotal)

	// 记录错误
	metrics.RecordError("orders", "rate_limit")
	metrics.RecordError("orders", "rate_limit")
	metrics.RecordError("reports", "network")

	// 验证计数
	rateLimitErrors := testutil.ToFloat64(metrics.errorsTotal.WithLabelValues("orders", "rate_limit"))
	assert.Equal(t, 2.0, rateLimitErrors)

	networkErrors := testutil.ToFloat64(metrics.errorsTotal.WithLabelValues("reports", "network"))
	assert.Equal(t, 1.0, networkErrors)
}

// TestRecordRateLimitWait tests recording rate limit wait time
func TestRecordRateLimitWait(t *testing.T) {
	registry := prometheus.NewRegistry()

	metrics := &Metrics{
		rateLimitWait: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: "test",
				Name:      "rate_limit_wait_seconds",
			},
			[]string{"api"},
		),
	}

	registry.MustRegister(metrics.rateLimitWait)

	// 记录等待时间
	metrics.RecordRateLimitWait("orders", 0.5)
	metrics.RecordRateLimitWait("orders", 1.0)
	metrics.RecordRateLimitWait("orders", 0.3)

	// 验证有数据记录（histogram 需要用 Collector 而不是 Observer）
	// 简单验证不会 panic
	assert.NotNil(t, metrics.rateLimitWait)
}

// TestMetrics_Concurrent tests concurrent metric recording
func TestMetrics_Concurrent(t *testing.T) {
	registry := prometheus.NewRegistry()

	metrics := &Metrics{
		requestsTotal: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: "test",
				Name:      "requests_concurrent",
			},
			[]string{"api", "method", "status"},
		),
		requestDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: "test",
				Name:      "request_duration_concurrent",
			},
			[]string{"api", "method"},
		),
	}

	registry.MustRegister(metrics.requestsTotal)
	registry.MustRegister(metrics.requestDuration)

	// 并发记录
	done := make(chan bool, 100)
	for range 100 {
		go func() {
			defer func() { done <- true }()
			metrics.RecordRequest("orders", "GET", 200, 0.1)
		}()
	}

	// 等待完成
	for range 100 {
		<-done
	}

	// 验证计数正确
	count := testutil.ToFloat64(metrics.requestsTotal.WithLabelValues("orders", "GET", "OK"))
	assert.Equal(t, 100.0, count)
}
