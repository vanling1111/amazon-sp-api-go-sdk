# 集成测试

集成测试需要真实的 Amazon SP-API 凭证。

## 运行集成测试

### 1. 设置环境变量

```bash
export SP_API_CLIENT_ID="your-client-id"
export SP_API_CLIENT_SECRET="your-client-secret"
export SP_API_REFRESH_TOKEN="your-refresh-token"
export RUN_INTEGRATION_TESTS=1
```

Windows PowerShell:
```powershell
$env:SP_API_CLIENT_ID="your-client-id"
$env:SP_API_CLIENT_SECRET="your-client-secret"
$env:SP_API_REFRESH_TOKEN="your-refresh-token"
$env:RUN_INTEGRATION_TESTS=1
```

### 2. 运行测试

```bash
# 运行所有集成测试
go test -tags=integration ./tests/integration/...

# 运行单个测试
go test -tags=integration ./tests/integration -run TestOrders

# 带详细输出
go test -tags=integration -v ./tests/integration/...
```

## 注意事项

1. ⚠️ 集成测试会调用真实的 Amazon SP-API
2. ⚠️ 可能会受到速率限制
3. ⚠️ 某些操作可能会产生实际影响（如创建订单、Feed等）
4. ✅ 建议在沙盒环境或测试账号中运行
5. ✅ 默认情况下集成测试会被跳过

## 测试覆盖

当前集成测试覆盖的 API:
- [x] Orders v0 - 订单管理
- [x] Feeds v2021-06-30 - Feed 处理
- [x] Reports v2021-06-30 - 报告生成
- [x] Catalog Items v2022-04-01 - 商品目录
- [x] Product Pricing v0 - 商品定价
- [x] FBA Inventory v1 - 库存管理
- [x] Listings Items v2021-08-01 - Listings 管理
- [x] Sellers v1 - 卖家信息
- [x] Tokens v2021-03-01 - RDT 令牌
- [x] Notifications v1 - 通知订阅（Grantless）

总计: **10 个核心 API 的集成测试**

