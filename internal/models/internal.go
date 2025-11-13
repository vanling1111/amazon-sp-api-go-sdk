// Package models 提供内部通用模型和数据结构。
//
// 此包包含在多个内部模块之间共享的通用数据结构，
// 不对外暴露。公开的 API 模型定义在 pkg/spapi 目录下。
//
// 注意：Region 和 Marketplace 已移至 pkg/spapi 包，作为公开 API。
package models

import (
	"time"
)

// RateLimitInfo 表示速率限制信息。
//
// 用于内部速率限制管理，不对外暴露。
type RateLimitInfo struct {
	// Rate 是允许的请求速率（请求数/时间单位）
	Rate float64

	// Burst 是允许的突发请求数
	Burst int

	// RefillInterval 是令牌桶的补充间隔
	RefillInterval time.Duration
}

// RequestMetadata 表示请求的元数据。
//
// 用于内部请求追踪和日志记录。
type RequestMetadata struct {
	// RequestID 是请求的唯一标识符
	RequestID string

	// Timestamp 是请求时间戳
	Timestamp time.Time

	// Marketplace 是请求的市场
	Marketplace string

	// Endpoint 是请求的端点
	Endpoint string

	// Method 是 HTTP 方法
	Method string
}

// ErrorDetail 表示错误详情。
//
// 用于内部错误处理和日志记录。
type ErrorDetail struct {
	// Code 是错误代码
	Code string

	// Message 是错误消息
	Message string

	// Details 是额外的错误详情
	Details map[string]interface{}
}
