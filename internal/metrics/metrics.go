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

// Package metrics 提供 SDK 的可观测性接口。
//
// 此包定义了 Metrics 接口，允许用户集成自己的监控系统（如 Prometheus、StatsD 等）。
// SDK 的各个组件会通过这个接口报告关键指标。
package metrics

import (
	"time"
)

// Recorder 定义指标记录器接口。
//
// 用户可以实现此接口以集成自己的监控系统。
// SDK 的各个组件将使用此接口记录关键指标。
//
// 示例:
//
//	type PrometheusRecorder struct { ... }
//
//	func (r *PrometheusRecorder) RecordCounter(name string, value float64, labels map[string]string) {
//	    // 实现 Prometheus counter 记录
//	}
type Recorder interface {
	// RecordCounter 记录计数器指标。
	//
	// 参数:
	//   - name: 指标名称
	//   - value: 增量值
	//   - labels: 标签（如 operation、region 等）
	RecordCounter(name string, value float64, labels map[string]string)

	// RecordGauge 记录仪表盘指标。
	//
	// 参数:
	//   - name: 指标名称
	//   - value: 当前值
	//   - labels: 标签
	RecordGauge(name string, value float64, labels map[string]string)

	// RecordHistogram 记录直方图指标（用于延迟分布等）。
	//
	// 参数:
	//   - name: 指标名称
	//   - value: 观测值
	//   - labels: 标签
	RecordHistogram(name string, value float64, labels map[string]string)

	// RecordTiming 记录时间指标。
	//
	// 参数:
	//   - name: 指标名称
	//   - duration: 持续时间
	//   - labels: 标签
	RecordTiming(name string, duration time.Duration, labels map[string]string)
}

// NoOpRecorder 是一个空操作的指标记录器。
//
// 当用户不需要指标收集时，可以使用此实现。
type NoOpRecorder struct{}

// RecordCounter 空操作。
func (n *NoOpRecorder) RecordCounter(name string, value float64, labels map[string]string) {}

// RecordGauge 空操作。
func (n *NoOpRecorder) RecordGauge(name string, value float64, labels map[string]string) {}

// RecordHistogram 空操作。
func (n *NoOpRecorder) RecordHistogram(name string, value float64, labels map[string]string) {}

// RecordTiming 空操作。
func (n *NoOpRecorder) RecordTiming(name string, duration time.Duration, labels map[string]string) {}

// DefaultRecorder 是默认的指标记录器（NoOp）。
var DefaultRecorder Recorder = &NoOpRecorder{}

// 预定义的指标名称常量。
const (
	// MetricRequestTotal 是请求总数指标。
	MetricRequestTotal = "spapi_request_total"

	// MetricRequestDuration 是请求延迟指标。
	MetricRequestDuration = "spapi_request_duration_seconds"

	// MetricRequestErrors 是请求错误数指标。
	MetricRequestErrors = "spapi_request_errors_total"

	// MetricAuthTokenRefresh 是令牌刷新次数指标。
	MetricAuthTokenRefresh = "spapi_auth_token_refresh_total"

	// MetricRateLimitWait 是速率限制等待时间指标。
	MetricRateLimitWait = "spapi_ratelimit_wait_seconds"

	// MetricRateLimitActive 是活跃的速率限制器数量指标。
	MetricRateLimitActive = "spapi_ratelimit_active_limiters"
)

// 预定义的标签键常量。
const (
	// LabelOperation 是 API 操作标签键。
	LabelOperation = "operation"

	// LabelRegion 是 AWS 区域标签键。
	LabelRegion = "region"

	// LabelMarketplace 是市场标签键。
	LabelMarketplace = "marketplace"

	// LabelStatusCode 是 HTTP 状态码标签键。
	LabelStatusCode = "status_code"

	// LabelErrorType 是错误类型标签键。
	LabelErrorType = "error_type"

	// LabelGrantType 是授权类型标签键（refresh_token / client_credentials）。
	LabelGrantType = "grant_type"
)
