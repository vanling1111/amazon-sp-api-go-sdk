# 指标监控指南

## 概述

本 SDK 提供了内置的指标监控接口。

## 指标接口

```go
// Recorder 定义指标记录接口
type Recorder interface {
    RecordRequest(operation string, duration time.Duration, err error)
    RecordAuth(grantType string, duration time.Duration, success bool)
    RecordRateLimit(operation string, waitDuration time.Duration)
}
```

## 内置实现

### NoOpRecorder
默认实现，不记录任何指标。

### 自定义实现

您可以实现自己的 Recorder：

```go
type MyRecorder struct {
    // 您的指标存储
}

func (r *MyRecorder) RecordRequest(operation string, duration time.Duration, err error) {
    // 记录请求指标
}

func (r *MyRecorder) RecordAuth(grantType string, duration time.Duration, success bool) {
    // 记录认证指标
}

func (r *MyRecorder) RecordRateLimit(operation string, waitDuration time.Duration) {
    // 记录速率限制指标
}
```

## 使用方法

```go
// 创建自定义 recorder
recorder := NewMyRecorder()

// 创建客户端时注入
client, err := spapi.NewClient(
    spapi.WithRegion(models.RegionNA),
    spapi.WithCredentials(...),
    spapi.WithMetrics(recorder),
)
```

## 可记录的指标

### 请求指标
- 操作名称
- 请求时长
- 成功/失败

### 认证指标
- Grant Type
- 认证时长
- 成功/失败

### 速率限制指标
- 操作名称
- 等待时长

## 工具

项目提供了监控工具示例：
- `tools/monitoring/metrics.go` - 指标收集器
- `tools/performance/profiler.go` - 性能分析器

## 参考

- [Go Metrics](https://pkg.go.dev/runtime/metrics)
- [Prometheus](https://prometheus.io/)

