# 架构设计

## 概述

本 SDK 采用分层架构设计，基于 Amazon SP-API 官方规范实现。

## 设计原则

1. **基于官方规范** - 从 OpenAPI 规范自动生成代码
2. **Go 最佳实践** - 遵循 Go 社区惯用法
3. **高质量** - 完整的测试和错误处理
4. **零依赖** - 仅使用 Go 标准库

## 分层架构

```
┌─────────────────────────────────┐
│  Public API Layer (pkg/spapi)   │
│  - Orders, Feeds, Reports...    │
└─────────────────────────────────┘
           ↓
┌─────────────────────────────────┐
│  Rate Limit (internal/ratelimit)│
│  - Token Bucket Algorithm       │
└─────────────────────────────────┘
           ↓
┌─────────────────────────────────┐
│  Signer (internal/signer)       │
│  - LWA & RDT Signing            │
└─────────────────────────────────┘
           ↓
┌─────────────────────────────────┐
│  Transport (internal/transport) │
│  - HTTP Client & Middleware     │
└─────────────────────────────────┘
           ↓
┌─────────────────────────────────┐
│  Auth (internal/auth)           │
│  - LWA Authentication           │
└─────────────────────────────────┘
```

## 核心组件

### Auth Layer - 认证层
- LWA 认证
- 令牌缓存
- Grantless 操作支持

### Transport Layer - 传输层
- HTTP 请求/响应
- 中间件系统
- 重试逻辑

### Signer Layer - 签名层
- 请求签名
- RDT 支持

### Rate Limit Layer - 速率限制
- Token Bucket 算法
- 动态速率调整

### Public API Layer - 公开 API
- 类型安全的 API 接口
- 57 个 API 版本支持

## 错误处理

- 完整的错误类型定义
- 错误包装和链式传递
- 可重试错误判断

## 并发安全

所有共享资源使用 `sync.Mutex` 保护，确保并发安全。

## 参考

- [Amazon SP-API 文档](https://developer-docs.amazon.com/sp-api/docs/)
- [Go 官方文档](https://go.dev/doc/)

