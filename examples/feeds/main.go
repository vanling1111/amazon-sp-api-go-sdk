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
	feeds "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi/feeds-v2021-06-30"
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

	// 创建 Feeds API 客户端
	feedsClient := feeds.NewClient(baseClient)

	ctx := context.Background()

	// 示例 1: 创建 Feed 文档
	fmt.Println("=== 示例 1: 创建 Feed 文档 ===")
	docRequest := map[string]interface{}{
		"contentType": "text/xml; charset=UTF-8",
	}

	docResult, err := feedsClient.CreateFeedDocument(ctx, docRequest)
	if err != nil {
		log.Printf("创建 Feed 文档失败: %v", err)
	} else {
		jsonData, _ := json.MarshalIndent(docResult, "", "  ")
		fmt.Printf("Feed 文档:\n%s\n\n", jsonData)
	}

	// 示例 2: 创建 Feed
	fmt.Println("=== 示例 2: 创建 Feed ===")
	feedRequest := map[string]interface{}{
		"feedType":            "POST_PRODUCT_DATA",
		"marketplaceIds":      []string{"ATVPDKIKX0DER"},
		"inputFeedDocumentId": "amzn1.tortuga.3.example", // 替换为实际的文档ID
	}

	feedResult, err := feedsClient.CreateFeed(ctx, feedRequest)
	if err != nil {
		log.Printf("创建 Feed 失败: %v", err)
	} else {
		jsonData, _ := json.MarshalIndent(feedResult, "", "  ")
		fmt.Printf("Feed 创建结果:\n%s\n\n", jsonData)
	}

	// 示例 3: 获取 Feed 列表
	fmt.Println("=== 示例 3: 获取 Feed 列表 ===")
	queryParams := map[string]string{
		"feedTypes":  "POST_PRODUCT_DATA",
		"maxResults": "10",
	}

	listResult, err := feedsClient.GetFeeds(ctx, queryParams)
	if err != nil {
		log.Printf("获取 Feed 列表失败: %v", err)
	} else {
		jsonData, _ := json.MarshalIndent(listResult, "", "  ")
		fmt.Printf("Feed 列表:\n%s\n\n", jsonData)
	}

	// 示例 4: 获取 Feed 详情
	fmt.Println("=== 示例 4: 获取 Feed 详情 ===")
	feedID := "12345" // 替换为实际的 Feed ID

	detailResult, err := feedsClient.GetFeed(ctx, feedID, nil)
	if err != nil {
		log.Printf("获取 Feed 详情失败: %v", err)
	} else {
		jsonData, _ := json.MarshalIndent(detailResult, "", "  ")
		fmt.Printf("Feed 详情:\n%s\n\n", jsonData)
	}

	fmt.Println("\n✓ Feeds API 示例完成")
}
