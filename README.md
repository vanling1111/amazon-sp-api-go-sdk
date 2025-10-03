# Amazon SP-API Go SDK

[![Go Version](https://img.shields.io/badge/go-1.25+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/license-AGPL--3.0%20%7C%20Commercial-blue.svg)](LICENSE)
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

本项目采用**双许可证**模式（Dual License）：

### 🆓 许可证选项 1：AGPL-3.0（开源许可证）

**适用于开源项目 - 免费使用**

#### ✅ 你可以（免费）：
- ✅ 用于个人学习和研究
- ✅ 用于学术研究项目
- ✅ 用于开源项目
- ✅ 修改源代码
- ✅ 分发软件

#### ⚠️ 你必须：
- **开源你的完整项目**（使用 AGPL-3.0 许可证）
- 包含原始版权声明
- 说明所有修改
- **向所有用户提供源代码**（包括网络用户）

#### 🚨 重要 - 网络 Copyleft：
如果你在服务器上运行此软件并让用户通过网络访问，你**必须**向这些用户提供完整源代码。这是 AGPL 与 GPL 的关键区别 - AGPL 关闭了"SaaS 漏洞"。

**适合以下场景**：
- 个人开发者的个人项目
- 开源社区项目
- 学术研究
- 愿意开源整个应用的项目

---

### 💰 许可证选项 2：商业许可证（Commercial License）

**适用于商业/专有软件使用 - 需要付费**

#### ✅ 商业许可证的好处：
- ✅ **无需开源** - 保持源代码私有
- ✅ **商业使用** - 用于商业产品和服务
- ✅ **SaaS/托管服务** - 提供托管服务
- ✅ **集成到专有软件** - 嵌入到闭源产品
- ✅ **企业支持** - 优先技术支持
- ✅ **法律保护** - 赔偿条款和保修选项

#### 💼 谁需要商业许可证？

如果你符合以下任何情况，需要购买商业许可证：

1. **开发商业软件**
   - 构建销售的产品
   - 创建专有应用
   - 集成到闭源系统

2. **运营 SaaS/云服务**
   - 为客户托管应用
   - 提供 API 服务
   - 运营 Web 应用

3. **企业/公司使用**
   - 在公司生产系统中使用
   - 内部业务应用
   - 企业软件开发

4. **无法满足 AGPL 要求**
   - 无法公开源代码
   - 受公司政策限制
   - 保护商业知识产权

#### 💵 商业许可定价

我们提供灵活的定价方案：
- **初创企业许可** - 适用于 < 10 人的公司
- **商业许可** - 适用于 10-100 人的公司
- **企业许可** - 适用于 > 100 人的公司
- **OEM 许可** - 随产品分发
- **永久许可** - 一次性付款
- **订阅许可** - 年付/月付

#### 📧 获取商业许可

**联系我们获取报价**：
- 📧 邮箱：vanling1111@gmail.com
- 💬 GitHub：[提交咨询](https://github.com/vanling1111/amazon-sp-api-go-sdk/issues)
- 📄 完整许可证：[LICENSE](LICENSE)

**请在咨询中包含**：
1. 公司名称和规模
2. 使用场景（SaaS、内部、OEM 等）
3. 预期部署规模
4. 期望的许可期限

通常 1-2 个工作日内回复。

---

### 📊 许可证对比

| 特性 | AGPL-3.0（免费） | 商业许可证 |
|------|----------------|----------|
| **费用** | 免费 | 付费 |
| **个人使用** | ✅ 是 | ✅ 是 |
| **商业使用** | ❌ 否（必须开源） | ✅ 是 |
| **保持代码私有** | ❌ 否 | ✅ 是 |
| **修改源码** | ✅ 是 | ✅ 是 |
| **分发** | ✅ 是（需附源码） | ✅ 是（无需源码） |
| **SaaS/托管** | ❌ 否（必须开源） | ✅ 是 |
| **企业支持** | ❌ 否 | ✅ 是 |
| **保修** | ❌ 否 | ✅ 可选 |

---

### ❓ 常见问题

**Q: 我是独立开发者，需要商业许可吗？**  
A: 如果你愿意开源你的项目（AGPL-3.0），则不需要。如果想保持代码私有，则需要商业许可。

**Q: 可以用于公司内部工具吗？**  
A: 需要商业许可。AGPL 要求即使是内部网络服务也要公开源代码。

**Q: 如果我修改了代码？**  
A: AGPL-3.0 要求公开修改。商业许可允许保持修改私有。

**Q: 可以作为产品的一部分分发吗？**  
A: 需要商业许可。AGPL 会要求你的整个产品也使用 AGPL-3.0。

**Q: 初创公司/非营利组织有优惠吗？**  
A: 有！联系我们 - 我们为初创公司提供折扣，为注册的非营利组织提供免费许可。

---

### ⚠️ 重要提醒

- 使用前请仔细阅读 [LICENSE](LICENSE) 文件
- 选择适合你的许可证选项
- 违反许可证条款将导致法律后果
- 我们保留对未授权商业使用采取法律行动的权利

---

## 🌟 致谢

感谢所有贡献者的付出！

## ⚖️ 免责声明

本项目是独立开发的 SDK，不隶属于 Amazon。使用本 SDK 时请遵守 [Amazon Selling Partner API 使用协议](https://developer-docs.amazon.com/sp-api/)。

本软件按"现状"提供，不提供任何明示或暗示的保证。使用本软件的风险由您自行承担。

---

**关注本项目** ⭐ 以获取最新进展！

**注意**：请在使用前仔细阅读 [LICENSE](LICENSE) 文件，确保您的使用场景符合许可证要求。
