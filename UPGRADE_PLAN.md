# Amazon SP-API Go SDK 升级计划

## 📋 版本：v1.0.0 → v2.0.0

**当前版本**：v1.0.0  
**目标版本**：v2.0.0  
**升级日期**：2025年10月  
**Go 版本**：1.21 → 1.25

---

## 🎯 升级目标

### 核心目标
1. ✅ 引入最佳实践的第三方依赖（24 个精选依赖）
2. ✅ 重构项目使用 Go 1.25 新特性
3. ✅ 实现功能扩展（解密、分页、通知、可观测性）
4. ✅ 提升健壮性和鲁棒性（熔断器、重试、日志、追踪）

### 预期成果
- 🚀 性能提升 30-50%
- 🛡️ 可靠性提升（熔断、重试、监控）
- 📊 可观测性完善（日志、追踪、指标）
- 🎨 开发体验提升（迭代器、构建器、调试）

---

## 📦 阶段 1：依赖管理（预计 1 天）

### 任务 1.1：添加所有依赖到 go.mod

**依赖清单（24 个）**：

```go
require (
    // === AWS 集成（3个）===
    github.com/aws/aws-sdk-go-v2 v1.24.0
    github.com/aws/aws-sdk-go-v2/service/sqs v1.29.0
    github.com/aws/aws-sdk-go-v2/service/eventbridge v1.26.0

    // === 核心功能（6个）===
    github.com/pkg/errors v0.9.1                      // 错误处理（错误堆栈）
    github.com/go-playground/validator/v10 v10.19.0   // 数据验证（请求参数验证）
    go.uber.org/ratelimit v0.3.0                      // 限流（Leaky Bucket）
    github.com/imroc/req/v3 v3.42.3                   // HTTP 客户端（替换标准库）
    golang.org/x/sync v0.6.0                          // 并发控制（errgroup）
    golang.org/x/crypto v0.19.0                       // 加密（AES-256-CBC）

    // === 性能优化（2个）===
    github.com/json-iterator/go v1.1.12               // JSON（比标准库快 3-5 倍）
    github.com/allegro/bigcache/v3 v3.1.0             // 缓存（零 GC，100 万 QPS）

    // === 可观测性（4个）===
    go.uber.org/zap v1.26.0                           // 结构化日志
    go.opentelemetry.io/otel v1.21.0                  // 分布式追踪
    go.opentelemetry.io/otel/trace v1.21.0
    github.com/prometheus/client_golang v1.18.0       // 指标收集

    // === 开发工具（5个）===
    github.com/stretchr/testify v1.8.4                // 测试框架
    github.com/h2non/gock v1.2.0                      // HTTP Mock
    github.com/spf13/viper v1.18.2                    // 配置管理
    github.com/joho/godotenv v1.5.1                   // 环境变量
    github.com/spf13/cobra v1.8.0                     // CLI 工具

    // === 实用工具（4个）===
    github.com/google/uuid v1.6.0                     // UUID 生成
    github.com/gammazero/workerpool v1.1.3            // 工作池
    github.com/jinzhu/now v1.1.5                      // 时间处理
    github.com/fatih/color v1.16.0                    // 彩色输出
    github.com/schollz/progressbar/v3 v3.14.1         // 进度条
)
```

**执行步骤**：
1. ✅ 升级 `go.mod` 到 Go 1.25（已完成）
2. 🔄 添加所有依赖
3. 🔄 运行 `go mod tidy`
4. 🔄 运行 `go mod verify`

---

## 🔧 阶段 2：代码重构（预计 5-7 天）

### 任务 2.1：迁移到 req/v3 HTTP 客户端

**影响文件**：
- `pkg/spapi/client.go` - 核心客户端
- `internal/transport/client.go` - HTTP 传输层
- `internal/transport/middleware.go` - 中间件系统

