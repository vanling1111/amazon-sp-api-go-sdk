# Metrics 集成指南

本文档介绍如何在 Amazon SP-API Go SDK 中集成和使用 Metrics（指标）功能。

---

## 📊 **概述**

SDK 提供了一个可选的 Metrics 接口，允许您集成自己的监控系统（如 Prometheus、StatsD、DataDog 等）。

默认情况下，SDK 使用 NoOp 记录器（不执行任何操作），不会影响性能。

---

## 🎯 **支持的指标**

### 1. **请求指标**

| 指标名称 | 类型 | 说明 |
|---------|------|------|
| `spapi_request_total` | Counter | 请求总数 |
| `spapi_request_duration_seconds` | Timing | 请求延迟 |
| `spapi_request_errors_total` | Counter | 请求错误数 |

### 2. **认证指标**

| 指标名称 | 类型 | 说明 |
|---------|------|------|
| `spapi_auth_token_refresh_total` | Counter | 令牌刷新次数 |

### 3. **速率限制指标**

| 指标名称 | 类型 | 说明 |
|---------|------|------|
| `spapi_ratelimit_wait_seconds` | Timing | 速率限制等待时间 |
| `spapi_ratelimit_active_limiters` | Gauge | 活跃的限制器数量 |

---

## 🔧 **实现自定义 Metrics Recorder**

### 示例 1: Prometheus 集成

```go
package main

import (
	"time"

	"github.com/vanling1111/amazon-sp-api-go-sdk/internal/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

// PrometheusRecorder 是 Prometheus 的指标记录器实现。
type PrometheusRecorder struct {
	requestTotal    *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
	requestErrors   *prometheus.CounterVec
	tokenRefresh    *prometheus.CounterVec
}

func NewPrometheusRecorder() *PrometheusRecorder {
	rec := &PrometheusRecorder{
		requestTotal: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "spapi_request_total",
				Help: "Total number of SP-API requests",
			},
			[]string{"operation", "status_code"},
		),
		requestDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "spapi_request_duration_seconds",
				Help:    "SP-API request duration in seconds",
				Buckets: prometheus.DefBuckets,
			},
			[]string{"operation"},
		),
		requestErrors: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "spapi_request_errors_total",
				Help: "Total number of SP-API request errors",
			},
			[]string{"operation", "error_type"},
		),
		tokenRefresh: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "spapi_auth_token_refresh_total",
				Help: "Total number of LWA token refreshes",
			},
			[]string{"grant_type"},
		),
	}

	// 注册所有指标
	prometheus.MustRegister(rec.requestTotal)
	prometheus.MustRegister(rec.requestDuration)
	prometheus.MustRegister(rec.requestErrors)
	prometheus.MustRegister(rec.tokenRefresh)

	return rec
}

func (r *PrometheusRecorder) RecordCounter(name string, value float64, labels map[string]string) {
	switch name {
	case metrics.MetricRequestTotal:
		r.requestTotal.With(prometheus.Labels(labels)).Add(value)
	case metrics.MetricRequestErrors:
		r.requestErrors.With(prometheus.Labels(labels)).Add(value)
	case metrics.MetricAuthTokenRefresh:
		r.tokenRefresh.With(prometheus.Labels(labels)).Add(value)
	}
}

func (r *PrometheusRecorder) RecordGauge(name string, value float64, labels map[string]string) {
	// 实现 Gauge 指标
}

func (r *PrometheusRecorder) RecordHistogram(name string, value float64, labels map[string]string) {
	// 实现 Histogram 指标
}

func (r *PrometheusRecorder) RecordTiming(name string, duration time.Duration, labels map[string]string) {
	switch name {
	case metrics.MetricRequestDuration:
		r.requestDuration.With(prometheus.Labels(labels)).Observe(duration.Seconds())
	}
}
```

### 示例 2: StatsD 集成

