# 分页迭代器使用指南

本指南介绍如何使用 SDK 的 Go 1.25 分页迭代器功能。

## 📋 概述

Amazon SP-API 的许多接口返回分页数据（如订单列表、报告列表等）。传统方式需要手动管理 `NextToken`，代码繁琐且容易出错。

从 v1.1.0 开始，SDK 为所有 27 个分页 API 提供了**自动分页迭代器**，基于 Go 1.25 的迭代器特性。

## ✨ 核心优势

- ✅ **代码减少 70%** - 无需手动管理 NextToken
- ✅ **自动错误处理** - 错误即时返回
- ✅ **支持提前退出** - 使用 break 随时中断
- ✅ **类型安全** - 编译时检查
- ✅ **惯用 Go 语法** - 原生 for-range 循环

## 🚀 快速开始

### 基础使用

```go
package main

import (
    "context"
    "log"
    
    "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
    "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi/orders-v0"
)

func main() {
    // 创建客户端
    baseClient, _ := spapi.NewClient(...)
    ordersClient := orders_v0.NewClient(baseClient)
    
    // 查询参数
    query := map[string]string{
        "MarketplaceIds": "ATVPDKIKX0DER",
        "CreatedAfter":   "2025-01-01T00:00:00Z",
    }
    
    // 使用迭代器（自动处理所有分页）
    for order, err := range ordersClient.IterateOrders(context.Background(), query) {
        if err != nil {
            log.Fatal(err)
        }
        
        // 处理订单
        orderID := order["AmazonOrderId"]
        log.Printf("Processing order: %s", orderID)
    }
}
```

## 📖 支持的 API

### 高频使用 API

#### Orders API
```go
ordersClient := orders_v0.NewClient(baseClient)

// 迭代订单
for order, err := range ordersClient.IterateOrders(ctx, query) {
    // 处理订单
}

// 迭代订单项
for item, err := range ordersClient.IterateOrderItems(ctx, orderID, nil) {
    // 处理订单项
}
```

#### Reports API
```go
reportsClient := reports_v2021_06_30.NewClient(baseClient)

// 迭代报告
for report, err := range reportsClient.IterateReports(ctx, query) {
    reportID := report["reportId"]
    status := report["processingStatus"]
    // 处理报告
}
```

#### Feeds API
```go
feedsClient := feeds_v2021_06_30.NewClient(baseClient)

// 迭代 Feed
for feed, err := range feedsClient.IterateFeeds(ctx, query) {
    feedID := feed["feedId"]
    // 处理 Feed
}
```

#### Catalog Items API
```go
catalogClient := catalog_items_v2022_04_01.NewClient(baseClient)

// 迭代商品
for item, err := range catalogClient.IterateCatalogItems(ctx, query) {
    asin := item["asin"]
    // 处理商品
}
```

#### FBA Inventory API
```go
inventoryClient := fba_inventory_v1.NewClient(baseClient)

// 迭代库存
for inventory, err := range inventoryClient.IterateInventorySummaries(ctx, query) {
    sku := inventory["sellerSku"]
    quantity := inventory["totalQuantity"]
    // 处理库存
}
```

#### Finances API
```go
financesClient := finances_v0.NewClient(baseClient)

// 迭代财务事件
for event, err := range financesClient.IterateFinancialEvents(ctx, query) {
    // 处理财务事件
}

// 迭代财务事件组
for group, err := range financesClient.IterateFinancialEventGroups(ctx, query) {
    // 处理事件组
}
```

### 完整 API 列表（27 个）

所有有分页的 API 都支持迭代器：

1. orders-v0
2. reports-v2021-06-30
3. feeds-v2021-06-30
4. catalog-items-v0
5. catalog-items-v2020-12-01
6. catalog-items-v2022-04-01
7. fba-inventory-v1
8. finances-v0
9. finances-v2024-06-19
10. fulfillment-inbound-v0
11. fulfillment-inbound-v2024-03-20
12. fulfillment-outbound-v2020-07-01
13. listings-items-v2021-08-01
14. services-v1
15. invoices-v2024-06-19
16. seller-wallet-v2024-03-01
17. supply-sources-v2020-07-01
18. vehicles-v2024-11-01
19. aplus-content-v2020-11-01
20. data-kiosk-v2023-11-15
21. amazon-warehousing-and-distribution-v2024-05-09
22-27. vendor-* 系列（6 个）

## 🎯 高级用法

### 提前退出

