# SP-API 公开包

本目录包含对外暴露的公开 API，是 SDK 的主要入口。

## 包结构

```
pkg/spapi/
├── client.go                      # 主客户端
├── client_test.go                 # 主客户端测试
├── config.go                      # 配置选项
├── config_test.go                 # 配置测试
├── errors.go                      # 错误定义
├── regions.go                     # 区域定义
├── marketplaces.go                # 市场定义
│
├── orders/                        # Orders API
│   ├── client.go                  # Orders API 客户端
│   ├── types.go                   # 手写类型定义
│   ├── types_gen.go               # 自动生成的类型
│   ├── client_test.go             # 单元测试
│   └── examples_test.go           # 示例测试
│
├── reports/                       # Reports API
│   ├── client.go
│   ├── types.go
│   ├── types_gen.go
│   └── client_test.go
│
├── feeds/                         # Feeds API
│   ├── client.go
│   ├── types.go
│   ├── types_gen.go
│   └── client_test.go
│
├── listings/                      # Listings Items API
│   ├── client.go
│   ├── types.go
│   └── client_test.go
│
├── notifications/                 # Notifications API
│   ├── client.go
│   ├── types.go
│   └── client_test.go
│
├── catalog/                       # Catalog Items API
│   ├── client.go
│   ├── types.go
│   └── client_test.go
│
├── pricing/                       # Product Pricing API
│   ├── client.go
│   ├── types.go
│   └── client_test.go
│
├── fba/                           # FBA (Fulfillment by Amazon)
│   ├── inventory/                 # FBA Inventory API
│   │   ├── client.go
│   │   ├── types.go
│   │   └── client_test.go
│   ├── inbound/                   # FBA Inbound API
│   │   ├── client.go
│   │   ├── types.go
│   │   └── client_test.go
│   └── outbound/                  # FBA Outbound API
│       ├── client.go
│       ├── types.go
│       └── client_test.go
│
├── sellers/                       # Sellers API
│   ├── client.go
│   ├── types.go
│   └── client_test.go
│
└── tokens/                        # Tokens API (RDT)
    ├── client.go
    ├── types.go
    └── client_test.go
```

## 使用方式

### 基本用法

```go
package main

import (
    "context"
    "log"
    "time"

    "github.com/yourusername/amazon-sp-api-go-sdk/pkg/spapi"
    "github.com/yourusername/amazon-sp-api-go-sdk/pkg/spapi/orders"
)

func main() {
    // 创建客户端
    client, err := spapi.NewClient(
        spapi.WithCredentials("client-id", "client-secret", "refresh-token"),
        spapi.WithRegion(spapi.RegionNA),
        spapi.WithMarketplace(spapi.MarketplaceUS),
    )
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()

    // 使用 Orders API
    ctx := context.Background()
    ordersAPI := orders.NewClient(client)
    
    resp, err := ordersAPI.GetOrders(ctx, &orders.GetOrdersRequest{
        MarketplaceIDs: []string{spapi.MarketplaceUS},
        CreatedAfter:   time.Now().Add(-24 * time.Hour),
    })
    if err != nil {
        log.Fatal(err)
    }

    for _, order := range resp.Orders {
        log.Printf("Order ID: %s, Status: %s", order.AmazonOrderID, order.OrderStatus)
    }
}
```

### 配置选项

```go
client, err := spapi.NewClient(
    // 必需：LWA 凭证
    spapi.WithCredentials(clientID, clientSecret, refreshToken),
    
    // 必需：区域
    spapi.WithRegion(spapi.RegionNA),
    
    // 可选：默认市场
    spapi.WithMarketplace(spapi.MarketplaceUS),
    
    // 可选：超时配置
    spapi.WithTimeout(30 * time.Second),
    
    // 可选：重试配置
    spapi.WithRetry(3, 1*time.Second),
    
    // 可选：自定义 User-Agent
    spapi.WithUserAgent("MyApp/1.0"),
    
    // 可选：速率限制
    spapi.WithRateLimit(10, 60), // 10 请求/60秒
)
```

### Grantless 操作

```go
// 创建 Grantless 客户端
client, err := spapi.NewClient(
    spapi.WithGrantlessCredentials(clientID, clientSecret, []string{
        spapi.ScopeNotifications,
    }),
    spapi.WithRegion(spapi.RegionNA),
)

// 使用 Notifications API
subscriptions, err := client.Notifications.GetSubscriptions(ctx, "ORDER_CHANGE")
```

## API 模块列表

| API | 导入路径 | 状态 | 版本 |
|-----|---------|------|------|
| Orders API | `pkg/spapi/orders` | 📅 计划中 | v0 |
| Reports API | `pkg/spapi/reports` | 📅 计划中 | v2021-06-30 |
| Feeds API | `pkg/spapi/feeds` | 📅 计划中 | v2021-06-30 |
| Listings API | `pkg/spapi/listings` | 📅 计划中 | v2021-08-01 |
| Notifications API | `pkg/spapi/notifications` | 📅 计划中 | v1 |
| Catalog Items API | `pkg/spapi/catalog` | 📅 计划中 | v2022-04-01 |
| Product Pricing API | `pkg/spapi/pricing` | 📅 计划中 | v0 |
| FBA Inventory API | `pkg/spapi/fba/inventory` | 📅 计划中 | v1 |
| FBA Inbound API | `pkg/spapi/fba/inbound` | 📅 计划中 | v0 |
| FBA Outbound API | `pkg/spapi/fba/outbound` | 📅 计划中 | v2020-07-01 |
| Sellers API | `pkg/spapi/sellers` | 📅 计划中 | v1 |
| Tokens API | `pkg/spapi/tokens` | 📅 计划中 | v2021-03-01 |

**图例**: ✅ 已实现 | 🔄 进行中 | 📅 计划中

## 错误处理

```go
orders, err := client.Orders.GetOrders(ctx, req)
if err != nil {
    // 检查特定错误类型
    var apiErr *spapi.APIError
    if errors.As(err, &apiErr) {
        log.Printf("API Error: %s - %s", apiErr.Code, apiErr.Message)
        
        // 检查是否是速率限制错误
        if apiErr.Code == "QuotaExceeded" {
            log.Println("Rate limit exceeded, retry later")
        }
        return
    }
    
    // 其他错误
    log.Printf("Request failed: %v", err)
    return
}
```

## 并发安全

所有公开 API 都是并发安全的，可以在多个 goroutine 中共享同一个客户端：

```go
client, _ := spapi.NewClient(...)

var wg sync.WaitGroup
for i := 0; i < 10; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        orders, _ := client.Orders.GetOrders(ctx, req)
        // 处理订单...
    }()
}
wg.Wait()
```

## 官方文档

- [SP-API 官方文档](https://developer-docs.amazon.com/sp-api/docs/)
- [SP-API 参考](https://developer-docs.amazon.com/sp-api/reference/)

## 导入路径

```go
import "github.com/yourusername/amazon-sp-api-go-sdk/pkg/spapi"
```

## 许可证

MIT License

