// Copyright 2025 Amazon SP-API Go SDK Authors.
//
// This file is part of Amazon SP-API Go SDK.
//
// Amazon SP-API Go SDK is dual-licensed:
//
// 1. GNU Affero General Public License v3.0 (AGPL-3.0) for open source use
//    - Free for personal, educational, and open source projects
//    - Your project must also be open sourced under AGPL-3.0
//    - See: https://www.gnu.org/licenses/agpl-3.0.html
//
// 2. Commercial License for proprietary/commercial use
//    - Required for any commercial, enterprise, or proprietary use
//    - Allows closed source distribution
//    - Contact: vanling1111@gmail.com
//
// Unless you have obtained a commercial license, this file is licensed
// under AGPL-3.0. By using this software, you agree to comply with the
// terms of the applicable license.
//
// This file is part of Amazon SP-API Go SDK.
//
// Amazon SP-API Go SDK is dual-licensed:
//
// 1. GNU Affero General Public License v3.0 (AGPL-3.0) for open source use
//    - Free for personal, educational, and open source projects
//    - Your project must also be open sourced under AGPL-3.0
//    - See: https://www.gnu.org/licenses/agpl-3.0.html
//
// 2. Commercial License for proprietary/commercial use
//    - Required for any commercial, enterprise, or proprietary use
//    - Allows closed source distribution
//    - Contact: vanling1111@gmail.com
//
// Unless you have obtained a commercial license, this file is licensed
// under AGPL-3.0. By using this software, you agree to comply with the
// terms of the applicable license.

// Package main 演示 Go 1.25 迭代器的使用。
//
// 此示例展示如何使用 SDK 的分页迭代器自动处理 Amazon SP-API 的分页响应。
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi/catalog-items-v2022-04-01"
	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi/orders-v0"
	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi/reports-v2021-06-30"
)

func main() {
	// 从环境变量获取配置
	clientID := os.Getenv("SP_API_CLIENT_ID")
	clientSecret := os.Getenv("SP_API_CLIENT_SECRET")
	refreshToken := os.Getenv("SP_API_REFRESH_TOKEN")

	if clientID == "" || clientSecret == "" || refreshToken == "" {
		log.Fatal("缺少必要的环境变量: SP_API_CLIENT_ID, SP_API_CLIENT_SECRET, SP_API_REFRESH_TOKEN")
	}

	// 创建客户端
	baseClient, err := spapi.NewClient(
		spapi.WithRegion(spapi.RegionNA),
		spapi.WithCredentials(clientID, clientSecret, refreshToken),
	)
	if err != nil {
		log.Fatalf("创建客户端失败: %v", err)
	}
	defer baseClient.Close()

	ctx := context.Background()

	// 示例 1: 迭代订单（自动分页）
	fmt.Println("=== 示例 1: 迭代订单 ===")
	iterateOrdersExample(ctx, baseClient)

	// 示例 2: 迭代订单项
	fmt.Println("\n=== 示例 2: 迭代订单项 ===")
	iterateOrderItemsExample(ctx, baseClient)

	// 示例 3: 迭代报告
	fmt.Println("\n=== 示例 3: 迭代报告 ===")
	iterateReportsExample(ctx, baseClient)

	// 示例 4: 迭代商品目录
	fmt.Println("\n=== 示例 4: 迭代商品目录 ===")
	iterateCatalogItemsExample(ctx, baseClient)

	// 示例 5: 提前退出迭代
	fmt.Println("\n=== 示例 5: 提前退出 ===")
	earlyExitExample(ctx, baseClient)

	// 示例 6: 并发处理
	fmt.Println("\n=== 示例 6: 并发处理 ===")
	concurrentProcessingExample(ctx, baseClient)
}

// iterateOrdersExample 演示订单迭代器
func iterateOrdersExample(ctx context.Context, baseClient *spapi.Client) {
	ordersClient := orders_v0.NewClient(baseClient)

	query := map[string]string{
		"MarketplaceIds": "ATVPDKIKX0DER",
		"CreatedAfter":   "2025-01-01T00:00:00Z",
	}

	// 使用 Go 1.25 迭代器：自动处理所有分页
	count := 0
	for order, err := range ordersClient.IterateOrders(ctx, query) {
		if err != nil {
			log.Printf("迭代错误: %v", err)
			break
		}

		count++
		orderID := order["AmazonOrderId"]
		orderTotal := order["OrderTotal"]
		fmt.Printf("  订单 %d: %s - %v\n", count, orderID, orderTotal)

		// 可以随时中断
		if count >= 10 {
			fmt.Println("  (仅显示前 10 个订单)")
			break
		}
	}

	fmt.Printf("总计处理订单: %d\n", count)
}