```go
// 查找特定订单后退出
targetOrderID := "123-456789-123456"

for order, err := range ordersClient.IterateOrders(ctx, query) {
    if err != nil {
        return err
    }
    
    orderID := order["AmazonOrderId"]
    if orderID == targetOrderID {
        log.Printf("Found target order: %s", orderID)
        break  // 提前退出，不继续迭代
    }
}
```

### 限制数量

```go
// 只处理前 100 个订单
count := 0
for order, err := range ordersClient.IterateOrders(ctx, query) {
    if err != nil {
        return err
    }
    
    process(order)
    
    count++
    if count >= 100 {
        break
    }
}
```

### 并发处理

```go
// 使用 channel 收集数据
ordersChan := make(chan map[string]interface{}, 100)

// 在 goroutine 中迭代
go func() {
    defer close(ordersChan)
    
    for order, err := range ordersClient.IterateOrders(ctx, query) {
        if err != nil {
            log.Printf("Error: %v", err)
            return
        }
        ordersChan <- order
    }
}()

// 并发处理订单（Go 1.25 自动正确捕获变量）
for order := range ordersChan {
    go func() {
        // 处理订单（不再需要 order := order）
        processOrder(order)
    }()
}
```

### 错误处理

```go
for order, err := range ordersClient.IterateOrders(ctx, query) {
    if err != nil {
        // 错误类型判断
        if apiErr, ok := err.(*spapi.APIError); ok {
            log.Printf("API Error: %s - %s", apiErr.Code, apiErr.Message)
            
            // 根据错误类型决定是否继续
            if apiErr.StatusCode == 429 {
                // 速率限制，等待后重试
                time.Sleep(1 * time.Minute)
                continue
            }
        }
        
        // 其他错误，中断
        return err
    }
    
    process(order)
}
```

### 上下文控制

```go
// 使用 context 超时控制
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
defer cancel()

for order, err := range ordersClient.IterateOrders(ctx, query) {
    if err != nil {
        if err == context.DeadlineExceeded {
            log.Println("Timeout: processing took too long")
        }
        return err
    }
    
    process(order)
}
```

## 🔍 实现原理

### Go 1.25 迭代器

迭代器使用 Go 1.25 的 `iter.Seq2` 类型：

```go
import "iter"

func (c *Client) IterateOrders(ctx context.Context, query map[string]string) iter.Seq2[map[string]interface{}, error] {
    return func(yield func(map[string]interface{}, error) bool) {
        nextToken := ""
        for {
            // 获取当前页
            result, err := c.GetOrders(ctx, query)
            if err != nil {
                yield(nil, err)
                return
            }
            
            // 遍历当前页数据
            for _, order := range result.Orders {
                if !yield(order, nil) {
                    return  // 用户调用 break
                }
            }
            
            // 检查下一页
            if result.NextToken == "" {
                break
            }
            nextToken = result.NextToken
        }
    }
}
```

### 自动分页

迭代器内部自动：
1. 调用 API 获取当前页
2. 遍历当前页数据
3. 检查是否有下一页
4. 自动设置 NextToken
5. 重复直到没有更多数据

用户完全无感知分页逻辑！

## ⚠️ 注意事项

### 1. 数据类型

迭代器返回 `map[string]interface{}`，需要类型断言：

```go
for order, err := range ordersClient.IterateOrders(ctx, query) {
    if err != nil {
        return err
    }
    
    // 类型断言
    orderID, ok := order["AmazonOrderId"].(string)
    if !ok {
        log.Println("Invalid order ID type")
        continue
    }
    
    orderTotal, ok := order["OrderTotal"].(map[string]interface{})
    // ...
}
```

### 2. 内存使用

迭代器是**流式处理**，不会一次性加载所有数据到内存：

```go
// ✅ 好：流式处理，内存占用低
for order, err := range ordersClient.IterateOrders(ctx, query) {
    process(order)  // 处理后即可丢弃
}

// ❌ 差：全部加载到内存
allOrders := []map[string]interface{}{}
for order, err := range ordersClient.IterateOrders(ctx, query) {
    allOrders = append(allOrders, order)  // 占用大量内存
}
```

### 3. API 配额

迭代器会自动获取所有页的数据，这会消耗 API 配额：

```go
// 如果有 1000 个订单，每页 100 个
// 迭代器会自动调用 10 次 API

for order, err := range ordersClient.IterateOrders(ctx, query) {
    // 处理 1000 个订单
    // 消耗了 10 次 API 调用
}
```

