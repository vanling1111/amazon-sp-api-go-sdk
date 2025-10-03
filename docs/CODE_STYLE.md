# 代码风格指南

## 基本规范

遵循 Go 官方代码规范和社区最佳实践。

## 命名规范

### 包名
- 小写单词
- 简短清晰
- 避免下划线

```go
package auth
package transport
package ratelimit
```

### 函数名
- 驼峰命名
- 公开函数首字母大写
- 私有函数首字母小写

```go
// 公开函数
func NewClient() *Client
func GetOrders() error

// 私有函数
func buildRequest() error
func parseResponse() error
```

### 变量名
- 驼峰命名
- 见名知意

```go
var maxRetries int
var defaultTimeout time.Duration
var httpClient *http.Client
```

## 注释规范

### 包注释

```go
// Package auth 提供 Amazon SP-API 的 LWA 认证功能。
//
// 支持 Regular 和 Grantless 两种操作模式。
package auth
```

### 函数注释

```go
// GetAccessToken 获取 LWA 访问令牌。
//
// 参数：
//   - ctx: 请求上下文
//
// 返回：
//   - string: 访问令牌
//   - error: 错误信息
func GetAccessToken(ctx context.Context) (string, error)
```

## 错误处理

### 定义错误

```go
var (
    ErrInvalidCredentials = errors.New("invalid credentials")
    ErrTokenExpired = errors.New("token expired")
)
```

### 包装错误

```go
if err != nil {
    return nil, fmt.Errorf("fetch token: %w", err)
}
```

## 代码组织

### 导入顺序

1. 标准库
2. 第三方库（如有）
3. 项目内部包

```go
import (
    "context"
    "fmt"
    "net/http"
    
    "github.com/vanling1111/amazon-sp-api-go-sdk/internal/auth"
)
```

## 测试

### 测试文件命名

```go
client.go       // 实现
client_test.go  // 测试
```

### 测试函数命名

```go
func TestNewClient(t *testing.T)
func TestGetOrders_Success(t *testing.T)
func TestGetOrders_InvalidParams(t *testing.T)
```

## 工具

```bash
# 格式化代码
gofmt -w .

# 导入整理
goimports -w .

# 代码检查
go vet ./...

# Linter（可选）
golangci-lint run
```

## 参考

- [Effective Go](https://go.dev/doc/effective_go)
- [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments)