// iterateOrderItemsExample 演示订单项迭代器
func iterateOrderItemsExample(ctx context.Context, baseClient *spapi.Client) {
	ordersClient := orders_v0.NewClient(baseClient)

	// 假设已知的订单 ID
	orderID := "123-4567890-1234567"

	count := 0
	for item, err := range ordersClient.IterateOrderItems(ctx, orderID, nil) {
		if err != nil {
			log.Printf("迭代错误: %v", err)
			break
		}

		count++
		sku := item["SellerSKU"]
		qty := item["QuantityOrdered"]
		fmt.Printf("  商品 %d: SKU=%s, 数量=%v\n", count, sku, qty)
	}

	fmt.Printf("总计订单项: %d\n", count)
}

// iterateReportsExample 演示报告迭代器
func iterateReportsExample(ctx context.Context, baseClient *spapi.Client) {
	reportsClient := reports_v2021_06_30.NewClient(baseClient)

	query := map[string]string{
		"reportTypes":    "GET_FLAT_FILE_ALL_ORDERS_DATA_BY_ORDER_DATE",
		"marketplaceIds": "ATVPDKIKX0DER",
	}

	count := 0
	for report, err := range reportsClient.IterateReports(ctx, query) {
		if err != nil {
			log.Printf("迭代错误: %v", err)
			break
		}

		count++
		reportID := report["reportId"]
		status := report["processingStatus"]
		fmt.Printf("  报告 %d: %s - %s\n", count, reportID, status)

		if count >= 5 {
			fmt.Println("  (仅显示前 5 个报告)")
			break
		}
	}

	fmt.Printf("总计处理报告: %d\n", count)
}

// iterateCatalogItemsExample 演示商品目录迭代器
func iterateCatalogItemsExample(ctx context.Context, baseClient *spapi.Client) {
	catalogClient := catalog_items_v2022_04_01.NewClient(baseClient)

	query := map[string]string{
		"keywords":       "laptop",
		"marketplaceIds": "ATVPDKIKX0DER",
	}

	count := 0
	for item, err := range catalogClient.IterateCatalogItems(ctx, query) {
		if err != nil {
			log.Printf("迭代错误: %v", err)
			break
		}

		count++
		asin := item["asin"]
		fmt.Printf("  商品 %d: ASIN=%s\n", count, asin)

		if count >= 20 {
			fmt.Println("  (仅显示前 20 个商品)")
			break
		}
	}

	fmt.Printf("总计商品: %d\n", count)
}

// earlyExitExample 演示提前退出迭代
func earlyExitExample(ctx context.Context, baseClient *spapi.Client) {
	ordersClient := orders_v0.NewClient(baseClient)

	query := map[string]string{
		"MarketplaceIds": "ATVPDKIKX0DER",
		"CreatedAfter":   "2025-01-01T00:00:00Z",
	}

	// 查找特定订单然后退出
	targetOrderID := "target-order-id"

	for order, err := range ordersClient.IterateOrders(ctx, query) {
		if err != nil {
			log.Printf("错误: %v", err)
			break
		}

		orderID := order["AmazonOrderId"]
		if orderID == targetOrderID {
			fmt.Printf("找到目标订单: %s\n", orderID)
			break // 提前退出，不再继续迭代
		}
	}
}

// concurrentProcessingExample 演示并发处理订单
func concurrentProcessingExample(ctx context.Context, baseClient *spapi.Client) {
	ordersClient := orders_v0.NewClient(baseClient)

	query := map[string]string{
		"MarketplaceIds": "ATVPDKIKX0DER",
		"CreatedAfter":   "2025-01-01T00:00:00Z",
	}

	// 使用 channel 收集订单
	ordersChan := make(chan map[string]interface{}, 100)

	// Go 1.25: 在循环中启动 goroutine 不再需要 item := item
	go func() {
		defer close(ordersChan)

		for order, err := range ordersClient.IterateOrders(ctx, query) {
			if err != nil {
				log.Printf("错误: %v", err)
				return
			}
			ordersChan <- order
		}
	}()

	// 并发处理订单
	count := 0
	for order := range ordersChan {
		count++

		// 启动 goroutine 处理订单（Go 1.25 自动正确捕获变量）
		go func() {
			orderID := order["AmazonOrderId"]
			fmt.Printf("  并发处理订单: %s\n", orderID)
			// 处理订单的业务逻辑
		}()
	}

	fmt.Printf("总计提交 %d 个订单处理任务\n", count)
}
