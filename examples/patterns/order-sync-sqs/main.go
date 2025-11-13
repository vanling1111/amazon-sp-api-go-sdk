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

// Amazon SP-API 订单实时同步服务
//
// 这是一个生产级的订单同步服务示例，展示如何：
// 1. 通过 SQS 接收订单变更通知（10-30秒延迟）
// 2. 调用 Orders API 获取订单详情
// 3. 使用 Go 1.25 迭代器自动处理分页
// 4. 推送订单到 ERP 系统
//
// 特点：
// - 准实时（10-30秒延迟，这是 Amazon SP-API 的架构限制）
// - 可靠（SQS 消息持久化）
// - 低成本（不浪费 Orders API 配额）
// - 生产级错误处理和重试
//
// 使用方法：
//  1. 配置环境变量（见 config.yaml.example）
//  2. go run main.go
//  3. 或使用 Docker: docker-compose up
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/vanling1111/amazon-sp-api-go-sdk/examples/patterns/order-sync-sqs/poller"
	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi/orders-v0"
)

// OrderSyncService 订单同步服务
type OrderSyncService struct {
	spapiClient  *spapi.Client
	ordersClient *orders_v0.Client
	poller       *poller.Poller
}

func main() {
	log.Println("=== Amazon SP-API Order Sync Service ===")
	log.Println("Using SQS notifications for real-time order updates")
	log.Println("Latency: 10-30 seconds (Amazon SP-API design limitation)")
	log.Println("")

	// 1. 加载配置
	spapiConfig := loadSPAPIConfig()
	sqsQueueURL := os.Getenv("SQS_QUEUE_URL")
	if sqsQueueURL == "" {
		log.Fatal("SQS_QUEUE_URL environment variable is required")
	}

	// 2. 创建 SP-API 客户端
	spapiClient, err := spapi.NewClient(
		spapi.WithRegion(spapi.RegionNA),
		spapi.WithCredentials(
			spapiConfig.ClientID,
			spapiConfig.ClientSecret,
			spapiConfig.RefreshToken,
		),
	)
	if err != nil {
		log.Fatalf("Failed to create SP-API client: %v", err)
	}
	defer spapiClient.Close()

	// 3. 创建 AWS SQS 客户端
	ctx := context.Background()
	awsConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("Failed to load AWS config: %v", err)
	}
	sqsClient := sqs.NewFromConfig(awsConfig)

	// 4. 创建订单同步服务
	service := &OrderSyncService{
		spapiClient:  spapiClient,
		ordersClient: orders_v0.NewClient(spapiClient),
		poller: poller.NewPoller(sqsClient, &poller.Config{
			QueueURL:     sqsQueueURL,
			PollInterval: 10 * time.Second, // 每 10 秒轮询一次
			MaxMessages:  10,               // 每次最多 10 条消息
			WaitTime:     20,               // Long polling 20 秒
		}),
	}

	// 5. 注册事件处理器
	service.poller.RegisterHandler("ORDER_CHANGE", service.handleOrderChange)
	service.poller.RegisterHandler("FEED_PROCESSING_FINISHED", service.handleFeedDone)

	// 6. 注册错误处理器
	service.poller.OnError(func(err error) {
		log.Printf("[ERROR] %v", err)
		// TODO: 发送告警到监控系统
	})

	// 7. 启动轮询器（阻塞）
	log.Println("Starting SQS poller...")
	log.Println("Listening for order notifications...")
	log.Println("Press Ctrl+C to stop")

	// 处理优雅退出
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Println("\nReceived shutdown signal...")
		cancel()
	}()

	// 启动轮询
	if err := service.poller.Start(ctx); err != nil && err != context.Canceled {
		log.Fatalf("Poller error: %v", err)
	}

	log.Println("Service stopped gracefully")
}

