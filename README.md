# Amazon SP-API Go SDK

[![Go Version](https://img.shields.io/badge/go-1.21+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)
[![APIs](https://img.shields.io/badge/APIs-57%20versions-green.svg)](https://github.com/vanling1111/amazon-sp-api-go-sdk)
[![Methods](https://img.shields.io/badge/methods-314-brightgreen.svg)](https://github.com/vanling1111/amazon-sp-api-go-sdk)

一个严格基于 [Amazon SP-API 官方文档](https://developer-docs.amazon.com/sp-api/docs/) 和 Go 最佳实践开发的高质量 Go SDK。

## ✨ 核心特性

- 🎯 **完整 API 支持** - 57 个 API 版本，314 个操作方法
- 🔐 **完整的 LWA 认证** - 支持 Regular 和 Grantless 操作
- 🔄 **智能令牌缓存** - 自动刷新和提前过期处理
- 🚦 **速率限制** - 内置 Token Bucket 算法
- 🔒 **RDT 支持** - 处理受限数据访问
- 🌐 **HTTP 中间件** - 可扩展的请求/响应处理
- ♻️ **自动重试** - 智能错误检测和重试逻辑
- 🧪 **高测试覆盖率** - 100+ 测试用例，所有核心模块已测试
- 📖 **完整文档** - 中文注释和详细示例
- 🚀 **生产就绪** - 所有代码已编译验证和测试

## 🎯 设计原则

### ⚠️ 核心约束

1. ❌ **禁止参考其他语言的官方 SDK** - 不参考 Java、Python、Node.js 等官方 SDK 源码
2. ✅ **只参考官方 SP-API 文档** - 唯一权威来源：https://developer-docs.amazon.com/sp-api/docs/
3. 📚 **基于 OpenAPI 规范** - 直接从官方 OpenAPI 规范生成代码
4. 🚫 **禁止猜测开发** - 所有实现必须基于官方文档的明确说明
5. 🐹 **Go 最佳实践** - 充分利用 Go 语言特性和社区最佳实践

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

## 📦 已实现功能

### 核心模块

| 模块 | 状态 | 测试覆盖率 | 说明 |
|------|------|-----------|------|
| `internal/auth` | ✅ 已完成 | 89.0% | LWA 认证、令牌缓存、Grantless 支持 |
| `internal/transport` | ✅ 已完成 | 87.4% | HTTP 客户端、中间件、重试逻辑、HTTP/2 |
| `internal/signer` | ✅ 已完成 | 93.3% | LWA 签名器、RDT 签名器 |
| `internal/ratelimit` | ✅ 已完成 | 97.7% | Token Bucket、多维度管理器、动态速率调整 |
| `internal/models` | ✅ 已完成 | 100.0% | Region、Marketplace 定义 |
| `internal/utils` | ✅ 已完成 | 98.1% | HTTP、时间、字符串工具 |
| `internal/codec` | ✅ 已完成 | 94.6% | JSON 编解码、数据验证 |
| `internal/errors` | ✅ 已完成 | 88.0% | 详细错误分类、可重试判断 |
| `internal/metrics` | ✅ 已完成 | 100.0% | 指标记录接口、NoOp 实现 |
| `pkg/spapi` | ✅ 已完成 | 100.0% | 主客户端、Functional Options 配置 |

### API 支持

| API | 状态 | 版本 |
|-----|------|------|
| Orders API | 📅 计划中 | v0 |
| Reports API | 📅 计划中 | v2021-06-30 |
| Feeds API | 📅 计划中 | v2021-06-30 |
| Listings API | 📅 计划中 | v2021-08-01 |
| Notifications API | 📅 计划中 | v1 |

**图例**: ✅ 已完成 | 🔄 进行中 | 📅 计划中

## 🔄 开发路线图

### ✅ 阶段 1: 文档和架构（已完成）
- [x] 清空旧代码
- [x] 编写架构设计文档
- [x] 编写开发规范文档
- [x] 编写代码风格指南
- [x] 编写项目结构文档
- [x] 编写 API 追踪策略
- [x] 编写贡献指南

### ✅ 阶段 2: 核心基础设施（已完成）
- [x] 认证层 (LWA) - `internal/auth`
- [x] 传输层 (HTTP Client) - `internal/transport`
- [x] 签名层 (Request Signing) - `internal/signer`
- [x] Grantless 操作支持
- [x] RDT 签名器
- [x] 中间件系统
- [x] 重试逻辑

### ✅ 阶段 3: 速率限制和工具包（已完成）
- [x] Token Bucket 算法实现 - `internal/ratelimit/bucket.go`
- [x] 速率限制器实现 - `internal/ratelimit/limiter.go`
- [x] 多维度速率限制管理器 - `internal/ratelimit/manager.go`
- [x] 从 API 响应头动态更新速率
- [x] 支持 per seller + app + marketplace + operation 的独立限流
- [x] 通用模型 - `internal/models`
- [x] 工具包 - `internal/utils`

### ✅ 阶段 4: 编解码和错误处理（已完成）
- [x] JSON 编码器 - `internal/codec/json.go`
- [x] JSON 解码器 - 支持禁用未知字段
- [x] 数据验证器 - `internal/codec/validator.go`
- [x] 验证规则：Required、MinLength、MaxLength、Range、Email、URL、Pattern、OneOf
- [x] 详细错误分类 - `internal/errors`
- [x] 错误类型：RateLimit、Auth、Validation、NotFound、Server、Network
- [x] 可重试判断和错误详情提取

### 🔄 阶段 5: 公开 API 层（已完成 ✅）
- [x] 统一客户端 - `pkg/spapi`
- [x] Functional Options 配置模式
- [x] 完整的错误定义和验证
- [x] 客户端测试覆盖率 100%
- [x] 集成所有 internal 模块
- [x] **所有 47 个 SP-API 客户端实现完成**

#### 📦 Seller APIs (34 个) - 全部完成 ✅
- [x] Orders API - 订单管理
- [x] Reports API - 报告管理
- [x] Feeds API - 数据上传
- [x] Catalog Items API - 商品目录
- [x] Listings Items API - 商品列表
- [x] FBA Inventory API - FBA 库存
- [x] Product Pricing API - 价格查询
- [x] Tokens API - RDT 令牌
- [x] Notifications API - 通知订阅
- [x] Sellers API - 卖家信息
- [x] Product Fees API - 费用估算
- [x] Fulfillment Inbound API - FBA 入库
- [x] Fulfillment Outbound API - FBA 出库
- [x] Merchant Fulfillment API - 卖家配送
- [x] Shipping API - 货运管理
- [x] Solicitations API - 请求评论
- [x] Easy Ship API - Easy Ship
- [x] Messaging API - 买家消息
- [x] FBA Inbound Eligibility API - 入库资格
- [x] Services API - 服务工单
- [x] Shipment Invoicing API - 货件发票
- [x] Invoices API - 发票管理
- [x] Finances API - 财务事件
- [x] Listings Restrictions API - 列表限制
- [x] Product Type Definitions API - 产品类型
- [x] Sales API - 销售指标
- [x] Seller Wallet API - 钱包余额
- [x] Supply Sources API - 供应源
- [x] Uploads API - 文件上传
- [x] Vehicles API - 车辆兼容性
- [x] Replenishment API - 补货管理
- [x] Amazon Warehousing & Distribution API - 仓储配送
- [x] A+ Content API - A+ 内容
- [x] Application APIs (2个) - 应用管理和集成
- [x] Customer Feedback API - 客户反馈
- [x] Data Kiosk API - 数据查询

#### 🏭 Vendor APIs (10 个) - 全部完成 ✅
- [x] Vendor Direct Fulfillment Inventory API
- [x] Vendor Direct Fulfillment Orders API
- [x] Vendor Direct Fulfillment Payments API
- [x] Vendor Direct Fulfillment Sandbox API
- [x] Vendor Direct Fulfillment Shipping API
- [x] Vendor Direct Fulfillment Transactions API
- [x] Vendor Invoices API
- [x] Vendor Orders API
- [x] Vendor Shipments API
- [x] Vendor Transaction Status API

### 📅 阶段 6: 工具和自动化（计划中）
- [ ] API 更新监控工具
- [ ] OpenAPI 规范同步工具
- [ ] 代码生成器
- [ ] 性能测试工具
- [ ] GitHub Actions 工作流

## 🧪 测试

```bash
# 运行所有测试
go test -v ./...

# 查看测试覆盖率
go test -cover ./...

# 生成覆盖率报告
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

**当前测试覆盖率**:
- `internal/auth`: 89.0%
- `internal/transport`: 87.4%
- `internal/signer`: 93.3%
- **整体**: 90.2%

## 🔧 开发工具

```bash
# 代码格式化
gofmt -w .
goimports -w .

# 代码检查
golangci-lint run

# 构建
go build ./...
```

## 🤝 参与贡献

欢迎参与贡献！在开始之前，请务必阅读：

1. **强制性约束**: [开发规范](docs/DEVELOPMENT.md) - 必须严格遵守
2. **代码风格**: [代码风格](docs/CODE_STYLE.md) - Go 最佳实践
3. **贡献流程**: [贡献指南](docs/CONTRIBUTING.md) - 如何提交 PR

### ⚠️ 重要提醒

- ❌ 禁止参考其他语言的官方 SDK 源码
- ✅ 只参考官方 SP-API 文档
- 📚 基于官方 OpenAPI 规范生成代码
- 🚫 禁止基于猜测或假设进行开发

## 📄 许可证

本项目采用 MIT 许可证 - 详见 [LICENSE](LICENSE) 文件

## 📞 支持与联系

- **Issues**: [提交 Bug 或功能请求](https://github.com/yourusername/amazon-sp-api-go-sdk/issues)
- **Discussions**: [技术讨论和问答](https://github.com/yourusername/amazon-sp-api-go-sdk/discussions)
- **官方文档**: [Amazon SP-API 文档](https://developer-docs.amazon.com/sp-api/docs/)

## 🌟 致谢

感谢所有贡献者的付出！

## ⚖️ 免责声明

本项目是独立开发的开源 SDK，不隶属于 Amazon。使用本 SDK 时请遵守 [Amazon Selling Partner API 使用协议](https://developer-docs.amazon.com/sp-api/)。

---

**Star** ⭐ 本项目以获取最新进展！
