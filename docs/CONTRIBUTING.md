# 贡献指南

欢迎参与 Amazon SP-API Go SDK 项目的开发！

---

## 🚨 核心原则 - 必须严格遵守

### 1. ❌ 禁止参考其他语言的官方 SDK
- 不得查看或参考 Java、Python、Node.js、C#、PHP 等任何语言的官方 SDK 源码
- 不得基于其他 SDK 的实现来推断 API 行为

### 2. ✅ 只参考官方 SP-API 文档
- **唯一权威来源**: https://developer-docs.amazon.com/sp-api/docs/
- **文档验证**: 直接访问和参考官方文档内容
- **所有实现必须有文档依据**: 每个功能都必须能追溯到官方文档的对应章节

### 3. 🚫 禁止猜测开发
- 不得基于假设、推测或个人经验进行开发
- 如果文档不明确，查找更多官方资料和OpenAPI规范
- 无法从官方文档确认时，应提出问题而不是盲目实现

---

## 如何贡献

### 1. Fork 项目

```bash
# 1. Fork 仓库到你的 GitHub 账号

# 2. 克隆你的 fork
git clone https://github.com/your-username/amazon-sp-api-go-sdk.git
cd amazon-sp-api-go-sdk

# 3. 添加上游仓库
git remote add upstream https://github.com/original-owner/amazon-sp-api-go-sdk.git
```

---

### 2. 创建分支

```bash
# 从 main 分支创建新分支
git checkout -b feature/your-feature-name

# 或者修复 bug
git checkout -b fix/issue-number-description
```

**分支命名规范**:
- `feature/` - 新功能
- `fix/` - Bug 修复
- `docs/` - 文档更新
- `refactor/` - 代码重构
- `test/` - 测试相关

---

### 3. 开发前准备

#### 阅读官方文档

**必须完成**:
- [ ] 访问并阅读相关的官方文档章节
- [ ] 完整理解 API 的请求格式、响应格式、错误处理
- [ ] 记录官方文档的关键要求

**示例**:
```bash
# 访问官方文档
# 在浏览器中打开文档页面阅读
https://developer-docs.amazon.com/sp-api/docs/connecting-to-the-selling-partner-api
```

#### 阅读项目文档

- [ ] [开发规范](DEVELOPMENT.md)
- [ ] [架构设计](ARCHITECTURE.md)
- [ ] [代码风格](CODE_STYLE.md)
- [ ] [项目结构](PROJECT_STRUCTURE.md)

---

### 4. 开发流程

#### 步骤 1: 设计

1. **创建 Issue** 描述你要实现的功能或修复的 bug
2. **讨论方案** 在 Issue 中与维护者讨论实现方案
3. **获得批准** 等待维护者批准后再开始开发

#### 步骤 2: 实现

**强制要求**:
1. ✅ 所有实现必须基于官方文档
2. ✅ 添加完整的中文注释（Google 风格）
3. ✅ 实现完整的错误处理
4. ✅ 代码符合 Go 官方规范
5. ✅ 编写单元测试（覆盖率 > 90%）
6. ✅ 添加使用示例

**代码示例**:
```go
// GetOrders 获取订单列表。
//
// 此方法根据提供的查询参数获取订单列表，
// 支持按创建时间、更新时间等条件过滤。
//
// 参数:
//   - ctx: 请求上下文
//   - req: 查询请求参数
//
// 返回值:
//   - *GetOrdersResponse: 订单列表响应
//   - error: 如果请求失败，返回错误
//
// 示例:
//
//	req := &GetOrdersRequest{
//	    MarketplaceIDs: []string{"ATVPDKIKX0DER"},
//	    CreatedAfter:   time.Now().Add(-24 * time.Hour),
//	}
//	resp, err := client.GetOrders(ctx, req)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, order := range resp.Orders {
//	    fmt.Printf("Order ID: %s\n", order.AmazonOrderID)
//	}
//
// 官方文档:
//   - https://developer-docs.amazon.com/sp-api/docs/orders-api-v0-reference#getorders
func (c *OrdersAPI) GetOrders(ctx context.Context, req *GetOrdersRequest) (*GetOrdersResponse, error) {
    // 实现...
}
```

