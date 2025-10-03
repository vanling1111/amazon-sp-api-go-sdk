// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/vanling1111/amazon-sp-api-go-sdk/internal/models"
	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
	listings "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi/listings-items-v2021-08-01"
)

func main() {
	// 创建基础客户端
	baseClient, err := spapi.NewClient(
		spapi.WithRegion(models.RegionNA),
		spapi.WithCredentials(
			"your-client-id",
			"your-client-secret",
			"your-refresh-token",
		),
	)
	if err != nil {
		log.Fatalf("创建客户端失败: %v", err)
	}
	defer baseClient.Close()

	// 创建 Listings Items API 客户端
	listingsClient := listings.NewClient(baseClient)

	ctx := context.Background()

	// 示例 1: 搜索 Listings
	fmt.Println("=== 示例 1: 搜索 Listings ===")
	sellerId := "A1234567890123" // 替换为实际的 Seller ID
	queryParams := map[string]string{
		"marketplaceIds": "ATVPDKIKX0DER",
		"pageSize":       "10",
	}

	searchResult, err := listingsClient.SearchListingsItems(ctx, sellerId, queryParams)
	if err != nil {
		log.Printf("搜索商品失败: %v", err)
	} else {
		jsonData, _ := json.MarshalIndent(searchResult, "", "  ")
		fmt.Printf("搜索结果:\n%s\n\n", jsonData)
	}

	// 示例 2: 获取 Listing 详情
	fmt.Println("=== 示例 2: 获取 Listing 详情 ===")
	sku := "MY-SKU-001"

	getParams := map[string]string{
		"marketplaceIds": "ATVPDKIKX0DER",
		"includedData":   "summaries,attributes,issues",
	}

	itemResult, err := listingsClient.GetListingsItem(ctx, sellerId, sku, getParams)
	if err != nil {
		log.Printf("获取 Listing 失败: %v", err)
	} else {
		jsonData, _ := json.MarshalIndent(itemResult, "", "  ")
		fmt.Printf("Listing 详情:\n%s\n\n", jsonData)
	}

	// 示例 3: 更新 Listing（PATCH）
	fmt.Println("=== 示例 3: 更新 Listing ===")
	patchRequest := map[string]interface{}{
		"productType": "PRODUCT",
		"patches": []map[string]interface{}{
			{
				"op":   "replace",
				"path": "/attributes/fulfillment_availability",
				"value": []map[string]interface{}{
					{
						"fulfillment_channel_code": "DEFAULT",
						"quantity":                 100,
					},
				},
			},
		},
	}

	patchResult, err := listingsClient.PatchListingsItem(ctx, sellerId, sku, patchRequest)
	if err != nil {
		log.Printf("更新 Listing 失败: %v", err)
	} else {
		jsonData, _ := json.MarshalIndent(patchResult, "", "  ")
		fmt.Printf("更新结果:\n%s\n\n", jsonData)
	}

	fmt.Println("\n✓ Listings API 示例完成")
}
