# Feed 批量上传工具

这是一个生产级的 Feed 上传工具示例，展示如何批量更新库存、价格、Listing 等。

## 📋 功能特点

- ✅ 完整的 Feed 上传流程
- ✅ 支持大文件上传（100MB+）
- ✅ 自动监控处理状态
- ✅ 处理 Feed 结果
- ✅ XML Feed 生成示例

## 🚀 快速开始

### 1. 配置环境变量

```bash
export SP_API_CLIENT_ID="amzn1.application-oa2-client.xxxxx"
export SP_API_CLIENT_SECRET="xxxxx"
export SP_API_REFRESH_TOKEN="Atzr|xxxxx"
```

### 2. 运行

```bash
go run main.go
```

## 📖 支持的 Feed 类型

### 库存 Feed
- `POST_INVENTORY_AVAILABILITY_DATA` - 更新库存数量
- `POST_FBA_INBOUND_CARTON_CONTENTS` - FBA 入库箱内容

### 价格 Feed  
- `POST_PRODUCT_PRICING_DATA` - 更新价格
- `POST_PRODUCT_OVERRIDES_DATA` - 价格覆盖

### Listing Feed
- `POST_PRODUCT_DATA` - 创建/更新商品
- `POST_PRODUCT_IMAGE_DATA` - 上传商品图片
- `POST_PRODUCT_RELATIONSHIP_DATA` - 变体关系

### 订单 Feed
- `POST_ORDER_ACKNOWLEDGEMENT_DATA` - 确认订单
- `POST_ORDER_FULFILLMENT_DATA` - 订单发货

完整列表：https://developer-docs.amazon.com/sp-api/docs/feed-type-values

## 💡 使用场景

### 场景 1：批量更新库存

```go
// 从 ERP 系统获取库存数据
inventory := getInventoryFromERP()

// 生成 Feed
feedContent := generateInventoryFeed(inventory)

// 上传
feedID, _ := uploadFeed(ctx, feedsClient, feedContent, "POST_INVENTORY_AVAILABILITY_DATA")

// 等待处理完成
monitorFeedProcessing(ctx, feedsClient, feedID)
```

### 场景 2：批量更新价格

```go
// 从定价系统获取价格
prices := getPricesFromSystem()

// 生成价格 Feed
feedContent := generatePricingFeed(prices)

// 上传
feedID, _ := uploadFeed(ctx, feedsClient, feedContent, "POST_PRODUCT_PRICING_DATA")
```

### 场景 3：批量创建 Listing

```go
// 从产品数据库获取商品信息
products := getProductsFromDB()

// 生成商品 Feed
feedContent := generateProductFeed(products)

// 上传
feedID, _ := uploadFeed(ctx, feedsClient, feedContent, "POST_PRODUCT_DATA")
```

## 🔧 自定义

### 修改 Feed 类型

```go
// 在 main.go 中修改
feedType := "POST_PRODUCT_PRICING_DATA"  // 改为价格 Feed
```

### 处理大文件

对于 100MB+ 的 Feed：

```go
// TODO: v1.2.0 将添加分片上传支持
// uploader := transfer.NewChunkedUploader(client)
// uploader.Upload(ctx, largeFile, feedType)
```

### 并发上传多个 Feed

```go
feedTypes := []string{
    "POST_INVENTORY_AVAILABILITY_DATA",
    "POST_PRODUCT_PRICING_DATA",
}

for _, feedType := range feedTypes {
    go func(ft string) {
        feedContent := generateFeed(ft)
        uploadFeed(ctx, feedsClient, feedContent, ft)
    }(feedType)
}
```

## 📚 相关文档

- [Feeds API 文档](https://developer-docs.amazon.com/sp-api/docs/feeds-api-v2021-06-30-reference)
- [Feed 类型列表](https://developer-docs.amazon.com/sp-api/docs/feed-type-values)
- [XML Feed 规范](https://developer-docs.amazon.com/sp-api/docs/xml-feed-format)

