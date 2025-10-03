# SDK 功能清单

## 📋 完整功能列表

本文档列出 Amazon SP-API Go SDK 的所有功能特性。

---

## 🎯 核心功能

### 1. API 支持

| 功能 | 状态 | 说明 |
|------|------|------|
| **API 覆盖** | ✅ 完整 | 57 个 API 版本，314 个操作方法 |
| **多版本支持** | ✅ 完整 | 同一 API 的多个版本可共存（如 Catalog v0/v2020/v2022） |
| **类型安全** | ✅ 完整 | 所有 API 都有完整的类型定义 |
| **错误处理** | ✅ 完整 | 统一的错误类型和错误堆栈 |

### 2. 认证和授权

| 功能 | 状态 | 说明 |
|------|------|------|
| **LWA 认证** | ✅ 完整 | Login with Amazon OAuth 2.0 |
| **Regular 操作** | ✅ 完整 | 使用 Refresh Token 的常规操作 |
| **Grantless 操作** | ✅ 完整 | 无需卖家授权的操作（如通知订阅） |
| **Token 缓存** | ✅ 完整 | 自动缓存和刷新 Access Token |
| **Token 提前刷新** | ✅ 完整 | 过期前 1 分钟自动刷新 |
| **RDT 支持** | ✅ 完整 | Restricted Data Token 用于 PII 数据访问 |

### 3. 速率限制

| 功能 | 状态 | 说明 |
|------|------|------|
| **Token Bucket 算法** | ✅ 完整 | 标准的速率限制实现 |
| **突发支持** | ✅ 完整 | 支持配置突发请求数量 |
| **动态更新** | ✅ 完整 | 从 API 响应头自动更新速率配置 |
| **并发安全** | ✅ 完整 | 可在多个 goroutine 中共享 |
| **多维度限流** | ✅ 完整 | 支持按卖家、应用、市场、操作分别限流 |

---

## 🆕 v1.1.0 新增功能

### 4. Go 1.25 分页迭代器

| 功能 | 状态 | 覆盖率 |
|------|------|--------|
| **自动分页** | ✅ 完整 | 27/27 APIs (100%) |
| **Orders API** | ✅ | IterateOrders(), IterateOrderItems() |
| **Reports API** | ✅ | IterateReports() |
| **Feeds API** | ✅ | IterateFeeds() |
| **Catalog Items API** | ✅ | IterateCatalogItems() (3 个版本) |
| **FBA Inventory API** | ✅ | IterateInventorySummaries() |
| **Finances API** | ✅ | IterateFinancialEvents(), IterateFinancialEventGroups() |
| **Fulfillment API** | ✅ | Inbound/Outbound 多个迭代器 |
| **Listings Items API** | ✅ | IterateListingsItems() |
| **Services API** | ✅ | IterateServiceJobs() |
| **Vendor API** | ✅ | 11 个 Vendor API 迭代器 |
| **其他 API** | ✅ | Invoices, Seller Wallet, Supply Sources, Vehicles 等 |

**特性**：
- ✅ 自动处理 NextToken/pageToken
- ✅ 用户代码减少 70%
- ✅ 支持提前退出（break）
- ✅ 流式处理（低内存占用）
- ✅ 完整错误处理

### 5. 报告加密/解密

| 功能 | 状态 | 说明 |
|------|------|------|
| **报告解密** | ✅ 完整 | AES-256-CBC 自动解密 |
| **自动下载** | ✅ 完整 | 从 URL 自动下载报告内容 |
| **加密检测** | ✅ 完整 | 自动判断是否加密 |
| **未加密支持** | ✅ 完整 | 同时支持加密和未加密报告 |
| **PKCS7 填充** | ✅ 完整 | 自动添加/移除填充 |
| **文档加密** | ✅ 完整 | EncryptDocument() 用于上传 |
| **参数验证** | ✅ 完整 | ValidateEncryptionDetails() |

**集成 API**：
- ✅ Reports API - `GetReportDocumentDecrypted()`
- ⚠️ Services API - 加密模块可用，待集成便利方法

---

## 🛠️ 基础设施功能

### 6. HTTP 传输

