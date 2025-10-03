# API 追踪策略

## 概述

本文档说明如何追踪和同步 Amazon SP-API 官方文档的更新，确保 SDK 始终与官�?API 保持一致�?

## 核心原则

### �?唯一权威来源
- **官方文档**: https://developer-docs.amazon.com/sp-api/docs/
- **官方参�?*: https://developer-docs.amazon.com/sp-api/reference/
- **官方 OpenAPI 规范**: https://github.com/amzn/selling-partner-api-models

### �?不参考的资源
- 其他语言的官�?SDK（Java、Python、Node.js、C#、PHP 等）
- 第三方实�?
- Stack Overflow 讨论（除非引用官方文档）

---

## 监控目标

### 1. 官方文档更新
**监控页面**:
- [Welcome Guide](https://developer-docs.amazon.com/sp-api/docs/welcome)
- [Connecting to SP-API](https://developer-docs.amazon.com/sp-api/docs/connecting-to-the-selling-partner-api)
- [Usage Plans](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits)
- [API Reference](https://developer-docs.amazon.com/sp-api/reference/)

**监控方式**:
- 定期访问官方文档网站
- 提取关键章节内容
- 计算内容哈希值检测变�?

### 2. OpenAPI 规范更新
**GitHub 仓库**:
- https://github.com/amzn/selling-partner-api-models

**监控方式**:
- 监控 `models/` 目录下的 JSON 文件
- 使用 GitHub API 获取最�?commit
- 对比文件内容差异

---

## 自动化工�?

### 1. 文档监控工具

**位置**: `tools/monitoring/api_monitor.go`

**功能**:
- 定期访问官方文档页面
- 提取关键信息
- 检测内容变�?
- 发送通知

**使用方式**:
```bash
go run tools/monitoring/api_monitor.go
```

**配置** (`config/monitor.yml`):
```yaml
interval: 24h  # 检查间�?
pages:
  - url: https://developer-docs.amazon.com/sp-api/docs/connecting-to-the-selling-partner-api
    selectors:
      - "#main-content"
  - url: https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits
    selectors:
      - "#main-content"

notifications:
  - type: github-issue
    repo: vanling1111/amazon-sp-api-go-sdk
```

---

### 2. OpenAPI 规范同步工具

**位置**: `tools/monitoring/openapi_sync.go`

**功能**:
- �?GitHub 拉取最�?OpenAPI 规范
- 对比本地版本
- 标记需要更新的模型

**使用方式**:
```bash
go run tools/monitoring/openapi_sync.go --check
go run tools/monitoring/openapi_sync.go --sync
```

---

## GitHub Actions 工作�?

### 1. 文档更新检�?

**文件**: `.github/workflows/doc-check.yml`

```yaml
name: Documentation Update Check

on:
  schedule:
    # 每天 UTC 时间 00:00 运行
    - cron: '0 0 * * *'
  workflow_dispatch:  # 支持手动触发

jobs:
  check-docs:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.21'
      
      - name: Run Documentation Monitor
        run: go run tools/monitoring/api_monitor.go
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
      
      - name: Create Issue if Changes Detected
        if: ${{ steps.monitor.outputs.changed == 'true' }}
        uses: actions/github-script@v7
        with:
          script: |
            github.rest.issues.create({
              owner: context.repo.owner,
              repo: context.repo.repo,
              title: '🚨 官方 SP-API 文档已更�?,
              body: '检测到官方文档有更新，请检查并同步修改。\n\n详情�? ${{ steps.monitor.outputs.details }}',
              labels: ['documentation', 'needs-review']
            })
```

---

### 2. OpenAPI 规范同步

**文件**: `.github/workflows/openapi-sync.yml`

```yaml
name: OpenAPI Spec Sync

on:
  schedule:
    # 每周一 UTC 时间 00:00 运行
    - cron: '0 0 * * 1'
  workflow_dispatch:

jobs:
  sync-openapi:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.21'
      
      - name: Check OpenAPI Updates
        id: check
        run: |
          go run tools/monitoring/openapi_sync.go --check
      
      - name: Create PR if Updates Found
        if: ${{ steps.check.outputs.changed == 'true' }}
        uses: peter-evans/create-pull-request@v6
        with:
          commit-message: 'chore: sync OpenAPI specifications'
          title: '🔄 同步 OpenAPI 规范'
          body: |
            自动同步官方 OpenAPI 规范�?
            
            **变更文件**:
            ${{ steps.check.outputs.files }}
            
            **请审�?*:
            - [ ] 检查模型变�?
            - [ ] 更新相关代码
            - [ ] 更新测试
            - [ ] 更新文档
          branch: sync-openapi
          labels: |
            openapi
            automated
```

---

## 手动检查流�?

### 1. 每周检查（推荐�?

**检查项**:
- [ ] 访问官方文档首页，查�?"What's New" 部分
- [ ] 检�?OpenAPI 规范仓库的最�?commit
- [ ] 查看官方 SDK �?Release Notes（仅作参考，不参考代码）

**操作步骤**:
```bash
# 1. 运行文档监控工具
go run tools/monitoring/api_monitor.go

# 2. 运行 OpenAPI 同步检�?
go run tools/monitoring/openapi_sync.go --check

# 3. 如果有变更，查看详情
go run tools/monitoring/openapi_sync.go --diff
```

---

### 2. 发布前检查（必须�?

**发布新版本前必须执行**:
```bash
# 1. 确保所有文档是最新的
go run tools/monitoring/api_monitor.go --force-check

# 2. 同步 OpenAPI 规范
go run tools/monitoring/openapi_sync.go --sync

# 3. 运行完整测试
make test
make test-integration

# 4. 更新 CHANGELOG
# 记录所�?API 变更
```

---

## 变更处理流程

### 1. 发现文档变更

**步骤**:
1. 访问变更的文档页�?
2. 提取变更的内�?
3. 创建 GitHub Issue 记录变更
4. 标记需要同步的模块

**Issue 模板**:
```markdown
## 📄 官方文档更新

**变更页面**: [页面 URL]

**变更摘要**:
- 变更 1: ...
- 变更 2: ...

**影响模块**:
- [ ] internal/auth
- [ ] internal/signer
- [ ] internal/transport
- [ ] pkg/spapi

**处理计划**:
1. 阅读完整变更内容
2. 评估影响范围
3. 更新代码实现
4. 更新测试
5. 更新文档

**参�?*:
- 官方文档: [URL]
- 变更详情: [详细描述]
```

---

### 2. OpenAPI 规范变更

**步骤**:
1. 运行 `openapi_sync.go --diff` 查看变更
2. 评估变更影响（新增、修改、删除）
3. 重新生成受影响的模型
4. 更新相关代码
5. 更新测试

**示例**:
```bash
# 查看变更
go run tools/monitoring/openapi_sync.go --diff models/orders-api-model.json

# 输出示例:
# Changes detected in orders-api-model.json:
# + Added: Order.BuyerInfo.BuyerCounty
# * Modified: Order.OrderStatus (new value: "Shipped")
# - Removed: Order.DeprecatedField

# 重新生成模型
go run cmd/generator/main.go -input models/orders-api-model.json -output api/orders

# 运行测试
go test ./api/orders/... -v
```

---

### 3. API 新增或废�?

**新增 API**:
1. 获取�?API �?OpenAPI 规范
2. 生成 Go 模型
3. 实现 API 客户�?
4. 添加测试和示�?
5. 更新文档

**废弃 API**:
1. 在代码中标记�?`Deprecated`
2. 添加弃用警告日志
3. 更新文档说明替代方案
4. 在下一个主版本中移�?

**示例**:
```go
// Deprecated: GetOrderMetrics 已被官方废弃，请使用 GetOrderMetricsV2
//
// 官方文档: https://developer-docs.amazon.com/sp-api/docs/...
//
// 此方法将�?v2.0.0 中移�?
func (c *Client) GetOrderMetrics(ctx context.Context, req *GetOrderMetricsRequest) (*GetOrderMetricsResponse, error) {
    log.Warn("GetOrderMetrics is deprecated, use GetOrderMetricsV2 instead")
    // ...
}
```

---

## 版本管理

### 语义化版�?

**规则**:
- `MAJOR.MINOR.PATCH`
- **MAJOR**: 不兼容的 API 变更
- **MINOR**: 向后兼容的功能新�?
- **PATCH**: 向后兼容�?Bug 修复

**示例**:
```
v1.0.0  - 初始版本
v1.1.0  - 新增 Listings API
v1.1.1  - 修复 Orders API �?bug
v2.0.0  - 移除已弃用的 API
```

---

### CHANGELOG 管理

**格式**:
```markdown
# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

## [1.2.0] - 2025-01-15

### Added
- 新增 Notifications API 支持 (#123)
- 新增 Grantless Operations 支持 (#124)

### Changed
- 根据官方文档更新 LWA 认证流程 (#125)
- 优化 Token 缓存策略 (#126)

### Fixed
- 修复 RDT Signer �?data elements 提取逻辑 (#127)

### Official Documentation Changes
- [2025-01-10] 官方文档更新�?Rate Limits 说明
  - 链接: https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits
  - 影响: internal/ratelimit
  - 状�? �?已同�?
```

---

## 通知机制

### 1. GitHub Issues
自动创建 Issue 追踪文档变更�?

### 2. GitHub Discussions
重大变更�?Discussions 中讨论�?

### 3. Release Notes
每次发布时包含完整的变更说明�?

---

## 工具实现参�?

### 文档内容哈希

```go
package monitoring

import (
    "crypto/sha256"
    "fmt"
    "io"
)

// CalculateContentHash 计算文档内容哈希
func CalculateContentHash(content string) string {
    h := sha256.New()
    io.WriteString(h, content)
    return fmt.Sprintf("%x", h.Sum(nil))
}
```

### HTTP 文档获取

```go
package monitoring

import (
    "context"
    "io"
    "net/http"
)

// FetchDocumentContent 获取文档内容
func FetchDocumentContent(ctx context.Context, url string) (string, error) {
    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return "", err
    }

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    return string(body), nil
}
```

---

## 参考资�?

- [Amazon SP-API 官方文档](https://developer-docs.amazon.com/sp-api/docs/)
- [OpenAPI 规范仓库](https://github.com/amzn/selling-partner-api-models)
- [语义化版本](https://semver.org/lang/zh-CN/)
- [Keep a Changelog](https://keepachangelog.com/zh-CN/1.0.0/)

