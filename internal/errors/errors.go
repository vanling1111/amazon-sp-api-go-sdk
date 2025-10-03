// Package errors 提供 SP-API 错误类型和处理。
//
// 此包定义了详细的错误分类，帮助应用程序更好地处理各种错误情况。
//
// 官方文档:
//   - https://developer-docs.amazon.com/sp-api/docs/response-format
package errors

import (
	"fmt"
	"net/http"
)

// ErrorType 表示错误类型。
type ErrorType string

const (
	// ErrorTypeRateLimit 表示速率限制错误（429）
	ErrorTypeRateLimit ErrorType = "RateLimit"

	// ErrorTypeAuth 表示认证错误（401, 403）
	ErrorTypeAuth ErrorType = "Authentication"

	// ErrorTypeValidation 表示验证错误（400）
	ErrorTypeValidation ErrorType = "Validation"

	// ErrorTypeNotFound 表示资源未找到（404）
	ErrorTypeNotFound ErrorType = "NotFound"

	// ErrorTypeServer 表示服务器错误（5xx）
	ErrorTypeServer ErrorType = "Server"

	// ErrorTypeNetwork 表示网络错误
	ErrorTypeNetwork ErrorType = "Network"

	// ErrorTypeUnknown 表示未知错误
	ErrorTypeUnknown ErrorType = "Unknown"
)

// SPAPIError 表示 SP-API 错误。
//
// 包含详细的错误信息，帮助应用程序识别和处理错误。
type SPAPIError struct {
	// StatusCode 是 HTTP 状态码
	StatusCode int

	// ErrorCode 是 SP-API 错误代码
	ErrorCode string

	// Message 是错误消息
	Message string

	// RequestID 是请求 ID（用于追踪）
	RequestID string

	// Type 是错误类型
	Type ErrorType

	// Retryable 指示错误是否可重试
	Retryable bool

	// Details 包含额外的错误详情
	Details map[string]interface{}
}

// Error 实现 error 接口。
func (e *SPAPIError) Error() string {
	if e.RequestID != "" {
		return fmt.Sprintf("SP-API error [%s]: %s (status: %d, request ID: %s)",
			e.Type, e.Message, e.StatusCode, e.RequestID)
	}
	return fmt.Sprintf("SP-API error [%s]: %s (status: %d)",
		e.Type, e.Message, e.StatusCode)
}

// IsRetryable 返回错误是否可重试。
func (e *SPAPIError) IsRetryable() bool {
	return e.Retryable
}

// NewSPAPIError 创建新的 SP-API 错误。
//
// 参数:
//   - statusCode: HTTP 状态码
//   - message: 错误消息
//
// 返回值:
//   - *SPAPIError: SP-API 错误实例
//
// 示例:
//
//	err := errors.NewSPAPIError(429, "Request was throttled")
func NewSPAPIError(statusCode int, message string) *SPAPIError {
	err := &SPAPIError{
		StatusCode: statusCode,
		Message:    message,
		Details:    make(map[string]interface{}),
	}

	// 根据状态码设置错误类型和可重试性
	err.Type = classifyErrorType(statusCode)
	err.Retryable = isRetryableStatusCode(statusCode)

	return err
}

// NewSPAPIErrorFromResponse 从 HTTP 响应创建 SP-API 错误。
//
// 参数:
//   - resp: HTTP 响应
//   - message: 错误消息
//
// 返回值:
//   - *SPAPIError: SP-API 错误实例
//
// 示例:
//
//	err := errors.NewSPAPIErrorFromResponse(resp, "Request failed")
func NewSPAPIErrorFromResponse(resp *http.Response, message string) *SPAPIError {
	err := NewSPAPIError(resp.StatusCode, message)

	// 提取请求 ID
	if requestID := resp.Header.Get("x-amzn-requestid"); requestID != "" {
		err.RequestID = requestID
	}

	// 提取速率限制信息（如果是 429 错误）
	if resp.StatusCode == http.StatusTooManyRequests {
		if rateLimit := resp.Header.Get("x-amzn-ratelimit-limit"); rateLimit != "" {
			err.Details["rateLimit"] = rateLimit
		}
	}

	return err
}

// WithErrorCode 设置错误代码。
func (e *SPAPIError) WithErrorCode(code string) *SPAPIError {
	e.ErrorCode = code
	return e
}

// WithRequestID 设置请求 ID。
func (e *SPAPIError) WithRequestID(requestID string) *SPAPIError {
	e.RequestID = requestID
	return e
}

// WithDetail 添加错误详情。
func (e *SPAPIError) WithDetail(key string, value interface{}) *SPAPIError {
	e.Details[key] = value
	return e
}

// classifyErrorType 根据状态码分类错误类型。
func classifyErrorType(statusCode int) ErrorType {
	switch {
	case statusCode == http.StatusTooManyRequests:
		return ErrorTypeRateLimit
	case statusCode == http.StatusUnauthorized || statusCode == http.StatusForbidden:
		return ErrorTypeAuth
	case statusCode == http.StatusBadRequest:
		return ErrorTypeValidation
	case statusCode == http.StatusNotFound:
		return ErrorTypeNotFound
	case statusCode >= 500:
		return ErrorTypeServer
	default:
		return ErrorTypeUnknown
	}
}

// isRetryableStatusCode 判断状态码是否可重试。
func isRetryableStatusCode(statusCode int) bool {
	switch statusCode {
	case http.StatusTooManyRequests: // 429
		return true
	case http.StatusInternalServerError: // 500
		return true
	case http.StatusBadGateway: // 502
		return true
	case http.StatusServiceUnavailable: // 503
		return true
	case http.StatusGatewayTimeout: // 504
		return true
	default:
		return false
	}
}

// IsRateLimitError 检查错误是否为速率限制错误。
func IsRateLimitError(err error) bool {
	if apiErr, ok := err.(*SPAPIError); ok {
		return apiErr.Type == ErrorTypeRateLimit
	}
	return false
}

// IsAuthError 检查错误是否为认证错误。
func IsAuthError(err error) bool {
	if apiErr, ok := err.(*SPAPIError); ok {
		return apiErr.Type == ErrorTypeAuth
	}
	return false
}

// IsValidationError 检查错误是否为验证错误。
func IsValidationError(err error) bool {
	if apiErr, ok := err.(*SPAPIError); ok {
		return apiErr.Type == ErrorTypeValidation
	}
	return false
}

// IsServerError 检查错误是否为服务器错误。
func IsServerError(err error) bool {
	if apiErr, ok := err.(*SPAPIError); ok {
		return apiErr.Type == ErrorTypeServer
	}
	return false
}

// IsRetryable 检查错误是否可重试。
func IsRetryable(err error) bool {
	if apiErr, ok := err.(*SPAPIError); ok {
		return apiErr.Retryable
	}
	return false
}

