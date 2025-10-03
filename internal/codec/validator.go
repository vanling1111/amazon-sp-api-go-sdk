// Package codec 提供数据编解码和验证功能。
package codec

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

// ValidationError 表示验证错误。
type ValidationError struct {
	// Field 是字段名
	Field string

	// Message 是错误消息
	Message string
}

// Error 实现 error 接口。
func (e *ValidationError) Error() string {
	if e.Field != "" {
		return fmt.Sprintf("validation error: field '%s': %s", e.Field, e.Message)
	}
	return fmt.Sprintf("validation error: %s", e.Message)
}

// ValidationErrors 表示多个验证错误。
type ValidationErrors []ValidationError

// Error 实现 error 接口。
func (e ValidationErrors) Error() string {
	if len(e) == 0 {
		return "no validation errors"
	}

	if len(e) == 1 {
		return e[0].Error()
	}

	var messages []string
	for _, err := range e {
		messages = append(messages, err.Error())
	}

	return fmt.Sprintf("validation errors: %s", strings.Join(messages, "; "))
}

// Validator 表示数据验证器。
type Validator struct {
	errors ValidationErrors
}

// NewValidator 创建新的验证器。
//
// 返回值:
//   - *Validator: 验证器实例
//
// 示例:
//
//	validator := codec.NewValidator()
//	validator.Required("orderId", order.OrderID)
//	validator.MaxLength("status", order.Status, 50)
//
//	if err := validator.Error(); err != nil {
//	    return err
//	}
func NewValidator() *Validator {
	return &Validator{
		errors: make(ValidationErrors, 0),
	}
}

// Required 验证字段是否非空。
//
// 参数:
//   - field: 字段名
//   - value: 字段值
//
// 示例:
//
//	validator.Required("orderId", order.OrderID)
func (v *Validator) Required(field string, value string) {
	if strings.TrimSpace(value) == "" {
		v.errors = append(v.errors, ValidationError{
			Field:   field,
			Message: "is required",
		})
	}
}

// MinLength 验证字符串最小长度。
//
// 参数:
//   - field: 字段名
//   - value: 字段值
//   - min: 最小长度
//
// 示例:
//
//	validator.MinLength("orderId", order.OrderID, 5)
func (v *Validator) MinLength(field string, value string, min int) {
	if len(value) < min {
		v.errors = append(v.errors, ValidationError{
			Field:   field,
			Message: fmt.Sprintf("must be at least %d characters long", min),
		})
	}
}

// MaxLength 验证字符串最大长度。
//
// 参数:
//   - field: 字段名
//   - value: 字段值
//   - max: 最大长度
//
// 示例:
//
//	validator.MaxLength("status", order.Status, 50)
func (v *Validator) MaxLength(field string, value string, max int) {
	if len(value) > max {
		v.errors = append(v.errors, ValidationError{
			Field:   field,
			Message: fmt.Sprintf("must be at most %d characters long", max),
		})
	}
}

// Min 验证数值最小值。
//
// 参数:
//   - field: 字段名
//   - value: 字段值
//   - min: 最小值
//
// 示例:
//
//	validator.Min("quantity", order.Quantity, 1)
func (v *Validator) Min(field string, value int, min int) {
	if value < min {
		v.errors = append(v.errors, ValidationError{
			Field:   field,
			Message: fmt.Sprintf("must be at least %d", min),
		})
	}
}

// Max 验证数值最大值。
//
// 参数:
//   - field: 字段名
//   - value: 字段值
//   - max: 最大值
//
// 示例:
//
//	validator.Max("quantity", order.Quantity, 1000)
func (v *Validator) Max(field string, value int, max int) {
	if value > max {
		v.errors = append(v.errors, ValidationError{
			Field:   field,
			Message: fmt.Sprintf("must be at most %d", max),
		})
	}
}