// handleOrderChange 处理订单变更事件
func (s *OrderSyncService) handleOrderChange(ctx context.Context, event *poller.Event) error {
	log.Printf("[ORDER_CHANGE] Received notification: %s", event.NotificationID)

	// 解析订单变更负载
	var payload struct {
		OrderChangeNotification struct {
			AmazonOrderID string `json:"AmazonOrderId"`
			OrderStatus   string `json:"OrderStatus"`
			MarketplaceID string `json:"MarketplaceId"`
		} `json:"OrderChangeNotification"`
	}

	if err := event.ParsePayload(&payload); err != nil {
		return fmt.Errorf("parse payload: %w", err)
	}

	orderID := payload.OrderChangeNotification.AmazonOrderID
	log.Printf("[ORDER_CHANGE] Order ID: %s, Status: %s",
		orderID, payload.OrderChangeNotification.OrderStatus)

	// 获取完整订单详情
	order, err := s.ordersClient.GetOrder(ctx, orderID, nil)
	if err != nil {
		return fmt.Errorf("get order details: %w", err)
	}

	// 获取订单项（使用 Go 1.25 迭代器）
	items := []map[string]interface{}{}
	for item, err := range s.ordersClient.IterateOrderItems(ctx, orderID, nil) {
		if err != nil {
			return fmt.Errorf("iterate order items: %w", err)
		}
		items = append(items, item)
	}

	log.Printf("[ORDER_CHANGE] Order has %d items", len(items))

	// 推送到 ERP
	if err := s.pushToERP(order, items); err != nil {
		return fmt.Errorf("push to ERP: %w", err)
	}

	log.Printf("[ORDER_CHANGE] Successfully synced order: %s", orderID)
	return nil
}

// handleFeedDone 处理 Feed 处理完成事件
func (s *OrderSyncService) handleFeedDone(ctx context.Context, event *poller.Event) error {
	log.Printf("[FEED_DONE] Feed processing finished: %s", event.NotificationID)

	var payload struct {
		FeedID string `json:"feedId"`
		Status string `json:"processingStatus"`
	}

	if err := event.ParsePayload(&payload); err != nil {
		return fmt.Errorf("parse payload: %w", err)
	}

	log.Printf("[FEED_DONE] Feed ID: %s, Status: %s", payload.FeedID, payload.Status)

	// TODO: 处理 Feed 结果

	return nil
}

// pushToERP 推送订单到 ERP 系统
func (s *OrderSyncService) pushToERP(order interface{}, items []map[string]interface{}) error {
	// 这里实现推送到 ERP 的逻辑
	// 方式 1: HTTP POST 到 ERP 的 Webhook
	// 方式 2: 写入数据库
	// 方式 3: 发送到消息队列（Kafka/RabbitMQ）

	orderJSON, _ := json.MarshalIndent(map[string]interface{}{
		"order": order,
		"items": items,
	}, "", "  ")

	log.Printf("[ERP] Would push order to ERP:\n%s", orderJSON)

	// TODO: 实际的 ERP 推送逻辑
	// Example:
	// resp, err := http.Post("https://your-erp.com/api/orders", "application/json", bytes.NewBuffer(orderJSON))

	return nil
}

// loadSPAPIConfig 从环境变量加载 SP-API 配置
func loadSPAPIConfig() struct {
	ClientID     string
	ClientSecret string
	RefreshToken string
} {
	clientID := os.Getenv("SP_API_CLIENT_ID")
	clientSecret := os.Getenv("SP_API_CLIENT_SECRET")
	refreshToken := os.Getenv("SP_API_REFRESH_TOKEN")

	if clientID == "" || clientSecret == "" || refreshToken == "" {
		log.Fatal("Missing required environment variables: SP_API_CLIENT_ID, SP_API_CLIENT_SECRET, SP_API_REFRESH_TOKEN")
	}

	return struct {
		ClientID     string
		ClientSecret string
		RefreshToken string
	}{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RefreshToken: refreshToken,
	}
}
