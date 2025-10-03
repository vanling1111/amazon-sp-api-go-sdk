# 贡献指南

感谢您对 Amazon SP-API Go SDK 项目的关注！

## 如何贡献

### 1. 提交 Issue

发现 Bug 或有功能建议？欢迎创建 Issue：
- 描述问题或建议
- 提供复现步骤（如果是 Bug）
- 说明期望的行为

### 2. 提交 Pull Request

#### 准备工作

```bash
# Fork 并克隆仓库
git clone https://github.com/your-username/amazon-sp-api-go-sdk.git
cd amazon-sp-api-go-sdk

# 创建新分支
git checkout -b feature/your-feature-name
```

#### 开发流程

1. 编写代码
2. 添加测试
3. 运行测试确保通过
4. 提交代码

```bash
# 运行测试
go test ./...

# 代码检查
go vet ./...

# 格式化代码
gofmt -w .
```

#### 提交规范

使用清晰的 commit 消息：

```
feat: add new API support
fix: correct token caching issue
docs: update README
```

### 3. 代码要求

- 遵循 Go 最佳实践
- 添加适当的测试
- 编写清晰的注释
- 确保所有测试通过

## 获取帮助

- [GitHub Issues](https://github.com/vanling1111/amazon-sp-api-go-sdk/issues)
- [GitHub Discussions](https://github.com/vanling1111/amazon-sp-api-go-sdk/discussions)
- [官方文档](https://developer-docs.amazon.com/sp-api/docs/)

## 许可证

贡献的代码将采用 Apache License 2.0。

---

感谢您的贡献！🎉
