# 版本追踪

本文档记录 Amazon SP-API 官方文档和本 SDK 的版本历史。

---

## 官方文档版本

### 文档更新历史

| 日期 | 文档页面 | 变更摘要 | SDK 状态 |
|------|---------|---------|---------|
| 2025-01-10 | [Usage Plans](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits) | 更新速率限制说明 | ✅ v1.2.0 已同步 |
| 2024-12-15 | [Connecting to SP-API](https://developer-docs.amazon.com/sp-api/docs/connecting-to-the-selling-partner-api) | 新增 Grantless 操作说明 | ✅ v1.1.0 已同步 |
| 2024-11-20 | [Tokens API](https://developer-docs.amazon.com/sp-api/docs/tokens-api) | 新增 RDT 数据元素说明 | ✅ v1.0.5 已同步 |

---

### OpenAPI 规范版本

| API | 当前版本 | 规范文件 | 最后更新 | SDK 状态 |
|-----|---------|---------|---------|---------|
| Orders API | v0 | [orders-api-model.json](https://github.com/amzn/selling-partner-api-models/blob/main/models/orders-api-model/ordersV0.json) | 2024-12-01 | ✅ v1.2.0 |
| Reports API | v2021-06-30 | [reports-api-model.json](https://github.com/amzn/selling-partner-api-models/blob/main/models/reports-api-model/reports_2021-06-30.json) | 2024-11-15 | ✅ v1.2.0 |
| Feeds API | v2021-06-30 | [feeds-api-model.json](https://github.com/amzn/selling-partner-api-models/blob/main/models/feeds-api-model/feeds_2021-06-30.json) | 2024-10-20 | ✅ v1.1.0 |
| Listings API | v2021-08-01 | [listings-api-model.json](https://github.com/amzn/selling-partner-api-models/blob/main/models/listings-items-api-model/listingsItems_2021-08-01.json) | 2024-09-30 | 🔄 计划中 |
| Notifications API | v1 | [notifications-api-model.json](https://github.com/amzn/selling-partner-api-models/blob/main/models/notifications-api-model/notificationsV1.json) | 2024-08-15 | 🔄 计划中 |

**图例**:
- ✅ 已实现并同步
- 🔄 计划中
- ❌ 暂不支持

---

## SDK 版本历史

### v1.2.0 - 2025-01-15

**新增功能**:
- ✅ Grantless Operations 支持
- ✅ Notifications API 基础框架

**变更**:
- 📝 根据官方文档更新 LWA 认证流程
- 📝 优化 Token 缓存策略

**修复**:
- 🐛 修复 RDT Signer 的 data elements 提取逻辑
- 🐛 修复并发场景下的 Token 缓存竞态问题

**官方文档变更**:
- [2025-01-10] 官方文档更新了 Rate Limits 说明
  - 链接: https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits
  - 影响: `internal/ratelimit`
  - 状态: ✅ 已同步

**测试**:
- 测试覆盖率: 91.5%
- 新增测试: 45 个

---

### v1.1.0 - 2024-12-20

**新增功能**:
- ✅ Reports API 支持
- ✅ Feeds API 支持

**变更**:
- 📝 重构 HTTP Transport 层
- 📝 优化重试逻辑

**修复**:
- 🐛 修复 LWA Token 过期处理
- 🐛 修复中间件链执行顺序

**官方文档变更**:
- [2024-12-15] 官方文档新增 Grantless 操作说明
  - 链接: https://developer-docs.amazon.com/sp-api/docs/connecting-to-the-selling-partner-api
  - 影响: `internal/auth`
  - 状态: ✅ 已同步 (v1.2.0)

**测试**:
- 测试覆盖率: 89.2%
- 新增测试: 62 个

---

### v1.0.0 - 2024-11-01

**初始版本**:
- ✅ LWA 认证
- ✅ HTTP Transport
- ✅ Request Signing (LWA + RDT)
- ✅ Rate Limiting
- ✅ Orders API

**官方文档依据**:
- [Connecting to SP-API](https://developer-docs.amazon.com/sp-api/docs/connecting-to-the-selling-partner-api)
- [Usage Plans and Rate Limits](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits)
- [Tokens API](https://developer-docs.amazon.com/sp-api/docs/tokens-api)
- [Orders API](https://developer-docs.amazon.com/sp-api/docs/orders-api-v0-reference)

**测试**:
- 测试覆盖率: 87.5%
- 单元测试: 145 个
- 集成测试: 12 个

---

## 版本规划

### v1.3.0 - 计划中 (2025-02)

**计划新增**:
- [ ] Listings API 完整支持
- [ ] Notifications API 完整支持
- [ ] WebSocket 支持（Notifications）

**计划优化**:
- [ ] 性能优化
- [ ] 内存优化
- [ ] 并发优化

---

### v1.4.0 - 计划中 (2025-03)

**计划新增**:
- [ ] Catalog Items API
- [ ] Product Pricing API
- [ ] FBA Inventory API

---

### v2.0.0 - 长期计划

**重大变更**:
- [ ] 移除所有已弃用的 API
- [ ] 重新设计公开 API 接口
- [ ] 支持更多高级特性

---

## 官方文档监控

### 自动监控

**GitHub Actions**:
- `.github/workflows/doc-check.yml` - 每天检查文档更新
- `.github/workflows/openapi-sync.yml` - 每周检查 OpenAPI 规范

**监控工具**:
- `tools/monitoring/api_monitor.go` - 文档内容哈希监控
- `tools/monitoring/openapi_sync.go` - OpenAPI 规范同步

---

### 手动检查清单

**每周检查**:
- [ ] 访问 [What's New](https://developer-docs.amazon.com/sp-api/docs/welcome)
- [ ] 检查 [OpenAPI 规范仓库](https://github.com/amzn/selling-partner-api-models/commits/main)
- [ ] 运行 `go run tools/monitoring/api_monitor.go`

**每月检查**:
- [ ] 审查所有 API 文档页面
- [ ] 验证 SDK 实现符合性
- [ ] 更新测试用例

**发布前检查**:
- [ ] 运行完整测试套件
- [ ] 验证所有 OpenAPI 规范是最新的
- [ ] 更新 CHANGELOG
- [ ] 更新本文档

---

## 兼容性策略

### 向后兼容

**保证**:
- ✅ 同一主版本内保持 API 兼容性
- ✅ 废弃的 API 至少保留一个主版本周期
- ✅ 所有重大变更在 CHANGELOG 中明确标注

**示例**:
```go
// v1.1.0 引入新方法
func (c *Client) GetOrdersV2(ctx context.Context, req *GetOrdersRequest) (*GetOrdersResponse, error)

// v1.2.0 标记旧方法为废弃
// Deprecated: Use GetOrdersV2 instead. Will be removed in v2.0.0.
func (c *Client) GetOrders(ctx context.Context, req *GetOrdersRequest) (*GetOrdersResponse, error)

// v2.0.0 移除旧方法
// GetOrders 方法已被移除，请使用 GetOrdersV2
```

---

### 迁移指南

**大版本升级时提供**:
1. 详细的迁移文档
2. 代码示例对比
3. 自动化迁移工具（如果可能）

---

## 发布流程

### 1. 准备发布

```bash
# 1. 确保所有测试通过
make test
make test-integration

# 2. 确保 linter 通过
make lint

# 3. 检查文档更新
go run tools/monitoring/api_monitor.go

# 4. 检查 OpenAPI 规范
go run tools/monitoring/openapi_sync.go --check
```

---

### 2. 更新版本号

**更新文件**:
- `go.mod` - 模块版本
- `CHANGELOG.md` - 变更日志
- `VERSION_TRACKING.md` - 本文档

---

### 3. 创建 Git Tag

```bash
# 创建标签
git tag -a v1.2.0 -m "Release v1.2.0

New Features:
- Grantless Operations support
- Notifications API framework

Changes:
- Update LWA auth flow per official docs
- Optimize token caching

Fixes:
- Fix RDT Signer data elements extraction
- Fix token cache race condition

Official Docs Changes:
- [2025-01-10] Rate Limits updated
"

# 推送标签
git push origin v1.2.0
```

---

### 4. 发布到 GitHub

**GitHub Release**:
1. 在 GitHub 上创建 Release
2. 附上 CHANGELOG
3. 标注重大变更
4. 提供迁移指南（如果需要）

---

### 5. 发布公告

**渠道**:
- GitHub Discussions
- README 更新
- 社区通知

---

## 依赖版本

### Go 版本

| SDK 版本 | 最低 Go 版本 | 推荐 Go 版本 |
|---------|------------|------------|
| v1.0.x  | 1.21       | 1.21       |
| v1.1.x  | 1.21       | 1.21       |
| v1.2.x  | 1.21       | 1.21       |
| v2.0.x  | 1.22       | 1.22       |

---

### 第三方依赖

**原则**: 尽量不依赖第三方库，只使用 Go 标准库

**例外**:
- 测试工具: `testify` (可选)
- 代码生成: `openapi-generator` (开发工具)

---

## 参考资料

- [语义化版本](https://semver.org/lang/zh-CN/)
- [Keep a Changelog](https://keepachangelog.com/zh-CN/1.0.0/)
- [Amazon SP-API 官方文档](https://developer-docs.amazon.com/sp-api/docs/)
- [OpenAPI 规范仓库](https://github.com/amzn/selling-partner-api-models)

