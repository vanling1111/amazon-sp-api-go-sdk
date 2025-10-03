# Amazon SP-API 订单实时同步服务

这是一个**生产级**的订单同步服务示例，展示如何通过 Amazon SQS 实现订单的准实时同步。

## 📋 功能特点

- ✅ **准实时** - 10-30 秒延迟（Amazon SP-API 架构限制）
- ✅ **可靠** - SQS 消息持久化，不会丢失订单
- ✅ **低成本** - 不浪费 Orders API 配额
- ✅ **生产级** - 完整的错误处理和重试
- ✅ **Go 1.25** - 使用最新的迭代器特性
- ✅ **可复制** - 代码可以直接复制到你的项目

## 🚫 重要说明

### **为什么是"准实时"而不是"实时"？**

Amazon SP-API **不支持真正的实时推送**（WebSocket/SSE），只支持：
- EventBridge → SQS
- 你的应用轮询 SQS

**延迟来源**：
1. EventBridge 事件延迟：5-15 秒
2. SQS 轮询间隔：10 秒
3. **总延迟：10-30 秒**

这是 **Amazon 的设计限制，无法绕过**！

### **为什么不在 SDK 核心实现？**

- SQS 轮询属于消息队列集成，不属于 HTTP API 封装
- 官方 SDK（Java/Python/JS）都没有实现
- 用户可能使用其他消息队列（Kafka/RabbitMQ）
- **保持 SDK 边界清晰**

因此，我们提供**生产级示例代码**，用户可以：
- 直接复制使用
- 根据需求修改
- 集成到自己的架构

## 🚀 快速开始

### 1. 配置 AWS

创建 SQS 队列并配置 EventBridge：

```bash
# 1. 创建 SQS 队列
aws sqs create-queue --queue-name sp-api-notifications

# 2. 获取队列 ARN
aws sqs get-queue-attributes \
  --queue-url https://sqs.us-east-1.amazonaws.com/123456789/sp-api-notifications \
  --attribute-names QueueArn

# 3. 使用 SP-API Notifications API 创建订阅（见下文）
```

### 2. 配置 SP-API 通知订阅

```go
package main

import (
    "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
    notifications "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi/notifications-v1"
)

func setupNotifications() {
    client, _ := spapi.NewClient(...)
    notifClient := notifications.NewClient(client)
    
    // 创建 SQS 目标
    notifClient.CreateDestination(ctx, map[string]interface{}{
        "resourceSpecification": map[string]interface{}{
            "sqs": map[string]interface{}{
                "arn": "arn:aws:sqs:us-east-1:123456789:sp-api-notifications",
            },
        },
        "name": "OrderNotifications",
    })
    
    // 订阅订单变更通知
    notifClient.CreateSubscription(ctx, "ORDER_CHANGE", map[string]interface{}{
        "payloadVersion": "1.0",
        "destinationId": "your-destination-id",
    })
}
```

### 3. 配置环境变量

创建 `.env` 文件：

```bash
# SP-API 凭证
SP_API_CLIENT_ID=amzn1.application-oa2-client.xxxxx
SP_API_CLIENT_SECRET=xxxxx
SP_API_REFRESH_TOKEN=Atzr|xxxxx

# SQS 配置
SQS_QUEUE_URL=https://sqs.us-east-1.amazonaws.com/123456789/sp-api-notifications
AWS_REGION=us-east-1

# ERP 配置（可选）
ERP_WEBHOOK_URL=https://your-erp.com/api/orders
```

### 4. 运行服务

```bash
# 方式 1: 直接运行
go run main.go

# 方式 2: 使用 Docker
docker-compose up

# 方式 3: 构建二进制
go build -o order-sync
./order-sync
```

## 📖 代码说明

### 核心流程

```
1. SQS 轮询（10秒间隔）
      ↓
2. 收到 ORDER_CHANGE 通知
      ↓
3. 解析订单 ID
      ↓
4. 调用 Orders API 获取详情
      ↓
5. 使用迭代器获取订单项
      ↓
6. 推送到 ERP
```

### 关键组件

#### **SQS 轮询器**（`poller/poller.go`）
- Long Polling（20 秒）减少空轮询
- 自动消息删除
- 错误处理和重试
- 优雅退出

#### **事件处理器**（`handleOrderChange`）
- 解析通知负载
- 调用 Orders API
- 使用 Go 1.25 迭代器自动分页
- 推送到 ERP

#### **ERP 集成**（`pushToERP`）
- HTTP Webhook 示例
- 可替换为数据库写入
- 可替换为消息队列

## 🔧 自定义和扩展

### 修改轮询间隔

```go
poller.NewPoller(sqsClient, &poller.Config{
    PollInterval: 5 * time.Second,  // 改为 5 秒
    MaxMessages:  5,                 // 每次 5 条
})
```

### 添加更多事件类型

```go
service.poller.RegisterHandler("FEED_PROCESSING_FINISHED", handleFeed)
service.poller.RegisterHandler("REPORT_PROCESSING_FINISHED", handleReport)
service.poller.RegisterHandler("ANY_OFFER_CHANGED", handleOfferChange)
```

### 更换消息队列

如果你使用 Kafka/RabbitMQ 而不是 SQS：

```go
// 1. 修改 poller/poller.go
// 2. 替换 sqsClient 为你的消息队列客户端
// 3. 实现相同的轮询接口
```

## 📊 性能和成本

### 延迟对比

| 方案 | 延迟 | API 调用/小时 | 成本 |
|------|------|--------------|------|
| **SQS 通知（本方案）** | 10-30秒 | ~0（仅处理实际订单） | 极低 |
| 轮询 Orders API（5分钟） | 5-10分钟 | 12 次 | 中 |
| 轮询 Orders API（1分钟） | 1-2分钟 | 60 次 | 高（浪费配额） |

### SQS 成本

- 前 100 万次请求：免费
- 之后每 100 万次请求：$0.40
- **几乎可以忽略不计**

## ⚠️ 注意事项

### 1. 消息顺序不保证

SQS 标准队列不保证顺序。如果需要顺序：
- 使用 SQS FIFO 队列
- 或在应用层处理

### 2. 重复消息

SQS 可能重复投递消息。需要：
- 记录已处理的 NotificationID
- 实现幂等性

### 3. 消息丢失

极少情况下 SQS 可能丢消息。建议：
- 定期执行全量同步（每天一次）
- 作为 SQS 的补充

## 🐛 故障排除

### 收不到通知

1. 检查 SQS 队列是否有消息：`aws sqs receive-message --queue-url ...`
2. 检查 EventBridge 订阅是否正确
3. 检查 SP-API Notifications 订阅状态

### 处理失败

1. 查看日志中的错误信息
2. 检查 Orders API 权限
3. 检查网络连接

### 延迟过高（> 1 分钟）

1. 检查 EventBridge 是否正常
2. 检查 SQS 队列是否积压
3. 减小 PollInterval

## 📚 相关文档

- [SP-API Notifications API](https://developer-docs.amazon.com/sp-api/docs/notifications-api-v1-use-case-guide)
- [AWS SQS Go SDK](https://docs.aws.amazon.com/sdk-for-go/v2/developer-guide/go_sqs_code_examples.html)
- [EventBridge 设置](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-what-is.html)

## 📄 许可证

Apache 2.0 - 可自由使用和修改

