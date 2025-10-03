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
	notifications "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi/notifications-v1"
)

func main() {
	// 创建 Grantless 操作客户端
	// Grantless 操作不需要卖家的 refresh token
	// 只需要应用的 Client ID 和 Client Secret，以及相应的 scopes
	client, err := spapi.NewClient(
		spapi.WithRegion(models.RegionNA),
		spapi.WithGrantlessCredentials(
			"your-client-id",
			"your-client-secret",
			[]string{
				"sellingpartnerapi::notifications",
			},
		),
	)
	if err != nil {
		log.Fatalf("创建 Grantless 客户端失败: %v", err)
	}
	defer client.Close()

	// 创建 Notifications API 客户端
	notificationsClient := notifications.NewClient(client)

	ctx := context.Background()

	// 示例 1: 创建通知目标
	fmt.Println("=== 示例 1: 创建 SQS 通知目标 ===")
	destinationRequest := map[string]interface{}{
		"resourceSpecification": map[string]interface{}{
			"sqs": map[string]interface{}{
				"arn": "arn:aws:sqs:us-east-1:123456789012:your-queue-name",
			},
		},
		"name": "MyNotificationDestination",
	}

	result, err := notificationsClient.CreateDestination(ctx, destinationRequest)
	if err != nil {
		log.Printf("创建通知目标失败: %v", err)
	} else {
		jsonData, _ := json.MarshalIndent(result, "", "  ")
		fmt.Printf("通知目标:\n%s\n\n", jsonData)
	}

	// 示例 2: 获取通知目标列表
	fmt.Println("=== 示例 2: 获取通知目标列表 ===")
	listResult, err := notificationsClient.GetDestinations(ctx, nil)
	if err != nil {
		log.Printf("获取通知目标失败: %v", err)
	} else {
		jsonData, _ := json.MarshalIndent(listResult, "", "  ")
		fmt.Printf("通知目标列表:\n%s\n\n", jsonData)
	}

	// 示例 3: 创建订阅
	fmt.Println("=== 示例 3: 创建通知订阅 ===")
	subscriptionRequest := map[string]interface{}{
		"payloadVersion": "1.0",
		"destinationId":  "destination-id-from-step-1",
	}

	subResult, err := notificationsClient.CreateSubscription(ctx, "ANY_OFFER_CHANGED", subscriptionRequest)
	if err != nil {
		log.Printf("创建订阅失败: %v", err)
	} else {
		jsonData, _ := json.MarshalIndent(subResult, "", "  ")
		fmt.Printf("订阅结果:\n%s\n\n", jsonData)
	}

	fmt.Println("\n✓ Grantless 操作示例完成")
	fmt.Println("\n支持的 Grantless scopes:")
	fmt.Println("  - sellingpartnerapi::notifications")
	fmt.Println("  - sellingpartnerapi::migration")
}
