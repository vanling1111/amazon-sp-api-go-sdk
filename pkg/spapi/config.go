// Copyright 2025 Amazon SP-API Go SDK Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package spapi 提供 Amazon Selling Partner API 的 Go SDK。
//
// 此包是 SDK 的公开接口，提供简洁、易用的 API 访问方式。
//
// 基于官方 SP-API 文档:
//   - https://developer-docs.amazon.com/sp-api/docs/
package spapi

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/vanling1111/amazon-sp-api-go-sdk/internal/metrics"
	"github.com/vanling1111/amazon-sp-api-go-sdk/internal/models"
)

// Config 定义 SP-API 客户端的配置。
type Config struct {
	// Region 是 SP-API 的区域（如 NA, EU, FE）。
	Region models.Region `validate:"required"`

	// ClientID 是 LWA 客户端 ID。
	ClientID string `validate:"required"`

	// ClientSecret 是 LWA 客户端密钥。
	ClientSecret string `validate:"required"`

	// RefreshToken 是 LWA 刷新令牌（用于 Regular 操作）。
	// 如果使用 Grantless 操作，可以为空。
	RefreshToken string `validate:"required_without=Scopes"`

	// Scopes 是 Grantless 操作所需的权限范围。
	// 如果为空，则使用 RefreshToken 进行 Regular 操作。
	Scopes []string `validate:"required_without=RefreshToken,dive,required"`

	// SellerID 是卖家 ID（可选）。
	// 用于速率限制的多维度管理。如果未设置，将使用 ClientID。
	SellerID string

	// HTTPTimeout 是 HTTP 请求超时时间。
	HTTPTimeout time.Duration `validate:"min=1s,max=5m"`

	// MaxRetries 是请求失败时的最大重试次数。
	MaxRetries int `validate:"min=0,max=10"`

	// RateLimitBuffer 是速率限制的缓冲比例（0.0-1.0）。
	// 例如 0.1 表示保留 10% 的速率限制作为缓冲。
	RateLimitBuffer float64 `validate:"min=0,max=1"`

	// Debug 启用调试模式（详细日志）。
	Debug bool

	// MetricsRecorder 是可选的指标记录器（如 Prometheus）。
	MetricsRecorder metrics.Recorder `validate:"-"`

	// Logger 是可选的日志器（如 Zap）。
	// 如果为 nil，使用默认的 NopLogger（不输出日志）。
	Logger interface {
		Debug(msg string, fields ...interface{})
		Info(msg string, fields ...interface{})
		Warn(msg string, fields ...interface{})
		Error(msg string, fields ...interface{})
	} `validate:"-"`
}

// validate 全局验证器实例
var validate = validator.New()

// DefaultConfig 返回默认配置。
//
// 默认配置：
//   - HTTPTimeout: 30s
//   - MaxRetries: 3
//   - RateLimitBuffer: 0.1 (10%)
//   - Debug: false
func DefaultConfig() *Config {
	return &Config{
		HTTPTimeout:     30 * time.Second,
		MaxRetries:      3,
		RateLimitBuffer: 0.1,
		Debug:           false,
		MetricsRecorder: metrics.DefaultRecorder,
	}
}

// Validate 验证配置的有效性。
//
// 使用 validator 库进行声明式验证，支持丰富的验证规则。
//
// 返回值:
//   - error: 如果配置无效，返回错误
func (c *Config) Validate() error {
	// 使用 validator 进行结构体验证
	if err := validate.Struct(c); err != nil {
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			return formatValidationErrors(validationErrs)
		}
		return err
	}

	// 额外的业务逻辑验证
	if c.Region.Code == "" || c.Region.Endpoint == "" || c.Region.LWAEndpoint == "" {
		return ErrInvalidRegion
	}

	return nil
}

// formatValidationErrors 格式化验证错误
func formatValidationErrors(errs validator.ValidationErrors) error {
	messages := make([]string, 0, len(errs))
	
	for _, err := range errs {
		var msg string
		field := err.Field()
		
		switch err.Tag() {
		case "required":
			msg = fmt.Sprintf("%s is required", field)
		case "required_without":
			msg = fmt.Sprintf("%s is required when %s is not provided", field, err.Param())
		case "min":
			msg = fmt.Sprintf("%s must be at least %s", field, err.Param())
		case "max":
			msg = fmt.Sprintf("%s must not exceed %s", field, err.Param())
		case "dive":
			msg = fmt.Sprintf("%s contains invalid elements", field)
		default:
			msg = fmt.Sprintf("%s validation failed: %s", field, err.Tag())
		}
		
		messages = append(messages, msg)
	}

	if len(messages) == 1 {
		return fmt.Errorf("config validation failed: %s", messages[0])
	}
	
	return fmt.Errorf("config validation failed: %v", messages)
}

