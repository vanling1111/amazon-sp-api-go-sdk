# 项目结构

## 📌 重要说明

### 自动生成�?API 类型文件

本项目的 API 类型定义（`model_*.go`）是使用官方推荐�?`swagger-codegen` 工具�?OpenAPI 规范自动生成的�?

**关键特点**�?
- 📂 **每个 API �?~70 �?`model_*.go` 文件**
- �?与官�?Python SDK�?0个）、PHP SDK�?0个）完全一�?
- �?符合 Go 语言"一个类型一个文�?的最佳实�?
- �?清晰、可维护、易于协�?

**为什么这么多文件�?*
- Go、Python、PHP 的官�?SDK 都采用这种结�?
- Java SDK �?81 个文件（因为 Java 为数组类型创建额外的 List 类）
- 社区 Go SDK（renabled）有 83 个文件（过度设计�?
- **我们�?70 个文件是最优的** �?

**不要惊慌�?* 这是正常且推荐的结构�?

---

## 目录组织

```
amazon-sp-api-go-sdk/
├── .github/                    # GitHub 配置
�?  └── workflows/             # GitHub Actions 工作�?
�?      ├── ci.yml            # 持续集成
�?      ├── release.yml       # 发布管理
�?      └── doc-check.yml     # 文档更新检�?
�?
├── api/                       # API 模型定义（自动生成）
�?  ├── orders/               # Orders API 模型
�?  ├── reports/              # Reports API 模型
�?  ├── feeds/                # Feeds API 模型
�?  ├── listings/             # Listings API 模型
�?  └── ...                   # 其他 API 模型
�?
├── cmd/                      # 命令行工�?
�?  └── generator/           # 代码生成�?
�?
├── docs/                     # 项目文档
�?  ├── ARCHITECTURE.md      # 架构设计
�?  ├── DEVELOPMENT.md       # 开发规�?
�?  ├── PROJECT_STRUCTURE.md # 本文�?
�?  ├── API_TRACKING.md      # API 追踪策略
�?  ├── CODE_STYLE.md        # 代码风格
�?  ├── CONTRIBUTING.md      # 贡献指南
�?  ├── VERSION_TRACKING.md  # 版本追踪
�?  └── GRANTLESS_OPERATIONS_GUIDE.md # Grantless 操作指南
�?
├── examples/                 # 示例代码
�?  ├── basic_usage/         # 基础用法
�?  ├── advanced_usage/      # 高级用法
�?  ├── orders/              # Orders API 示例
�?  ├── reports/             # Reports API 示例
�?  └── README.md            # 示例说明
�?
├── internal/                 # 内部包（不对外暴露）
�?  ├── auth/                # LWA 认证
�?  �?  ├── client.go        # LWA 客户�?
�?  �?  ├── client_test.go   # 单元测试
�?  �?  ├── credentials.go   # 凭证管理
�?  �?  ├── credentials_test.go
�?  �?  ├── token.go         # 令牌结构
�?  �?  ├── token_test.go
�?  �?  └── errors.go        # 错误定义
�?  �?
�?  ├── transport/           # HTTP 传输�?
�?  �?  ├── client.go        # HTTP 客户�?
�?  �?  ├── client_test.go
�?  �?  ├── middleware.go    # 中间�?
�?  �?  ├── middleware_test.go
�?  �?  ├── retry.go         # 重试逻辑
�?  �?  └── retry_test.go
�?  �?
�?  ├── signer/              # 请求签名
�?  �?  ├── lwa.go           # LWA 签名�?
�?  �?  ├── lwa_test.go
�?  �?  ├── rdt.go           # RDT 签名�?
�?  �?  ├── rdt_test.go
�?  �?  ├── chain.go         # 签名器链
�?  �?  └── signer.go        # 签名器接�?
�?  �?
�?  ├── ratelimit/           # 速率限制
�?  �?  ├── limiter.go       # 限流�?
�?  �?  ├── limiter_test.go
�?  �?  ├── bucket.go        # Token Bucket
�?  �?  └── bucket_test.go
�?  �?
�?  ├── codec/               # 编解�?
�?  �?  ├── json.go          # JSON 编解�?
�?  �?  ├── json_test.go
�?  �?  └── validator.go     # 数据验证
�?  �?
�?  ├── models/              # 内部模型
�?  �?  └── common.go        # 通用模型
�?  �?
�?  └── utils/               # 工具函数
�?      ├── http.go          # HTTP 工具
�?      ├── time.go          # 时间工具
�?      └── string.go        # 字符串工�?
�?
├── pkg/                     # 公开包（对外暴露�?
�?  └── spapi/              # SP-API 客户�?
�?      ├── client.go        # 主客户端
�?      ├── client_test.go   # 主客户端测试
�?      ├── config.go        # 配置选项
�?      ├── config_test.go   # 配置测试
�?      ├── errors.go        # 公开错误类型
�?      ├── regions.go       # 区域定义
�?      ├── marketplaces.go  # 市场定义
�?      �?
�?      ├── orders/          # Orders API
�?      �?  ├── client.go    # Orders API 客户端（手写�?
�?      �?  ├── client_test.go # 单元测试
�?      �?  ├── examples_test.go # 示例测试
�?      �?  └── model_*.go   # 自动生成的类型定义（~70个文件）
�?      �?
�?      ├── reports/         # Reports API
�?      �?  ├── client.go
�?      �?  ├── client_test.go
�?      �?  └── model_*.go   # 自动生成的类型定�?
�?      �?
�?      ├── feeds/           # Feeds API
�?      �?  ├── client.go
�?      �?  ├── client_test.go
�?      �?  └── model_*.go   # 自动生成的类型定�?
�?      �?
�?      ├── catalog-items/   # Catalog Items API
�?      �?  ├── client.go
�?      �?  ├── client_test.go
�?      �?  └── model_*.go   # 自动生成的类型定�?
�?      �?
�?      ├── listings-items/  # Listings Items API
�?      �?  ├── client.go
�?      �?  ├── client_test.go
�?      �?  └── model_*.go   # 自动生成的类型定�?
�?      �?
�?      ├── notifications/   # Notifications API
�?      �?  ├── client.go
�?      �?  ├── client_test.go
�?      �?  └── model_*.go   # 自动生成的类型定�?
�?      �?
�?      ├── pricing/         # Product Pricing API
�?      �?  ├── client.go
�?      �?  ├── client_test.go
�?      �?  └── model_*.go   # 自动生成的类型定�?
�?      �?
�?      ├── fba-inventory/   # FBA Inventory API
�?      �?  ├── client.go
�?      �?  ├── client_test.go
�?      �?  └── model_*.go   # 自动生成的类型定�?
�?      �?
�?      ├── fulfillment-inbound/ # FBA Inbound API
�?      �?  ├── client.go
�?      �?  ├── client_test.go
�?      �?  └── model_*.go   # 自动生成的类型定�?
�?      �?
�?      ├── fulfillment-outbound/ # FBA Outbound API
�?      �?  ├── client.go
�?      �?  ├── client_test.go
�?      �?  └── model_*.go   # 自动生成的类型定�?
�?      �?
�?      ├── sellers/         # Sellers API
�?      �?  ├── client.go
�?      �?  ├── client_test.go
�?      �?  └── model_*.go   # 自动生成的类型定�?
�?      �?
�?      ├── tokens/          # Tokens API (RDT)
�?      �?  ├── client.go
�?      �?  ├── client_test.go
�?      �?  └── model_*.go   # 自动生成的类型定�?
�?      �?
�?      └── ...              # 其他47个SP-API（总计47个API�?
�?
├── tests/                   # 集成测试
�?  ├── integration/         # 集成测试
�?  �?  ├── orders_test.go
�?  �?  ├── reports_test.go
�?  �?  └── ...
�?  �?
�?  └── benchmarks/          # 性能测试
�?      └── client_bench_test.go
�?
├── tools/                   # 开发工�?
�?  ├── monitoring/          # 监控工具
�?  �?  └── api_monitor.go  # API 更新监控
�?  ├── performance/         # 性能分析
�?  �?  ├── profiler.go
�?  �?  └── memory.go
�?  └── profiling/           # 性能分析
�?      └── cpu.go
�?
├── .gitignore              # Git 忽略文件
├── .golangci.yml           # Linter 配置
├── go.mod                  # Go 模块
├── go.sum                  # Go 依赖锁定
├── LICENSE                 # 许可�?
├── Makefile                # 构建脚本
└── README.md               # 项目说明

```

