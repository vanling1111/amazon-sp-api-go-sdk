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

// Package main 演示报告自动解密功能。
//
// 此示例展示如何使用 SDK 的自动解密功能下载和处理 Amazon SP-API 加密报告。
package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
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

	reportsClient := reports_v2021_06_30.NewClient(baseClient)
	ctx := context.Background()

	// 示例 1: 创建订单报告
	fmt.Println("=== 步骤 1: 创建订单报告 ===")
	reportID, err := createOrdersReport(ctx, reportsClient)
	if err != nil {
		log.Fatalf("创建报告失败: %v", err)
	}
	fmt.Printf("报告 ID: %s\n", reportID)

	// 示例 2: 等待报告生成
	fmt.Println("\n=== 步骤 2: 等待报告生成 ===")
	reportDocumentID, err := waitForReportCompletion(ctx, reportsClient, reportID)
	if err != nil {
		log.Fatalf("等待报告失败: %v", err)
	}
	fmt.Printf("报告文档 ID: %s\n", reportDocumentID)

	// 示例 3: 自动下载并解密报告
	fmt.Println("\n=== 步骤 3: 下载并解密报告 ===")
	decryptedData, err := reportsClient.GetReportDocumentDecrypted(ctx, reportDocumentID)
	if err != nil {
		log.Fatalf("下载/解密报告失败: %v", err)
	}
	fmt.Printf("报告大小: %d bytes\n", len(decryptedData))

	// 示例 4: 解析 CSV 报告
	fmt.Println("\n=== 步骤 4: 解析 CSV 数据 ===")
	if err := parseCSVReport(decryptedData); err != nil {
		log.Fatalf("解析报告失败: %v", err)
	}

	// 示例 5: 保存报告到文件
	fmt.Println("\n=== 步骤 5: 保存报告 ===")
	filename := fmt.Sprintf("order_report_%s.csv", time.Now().Format("20060102_150405"))
	if err := os.WriteFile(filename, decryptedData, 0644); err != nil {
		log.Fatalf("保存报告失败: %v", err)
	}
	fmt.Printf("报告已保存到: %s\n", filename)
}

// createOrdersReport 创建订单报告
func createOrdersReport(ctx context.Context, client *reports_v2021_06_30.Client) (string, error) {
	request := map[string]interface{}{
		"reportType":     "GET_FLAT_FILE_ALL_ORDERS_DATA_BY_ORDER_DATE",
		"marketplaceIds": []string{"ATVPDKIKX0DER"},
		"dataStartTime":  time.Now().Add(-30 * 24 * time.Hour).Format(time.RFC3339),
		"dataEndTime":    time.Now().Format(time.RFC3339),
	}

	result, err := client.CreateReport(ctx, request)
	if err != nil {
		return "", err
	}

	// 解析 reportId
	resultMap, ok := result.(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("unexpected response format")
	}

	reportID, ok := resultMap["reportId"].(string)
	if !ok {
		return "", fmt.Errorf("reportId not found in response")
	}

	return reportID, nil
}

// waitForReportCompletion 等待报告生成完成
func waitForReportCompletion(ctx context.Context, client *reports_v2021_06_30.Client, reportID string) (string, error) {
	maxAttempts := 60 // 最多等待 10 分钟
	interval := 10 * time.Second

	for attempt := range maxAttempts {
		// 获取报告状态
		result, err := client.GetReport(ctx, reportID, nil)
		if err != nil {
			return "", err
		}

		resultMap, ok := result.(map[string]interface{})
		if !ok {
			return "", fmt.Errorf("unexpected response format")
		}

		status, _ := resultMap["processingStatus"].(string)
		fmt.Printf("  尝试 %d/%d: 状态=%s\n", attempt+1, maxAttempts, status)

		switch status {
		case "DONE":
			// 报告生成完成
			reportDocumentID, ok := resultMap["reportDocumentId"].(string)
			if !ok {
				return "", fmt.Errorf("reportDocumentId not found")
			}
			return reportDocumentID, nil

		case "FATAL", "CANCELLED":
			// 报告生成失败
			return "", fmt.Errorf("report generation failed with status: %s", status)

		case "IN_QUEUE", "IN_PROGRESS":
			// 继续等待
			time.Sleep(interval)

		default:
			return "", fmt.Errorf("unknown status: %s", status)
		}
	}

	return "", fmt.Errorf("timeout waiting for report completion")
}

// parseCSVReport 解析 CSV 格式的报告
func parseCSVReport(data []byte) error {
	reader := csv.NewReader(strings.NewReader(string(data)))
	reader.Comma = '\t' // Amazon 报告通常使用 Tab 分隔

	// 读取表头
	headers, err := reader.Read()
	if err != nil {
		return fmt.Errorf("failed to read headers: %w", err)
	}
	fmt.Printf("  列数: %d\n", len(headers))
	fmt.Printf("  列名: %v\n", headers[:min(5, len(headers))]) // 显示前 5 列

	// 读取数据行
	rowCount := 0
	for {
		row, err := reader.Read()
		if err != nil {
			break // EOF or error
		}
		rowCount++

		// 只显示前 3 行
		if rowCount <= 3 {
			fmt.Printf("  第 %d 行: %v\n", rowCount, row[:min(3, len(row))])
		}
	}

	fmt.Printf("  总计行数: %d\n", rowCount)
	return nil
}

// min 返回两个整数中的较小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