| 功能 | 状态 | 说明 |
|------|------|------|
| **HTTP/1.1** | ✅ 完整 | 基于标准库 net/http |
| **HTTP/2** | ✅ 完整 | 自动支持 |
| **连接池** | ✅ 完整 | 自动连接复用 |
| **超时控制** | ✅ 完整 | 可配置超时时间 |
| **中间件系统** | ✅ 完整 | 可扩展的请求/响应拦截 |
| **User-Agent** | ✅ 完整 | 自动添加 SDK 标识 |

### 7. 重试机制

| 功能 | 状态 | 说明 |
|------|------|------|
| **自动重试** | ✅ 完整 | 网络错误、5xx、429 自动重试 |
| **指数退避** | ✅ 完整 | 重试间隔指数增长 |
| **最大重试次数** | ✅ 完整 | 可配置（默认 3 次） |
| **可自定义** | ✅ 完整 | 可自定义重试条件 |
| **Context 支持** | ✅ 完整 | 支持取消和超时 |

### 8. 错误处理

| 功能 | 状态 | 说明 |
|------|------|------|
| **统一错误类型** | ✅ 完整 | APIError 包含所有错误信息 |
| **错误堆栈** | ✅ 完整 | 使用 pkg/errors 提供堆栈追踪 |
| **错误分类** | ✅ 完整 | 按 HTTP 状态码分类 |
| **错误上下文** | ✅ 完整 | errors.Wrap 添加上下文信息 |
| **可重试判断** | ✅ 完整 | 自动判断错误是否可重试 |

---

## 📊 数据处理

### 9. JSON 处理

| 功能 | 状态 | 说明 |
|------|------|------|
| **编码/解码** | ✅ 完整 | 基于标准库 encoding/json |
| **验证** | ✅ 完整 | 内置 JSON 验证 |
| **错误处理** | ✅ 完整 | 详细的解析错误信息 |
| **性能优化** | ✅ 完整 | json-iterator (v1.2.0) |

### 10. 数据模型

| 功能 | 状态 | 说明 |
|------|------|------|
| **自动生成** | ✅ 完整 | 从 OpenAPI 规范自动生成 |
| **类型定义** | ✅ 完整 | 所有 API 都有完整类型 |
| **JSON 标签** | ✅ 完整 | 所有字段都有 JSON 标签 |
| **文档注释** | ✅ 完整 | 所有类型都有注释 |

---

## 🔍 可观测性

### 11. 日志

| 功能 | 状态 | 说明 |
|------|------|------|
| **基础日志** | ✅ 有 | 使用标准库 log |
| **结构化日志** | ✅ 完整 | Zap 集成 (v1.2.0) |
| **日志级别** | ✅ 完整 | Debug/Info/Warn/Error |
| **日志中间件** | ✅ 完整 | HTTP 请求/响应日志 |

### 12. 指标

| 功能 | 状态 | 说明 |
|------|------|------|
| **基础指标** | ✅ 完整 | 内置 Metrics 收集器 |
| **API 调用统计** | ✅ 完整 | 统计调用次数、成功/失败率 |
| **Prometheus** | ✅ 完整 | 指标导出 (v1.3.0) |
| **自定义指标** | ✅ 完整 | 可扩展的指标系统 |

### 13. 追踪

| 功能 | 状态 | 说明 |
|------|------|------|
| **分布式追踪** | ✅ 完整 | OpenTelemetry (v1.3.0) |
| **请求 ID** | ✅ 完整 | 每个请求自动生成 ID |
| **调用链** | ✅ 完整 | 支持 Jaeger/Zipkin/Tempo |

---

## 🧪 测试和调试

### 14. 测试支持

| 功能 | 状态 | 说明 |
|------|------|------|
| **单元测试** | ✅ 完整 | 154+ 测试用例 |
| **测试覆盖率** | ✅ 完整 | 核心模块 92%+ |
| **Testify 集成** | ✅ 完整 | 使用 testify 断言 |
| **HTTP Mock** | ⚠️ 待定 | 未来版本考虑 |
| **并发测试** | ✅ 完整 | 所有并发代码都已测试 |

### 15. 调试工具