---

## 目录说明

### `.github/`
存放 GitHub 相关配置文件�?

#### `workflows/`
- **`ci.yml`**: 持续集成工作流，每次 push �?PR 时运行测试、linter
- **`release.yml`**: 自动发布工作流，�?tag 时自动发布新版本
- **`doc-check.yml`**: 定期检查官方文档更�?

---

### `api/`
存放�?OpenAPI/Swagger 规范自动生成�?API 模型�?

**特点**�?
- �?自动生成，不手动编辑
- �?每个 API 一个子目录
- �?包含请求/响应结构�?
- �?包含枚举和常�?

**生成命令**�?
```bash
make generate-models
```

---

### `cmd/`
存放命令行工具和可执行文件�?

#### `generator/`
代码生成器工具，用于�?OpenAPI 规范生成 Go 代码�?

**使用方式**�?
```bash
go run cmd/generator/main.go -input openapi.json -output api/orders
```

---

### `docs/`
存放项目文档�?

| 文档 | 说明 |
|------|------|
| `ARCHITECTURE.md` | 架构设计和分层说�?|
| `DEVELOPMENT.md` | 开发规范和流程 |
| `PROJECT_STRUCTURE.md` | 项目结构说明（本文档�?|
| `API_TRACKING.md` | API 更新追踪策略 |
| `CODE_STYLE.md` | 代码风格和命名规�?|
| `CONTRIBUTING.md` | 如何参与项目开�?|
| `VERSION_TRACKING.md` | 官方 SDK 版本追踪 |
| `GRANTLESS_OPERATIONS_GUIDE.md` | Grantless 操作指南 |

