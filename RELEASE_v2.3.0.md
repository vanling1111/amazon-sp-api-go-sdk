# v2.3.0 - Facade设计模式和架构优化

## 🏗️ 架构优化

### Added

#### Facade设计模式
- 新增 `internal/core/facade.go` - 核心门面层
- 封装所有内部组件（auth, transport, signer, ratelimit）
- 提供统一的访问接口
- 隐藏内部复杂性

### Changed

#### Client结构优化
- 简化 `Client` 结构，使用 Facade 封装
- 降低 `pkg/spapi` 与 `internal/*` 的直接耦合
- 通过 Facade 访问内部组件
- 更清晰的代码组织

**Before**:
```go
type Client struct {
    config           *Config
    lwaClient        *auth.Client      // 直接依赖
    httpClient       *transport.Client // 直接依赖
    signer           signer.Signer     // 直接依赖
    rateLimitManager *ratelimit.Manager // 直接依赖
}
```

**After**:
```go
type Client struct {
    config *Config
    facade *core.Facade  // ✅ Facade封装
}
```

### Benefits

- ✅ **更好的封装性** - 内部组件通过 Facade 统一访问
- ✅ **降低耦合度** - pkg/spapi 不再直接依赖多个 internal 包
- ✅ **提高可维护性** - 修改内部实现不影响外部接口
- ✅ **符合设计模式** - 遵循 Facade 模式最佳实践

### Example

```go
import "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"

// 使用方式完全兼容，无需修改代码
client, err := spapi.NewClient(
    spapi.WithRegion(spapi.RegionNA),
    spapi.WithCredentials("client-id", "secret", "token"),
)

// 内部通过 Facade 访问组件
// facade.GetLWAClient()
// facade.GetHTTPClient()
// facade.GetSigner()
// facade.GetRateLimitManager()
```

## 🐛 Bug Fixes

- 修复 linting 警告
- 修复错误处理
- 代码格式优化

## 📊 测试

- ✅ 所有测试通过
- ✅ 无 linting 警告
- ✅ 100% 向后兼容

## 📚 文档

- 更新 CHANGELOG.md
- 更新架构文档
- 添加 Facade 模式说明

## 🚀 安装

```bash
go get github.com/vanling1111/amazon-sp-api-go-sdk@v2.3.0
```

## 🔄 迁移指南

**无需任何代码修改！** 此版本完全向后兼容。

## 📝 完整变更日志

查看 [CHANGELOG.md](https://github.com/vanling1111/amazon-sp-api-go-sdk/blob/main/CHANGELOG.md) 获取完整的变更历史。

---

**Full Changelog**: https://github.com/vanling1111/amazon-sp-api-go-sdk/compare/v2.2.0...v2.3.0
