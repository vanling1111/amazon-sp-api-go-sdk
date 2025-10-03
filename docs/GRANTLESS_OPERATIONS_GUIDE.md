# Grantless 操作指南

## 什么是 Grantless 操作

Grantless 操作是不需要特定卖家授权的 API 操作，使用应用的 Client ID 和 Client Secret 即可调用。

## 支持的 Grantless 操作

### 1. Notifications API
- 创建和管理通知订阅
- 不需要卖家的 refresh token

### 2. Application Management API  
- 管理应用配置
- 应用级别的操作

## 使用方法

### 创建 Grantless 客户端

```go
package main

import (
    "github.com/vanling1111/amazon-sp-api-go-sdk/internal/models"
    "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
    notifications "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi/notifications-v1"
)

func main() {
    // 创建 Grantless 客户端
    client, err := spapi.NewClient(
        spapi.WithRegion(models.RegionNA),
        spapi.WithGrantlessCredentials(
            "your-client-id",
            "your-client-secret",
            []string{"sellingpartnerapi::notifications"},
        ),
    )
    if err != nil {
        panic(err)
    }
    defer client.Close()
    
    // 使用 Notifications API
    notifClient := notifications.NewClient(client)
    
    // 创建通知订阅
    // ...
}
```

## Scopes 说明

### 可用的 Scopes

```
sellingpartnerapi::notifications  # Notifications API
sellingpartnerapi::migration      # 数据迁移（少用）
```

## Regular vs Grantless

### Regular 操作
- 需要卖家授权
- 使用 refresh token
- 访问卖家特定数据

### Grantless 操作
- 不需要卖家授权
- 使用 client credentials
- 应用级别操作

## 示例代码

完整示例请查看：
- `examples/grantless/main.go`

## 参考文档

- [Grantless Operations](https://developer-docs.amazon.com/sp-api/docs/grantless-operations)
- [Notifications API](https://developer-docs.amazon.com/sp-api/docs/notifications-api)

---

更新时间：2025-10-03