---

### `examples/`
存放示例代码�?

#### 目录结构
```
examples/
├── basic_usage/          # 基础用法
�?  └── main.go          # 基本�?API 调用
├── advanced_usage/       # 高级用法
�?  └── main.go          # 中间件、重试、RDT �?
├── orders/              # Orders API 专项示例
�?  ├── get_orders.go
�?  └── get_order_items.go
├── reports/             # Reports API 专项示例
�?  ├── create_report.go
�?  └── get_report.go
└── README.md            # 示例说明
```

**运行示例**�?
```bash
cd examples/basic_usage
go run main.go
```

---

### `internal/`
存放内部实现包，**不对外暴�?*�?

> ⚠️ **重要**: `internal/` 下的包只能被本项目内部使用，外部项目无法导入�?

#### `auth/` - 认证�?
- `client.go` - LWA 客户端实�?
- `credentials.go` - 凭证管理
- `token.go` - 令牌结构和缓�?
- `errors.go` - 认证相关错误

#### `transport/` - 传输�?
- `client.go` - HTTP 客户�?
- `middleware.go` - 中间件（UserAgent, Date, Logging 等）
- `retry.go` - 重试逻辑

#### `signer/` - 签名�?
- `lwa.go` - LWA 签名器（常规操作�?
- `rdt.go` - RDT 签名器（受限操作�?
- `chain.go` - 签名器链（组合多个签名器�?

