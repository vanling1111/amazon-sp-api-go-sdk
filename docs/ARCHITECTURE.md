# 架构设计

## 概述

本 SDK 采用分层架构设计，严格基于 [Amazon SP-API 官方文档](https://developer-docs.amazon.com/sp-api/docs/) 实现，充分利用 Go 语言特性。

## 设计原则

### 1. 官方文档驱动
- ✅ 唯一权威来源：Amazon SP-API 官方文档
- ❌ 不参考其他语言的官方 SDK 实现
- ✅ 直接访问和验证官方文档内容
- ✅ 所有实现都可追溯到官方文档章节

### 2. Go 语言最佳实践
- Context-First：所有方法接受 `context.Context`
- Interface-Driven：关键组件通过接口定义
- Functional Options：灵活的配置模式
- 并发安全：所有共享资源使用 `sync.Mutex` 保护

### 3. 高质量代码
- 测试覆盖率 > 90%
- 完整的错误处理和错误包装
- Google 风格的中文注释
- 完整的示例和文档

---

## 分层架构

```
┌─────────────────────────────────────────────────┐
│          Public API Layer (spapi)               │
│  - Orders API, Reports API, Feeds API...        │
└─────────────────────────────────────────────────┘
                      ↓
┌─────────────────────────────────────────────────┐
│         Codec Layer (internal/codec)            │
│  - Request Encoding, Response Decoding          │
└─────────────────────────────────────────────────┘
                      ↓
┌─────────────────────────────────────────────────┐
│      Rate Limit Layer (internal/ratelimit)      │
│  - Token Bucket, Request Throttling             │
└─────────────────────────────────────────────────┘
                      ↓
┌─────────────────────────────────────────────────┐
│       Signer Layer (internal/signer)            │
│  - LWA Signer, RDT Signer, Chain Signer         │
└─────────────────────────────────────────────────┘
                      ↓
┌─────────────────────────────────────────────────┐
│     Transport Layer (internal/transport)        │
│  - HTTP Client, Middleware, Retry Logic         │
└─────────────────────────────────────────────────┘
                      ↓
┌─────────────────────────────────────────────────┐
│         Auth Layer (internal/auth)              │
│  - LWA Authentication, Token Cache              │
└─────────────────────────────────────────────────┘
```

---

## 核心组件

### 1. Auth Layer - 认证层

**职责**：
- LWA (Login with Amazon) 认证
- 访问令牌获取和缓存
- 支持 Regular 和 Grantless 操作

**官方文档**：
- [Connect to the SP-API - Step 1](https://developer-docs.amazon.com/sp-api/docs/connecting-to-the-selling-partner-api#step-1-request-a-login-with-amazon-access-token)

**主要接口**：
```go
type Client interface {
    GetAccessToken(ctx context.Context) (string, error)
    RefreshToken(ctx context.Context) (string, error)
}
```

**实现细节**：
- 请求格式：`application/x-www-form-urlencoded`
- 响应格式：JSON
- 缓存策略：内存缓存 + 提前过期（60秒）
- Grant Types：
  - `refresh_token` - 需要卖家授权的操作
  - `client_credentials` - Grantless 操作

---

### 2. Transport Layer - 传输层

**职责**：
- HTTP 请求发送
- 中间件管理
- 重试逻辑

**官方文档**：
- [Connect to the SP-API](https://developer-docs.amazon.com/sp-api/docs/connecting-to-the-selling-partner-api)

**主要接口**：
```go
type Client interface {
    Do(ctx context.Context, req *http.Request) (*http.Response, error)
    Use(middleware Middleware)
}
```

**内置中间件**：
- `UserAgentMiddleware` - 设置 User-Agent 头
- `DateMiddleware` - 设置 x-amz-date 头
- `HeaderMiddleware` - 添加自定义头
- `LoggingMiddleware` - 请求/响应日志
- `RetryMiddleware` - 重试逻辑

---

### 3. Signer Layer - 签名层

**职责**：
- 请求签名
- LWA 令牌注入
- RDT (Restricted Data Token) 处理

**官方文档**：
- [Connect to the SP-API - Step 3](https://developer-docs.amazon.com/sp-api/docs/connecting-to-the-selling-partner-api#step-3-add-headers-to-the-uri)
- [Tokens API](https://developer-docs.amazon.com/sp-api/docs/tokens-api)

**主要接口**：
```go
type Signer interface {
    Sign(ctx context.Context, req *http.Request) error
}
```

**实现**：
- `LWASigner` - 添加 `x-amz-access-token` 头（常规操作）
- `RDTSigner` - 添加 `x-amz-access-token` 头（受限操作，使用 RDT）
- `ChainSigner` - 组合多个签名器

---

### 4. Rate Limit Layer - 速率限制层

**职责**：
- 请求速率控制
- Token Bucket 算法
- 防止超出 API 限制

**官方文档**：
- [Usage Plans and Rate Limits](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits)

**实现**：
- 基于 Token Bucket 算法
- 支持动态速率调整
- 支持多个速率限制器

---

### 5. Codec Layer - 编解码层

**职责**：
- 请求序列化
- 响应反序列化
- 数据验证

**实现**：
- JSON 编解码
- 错误响应解析
- 数据类型转换

---

### 6. Public API Layer - 公开 API 层

**职责**：
- 提供类型安全的 API
- 封装底层复杂性
- 提供友好的开发体验

**示例**：
```go
// Orders API
type OrdersAPI interface {
    GetOrders(ctx context.Context, req *GetOrdersRequest) (*GetOrdersResponse, error)
    GetOrder(ctx context.Context, orderID string) (*Order, error)
}
```

---

## 数据流

### 请求流程

```
用户代码
  ↓
Public API Layer
  ↓
Codec Layer (Encode Request)
  ↓
Rate Limit Layer (Check & Wait)
  ↓
Signer Layer (Add Auth Headers)
  ↓
Transport Layer (Send HTTP Request)
  ↓
Auth Layer (Provide Token)
  ↓
Amazon SP-API
```

### 响应流程

```
Amazon SP-API
  ↓
Transport Layer (Receive Response)
  ↓
Transport Layer (Retry if needed)
  ↓
Codec Layer (Decode Response)
  ↓
Public API Layer
  ↓
用户代码
```

---

## 错误处理

### 错误类型

```go
// 认证错误
type AuthError struct {
    Code    string
    Message string
}

// API 错误
type APIError struct {
    Code    string
    Message string
    Details map[string]interface{}
}

// 速率限制错误
type RateLimitError struct {
    RetryAfter time.Duration
}

// 网络错误
type NetworkError struct {
    Err error
}
```

### 错误处理策略

1. **可重试错误**：
   - 5xx 服务器错误
   - 429 速率限制错误
   - 网络超时错误

2. **不可重试错误**：
   - 4xx 客户端错误（除 429）
   - 认证错误
   - 参数验证错误

3. **错误包装**：
   - 使用 `fmt.Errorf("...: %w", err)` 包装错误
   - 保留错误链以便调试

---

## 并发安全

### 线程安全组件

1. **Token Cache**：
   - 使用 `sync.RWMutex` 保护
   - 支持并发读取

2. **Rate Limiter**：
   - 使用 `sync.Mutex` 保护
   - Token bucket 操作原子性

3. **HTTP Client**：
   - Go 标准库 `http.Client` 本身是并发安全的
   - 连接池自动管理

---

## 性能优化

### 1. 连接复用
```go
// HTTP/2 支持
transport := &http.Transport{
    MaxIdleConns:        100,
    MaxIdleConnsPerHost: 10,
    IdleConnTimeout:     90 * time.Second,
}
```

### 2. 令牌缓存
- 内存缓存避免重复 LWA 请求
- 提前 60 秒过期确保令牌有效性

### 3. 请求批处理
- 支持批量操作减少 API 调用
- 自动处理分页

---

## 扩展性

### 添加新 API

1. 在 `spapi/` 下创建新的 API 包
2. 定义请求/响应结构
3. 实现 API 接口
4. 添加测试和文档

### 自定义中间件

```go
func CustomMiddleware() transport.Middleware {
    return func(next transport.Handler) transport.Handler {
        return func(ctx context.Context, req *http.Request) (*http.Response, error) {
            // 前置处理
            resp, err := next(ctx, req)
            // 后置处理
            return resp, err
        }
    }
}
```

### 自定义签名器

```go
type CustomSigner struct {}

func (s *CustomSigner) Sign(ctx context.Context, req *http.Request) error {
    // 自定义签名逻辑
    return nil
}
```

---

## 测试策略

### 1. 单元测试
- 每个组件独立测试
- Mock 外部依赖
- 覆盖率 > 90%

### 2. 集成测试
- 测试组件协作
- 使用 Sandbox 环境

### 3. 性能测试
- 并发测试
- 压力测试
- 内存泄漏检测

---

## 参考资料

- [Amazon SP-API 官方文档](https://developer-docs.amazon.com/sp-api/docs/)
- [Go 官方文档](https://go.dev/doc/)
- [Google Go 风格指南](https://google.github.io/styleguide/go/)

