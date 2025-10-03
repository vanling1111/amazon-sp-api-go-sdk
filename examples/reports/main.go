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

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/vanling1111/amazon-sp-api-go-sdk/internal/models"
	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
	reports "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi/reports-v2021-06-30"
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

	// 创建 Reports API 客户端
	reportsClient := reports.NewClient(baseClient)

	ctx := context.Background()

	// 示例 1: 创建报告
	fmt.Println("=== 示例 1: 创建库存报告 ===")
	reportRequest := map[string]interface{}{
		"reportType":     "GET_MERCHANT_LISTINGS_ALL_DATA",
		"marketplaceIds": []string{"ATVPDKIKX0DER"},
	}

	reportResult, err := reportsClient.CreateReport(ctx, reportRequest)
	if err != nil {
		log.Printf("创建报告失败: %v", err)
	} else {
		jsonData, _ := json.MarshalIndent(reportResult, "", "  ")
		fmt.Printf("报告创建结果:\n%s\n\n", jsonData)
	}

	// 示例 2: 获取报告列表
	fmt.Println("=== 示例 2: 获取报告列表 ===")
	queryParams := map[string]string{
		"reportTypes":  "GET_MERCHANT_LISTINGS_ALL_DATA",
		"createdSince": time.Now().Add(-30 * 24 * time.Hour).Format(time.RFC3339),
		"pageSize":     "10",
	}

	listResult, err := reportsClient.GetReports(ctx, queryParams)
	if err != nil {
		log.Printf("获取报告列表失败: %v", err)
	} else {
		jsonData, _ := json.MarshalIndent(listResult, "", "  ")
		fmt.Printf("报告列表:\n%s\n\n", jsonData)
	}

	// 示例 3: 获取报告详情
	fmt.Println("=== 示例 3: 获取报告详情 ===")
	reportID := "12345" // 替换为实际的报告ID

	detailResult, err := reportsClient.GetReport(ctx, reportID, nil)
	if err != nil {
		log.Printf("获取报告详情失败: %v", err)
	} else {
		jsonData, _ := json.MarshalIndent(detailResult, "", "  ")
		fmt.Printf("报告详情:\n%s\n\n", jsonData)
	}

	// 示例 4: 获取报告文档
	fmt.Println("=== 示例 4: 获取报告文档 ===")
	reportDocumentID := "amzn1.tortuga.3.example" // 替换为实际的文档ID

	docResult, err := reportsClient.GetReportDocument(ctx, reportDocumentID, nil)
	if err != nil {
		log.Printf("获取报告文档失败: %v", err)
	} else {
		jsonData, _ := json.MarshalIndent(docResult, "", "  ")
		fmt.Printf("报告文档信息:\n%s\n\n", jsonData)
	}

	fmt.Println("\n✓ Reports API 示例完成")
}
