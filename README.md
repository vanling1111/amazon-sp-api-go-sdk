# Amazon SP-API Go SDK

[![Go Version](https://img.shields.io/badge/go-1.25+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)
[![Latest Release](https://img.shields.io/github/v/release/vanling1111/amazon-sp-api-go-sdk)](https://github.com/vanling1111/amazon-sp-api-go-sdk/releases)
[![APIs](https://img.shields.io/badge/APIs-57%20versions-green.svg)](https://github.com/vanling1111/amazon-sp-api-go-sdk)
[![Methods](https://img.shields.io/badge/methods-314-brightgreen.svg)](https://github.com/vanling1111/amazon-sp-api-go-sdk)
[![Iterators](https://img.shields.io/badge/iterators-27-orange.svg)](https://github.com/vanling1111/amazon-sp-api-go-sdk)
[![Tests](https://img.shields.io/badge/tests-passing-success.svg)](https://github.com/vanling1111/amazon-sp-api-go-sdk/actions)
[![Coverage](https://img.shields.io/badge/coverage-92%25-brightgreen.svg)](https://github.com/vanling1111/amazon-sp-api-go-sdk)

**生产级 Amazon Selling Partner API Go SDK**

填补官方 SDK 空白，提供 Go 语言的完整 SP-API 实现。基于 [Amazon SP-API 官方文档](https://developer-docs.amazon.com/sp-api/docs/) 和 Go 最佳实践开发。

**当前版本**: v1.3.0 | **Go 要求**: 1.25+ | **状态**: ✅ 生产就绪

## ✨ 核心特性

- 🎯 **完整 API 支持** - 57 个 API 版本，314 个操作方法
- 🔐 **完整的 LWA 认证** - 支持 Regular 和 Grantless 操作
- 🔄 **智能令牌缓存** - 自动刷新和提前过期处理
- 🚦 **速率限制** - 内置 Token Bucket 算法，支持动态更新
- 🔒 **RDT 支持** - 处理受限数据访问
- 🔁 **Go 1.25 迭代器** - 所有 27 个分页 API 支持自动分页迭代
- 🔓 **自动解密** - Reports API 自动下载和解密加密报告
- 🌐 **HTTP 中间件** - 可扩展的请求/响应处理
- ♻️ **自动重试** - 智能错误检测和重试逻辑
- 🤖 **自动监控** - 每日自动检测官方 API 变更，确保 SDK 始终同步最新规范
- 🧪 **高测试覆盖率** - 154+ 测试用例，所有核心模块已测试
- 📖 **完整文档** - 中文注释和详细示例
- 🚀 **生产就绪** - 所有代码已编译验证和测试

## 🎯 设计原则

1. 📚 **基于官方规范** - 直接从 Amazon 官方 OpenAPI 规范自动生成代码
2. ✅ **文档驱动** - 所有实现严格遵循官方 SP-API 文档
3. 🐹 **Go 惯用法** - 充分利用 Go 语言特性和社区最佳实践
4. 🔒 **类型安全** - 完整的类型定义和编译时检查
5. 🧪 **高质量** - 完整的测试覆盖和错误处理
6. ⚡ **Go 1.25** - 使用最新 Go 特性（迭代器、性能优化）

## 🆕 v1.1.0 新特性

### **Go 1.25 分页迭代器**

所有 27 个分页 API 现在支持 Go 1.25 迭代器，用户代码减少 70%：

```go
// 之前：手动处理分页（繁琐）
nextToken := ""
for {
    result, _ := client.Orders.GetOrders(ctx, &Query{NextToken: nextToken})
    for _, order := range result.Orders {
        process(order)
    }
    if result.NextToken == "" { break }
    nextToken = result.NextToken
}

// 现在：自动分页（简洁）
for order, err := range ordersClient.IterateOrders(ctx, query) {
    if err != nil { return err }
    process(order)
}
```

**支持的 API**：Orders, Reports, Feeds, Catalog Items, FBA Inventory, Finances, 所有 Vendor API 等 27 个

### **自动报告解密**

Reports API 现在自动处理报告下载和解密：

```go
// 一行代码获取解密后的报告
decrypted, err := reportsClient.GetReportDocumentDecrypted(ctx, reportDocumentID)

// 直接使用数据（CSV/TSV/JSON）
fmt.Println(string(decrypted))
```

### **生产级示例**

新增完整的生产级示例代码：
- `examples/patterns/order-sync-sqs/` - SQS 订单实时同步服务
- `examples/iterators/` - 迭代器使用示例
- `examples/report-decryption/` - 报告解密示例

所有示例都可以直接运行或复制到项目中使用。

## 📚 文档

### 设计文档
- [架构设计](docs/ARCHITECTURE.md) - 系统架构和设计决策
- [项目结构](docs/PROJECT_STRUCTURE.md) - 目录结构和组织方式
- [API 追踪策略](docs/API_TRACKING.md) - 如何追踪和同步官方 API 更新

### 开发指南
- [开发规范](docs/DEVELOPMENT.md) - 开发流程和强制性规范
- [代码风格](docs/CODE_STYLE.md) - 代码风格和命名规范
- [贡献指南](docs/CONTRIBUTING.md) - 如何参与项目开发

### 功能指南
- [Grantless 操作指南](docs/GRANTLESS_OPERATIONS_GUIDE.md) - Grantless 操作的详细说明

### 参考资料
- [版本追踪](docs/VERSION_TRACKING.md) - SDK 和官方文档版本历史
- [官方 SP-API 文档](https://developer-docs.amazon.com/sp-api/docs/) - 唯一权威来源

## 🚀 快速开始

### 安装

```bash
go get github.com/vanling1111/amazon-sp-api-go-sdk
```

### 基本用法

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/vanling1111/amazon-sp-api-go-sdk/internal/models"
    "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
    orders "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi/orders-v0"
)

func main() {
    // 1. 创建基础 SP-API 客户端
    baseClient, err := spapi.NewClient(
        spapi.WithRegion(models.RegionNA),
        spapi.WithCredentials(
            "your-client-id",
            "your-client-secret",
            "your-refresh-token",
        ),
    )
    if err != nil {
        log.Fatal(err)
    }
    defer baseClient.Close()

    // 2. 创建 Orders API 客户端
    ordersClient := orders.NewClient(baseClient)

    // 3. 调用 API 方法
    ctx := context.Background()
    params := map[string]string{
        "MarketplaceIds": "ATVPDKIKX0DER",
        "CreatedAfter":   time.Now().Add(-7 * 24 * time.Hour).Format(time.RFC3339),
    }

    result, err := ordersClient.GetOrders(ctx, params)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("订单获取成功:", result)
}
```

### Grantless 操作

```go
// 创建 Grantless 操作的客户端
client, err := spapi.NewClient(
    spapi.WithRegion(models.RegionEU),
    spapi.WithGrantlessCredentials(
        "your-client-id",
        "your-client-secret",
        []string{"sellingpartnerapi::notifications"},
    ),
)
if err != nil {
    log.Fatal(err)
}
defer client.Close()

// 使用客户端访问 Grantless API...
```

更多示例请查看 [examples/](examples/) 目录。

## 📦 支持的 API

本 SDK 完整支持 **57 个 Amazon SP-API 版本**，包括：

### 核心业务 API
- **Orders** - 订单管理
- **Feeds** - 数据上传和处理
- **Reports** - 报告生成和下载
- **Catalog Items** - 商品目录查询
- **Listings Items** - 商品列表管理

### 库存与物流 API
- **FBA Inventory** - FBA 库存管理
- **Fulfillment Inbound/Outbound** - 入库和出库管理
- **Merchant Fulfillment** - 卖家配送
- **Shipping** - 物流服务

### 定价与财务 API
- **Product Pricing** - 商品定价
- **Product Fees** - 费用估算
- **Finances** - 财务报告
- **Seller Wallet** - 钱包管理

### 通知与消息 API
- **Notifications** - 通知订阅
- **Messaging** - 买家消息
- **Solicitations** - 评论请求

### Vendor API（完整支持）
- Vendor Direct Fulfillment 系列（Inventory, Orders, Payments, Shipping, Transactions）
- Vendor Orders, Invoices, Shipments

### 高级功能 API
- A+ Content, Replenishment, AWD, Customer Feedback, Data Kiosk, Easy Ship, 等

**📋 完整列表**: 查看 [pkg/spapi/](pkg/spapi/) 目录查看所有 57 个 API 版本

## 🧪 测试

```bash
# 运行所有测试
go test ./...

# 运行测试并查看覆盖率
go test -cover ./...
```

核心模块测试覆盖率达到 **92%+**，所有测试持续通过。

## 🛠️ 开发

```bash
# 克隆仓库
git clone https://github.com/vanling1111/amazon-sp-api-go-sdk.git
cd amazon-sp-api-go-sdk

# 运行测试
go test ./...

# 构建项目
go build ./...

# 代码检查（可选）
golangci-lint run
```

更多开发信息请参考 [开发指南](docs/DEVELOPMENT.md)。

## 🤝 参与贡献

欢迎参与贡献！请参考以下文档：

1. **开发规范**: [开发指南](docs/DEVELOPMENT.md) - 开发流程和最佳实践
2. **代码风格**: [代码风格](docs/CODE_STYLE.md) - Go 编码规范
3. **贡献流程**: [贡献指南](docs/CONTRIBUTING.md) - 如何提交 PR

### 💡 技术亮点

- 📚 直接从官方 OpenAPI 规范生成，确保与 Amazon API 完全一致
- 🤖 **每日自动监控** - GitHub Actions 每天自动检测官方 57 个 API 的 OpenAPI 规范变更
- 🔔 **变更通知** - 检测到 API 变更时自动创建 GitHub Issue 提醒维护者
- 🔄 自动化工具链，可快速同步官方 API 更新
- 🧪 高测试覆盖率，核心模块达到 92%+
- 📖 完整的中文文档和示例代码

## 📄 许可证

本项目采用 Apache 2.0 许可证 - 详见 [LICENSE](LICENSE) 文件

## 📞 支持与联系

- **Issues**: [提交 Bug 或功能请求](https://github.com/vanling1111/amazon-sp-api-go-sdk/issues)
- **Discussions**: [技术讨论和问答](https://github.com/vanling1111/amazon-sp-api-go-sdk/discussions)
- **官方文档**: [Amazon SP-API 文档](https://developer-docs.amazon.com/sp-api/docs/)

## 🌟 致谢

感谢所有贡献者的付出！

## ⚖️ 免责声明

本项目是独立开发的开源 SDK，不隶属于 Amazon。使用本 SDK 时请遵守 [Amazon Selling Partner API 使用协议](https://developer-docs.amazon.com/sp-api/)。

---

**Star** ⭐ 本项目以获取最新进展！