| 功能 | 状态 | 说明 |
|------|------|------|
| **Debug 模式** | ✅ 完整 | 可配置 Debug 选项 |
| **请求 Dump** | ⚠️ 手动 | 需要用户自己实现 |
| **响应 Dump** | ⚠️ 手动 | 需要用户自己实现 |
| **性能分析** | ✅ 有 | Benchmark 测试 |

---

## 🔧 高级功能

### 16. 并发控制

| 功能 | 状态 | 说明 |
|------|------|------|
| **并发安全** | ✅ 完整 | 所有核心组件并发安全 |
| **Context 支持** | ✅ 完整 | 所有方法支持 context |
| **Goroutine 池** | ⚠️ 示例 | examples/ 中有示例 |
| **并发限制** | ⚠️ 用户实现 | 可使用 golang.org/x/sync/errgroup |

### 17. 配置管理

| 功能 | 状态 | 说明 |
|------|------|------|
| **Functional Options** | ✅ 完整 | 灵活的配置方式 |
| **环境变量** | ✅ 支持 | 可从环境变量加载配置 |
| **配置验证** | ✅ 完整 | 自动验证配置有效性 |
| **配置文件** | ⚠️ 示例 | examples/ 中有 YAML 示例 |
| **热重载** | ⚠️ 待定 | 未来版本考虑 |

### 18. 可靠性

| 功能 | 状态 | 说明 |
|------|------|------|
| **熔断器** | ✅ 完整 | Circuit Breaker (v1.2.0) |
| **降级策略** | ❌ 无 | 用户自己实现 |
| **健康检查** | ⚠️ 示例 | examples/ 中有示例 |
| **优雅退出** | ✅ 完整 | Close() 方法 |

---

## 📦 特定 API 功能

### 19. Orders API

| 功能 | 状态 | 说明 |
|------|------|------|
| **获取订单** | ✅ 完整 | GetOrder(), GetOrders() |
| **订单项** | ✅ 完整 | GetOrderItems() |
| **买家信息** | ✅ 完整 | GetOrderBuyerInfo() (需 RDT) |
| **地址信息** | ✅ 完整 | GetOrderAddress() (需 RDT) |
| **分页迭代** | ✅ 完整 | IterateOrders(), IterateOrderItems() |
| **订单更新** | ✅ 完整 | UpdateShipmentStatus() |

### 20. Reports API

| 功能 | 状态 | 说明 |
|------|------|------|
| **创建报告** | ✅ 完整 | CreateReport() |
| **查询报告** | ✅ 完整 | GetReports(), GetReport() |
| **下载报告** | ✅ 完整 | GetReportDocument() |
| **自动解密** | ✅ 完整 | GetReportDocumentDecrypted() |
| **报告调度** | ✅ 完整 | CreateReportSchedule() |
| **取消报告** | ✅ 完整 | CancelReport() |
| **分页迭代** | ✅ 完整 | IterateReports() |

### 21. Feeds API

| 功能 | 状态 | 说明 |
|------|------|------|
| **创建 Feed** | ✅ 完整 | CreateFeed() |
| **上传文档** | ✅ 完整 | CreateFeedDocument() |
| **查询 Feed** | ✅ 完整 | GetFeeds(), GetFeed() |
| **Feed 结果** | ✅ 完整 | GetFeedDocument() |
| **取消 Feed** | ✅ 完整 | CancelFeed() |
| **分页迭代** | ✅ 完整 | IterateFeeds() |
| **大文件传输** | ✅ 完整 | 上传/下载工具 (v1.2.0) |

### 22. Catalog Items API

| 功能 | 状态 | 说明 |
|------|------|------|
| **搜索商品** | ✅ 完整 | SearchCatalogItems() |
| **获取商品** | ✅ 完整 | GetCatalogItem() |
| **多版本支持** | ✅ 完整 | v0, v2020-12-01, v2022-04-01 |
| **分页迭代** | ✅ 完整 | IterateCatalogItems() (3 个版本) |

### 23. Notifications API

