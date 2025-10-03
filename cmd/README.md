# 命令行工具目录

本目录存放项目的命令行工具和可执行程序。

## 工具列表

### generator - 代码生成器

从 OpenAPI 规范生成 Go 代码的工具。

**用法**:
```bash
go run cmd/generator/main.go -input <openapi.json> -output <output_dir>
```

**示例**:
```bash
# 生成 Orders API 模型
go run cmd/generator/main.go \
  -input https://github.com/amzn/selling-partner-api-models/raw/main/models/orders-api-model/ordersV0.json \
  -output api/orders
```

**功能**:
- ✅ 从 OpenAPI 规范生成请求/响应结构体
- ✅ 生成枚举和常量
- ✅ 生成 API 客户端接口
- ✅ 自动添加中文注释

## 添加新工具

在此目录下创建新的子目录，每个工具一个 `main.go` 文件：

```
cmd/
├── generator/
│   └── main.go
├── validator/          # 配置验证工具
│   └── main.go
└── monitor/            # API 监控工具
    └── main.go
```

## 构建

```bash
# 构建所有工具
go build -o bin/ ./cmd/...

# 构建特定工具
go build -o bin/generator ./cmd/generator
```

## 官方文档

- [OpenAPI Specification](https://spec.openapis.org/oas/v3.1.0)
- [SP-API Models](https://github.com/amzn/selling-partner-api-models)

