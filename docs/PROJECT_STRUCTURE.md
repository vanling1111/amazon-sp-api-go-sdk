# 项目结构

## 目录说明

```
amazon-sp-api-go-sdk/
├── cmd/                  # 命令行工具
│   ├── api-monitor/     # API 监控工具
│   └── generator/       # 代码生成器
├── internal/            # 内部核心模块
│   ├── auth/           # 认证
│   ├── signer/         # 签名
│   ├── ratelimit/      # 速率限制
│   ├── transport/      # HTTP 传输
│   ├── codec/          # 编解码
│   ├── errors/         # 错误处理
│   ├── metrics/        # 指标
│   ├── models/         # 通用模型
│   └── utils/          # 工具函数
├── pkg/spapi/          # 公开 API
│   ├── client.go       # 主客户端
│   ├── config.go       # 配置
│   ├── orders-v0/      # Orders API
│   ├── feeds-v2021-06-30/  # Feeds API
│   └── ...             # 其他 57 个 API
├── examples/           # 使用示例
│   ├── basic_usage/   # 基本用法
│   ├── orders/        # 订单示例
│   └── ...
├── tests/             # 测试
│   ├── integration/   # 集成测试
│   └── benchmarks/    # 基准测试
├── tools/             # 工具
│   ├── monitoring/    # 监控工具
│   ├── performance/   # 性能工具
│   └── validation/    # 验证工具
├── scripts/           # 脚本
│   └── *.ps1         # PowerShell 脚本
├── docs/              # 文档
├── README.md          # 项目说明
├── CHANGELOG.md       # 变更日志
└── LICENSE            # 许可证
```

## 模块说明

### internal/ - 核心模块

不对外暴露的内部实现：
- `auth` - LWA 认证和令牌管理
- `signer` - 请求签名
- `ratelimit` - 速率限制
- `transport` - HTTP 客户端
- `codec` - JSON 编解码
- `errors` - 错误定义
- `metrics` - 指标记录
- `models` - Region/Marketplace 定义
- `utils` - HTTP/时间/字符串工具

### pkg/spapi/ - 公开 API

对外暴露的 API：
- `client.go` - 主客户端
- `config.go` - 配置选项
- `errors.go` - 公开错误
- `*-v*/` - 57 个 API 版本目录
  - `client.go` - API 客户端方法
  - `client_test.go` - 单元测试
  - `model_*.go` - 数据模型

### cmd/ - 命令行工具

可执行程序：
- `api-monitor` - 监控官方 API 变更
- `generator` - 代码生成工具

### examples/ - 示例代码

7 个完整的使用示例：
- 基本用法
- Orders API
- Feeds API
- Reports API
- Listings API
- Grantless 操作
- 高级用法

### tests/ - 测试

- `integration/` - 集成测试（需要真实凭证）
- `benchmarks/` - 性能基准测试

### tools/ - 工具库

辅助工具：
- `monitoring` - 监控和指标收集
- `performance` - 性能分析
- `validation` - 数据验证

### scripts/ - 自动化脚本

PowerShell 脚本：
- `generate-apis-versioned.ps1` - 生成 API 模型
- `generate-all-api-clients.ps1` - 生成 API 客户端
- `generate-api-client-tests.ps1` - 生成测试

## 文件命名

- Go 源文件：`snake_case.go`
- 测试文件：`*_test.go`
- 文档文件：`UPPERCASE.md`
- 配置文件：`.lowercase`

## 依赖管理

使用 Go Modules：
```bash
# 下载依赖
go mod download

# 整理依赖
go mod tidy

# 验证依赖
go mod verify
```

项目无外部依赖，仅使用 Go 标准库。

