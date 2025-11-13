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

// Feed 批量上传工具
//
// 这是一个生产级的 Feed 上传工具，展示如何：
// 1. 创建 Feed 文档上传目标
// 2. 上传大文件（处理 100MB+ 文件）
// 3. 创建 Feed
// 4. 监控 Feed 处理状态
// 5. 处理 Feed 结果
//
// 适用场景：
// - 批量更新库存
// - 批量更新价格
// - 批量创建 Listing
// - 批量更新订单状态
package main

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi/feeds-v2021-06-30"
)

func main() {
	log.Println("=== Amazon SP-API Feed Uploader ===")

	// 1. 创建客户端
	client, err := spapi.NewClient(
		spapi.WithRegion(spapi.RegionNA),
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

	feedsClient := feeds_v2021_06_30.NewClient(client)
	ctx := context.Background()

	// 2. 准备 Feed 数据（库存更新示例）
	feedContent := generateInventoryFeed()
	log.Printf("Feed size: %d bytes", len(feedContent))

	// 3. 上传 Feed
	feedID, err := uploadFeed(ctx, feedsClient, feedContent, "POST_INVENTORY_AVAILABILITY_DATA")
	if err != nil {
		log.Fatalf("Upload failed: %v", err)
	}
	log.Printf("Feed uploaded: %s", feedID)

	// 4. 监控 Feed 处理状态
	log.Println("Monitoring feed processing...")
	if err := monitorFeedProcessing(ctx, feedsClient, feedID); err != nil {
		log.Fatalf("Monitoring failed: %v", err)
	}

	log.Println("Feed processed successfully!")
}

// uploadFeed 上传 Feed 的完整流程
func uploadFeed(ctx context.Context, client *feeds_v2021_06_30.Client, content []byte, feedType string) (string, error) {
	// 步骤 1: 创建 Feed 文档上传目标
	log.Println("Step 1: Creating feed document upload destination...")

	docResult, err := client.CreateFeedDocument(ctx, map[string]interface{}{
		"contentType": "text/tab-separated-values; charset=UTF-8",
	})
	if err != nil {
		return "", fmt.Errorf("create feed document: %w", err)
	}

	docResp := docResult.(map[string]interface{})
	feedDocumentID := docResp["feedDocumentId"].(string)
	uploadURL := docResp["url"].(string)

	log.Printf("  Feed document ID: %s", feedDocumentID)

	// 步骤 2: 上传文件到 S3
	log.Println("Step 2: Uploading feed content to S3...")

	req, _ := http.NewRequestWithContext(ctx, "PUT", uploadURL, bytes.NewReader(content))
	req.Header.Set("Content-Type", "text/tab-separated-values; charset=UTF-8")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("upload to S3: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("S3 upload failed: %d - %s", resp.StatusCode, body)
	}

	log.Println("  Upload successful")

	// 步骤 3: 创建 Feed
	log.Println("Step 3: Creating feed...")

	feedResult, err := client.CreateFeed(ctx, map[string]interface{}{
		"feedType":            feedType,
		"marketplaceIds":      []string{"ATVPDKIKX0DER"},
		"inputFeedDocumentId": feedDocumentID,
	})
	if err != nil {
		return "", fmt.Errorf("create feed: %w", err)
	}

	feedResp := feedResult.(map[string]interface{})
	feedID := feedResp["feedId"].(string)

	log.Printf("  Feed created: %s", feedID)

	return feedID, nil
}

// monitorFeedProcessing 监控 Feed 处理状态
func monitorFeedProcessing(ctx context.Context, client *feeds_v2021_06_30.Client, feedID string) error {
	maxAttempts := 60
	interval := 10 * time.Second

	for attempt := range maxAttempts {
		result, err := client.GetFeed(ctx, feedID, nil)
		if err != nil {
			return err
		}

		feed := result.(map[string]interface{})
		status := feed["processingStatus"].(string)

		log.Printf("  Attempt %d/%d: Status=%s", attempt+1, maxAttempts, status)

		switch status {
		case "DONE":
			// 处理完成，获取结果
			if resultDocID, ok := feed["resultFeedDocumentId"].(string); ok {
				return processFeedResult(ctx, client, resultDocID)
			}
			return nil

		case "FATAL", "CANCELLED":
			return fmt.Errorf("feed processing failed: %s", status)

		case "IN_QUEUE", "IN_PROGRESS":
			time.Sleep(interval)

		default:
			return fmt.Errorf("unknown status: %s", status)
		}
	}

	return fmt.Errorf("timeout waiting for feed processing")
}

// processFeedResult 处理 Feed 结果
func processFeedResult(ctx context.Context, client *feeds_v2021_06_30.Client, resultDocID string) error {
	log.Println("Downloading feed result...")

	// 获取结果文档
	result, err := client.GetFeedDocument(ctx, resultDocID, nil)
	if err != nil {
		return err
	}

	doc := result.(map[string]interface{})
	url := doc["url"].(string)

	// 下载结果
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	resultContent, _ := io.ReadAll(resp.Body)

	log.Printf("Feed result:\n%s", string(resultContent))

	return nil
}

// generateInventoryFeed 生成库存更新 Feed（示例）
func generateInventoryFeed() []byte {
	// 使用 XML 格式的库存 Feed
	type Message struct {
		XMLName       xml.Name `xml:"Message"`
		MessageID     int      `xml:"MessageID"`
		OperationType string   `xml:"OperationType"`
		Inventory     struct {
			SKU                string `xml:"SKU"`
			Quantity           int    `xml:"Quantity"`
			FulfillmentLatency int    `xml:"FulfillmentLatency"`
		} `xml:"Inventory"`
	}

	type Envelope struct {
		XMLName xml.Name `xml:"AmazonEnvelope"`
		NS      string   `xml:"xmlns:xsi,attr"`
		Header  struct {
			DocumentVersion    string `xml:"DocumentVersion"`
			MerchantIdentifier string `xml:"MerchantIdentifier"`
		} `xml:"Header"`
		MessageType string    `xml:"MessageType"`
		Messages    []Message `xml:"Message"`
	}

	envelope := Envelope{
		NS:          "http://www.w3.org/2001/XMLSchema-instance",
		MessageType: "Inventory",
	}
	envelope.Header.DocumentVersion = "1.01"
	envelope.Header.MerchantIdentifier = "M_EXAMPLE_123456"

	// 添加库存更新消息
	envelope.Messages = []Message{
		{
			MessageID:     1,
			OperationType: "Update",
			Inventory: struct {
				SKU                string `xml:"SKU"`
				Quantity           int    `xml:"Quantity"`
				FulfillmentLatency int    `xml:"FulfillmentLatency"`
			}{
				SKU:                "MY-SKU-001",
				Quantity:           100,
				FulfillmentLatency: 2,
			},
		},
		{
			MessageID:     2,
			OperationType: "Update",
			Inventory: struct {
				SKU                string `xml:"SKU"`
				Quantity           int    `xml:"Quantity"`
				FulfillmentLatency int    `xml:"FulfillmentLatency"`
			}{
				SKU:                "MY-SKU-002",
				Quantity:           50,
				FulfillmentLatency: 2,
			},
		},
	}

	data, _ := xml.MarshalIndent(envelope, "", "  ")
	return append([]byte(xml.Header), data...)
}