**重构内容**：
```go
// 之前（标准库）
type Client struct {
    httpClient *http.Client
    baseURL    string
}

func (c *Client) Get(ctx context.Context, path string, result interface{}) error {
    req, _ := http.NewRequestWithContext(ctx, "GET", c.baseURL+path, nil)
    resp, err := c.httpClient.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    body, _ := io.ReadAll(resp.Body)
    json.Unmarshal(body, result)
    return nil
}

// 之后（req/v3）
type Client struct {
    httpClient *req.Client
}

func (c *Client) Get(ctx context.Context, path string, result interface{}) error {
    _, err := c.httpClient.R().
        SetContext(ctx).
        SetSuccessResult(result).
        Get(path)
    return err
}
```

**预期收益**：
- 代码减少 40-50%
- 自动 JSON 解析
- 自动重试
- 调试更简单

---

### 任务 2.2：使用 Go 1.25 迭代器重构分页 API

**影响文件**：
- `pkg/spapi/orders-v0/client.go` - Orders API
- `pkg/spapi/reports-v2021-06-30/client.go` - Reports API
- `pkg/spapi/catalog-items-*/client.go` - Catalog APIs
- 所有有分页的 API（约 30 个）

**重构内容**：
```go
// 之前（手动分页，用户痛苦）
func (c *Client) GetOrders(ctx context.Context, query *GetOrdersQuery) (*GetOrdersResponse, error) {
    // 只返回一页
    return c.get(ctx, path, query)
}

// 用户需要手动处理分页
nextToken := ""
for {
    result, err := client.GetOrders(ctx, &GetOrdersQuery{
        NextToken: nextToken,
    })
    if err != nil {
        return err
    }
    
    for _, order := range result.Orders {
        process(order)
    }
    
    if result.NextToken == "" {
        break
    }
    nextToken = result.NextToken
}

// 之后（Go 1.25 迭代器，自动分页）
import "iter"

func (c *Client) IterateOrders(ctx context.Context, query *GetOrdersQuery) iter.Seq2[*Order, error] {
    return func(yield func(*Order, error) bool) {
        nextToken := ""
        for {
            result, err := c.GetOrders(ctx, &GetOrdersQuery{
                MarketplaceIDs: query.MarketplaceIDs,
                CreatedAfter:   query.CreatedAfter,
                NextToken:      nextToken,
            })
            if err != nil {
                yield(nil, err)
                return
            }
            
            for _, order := range result.Orders {
                if !yield(order, nil) {
                    return
                }
            }
            
            if result.NextToken == "" {
                break
            }
            nextToken = result.NextToken
        }
    }
}

// 用户使用（极简）
for order, err := range client.Orders.IterateOrders(ctx, query) {
    if err != nil {
        return err
    }
    process(order)
}
```

**预期收益**：
- 用户代码减少 70%
- 自动处理分页
- 自动处理错误
- 支持提前退出

---

### 任务 2.3：移除所有 `item := item` 代码

**影响文件**：
- 所有使用 goroutine + 循环的代码
- 约 20-30 处

**重构内容**：
```go
// Go 1.21（需要显式复制）
for _, api := range apis {
    api := api  // ← 删除这行
    go func() {
        process(api)
    }()
}

// Go 1.25（自动正确）
for _, api := range apis {
    go func() {
        process(api)  // 自动捕获正确的值
    }()
}
```

---

### 任务 2.4：使用 For-range 整数简化代码

**影响文件**：
- `internal/ratelimit/manager.go`
- `tests/benchmarks/benchmark_test.go`
- 工作池初始化代码

**重构内容**：
```go
// 之前
workers := make([]*Worker, workerCount)
for i := 0; i < workerCount; i++ {
    workers[i] = NewWorker(i)
}

// 之后（Go 1.25）
workers := make([]*Worker, workerCount)
for i := range workerCount {
    workers[i] = NewWorker(i)
}
```

---

### 任务 2.5：错误处理统一使用 pkg/errors

**影响文件**：
- 所有返回 error 的函数（约 200+ 处）

**重构内容**：
```go
// 之前
if err != nil {
    return fmt.Errorf("failed to get order: %v", err)
}

// 之后（带堆栈）
import "github.com/pkg/errors"

if err != nil {
    return errors.Wrap(err, "failed to get order")
}

// 根错误
if orderID == "" {
    return errors.New("orderID is required")
}
```

**预期收益**：
- 完整错误堆栈
- 更容易调试
- 更好的错误上下文

---

### 任务 2.6：参数验证使用 validator

