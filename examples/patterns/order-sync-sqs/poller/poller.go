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

// Package poller 提供 SQS 消息轮询功能。
//
// 这是一个生产级的 SQS 轮询器实现，可以直接复制到你的项目中使用。
package poller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

// EventHandler 事件处理器函数类型
type EventHandler func(ctx context.Context, event *Event) error

// Poller SQS 消息轮询器
type Poller struct {
	sqsClient    *sqs.Client
	queueURL     string
	pollInterval time.Duration
	maxMessages  int32
	waitTime     int32 // Long polling wait time (1-20 seconds)
	handlers     map[string]EventHandler
	errorHandler func(error)
}

// Config 轮询器配置
type Config struct {
	QueueURL     string        // SQS 队列 URL
	PollInterval time.Duration // 轮询间隔（默认 10 秒）
	MaxMessages  int32         // 每次获取的最大消息数（1-10，默认 10）
	WaitTime     int32         // Long polling 等待时间（1-20 秒，默认 20）
}

// NewPoller 创建 SQS 轮询器
func NewPoller(sqsClient *sqs.Client, config *Config) *Poller {
	if config.PollInterval == 0 {
		config.PollInterval = 10 * time.Second
	}
	if config.MaxMessages == 0 {
		config.MaxMessages = 10
	}
	if config.WaitTime == 0 {
		config.WaitTime = 20 // Long polling
	}

	return &Poller{
		sqsClient:    sqsClient,
		queueURL:     config.QueueURL,
		pollInterval: config.PollInterval,
		maxMessages:  config.MaxMessages,
		waitTime:     config.WaitTime,
		handlers:     make(map[string]EventHandler),
	}
}

// RegisterHandler 注册事件处理器
//
// 参数:
//   notificationType: 事件类型（如 "ORDER_CHANGE", "FEED_PROCESSING_FINISHED"）
//   handler: 处理函数
func (p *Poller) RegisterHandler(notificationType string, handler EventHandler) {
	p.handlers[notificationType] = handler
}

// OnError 注册错误处理器
func (p *Poller) OnError(handler func(error)) {
	p.errorHandler = handler
}

// Start 开始轮询
//
// 这个方法会阻塞，直到 context 被取消
func (p *Poller) Start(ctx context.Context) error {
	log.Printf("[Poller] Starting SQS poller for queue: %s", p.queueURL)
	log.Printf("[Poller] Poll interval: %v, Max messages: %d", p.pollInterval, p.maxMessages)

	ticker := time.NewTicker(p.pollInterval)
	defer ticker.Stop()

	// 立即执行一次
	p.poll(ctx)

	// 定时轮询
	for {
		select {
		case <-ctx.Done():
			log.Println("[Poller] Stopping poller...")
			return ctx.Err()
		case <-ticker.C:
			p.poll(ctx)
		}
	}
}

// poll 执行一次轮询
func (p *Poller) poll(ctx context.Context) {
	// 从 SQS 接收消息（Long Polling）
	result, err := p.sqsClient.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(p.queueURL),
		MaxNumberOfMessages: p.maxMessages,
		WaitTimeSeconds:     p.waitTime, // Long polling
		MessageAttributeNames: []string{
			string(types.QueueAttributeNameAll),
		},
	})

	if err != nil {
		p.handleError(fmt.Errorf("failed to receive messages: %w", err))
		return
	}

	if len(result.Messages) == 0 {
		// 没有消息
		return
	}

	log.Printf("[Poller] Received %d messages", len(result.Messages))

	// 处理每条消息
	for _, msg := range result.Messages {
		if err := p.processMessage(ctx, msg); err != nil {
			p.handleError(fmt.Errorf("failed to process message: %w", err))
			continue
		}

		// 删除已处理的消息
		if err := p.deleteMessage(ctx, msg.ReceiptHandle); err != nil {
			p.handleError(fmt.Errorf("failed to delete message: %w", err))
		}
	}
}

// processMessage 处理单条消息
func (p *Poller) processMessage(ctx context.Context, msg types.Message) error {
	if msg.Body == nil {
		return fmt.Errorf("message body is nil")
	}

	// 解析消息
	event, err := ParseSQSMessage(*msg.Body)
	if err != nil {
		return fmt.Errorf("failed to parse message: %w", err)
	}

	// 查找对应的处理器
	handler, ok := p.handlers[event.NotificationType]
	if !ok {
		log.Printf("[Poller] No handler for notification type: %s", event.NotificationType)
		return nil // 不是错误，只是没有处理器
	}

	// 执行处理器
	log.Printf("[Poller] Processing %s event: %s", event.NotificationType, event.NotificationID)
	if err := handler(ctx, event); err != nil {
		return fmt.Errorf("handler failed: %w", err)
	}

	log.Printf("[Poller] Successfully processed event: %s", event.NotificationID)
	return nil
}

// deleteMessage 删除 SQS 消息
func (p *Poller) deleteMessage(ctx context.Context, receiptHandle *string) error {
	_, err := p.sqsClient.DeleteMessage(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(p.queueURL),
		ReceiptHandle: receiptHandle,
	})
	return err
}

// handleError 处理错误
func (p *Poller) handleError(err error) {
	if p.errorHandler != nil {
		p.errorHandler(err)
	} else {
		log.Printf("[Poller] ERROR: %v", err)
	}
}

// Event Amazon SP-API 通知事件
type Event struct {
	NotificationVersion string          `json:"NotificationVersion"`
	NotificationType    string          `json:"NotificationType"`
	PayloadVersion      string          `json:"PayloadVersion"`
	EventTime           string          `json:"EventTime"`
	NotificationID      string          `json:"NotificationMetadata.NotificationId"`
	Payload             json.RawMessage `json:"Payload"`
}

// ParseSQSMessage 解析 SQS 消息为 SP-API 事件
func ParseSQSMessage(body string) (*Event, error) {
	// SQS 消息格式:
	// {
	//   "Message": "{...}",  // SNS 格式
	//   或直接是 EventBridge 格式
	// }

	var sqsMsg struct {
		Message string `json:"Message"`
	}

	// 尝试解析为 SNS 格式
	if err := json.Unmarshal([]byte(body), &sqsMsg); err == nil && sqsMsg.Message != "" {
		// SNS 包装格式，解包
		body = sqsMsg.Message
	}

	// 解析为 SP-API 事件
	var event Event
	if err := json.Unmarshal([]byte(body), &event); err != nil {
		return nil, fmt.Errorf("failed to unmarshal event: %w", err)
	}

	return &event, nil
}

// ParsePayload 解析事件负载为指定类型
func (e *Event) ParsePayload(v interface{}) error {
	return json.Unmarshal(e.Payload, v)
}
