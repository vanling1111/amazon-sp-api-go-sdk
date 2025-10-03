// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/vanling1111/amazon-sp-api-go-sdk/internal/models"
	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
	orders "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi/orders-v0"
	pricing "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi/product-pricing-v2022-05-01"
)

func main() {
	// 高级用法示例：
	// 1. 自定义配置
	// 2. 多个 API 组合使用
	// 3. 错误处理
	// 4. 速率限制管理

	// 创建客户端（带高级配置）
	client, err := spapi.NewClient(
		spapi.WithRegion(models.RegionNA),
		spapi.WithCredentials(
			"your-client-id",
			"your-client-secret",
			"your-refresh-token",
		),
		spapi.WithHTTPTimeout(60*time.Second), // 更长的超时时间
		spapi.WithMaxRetries(5),               // 更多重试次数
		spapi.WithDebug(),                     // 启用调试模式
	)
	if err != nil {
		log.Fatalf("创建客户端失败: %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	// 示例 1: 组合使用多个 API
	fmt.Println("=== 示例 1: 获取订单并查询价格 ===")

	// 创建多个 API 客户端
	ordersClient := orders.NewClient(client)
	pricingClient := pricing.NewClient(client)

	// 获取订单
	orderParams := map[string]string{
		"MarketplaceIds": "ATVPDKIKX0DER",
		"CreatedAfter":   time.Now().Add(-24 * time.Hour).Format(time.RFC3339),
	}

	_, err = ordersClient.GetOrders(ctx, orderParams)
	if err != nil {
		log.Printf("获取订单失败: %v", err)
	} else {
		fmt.Println("✓ 获取订单成功")

		// 假设我们获取到订单中的 SKU，查询价格
		sku := "MY-SKU-001"
		pricingParams := map[string]string{
			"marketplaceId": "ATVPDKIKX0DER",
			"itemType":      "Sku",
		}

		priceResult, err := pricingClient.GetCompetitiveSummary(ctx, pricingParams)
		if err != nil {
			log.Printf("获取价格失败: %v", err)
		} else {
			fmt.Printf("✓ SKU %s 的价格信息获取成功\n", sku)
			jsonData, _ := json.MarshalIndent(priceResult, "", "  ")
			fmt.Printf("%s\n", jsonData)
		}
	}

	// 示例 2: 速率限制管理
	fmt.Println("\n=== 示例 2: 速率限制监控 ===")
	rateLimitMgr := client.RateLimitManager()
	count := rateLimitMgr.Count()
	fmt.Printf("活跃的速率限制器数量: %d\n", count)

	// 示例 3: 错误处理最佳实践
	fmt.Println("\n=== 示例 3: 错误处理 ===")
	_, err = ordersClient.GetOrder(ctx, "invalid-order-id", nil)
	if err != nil {
		// 打印错误信息
		fmt.Printf("获取订单失败: %v\n", err)

		// 可以在这里根据错误类型进行特殊处理
		fmt.Println("可以检查错误类型并进行相应的处理")
	}

	// 示例 4: 并发请求（使用 goroutine）
	fmt.Println("\n=== 示例 4: 并发请求 ===")

	orderIDs := []string{"111-1111111-1111111", "222-2222222-2222222", "333-3333333-3333333"}
	results := make(chan interface{}, len(orderIDs))
	errors := make(chan error, len(orderIDs))

	for _, orderID := range orderIDs {
		go func(id string) {
			result, err := ordersClient.GetOrder(ctx, id, nil)
			if err != nil {
				errors <- err
			} else {
				results <- result
			}
		}(orderID)
	}

	// 收集结果
	successCount := 0
	errorCount := 0
	for i := 0; i < len(orderIDs); i++ {
		select {
		case <-results:
			successCount++
		case <-errors:
			errorCount++
		case <-time.After(10 * time.Second):
			fmt.Println("⚠ 请求超时")
			break
		}
	}

	fmt.Printf("并发请求完成: 成功 %d, 失败 %d\n", successCount, errorCount)

	fmt.Println("\n✓ 高级用法示例完成")
}
