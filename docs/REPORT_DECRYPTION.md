# 报告解密使用指南

本指南介绍如何使用 SDK 的自动报告解密功能。

## 📋 背景

Amazon SP-API 从 2020-09-04 版本开始，返回的报告数据使用 **AES-256-CBC** 加密。这是为了保护敏感的订单、客户信息。

传统方式需要：
1. 调用 `GetReportDocument` API
2. 从响应中获取加密密钥和 IV
3. 下载加密的报告内容
4. 手动解密（AES-256-CBC）
5. 移除 PKCS7 填充

从 v1.1.0 开始，SDK 提供**一行代码自动解密**功能！

## ✨ 核心优势

- ✅ **一行代码** - 自动下载+解密
- ✅ **自动检测** - 自动判断是否加密
- ✅ **完整实现** - AES-256-CBC + PKCS7 padding
- ✅ **错误处理** - 完整的验证和错误提示
- ✅ **生产级** - 13 个单元测试验证

## 🚀 快速开始

### 基础用法

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
    "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi/reports-v2021-06-30"
)

func main() {
    // 创建客户端
    baseClient, _ := spapi.NewClient(...)
    reportsClient := reports_v2021_06_30.NewClient(baseClient)
    ctx := context.Background()
    
    // 假设你已经有 reportDocumentID
    reportDocumentID := "amzn1.tortuga.xxx"
    
    // 一行代码获取解密后的报告
    decrypted, err := reportsClient.GetReportDocumentDecrypted(ctx, reportDocumentID)
    if err != nil {
        log.Fatal(err)
    }
    
    // 直接使用解密后的数据
    fmt.Println(string(decrypted))
}
```

## 📖 完整流程

### 1. 创建报告

```go
// 创建订单报告
request := map[string]interface{}{
    "reportType": "GET_FLAT_FILE_ALL_ORDERS_DATA_BY_ORDER_DATE",
    "marketplaceIds": []string{"ATVPDKIKX0DER"},
    "dataStartTime": "2025-01-01T00:00:00Z",
    "dataEndTime":   "2025-01-31T23:59:59Z",
}

result, err := reportsClient.CreateReport(ctx, request)
if err != nil {
    log.Fatal(err)
}

reportID := result.(map[string]interface{})["reportId"].(string)
log.Printf("Report ID: %s", reportID)
```

### 2. 等待报告生成

```go
// 轮询报告状态
for {
    result, err := reportsClient.GetReport(ctx, reportID, nil)
    if err != nil {
        log.Fatal(err)
    }
    
    report := result.(map[string]interface{})
    status := report["processingStatus"].(string)
    
    log.Printf("Status: %s", status)
    
    if status == "DONE" {
        reportDocumentID = report["reportDocumentId"].(string)
        break
    } else if status == "FATAL" || status == "CANCELLED" {
        log.Fatal("Report generation failed")
    }
    
    time.Sleep(10 * time.Second)
}
```

### 3. 自动下载和解密

```go
// SDK 自动处理：
// 1. 调用 GetReportDocument 获取元数据
// 2. 从 URL 下载加密内容
// 3. 检测是否加密
// 4. 如果加密，使用 AES-256-CBC 解密
// 5. 移除 PKCS7 填充
// 6. 返回原始数据

decrypted, err := reportsClient.GetReportDocumentDecrypted(ctx, reportDocumentID)
if err != nil {
    log.Fatal(err)
}

// 解密后的数据可以直接使用
fmt.Println("Report size:", len(decrypted), "bytes")
```

### 4. 解析报告数据

```go
import (
    "encoding/csv"
    "strings"
)

// 大多数报告是 TSV 格式（Tab 分隔）
reader := csv.NewReader(strings.NewReader(string(decrypted)))
reader.Comma = '\t'

// 读取表头
headers, err := reader.Read()
if err != nil {
    log.Fatal(err)
}
fmt.Println("Columns:", headers)

