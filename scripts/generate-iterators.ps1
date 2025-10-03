# 自动为所有有分页的 API 生成迭代器

$ErrorActionPreference = "Stop"

$PKG_DIR = "C:\Users\Administrator\amazon-sp-api-go-sdk\pkg\spapi"

# API 配置：定义每个 API 的分页方法和响应结构
$APIConfigs = @(
    # 已完成的（跳过）
    # @{API="orders-v0"; Methods=@(...)}  # 已有
    # @{API="reports-v2021-06-30"; Methods=@(...)}  # 已有
    # @{API="catalog-items-v2022-04-01"; Methods=@(...)}  # 已有
    
    # 高优先级
    @{
        API="feeds-v2021-06-30"
        Methods=@(
            @{Name="GetFeeds"; ResponseField="feeds"; TokenField="nextToken"}
        )
    },
    @{
        API="fba-inventory-v1"
        Methods=@(
            @{Name="GetInventorySummaries"; ResponseField="inventorySummaries"; TokenField="nextToken"}
        )
    },
    @{
        API="finances-v0"
        Methods=@(
            @{Name="ListFinancialEvents"; ResponseField="FinancialEvents"; TokenField="NextToken"}
            @{Name="ListFinancialEventGroups"; ResponseField="FinancialEventGroupList"; TokenField="NextToken"}
        )
    },
    @{
        API="fulfillment-inbound-v0"
        Methods=@(
            @{Name="GetShipments"; ResponseField="ShipmentData"; TokenField="NextToken"}
            @{Name="GetShipmentItems"; ResponseField="ItemData"; TokenField="NextToken"}
        )
    },
    @{
        API="fulfillment-outbound-v2020-07-01"
        Methods=@(
            @{Name="ListAllFulfillmentOrders"; ResponseField="fulfillmentOrders"; TokenField="nextToken"}
        )
    },
    @{
        API="catalog-items-v2020-12-01"
        Methods=@(
            @{Name="SearchCatalogItems"; ResponseField="items"; TokenField="nextToken"}
        )
    },
    @{
        API="listings-items-v2021-08-01"
        Methods=@(
            @{Name="GetListingsItem"; ResponseField="items"; TokenField="nextToken"}
        )
    },
    @{
        API="finances-v2024-06-19"
        Methods=@(
            @{Name="ListTransactions"; ResponseField="transactions"; TokenField="nextToken"}
        )
    },
    
    # 中优先级
    @{
        API="fulfillment-inbound-v2024-03-20"
        Methods=@(
            @{Name="ListInboundPlans"; ResponseField="inboundPlans"; TokenField="nextToken"}
        )
    },
    @{
        API="invoices-v2024-06-19"
        Methods=@(
            @{Name="GetInvoices"; ResponseField="invoices"; TokenField="nextToken"}
        )
    },
    @{
        API="seller-wallet-v2024-03-01"
        Methods=@(
            @{Name="ListTransactions"; ResponseField="transactions"; TokenField="nextToken"}
        )
    },
    @{
        API="services-v1"
        Methods=@(
            @{Name="GetServiceJobs"; ResponseField="jobs"; TokenField="nextToken"}
        )
    },
    @{
        API="supply-sources-v2020-07-01"
        Methods=@(
            @{Name="GetSupplySources"; ResponseField="supplySources"; TokenField="nextToken"}
        )
    },
    @{
        API="vehicles-v2024-11-01"
        Methods=@(
            @{Name="GetVehicles"; ResponseField="vehicles"; TokenField="nextToken"}
        )
    },
    @{
        API="aplus-content-v2020-11-01"
        Methods=@(
            @{Name="SearchContentDocuments"; ResponseField="contentMetadataRecords"; TokenField="nextToken"}
        )
    },
    @{
        API="data-kiosk-v2023-11-15"
        Methods=@(
            @{Name="GetQueries"; ResponseField="queries"; TokenField="nextToken"}
        )
    },
    
    # Vendor 系列
    @{
        API="vendor-orders-v1"
        Methods=@(
            @{Name="GetPurchaseOrders"; ResponseField="orders"; TokenField="nextToken"}
        )
    },
    @{
        API="vendor-shipments-v1"
        Methods=@(
            @{Name="GetShipments"; ResponseField="shipments"; TokenField="nextToken"}
        )
    },
    @{
        API="vendor-direct-fulfillment-orders-v1"
        Methods=@(
            @{Name="GetOrders"; ResponseField="orders"; TokenField="nextToken"}
        )
    },
    @{
        API="vendor-direct-fulfillment-shipping-v1"
        Methods=@(
            @{Name="GetShippingLabels"; ResponseField="shippingLabels"; TokenField="nextToken"}
        )
    }
)

