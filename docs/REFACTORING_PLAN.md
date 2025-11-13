# SDK 长期重构优化计划

## 目标
符合Amazon SP-API官方要求，打造生产级Go SDK

## 官方要求对照

| 要求 | 当前状态 | 目标状态 |
|------|---------|---------|
| OpenAPI规范 | ✅ 已实现 | ✅ 保持 |
| 区域支持 | ✅ NA/EU/FE | ✅ 保持 |
| LWA认证 | ✅ 已实现 | ✅ 优化 |
| 速率限制 | ✅ Token Bucket | ✅ 保持 |
| 错误处理 | ✅ 已实现 | 🔄 标准化 |
| Marketplace | ✅ 已实现 | 🔄 公开API |
| RDT支持 | ✅ 已实现 | ✅ 保持 |
| Sandbox | ⚠️ 部分支持 | 🔄 完善 |

---

## 阶段1: 核心重构（P0 - 本周完成）

### 1.1 统一类型定义
**问题**: Region和Marketplace定义重复

**方案**:
```go
// 移除 internal/models.Region
// 统一使用 pkg/spapi.Region

pkg/spapi/
  ├── region.go       // Region定义和常量
  ├── marketplace.go  // Marketplace定义和常量（新增公开）
  └── types.go        // 其他公开类型
```

**影响**:
- ✅ 简化架构
- ✅ 无转换开销
- ✅ 易于维护
- ⚠️ Breaking Change（需要v2.0.0）

---

### 1.2 清理internal/models
**当前**:
```go
internal/models/
  └── common.go
      ├── Region          // 删除（已在pkg/spapi）
      ├── Marketplace     // 移动到pkg/spapi
      ├── RateLimitInfo   // 保留（内部使用）
      ├── RequestMetadata // 保留（内部使用）
      └── ErrorDetail     // 保留（内部使用）
```

**重构后**:
```go
internal/models/
  └── internal.go       // 重命名，明确内部使用
      ├── RateLimitInfo
      ├── RequestMetadata
      └── ErrorDetail

pkg/spapi/
  ├── region.go         // Region + 常量
  └── marketplace.go    // Marketplace + 常量（新增）
```

---

### 1.3 更新文档
**修正内容**:
1. 移除"零依赖"声明
2. 添加"精选依赖"说明
3. 列出所有依赖及用途
4. 更新示例代码

---

## 阶段2: 架构优化（P1 - 下周完成）

### 2.1 接口抽象层
**目标**: 提高可测试性和可扩展性

```go
// pkg/spapi/interfaces.go
package spapi

// HTTPClient 定义HTTP客户端接口
type HTTPClient interface {
    Do(*http.Request) (*http.Response, error)
}

// Signer 定义签名器接口
type Signer interface {
    Sign(context.Context, *http.Request) error
}

// RateLimiter 定义速率限制器接口
type RateLimiter interface {
    Wait(context.Context, string) error
    Update(api string, rate float64, burst int)
}

// Logger 定义日志接口
type Logger interface {
    Debug(msg string, fields ...Field)
    Info(msg string, fields ...Field)
    Warn(msg string, fields ...Field)
    Error(msg string, fields ...Field)
}

// MetricsCollector 定义指标收集接口
type MetricsCollector interface {
    RecordRequest(api, method string, duration time.Duration)
    RecordError(api, errorType string)
}
```

**优点**:
- ✅ 易于mock测试
- ✅ 用户可替换实现
- ✅ 解耦具体实现

---

### 2.2 可选依赖
**目标**: 监控和日志可选化

```go
// pkg/spapi/options.go

// WithLogger 设置自定义日志器（可选）
func WithLogger(logger Logger) ClientOption {
    return func(c *Config) error {
        c.Logger = logger
        return nil
    }
}

// WithMetrics 设置自定义指标收集器（可选）
func WithMetrics(metrics MetricsCollector) ClientOption {
    return func(c *Config) error {
        c.Metrics = metrics
        return nil
    }
}

// 默认使用no-op实现
type noOpLogger struct{}
func (n *noOpLogger) Debug(string, ...Field) {}
func (n *noOpLogger) Info(string, ...Field) {}
// ...

type noOpMetrics struct{}
func (n *noOpMetrics) RecordRequest(string, string, time.Duration) {}
// ...
```

**用户使用**:
```go
// 默认（无日志无监控）
client := spapi.NewClient(
    spapi.WithRegion(spapi.RegionNA),
    spapi.WithCredentials(...),
)

// 使用Zap日志
client := spapi.NewClient(
    spapi.WithRegion(spapi.RegionNA),
    spapi.WithCredentials(...),
    spapi.WithLogger(zapLogger),
)

// 使用自定义实现
client := spapi.NewClient(
    spapi.WithRegion(spapi.RegionNA),
    spapi.WithCredentials(...),
    spapi.WithLogger(myLogger),
    spapi.WithMetrics(myMetrics),
)
```