| 功能 | 状态 | 说明 |
|------|------|------|
| **创建目标** | ✅ 完整 | CreateDestination() (SQS/EventBridge) |
| **订阅通知** | ✅ 完整 | CreateSubscription() |
| **管理订阅** | ✅ 完整 | GetSubscription(), DeleteSubscription() |
| **SQS 轮询器** | ⚠️ 示例 | examples/patterns/order-sync-sqs/ |
| **事件解析** | ⚠️ 示例 | 完整的消息解析器示例 |

### 24. FBA Inventory API

| 功能 | 状态 | 说明 |
|------|------|------|
| **库存摘要** | ✅ 完整 | GetInventorySummaries() |
| **分页迭代** | ✅ 完整 | IterateInventorySummaries() |

### 25. Finances API

| 功能 | 状态 | 说明 |
|------|------|------|
| **财务事件** | ✅ 完整 | ListFinancialEvents() |
| **事件组** | ✅ 完整 | ListFinancialEventGroups() |
| **多版本** | ✅ 完整 | v0, v2024-06-19 |
| **分页迭代** | ✅ 完整 | 2 个迭代器方法 |

---

## 🔐 安全功能

### 26. 数据保护

| 功能 | 状态 | 说明 |
|------|------|------|
| **HTTPS** | ✅ 完整 | 所有请求使用 HTTPS |
| **Token 安全** | ✅ 完整 | Token 不写入日志 |
| **RDT 支持** | ✅ 完整 | 受限数据访问 |
| **加密传输** | ✅ 完整 | TLS 1.2+ |

---

## 📖 文档和示例

### 27. 文档

| 类型 | 数量 | 状态 |
|------|------|------|
| **技术文档** | 11 个 | ✅ 完整 |
| **API 文档** | 完整 | ✅ GoDoc 注释 |
| **使用指南** | 3 个 | ✅ 完整 |
| **贡献指南** | 1 个 | ✅ 完整 |

**文档列表**：
- README.md - 项目介绍
- CHANGELOG.md - 版本历史
- UPGRADE_PLAN.md - 升级计划
- docs/ARCHITECTURE.md - 架构设计
- docs/DEVELOPMENT.md - 开发指南
- docs/CODE_STYLE.md - 代码风格
- docs/CONTRIBUTING.md - 贡献指南
- docs/PAGINATION_GUIDE.md - 分页迭代器指南 🆕
- docs/REPORT_DECRYPTION.md - 报告解密指南 🆕
- docs/GRANTLESS_OPERATIONS_GUIDE.md - Grantless 操作
- docs/API_TRACKING.md - API 追踪策略

### 28. 示例代码

| 类型 | 数量 | 状态 |
|------|------|------|
| **基础示例** | 7 个 | ✅ 完整 |
| **高级示例** | 1 个 | ✅ 完整 |
| **生产级示例** | 3 个 | ✅ 完整 |

**示例列表**：
- examples/basic_usage/ - 基础用法
- examples/advanced_usage/ - 高级用法
- examples/orders/ - 订单 API
- examples/reports/ - 报告 API
- examples/feeds/ - Feed API
- examples/listings/ - Listing API
- examples/grantless/ - Grantless 操作
- examples/iterators/ - 迭代器示例 🆕
- examples/report-decryption/ - 报告解密 🆕
- examples/patterns/order-sync-sqs/ - SQS 订单同步 🆕

---

## 🚀 自动化功能

### 29. CI/CD

| 功能 | 状态 | 说明 |
|------|------|------|
| **GitHub Actions** | ✅ 完整 | 自动测试和构建 |
| **自动测试** | ✅ 完整 | 每次提交自动测试 |
| **代码检查** | ✅ 完整 | go vet, golangci-lint |
| **API 监控** | ✅ 完整 | 每日监控官方 API 变更 |
| **自动 Issue** | ✅ 完整 | 检测到变更自动创建 Issue |

### 30. 工具

| 工具 | 状态 | 说明 |
|------|------|------|
| **API 生成器** | ✅ 完整 | 自动生成 API 代码 |
| **迭代器生成器** | ✅ 完整 | 批量生成迭代器 🆕 |
| **监控工具** | ✅ 完整 | cmd/api-monitor/ |
| **验证工具** | ✅ 完整 | tools/validation/ |
| **性能工具** | ✅ 完整 | tools/performance/ |