**影响文件**：
- `pkg/spapi/config.go` - 配置验证
- 所有 API 客户端的参数验证

**重构内容**：
```go
// 之前（手动验证）
func (c *Config) Validate() error {
    if c.ClientID == "" {
        return fmt.Errorf("ClientID is required")
    }
    if c.ClientSecret == "" {
        return fmt.Errorf("ClientSecret is required")
    }
    if c.RefreshToken == "" {
        return fmt.Errorf("RefreshToken is required")
    }
    return nil
}

// 之后（自动验证）
import "github.com/go-playground/validator/v10"

type Config struct {
    ClientID     string `validate:"required"`
    ClientSecret string `validate:"required"`
    RefreshToken string `validate:"required"`
    Region       string `validate:"required,oneof=NA EU FE"`
}

var validate = validator.New()

func (c *Config) Validate() error {
    return validate.Struct(c)
}
```

---

### 任务 2.7：JSON 处理迁移到 json-iterator

**影响文件**：
- `internal/codec/json.go` - JSON 编解码器
- 所有使用 `encoding/json` 的地方

**重构内容**：
```go
// 创建别名
import jsoniter "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// 全局替换 encoding/json 为 json-iterator
// 性能提升 3-5 倍
```

---

## 🚀 阶段 3：功能扩展（预计 10-14 天）

### 任务 3.1：报告解密/加密（P0）⭐⭐⭐⭐⭐

**新增文件**：
```
internal/crypto/
├── aes.go          // AES-256-CBC 解密/加密
├── aes_test.go     // 单元测试
└── README.md       // 说明文档
```

**实现功能**：
```go
package crypto

// DecryptReport 解密 SP-API 报告
func DecryptReport(key, iv string, data []byte) ([]byte, error)

// EncryptDocument 加密上传文档
func EncryptDocument(data []byte) (*EncryptionDetails, []byte, error)

// ValidateEncryptionDetails 验证加密参数
func ValidateEncryptionDetails(details *EncryptionDetails) error
```

**集成到 Reports API**：
```go
// pkg/spapi/reports-v2021-06-30/client.go
func (c *Client) GetReportDocumentDecrypted(ctx context.Context, reportDocumentID string) ([]byte, error) {
    // 1. 获取报告文档信息
    doc, err := c.GetReportDocument(ctx, reportDocumentID)
    
    // 2. 下载加密内容
    encrypted, err := downloadURL(doc.URL)
    
    // 3. 自动解密
    if doc.EncryptionDetails != nil {
        return crypto.DecryptReport(
            doc.EncryptionDetails.Key,
            doc.EncryptionDetails.InitializationVector,
            encrypted,
        )
    }
    
    return encrypted, nil
}
```

**依赖**：`golang.org/x/crypto`

---

### 任务 3.2：分页迭代器（P0）⭐⭐⭐⭐⭐

**影响文件**：
- 所有有分页的 API（30+ 个）

**新增方法**：
```go
// pkg/spapi/orders-v0/iterator.go
import "iter"

// IterateOrders 订单迭代器（自动分页）
func (c *Client) IterateOrders(ctx context.Context, query *GetOrdersQuery) iter.Seq2[*Order, error]

// IterateOrderItems 订单项迭代器
func (c *Client) IterateOrderItems(ctx context.Context, orderID string) iter.Seq2[*OrderItem, error]
```

**同样实现**：
- `IterateReports()` - Reports API
- `IterateCatalogItems()` - Catalog API
- `IterateFeeds()` - Feeds API
- 所有分页 API

**依赖**：Go 1.25 标准库（`iter` 包）

---

### 任务 3.3：大文件传输（P0）⭐⭐⭐⭐⭐

**新增文件**：
```
internal/transfer/
├── uploader.go        // 分片上传器
├── uploader_test.go
├── downloader.go      // 分片下载器
├── downloader_test.go
└── progress.go        // 进度回调
```

