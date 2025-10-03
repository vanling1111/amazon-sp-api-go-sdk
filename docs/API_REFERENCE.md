# API 参考文档

本 SDK 提供完整的 GoDoc 注释，可通过多种方式查看 API 文档。

## 📖 在线查看

### 1. pkg.go.dev（推荐）

SDK 发布到 GitHub 后，会自动同步到 pkg.go.dev：

```
https://pkg.go.dev/github.com/vanling1111/amazon-sp-api-go-sdk
```

**特点**：
- ✅ 官方托管，自动更新
- ✅ 完整的代码索引和搜索
- ✅ 支持版本切换
- ✅ 跨包引用跳转

### 2. 本地查看

使用 `go doc` 命令查看：

```bash
# 查看包文档
go doc github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi

# 查看特定类型
go doc github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi.Client

# 查看特定方法
go doc github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi.NewClient

# 查看 Orders API
go doc github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi/orders-v0

# 查看所有导出的符号
go doc -all github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi
```

### 3. 本地 Web 服务器

启动本地文档服务器：

```bash
# 安装 godoc（Go 1.13+）
go install golang.org/x/tools/cmd/godoc@latest

# 启动服务器
godoc -http=:6060

# 浏览器访问
# http://localhost:6060/pkg/github.com/vanling1111/amazon-sp-api-go-sdk/
```

---

## 📦 核心包文档

### 主包

| 包路径 | 说明 | 文档链接 |
|--------|------|----------|
| `pkg/spapi` | SDK 主入口 | [查看](https://pkg.go.dev/github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi) |

### API 包（57 个）

#### 核心业务

| API | 包路径 | 说明 |
|-----|--------|------|
| Orders API | `pkg/spapi/orders-v0` | 订单管理 |
| Feeds API | `pkg/spapi/feeds-v2021-06-30` | 数据上传 |
| Reports API | `pkg/spapi/reports-v2021-06-30` | 报告下载 |
| Catalog Items | `pkg/spapi/catalog-items-v2022-04-01` | 商品目录 |
| Listings Items | `pkg/spapi/listings-items-v2021-08-01` | 商品列表 |

#### 库存物流

| API | 包路径 | 说明 |
|-----|--------|------|
| FBA Inventory | `pkg/spapi/fba-inventory-v1` | FBA 库存 |
| Fulfillment Inbound | `pkg/spapi/fulfillment-inbound-v2024-03-20` | 入库管理 |
| Fulfillment Outbound | `pkg/spapi/fulfillment-outbound-v2020-07-01` | 出库管理 |

#### 定价财务

| API | 包路径 | 说明 |
|-----|--------|------|
| Product Pricing | `pkg/spapi/product-pricing-v2022-05-01` | 商品定价 |
| Product Fees | `pkg/spapi/product-fees-v0` | 费用估算 |
| Finances | `pkg/spapi/finances-v2024-06-19` | 财务报告 |

**📋 完整列表**：查看 [pkg/spapi/](https://github.com/vanling1111/amazon-sp-api-go-sdk/tree/main/pkg/spapi) 目录

---

## 🔍 内部包文档

高级用户和贡献者可能需要的内部包：

| 包路径 | 说明 |
|--------|------|
| `internal/auth` | LWA 认证和 Token 管理 |
| `internal/ratelimit` | 速率限制（Token Bucket） |
| `internal/signer` | AWS 签名 v4 |
| `internal/transport` | HTTP 传输和重试 |
| `internal/logging` | 结构化日志（Zap） |
| `internal/circuit` | 熔断器 |
| `internal/crypto` | AES 加密解密 |
| `internal/tracing` | OpenTelemetry 追踪 |
| `internal/metrics` | 指标收集 |

---

## 📖 使用示例

### 快速查找 API

```bash
# 查找所有 Orders 相关的包
go list github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi/orders-v0

# 查看 Orders API 的所有方法
go doc -all github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi/orders-v0

# 查看特定方法的签名和文档
go doc github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi/orders-v0.Client.GetOrders
```

### IDE 集成

所有现代 Go IDE 都支持 GoDoc：

- **VS Code / Cursor**: 鼠标悬停查看文档，`Ctrl+Click` 跳转
- **GoLand**: `Ctrl+Q` 查看快速文档
- **Vim-go**: `:GoDoc` 命令

---

## 🌐 在线发布

### 自动发布到 pkg.go.dev

1. **确保项目公开**：GitHub 仓库设置为 public
2. **打标签**：推送 Git 标签触发索引
   ```bash
   git tag v1.3.0
   git push origin v1.3.0
   ```
3. **手动触发**：访问 https://pkg.go.dev/github.com/vanling1111/amazon-sp-api-go-sdk@v1.3.0
4. **等待索引**：通常 10-30 分钟完成

### 查看发布状态

访问：https://pkg.go.dev/github.com/vanling1111/amazon-sp-api-go-sdk

如果看到 "Module not found"，点击 "Request" 按钮手动请求索引。

---

## 📚 相关文档

- [官方 GoDoc 规范](https://go.dev/blog/godoc)
- [编写优质文档](https://go.dev/doc/comment)
- [pkg.go.dev 使用指南](https://go.dev/about)

---

**提示**：所有代码注释都使用中文编写，GoDoc 完全支持 UTF-8 显示。