#### 步骤 3: 测试

**必须完成**:
```bash
# 1. 运行单元测试
go test -v ./...

# 2. 检查测试覆盖率
go test -cover ./...

# 3. 运行 linter
golangci-lint run

# 4. 格式化代码
gofmt -w .
goimports -w .
```

**测试覆盖率要求**:
- 新代码: **≥ 90%**
- 现有代码: 不降低整体覆盖率

#### 步骤 4: 文档

**必须完成**:
- [ ] 更新相关 API 文档
- [ ] 添加使用示例到 `examples/`
- [ ] 更新 `README.md`（如果需要）
- [ ] 记录官方文档来源

**示例目录**:
```
examples/
  orders/
    get_orders.go          # 获取订单列表示例
    get_order_items.go     # 获取订单项示例
    README.md              # 示例说明（包含官方文档链接）
```

---

### 5. 提交代码

#### Commit 规范

**格式**: `<type>(<scope>): <subject>`

**类型**:
- `feat`: 新功能
- `fix`: Bug 修复
- `docs`: 文档更新
- `style`: 代码格式（不影响代码运行的变动）
- `refactor`: 重构
- `test`: 测试相关
- `chore`: 构建过程或辅助工具的变动

**示例**:
```bash
git commit -m "feat(orders): add GetOrders API support

- Implement GetOrders method
- Add request/response models
- Add unit tests (coverage: 95%)
- Add usage example

Official docs: https://developer-docs.amazon.com/sp-api/docs/orders-api-v0-reference#getorders"
```

#### 推送代码

```bash
# 推送到你的 fork
git push origin feature/your-feature-name
```

---

### 6. 创建 Pull Request

#### PR 标题

**格式**: `<type>: <brief description>`

**示例**:
- `feat: add Orders API support`
- `fix: correct LWA token caching logic`
- `docs: update README with grantless operations`

#### PR 描述模板

```markdown
## 📝 变更说明

简要描述这个 PR 做了什么。

## 🔗 官方文档依据

- [连接到 SP-API](https://developer-docs.amazon.com/sp-api/docs/connecting-to-the-selling-partner-api)
- [Orders API Reference](https://developer-docs.amazon.com/sp-api/docs/orders-api-v0-reference)

## ✅ 检查清单

开发前:
- [ ] 已访问并阅读相关官方文档
- [ ] 已阅读 DEVELOPMENT.md 和 CODE_STYLE.md
- [ ] 在 Issue 中讨论并获得批准

实现:
- [ ] 代码基于官方文档，未参考其他语言 SDK
- [ ] 添加了完整的中文注释（Google 风格）
- [ ] 实现了完整的错误处理
- [ ] 代码符合 Go 官方规范

测试:
- [ ] 添加了单元测试（覆盖率 ≥ 90%）
- [ ] 所有测试通过 (`go test ./...`)
- [ ] Linter 通过 (`golangci-lint run`)
- [ ] 代码已格式化 (`gofmt`, `goimports`)

文档:
- [ ] 更新了相关文档
- [ ] 添加了使用示例
- [ ] 记录了官方文档来源

## 📊 测试结果

```bash
# 测试覆盖率
$ go test -cover ./internal/orders/...
ok      internal/orders    0.234s    coverage: 95.2% of statements

# Linter 结果
$ golangci-lint run
# 无问题
```

## 📸 截图/示例（如适用）

```go
// 使用示例
client := spapi.NewClient(config)
orders, err := client.Orders.GetOrders(ctx, req)
```

## 🔍 相关 Issue

Closes #123
```

---

### 7. 代码审查

**审查重点**:
1. ✅ 是否严格基于官方文档
2. ✅ 是否有完整的中文注释
3. ✅ 是否有完整的错误处理
4. ✅ 测试覆盖率是否达标
5. ✅ 是否符合代码风格

