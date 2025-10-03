# API 追踪策略

## 概述

本文档说明如何追踪 Amazon SP-API 的更新。

## 自动监控

项目配置了 GitHub Actions 自动监控：

### 监控内容
- 57 个 API 的 OpenAPI 规范
- 每日自动检查
- 检测到变更时自动创建 Issue

### 工作流
- 文件：`.github/workflows/api-monitor.yml`
- 时间：每天 UTC 00:00
- 工具：`cmd/api-monitor/main.go`

## 监控的 API

监控所有 57 个 API 版本，包括：
- Orders, Feeds, Reports
- Catalog Items, Listings
- FBA Inventory, Fulfillment
- Product Pricing, Fees
- Finances, Seller Wallet
- 所有 Vendor API
- 等等

## 变更处理

当检测到 API 变更时：

1. 系统自动创建 GitHub Issue
2. 审查变更内容
3. 更新受影响的代码
4. 运行测试
5. 发布新版本

## 版本管理

遵循语义化版本：
- **MAJOR** - 不兼容的 API 变更
- **MINOR** - 新增功能
- **PATCH** - Bug 修复

## 工具使用

```bash
# 手动运行监控
cd cmd/api-monitor
go run main.go
```

## 参考资源

- [OpenAPI 规范仓库](https://github.com/amzn/selling-partner-api-models)
- [SP-API 文档](https://developer-docs.amazon.com/sp-api/docs/)
- [语义化版本](https://semver.org/)
