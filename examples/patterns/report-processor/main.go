// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

// 报告自动处理器
//
// 这是一个生产级的报告处理工具，展示如何：
// 1. 定期创建报告
// 2. 监控报告生成状态
// 3. 自动下载和解密报告
// 4. 解析报告数据
// 5. 存储到数据库或数据仓库
//
// 适用场景：
// - 每日订单数据同步
// - 财务对账
// - 库存分析
// - 销售数据分析
package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/vanling1111/amazon-sp-api-go-sdk/internal/models"
	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi/reports-v2021-06-30"
)

func main() {
	log.Println("=== Amazon SP-API Report Processor ===")

	// 创建客户端
	client, err := spapi.NewClient(
		spapi.WithRegion(models.RegionNA),
		spapi.WithCredentials(
			os.Getenv("SP_API_CLIENT_ID"),
			os.Getenv("SP_API_CLIENT_SECRET"),
			os.Getenv("SP_API_REFRESH_TOKEN"),
		),
	)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	reportsClient := reports_v2021_06_30.NewClient(client)
	ctx := context.Background()

	// 示例 1: 处理订单报告
	log.Println("\n=== Processing Order Report ===")
	if err := processOrderReport(ctx, reportsClient); err != nil {
		log.Printf("Order report failed: %v", err)
	}

	// 示例 2: 处理财务报告
	log.Println("\n=== Processing Finance Report ===")
	if err := processFinanceReport(ctx, reportsClient); err != nil {
		log.Printf("Finance report failed: %v", err)
	}

	// 示例 3: 定时处理所有报告类型
	log.Println("\n=== Starting Scheduled Report Processing ===")
	startScheduledProcessing(ctx, reportsClient)
}

// processOrderReport 处理订单报告
func processOrderReport(ctx context.Context, client *reports_v2021_06_30.Client) error {
	// 1. 创建报告
	log.Println("Creating order report...")
	reportID, err := createReport(ctx, client, "GET_FLAT_FILE_ALL_ORDERS_DATA_BY_ORDER_DATE", 30)
	if err != nil {
		return err
	}

	// 2. 等待生成
	log.Println("Waiting for report generation...")
	reportDocID, err := waitForReport(ctx, client, reportID)
	if err != nil {
		return err
	}

	// 3. 下载并解密（一行代码！）
	log.Println("Downloading and decrypting...")
	decrypted, err := client.GetReportDocumentDecrypted(ctx, reportDocID)
	if err != nil {
		return err
	}

	// 4. 解析数据
	log.Println("Parsing report data...")
	orders, err := parseOrderReport(decrypted)
	if err != nil {
		return err
	}

	log.Printf("Processed %d orders", len(orders))

	// 5. 存储数据
	if err := saveToDatabase(orders); err != nil {
		return err
	}

	log.Println("Order report processed successfully!")
	return nil
}

// processFinanceReport 处理财务报告
func processFinanceReport(ctx context.Context, client *reports_v2021_06_30.Client) error {
	log.Println("Creating finance report...")
	reportID, err := createReport(ctx, client, "GET_V2_SETTLEMENT_REPORT_DATA_FLAT_FILE", 7)
	if err != nil {
		return err
	}

	reportDocID, err := waitForReport(ctx, client, reportID)
	if err != nil {
		return err
	}

	// 自动解密
	decrypted, err := client.GetReportDocumentDecrypted(ctx, reportDocID)
	if err != nil {
		return err
	}

	// 解析财务数据
	log.Printf("Finance report size: %d bytes", len(decrypted))
	// TODO: 解析财务数据并保存

	return nil
}

// createReport 创建报告
func createReport(ctx context.Context, client *reports_v2021_06_30.Client, reportType string, daysBack int) (string, error) {
	endTime := time.Now()
	startTime := endTime.Add(-time.Duration(daysBack) * 24 * time.Hour)

	result, err := client.CreateReport(ctx, map[string]interface{}{
		"reportType":     reportType,
		"marketplaceIds": []string{"ATVPDKIKX0DER"},
		"dataStartTime":  startTime.Format(time.RFC3339),
		"dataEndTime":    endTime.Format(time.RFC3339),
	})
	if err != nil {
		return "", err
	}

	resp := result.(map[string]interface{})
	return resp["reportId"].(string), nil
}

// waitForReport 等待报告生成
func waitForReport(ctx context.Context, client *reports_v2021_06_30.Client, reportID string) (string, error) {
	for attempt := range 60 {
		result, err := client.GetReport(ctx, reportID, nil)
		if err != nil {
			return "", err
		}

		report := result.(map[string]interface{})
		status := report["processingStatus"].(string)

		log.Printf("  Status: %s (attempt %d)", status, attempt+1)

		if status == "DONE" {
			return report["reportDocumentId"].(string), nil
		} else if status == "FATAL" || status == "CANCELLED" {
			return "", fmt.Errorf("report failed: %s", status)
		}

		time.Sleep(10 * time.Second)
	}

	return "", fmt.Errorf("timeout")
}

// parseOrderReport 解析订单报告（TSV 格式）
func parseOrderReport(data []byte) ([]map[string]string, error) {
	reader := csv.NewReader(strings.NewReader(string(data)))
	reader.Comma = '\t'
	reader.LazyQuotes = true

	// 读取表头
	headers, err := reader.Read()
	if err != nil {
		return nil, err
	}

	// 读取数据行
	var orders []map[string]string
	for {
		row, err := reader.Read()
		if err != nil {
			break
		}

		order := make(map[string]string)
		for i, value := range row {
			if i < len(headers) {
				order[headers[i]] = value
			}
		}
		orders = append(orders, order)
	}

	return orders, nil
}

// saveToDatabase 保存到数据库
func saveToDatabase(orders []map[string]string) error {
	// TODO: 实现数据库保存逻辑
	log.Printf("Would save %d orders to database", len(orders))
	return nil
}

// startScheduledProcessing 定时处理报告
func startScheduledProcessing(ctx context.Context, client *reports_v2021_06_30.Client) {
	ticker := time.NewTicker(24 * time.Hour) // 每天执行一次
	defer ticker.Stop()

	reportTypes := []string{
		"GET_FLAT_FILE_ALL_ORDERS_DATA_BY_ORDER_DATE",
		"GET_V2_SETTLEMENT_REPORT_DATA_FLAT_FILE",
		"GET_FBA_INVENTORY_AGED_DATA",
	}

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			log.Println("Starting scheduled report processing...")

			for _, reportType := range reportTypes {
				log.Printf("Processing %s...", reportType)

				reportID, err := createReport(ctx, client, reportType, 1)
				if err != nil {
					log.Printf("Create failed: %v", err)
					continue
				}

				reportDocID, err := waitForReport(ctx, client, reportID)
				if err != nil {
					log.Printf("Wait failed: %v", err)
					continue
				}

				decrypted, err := client.GetReportDocumentDecrypted(ctx, reportDocID)
				if err != nil {
					log.Printf("Decrypt failed: %v", err)
					continue
				}

				log.Printf("Processed %s (%d bytes)", reportType, len(decrypted))
			}
		}
	}
}
