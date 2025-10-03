// Package utils 提供内部工具函数。
package utils

import (
	"fmt"
	"net/http"
	"strings"
)

// BuildURL 构建完整的 URL。
//
// 此函数将基础 URL 和路径组合成完整的 URL，
// 并确保路径分隔符的正确性。
//
// 参数:
//   - baseURL: 基础 URL（如 https://sellingpartnerapi-na.amazon.com）
//   - path: 路径（如 /orders/v0/orders）
//
// 返回值:
//   - string: 完整的 URL
//
// 示例:
//
//	url := BuildURL("https://sellingpartnerapi-na.amazon.com", "/orders/v0/orders")
//	// 返回: "https://sellingpartnerapi-na.amazon.com/orders/v0/orders"
func BuildURL(baseURL, path string) string {
	// 移除 baseURL 末尾的斜杠
	baseURL = strings.TrimRight(baseURL, "/")

	// 确保 path 以斜杠开头
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	return baseURL + path
}

// GetContentType 从 HTTP 响应中获取 Content-Type。
//
// 参数:
//   - resp: HTTP 响应
//
// 返回值:
//   - string: Content-Type 值，如果不存在返回空字符串
//
// 示例:
//
//	contentType := GetContentType(resp)
//	if contentType == "application/json" {
//	    // 处理 JSON 响应
//	}
func GetContentType(resp *http.Response) string {
	return resp.Header.Get("Content-Type")
}

// IsJSONResponse 检查 HTTP 响应是否为 JSON 格式。
//
// 参数:
//   - resp: HTTP 响应
//
// 返回值:
//   - bool: 如果是 JSON 响应返回 true，否则返回 false
func IsJSONResponse(resp *http.Response) bool {
	contentType := GetContentType(resp)
	return strings.Contains(strings.ToLower(contentType), "application/json")
}

// GetRequestID 从 HTTP 响应中提取请求 ID。
//
// SP-API 在响应头中返回 x-amzn-requestid，可用于调试和追踪。
//
// 参数:
//   - resp: HTTP 响应
//
// 返回值:
//   - string: 请求 ID，如果不存在返回空字符串
//
// 官方文档:
//   - https://developer-docs.amazon.com/sp-api/docs/connecting-to-the-selling-partner-api
func GetRequestID(resp *http.Response) string {
	return resp.Header.Get("x-amzn-requestid")
}

// GetRateLimitHeader 从 HTTP 响应中提取速率限制信息。
//
// SP-API 在响应头中返回速率限制信息：
//   - x-amzn-ratelimit-limit: 速率限制（请求数/时间单位）
//
// 参数:
//   - resp: HTTP 响应
//
// 返回值:
//   - string: 速率限制信息，如果不存在返回空字符串
//
// 官方文档:
//   - https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits
func GetRateLimitHeader(resp *http.Response) string {
	return resp.Header.Get("x-amzn-ratelimit-limit")
}

// FormatHTTPError 格式化 HTTP 错误消息。
//
// 参数:
//   - resp: HTTP 响应
//
// 返回值:
//   - string: 格式化的错误消息
//
// 示例:
//
//	if resp.StatusCode != http.StatusOK {
//	    errMsg := FormatHTTPError(resp)
//	    return fmt.Errorf("API request failed: %s", errMsg)
//	}
func FormatHTTPError(resp *http.Response) string {
	requestID := GetRequestID(resp)
	if requestID != "" {
		return fmt.Sprintf("status %d (%s), request ID: %s",
			resp.StatusCode, resp.Status, requestID)
	}
	return fmt.Sprintf("status %d (%s)", resp.StatusCode, resp.Status)
}

// IsRetryableStatusCode 检查 HTTP 状态码是否应该重试。
//
// 以下状态码被认为是可重试的：
//   - 429 Too Many Requests
//   - 500 Internal Server Error
//   - 502 Bad Gateway
//   - 503 Service Unavailable
//   - 504 Gateway Timeout
//
// 参数:
//   - statusCode: HTTP 状态码
//
// 返回值:
//   - bool: 如果应该重试返回 true，否则返回 false
func IsRetryableStatusCode(statusCode int) bool {
	switch statusCode {
	case http.StatusTooManyRequests, // 429
		http.StatusInternalServerError, // 500
		http.StatusBadGateway,          // 502
		http.StatusServiceUnavailable,  // 503
		http.StatusGatewayTimeout:      // 504
		return true
	default:
		return false
	}
}