---

### 2.3 改进分层设计
**当前问题**: pkg/spapi直接依赖所有internal包

**优化方案**:
```go
internal/
  ├── core/           // 新增：核心协调层
  │   └── facade.go   // 统一入口，封装所有内部包
  ├── auth/
  ├── transport/
  ├── signer/
  └── ratelimit/

pkg/spapi/
  └── client.go       // 只依赖 internal/core
```

**facade.go示例**:
```go
package core

// Facade 封装所有内部组件
type Facade struct {
    auth      *auth.Client
    transport *transport.Client
    signer    *signer.LWASigner
    ratelimit *ratelimit.Manager
}

// NewFacade 创建核心门面
func NewFacade(config *Config) (*Facade, error) {
    // 初始化所有组件
    return &Facade{...}, nil
}

// DoRequest 统一请求入口
func (f *Facade) DoRequest(ctx context.Context, req *Request) (*Response, error) {
    // 协调所有组件
    return nil, nil
}
```

---

## 阶段3: 功能完善（P2 - 两周内完成）

### 3.1 Sandbox完整支持
**当前**: 部分支持

**目标**:
```go
// pkg/spapi/sandbox.go

// WithSandbox 启用Sandbox模式
func WithSandbox() ClientOption {
    return func(c *Config) error {
        c.Sandbox = true
        // 自动切换到sandbox endpoints
        return nil
    }
}

// 使用
client := spapi.NewClient(
    spapi.WithRegion(spapi.RegionNA),
    spapi.WithCredentials(...),
    spapi.WithSandbox(),  // 自动使用测试环境
)
```

---

### 3.2 项目结构优化
**当前**:
```
pkg/spapi/
  ├── client.go
  ├── config.go
  ├── orders-v0/
  ├── feeds-v2021-06-30/
  └── ... (57个API目录)
```

**优化后**:
```
pkg/spapi/
  ├── core/              // 核心类型
  │   ├── client.go
  │   ├── config.go
  │   ├── interfaces.go
  │   └── options.go
  ├── types/             // 公开类型
  │   ├── region.go
  │   ├── marketplace.go
  │   └── common.go
  ├── apis/              // 所有API（保持原结构）
  │   ├── orders-v0/
  │   ├── feeds-v2021-06-30/
  │   └── ...
  └── spapi.go          // 包入口，re-export常用类型
```

**spapi.go示例**:
```go
package spapi

// Re-export核心类型
import (
    "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi/core"
    "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi/types"
)

// 类型别名，简化导入
type (
    Client = core.Client
    Config = core.Config
    Region = types.Region
    Marketplace = types.Marketplace
)

// 常量re-export
var (
    RegionNA = types.RegionNA
    RegionEU = types.RegionEU
    RegionFE = types.RegionFE
)

// 函数re-export
var (
    NewClient = core.NewClient
    WithRegion = core.WithRegion
    WithCredentials = core.WithCredentials
)
```

**用户使用**:
```go
// 简化导入
import "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"

client := spapi.NewClient(
    spapi.WithRegion(spapi.RegionNA),
    spapi.WithCredentials(...),
)
```

---

### 3.3 版本管理
**目标**: 清晰的API版本管理

```
pkg/spapi/
  ├── apis/
  │   ├── stable/        // 稳定版本
  │   │   ├── orders-v0/
  │   │   └── ...
  │   ├── deprecated/    // 废弃的API
  │   │   ├── README.md  // 迁移指南
  │   │   └── ...
  │   └── experimental/  // 实验性API
  │       └── ...
  └── MIGRATION.md       // 版本迁移指南
```

---

### 3.4 测试体系完善
**当前**:
```
tests/
  ├── integration/
  └── benchmarks/
```

**优化后**:
```
tests/
  ├── unit/              // 单元测试（快速）
  │   ├── auth_test.go
  │   ├── signer_test.go
  │   └── ...
  ├── integration/       // 集成测试（需要凭证）
  │   ├── orders_test.go
  │   └── ...
  ├── e2e/              // 端到端测试
  │   └── workflow_test.go
  ├── benchmarks/       // 性能测试
  │   └── benchmark_test.go
  ├── fixtures/         // 测试数据
  │   ├── requests/
  │   └── responses/
  └── mocks/            // Mock实现
      ├── http_client.go
      ├── signer.go
      └── ...
```