建议：
- 使用日期范围限制（`CreatedAfter`, `CreatedBefore`）
- 使用 break 在达到预期数量后退出
- 监控 API 使用情况

## 📚 相关文档

- [Go 1.25 迭代器官方文档](https://go.dev/blog/range-functions)
- [SP-API 分页文档](https://developer-docs.amazon.com/sp-api/docs/)
- [完整示例代码](../examples/iterators/main.go)

## 🆚 对比：迭代器 vs 手动分页

| 特性 | 手动分页 | 迭代器 |
|------|---------|--------|
| **代码行数** | ~15 行 | ~5 行 |
| **NextToken 管理** | 手动 | 自动 |
| **错误处理** | 手动 | 自动 |
| **提前退出** | 需要 return | break |
| **可读性** | ⭐⭐ | ⭐⭐⭐⭐⭐ |
| **维护性** | ⭐⭐ | ⭐⭐⭐⭐⭐ |

## 💡 最佳实践

### 1. 使用上下文超时

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
defer cancel()

for order, err := range ordersClient.IterateOrders(ctx, query) {
    // ...
}
```

### 2. 限制数据范围

```go
query := map[string]string{
    "MarketplaceIds": "ATVPDKIKX0DER",
    "CreatedAfter":   time.Now().Add(-7*24*time.Hour).Format(time.RFC3339),  // 只获取最近 7 天
    "CreatedBefore":  time.Now().Format(time.RFC3339),
}
```

### 3. 批量处理

```go
batch := []map[string]interface{}{}
batchSize := 100

for order, err := range ordersClient.IterateOrders(ctx, query) {
    if err != nil {
        return err
    }
    
    batch = append(batch, order)
    
    // 达到批次大小，处理一批
    if len(batch) >= batchSize {
        processBatch(batch)
        batch = batch[:0]  // 清空
    }
}

// 处理剩余的
if len(batch) > 0 {
    processBatch(batch)
}
```

### 4. 进度显示

```go
count := 0
for order, err := range ordersClient.IterateOrders(ctx, query) {
    if err != nil {
        return err
    }
    
    count++
    if count%100 == 0 {
        log.Printf("Processed %d orders...", count)
    }
    
    process(order)
}

log.Printf("Total processed: %d orders", count)
```

## 🐛 常见问题

### Q: 迭代器会自动处理速率限制吗？

A: 是的！SDK 的速率限制器会自动应用到所有 API 调用，包括迭代器内部的调用。

### Q: 如何知道总共有多少数据？

A: Amazon SP-API 不提供总数。你需要完整遍历或使用计数器：

```go
count := 0
for _, err := range ordersClient.IterateOrders(ctx, query) {
    if err != nil {
        return err
    }
    count++
}
log.Printf("Total: %d orders", count)
```

### Q: 迭代器性能如何？

A: 迭代器是**流式处理**，性能优秀：
- 不会一次性加载所有数据
- 内存占用低
- 只在需要时调用 API

### Q: 可以嵌套使用迭代器吗？

A: 可以！

```go
for order, err := range ordersClient.IterateOrders(ctx, query) {
    if err != nil {
        return err
    }
    
    orderID := order["AmazonOrderId"]
    
    // 嵌套：获取每个订单的订单项
    for item, err := range ordersClient.IterateOrderItems(ctx, orderID, nil) {
        if err != nil {
            return err
        }
        // 处理订单项
    }
}
```

### Q: 迭代器支持并发吗？

A: 单个迭代器实例**不是**并发安全的，但你可以：

```go
// 方式 1：为每个 goroutine 创建独立的客户端
for i := range 10 {
    go func() {
        client := orders_v0.NewClient(baseClient)
        for order, err := range client.IterateOrders(ctx, query) {
            // ...
        }
    }()
}

// 方式 2：使用 channel 传递数据
ordersChan := make(chan map[string]interface{}, 100)
go func() {
    for order, err := range ordersClient.IterateOrders(ctx, query) {
        if err == nil {
            ordersChan <- order
        }
    }
    close(ordersChan)
}()

// 并发处理
for order := range ordersChan {
    go processOrder(order)
}
```

## 📝 完整示例

参见 [examples/iterators/main.go](../examples/iterators/main.go) 获取完整的可运行示例。

## 🔗 相关指南

- [API 使用示例](../examples/)
- [错误处理指南](ERROR_HANDLING.md)
- [性能优化指南](PERFORMANCE.md)