Write-Host "=== Generating Iterators for All APIs ===" -ForegroundColor Cyan
Write-Host "Total APIs to process: $($APIConfigs.Count)" -ForegroundColor Green
Write-Host ""

$SuccessCount = 0
$SkippedCount = 0

foreach ($config in $APIConfigs) {
    $apiName = $config.API
    $apiDir = Join-Path $PKG_DIR $apiName
    $iteratorFile = Join-Path $apiDir "iterator.go"
    
    Write-Host "[$($SuccessCount + $SkippedCount + 1)/$($APIConfigs.Count)] $apiName" -ForegroundColor Yellow
    
    # 检查 API 目录是否存在
    if (!(Test-Path $apiDir)) {
        Write-Host "  ⚠ Skipped: API directory not found" -ForegroundColor Gray
        $SkippedCount++
        continue
    }
    
    # 检查是否已有迭代器
    if (Test-Path $iteratorFile) {
        Write-Host "  ✓ Skipped: Iterator already exists" -ForegroundColor Gray
        $SkippedCount++
        continue
    }
    
    # 生成包名（替换 - 为 _）
    $packageName = $apiName -replace '-','_'
    
    # 生成迭代器代码
    $iteratorContent = @"
// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package $packageName

import (
	"context"
	"encoding/json"
	"iter"

	"github.com/pkg/errors"
)

"@
    
    # 为每个方法生成迭代器
    foreach ($method in $config.Methods) {
        $methodName = $method.Name
        $responseField = $method.ResponseField
        $tokenField = $method.TokenField
        $iteratorName = "Iterate" + $methodName.Replace("Get", "").Replace("List", "").Replace("Search", "")
        
        $iteratorContent += @"

// $iteratorName 返回迭代器，自动处理分页。
//
// 使用 Go 1.25 迭代器特性，自动处理 $tokenField 分页逻辑。
//
// 示例:
//   for item, err := range client.$iteratorName(ctx, query) {
//       if err != nil { return err }
//       process(item)
//   }
func (c *Client) $iteratorName(ctx context.Context, query map[string]string) iter.Seq2[map[string]interface{}, error] {
	return func(yield func(map[string]interface{}, error) bool) {
		currentQuery := make(map[string]string)
		for k, v := range query {
			currentQuery[k] = v
		}

		for {
			result, err := c.$methodName(ctx, currentQuery)
			if err != nil {
				yield(nil, errors.Wrap(err, "failed to call $methodName"))
				return
			}

			resultBytes, err := json.Marshal(result)
			if err != nil {
				yield(nil, errors.Wrap(err, "failed to marshal result"))
				return
			}

			var response map[string]interface{}
			if err := json.Unmarshal(resultBytes, &response); err != nil {
				yield(nil, errors.Wrap(err, "failed to unmarshal response"))
				return
			}

			// 获取数据数组
			items, ok := response["$responseField"].([]interface{})
			if !ok || items == nil {
				break
			}

			// 遍历当前页
			for _, item := range items {
				itemMap, ok := item.(map[string]interface{})
				if !ok {
					continue
				}
				if !yield(itemMap, nil) {
					return
				}
			}

			// 检查下一页
			nextToken, _ := response["$tokenField"].(string)
			if nextToken == "" {
				break
			}

			currentQuery["$tokenField"] = nextToken
		}
	}
}

"@
    }
    
    # 写入文件
    Set-Content -Path $iteratorFile -Value $iteratorContent -Encoding UTF8
    
    Write-Host "  ✓ Generated iterator.go" -ForegroundColor Green
    $SuccessCount++
}

Write-Host ""
Write-Host "=== Generation Complete ===" -ForegroundColor Cyan
Write-Host "Success: $SuccessCount" -ForegroundColor Green
Write-Host "Skipped: $SkippedCount" -ForegroundColor Gray
Write-Host ""

