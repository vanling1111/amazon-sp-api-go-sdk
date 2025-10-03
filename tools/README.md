# 开发工具目录

本目录包含项目开发和维护所需的工具。

## 工具列表

### monitoring/ - API 监控工具

监控官方 SP-API 文档和 OpenAPI 规范的更新。

#### api_monitor.go - 文档监控

监控官方文档和API规范的变更。

**功能**:
- ✅ 定期访问官方文档页面
- ✅ 提取关键内容并计算哈希
- ✅ 检测内容变更
- ✅ 自动创建 GitHub Issue

**用法**:
```bash
go run tools/monitoring/api_monitor.go
```

**配置** (`config/monitor.yml`):
```yaml
interval: 24h
pages:
  - url: https://developer-docs.amazon.com/sp-api/docs/connecting-to-the-selling-partner-api
    selector: "#main-content"
  - url: https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits
    selector: "#main-content"

notifications:
  github:
    enabled: true
    repo: yourusername/amazon-sp-api-go-sdk
```

#### openapi_sync.go - OpenAPI 规范同步

同步官方 OpenAPI 规范文件。

**功能**:
- ✅ 从 GitHub 拉取最新规范
- ✅ 对比本地版本
- ✅ 标记需要更新的文件

**用法**:
```bash
# 检查更新
go run tools/monitoring/openapi_sync.go --check

# 同步规范
go run tools/monitoring/openapi_sync.go --sync

# 查看差异
go run tools/monitoring/openapi_sync.go --diff models/orders-api-model.json
```

### performance/ - 性能分析工具

分析和优化 SDK 性能。

#### profiler.go - 性能分析

CPU 和内存性能分析。

**用法**:
```bash
# CPU 分析
go run tools/performance/profiler.go -type=cpu -output=cpu.prof

# 内存分析
go run tools/performance/profiler.go -type=mem -output=mem.prof

# 查看分析结果
go tool pprof -http=:8080 cpu.prof
```

#### memory.go - 内存泄漏检测

检测潜在的内存泄漏问题。

**用法**:
```bash
go run tools/performance/memory.go -duration=5m
```

## GitHub Actions 集成

### 文档监控工作流

`.github/workflows/doc-check.yml`:
```yaml
name: Documentation Update Check

on:
  schedule:
    - cron: '0 0 * * *'  # 每天 UTC 00:00
  workflow_dispatch:

jobs:
  check-docs:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.21'
      
      - name: Run Monitor
        run: go run tools/monitoring/api_monitor.go
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

### OpenAPI 同步工作流

`.github/workflows/openapi-sync.yml`:
```yaml
name: OpenAPI Spec Sync

on:
  schedule:
    - cron: '0 0 * * 1'  # 每周一 UTC 00:00
  workflow_dispatch:

jobs:
  sync-openapi:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.21'
      
      - name: Check Updates
        id: check
        run: go run tools/monitoring/openapi_sync.go --check
      
      - name: Create PR
        if: steps.check.outputs.changed == 'true'
        uses: peter-evans/create-pull-request@v6
        with:
          title: 'chore: sync OpenAPI specifications'
          branch: sync-openapi
          labels: openapi,automated
```

## 工具开发指南

### 添加新工具

1. 在 `tools/` 下创建新目录
2. 创建 `main.go` 或相应的工具文件
3. 添加 README 说明
4. 在本文档中添加说明

### 目录结构示例

```
tools/
├── monitoring/
│   ├── api_monitor.go
│   ├── openapi_sync.go
│   └── README.md
├── performance/
│   ├── profiler.go
│   ├── memory.go
│   └── README.md
├── validation/
│   ├── config_validator.go
│   └── README.md
└── README.md  # 本文件
```

### 工具开发最佳实践

1. **单一职责** - 每个工具专注一个任务
2. **命令行参数** - 使用 `flag` 包处理参数
3. **清晰输出** - 提供有意义的日志和进度信息
4. **错误处理** - 完整的错误处理和退出码
5. **文档** - 添加使用说明和示例

### 示例工具模板

```go
// tools/example/main.go
package main

import (
    "flag"
    "fmt"
    "log"
    "os"
)

var (
    input  = flag.String("input", "", "输入文件路径")
    output = flag.String("output", "", "输出文件路径")
    debug  = flag.Bool("debug", false, "启用调试模式")
)

func main() {
    flag.Parse()

    if *input == "" {
        log.Fatal("必须指定输入文件")
    }

    if *debug {
        log.Println("调试模式已启用")
    }

    if err := run(); err != nil {
        log.Fatal(err)
        os.Exit(1)
    }

    fmt.Println("完成！")
}

func run() error {
    // 工具逻辑
    return nil
}
```

## 依赖管理

工具所需的依赖应添加到项目的 `go.mod` 中：

```go
// go.mod
module github.com/yourusername/amazon-sp-api-go-sdk

go 1.21

require (
    // 标准库已足够，无需外部依赖
)
```

## 参考资料

- [Go pprof](https://pkg.go.dev/runtime/pprof)
- [GitHub API](https://docs.github.com/en/rest)
- [Prometheus Metrics](https://prometheus.io/docs/introduction/overview/)

