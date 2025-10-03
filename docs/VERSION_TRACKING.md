# 版本追踪

## 当前版本

**v1.3.0** - 2025-10-03

云原生可观测性版本。

## 版本历史

### v1.3.0 (2025-10-03)

**新增功能：**
- OpenTelemetry 分布式追踪
- Prometheus 指标导出
- 完整的可观测性堆栈（日志 + 追踪 + 指标）

**依赖：**
- OpenTelemetry v1.33.0
- Prometheus client v1.23.2

### v1.2.0 (2025-10-03)

**新增功能：**
- 结构化日志（Zap）
- 熔断器（Circuit Breaker）
- 参数验证（validator）
- JSON 性能优化（3-5倍提升）
- 大文件传输工具

**依赖：**
- Zap v1.27.0
- validator v10.23.0
- json-iterator v1.1.12

### v1.1.0 (2025-10-03)

**新增功能：**
- Go 1.25 分页迭代器（27 个 API）
- 报告自动解密（AES-256-CBC）
- 生产级 SQS 订单同步示例
- 加密模块（internal/crypto）

**依赖：**
- pkg/errors v0.9.1
- testify v1.10.0
- AWS SDK v2

### v1.0.0 (2025-10-03)

**特性：**
- 57 个 API 版本完整支持
- 314 个 API 操作方法
- 完整的认证、签名、速率限制
- 92%+ 测试覆盖率
- 自动 API 监控

**文档：**
- 完整的中文文档
- 7 个使用示例
- 技术文档

## 语义化版本

本项目遵循 [语义化版本 2.0.0](https://semver.org/)：

- **MAJOR** - 不兼容的 API 变更
- **MINOR** - 向后兼容的新功能
- **PATCH** - 向后兼容的 Bug 修复

## 发布流程

1. 更新版本号
2. 更新 CHANGELOG.md
3. 运行完整测试
4. 创建 Git tag
5. 发布到 GitHub

```bash
git tag -a v1.1.0 -m "Release v1.1.0"
git push origin v1.1.0
```

## 升级指南

### 从 v1.0.x 升级到 v1.1.x

无需修改代码，向后兼容。

## 依赖版本

- Go: 1.21+
- 无外部依赖

## 官方 API 版本对应

本 SDK 支持的 Amazon SP-API 版本：

- Orders API: v0
- Feeds API: v2021-06-30
- Reports API: v2021-06-30
- Catalog Items API: v0, v2020-12-01, v2022-04-01
- 其他 53 个 API...

完整列表查看 `pkg/spapi/` 目录。

## 参考

- [Semantic Versioning](https://semver.org/)
- [Keep a Changelog](https://keepachangelog.com/)

