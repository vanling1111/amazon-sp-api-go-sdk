# API Monitoring Tools

监控 Amazon SP-API OpenAPI 规范的变更。

## 工具

### api_monitor.go

自动监控所有 57 个 API 的 OpenAPI 规范文件，检测是否有更新。

**功能**:
- 从 GitHub 获取最新的 OpenAPI 规范
- 计算文件的 SHA256 哈希值
- 与上次检查的哈希值对比
- 检测到变更时发出警告

**使用方式**:

```bash
# 手动运行监控
go run tools/monitoring/api_monitor.go

# 在 CI 中运行（自动）
# 参见 .github/workflows/api-monitor.yml
```

**输出**:
- 无变更: 退出码 0
- 有变更: 退出码 1，并输出变更的 API 列表

**状态文件**:
- `api-state.json` - 存储每个 API 的最后已知哈希值

## GitHub Actions 集成

项目配置了自动化工作流：

### 1. API 监控 (.github/workflows/api-monitor.yml)
- **触发**: 每天 00:00 UTC
- **功能**: 检测 API 规范变更
- **动作**: 发现变更时自动创建 GitHub Issue

### 2. 测试工作流 (.github/workflows/tests.yml)
- **触发**: Push 或 Pull Request
- **功能**: 运行所有测试
- **检查**: 编译、测试、基准测试

## 手动监控

如果需要手动检查特定 API:

```bash
# 下载最新规范
curl -o orders.json \
  https://raw.githubusercontent.com/amzn/selling-partner-api-models/main/models/orders-api-model/ordersV0.json

# 计算哈希
sha256sum orders.json

# 与上次的哈希对比
```

## 更新流程

当检测到 API 规范变更时:

1. **查看变更内容**
   ```bash
   # 对比新旧版本
   diff old-spec.json new-spec.json
   ```

2. **重新生成代码**
   ```powershell
   # Windows
   .\scripts\generate-apis-versioned.ps1
   .\scripts\generate-all-api-clients.ps1
   ```

3. **运行测试**
   ```bash
   go test ./...
   ```

4. **提交更新**
   ```bash
   git add .
   git commit -m "feat: Update API X to version Y"
   git push
   ```

## 配置

监控的 API 列表定义在:
- `api-list.json` - 所有 57 个 API 的配置
- `scripts/api-config.ps1` - PowerShell 脚本使用的配置

## 参考

- [OpenAPI Specs Repository](https://github.com/amzn/selling-partner-api-models)
- [SP-API Release Notes](https://developer-docs.amazon.com/sp-api/docs/sp-api-release-notes)