#### `ratelimit/` - 速率限制�?
- `limiter.go` - 速率限制�?
- `bucket.go` - Token Bucket 算法实现

#### `codec/` - 编解码层
- `json.go` - JSON 编解�?
- `validator.go` - 数据验证

#### `models/` - 内部模型
- `common.go` - 内部通用模型

#### `utils/` - 工具函数
- `http.go` - HTTP 相关工具
- `time.go` - 时间处理工具
- `string.go` - 字符串工�?

---

### `pkg/` - 公开�?
存放对外暴露�?API�?

> �?**公开 API**: `pkg/` 下的包是项目的公开接口，外部项目可以导入使用�?

#### `spapi/` - SP-API 客户�?

**主文�?*�?
- `client.go` - 主客户端，管理所�?API
- `config.go` - 配置选项（使�?Functional Options 模式�?
- `errors.go` - 公开错误类型
- `regions.go` - 区域和市场定�?

**API 子目录结�?*�?

每个 API 目录（如 `orders/`、`reports/` 等）包含�?
1. **`client.go`** - 手写�?API 客户端封�?
2. **`client_test.go`** - 单元测试
3. **`model_*.go`** - 自动生成的类型定义（多个文件�?

**示例：Orders API 目录结构**�?
```
orders/
├── client.go               # 手写：Orders API 客户�?
├── client_test.go          # 手写：单元测�?
├── model_address.go        # 生成：Address 类型
├── model_order.go          # 生成：Order 类型
├── model_order_item.go     # 生成：OrderItem 类型
├── model_money.go          # 生成：Money 类型
└── ...                     # ~70�?model_*.go 文件
```

**为什么每�?API 有这么多 `model_*.go` 文件�?*

1. **符合官方 SDK 标准**�?
   - 官方 Python SDK�?0个文�?API
   - 官方 PHP SDK�?0个文�?API
   - 官方 Java SDK�?1个文�?API
   - 我们�?Go SDK�?0个文�?API �?

2. **Go 最佳实�?*�?
   - 一个类型一个文件（清晰、可维护�?
   - 避免大型单体文件
   - 符合 `swagger-codegen` 标准输出

3. **优势**�?
   - �?清晰的文件结�?
   - �?易于查找和修�?
   - �?Git diff 更友�?
   - �?团队协作更高�?

**导入方式**�?
```go
import (
    "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
    "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi/orders"
)
```

---

### `tests/`
存放集成测试和性能测试�?

#### `integration/` - 集成测试
真实环境（Sandbox）的集成测试�?

**运行方式**�?
```bash
make test-integration
```

#### `benchmarks/` - 性能测试
基准测试和压力测试�?

**运行方式**�?
```bash
make benchmark
```

---

### `tools/`
存放开发工具�?

#### `monitoring/` - 监控工具
- `api_monitor.go` - 监控官方文档�?OpenAPI 规范更新

#### `performance/` - 性能分析
- `profiler.go` - CPU/内存分析
- `memory.go` - 内存泄漏检�?

#### `profiling/` - 性能分析
- `cpu.go` - CPU 性能分析

**使用方式**�?
```bash
go run tools/monitoring/api_monitor.go
```

---

## 文件命名规范

### Go 源文�?
- **小写 + 下划�?*: `http_client.go`
- **测试文件**: `http_client_test.go`
- **避免缩写**: 使用 `credentials.go` 而不�?`cred.go`

### 包名
- **小写单词**: `auth`, `transport`, `signer`
- **简短有意义**: 避免 `pkg`, `utils`, `common` 这类过于通用的名�?
- **与目录名一�?*: `internal/auth` �?`package auth`

### 常量和变�?
- **驼峰命名**: `maxRetries`, `defaultTimeout`
- **导出常量**: `MaxRetries`, `DefaultTimeout`
- **枚举前缀**: `GrantTypeRefreshToken`, `GrantTypeClientCredentials`

---

## 依赖管理

### `go.mod`
定义项目依赖�?