---

## 🎁 额外功能

### 31. 区域支持

| 区域 | 状态 | Endpoint |
|------|------|----------|
| **北美** | ✅ 完整 | sellingpartnerapi-na.amazon.com |
| **欧洲** | ✅ 完整 | sellingpartnerapi-eu.amazon.com |
| **远东** | ✅ 完整 | sellingpartnerapi-fe.amazon.com |

### 32. 市场支持

| 市场 | 状态 | 说明 |
|------|------|------|
| **美国** | ✅ 完整 | ATVPDKIKX0DER |
| **加拿大** | ✅ 完整 | A2EUQ1WTGCTBG2 |
| **英国** | ✅ 完整 | A1F83G8C2ARO7P |
| **德国** | ✅ 完整 | A1PA6795UKMFR9 |
| **日本** | ✅ 完整 | A1VC38T7YXB528 |
| **所有市场** | ✅ 完整 | 20+ 个市场 |

---

## 📊 功能统计

### 总览

| 类别 | 完成数 | 总数 | 完成率 |
|------|--------|------|--------|
| **核心功能** | 8 | 8 | 100% ✅ |
| **API 支持** | 57 | 57 | 100% ✅ |
| **分页迭代器** | 27 | 27 | 100% ✅ |
| **加密/解密** | 1 | 2 | 50% ⚠️ |
| **可观测性** | 3 | 3 | 100% ✅ |
| **文档** | 11 | 11 | 100% ✅ |
| **示例** | 10 | 10 | 100% ✅ |
| **测试** | 154+ | 154+ | 100% ✅ |

### v1.3.0 状态

- **已实现功能**：38 项
- **计划中功能**：3 项（v2.0.0）
- **核心功能完成度**：100%
- **整体完成度**：92%

---

## 🔮 计划中的功能

### v2.0.0 路线图

未来版本可能添加：

1. ⚠️ **GraphQL 风格查询** - 简化复杂查询
2. ⚠️ **CLI 工具** - spapi 命令行工具
3. ⚠️ **更多生产级示例** - 库存同步、价格监控等

### 已完成的功能

✅ v1.1.0 - Go 1.25 迭代器 + 报告解密  
✅ v1.2.0 - 结构化日志 + 熔断器 + 大文件传输 + 参数验证 + JSON 优化  
✅ v1.3.0 - OpenTelemetry + Prometheus 指标导出

---

## 🎯 功能对比

### vs 官方 Java SDK

| 功能 | 官方 Java SDK | 本 SDK (Go) |
|------|--------------|-------------|
| **API 支持** | ~59 | 57 ✅ |
| **LWA 认证** | ✅ | ✅ |
| **速率限制** | ✅ | ✅ |
| **RDT 支持** | ✅ | ✅ |
| **分页迭代器** | ❌ | ✅ (27 个) |
| **自动解密** | ⚠️ 手动 | ✅ 自动 |
| **零依赖** | ❌ (4+ 依赖) | ⚠️ (4 依赖) |
| **Go 语言** | ❌ | ✅ |

### vs amzapi SDK (68 stars)

| 功能 | amzapi SDK | 本 SDK |
|------|-----------|--------|
| **API 支持** | 20 | 57 ✅ |
| **分页迭代器** | ❌ | ✅ (27 个) |
| **自动解密** | ⚠️ 手动调用 | ✅ 自动 |
| **速率限制** | ❌ | ✅ |
| **RDT 支持** | ❌ | ✅ |
| **测试覆盖** | ⚠️ 部分 | ✅ 92%+ |
| **文档** | ⚠️ 简单 | ✅ 完整 |
| **自动监控** | ❌ | ✅ |

---

## 📞 获取帮助

- **GitHub Issues**: https://github.com/vanling1111/amazon-sp-api-go-sdk/issues
- **文档**: https://github.com/vanling1111/amazon-sp-api-go-sdk/tree/main/docs
- **示例**: https://github.com/vanling1111/amazon-sp-api-go-sdk/tree/main/examples

---

**最后更新**: 2025-10-03 (v1.3.0)