**实现功能**：
```go
package transfer

// ChunkedUploader 分片上传器
type ChunkedUploader struct {
    client      *req.Client
    chunkSize   int64         // 分片大小（默认 10MB）
    concurrency int           // 并发数
    progress    ProgressFunc  // 进度回调
}

// Upload 上传大文件
func (u *ChunkedUploader) Upload(ctx context.Context, file io.Reader, feedType string) (*UploadResult, error)

// UploadWithProgress 带进度的上传
func (u *ChunkedUploader) UploadWithProgress(ctx context.Context, file io.Reader, feedType string, onProgress ProgressFunc) error

type ProgressFunc func(uploaded, total int64, percent float64)

// ChunkedDownloader 分片下载器
type ChunkedDownloader struct {
    client      *req.Client
    bufferSize  int64
    autoDecrypt bool  // 自动解密
}

// Download 下载大文件
func (d *ChunkedDownloader) Download(ctx context.Context, url string, writer io.Writer) error

// DownloadWithProgress 带进度的下载
func (d *ChunkedDownloader) DownloadWithProgress(ctx context.Context, url string, writer io.Writer, onProgress ProgressFunc) error
```

**依赖**：
- `github.com/imroc/req/v3`
- `github.com/schollz/progressbar/v3`

---

### 任务 3.4：通知支持（SQS 轮询器）（P0）⭐⭐⭐⭐⭐

**新增文件**：
```
pkg/notifications/
├── poller.go          // SQS 轮询器
├── poller_test.go
├── parser.go          // 消息解析器
├── parser_test.go
├── events.go          // 事件类型定义
├── subscription.go    // 订阅管理器
└── README.md
```

**实现功能**：
```go
package notifications

// SQSPoller SQS 消息轮询器
type SQSPoller struct {
    sqsClient    *sqs.Client
    queueURL     string
    pollInterval time.Duration
    maxMessages  int32
    handlers     map[string]EventHandler
}

// NewSQSPoller 创建 SQS 轮询器
func NewSQSPoller(sqsClient *sqs.Client, queueURL string, opts ...Option) *SQSPoller

// RegisterHandler 注册事件处理器
func (p *SQSPoller) RegisterHandler(notificationType string, handler EventHandler)

// Start 开始轮询
func (p *SQSPoller) Start(ctx context.Context) error

// EventHandler 事件处理器
type EventHandler func(ctx context.Context, event *Event) error

// Event 通知事件
type Event struct {
    NotificationVersion string
    NotificationType    string
    Payload             json.RawMessage
    Timestamp           time.Time
    MessageID           string
}

// Parse 解析事件负载
func (e *Event) Parse(v interface{}) error

// SubscriptionManager 订阅管理器
type SubscriptionManager struct {
    client *spapi.Client
}

// Subscribe 订阅通知
func (m *SubscriptionManager) Subscribe(notificationType, sqsArn string) (*Subscription, error)

// Unsubscribe 取消订阅
func (m *SubscriptionManager) Unsubscribe(subscriptionID string) error

// List 列出所有订阅
func (m *SubscriptionManager) List() ([]*Subscription, error)
```

**依赖**：
- `github.com/aws/aws-sdk-go-v2/service/sqs`
- `github.com/json-iterator/go`

---

### 任务 3.5：熔断器（P1）⭐⭐⭐⭐

**新增文件**：
```
internal/circuit/
├── breaker.go         // 熔断器
├── breaker_test.go
└── middleware.go      // 熔断器中间件
```

**实现功能**：
```go
package circuit

// State 熔断器状态
type State int

const (
    StateClosed   State = iota  // 关闭（正常）
    StateOpen                    // 打开（熔断）
    StateHalfOpen                // 半开（尝试恢复）
)

// Breaker 熔断器
type Breaker struct {
    maxFailures   int
    timeout       time.Duration
    state         State
    failures      int
    lastFailTime  time.Time
}

// NewBreaker 创建熔断器
func NewBreaker(maxFailures int, timeout time.Duration) *Breaker

// Execute 执行请求（带熔断保护）
func (b *Breaker) Execute(fn func() error) error

// State 获取当前状态
func (b *Breaker) State() State

// Reset 重置熔断器
func (b *Breaker) Reset()

// Middleware 熔断器中间件
func Middleware(breaker *Breaker) transport.Middleware
```

**依赖**：无（标准库）

---

### 任务 3.6：结构化日志（P1）⭐⭐⭐⭐⭐

