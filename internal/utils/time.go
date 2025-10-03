package utils

import (
	"time"
)

// FormatISO8601 将时间格式化为 ISO 8601 格式。
//
// SP-API 要求某些时间戳使用 ISO 8601 格式。
//
// 参数:
//   - t: 时间
//
// 返回值:
//   - string: ISO 8601 格式的时间字符串
//
// 示例:
//
//	timestamp := FormatISO8601(time.Now())
//	// 返回: "2025-01-15T10:30:45Z"
//
// 官方文档:
//   - https://developer-docs.amazon.com/sp-api/docs/connecting-to-the-selling-partner-api
func FormatISO8601(t time.Time) string {
	return t.UTC().Format(time.RFC3339)
}

// FormatAMZDate 将时间格式化为 Amazon 日期格式。
//
// # SP-API 要求 x-amz-date 头使用特定格式：YYYYMMDDTHHmmssZ
//
// 参数:
//   - t: 时间
//
// 返回值:
//   - string: Amazon 日期格式的时间字符串
//
// 示例:
//
//	dateStr := FormatAMZDate(time.Now())
//	// 返回: "20250115T103045Z"
//
// 官方文档:
//   - https://developer-docs.amazon.com/sp-api/docs/connecting-to-the-selling-partner-api#step-3-add-headers-to-the-uri
func FormatAMZDate(t time.Time) string {
	return t.UTC().Format("20060102T150405Z")
}

// ParseISO8601 解析 ISO 8601 格式的时间字符串。
//
// 参数:
//   - s: ISO 8601 格式的时间字符串
//
// 返回值:
//   - time.Time: 解析后的时间
//   - error: 如果解析失败，返回错误
//
// 示例:
//
//	t, err := ParseISO8601("2025-01-15T10:30:45Z")
//	if err != nil {
//	    log.Fatal(err)
//	}
func ParseISO8601(s string) (time.Time, error) {
	return time.Parse(time.RFC3339, s)
}

// IsExpired 检查时间是否已过期。
//
// 参数:
//   - t: 过期时间
//   - buffer: 提前过期的缓冲时间
//
// 返回值:
//   - bool: 如果已过期（含缓冲）返回 true，否则返回 false
//
// 示例:
//
//	// 检查令牌是否在 60 秒内过期
//	if IsExpired(token.ExpiresAt, 60*time.Second) {
//	    // 刷新令牌
//	}
func IsExpired(t time.Time, buffer time.Duration) bool {
	return time.Now().Add(buffer).After(t)
}

// CalculateExpireTime 根据当前时间和过期秒数计算过期时间。
//
// 参数:
//   - expiresIn: 过期秒数
//
// 返回值:
//   - time.Time: 过期时间
//
// 示例:
//
//	// LWA 令牌过期时间
//	expiresAt := CalculateExpireTime(3600)  // 1 小时后过期
func CalculateExpireTime(expiresIn int) time.Time {
	return time.Now().Add(time.Duration(expiresIn) * time.Second)
}
