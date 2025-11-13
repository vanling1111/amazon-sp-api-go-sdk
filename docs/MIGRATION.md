# 迁移指南

本文档帮助您从旧版本迁移到最新版本的Amazon SP-API Go SDK。

---

## 从 v1.x 迁移到 v2.x

### 重大变更

#### 1. Region类型公开化 (v2.0.0)

**变更内容**:
- 移除 `internal/models.Region`
- 使用 `pkg/spapi.Region`

**迁移步骤**:

```go
// ❌ v1.x (旧)
import "github.com/vanling1111/amazon-sp-api-go-sdk/internal/models"

client := spapi.NewClient(
    spapi.WithRegion(models.RegionNA),
    spapi.WithCredentials(...),
)

// ✅ v2.0+ (新)
import "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"

client := spapi.NewClient(
    spapi.WithRegion(spapi.RegionNA),
    spapi.WithCredentials(...),
)
```

**查找替换**:
```bash
# 全局替换
models.RegionNA → spapi.RegionNA
models.RegionEU → spapi.RegionEU
models.RegionFE → spapi.RegionFE

# 移除导入
- import "github.com/vanling1111/amazon-sp-api-go-sdk/internal/models"
```

---

#### 2. MarketplaceID公开化 (v2.0.0)

**新增功能**:
- 19个预定义的MarketplaceID常量
- `MarketplaceID.Region()` 方法

**使用示例**:

```go
// v2.0+ 新增
marketplaceID := spapi.MarketplaceUS
region := marketplaceID.Region()  // 自动获取所属区域

// 在API调用中使用
params := map[string]string{
    "MarketplaceIds": string(spapi.MarketplaceUS),
}
```

---

#### 3. 接口抽象层 (v2.1.0)

**变更内容**:
- Logger改为接口类型
- 新增Metrics和Tracer接口

**迁移步骤**:

```go
// ❌ v1.x (旧) - 使用具体类型
config := &spapi.Config{
    Logger: zapLogger,  // 具体的zap logger
}

// ✅ v2.1+ (新) - 使用接口
client := spapi.NewClient(
    spapi.WithLogger(zapLogger),  // 实现Logger接口即可
)
```

**自定义Logger实现**:

```go
type MyLogger struct{}

func (l *MyLogger) Debug(msg string, fields ...spapi.Field) {}
func (l *MyLogger) Info(msg string, fields ...spapi.Field) {}
func (l *MyLogger) Warn(msg string, fields ...spapi.Field) {}
func (l *MyLogger) Error(msg string, fields ...spapi.Field) {}
func (l *MyLogger) With(fields ...spapi.Field) spapi.Logger { return l }

// 使用
client := spapi.NewClient(
    spapi.WithLogger(&MyLogger{}),
)
```

---

#### 4. 默认No-Op实现 (v2.1.0)

**变更内容**:
- 如果不提供Logger/Metrics/Tracer，自动使用no-op实现
- 不再强制依赖第三方库

**迁移步骤**:

```go
// ❌ v1.x (旧) - 必须提供logger
client := spapi.NewClient(
    spapi.WithRegion(spapi.RegionNA),
    spapi.WithCredentials(...),
    // 必须提供logger，否则可能panic
)

// ✅ v2.1+ (新) - 可选
client := spapi.NewClient(
    spapi.WithRegion(spapi.RegionNA),
    spapi.WithCredentials(...),
    // 自动使用no-op实现，不输出日志
)

// 或者提供自定义实现
client := spapi.NewClient(
    spapi.WithRegion(spapi.RegionNA),
    spapi.WithCredentials(...),
    spapi.WithLogger(myLogger),  // 可选
)
```

---

### 新增功能

#### 1. Sandbox支持 (v2.2.0)

**使用方法**:

```go
// 自动切换到Sandbox环境
client := spapi.NewClient(
    spapi.WithRegion(spapi.RegionNA),
    spapi.WithSandbox(),  // 自动切换到 RegionNASandbox
    spapi.WithCredentials(...),
)

// 或者直接使用Sandbox区域
client := spapi.NewClient(
    spapi.WithRegion(spapi.RegionNASandbox),
    spapi.WithCredentials(...),
)

// 检查是否为Sandbox
if client.Config().Region.IsSandbox() {
    fmt.Println("Running in sandbox mode")
}
```

---

#### 2. 中间件机制 (v2.2.0)

**使用方法**:

```go
// 使用内置中间件
client := spapi.NewClient(
    spapi.WithRegion(spapi.RegionNA),
    spapi.WithCredentials(...),
    spapi.WithMiddleware(
        spapi.LoggingMiddleware(logger),
        spapi.MetricsMiddleware(metrics),
        spapi.TracingMiddleware(tracer),
    ),
)

// 自定义中间件
func CustomMiddleware(next spapi.Handler) spapi.Handler {
    return func(ctx context.Context, req *http.Request) (*http.Response, error) {
        // 请求前处理
        fmt.Println("Before request:", req.URL)
        
        // 执行请求
        resp, err := next(ctx, req)
        
        // 请求后处理
        if err == nil {
            fmt.Println("After request:", resp.StatusCode)
        }
        
        return resp, err
    }
}

client := spapi.NewClient(
    spapi.WithRegion(spapi.RegionNA),
    spapi.WithCredentials(...),
    spapi.WithMiddleware(CustomMiddleware),
)
```

---

## 完整迁移示例

### v1.x 代码

```go
package main

import (
    "context"
    "log"
    
    "github.com/vanling1111/amazon-sp-api-go-sdk/internal/models"
    "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
    orders "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi/orders-v0"
)

func main() {
    // v1.x 方式
    client, err := spapi.NewClient(
        spapi.WithRegion(models.RegionNA),
        spapi.WithCredentials("client-id", "client-secret", "refresh-token"),
    )
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()
    
    ordersClient := orders.NewClient(client)
    result, err := ordersClient.GetOrders(context.Background(), map[string]string{
        "MarketplaceIds": "ATVPDKIKX0DER",
    })
    if err != nil {
        log.Fatal(err)
    }
    
    log.Println(result)
}
```

### v2.2+ 代码（推荐）

```go
package main

import (
    "context"
    "log"
    
    "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
    orders "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi/orders-v0"
)

func main() {
    // v2.2+ 方式（功能更强大）
    client, err := spapi.NewClient(
        // 基本配置
        spapi.WithRegion(spapi.RegionNA),
        spapi.WithCredentials("client-id", "client-secret", "refresh-token"),
        
        // 可选：Sandbox测试环境
        // spapi.WithSandbox(),
        
        // 可选：自定义日志和指标
        // spapi.WithLogger(myLogger),
        // spapi.WithMetrics(myMetrics),
        
        // 可选：中间件
        // spapi.WithMiddleware(
        //     spapi.LoggingMiddleware(logger),
        //     spapi.MetricsMiddleware(metrics),
        // ),
    )
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()
    
    ordersClient := orders.NewClient(client)
    result, err := ordersClient.GetOrders(context.Background(), map[string]string{
        "MarketplaceIds": string(spapi.MarketplaceUS),  // 使用常量
    })
    if err != nil {
        log.Fatal(err)
    }
    
    log.Println(result)
}
```

---

## 常见问题

### Q1: 为什么要移除internal/models导入？

**A**: Go的`internal`包不允许外部导入。将Region等类型移到`pkg/spapi`使其成为公开API，更符合Go的最佳实践。

### Q2: 旧代码会报错吗？

**A**: 是的，v2.0是Breaking Change。但迁移很简单，只需要：
1. 移除`internal/models`导入
2. 将`models.RegionXX`改为`spapi.RegionXX`

### Q3: 必须使用新功能吗？

**A**: 不必须。新功能（Logger、Metrics、Sandbox、Middleware）都是可选的。最小化使用方式：

```go
client := spapi.NewClient(
    spapi.WithRegion(spapi.RegionNA),
    spapi.WithCredentials(...),
)
```

### Q4: 如何验证迁移成功？

**A**: 运行测试：

```bash
go test ./...
go build ./...
```

如果编译通过且测试通过，迁移成功。

---

## 迁移检查清单

- [ ] 移除所有`internal/models`导入
- [ ] 替换所有`models.RegionXX`为`spapi.RegionXX`
- [ ] 更新go.mod中的SDK版本
- [ ] 运行`go test ./...`确保测试通过
- [ ] 运行`go build ./...`确保编译通过
- [ ] 考虑使用新功能（Sandbox、中间件等）
- [ ] 更新文档和注释

---

## 获取帮助

- **文档**: [docs/](../docs/)
- **示例**: [examples/](../examples/)
- **问题**: [GitHub Issues](https://github.com/vanling1111/amazon-sp-api-go-sdk/issues)
- **讨论**: [GitHub Discussions](https://github.com/vanling1111/amazon-sp-api-go-sdk/discussions)
