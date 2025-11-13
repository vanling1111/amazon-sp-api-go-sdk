# 快速开始指南

5分钟快速上手Amazon SP-API Go SDK。

---

## 前置要求

1. **Go 1.25+**
2. **Amazon SP-API凭证**:
   - Client ID
   - Client Secret
   - Refresh Token
3. **市场ID**: 例如 `ATVPDKIKX0DER` (美国)

---

## 安装

```bash
go get github.com/vanling1111/amazon-sp-api-go-sdk@latest
```

---

## 第一个程序

### 1. 创建客户端

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
    orders "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi/orders-v0"
)

func main() {
    // 创建客户端
    client, err := spapi.NewClient(
        spapi.WithRegion(spapi.RegionNA),  // 北美区域
        spapi.WithCredentials(
            "your-client-id",
            "your-client-secret",
            "your-refresh-token",
        ),
    )
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()
    
    fmt.Println("✅ 客户端创建成功！")
}
```

### 2. 调用API

```go
// 创建Orders API客户端
ordersClient := orders.NewClient(client)

// 获取订单
result, err := ordersClient.GetOrders(context.Background(), map[string]string{
    "MarketplaceIds": string(spapi.MarketplaceUS),
    "CreatedAfter":   "2024-01-01T00:00:00Z",
})
if err != nil {
    log.Fatal(err)
}

fmt.Printf("获取到 %d 个订单\n", len(result.Orders))
```

---

## 常用场景

### 场景1: 获取订单列表

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"
    
    "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
    orders "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi/orders-v0"
)

func main() {
    client, _ := spapi.NewClient(
        spapi.WithRegion(spapi.RegionNA),
        spapi.WithCredentials("client-id", "client-secret", "refresh-token"),
    )
    defer client.Close()
    
    ordersClient := orders.NewClient(client)
    
    // 获取最近7天的订单
    result, err := ordersClient.GetOrders(context.Background(), map[string]string{
        "MarketplaceIds": string(spapi.MarketplaceUS),
        "CreatedAfter":   time.Now().Add(-7 * 24 * time.Hour).Format(time.RFC3339),
    })
    if err != nil {
        log.Fatal(err)
    }
    
    for _, order := range result.Orders {
        fmt.Printf("订单ID: %s, 金额: %s\n", 
            order.AmazonOrderId, 
            order.OrderTotal.Amount)
    }
}
```

### 场景2: 使用Sandbox测试

```go
client, _ := spapi.NewClient(
    spapi.WithRegion(spapi.RegionNA),
    spapi.WithSandbox(),  // 🔧 启用Sandbox模式
    spapi.WithCredentials("client-id", "client-secret", "refresh-token"),
)

// Sandbox环境不会影响生产数据
```

### 场景3: 添加日志

```go
// 自定义Logger
type SimpleLogger struct{}

func (l *SimpleLogger) Debug(msg string, fields ...spapi.Field) {
    fmt.Println("[DEBUG]", msg)
}
func (l *SimpleLogger) Info(msg string, fields ...spapi.Field) {
    fmt.Println("[INFO]", msg)
}
func (l *SimpleLogger) Warn(msg string, fields ...spapi.Field) {
    fmt.Println("[WARN]", msg)
}
func (l *SimpleLogger) Error(msg string, fields ...spapi.Field) {
    fmt.Println("[ERROR]", msg)
}
func (l *SimpleLogger) With(fields ...spapi.Field) spapi.Logger {
    return l
}

// 使用
client, _ := spapi.NewClient(
    spapi.WithRegion(spapi.RegionNA),
    spapi.WithCredentials(...),
    spapi.WithLogger(&SimpleLogger{}),  // 📝 添加日志
)
```

### 场景4: 处理多个市场

```go
marketplaces := []spapi.MarketplaceID{
    spapi.MarketplaceUS,
    spapi.MarketplaceCA,
    spapi.MarketplaceMX,
}

for _, marketplace := range marketplaces {
    result, err := ordersClient.GetOrders(ctx, map[string]string{
        "MarketplaceIds": string(marketplace),
        "CreatedAfter":   "2024-01-01T00:00:00Z",
    })
    if err != nil {
        log.Printf("获取 %s 订单失败: %v", marketplace, err)
        continue
    }
    fmt.Printf("%s: %d 个订单\n", marketplace, len(result.Orders))
}
```

---

## 配置选项

### 基本配置