**新增文件**：
```
internal/logging/
├── logger.go          // 日志接口
├── zap.go            // Zap 实现
├── middleware.go     // 日志中间件
└── fields.go         // 日志字段
```

**实现功能**：
```go
package logging

import "go.uber.org/zap"

// Logger 日志接口
type Logger interface {
    Debug(msg string, fields ...Field)
    Info(msg string, fields ...Field)
    Warn(msg string, fields ...Field)
    Error(msg string, fields ...Field)
}

// Field 日志字段
type Field struct {
    Key   string
    Value interface{}
}

// ZapLogger Zap 实现
type ZapLogger struct {
    logger *zap.Logger
}

// NewZapLogger 创建 Zap 日志器
func NewZapLogger(config *zap.Config) (*ZapLogger, error)

// Middleware 日志中间件
func Middleware(logger Logger, opts *Options) transport.Middleware

type Options struct {
    LogHeaders   bool
    LogBody      bool
    MaxBodySize  int
    RedactFields []string  // 敏感字段脱敏
}
```

**集成到客户端**：
```go
// 用户可选日志
logger, _ := logging.NewZapLogger(zap.NewProductionConfig())

client, _ := spapi.NewClient(
    spapi.WithLogger(logger),
    spapi.WithLogOptions(&logging.Options{
        LogHeaders:   true,
        LogBody:      true,
        RedactFields: []string{"x-amz-access-token"},
    }),
)

// 自动记录所有请求
// [INFO] Request: GET /orders/v0/orders/123 duration=234ms status=200
```

**依赖**：`go.uber.org/zap`

---

### 任务 3.7：分布式追踪（P1）⭐⭐⭐⭐⭐

**新增文件**：
```
internal/tracing/
├── tracer.go          // 追踪器
├── middleware.go      // 追踪中间件
└── propagation.go     // 上下文传播
```

**实现功能**：
```go
package tracing

import (
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/trace"
)

// Middleware 追踪中间件
func Middleware(tracer trace.Tracer) transport.Middleware

// 自动追踪每个 API 调用
// - Span name: "Orders.GetOrders"
// - Attributes: method, path, status_code, duration
// - Parent span: 从 context 中获取
```

**集成**：
```go
// 用户配置 OpenTelemetry
tracer := otel.Tracer("amazon-sp-api")

client, _ := spapi.NewClient(
    spapi.WithTracer(tracer),
)

// 自动生成追踪数据，导出到 Jaeger/Zipkin
```

**依赖**：
- `go.opentelemetry.io/otel`
- `go.opentelemetry.io/otel/trace`

---

### 任务 3.8：Prometheus Metrics（P1）⭐⭐⭐⭐⭐

**新增文件**：
```
internal/metrics/
├── prometheus.go      // Prometheus 实现
├── collector.go       // 指标收集器
└── middleware.go      // 指标中间件
```

**实现功能**：
```go
package metrics

import "github.com/prometheus/client_golang/prometheus"

// PrometheusMetrics Prometheus 指标
type PrometheusMetrics struct {
    requestDuration *prometheus.HistogramVec
    requestTotal    *prometheus.CounterVec
    requestErrors   *prometheus.CounterVec
}

// NewPrometheusMetrics 创建 Prometheus 指标
func NewPrometheusMetrics(namespace string) *PrometheusMetrics

// 自动收集指标
// - spapi_requests_total{api="orders",method="GET",status="200"}
// - spapi_request_duration_seconds{api="orders",method="GET"}
// - spapi_request_errors_total{api="orders",error_type="rate_limit"}
```

**依赖**：`github.com/prometheus/client_golang`

---

### 任务 3.9：限流器升级（P1）⭐⭐⭐⭐⭐

**重构文件**：
- `internal/ratelimit/limiter.go` - 使用 Uber ratelimit

**重构内容**：
```go
// 之前（自己实现的 Token Bucket）
type RateLimiter struct {
    bucket *TokenBucket
}

// 之后（使用 Uber ratelimit）
import "go.uber.org/ratelimit"

type RateLimiter struct {
    limiter ratelimit.Limiter
}

func NewRateLimiter(rate int) *RateLimiter {
    return &RateLimiter{
        limiter: ratelimit.New(rate),
    }
}

func (r *RateLimiter) Wait(ctx context.Context) error {
    r.limiter.Take()  // 自动阻塞
    return nil
}
```