```go
module github.com/vanling1111/amazon-sp-api-go-sdk

go 1.21

require (
    // 无外部依赖，只使�?Go 标准�?
)
```

### `go.sum`
依赖�?checksum 锁定文件�?

---

## 构建和脚�?

### `Makefile`
提供常用命令快捷方式�?

```makefile
.PHONY: test
test:
    go test -v -race -cover ./...

.PHONY: lint
lint:
    golangci-lint run

.PHONY: build
build:
    go build -o bin/spapi ./cmd/...

.PHONY: generate-models
generate-models:
    go run cmd/generator/main.go
```

**使用方式**�?
```bash
make test
make lint
make build
```

---

## Git 忽略

### `.gitignore`
```gitignore
# 二进制文�?
bin/
*.exe
*.dll
*.so
*.dylib

# 测试覆盖
*.out
coverage.txt

# IDE
.vscode/
.idea/
*.swp
*.swo

# 操作系统
.DS_Store
Thumbs.db

# 临时文件
tmp/
temp/
*.log
```

---

## 最佳实�?

### 1. 自动生成 API 类型定义
使用官方推荐�?`swagger-codegen` 工具�?OpenAPI 规范生成类型定义�?

**生成命令**�?
```bash
# 生成所�?API
powershell -ExecutionPolicy Bypass -File scripts/generate-apis-clean.ps1

# 或使�?Makefile
make generate-apis
```

**生成规则**�?
- �?只生�?`model_*.go` 文件（类型定义）
- �?不生�?`client.go`、`api.go`（我们手写）
- �?每个类型一个文件（符合 Go 惯例�?
- �?使用正确的包名（�?`package orders`�?

**重要**�?
- ⚠️ **不要手动编辑** `model_*.go` 文件
- ⚠️ 如果官方 OpenAPI 规范更新，重新运行生成脚�?
- ⚠️ 生成后立即提交到 Git

---

### 2. 添加�?API 支持
1. **生成类型定义**�?
   ```bash
   # 已在 scripts/generate-apis-clean.ps1 中自动完�?
   ```

2. **手写 API 客户�?*�?
   ```bash
   # �?pkg/spapi/<api-name>/ 目录创建 client.go
   touch pkg/spapi/<api-name>/client.go
   ```

3. **编写单元测试**�?
   ```bash
   touch pkg/spapi/<api-name>/client_test.go
   ```

4. **添加示例代码**�?
   ```bash
   mkdir examples/<api-name>
   touch examples/<api-name>/main.go
   ```

5. **添加集成测试**�?
   ```bash
   touch tests/integration/<api-name>_test.go
   ```

**示例流程**�?
```bash
# 1. 自动生成类型（已完成�?
# pkg/spapi/orders/model_*.go 已存�?

# 2. 手写客户�?
cat > pkg/spapi/orders/client.go << 'EOF'
package orders

import "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"

type Client struct {
    *spapi.Client
}

func NewClient(c *spapi.Client) *Client {
    return &Client{Client: c}
}

func (c *Client) GetOrders(...) (*GetOrdersResponse, error) {
    // 实现
}
EOF

# 3. 添加测试
cat > pkg/spapi/orders/client_test.go << 'EOF'
package orders

func TestClient_GetOrders(t *testing.T) {
    // 测试
}
EOF
```

---

### 3. 修改内部组件
1. 修改 `internal/` 下的对应文件
2. 更新对应的单元测�?
3. 更新相关文档

---

### 4. 添加新工�?
1. �?`tools/` 下创建对应目�?
2. 添加 `main.go`
3. �?`Makefile` 中添加构建命�?

---

## 参考资�?

- [Go 项目布局标准](https://github.com/golang-standards/project-layout)
- [Google Go 风格指南](https://google.github.io/styleguide/go/)
- [Amazon SP-API 官方文档](https://developer-docs.amazon.com/sp-api/docs/)