// ClientOption 定义客户端配置选项函数。
type ClientOption func(*Config)

// WithRegion 设置 API 区域。
//
// 参数:
//   - region: SP-API 区域（NA, EU, FE）
//
// 示例:
//
//	client := spapi.NewClient(spapi.WithRegion(models.RegionNorthAmerica))
func WithRegion(region models.Region) ClientOption {
	return func(c *Config) {
		c.Region = region
	}
}

// WithCredentials 设置 LWA 凭证（用于 Regular 操作）。
//
// 参数:
//   - clientID: LWA 客户端 ID
//   - clientSecret: LWA 客户端密钥
//   - refreshToken: LWA 刷新令牌
//
// 示例:
//
//	client := spapi.NewClient(
//	    spapi.WithCredentials("client-id", "client-secret", "refresh-token"),
//	)
func WithCredentials(clientID, clientSecret, refreshToken string) ClientOption {
	return func(c *Config) {
		c.ClientID = clientID
		c.ClientSecret = clientSecret
		c.RefreshToken = refreshToken
	}
}

// WithGrantlessCredentials 设置 Grantless 凭证。
//
// 参数:
//   - clientID: LWA 客户端 ID
//   - clientSecret: LWA 客户端密钥
//   - scopes: 权限范围（如 "sellingpartnerapi::notifications"）
//
// 示例:
//
//	client := spapi.NewClient(
//	    spapi.WithGrantlessCredentials("client-id", "client-secret", []string{
//	        "sellingpartnerapi::notifications",
//	    }),
//	)
func WithGrantlessCredentials(clientID, clientSecret string, scopes []string) ClientOption {
	return func(c *Config) {
		c.ClientID = clientID
		c.ClientSecret = clientSecret
		c.Scopes = scopes
	}
}

// WithSellerID 设置卖家 ID（可选）。
//
// Seller ID 用于速率限制的多维度管理。
// 如果不设置，将使用 ClientID 作为标识。
//
// 参数:
//   - sellerID: 卖家 ID
//
// 示例:
//
//	client := spapi.NewClient(
//	    spapi.WithRegion(models.RegionNA),
//	    spapi.WithCredentials("client-id", "client-secret", "refresh-token"),
//	    spapi.WithSellerID("AXXXXXXXXXXXXX"),
//	)
func WithSellerID(sellerID string) ClientOption {
	return func(c *Config) {
		c.SellerID = sellerID
	}
}

// WithHTTPTimeout 设置 HTTP 请求超时时间。
//
// 参数:
//   - timeout: 超时时间
//
// 示例:
//
//	client := spapi.NewClient(spapi.WithHTTPTimeout(60 * time.Second))
func WithHTTPTimeout(timeout time.Duration) ClientOption {
	return func(c *Config) {
		c.HTTPTimeout = timeout
	}
}

// WithMaxRetries 设置最大重试次数。
//
// 参数:
//   - maxRetries: 最大重试次数
//
// 示例:
//
//	client := spapi.NewClient(spapi.WithMaxRetries(5))
func WithMaxRetries(maxRetries int) ClientOption {
	return func(c *Config) {
		c.MaxRetries = maxRetries
	}
}

// WithRateLimitBuffer 设置速率限制缓冲比例。
//
// 参数:
//   - buffer: 缓冲比例（0.0-1.0）
//
// 示例:
//
//	client := spapi.NewClient(spapi.WithRateLimitBuffer(0.2)) // 20% 缓冲
func WithRateLimitBuffer(buffer float64) ClientOption {
	return func(c *Config) {
		c.RateLimitBuffer = buffer
	}
}

// WithDebug 启用调试模式。
//
// 示例:
//
//	client := spapi.NewClient(spapi.WithDebug())
func WithDebug() ClientOption {
	return func(c *Config) {
		c.Debug = true
	}
}

// WithMetrics 设置指标记录器。
//
// 参数:
//   - recorder: 指标记录器实现（如 Prometheus）
//
// 示例:
//
//	promRecorder := NewPrometheusRecorder()
//	client := spapi.NewClient(spapi.WithMetrics(promRecorder))
func WithMetrics(recorder metrics.Recorder) ClientOption {
	return func(c *Config) {
		if recorder != nil {
			c.MetricsRecorder = recorder
		}
	}
}

