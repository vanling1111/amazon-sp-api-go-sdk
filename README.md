# Amazon SP-API Go SDK

[![Go Version](https://img.shields.io/badge/go-1.25+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/license-AGPL--3.0%20%7C%20Commercial-blue.svg)](LICENSE)
[![Latest Release](https://img.shields.io/github/v/release/vanling1111/amazon-sp-api-go-sdk)](https://github.com/vanling1111/amazon-sp-api-go-sdk/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/vanling1111/amazon-sp-api-go-sdk)](https://goreportcard.com/report/github.com/vanling1111/amazon-sp-api-go-sdk)

[![APIs](https://img.shields.io/badge/APIs-57%20versions-green.svg)](https://github.com/vanling1111/amazon-sp-api-go-sdk)
[![Methods](https://img.shields.io/badge/methods-314-brightgreen.svg)](https://github.com/vanling1111/amazon-sp-api-go-sdk)
[![Iterators](https://img.shields.io/badge/iterators-27-orange.svg)](https://github.com/vanling1111/amazon-sp-api-go-sdk)
[![Tests](https://img.shields.io/badge/tests-passing-success.svg)](https://github.com/vanling1111/amazon-sp-api-go-sdk/actions)
[![Coverage](https://img.shields.io/badge/coverage-92%25-brightgreen.svg)](https://github.com/vanling1111/amazon-sp-api-go-sdk)
[![GitHub Stars](https://img.shields.io/github/stars/vanling1111/amazon-sp-api-go-sdk?style=social)](https://github.com/vanling1111/amazon-sp-api-go-sdk)
[![pkg.go.dev](https://pkg.go.dev/badge/github.com/vanling1111/amazon-sp-api-go-sdk.svg)](https://pkg.go.dev/github.com/vanling1111/amazon-sp-api-go-sdk)

**生产级 Amazon Selling Partner API Go SDK**

填补官方 SDK 空白，提供 Go 语言的完整 SP-API 实现。基于 [Amazon SP-API 官方文档](https://developer-docs.amazon.com/sp-api/docs/) 和 Go 最佳实践开发。

**当前版本**: v2.3.0 | **Go 要求**: 1.25+ | **状态**: ✅ 生产就绪

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
7. 🔧 **精选依赖** - 只依赖业界最佳实践库，不重复造轮子

## 🌟 最新特性

### v1.3.0 - 云原生可观测性 (2025-10-03)

- 📊 **OpenTelemetry** - 分布式追踪，兼容 Jaeger/Zipkin
- 📈 **Prometheus** - 标准指标导出，Grafana 就绪
- 🔍 **完整可观测性** - 日志 + 追踪 + 指标

### v1.2.0 - 企业级可靠性

- 🪵 **结构化日志** - Zap 集成
- 🔌 **熔断器** - Circuit Breaker 防止级联失败
- ⚡ **JSON 优化** - 性能提升 3-5 倍
- 📦 **大文件传输** - 流式上传/下载

### v1.1.0 - Go 1.25 增强

- 🔁 **自动分页迭代器** - 27 个 API 支持，代码减少 70%
- 🔓 **自动报告解密** - AES-256-CBC 一键解密
- 🚀 **生产级示例** - SQS 订单同步等

📖 **详细说明**: [完整功能清单](docs/FEATURES.md) | [更新日志](CHANGELOG.md)

## 📚 文档

| 类型 | 文档 | 说明 |
|------|------|------|
| 📘 **API 参考** | [pkg.go.dev](https://pkg.go.dev/github.com/vanling1111/amazon-sp-api-go-sdk) | 完整 API 文档 |
| 📘 **API 参考** | [本地查看](docs/API_REFERENCE.md) | GoDoc 使用指南 |
| 🚀 **快速入门** | [示例代码](examples/) | 10+ 可运行示例 |
| 📖 **功能指南** | [完整功能清单](docs/FEATURES.md) | 38 项功能详解 |
| 📖 **功能指南** | [分页迭代器](docs/PAGINATION_GUIDE.md) | Go 1.25 迭代器 |
| 📖 **功能指南** | [报告解密](docs/REPORT_DECRYPTION.md) | AES-256 解密 |
| 📖 **功能指南** | [Grantless 操作](docs/GRANTLESS_OPERATIONS_GUIDE.md) | 无需授权 API |
| 🏗️ **架构设计** | [系统架构](docs/ARCHITECTURE.md) | 设计决策 |
| 👨‍💻 **开发指南** | [开发规范](docs/DEVELOPMENT.md) | 开发流程 |
| 🤝 **贡献** | [贡献指南](docs/CONTRIBUTING.md) | 如何提交 PR |

📌 **官方文档**: [Amazon SP-API 文档](https://developer-docs.amazon.com/sp-api/docs/)

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

    "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
    orders "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi/orders-v0"
)

func main() {
    // 1. 创建基础 SP-API 客户端
    baseClient, err := spapi.NewClient(
        spapi.WithRegion(spapi.RegionNA),
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
        "MarketplaceIds": string(spapi.MarketplaceUS), // 使用公开的MarketplaceID常量
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
    spapi.WithRegion(spapi.RegionEU),
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

### 自定义日志和指标（v2.1.0新增）

```go
// 使用自定义Logger和Metrics
client, err := spapi.NewClient(
    spapi.WithRegion(spapi.RegionNA),
    spapi.WithCredentials("your-client-id", "your-client-secret", "your-refresh-token"),
    spapi.WithLogger(myLogger),      // 可选：自定义日志
    spapi.WithMetrics(myMetrics),    // 可选：自定义指标收集
    spapi.WithTracer(myTracer),      // 可选：自定义分布式追踪
)

// 默认情况下，SDK使用no-op实现（不输出日志、不收集指标）
// 这样可以保持零依赖，用户可以根据需要选择性启用
```

### Sandbox测试环境（v2.2.0新增）

```go
// 使用Sandbox环境进行测试
client, err := spapi.NewClient(
    spapi.WithRegion(spapi.RegionNA),
    spapi.WithSandbox(),  // 自动切换到测试环境
    spapi.WithCredentials("your-client-id", "your-client-secret", "your-refresh-token"),
)

// Sandbox环境不会影响生产数据，适合开发和测试
```

### 中间件扩展（v2.2.0新增）

```go
// 使用中间件添加自定义逻辑
client, err := spapi.NewClient(
    spapi.WithRegion(spapi.RegionNA),
    spapi.WithCredentials("your-client-id", "your-client-secret", "your-refresh-token"),
    spapi.WithMiddleware(
        spapi.LoggingMiddleware(logger),    // 日志记录
        spapi.MetricsMiddleware(metrics),   // 指标收集
        CustomMiddleware,                    // 自定义中间件
    ),
)
```

更多示例请查看 [examples/](examples/) 目录。

## 📚 文档

- 📖 [快速开始指南](docs/QUICKSTART.md) - 5分钟上手
- 🔄 [迁移指南](docs/MIGRATION.md) - 从v1.x迁移到v2.x
- 🏗️ [架构设计](docs/ARCHITECTURE.md) - SDK架构说明
- 📁 [项目结构](docs/PROJECT_STRUCTURE.md) - 目录组织
- 🔧 [重构计划](docs/REFACTORING_PLAN.md) - 长期优化计划
- 💡 [功能特性](docs/FEATURES.md) - 详细功能说明
- 📄 [变更日志](CHANGELOG.md) - 版本历史

## 📦 支持的 API

本 SDK 完整支持 **57 个 Amazon SP-API 版本**，包括：

- 🛒 **核心业务**: Orders, Feeds, Reports, Catalog Items, Listings
- 📦 **库存物流**: FBA Inventory, Fulfillment, Merchant Fulfillment, Shipping
- 💰 **定价财务**: Product Pricing, Fees, Finances, Seller Wallet  
- 📢 **通知消息**: Notifications, Messaging, Solicitations
- 🏭 **Vendor API**: Direct Fulfillment 全系列, Orders, Invoices, Shipments
- ⚡ **高级功能**: A+ Content, Replenishment, AWD, Data Kiosk 等

**📋 完整列表**: [pkg/spapi/](pkg/spapi/) 目录 | **🤖 自动监控**: 每日检测官方 API 更新

## 📚 依赖说明

本SDK采用**精选依赖**策略，只依赖业界最佳实践库：

### 核心依赖
- **AWS SDK** - AWS服务集成（SQS等）
- **json-iterator** - 高性能JSON处理（比标准库快3-5倍）

### 企业级功能（推荐）
- **Zap** - Uber开源的高性能日志库
- **Prometheus** - CNCF监控标准
- **OpenTelemetry** - CNCF分布式追踪标准

### 设计理念
- ✅ 不重复造轮子，使用成熟方案
- ✅ 接口化设计，允许用户替换实现
- ✅ 所有依赖都是可选的（除核心功能）

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

## 📞 支持与联系

- **Issues**: [提交 Bug 或功能请求](https://github.com/vanling1111/amazon-sp-api-go-sdk/issues)
- **Discussions**: [技术讨论和问答](https://github.com/vanling1111/amazon-sp-api-go-sdk/discussions)
- **官方文档**: [Amazon SP-API 文档](https://developer-docs.amazon.com/sp-api/docs/)

## 📜 许可证

本项目采用**双许可证**模式：

### 🆓 AGPL-3.0（开源许可证）- 免费

✅ **适用于**：
- 个人学习和研究
- 学术研究项目
- 开源项目

⚠️ **要求**：
- 必须开源你的完整项目（AGPL-3.0）
- 包含原始版权声明
- 向所有用户提供源代码（包括网络用户）

### 💰 商业许可证 - 付费

✅ **适用于**：
- 商业产品和服务
- SaaS/托管服务
- 企业内部使用
- 闭源软件集成

📧 **获取商业许可**：vanling1111@gmail.com

---

📄 **详细条款**: [LICENSE](LICENSE) | 📊 **许可证对比**: 见 LICENSE 文件

---

## 🌟 致谢

感谢所有贡献者的付出！

## ⚖️ 免责声明

本项目是独立开发的 SDK，不隶属于 Amazon。使用本 SDK 时请遵守 [Amazon Selling Partner API 使用协议](https://developer-docs.amazon.com/sp-api/)。

本软件按"现状"提供，不提供任何明示或暗示的保证。使用本软件的风险由您自行承担。

---

**关注本项目** ⭐ 以获取最新进展！

**注意**：请在使用前仔细阅读 [LICENSE](LICENSE) 文件，确保您的使用场景符合许可证要求。