**审查流程**:
1. 维护者审查代码
2. 提出修改建议
3. 贡献者修改代码
4. 再次审查
5. 合并到 main

---

## 开发规范

### 1. 目录结构

新增 API 时:
```
pkg/spapi/
  orders.go           # Orders API
  orders_test.go      # 单元测试

examples/
  orders/
    get_orders.go     # 使用示例
    README.md         # 示例说明

api/
  orders/             # 自动生成的模型
    models.go
```

---

### 2. 错误处理

**定义错误**:
```go
var (
    // ErrOrderNotFound 表示订单不存在
    ErrOrderNotFound = errors.New("order not found")
    
    // ErrInvalidMarketplace 表示市场 ID 无效
    ErrInvalidMarketplace = errors.New("invalid marketplace ID")
)
```

**包装错误**:
```go
if err != nil {
    return nil, fmt.Errorf("fetch orders: %w", err)
}
```

---

### 3. 测试规范

**表驱动测试**:
```go
func TestGetOrders(t *testing.T) {
    tests := []struct {
        name    string
        req     *GetOrdersRequest
        want    *GetOrdersResponse
        wantErr bool
    }{
        {
            name: "success",
            req: &GetOrdersRequest{
                MarketplaceIDs: []string{"ATVPDKIKX0DER"},
            },
            wantErr: false,
        },
        {
            name: "missing marketplace ID",
            req:  &GetOrdersRequest{},
            wantErr: true,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := client.GetOrders(ctx, tt.req)
            if (err != nil) != tt.wantErr {
                t.Errorf("GetOrders() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            // 验证结果...
        })
    }
}
```

---

## 常见问题

### Q1: 如何访问官方文档？

**访问官方文档**:
```bash
# 访问官方文档网站
go run tools/doc_reader/main.go \
  --url "https://developer-docs.amazon.com/sp-api/docs/..."
```

---

### Q2: 如何验证实现符合官方规范？

**创建验证清单**:
1. 列出官方文档的所有要求
2. 逐项验证代码实现
3. 记录验证结果

**示例清单**:
```markdown
## LWA 认证验证清单

- [x] 请求格式: application/x-www-form-urlencoded
- [x] Grant Type: refresh_token
- [x] 必填字段: client_id, client_secret, refresh_token
- [x] 响应格式: JSON
- [x] 令牌头: x-amz-access-token
- [x] 过期处理: 提前 60 秒刷新
```

---

### Q3: 如何处理官方文档不明确的情况？

**步骤**:
1. 搜索更多官方资料和文档
2. 查看官方 OpenAPI 规范
3. 在 Issue 中提出问题
4. 等待官方明确或社区讨论

**不要**:
- ❌ 猜测实现
- ❌ 参考其他语言 SDK
- ❌ 基于假设开发

---

### Q4: 测试覆盖率如何达到 90%？

**策略**:
1. 测试所有公开方法
2. 测试所有错误路径
3. 测试边界条件
4. 使用表驱动测试

**查看覆盖率**:
```bash
# 生成覆盖率报告
go test -coverprofile=coverage.out ./...

# 查看详细报告
go tool cover -html=coverage.out
```

---

## 联系维护者

- **GitHub Issues**: 提交 Bug 或功能请求
- **GitHub Discussions**: 技术讨论
- **Email**: (维护者邮箱)

---

## 行为准则

参与本项目即表示你同意遵守我们的行为准则:

1. **尊重他人**: 尊重所有贡献者和用户
2. **建设性反馈**: 提供具体、有帮助的反馈
3. **专业态度**: 保持专业和友好的交流
4. **遵守规范**: 严格遵守项目的开发规范

---

## 许可证

贡献的代码将采用与项目相同的许可证 (Apache License 2.0)。

---

## 致谢

感谢所有贡献者的付出！

每个 PR 都会在 Release Notes 中致谢。

---

**祝开发愉快！🎉**

