# 开发指南

## 🎯 开发原则

本项目遵循以下开发原则以确保高质量和可维护性：

### 📚 基于官方规范

- **OpenAPI 驱动**: 所有 API 模型和接口定义直接从 Amazon 官方 OpenAPI 规范自动生成
- **文档参考**: https://developer-docs.amazon.com/sp-api/docs/
- **规范仓库**: https://github.com/amzn/selling-partner-api-models

### 🐹 Go 最佳实践

- **惯用法**: 遵循 Go 社区公认的最佳实践和代码风格
- **标准库优先**: 尽可能使用标准库，减少外部依赖
- **接口设计**: 使用接口实现模块解耦
- **Context 优先**: 所有 API 方法接受 `context.Context`

### 🧪 质量保证

- **测试驱动**: 核心模块测试覆盖率 > 90%
- **代码审查**: 所有 PR 需要通过代码审查
- **持续集成**: GitHub Actions 自动运行测试
- **文档完整**: 每个公开函数都有完整的文档注释

### 📚 **推荐参考资源**

**官方资源**:
- [SP-API 官方文档](https://developer-docs.amazon.com/sp-api/docs/)
- [SP-API 参考手册](https://developer-docs.amazon.com/sp-api/reference/)
- [OpenAPI 规范仓库](https://github.com/amzn/selling-partner-api-models)

**Go 语言资源**:
- [Go 官方文档](https://go.dev/doc/)
- [Go 标准库](https://pkg.go.dev/std)
- [Effective Go](https://go.dev/doc/effective_go)

**相关标准**:
- [OAuth 2.0 RFC](https://oauth.net/2/)
- [OpenAPI Specification](https://swagger.io/specification/)

### 📝 **代码审查检查清单**

提交 PR 前请确认：
- [ ] 代码遵循 Go 最佳实践和项目代码风格
- [ ] 添加了适当的测试用例
- [ ] 所有测试通过
- [ ] 添加了必要的文档注释
- [ ] 更新了相关文档（如有需要）

---

## 1. 开发环境要求

### 1.1 基础环境
- **Go 版本**: 1.21 或更高
- **Git**: 最新版本
- **IDE**: VSCode / GoLand（推荐）
- **操作系统**: Windows / macOS / Linux

### 1.2 开发工具
```bash
# 安装 golangci-lint（代码检查）
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# 安装 mockgen（生成 mock）
go install github.com/golang/mock/mockgen@latest

# 安装 goimports（格式化导入）
go install golang.org/x/tools/cmd/goimports@latest
```

## 2. 开发流程

### 2.1 准备工作

#### 克隆仓库
```bash
git clone https://github.com/your-repo/amazon-sp-api-go-sdk.git
cd amazon-sp-api-go-sdk
```

#### 安装依赖
```bash
go mod download
go mod verify
```

### 2.2 开发新功能

#### 步骤 1: 阅读官方文档
在开发任何功能之前，**必须**先阅读官方文档：

1. **访问官方 SP-API 文档**
   - URL: https://developer-docs.amazon.com/sp-api/
   - 找到对应的 API 分类（如 Orders API, Reports API）

2. **理解 API 规范**
   - 端点路径和 HTTP 方法
   - 请求参数（查询参数、路径参数、请求体）
   - 响应格式和字段定义
   - 错误代码和含义
   - 速率限制规则

3. **检查 OpenAPI 规范**
   - 如果有 OpenAPI/Swagger 规范，优先参考
   - 位置: https://github.com/amzn/selling-partner-api-models

#### 步骤 2: 设计 Go 接口

根据官方文档设计 Go 类型和接口：

```go
// 1. 定义请求类型
type GetOrdersRequest struct {
    MarketplaceIDs    []string   `json:"marketplace_ids"`
    CreatedAfter      *time.Time `json:"created_after,omitempty"`
    CreatedBefore     *time.Time `json:"created_before,omitempty"`
    OrderStatuses     []string   `json:"order_statuses,omitempty"`
    MaxResultsPerPage *int       `json:"max_results_per_page,omitempty"`
    NextToken         *string    `json:"next_token,omitempty"`
}

// 2. 定义响应类型
type GetOrdersResponse struct {
    Payload struct {
        Orders     []Order `json:"orders"`
        NextToken  *string `json:"next_token,omitempty"`
    } `json:"payload"`
    Errors []Error `json:"errors,omitempty"`
}

// 3. 定义接口
type OrdersAPI interface {
    GetOrders(ctx context.Context, req *GetOrdersRequest) (*GetOrdersResponse, error)
    GetOrder(ctx context.Context, orderID string) (*GetOrderResponse, error)
}
```

#### 步骤 3: 实现功能

```go
// client.go
package orders

import (
    "context"
    "fmt"
    "net/http"
    
    "github.com/amazon-sp-api-go-sdk/internal/transport"
)

type Client struct {
    transport *transport.Client
}

func NewClient(transport *transport.Client) *Client {
    return &Client{transport: transport}
}

func (c *Client) GetOrders(ctx context.Context, req *GetOrdersRequest) (*GetOrdersResponse, error) {
    // 1. 验证请求
    if err := req.Validate(); err != nil {
        return nil, fmt.Errorf("invalid request: %w", err)
    }
    
    // 2. 构建 HTTP 请求
    httpReq, err := c.buildGetOrdersRequest(ctx, req)
    if err != nil {
        return nil, err
    }
    
    // 3. 发送请求
    httpResp, err := c.transport.Do(ctx, httpReq)
    if err != nil {
        return nil, err
    }
    defer httpResp.Body.Close()
    
    // 4. 解析响应
    var resp GetOrdersResponse
    if err := c.transport.DecodeResponse(httpResp, &resp); err != nil {
        return nil, err
    }
    
    return &resp, nil
}
```

#### 步骤 4: 编写测试

```go
// client_test.go
package orders_test

import (
    "context"
    "testing"
    "time"
    
    "github.com/amazon-sp-api-go-sdk/spapi/orders"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestClient_GetOrders(t *testing.T) {
    // 准备
    ctx := context.Background()
    client := setupTestClient(t)
    
    req := &orders.GetOrdersRequest{
        MarketplaceIDs: []string{"ATVPDKIKX0DER"},
        CreatedAfter:   timePtr(time.Now().Add(-24 * time.Hour)),
    }
    
    // 执行
    resp, err := client.GetOrders(ctx, req)
    
    // 断言
    require.NoError(t, err)
    assert.NotNil(t, resp)
    assert.NotEmpty(t, resp.Payload.Orders)
}
```

#### 步骤 5: 编写文档

```go
// GetOrders 获取订单列表。
//
// 此方法从 Amazon SP-API 检索订单，支持多种筛选条件。
//
// 参数:
//   - ctx: 请求上下文，用于超时和取消控制
//   - req: 订单查询请求，必须包含至少一个 MarketplaceID
//
// 返回值:
//   - *GetOrdersResponse: 包含订单列表的响应
//   - error: 如果请求失败，返回错误
//
// 错误:
//   - ErrInvalidRequest: 请求参数无效
//   - ErrUnauthorized: 认证失败
//   - ErrRateLimit: 超过速率限制
//   - ErrServerError: 服务器内部错误
//
// 示例:
//   req := &GetOrdersRequest{
//       MarketplaceIDs: []string{"ATVPDKIKX0DER"},
//       CreatedAfter:   timePtr(time.Now().Add(-7 * 24 * time.Hour)),
//   }
//   resp, err := client.GetOrders(ctx, req)
//   if err != nil {
//       log.Fatal(err)
//   }
//   fmt.Printf("Found %d orders\n", len(resp.Payload.Orders))
//
// 参考:
//   - https://developer-docs.amazon.com/sp-api/docs/orders-api-v0-reference#getorders
func (c *Client) GetOrders(ctx context.Context, req *GetOrdersRequest) (*GetOrdersResponse, error) {
    // ...
}
```

### 2.3 代码审查

#### 自查清单

在提交代码前，请确认：

- [ ] **阅读了官方文档** - 确保理解 API 的行为
- [ ] **类型定义准确** - 与官方文档的字段定义一致
- [ ] **错误处理完整** - 处理了所有可能的错误情况
- [ ] **添加了单元测试** - 覆盖主要流程和边界情况
- [ ] **编写了完整文档** - 包含参数说明、示例和错误说明
- [ ] **代码通过 lint** - 运行 `golangci-lint run`
- [ ] **代码格式化** - 运行 `gofmt -s -w .` 和 `goimports -w .`
- [ ] **测试通过** - 运行 `go test ./...`

#### 代码检查命令

```bash
# 格式化代码
gofmt -s -w .
goimports -w .

# 运行 linter
golangci-lint run

# 运行测试
go test ./... -v -race -coverprofile=coverage.out

# 查看覆盖率
go tool cover -html=coverage.out
```

## 3. 测试规范

### 3.1 测试类型

#### 单元测试
```go
// 测试单个函数的行为
func TestValidateRequest(t *testing.T) {
    tests := []struct {
        name    string
        req     *GetOrdersRequest
        wantErr bool
    }{
        {
            name: "valid request",
            req:  &GetOrdersRequest{MarketplaceIDs: []string{"ATVPDKIKX0DER"}},
            wantErr: false,
        },
        {
            name: "empty marketplace IDs",
            req:  &GetOrdersRequest{},
            wantErr: true,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := tt.req.Validate()
            if tt.wantErr {
                assert.Error(t, err)
            } else {
                assert.NoError(t, err)
            }
        })
    }
}
```

#### 集成测试
```go
// 测试多个组件的集成
func TestClient_Integration(t *testing.T) {
    if testing.Short() {
        t.Skip("skipping integration test")
    }
    
    // 使用真实的配置
    client := setupRealClient(t)
    
    // 测试完整流程
    ctx := context.Background()
    resp, err := client.GetOrders(ctx, &GetOrdersRequest{
        MarketplaceIDs: []string{"ATVPDKIKX0DER"},
    })
    
    require.NoError(t, err)
    assert.NotNil(t, resp)
}
```

#### Mock 测试
```go
// 使用 mock 测试
func TestClient_WithMock(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()
    
    mockTransport := mock.NewMockTransport(ctrl)
    client := NewClient(mockTransport)
    
    // 设置预期
    mockTransport.EXPECT().
        Do(gomock.Any(), gomock.Any()).
        Return(&http.Response{StatusCode: 200, Body: ...}, nil)
    
    // 执行测试
    resp, err := client.GetOrders(context.Background(), &GetOrdersRequest{...})
    
    require.NoError(t, err)
    assert.NotNil(t, resp)
}
```

### 3.2 测试覆盖率要求

- **核心代码**: > 80%
- **关键路径**: 100%
- **错误处理**: 100%

### 3.3 基准测试

```go
func BenchmarkClient_GetOrders(b *testing.B) {
    client := setupTestClient(b)
    ctx := context.Background()
    req := &GetOrdersRequest{MarketplaceIDs: []string{"ATVPDKIKX0DER"}}
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, err := client.GetOrders(ctx, req)
        if err != nil {
            b.Fatal(err)
        }
    }
}
```

## 4. 文档规范

### 4.1 代码注释

#### 包注释
```go
// Package orders 提供了 Amazon SP-API Orders API 的 Go 客户端实现。
//
// Orders API 允许卖家检索订单信息、更新订单状态和管理订单项。
//
// 使用示例:
//   client, err := spapi.NewClient(...)
//   orders := client.Orders
//   resp, err := orders.GetOrders(ctx, &GetOrdersRequest{...})
//
// 官方文档:
//   https://developer-docs.amazon.com/sp-api/docs/orders-api-v0-reference
package orders
```

#### 函数注释
```go
// GetOrders 获取订单列表。
//
// 参数:
//   - ctx: 请求上下文
//   - req: 订单查询请求
//
// 返回值:
//   - *GetOrdersResponse: 订单列表响应
//   - error: 错误信息
func (c *Client) GetOrders(ctx context.Context, req *GetOrdersRequest) (*GetOrdersResponse, error)
```

#### 类型注释
```go
// GetOrdersRequest 表示获取订单的请求参数。
type GetOrdersRequest struct {
    // MarketplaceIDs 是要查询的市场 ID 列表（必填）。
    MarketplaceIDs []string `json:"marketplace_ids"`
    
    // CreatedAfter 是订单创建的起始时间（可选）。
    CreatedAfter *time.Time `json:"created_after,omitempty"`
}
```

### 4.2 示例代码

每个主要功能都应提供示例：

```go
func ExampleClient_GetOrders() {
    // 创建客户端
    client, _ := spapi.NewClient(
        spapi.WithCredentials(clientID, clientSecret, refreshToken),
        spapi.WithRegion(spapi.RegionNA),
    )
    
    // 获取最近 7 天的订单
    resp, err := client.Orders.GetOrders(context.Background(), &GetOrdersRequest{
        MarketplaceIDs: []string{spapi.MarketplaceUS},
        CreatedAfter:   timePtr(time.Now().Add(-7 * 24 * time.Hour)),
    })
    
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Found %d orders\n", len(resp.Payload.Orders))
}
```

## 5. 版本管理

### 5.1 语义版本

遵循 [Semantic Versioning 2.0.0](https://semver.org/):

- **MAJOR**: 不兼容的 API 变更
- **MINOR**: 向后兼容的功能新增
- **PATCH**: 向后兼容的问题修复

### 5.2 变更日志

每次发布都要更新 `CHANGELOG.md`:

```markdown
## [1.2.0] - 2025-01-15

### Added
- 新增 Catalog Items API 支持
- 新增自动重试功能

### Changed
- 改进了速率限制器的性能

### Fixed
- 修复了 LWA 令牌缓存的竞态条件

### Deprecated
- `OldMethod` 已弃用，请使用 `NewMethod`
```

## 6. 提交规范

### 6.1 Commit 消息格式

```
<type>(<scope>): <subject>

<body>

<footer>
```

**类型 (type)**:
- `feat`: 新功能
- `fix`: 错误修复
- `docs`: 文档更新
- `style`: 代码格式（不影响功能）
- `refactor`: 重构（不是新功能，也不是修复）
- `perf`: 性能优化
- `test`: 测试相关
- `chore`: 构建过程或辅助工具的变动

**示例**:
```
feat(orders): add GetOrders API support

Implement GetOrders method for Orders API v0 according to
official SP-API documentation.

- Add GetOrdersRequest and GetOrdersResponse types
- Add comprehensive unit tests
- Add integration tests
- Add documentation and examples

Closes #123
```

## 7. 发布流程

### 7.1 发布检查清单

- [ ] 所有测试通过
- [ ] 文档已更新
- [ ] CHANGELOG.md 已更新
- [ ] 版本号已更新
- [ ] 创建 Git tag

### 7.2 发布命令

```bash
# 1. 更新版本号
# 编辑 version.go

# 2. 提交变更
git add .
git commit -m "chore: release v1.2.0"

# 3. 创建 tag
git tag -a v1.2.0 -m "Release v1.2.0"

# 4. 推送
git push origin main
git push origin v1.2.0
```

## 8. 持续集成

### 8.1 GitHub Actions

```yaml
# .github/workflows/ci.yml
name: CI

on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v4
        with:
          go-version: '1.21'
      
      - name: Run tests
        run: go test ./... -v -race -coverprofile=coverage.out
      
      - name: Upload coverage
        uses: codecov/codecov-action@v3
```

## 9. 常见问题

### Q: 如何知道官方 API 更新了？
A: 查看 [API 追踪文档](API_TRACKING.md)，我们有自动化工具监控。

### Q: 如何添加新的 API？
A: 参考本文档的"开发新功能"部分。

### Q: 如何调试？
A: 启用 Debug 模式：`spapi.WithDebug(true)`

---

**文档版本**: v1.0  
**最后更新**: 2025-01-02