// Range 验证数值范围。
//
// 参数:
//   - field: 字段名
//   - value: 字段值
//   - min: 最小值
//   - max: 最大值
//
// 示例:
//
//	validator.Range("quantity", order.Quantity, 1, 1000)
func (v *Validator) Range(field string, value int, min int, max int) {
	if value < min || value > max {
		v.errors = append(v.errors, ValidationError{
			Field:   field,
			Message: fmt.Sprintf("must be between %d and %d", min, max),
		})
	}
}

// Email 验证电子邮件格式。
//
// 参数:
//   - field: 字段名
//   - value: 字段值
//
// 示例:
//
//	validator.Email("email", user.Email)
func (v *Validator) Email(field string, value string) {
	// 简单的邮箱格式验证
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if value != "" && !emailRegex.MatchString(value) {
		v.errors = append(v.errors, ValidationError{
			Field:   field,
			Message: "must be a valid email address",
		})
	}
}

// URL 验证 URL 格式。
//
// 参数:
//   - field: 字段名
//   - value: 字段值
//
// 示例:
//
//	validator.URL("callbackUrl", config.CallbackURL)
func (v *Validator) URL(field string, value string) {
	if value != "" {
		_, err := url.Parse(value)
		if err != nil {
			v.errors = append(v.errors, ValidationError{
				Field:   field,
				Message: "must be a valid URL",
			})
		}
	}
}

// Pattern 验证字符串是否匹配正则表达式。
//
// 参数:
//   - field: 字段名
//   - value: 字段值
//   - pattern: 正则表达式模式
//   - message: 错误消息
//
// 示例:
//
//	validator.Pattern("orderID", order.ID, `^[0-9]{3}-[0-9]{7}-[0-9]{7}$`, "must be a valid Amazon order ID")
func (v *Validator) Pattern(field string, value string, pattern string, message string) {
	regex, err := regexp.Compile(pattern)
	if err != nil {
		v.errors = append(v.errors, ValidationError{
			Field:   field,
			Message: fmt.Sprintf("invalid pattern: %v", err),
		})
		return
	}

	if value != "" && !regex.MatchString(value) {
		v.errors = append(v.errors, ValidationError{
			Field:   field,
			Message: message,
		})
	}
}

// OneOf 验证字段值是否为指定值之一。
//
// 参数:
//   - field: 字段名
//   - value: 字段值
//   - allowed: 允许的值列表
//
// 示例:
//
//	validator.OneOf("status", order.Status, []string{"Pending", "Shipped", "Delivered"})
func (v *Validator) OneOf(field string, value string, allowed []string) {
	if value == "" {
		return
	}

	for _, a := range allowed {
		if value == a {
			return
		}
	}

	v.errors = append(v.errors, ValidationError{
		Field:   field,
		Message: fmt.Sprintf("must be one of: %s", strings.Join(allowed, ", ")),
	})
}

// Custom 添加自定义验证错误。
//
// 参数:
//   - field: 字段名
//   - message: 错误消息
//
// 示例:
//
//	if order.Quantity > order.AvailableStock {
//	    validator.Custom("quantity", "exceeds available stock")
//	}
func (v *Validator) Custom(field string, message string) {
	v.errors = append(v.errors, ValidationError{
		Field:   field,
		Message: message,
	})
}

// Error 返回验证错误。
//
// 如果没有错误，返回 nil。
//
// 返回值:
//   - error: 验证错误，如果没有错误则返回 nil
//
// 示例:
//
//	if err := validator.Error(); err != nil {
//	    return err
//	}
func (v *Validator) Error() error {
	if len(v.errors) == 0 {
		return nil
	}
	return v.errors
}

// HasErrors 返回是否有验证错误。
//
// 返回值:
//   - bool: 如果有错误返回 true，否则返回 false
func (v *Validator) HasErrors() bool {
	return len(v.errors) > 0
}

// Errors 返回所有验证错误。
//
// 返回值:
//   - ValidationErrors: 验证错误列表
func (v *Validator) Errors() ValidationErrors {
	return v.errors
}

// Clear 清空所有验证错误。
func (v *Validator) Clear() {
	v.errors = make(ValidationErrors, 0)
}
