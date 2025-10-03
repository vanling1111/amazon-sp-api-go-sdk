# Grantless Operations 使用指南

## 概述

Grantless operations 是 Amazon SP-API 中不需要卖家授权的特殊操作。这些操作使用 `grant_type=client_credentials` 而不是 `refresh_token`。

**官方文档**: [https://developer-docs.amazon.com/sp-api/docs/grantless-operations](https://developer-docs.amazon.com/sp-api/docs/grantless-operations)

---

## 支持的 Grantless Operations

根据[官方 SP-API 文档](https://developer-docs.amazon.com/sp-api/docs/connecting-to-the-selling-partner-api)，目前支持的 grantless operations 包括：

### 1. Notifications API
**Scope**: `sellingpartnerapi::notifications`

用于接收和管理 SP-API 通知。

### 2. Application Management API
**Scope**: `sellingpartnerapi::client_credential:rotation`

用于应用程序凭据的轮换和管理。

---

## 使用方法

### 1. 创建 Grantless Credentials

```go
package main

import (
    "log"
    "github.com/amazon-sp-api-go-sdk/internal/auth"
)

func main() {
    // Notifications API
    creds, err := auth.NewGrantlessCredentials(
        "your-client-id",
        "your-client-secret",
        []string{auth.ScopeNotifications},  // 使用预定义的 scope 常量
        auth.EndpointNA,
    )
    if err != nil {
        log.Fatal(err)
    }

    // 创建 LWA 客户端
    client := auth.NewClient(creds)

    // ... 使用 client
}
```

### 2. 使用多个 Scopes

```go
// 同时使用多个 scopes（如果需要）
creds, err := auth.NewGrantlessCredentials(
    "your-client-id",
    "your-client-secret",
    []string{
        auth.ScopeNotifications,
        auth.ScopeCredentialRotation,
    },
    auth.EndpointNA,
)
```

### 3. 获取访问令牌

```go
// 获取访问令牌（自动缓存）
ctx := context.Background()
accessToken, err := client.GetAccessToken(ctx)
if err != nil {
    log.Fatal(err)
}

// 使用访问令牌调用 API
// ...
```

---

## 与常规操作的区别

| 特性 | Regular Operations | Grantless Operations |
|------|-------------------|---------------------|
| **Grant Type** | `refresh_token` | `client_credentials` |
| **需要卖家授权** | ✅ 是 | ❌ 否 |
| **Refresh Token** | ✅ 必需 | ❌ 不需要 |
| **Scopes** | ❌ 不需要 | ✅ 必需 |
| **使用场景** | 大多数 SP-API 操作 | Notifications API, Application Management API |

---

## 代码示例

### 完整示例：使用 Notifications API

```go
package main

import (
    "context"
    "log"
    "net/http"

    "github.com/amazon-sp-api-go-sdk/internal/auth"
    "github.com/amazon-sp-api-go-sdk/internal/signer"
)

func main() {
    // 1. 创建 grantless credentials
    creds, err := auth.NewGrantlessCredentials(
        "your-client-id",
        "your-client-secret",
        []string{auth.ScopeNotifications},
        auth.EndpointNA,
    )
    if err != nil {
        log.Fatalf("Failed to create credentials: %v", err)
    }

    // 2. 创建 LWA 客户端
    lwaClient := auth.NewClient(creds)

    // 3. 创建签名器
    signer := signer.NewLWASigner(lwaClient)

    // 4. 创建 HTTP 请求
    req, err := http.NewRequest(
        http.MethodGet,
        "https://sellingpartnerapi-na.amazon.com/notifications/v1/subscriptions/ORDER_CHANGE",
        nil,
    )
    if err != nil {
        log.Fatalf("Failed to create request: %v", err)
    }

    // 5. 签名请求
    ctx := context.Background()
    if err := signer.Sign(ctx, req); err != nil {
        log.Fatalf("Failed to sign request: %v", err)
    }

    // 6. 发送请求
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Fatalf("Failed to send request: %v", err)
    }
    defer resp.Body.Close()

    // 7. 处理响应
    log.Printf("Response status: %d", resp.StatusCode)
}
```

---

## 令牌缓存

Grantless operations 的访问令牌也会被自动缓存，缓存键基于：
- Client ID
- Scopes

**缓存键格式**: `lwa_token:{client_id}:grantless:{scopes}`

示例:
```
lwa_token:amzn1.application.xxxx:grantless:sellingpartnerapi::notifications
```

---

## 常见问题

### 1. 什么时候使用 grantless credentials？

**使用 grantless credentials**:
- ✅ Notifications API
- ✅ Application Management API

**使用 regular credentials**:
- ✅ 所有其他 SP-API 操作（Orders, Catalog, Listings 等）

### 2. 可以同时使用 RefreshToken 和 Scopes 吗？

❌ **不可以**。凭据验证会返回 `ErrBothRefreshTokenAndScopes` 错误。

```go
// ❌ 错误：同时设置了 RefreshToken 和 Scopes
creds := &auth.Credentials{
    ClientID:     "client-id",
    ClientSecret: "client-secret",
    RefreshToken: "refresh-token",  // ❌
    Scopes:       []string{auth.ScopeNotifications},  // ❌
    Endpoint:     auth.EndpointNA,
}

err := creds.Validate()
// err == ErrBothRefreshTokenAndScopes
```

### 3. 如何判断凭据是否为 grantless？

```go
if creds.IsGrantless() {
    log.Println("This is a grantless credential")
} else {
    log.Println("This is a regular credential")
}
```

### 4. Grantless 令牌的有效期是多久？

根据官方文档，grantless 令牌的有效期与 regular 令牌相同，通常为 **1 小时** (3600 秒)。

---

## 最佳实践

### 1. 使用预定义的 Scope 常量

✅ **推荐**:
```go
scopes := []string{auth.ScopeNotifications}
```

❌ **不推荐**:
```go
scopes := []string{"sellingpartnerapi::notifications"}  // 硬编码字符串
```

### 2. 为不同的 Scopes 创建不同的客户端

```go
// Notifications 客户端
notifCreds, _ := auth.NewGrantlessCredentials(
    clientID, clientSecret,
    []string{auth.ScopeNotifications},
    auth.EndpointNA,
)
notifClient := auth.NewClient(notifCreds)

// Application Management 客户端
appCreds, _ := auth.NewGrantlessCredentials(
    clientID, clientSecret,
    []string{auth.ScopeCredentialRotation},
    auth.EndpointNA,
)
appClient := auth.NewClient(appCreds)
```

### 3. 重用客户端实例

```go
// ✅ 好：创建一次，重复使用
var notifClient = auth.NewClient(notifCreds)

func handleNotification1() {
    token, _ := notifClient.GetAccessToken(ctx)
    // ...
}

func handleNotification2() {
    token, _ := notifClient.GetAccessToken(ctx)  // 使用缓存的令牌
    // ...
}
```

---

## API 参考

### 常量

```go
// Scopes
const (
    ScopeNotifications      = "sellingpartnerapi::notifications"
    ScopeCredentialRotation = "sellingpartnerapi::client_credential:rotation"
)
```

### 函数

```go
// NewGrantlessCredentials 创建 grantless credentials
func NewGrantlessCredentials(
    clientID string,
    clientSecret string,
    scopes []string,
    endpoint string,
) (*Credentials, error)
```

### 方法

```go
// IsGrantless 判断是否为 grantless credentials
func (c *Credentials) IsGrantless() bool
```

---

## 参考资料

- [官方 SP-API 文档 - Connect to the SP-API](https://developer-docs.amazon.com/sp-api/docs/connecting-to-the-selling-partner-api)
- [官方 SP-API 文档 - Grantless Operations](https://developer-docs.amazon.com/sp-api/docs/grantless-operations)
- [Notifications API](https://developer-docs.amazon.com/sp-api/docs/notifications-api)
- [Application Management API](https://developer-docs.amazon.com/sp-api/docs/application-management-api)

---

**最后更新**: 2025-10-03  
**基于**: 官方 SP-API 文档

