package utils

import (
	"crypto/rand"
	"encoding/hex"
	"strings"
)

// GenerateRequestID 生成唯一的请求 ID。
//
// 请求 ID 用于追踪和调试 API 请求。
//
// 返回值:
//   - string: 请求 ID（32 字符的十六进制字符串）
//   - error: 如果生成失败，返回错误
//
// 示例:
//
//	requestID, err := GenerateRequestID()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	req.Header.Set("x-request-id", requestID)
func GenerateRequestID() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// TruncateString 截断字符串到指定长度。
//
// 如果字符串长度超过 maxLen，则截断并添加省略号。
//
// 参数:
//   - s: 原始字符串
//   - maxLen: 最大长度
//
// 返回值:
//   - string: 截断后的字符串
//
// 示例:
//
//	shortened := TruncateString("very long string here", 10)
//	// 返回: "very lo..."
func TruncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	if maxLen < 3 {
		return s[:maxLen]
	}
	return s[:maxLen-3] + "..."
}

// JoinNonEmpty 连接非空字符串。
//
// 参数:
//   - sep: 分隔符
//   - strs: 字符串列表
//
// 返回值:
//   - string: 连接后的字符串
//
// 示例:
//
//	result := JoinNonEmpty(", ", "a", "", "b", "", "c")
//	// 返回: "a, b, c"
func JoinNonEmpty(sep string, strs ...string) string {
	var nonEmpty []string
	for _, s := range strs {
		if s != "" {
			nonEmpty = append(nonEmpty, s)
		}
	}
	return strings.Join(nonEmpty, sep)
}

// SanitizeLogString 清理日志字符串，移除敏感信息。
//
// 此函数用于日志记录前清理可能包含敏感信息的字符串。
//
// 参数:
//   - s: 原始字符串
//
// 返回值:
//   - string: 清理后的字符串
//
// 示例:
//
//	safe := SanitizeLogString("token=abc123xyz")
//	// 返回: "token=***"
func SanitizeLogString(s string) string {
	// 简单实现：如果包含 token、secret、password 等敏感词，则脱敏
	sensitiveWords := []string{"token", "secret", "password", "key"}
	lower := strings.ToLower(s)

	for _, word := range sensitiveWords {
		if strings.Contains(lower, word) {
			// 查找 = 或 : 后的值并替换为 ***
			parts := strings.SplitN(s, "=", 2)
			if len(parts) == 2 {
				return parts[0] + "=***"
			}
			parts = strings.SplitN(s, ":", 2)
			if len(parts) == 2 {
				return parts[0] + ":***"
			}
			return "***"
		}
	}

	return s
}

// MaskString 掩码字符串的中间部分。
//
// 保留开头和结尾的字符，中间用星号替换。
//
// 参数:
//   - s: 原始字符串
//   - visibleStart: 开头可见字符数
//   - visibleEnd: 结尾可见字符数
//
// 返回值:
//   - string: 掩码后的字符串
//
// 示例:
//
//	masked := MaskString("1234567890", 2, 2)
//	// 返回: "12******90"
func MaskString(s string, visibleStart, visibleEnd int) string {
	if len(s) <= visibleStart+visibleEnd {
		return s
	}

	start := s[:visibleStart]
	end := s[len(s)-visibleEnd:]
	maskLen := len(s) - visibleStart - visibleEnd

	return start + strings.Repeat("*", maskLen) + end
}