```go
package main

import (
	"time"

	"github.com/vanling1111/amazon-sp-api-go-sdk/internal/metrics"
	"github.com/cactus/go-statsd-client/statsd"
)

// StatsDRecorder 是 StatsD 的指标记录器实现。
type StatsDRecorder struct {
	client statsd.Statter
}

func NewStatsDRecorder(addr string) (*StatsDRecorder, error) {
	client, err := statsd.NewClient(addr, "spapi")
	if err != nil {
		return nil, err
	}
	return &StatsDRecorder{client: client}, nil
}

func (r *StatsDRecorder) RecordCounter(name string, value float64, labels map[string]string) {
	// StatsD 不直接支持标签，可以将标签拼接到指标名称中
	r.client.Inc(name, int64(value), 1.0)
}

func (r *StatsDRecorder) RecordGauge(name string, value float64, labels map[string]string) {
	r.client.Gauge(name, int64(value), 1.0)
}

func (r *StatsDRecorder) RecordHistogram(name string, value float64, labels map[string]string) {
	r.client.Gauge(name, int64(value), 1.0)
}

func (r *StatsDRecorder) RecordTiming(name string, duration time.Duration, labels map[string]string) {
	r.client.Timing(name, int64(duration.Milliseconds()), 1.0)
}
```

---

## 🚀 **集成到 SDK**

### 在 Transport 客户端中使用

```go
package main

import (
	"context"
	"log"

	"github.com/vanling1111/amazon-sp-api-go-sdk/internal/transport"
)

func main() {
	// 1. 创建 Prometheus Recorder
	promRecorder := NewPrometheusRecorder()

	// 2. 创建 Transport 客户端
	client := transport.NewClient(
		"https://sellingpartnerapi-na.amazon.com",
		nil, // 使用默认配置
	)

	// 3. 设置 Metrics Recorder
	client.SetMetrics(promRecorder)

	// 现在所有请求都会自动记录指标
	// ...
}
```

---

## 📈 **查看指标**

### Prometheus 示例

启动 Prometheus 服务器并配置抓取端点：

```yaml
# prometheus.yml
scrape_configs:
  - job_name: 'spapi_sdk'
    static_configs:
      - targets: ['localhost:9090']
```

### 查询示例

```promql
# 总请求数
spapi_request_total

# 按操作分组的请求数
sum by (operation) (spapi_request_total)

# 请求延迟 P95
histogram_quantile(0.95, spapi_request_duration_seconds)

# 错误率
rate(spapi_request_errors_total[5m]) / rate(spapi_request_total[5m])
```

---

## 💡 **最佳实践**

1. **异步记录**: Metrics 记录应该是异步的，避免阻塞主流程
2. **标签基数**: 避免使用高基数标签（如 request_id），防止指标爆炸
3. **采样**: 对于高频 API，考虑采样记录（如每 10 次请求记录 1 次）
4. **性能**: NoOp 记录器性能开销接近零，可以安全使用

---

## 🔒 **禁用 Metrics**

默认情况下，SDK 使用 NoOp 记录器（无性能开销）。

如果需要明确禁用：

```go
client.SetMetrics(&metrics.NoOpRecorder{})
```

---

## 📊 **完整示例**

```go
package main

import (
	"context"
	"log"
	"net/http"

	"github.com/vanling1111/amazon-sp-api-go-sdk/internal/auth"
	"github.com/vanling1111/amazon-sp-api-go-sdk/internal/transport"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// 1. 创建 Prometheus Recorder
	promRecorder := NewPrometheusRecorder()

	// 2. 启动 Prometheus HTTP 服务器
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Fatal(http.ListenAndServe(":9090", nil))
	}()

	// 3. 创建 Transport 客户端
	client := transport.NewClient(
		"https://sellingpartnerapi-na.amazon.com",
		nil,
	)
	client.SetMetrics(promRecorder)

	// 4. 创建 Auth 客户端
	creds, _ := auth.NewCredentials(
		"client-id",
		"client-secret",
		"refresh-token",
		"https://api.amazon.com/auth/o2/token",
	)
	authClient := auth.NewClient(creds)

	// 5. 发送请求（自动记录指标）
	ctx := context.Background()
	req, _ := http.NewRequestWithContext(ctx, "GET", "/orders/v0/orders", nil)
	
	// 获取访问令牌
	token, _ := authClient.GetAccessToken(ctx)
	req.Header.Set("x-amz-access-token", token.AccessToken)

	// 发送请求（自动记录 metrics）
	resp, err := client.Do(ctx, req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	log.Printf("Request completed: %s", resp.Status)
	// 访问 http://localhost:9090/metrics 查看指标
}
```

---

## 🎯 **总结**

- ✅ SDK 提供了灵活的 Metrics 接口
- ✅ 默认使用 NoOp 记录器（零性能开销）
- ✅ 可以轻松集成 Prometheus、StatsD 等监控系统
- ✅ 自动记录请求、认证、速率限制等关键指标

---

**更多信息**: 参考 `internal/metrics/metrics.go` 中的接口定义。