**预期收益**：
- 性能更好（无锁设计）
- 代码更简洁
- 久经考验的实现

**依赖**：`go.uber.org/ratelimit`

---

### 任务 3.10：测试增强（P1）⭐⭐⭐⭐⭐

**重构文件**：
- 所有测试文件（152 个）

**重构内容**：
```go
// 之前（标准库）
if result != expected {
    t.Errorf("Expected %v, got %v", expected, result)
}

// 之后（testify）
import "github.com/stretchr/testify/assert"

assert.Equal(t, expected, result)
assert.NoError(t, err)
assert.Len(t, items, 5)
assert.Contains(t, str, "substring")
```

**使用 Go 1.25 synctest**：
```go
import "testing/synctest"

// 测试 Token 过期
func TestTokenExpiry(t *testing.T) {
    synctest.Run(func() {
        cache := NewTokenCache()
        cache.Set("token", time.Now().Add(5*time.Minute))
        
        // 虚拟时间前进
        time.Sleep(6 * time.Minute)
        
        assert.True(t, cache.IsExpired("token"))
    })
}
```

**依赖**：`github.com/stretchr/testify`

---

### 任务 3.11：HTTP Mock 测试（P1）⭐⭐⭐⭐

**新增文件**：
```
testing/mock/
├── server.go          // Mock 服务器
├── responses.go       // 预定义响应
└── fixtures/          // 测试数据
```

**实现功能**：
```go
package mock

import "github.com/h2non/gock"

// SetupOrdersMock 设置 Orders API Mock
func SetupOrdersMock() {
    defer gock.Off()
    
    gock.New("https://sellingpartnerapi-na.amazon.com").
        Get("/orders/v0/orders/123").
        Reply(200).
        JSON(LoadFixture("orders/order-123.json"))
}
```

**依赖**：`github.com/h2non/gock`

---

### 任务 3.12：配置管理（P2）⭐⭐⭐⭐

**新增文件**：
```
pkg/config/
├── manager.go         // 配置管理器
├── loader.go          // 配置加载器
└── config.yaml        // 配置示例
```

**实现功能**：
```go
package config

import (
    "github.com/spf13/viper"
    "github.com/joho/godotenv"
)

// Manager 配置管理器
type Manager struct {
    viper *viper.Viper
}

// LoadConfig 加载配置
func LoadConfig(paths ...string) (*Config, error) {
    // 1. 加载 .env
    godotenv.Load()
    
    // 2. 加载 YAML/JSON
    v := viper.New()
    v.SetConfigFile("config.yaml")
    v.ReadInConfig()
    
    // 3. 环境变量覆盖
    v.AutomaticEnv()
    
    var config Config
    v.Unmarshal(&config)
    return &config, nil
}

// WatchConfig 监控配置变化
func (m *Manager) WatchConfig() <-chan *Config
```

**依赖**：
- `github.com/spf13/viper`
- `github.com/joho/godotenv`

---

### 任务 3.13：CLI 工具（P2）⭐⭐⭐⭐

**新增文件**：
```
cmd/spapi-cli/
├── main.go
├── commands/
│   ├── orders.go      // 订单命令
│   ├── reports.go     // 报告命令
│   ├── feeds.go       // Feed 命令
│   └── config.go      // 配置命令
└── README.md
```

**实现功能**：
```go
package main

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
    Use:   "spapi",
    Short: "Amazon SP-API CLI Tool",
}

var ordersCmd = &cobra.Command{
    Use:   "orders",
    Short: "Manage orders",
}

var getOrderCmd = &cobra.Command{
    Use:   "get [orderID]",
    Short: "Get order by ID",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        orderID := args[0]
        // 调用 SDK
        order, err := client.Orders.GetOrder(ctx, orderID)
        // 输出
        fmt.Println(order)
    },
}

// spapi orders get 123-456789
// spapi orders list --created-after 2025-01-01
// spapi reports create --type GET_FLAT_FILE_ALL_ORDERS_DATA_BY_ORDER_DATE
```