// 读取数据行
for {
    row, err := reader.Read()
    if err != nil {
        break  // EOF
    }
    
    // 处理每一行
    fmt.Println(row)
}
```

## 🔐 加密详情

### 加密算法

Amazon 使用：
- **算法**：AES-256-CBC
- **密钥长度**：256 位（32 字节）
- **IV 长度**：128 位（16 字节）
- **填充**：PKCS7

### 加密响应格式

```json
{
  "reportDocumentId": "amzn1.tortuga.xxx",
  "url": "https://tortuga-prod-na.s3.amazonaws.com/...",
  "encryptionDetails": {
    "standard": "AES",
    "initializationVector": "BASE64_ENCODED_IV",
    "key": "BASE64_ENCODED_KEY"
  },
  "compressionAlgorithm": null
}
```

### 手动解密（如果需要）

如果你需要在其他地方手动解密（不通过 SDK）：

```go
import "github.com/vanling1111/amazon-sp-api-go-sdk/internal/crypto"

// 使用 SDK 的加密模块
decrypted, err := crypto.DecryptReport(
    encryptionDetails.Key,
    encryptionDetails.InitializationVector,
    encryptedData,
)
```

## 🎯 高级用法

### 保存解密后的报告

```go
decrypted, err := reportsClient.GetReportDocumentDecrypted(ctx, reportDocumentID)
if err != nil {
    log.Fatal(err)
}

// 保存到文件
filename := fmt.Sprintf("order_report_%s.csv", time.Now().Format("20060102"))
err = os.WriteFile(filename, decrypted, 0644)
if err != nil {
    log.Fatal(err)
}

log.Printf("Report saved to: %s", filename)
```

### 流式处理大报告

对于非常大的报告（100MB+），可以流式处理：

```go
// TODO: v1.2.0 将添加流式解密支持
// decryptedReader, err := reportsClient.GetReportDocumentDecryptedStream(ctx, reportDocumentID)
```

### 批量处理报告

```go
// 使用迭代器获取所有报告
for report, err := range reportsClient.IterateReports(ctx, query) {
    if err != nil {
        log.Printf("Error: %v", err)
        continue
    }
    
    status := report["processingStatus"].(string)
    if status != "DONE" {
        continue  // 跳过未完成的报告
    }
    
    reportDocumentID := report["reportDocumentId"].(string)
    
    // 下载并解密
    decrypted, err := reportsClient.GetReportDocumentDecrypted(ctx, reportDocumentID)
    if err != nil {
        log.Printf("Decrypt failed for %s: %v", reportDocumentID, err)
        continue
    }
    
    // 处理报告
    processReport(decrypted)
}
```

## ⚠️ 注意事项

### 1. 内存占用

解密会将整个报告加载到内存：

```go
// 如果报告是 500MB，会占用 500MB+ 内存
decrypted, _ := reportsClient.GetReportDocumentDecrypted(ctx, reportDocumentID)
```

建议：
- 对大报告使用流式处理（v1.2.0 将支持）
- 处理后及时释放内存
- 监控内存使用

### 2. 网络超时

大报告下载可能超时，建议设置更长的超时时间：

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
defer cancel()

decrypted, err := reportsClient.GetReportDocumentDecrypted(ctx, reportDocumentID)
```

### 3. 报告有效期

报告下载 URL 有效期通常是 5 分钟，要及时下载：

```go
// ✅ 好：立即下载
reportDocumentID := report["reportDocumentId"].(string)
decrypted, _ := reportsClient.GetReportDocumentDecrypted(ctx, reportDocumentID)

// ❌ 差：延迟下载可能失败
reportDocumentID := report["reportDocumentId"].(string)
time.Sleep(10 * time.Minute)  // URL 可能已过期
decrypted, _ := reportsClient.GetReportDocumentDecrypted(ctx, reportDocumentID)  // 失败
```

## 📚 支持的报告类型

所有 SP-API 报告都支持自动解密，包括：

- **订单报告** - `GET_FLAT_FILE_ALL_ORDERS_DATA_BY_ORDER_DATE`
- **库存报告** - `GET_FBA_INVENTORY_AGED_DATA`
- **财务报告** - `GET_V2_SETTLEMENT_REPORT_DATA_FLAT_FILE`
- **销售报告** - `GET_SALES_AND_TRAFFIC_REPORT`
- **Performance 报告** - `GET_V1_SELLER_PERFORMANCE_REPORT`
- 等所有报告类型

## 🔗 相关资源

- [完整示例代码](../examples/report-decryption/main.go)
- [Reports API 文档](https://developer-docs.amazon.com/sp-api/docs/reports-api-v2021-06-30-reference)
- [报告类型列表](https://developer-docs.amazon.com/sp-api/docs/report-type-values)

