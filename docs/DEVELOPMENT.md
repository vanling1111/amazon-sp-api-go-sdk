# 开发指南

## 环境要求

- Go 1.21+
- Git
- （可选）golangci-lint

## 快速开始

```bash
# 克隆仓库
git clone https://github.com/vanling1111/amazon-sp-api-go-sdk.git
cd amazon-sp-api-go-sdk

# 运行测试
go test ./...

# 构建项目
go build ./...
```

## 项目结构

```
├── cmd/           # CLI 工具
├── internal/      # 核心模块
├── pkg/spapi/     # API 实现
├── examples/      # 使用示例
├── tests/         # 测试
└── docs/          # 文档
```

## 开发流程

### 1. 创建分支

```bash
git checkout -b feature/my-feature
```

### 2. 开发功能

- 编写代码
- 添加测试
- 更新文档

### 3. 测试

```bash
# 运行测试
go test ./...

# 查看覆盖率
go test -cover ./...

# 代码检查
go vet ./...
```

### 4. 提交代码

```bash
git add .
git commit -m "feat: add new feature"
git push origin feature/my-feature
```

## 测试规范

- 为新功能添加测试
- 保持测试覆盖率
- 使用表驱动测试

## 代码风格

- 遵循 Go 官方代码规范
- 使用 `gofmt` 格式化代码
- 编写清晰的注释

## 参考资源

- [Go 官方文档](https://go.dev/doc/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Amazon SP-API 文档](https://developer-docs.amazon.com/sp-api/docs/)

---

更新时间：2025-10-03
