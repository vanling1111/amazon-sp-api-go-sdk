# 测试目录

本目录包含集成测试和性能测试。单元测试与源代码放在同一目录下（`*_test.go`）。

## 目录结构

```
tests/
├── integration/        # 集成测试
│   ├── orders_test.go
│   ├── reports_test.go
│   └── ...
└── benchmarks/         # 性能测试
    ├── client_bench_test.go
    └── ...
```

## 单元测试

单元测试与源代码放在同一目录：

```
internal/auth/
├── client.go
├── client_test.go      # 单元测试
├── credentials.go
└── credentials_test.go # 单元测试
```

**运行单元测试**:
```bash
# 运行所有单元测试
go test ./...

# 运行特定包的测试
go test ./internal/auth/...

# 查看详细输出
go test -v ./...

# 查看覆盖率
go test -cover ./...
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## 集成测试

集成测试使用真实的 SP-API Sandbox 环境。

### 配置

创建 `tests/integration/.env` 文件：

```bash
SP_API_CLIENT_ID=your-client-id
SP_API_CLIENT_SECRET=your-client-secret
SP_API_REFRESH_TOKEN=your-refresh-token
SP_API_REGION=na
SP_API_MARKETPLACE=ATVPDKIKX0DER
SP_API_SANDBOX=true  # 使用 Sandbox 环境
```

### 运行集成测试

```bash
# 运行所有集成测试
go test ./tests/integration/... -v

# 运行特定测试
go test ./tests/integration/... -run TestOrders -v

# 使用构建标签
go test -tags=integration ./tests/integration/... -v
```

### 编写集成测试

```go
// tests/integration/orders_test.go
// +build integration

package integration

import (
    "context"
    "testing"
    "time"

    "github.com/yourusername/amazon-sp-api-go-sdk/pkg/spapi"
)

func TestOrders_GetOrders(t *testing.T) {
    client := setupTestClient(t)
    defer client.Close()

    ctx := context.Background()
    resp, err := client.Orders.GetOrders(ctx, &spapi.GetOrdersRequest{
        MarketplaceIDs: []string{spapi.MarketplaceUS},
        CreatedAfter:   time.Now().Add(-24 * time.Hour),
    })

    if err != nil {
        t.Fatalf("GetOrders failed: %v", err)
    }

    if resp == nil {
        t.Fatal("Expected non-nil response")
    }

    t.Logf("Retrieved %d orders", len(resp.Orders))
}
```

## 性能测试

### 运行性能测试

```bash
# 运行所有性能测试
go test -bench=. ./tests/benchmarks/...

# 运行特定性能测试
go test -bench=BenchmarkClient ./tests/benchmarks/...

# 包含内存分配统计
go test -bench=. -benchmem ./tests/benchmarks/...

# 运行多次取平均值
go test -bench=. -benchtime=10s ./tests/benchmarks/...
```

### 编写性能测试

```go
// tests/benchmarks/client_bench_test.go
package benchmarks

import (
    "context"
    "testing"

    "github.com/yourusername/amazon-sp-api-go-sdk/pkg/spapi"
)

func BenchmarkClient_GetOrders(b *testing.B) {
    client := setupBenchClient(b)
    defer client.Close()

    ctx := context.Background()
    req := &spapi.GetOrdersRequest{
        MarketplaceIDs: []string{spapi.MarketplaceUS},
    }

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, err := client.Orders.GetOrders(ctx, req)
        if err != nil {
            b.Fatal(err)
        }
    }
}

func BenchmarkTokenCache(b *testing.B) {
    // 测试令牌缓存性能
}
```

## 测试覆盖率

### 生成覆盖率报告

```bash
# 生成覆盖率文件
go test -coverprofile=coverage.out ./...

# 查看总体覆盖率
go tool cover -func=coverage.out

# 生成 HTML 报告
go tool cover -html=coverage.out -o coverage.html

# 按包查看覆盖率
go test -coverprofile=coverage.out ./... && \
go tool cover -func=coverage.out | grep -E 'internal/(auth|transport|signer)'
```

### 覆盖率要求

- **新代码**: ≥ 90%
- **现有代码**: 不降低整体覆盖率
- **公开 API**: 100%

**当前覆盖率**:
- `internal/auth`: 89.0%
- `internal/transport`: 87.4%
- `internal/signer`: 93.3%
- **整体**: 90.2%

## Mock 测试

使用接口和 Mock 进行单元测试：

```go
// 定义接口
type TokenProvider interface {
    GetAccessToken(ctx context.Context) (string, error)
}

// Mock 实现
type mockTokenProvider struct {
    token string
    err   error
}

func (m *mockTokenProvider) GetAccessToken(ctx context.Context) (string, error) {
    return m.token, m.err
}

// 测试
func TestSigner_WithMock(t *testing.T) {
    mockProvider := &mockTokenProvider{
        token: "test-token",
        err:   nil,
    }
    
    signer := NewLWASigner(mockProvider)
    // 测试逻辑...
}
```

## 测试最佳实践

1. **表驱动测试** - 使用测试表覆盖多个场景
2. **子测试** - 使用 `t.Run()` 组织相关测试
3. **并发测试** - 使用 `t.Parallel()` 提高效率
4. **清理资源** - 使用 `defer` 或 `t.Cleanup()` 清理
5. **隔离测试** - 不依赖外部状态或顺序

## CI/CD 集成

在 `.github/workflows/ci.yml` 中配置：

```yaml
- name: Run Unit Tests
  run: go test -v -race -cover ./...

- name: Run Integration Tests
  run: go test -v -tags=integration ./tests/integration/...
  env:
    SP_API_CLIENT_ID: ${{ secrets.SP_API_CLIENT_ID }}
    SP_API_CLIENT_SECRET: ${{ secrets.SP_API_CLIENT_SECRET }}

- name: Upload Coverage
  uses: codecov/codecov-action@v3
  with:
    file: ./coverage.out
```

## 官方文档

- [Go Testing Package](https://pkg.go.dev/testing)
- [SP-API Sandbox](https://developer-docs.amazon.com/sp-api/docs/the-selling-partner-api-sandbox)

