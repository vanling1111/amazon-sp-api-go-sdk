# 代码风格

## 概述

本项目严格遵循 Go 官方代码风格和 Google Go 风格指南，所有注释使用中文。

---

## 基本原则

### 1. 官方规范
- ✅ 遵循 [Effective Go](https://go.dev/doc/effective_go)
- ✅ 遵循 [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments)
- ✅ 遵循 [Google Go Style Guide](https://google.github.io/styleguide/go/)

### 2. 格式化工具
- **必须使用**: `gofmt` 或 `goimports`
- **推荐使用**: `golangci-lint`

```bash
# 格式化代码
gofmt -w .

# 自动整理导入
goimports -w .

# 运行 linter
golangci-lint run
```

---

## 命名规范

### 1. 包名

**规则**:
- 小写单词
- 简短、有意义
- 避免下划线和驼峰
- 与目录名一致

**✅ 好的命名**:
```go
package auth
package transport
package signer
```

**❌ 不好的命名**:
```go
package authenticationService  // 太长、有驼峰
package auth_client            // 有下划线
package utils                  // 太通用
```

---

### 2. 文件名

**规则**:
- 小写字母
- 单词之间用下划线分隔
- 测试文件以 `_test.go` 结尾

**✅ 好的命名**:
```
client.go
http_client.go
credentials_test.go
lwa_signer.go
```

**❌ 不好的命名**:
```
Client.go           // 大写
httpClient.go       // 驼峰
credentials-test.go // 连字符
```

---

### 3. 变量和函数

**规则**:
- 驼峰命名
- 首字母大写表示导出（公开）
- 首字母小写表示未导出（私有）
- 缩写词全部大写或全部小写

**✅ 好的命名**:
```go
// 变量
var maxRetries int
var defaultTimeout time.Duration
var ErrInvalidCredentials = errors.New("invalid credentials")

// 函数
func GetAccessToken(ctx context.Context) (string, error)
func parseHTTPResponse(resp *http.Response) error

// 缩写词
var apiURL string      // 全部大写
var userID string      // 全部大写
var httpClient *http.Client  // 全部小写（未导出）
```

**❌ 不好的命名**:
```go
var MaxRetries int     // 私有变量不应大写
var default_timeout    // 应使用驼峰
var errInvalid         // 错误变量应以 Err 开头
func get_token()       // 应使用驼峰
var ApiUrl string      // 缩写词应全部大写或全部小写
```

---

### 4. 常量

**规则**:
- 驼峰命名
- 相关常量分组
- 使用 `const` 块

**✅ 好的命名**:
```go
const (
    // Grant Types
    GrantTypeRefreshToken      = "refresh_token"
    GrantTypeClientCredentials = "client_credentials"
    
    // Endpoints
    EndpointNA   = "https://api.amazon.com"
    EndpointEU   = "https://api.amazon.co.uk"
    EndpointFE   = "https://api.amazon.co.jp"
    
    // Timeouts
    DefaultTimeout = 30 * time.Second
    MaxRetries     = 3
)
```

---

### 5. 接口

**规则**:
- 单方法接口以 `-er` 结尾
- 多方法接口使用名词

**✅ 好的命名**:
```go
// 单方法接口
type Signer interface {
    Sign(ctx context.Context, req *http.Request) error
}

type TokenProvider interface {
    GetToken(ctx context.Context) (string, error)
}

// 多方法接口
type Client interface {
    Do(ctx context.Context, req *http.Request) (*http.Response, error)
    Use(middleware Middleware)
    Close() error
}
```

**❌ 不好的命名**:
```go
type ISigner interface {}      // 不使用 I 前缀
type SignerInterface interface {} // 不使用 Interface 后缀
```

---

### 6. 结构体

**规则**:
- 驼峰命名
- 使用名词
- 避免 `Data`, `Info`, `Manager` 等无意义后缀

**✅ 好的命名**:
```go
type Credentials struct {
    ClientID     string
    ClientSecret string
    RefreshToken string
}

type Token struct {
    AccessToken string
    ExpiresAt   time.Time
}
```

**❌ 不好的命名**:
```go
type CredentialsData struct {}  // 避免 Data 后缀
type TokenInfo struct {}        // 避免 Info 后缀
type AuthManager struct {}      // 避免 Manager 后缀
```

---

## 注释规范

### 1. 包注释

**位置**: 包名上方

**格式**: Google 风格，中文

```go
// Package auth 提供 Amazon SP-API 的 LWA (Login with Amazon) 认证功能。
//
// 此包实现了访问令牌的获取、缓存和刷新逻辑，
// 支持 refresh_token 和 client_credentials 两种授权模式。
//
// 基本用法:
//
//	creds, err := auth.NewCredentials(clientID, clientSecret, refreshToken, endpoint)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	client := auth.NewClient(creds)
//	token, err := client.GetAccessToken(context.Background())
//
// 官方文档:
//   - https://developer-docs.amazon.com/sp-api/docs/connecting-to-the-selling-partner-api
package auth
```

---

### 2. 函数注释

**格式**: Google 风格，中文

**必须包含**:
- 功能描述
- 参数说明（如果有）
- 返回值说明（如果有）
- 错误说明（如果有）
- 使用示例（推荐）
- 官方文档链接（如果相关）

**✅ 好的注释**:
```go
// GetAccessToken 获取 LWA 访问令牌。
//
// 此方法首先检查缓存，如果缓存中有有效令牌则直接返回。
// 否则，向 LWA 服务请求新的访问令牌。
//
// 参数:
//   - ctx: 请求上下文，用于取消和超时控制
//
// 返回值:
//   - string: 访问令牌
//   - error: 如果请求失败或令牌无效，返回错误
//
// 可能的错误:
//   - ErrInvalidCredentials: 凭证无效
//   - ErrNetworkError: 网络请求失败
//   - context.DeadlineExceeded: 请求超时
//
// 示例:
//
//	token, err := client.GetAccessToken(ctx)
//	if err != nil {
//	    log.Printf("failed to get token: %v", err)
//	    return err
//	}
//	fmt.Println("Access Token:", token)
//
// 官方文档:
//   - https://developer-docs.amazon.com/sp-api/docs/connecting-to-the-selling-partner-api#step-1-request-a-login-with-amazon-access-token
func (c *Client) GetAccessToken(ctx context.Context) (string, error) {
    // 实现...
}
```

**❌ 不好的注释**:
```go
// get token
func (c *Client) GetAccessToken(ctx context.Context) (string, error) {
    // 注释太简单，没有说明参数、返回值、错误
}

// GetAccessToken gets the access token from LWA service
func (c *Client) GetAccessToken(ctx context.Context) (string, error) {
    // 使用了英文注释
}
```

---

### 3. 结构体注释

**格式**:
```go
// Credentials 表示 LWA 认证凭证。
//
// 凭证包含客户端 ID、客户端密钥和刷新令牌，
// 用于向 LWA 服务请求访问令牌。
//
// 支持两种操作模式:
//   - Regular: 使用 RefreshToken（需要卖家授权）
//   - Grantless: 使用 Scopes（无需卖家授权）
//
// 官方文档:
//   - https://developer-docs.amazon.com/sp-api/docs/connecting-to-the-selling-partner-api
type Credentials struct {
    // ClientID 是应用的客户端 ID
    ClientID string

    // ClientSecret 是应用的客户端密钥
    ClientSecret string

    // RefreshToken 是 LWA 刷新令牌（Regular 操作必需）
    RefreshToken string

    // Scopes 是授权范围列表（Grantless 操作必需）
    Scopes []string

    // Endpoint 是 LWA 令牌端点 URL
    Endpoint string
}
```

---

### 4. 常量和变量注释

**格式**:
```go
const (
    // GrantTypeRefreshToken 表示使用 refresh_token 授权模式。
    // 此模式需要卖家授权，用于大多数 SP-API 操作。
    GrantTypeRefreshToken = "refresh_token"

    // GrantTypeClientCredentials 表示使用 client_credentials 授权模式。
    // 此模式无需卖家授权，用于 Grantless 操作。
    GrantTypeClientCredentials = "client_credentials"
)

var (
    // ErrInvalidCredentials 表示提供的凭证无效。
    ErrInvalidCredentials = errors.New("invalid credentials")

    // ErrTokenExpired 表示访问令牌已过期。
    ErrTokenExpired = errors.New("access token expired")
)
```

---

## 代码组织

### 1. 导入顺序

**顺序**:
1. 标准库
2. 第三方库
3. 本项目内部包

**使用 `goimports` 自动整理**

**✅ 好的顺序**:
```go
import (
    "context"
    "encoding/json"
    "fmt"
    "net/http"
    "time"

    "github.com/yourusername/amazon-sp-api-go-sdk/internal/auth"
    "github.com/yourusername/amazon-sp-api-go-sdk/internal/transport"
)
```

---

### 2. 结构体字段顺序

**推荐顺序**:
1. 导出字段（公开）
2. 未导出字段（私有）
3. 嵌入字段
4. 同步原语（`sync.Mutex` 等）放在最后

**✅ 好的顺序**:
```go
type Client struct {
    // 导出字段
    Timeout time.Duration

    // 未导出字段
    credentials *Credentials
    httpClient  *http.Client
    cache       map[string]*Token

    // 同步原语
    mu sync.RWMutex
}
```

---

### 3. 函数顺序

**推荐顺序**:
1. 构造函数 (`New...`)
2. 公开方法（首字母大写）
3. 私有方法（首字母小写）
4. 辅助函数

**✅ 好的顺序**:
```go
// 1. 构造函数
func NewClient(creds *Credentials) *Client {
    // ...
}

// 2. 公开方法
func (c *Client) GetAccessToken(ctx context.Context) (string, error) {
    // ...
}

func (c *Client) RefreshToken(ctx context.Context) (string, error) {
    // ...
}

// 3. 私有方法
func (c *Client) fetchToken(ctx context.Context) (*Token, error) {
    // ...
}

func (c *Client) cacheToken(token *Token) {
    // ...
}

// 4. 辅助函数
func buildTokenRequest(creds *Credentials) url.Values {
    // ...
}
```

---

## 错误处理

### 1. 错误定义

**使用 `errors.New` 或 `fmt.Errorf`**:
```go
var (
    ErrInvalidCredentials = errors.New("invalid credentials")
    ErrTokenExpired       = errors.New("access token expired")
    ErrNetworkError       = errors.New("network error")
)
```

**自定义错误类型**:
```go
// APIError 表示 SP-API 返回的错误。
type APIError struct {
    Code    string
    Message string
    Details map[string]interface{}
}

func (e *APIError) Error() string {
    return fmt.Sprintf("API error %s: %s", e.Code, e.Message)
}
```

---

### 2. 错误包装

**使用 `fmt.Errorf` 和 `%w`**:
```go
func (c *Client) fetchToken(ctx context.Context) (*Token, error) {
    resp, err := c.httpClient.Do(req)
    if err != nil {
        return nil, fmt.Errorf("send LWA request: %w", err)
    }
    
    // ...
    
    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("LWA request failed with status %d: %w", resp.StatusCode, ErrAuthFailed)
    }
    
    return token, nil
}
```

---

### 3. 错误检查

**使用 `errors.Is` 和 `errors.As`**:
```go
import "errors"

token, err := client.GetAccessToken(ctx)
if err != nil {
    // 检查特定错误类型
    if errors.Is(err, auth.ErrInvalidCredentials) {
        log.Println("凭证无效，请检查配置")
        return
    }
    
    // 提取错误详情
    var apiErr *auth.APIError
    if errors.As(err, &apiErr) {
        log.Printf("API 错误: %s - %s", apiErr.Code, apiErr.Message)
        return
    }
    
    // 其他错误
    log.Printf("未知错误: %v", err)
    return
}
```

---

## 测试规范

### 1. 测试文件

**规则**:
- 与源文件同一目录
- 文件名以 `_test.go` 结尾

```
auth/
  client.go
  client_test.go
  credentials.go
  credentials_test.go
```

---

### 2. 测试函数命名

**格式**: `Test<FunctionName>_<Scenario>`

**✅ 好的命名**:
```go
func TestNewCredentials_Success(t *testing.T) {}
func TestNewCredentials_MissingClientID(t *testing.T) {}
func TestGetAccessToken_CacheHit(t *testing.T) {}
func TestGetAccessToken_CacheMiss(t *testing.T) {}
```

**❌ 不好的命名**:
```go
func TestNewCredentials(t *testing.T) {}  // 太笼统
func TestCase1(t *testing.T) {}           // 无意义
```

---

### 3. 表驱动测试

**推荐使用表驱动测试**:
```go
func TestNewCredentials(t *testing.T) {
    tests := []struct {
        name      string
        clientID  string
        secret    string
        token     string
        endpoint  string
        wantErr   bool
        errType   error
    }{
        {
            name:     "success",
            clientID: "test-client-id",
            secret:   "test-secret",
            token:    "test-token",
            endpoint: "https://api.amazon.com/auth/o2/token",
            wantErr:  false,
        },
        {
            name:     "missing client ID",
            clientID: "",
            secret:   "test-secret",
            token:    "test-token",
            endpoint: "https://api.amazon.com/auth/o2/token",
            wantErr:  true,
            errType:  ErrInvalidCredentials,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            creds, err := NewCredentials(tt.clientID, tt.secret, tt.token, tt.endpoint)
            
            if (err != nil) != tt.wantErr {
                t.Errorf("NewCredentials() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            
            if tt.wantErr && !errors.Is(err, tt.errType) {
                t.Errorf("NewCredentials() error type = %T, want %T", err, tt.errType)
                return
            }
            
            if !tt.wantErr && creds == nil {
                t.Error("NewCredentials() returned nil credentials")
            }
        })
    }
}
```

---

## Linter 配置

### `.golangci.yml`

```yaml
linters:
  enable:
    - gofmt         # 代码格式化
    - goimports     # 导入整理
    - govet         # Go vet
    - errcheck      # 错误检查
    - staticcheck   # 静态分析
    - unused        # 未使用代码
    - gosimple      # 代码简化
    - ineffassign   # 无效赋值
    - misspell      # 拼写检查
    - gocritic      # Go 代码评审
    - revive        # 替代 golint

linters-settings:
  gofmt:
    simplify: true
  
  misspell:
    locale: US
  
  revive:
    rules:
      - name: exported
        arguments:
          - "disableStutteringCheck"

issues:
  exclude-use-default: false
  max-issues-per-linter: 0
  max-same-issues: 0

run:
  timeout: 5m
  skip-dirs:
    - vendor
    - testdata
```

**运行 linter**:
```bash
golangci-lint run
```

---

## 参考资料

- [Effective Go](https://go.dev/doc/effective_go)
- [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments)
- [Google Go Style Guide](https://google.github.io/styleguide/go/)
- [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)