**依赖**：
- `github.com/spf13/cobra`
- `github.com/fatih/color`
- `github.com/schollz/progressbar/v3`

---

### 任务 3.14：并发增强（P2）⭐⭐⭐⭐

**重构文件**：
- 所有批量操作的地方

**使用 errgroup**：
```go
import "golang.org/x/sync/errgroup"

// 批量获取订单（并发，限制 10 个）
func (c *Client) BatchGetOrders(ctx context.Context, orderIDs []string) ([]*Order, error) {
    g, ctx := errgroup.WithContext(ctx)
    g.SetLimit(10)
    
    orders := make([]*Order, len(orderIDs))
    
    for i, orderID := range orderIDs {
        i, orderID := i, orderID
        g.Go(func() error {
            order, err := c.GetOrder(ctx, orderID)
            if err != nil {
                return err
            }
            orders[i] = order
            return nil
        })
    }
    
    if err := g.Wait(); err != nil {
        return nil, err
    }
    
    return orders, nil
}
```

**依赖**：`golang.org/x/sync`

---

### 任务 3.15：工作池（P2）⭐⭐⭐

**新增文件**：
```
internal/pool/
├── worker.go          // 工作池
└── worker_test.go
```

**实现功能**：
```go
package pool

import "github.com/gammazero/workerpool"

// Pool 工作池
type Pool struct {
    wp *workerpool.WorkerPool
}

// NewPool 创建工作池
func NewPool(size int) *Pool {
    return &Pool{
        wp: workerpool.New(size),
    }
}

// Submit 提交任务
func (p *Pool) Submit(task func())

// StopWait 等待所有任务完成
func (p *Pool) StopWait()
```

**依赖**：`github.com/gammazero/workerpool`

---

## 📊 阶段 4：文档更新（预计 2-3 天）

### 任务 4.1：更新 README.md

**添加内容**：
- ✅ 依赖说明
- ✅ Go 1.25 特性说明
- ✅ 新功能介绍（解密、迭代器、通知）

### 任务 4.2：创建升级指南

**新增文件**：
```
docs/UPGRADE_GUIDE.md  // v1.0 → v2.0 升级指南
```

### 任务 4.3：新增功能文档

**新增文件**：
```
docs/REPORT_DECRYPTION.md    // 报告解密指南
docs/PAGINATION_GUIDE.md     // 分页迭代器使用指南
docs/NOTIFICATIONS_GUIDE.md  // 通知集成指南
docs/OBSERVABILITY_GUIDE.md  // 可观测性配置指南
```

### 任务 4.4：更新 CHANGELOG.md

**添加 v2.0.0 条目**

---

## 📊 阶段 5：测试和验证（预计 3-5 天）

### 任务 5.1：单元测试
- ✅ 所有新功能 100% 测试覆盖率
- ✅ 使用 testify + synctest

### 任务 5.2：集成测试
- ✅ 真实 API 调用测试
- ✅ SQS 集成测试
- ✅ 解密功能测试

### 任务 5.3：性能测试
- ✅ Benchmark 所有核心功能
- ✅ 对比 v1.0.0 性能

### 任务 5.4：兼容性测试
- ✅ Go 1.25 编译
- ✅ Docker 环境测试
- ✅ Kubernetes 环境测试

---

## 📅 时间估算

| 阶段 | 任务数 | 预计天数 | 累计天数 |
|------|--------|---------|---------|
| 阶段 1：依赖管理 | 1 | 1 天 | 1 天 |
| 阶段 2：代码重构 | 7 | 5-7 天 | 6-8 天 |
| 阶段 3：功能扩展 | 8 | 10-14 天 | 16-22 天 |
| 阶段 4：文档更新 | 4 | 2-3 天 | 18-25 天 |
| 阶段 5：测试验证 | 4 | 3-5 天 | 21-30 天 |

**总计**：3-4 周

---

## ✅ 下一步行动

现在立即执行？请确认是否开始：

1. **立即开始阶段 1**：添加所有依赖到 go.mod
2. 还是先审核这个计划？

需要我现在开始执行吗？