```go
client := spapi.NewClient(
    spapi.WithRegion(spapi.RegionNA),           // 必需：区域
    spapi.WithCredentials(...),                  // 必需：凭证
)
```

### 高级配置

```go
client := spapi.NewClient(
    // 基本配置
    spapi.WithRegion(spapi.RegionNA),
    spapi.WithCredentials(...),
    
    // 超时和重试
    spapi.WithHTTPTimeout(60 * time.Second),    // HTTP超时
    spapi.WithMaxRetries(5),                     // 最大重试次数
    
    // 速率限制
    spapi.WithRateLimitBuffer(0.2),             // 20%缓冲
    
    // 调试
    spapi.WithDebug(),                           // 启用调试模式
    
    // 可选功能
    spapi.WithSandbox(),                         // Sandbox模式
    spapi.WithLogger(myLogger),                  // 自定义日志
    spapi.WithMetrics(myMetrics),                // 自定义指标
    spapi.WithMiddleware(                        // 自定义中间件
        spapi.LoggingMiddleware(logger),
    ),
)
```

---

## 支持的区域

```go
// 生产环境
spapi.RegionNA  // 北美（美国、加拿大、墨西哥、巴西）
spapi.RegionEU  // 欧洲（英国、德国、法国等）
spapi.RegionFE  // 远东（日本、澳大利亚、新加坡、印度）

// Sandbox环境
spapi.RegionNASandbox
spapi.RegionEUSandbox
spapi.RegionFESandbox
```

---

## 支持的市场

```go
// 北美
spapi.MarketplaceUS  // 美国
spapi.MarketplaceCA  // 加拿大
spapi.MarketplaceMX  // 墨西哥
spapi.MarketplaceBR  // 巴西

// 欧洲
spapi.MarketplaceUK  // 英国
spapi.MarketplaceDE  // 德国
spapi.MarketplaceFR  // 法国
spapi.MarketplaceIT  // 意大利
spapi.MarketplaceES  // 西班牙
// ... 更多市场

// 远东
spapi.MarketplaceJP  // 日本
spapi.MarketplaceAU  // 澳大利亚
spapi.MarketplaceSG  // 新加坡
spapi.MarketplaceIN  // 印度
```

---

## 错误处理

```go
result, err := ordersClient.GetOrders(ctx, params)
if err != nil {
    // 检查错误类型
    if strings.Contains(err.Error(), "authentication failed") {
        log.Fatal("认证失败，请检查凭证")
    } else if strings.Contains(err.Error(), "rate limit") {
        log.Println("速率限制，等待后重试")
        time.Sleep(5 * time.Second)
        // 重试...
    } else {
        log.Printf("API调用失败: %v", err)
    }
    return
}

// 处理成功响应
for _, order := range result.Orders {
    // 处理订单...
}
```

---

## 最佳实践

### 1. 使用defer关闭客户端

```go
client, err := spapi.NewClient(...)
if err != nil {
    log.Fatal(err)
}
defer client.Close()  // ✅ 确保资源释放
```

### 2. 使用context控制超时

```go
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

result, err := ordersClient.GetOrders(ctx, params)
```

### 3. 处理分页

```go
// 使用Go 1.25迭代器（如果API支持）
for order := range ordersClient.ListOrders(ctx, params) {
    fmt.Println(order.AmazonOrderId)
}
```

### 4. 先在Sandbox测试

```go
// 开发阶段
client := spapi.NewClient(
    spapi.WithRegion(spapi.RegionNA),
    spapi.WithSandbox(),  // ✅ 使用Sandbox
    spapi.WithCredentials(...),
)

// 生产环境
client := spapi.NewClient(
    spapi.WithRegion(spapi.RegionNA),
    // ❌ 移除WithSandbox()
    spapi.WithCredentials(...),
)
```

---

## 下一步

- 📚 查看[完整文档](../README.md)
- 💡 浏览[示例代码](../examples/)
- 🔧 了解[高级功能](./FEATURES.md)
- 🐛 [报告问题](https://github.com/vanling1111/amazon-sp-api-go-sdk/issues)

---

## 获取帮助

遇到问题？

1. 查看[FAQ](./FAQ.md)
2. 搜索[Issues](https://github.com/vanling1111/amazon-sp-api-go-sdk/issues)
3. 提问[Discussions](https://github.com/vanling1111/amazon-sp-api-go-sdk/discussions)
4. 查看[官方文档](https://developer-docs.amazon.com/sp-api/docs/)