---

### 3.5 中间件/插件机制
**目标**: 允许用户扩展功能

```go
// pkg/spapi/middleware.go

// Middleware 定义中间件类型
type Middleware func(next Handler) Handler

// Handler 定义请求处理器
type Handler func(context.Context, *Request) (*Response, error)

// WithMiddleware 添加中间件
func WithMiddleware(middlewares ...Middleware) ClientOption {
    return func(c *Config) error {
        c.Middlewares = append(c.Middlewares, middlewares...)
        return nil
    }
}

// 内置中间件
func LoggingMiddleware(logger Logger) Middleware {
    return func(next Handler) Handler {
        return func(ctx context.Context, req *Request) (*Response, error) {
            start := time.Now()
            resp, err := next(ctx, req)
            logger.Info("request completed",
                Field{"duration", time.Since(start)},
                Field{"path", req.Path},
            )
            return resp, err
        }
    }
}

func MetricsMiddleware(metrics MetricsCollector) Middleware {
    return func(next Handler) Handler {
        return func(ctx context.Context, req *Request) (*Response, error) {
            start := time.Now()
            resp, err := next(ctx, req)
            metrics.RecordRequest(req.API, req.Method, time.Since(start))
            return resp, err
        }
    }
}

// 用户自定义中间件
func CustomMiddleware(next Handler) Handler {
    return func(ctx context.Context, req *Request) (*Response, error) {
        // 自定义逻辑
        return next(ctx, req)
    }
}
```

**使用示例**:
```go
client := spapi.NewClient(
    spapi.WithRegion(spapi.RegionNA),
    spapi.WithCredentials(...),
    spapi.WithMiddleware(
        LoggingMiddleware(logger),
        MetricsMiddleware(metrics),
        CustomMiddleware,
    ),
)
```

---

## 版本规划

### v2.0.0 (Breaking Changes)
- ✅ 统一Region/Marketplace定义
- ✅ 清理internal/models
- ✅ 接口抽象层
- ✅ 可选依赖
- ✅ 改进分层设计

### v2.1.0 (功能增强)
- ✅ 完整Sandbox支持
- ✅ 中间件机制
- ✅ 项目结构优化

### v2.2.0 (完善)
- ✅ 版本管理
- ✅ 测试体系
- ✅ 性能优化

---

## 迁移指南

### 从v1.x迁移到v2.0

**Region类型变更**:
```go
// v1.x
import "github.com/vanling1111/amazon-sp-api-go-sdk/internal/models"
spapi.WithRegion(models.RegionNA)

// v2.0
import "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
spapi.WithRegion(spapi.RegionNA)
```

**Marketplace公开**:
```go
// v2.0新增
marketplace := spapi.MarketplaceUS
fmt.Println(marketplace.ID) // ATVPDKIKX0DER
```

**可选功能**:
```go
// v2.0新增
client := spapi.NewClient(
    spapi.WithRegion(spapi.RegionNA),
    spapi.WithCredentials(...),
    spapi.WithLogger(logger),      // 可选
    spapi.WithMetrics(metrics),    // 可选
    spapi.WithMiddleware(...),     // 可选
)
```

---

## 执行时间表

| 阶段 | 任务 | 时间 | 状态 |
|------|------|------|------|
| P0 | 统一类型定义 | 1天 | 待开始 |
| P0 | 清理internal/models | 0.5天 | 待开始 |
| P0 | 更新文档 | 0.5天 | 待开始 |
| P1 | 接口抽象层 | 1天 | 待开始 |
| P1 | 可选依赖 | 1天 | 待开始 |
| P1 | 改进分层 | 1天 | 待开始 |
| P2 | Sandbox支持 | 0.5天 | 待开始 |
| P2 | 项目结构优化 | 1天 | 待开始 |
| P2 | 版本管理 | 0.5天 | 待开始 |
| P2 | 测试体系 | 1天 | 待开始 |
| P2 | 中间件机制 | 1天 | 待开始 |

**总计**: 约9天工作量

---

## 风险评估

| 风险 | 影响 | 缓解措施 |
|------|------|---------|
| Breaking Changes | 高 | 详细迁移指南 + v2.0.0 |
| 测试覆盖不足 | 中 | 完善测试体系 |
| 性能下降 | 低 | 基准测试验证 |
| 用户迁移成本 | 中 | 保持v1.x维护6个月 |

---

## 成功标准

1. ✅ 所有测试通过
2. ✅ 性能无下降（基准测试）
3. ✅ 文档完整更新
4. ✅ 迁移指南清晰
5. ✅ CI/CD通过
6. ✅ 符合官方SP-API要求
