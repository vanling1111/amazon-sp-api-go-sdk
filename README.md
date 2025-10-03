# Amazon SP-API Go SDK

[![Go Version](https://img.shields.io/badge/go-1.25+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/license-Proprietary-red.svg)](LICENSE)
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

- 🛒 **核心业务**: Orders, Feeds, Reports, Catalog Items, Listings
- 📦 **库存物流**: FBA Inventory, Fulfillment, Merchant Fulfillment, Shipping
- 💰 **定价财务**: Product Pricing, Fees, Finances, Seller Wallet  
- 📢 **通知消息**: Notifications, Messaging, Solicitations
- 🏭 **Vendor API**: Direct Fulfillment 全系列, Orders, Invoices, Shipments
- ⚡ **高级功能**: A+ Content, Replenishment, AWD, Data Kiosk 等

**📋 完整列表**: [pkg/spapi/](pkg/spapi/) 目录 | **🤖 自动监控**: 每日检测官方 API 更新

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

## 📜 许可证

本项目采用**专有软件许可证**（Proprietary License）：

### ✅ 允许使用（免费）

- **个人学习和研究** - 用于提升个人技能和知识
- **学术研究** - 用于教育机构的非盈利性研究
- **非商业项目** - 用于个人的非盈利开源项目
- **测试评估** - 用于测试和评估软件功能

### ❌ 禁止使用（需要商业许可）

未经书面授权，严格禁止以下用途：

- **任何商业使用** - 不得用于任何盈利性业务活动
- **企业使用** - 不得在企业、公司、组织的生产环境中使用
- **商业软件集成** - 不得集成或嵌入到商业软件产品中
- **提供商业服务** - 不得使用本软件向第三方提供商业服务
- **托管服务** - 不得作为 SaaS 服务提供
- **衍生商业产品** - 不得基于本软件创建商业产品

### 💰 商业使用授权

如需在以下场景中使用本软件，必须获取商业许可证：

- 在任何企业或公司环境中使用（包括内部使用）
- 集成到任何商业产品或服务中
- 用于任何形式的盈利性业务
- 为客户提供基于本软件的服务

**商业许可咨询**：
- 📧 邮箱：vanling1111@gmail.com
- 💬 GitHub Issues：[提交咨询](https://github.com/vanling1111/amazon-sp-api-go-sdk/issues)
- 📄 完整许可证：[LICENSE](LICENSE)

### ⚠️ 重要说明

- 违反许可证条款将导致许可自动终止
- 许可方保留对未授权商业使用采取法律行动的权利
- 继续使用本软件即表示您同意并接受本许可证的所有条款

---

## 🌟 致谢

感谢所有贡献者的付出！

## ⚖️ 免责声明

本项目是独立开发的 SDK，不隶属于 Amazon。使用本 SDK 时请遵守 [Amazon Selling Partner API 使用协议](https://developer-docs.amazon.com/sp-api/)。

本软件按"现状"提供，不提供任何明示或暗示的保证。使用本软件的风险由您自行承担。

---

**关注本项目** ⭐ 以获取最新进展！

**注意**：请在使用前仔细阅读 [LICENSE](LICENSE) 文件，确保您的使用场景符合许可证要求。
