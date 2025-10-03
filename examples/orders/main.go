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

	// 创建 Orders API 客户端
	ordersClient := orders.NewClient(baseClient)

	ctx := context.Background()

	// 示例 1: 获取订单列表
	fmt.Println("=== 示例 1: 获取最近的订单 ===")
	queryParams := map[string]string{
		"MarketplaceIds":  "ATVPDKIKX0DER", // US marketplace
		"CreatedAfter":    time.Now().Add(-7 * 24 * time.Hour).Format(time.RFC3339),
		"MaxResultsPerPage": "10",
	}

	result, err := ordersClient.GetOrders(ctx, queryParams)
	if err != nil {
		log.Printf("获取订单失败: %v", err)
	} else {
		jsonData, _ := json.MarshalIndent(result, "", "  ")
		fmt.Printf("订单列表:\n%s\n\n", jsonData)
	}

	// 示例 2: 获取单个订单详情
	fmt.Println("=== 示例 2: 获取订单详情 ===")
	orderID := "123-1234567-1234567" // 替换为实际的订单ID

	orderResult, err := ordersClient.GetOrder(ctx, orderID, nil)
	if err != nil {
		log.Printf("获取订单详情失败: %v", err)
	} else {
		jsonData, _ := json.MarshalIndent(orderResult, "", "  ")
		fmt.Printf("订单详情:\n%s\n\n", jsonData)
	}

	// 示例 3: 获取订单商品
	fmt.Println("=== 示例 3: 获取订单商品 ===")
	itemsResult, err := ordersClient.GetOrderItems(ctx, orderID, nil)
	if err != nil {
		log.Printf("获取订单商品失败: %v", err)
	} else {
		jsonData, _ := json.MarshalIndent(itemsResult, "", "  ")
		fmt.Printf("订单商品:\n%s\n\n", jsonData)
	}

	// 示例 4: 更新发货状态
	fmt.Println("=== 示例 4: 更新发货状态 ===")
	shipmentRequest := map[string]interface{}{
		"marketplaceId": "ATVPDKIKX0DER",
		"shipmentStatus": "Shipped",
	}

	_, err = ordersClient.UpdateShipmentStatus(ctx, orderID, shipmentRequest)
	if err != nil {
		log.Printf("更新发货状态失败: %v", err)
	} else {
		fmt.Println("✓ 发货状态更新成功")
	}

	fmt.Println("\n✓ Orders API 示例完成")
}

