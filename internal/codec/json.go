// Package codec 提供数据编解码和验证功能。
//
// 此包实现了 JSON 编解码、数据验证等功能，
// 确保与 SP-API 的数据交互符合规范。
//
// 官方文档:
//   - https://developer-docs.amazon.com/sp-api/docs/
package codec

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

// Encoder 表示 JSON 编码器。
//
// 用于将 Go 结构体编码为 JSON 格式。
type Encoder struct {
	// indent 指定是否使用缩进格式
	indent bool

	// indentPrefix 缩进前缀
	indentPrefix string

	// indentValue 缩进值
	indentValue string
}

// EncoderOption 表示编码器选项。
type EncoderOption func(*Encoder)

// WithIndent 设置使用缩进格式。
//
// 参数:
//   - prefix: 缩进前缀
//   - indent: 缩进值（通常为 "  " 或 "\t"）
//
// 示例:
//
//	encoder := codec.NewEncoder(codec.WithIndent("", "  "))
func WithIndent(prefix, indent string) EncoderOption {
	return func(e *Encoder) {
		e.indent = true
		e.indentPrefix = prefix
		e.indentValue = indent
	}
}

// NewEncoder 创建新的 JSON 编码器。
//
// 参数:
//   - opts: 编码器选项
//
// 返回值:
//   - *Encoder: 编码器实例
//
// 示例:
//
//	// 创建默认编码器
//	encoder := codec.NewEncoder()
//
//	// 创建带缩进的编码器
//	encoder := codec.NewEncoder(codec.WithIndent("", "  "))
func NewEncoder(opts ...EncoderOption) *Encoder {
	e := &Encoder{
		indent:       false,
		indentPrefix: "",
		indentValue:  "",
	}

	for _, opt := range opts {
		opt(e)
	}

	return e
}

// Marshal 将 Go 值编码为 JSON。
//
// 参数:
//   - v: 要编码的值
//
// 返回值:
//   - []byte: JSON 字节数组
//   - error: 如果编码失败，返回错误
//
// 示例:
//
//	type Order struct {
//	    OrderID string `json:"orderId"`
//	    Status  string `json:"status"`
//	}
//
//	order := Order{OrderID: "123", Status: "Shipped"}
//	encoder := codec.NewEncoder()
//	data, err := encoder.Marshal(order)
func (e *Encoder) Marshal(v interface{}) ([]byte, error) {
	if v == nil {
		return nil, fmt.Errorf("cannot marshal nil value")
	}

	if e.indent {
		return json.MarshalIndent(v, e.indentPrefix, e.indentValue)
	}

	return json.Marshal(v)
}

// Decoder 表示 JSON 解码器。
//
// 用于将 JSON 数据解码为 Go 结构体。
type Decoder struct {
	// disallowUnknownFields 指定是否拒绝未知字段
	disallowUnknownFields bool
}

// DecoderOption 表示解码器选项。
type DecoderOption func(*Decoder)

// WithDisallowUnknownFields 设置拒绝未知字段。
//
// 当 JSON 中包含结构体中未定义的字段时，解码会失败。
//
// 示例:
//
//	decoder := codec.NewDecoder(codec.WithDisallowUnknownFields())
func WithDisallowUnknownFields() DecoderOption {
	return func(d *Decoder) {
		d.disallowUnknownFields = true
	}
}

// NewDecoder 创建新的 JSON 解码器。
//
// 参数:
//   - opts: 解码器选项
//
// 返回值:
//   - *Decoder: 解码器实例
//
// 示例:
//
//	// 创建默认解码器
//	decoder := codec.NewDecoder()
//
//	// 创建拒绝未知字段的解码器
//	decoder := codec.NewDecoder(codec.WithDisallowUnknownFields())
func NewDecoder(opts ...DecoderOption) *Decoder {
	d := &Decoder{
		disallowUnknownFields: false,
	}

	for _, opt := range opts {
		opt(d)
	}

	return d
}

// Unmarshal 将 JSON 解码为 Go 值。
//
// 参数:
//   - data: JSON 字节数组
//   - v: 目标值的指针
//
// 返回值:
//   - error: 如果解码失败，返回错误
//
// 示例:
//
//	var order Order
//	decoder := codec.NewDecoder()
//	err := decoder.Unmarshal(data, &order)
func (d *Decoder) Unmarshal(data []byte, v interface{}) error {
	if len(data) == 0 {
		return fmt.Errorf("cannot unmarshal empty data")
	}

	if v == nil {
		return fmt.Errorf("cannot unmarshal into nil value")
	}

	if d.disallowUnknownFields {
		dec := json.NewDecoder(bytes.NewReader(data))
		dec.DisallowUnknownFields()
		return dec.Decode(v)
	}

	return json.Unmarshal(data, v)
}

// UnmarshalFromReader 从 io.Reader 解码 JSON。
//
// 参数:
//   - r: 数据源
//   - v: 目标值的指针
//
// 返回值:
//   - error: 如果解码失败，返回错误
//
// 示例:
//
//	var order Order
//	decoder := codec.NewDecoder()
//	err := decoder.UnmarshalFromReader(resp.Body, &order)
func (d *Decoder) UnmarshalFromReader(r io.Reader, v interface{}) error {
	if r == nil {
		return fmt.Errorf("cannot unmarshal from nil reader")
	}

	if v == nil {
		return fmt.Errorf("cannot unmarshal into nil value")
	}

	dec := json.NewDecoder(r)

	if d.disallowUnknownFields {
		dec.DisallowUnknownFields()
	}

	return dec.Decode(v)
}

// MarshalJSON 是便捷函数，使用默认编码器将值编码为 JSON。
//
// 参数:
//   - v: 要编码的值
//
// 返回值:
//   - []byte: JSON 字节数组
//   - error: 如果编码失败，返回错误
//
// 示例:
//
//	data, err := codec.MarshalJSON(order)
func MarshalJSON(v interface{}) ([]byte, error) {
	encoder := NewEncoder()
	return encoder.Marshal(v)
}

// UnmarshalJSON 是便捷函数，使用默认解码器将 JSON 解码为值。
//
// 参数:
//   - data: JSON 字节数组
//   - v: 目标值的指针
//
// 返回值:
//   - error: 如果解码失败，返回错误
//
// 示例:
//
//	var order Order
//	err := codec.UnmarshalJSON(data, &order)
func UnmarshalJSON(data []byte, v interface{}) error {
	decoder := NewDecoder()
	return decoder.Unmarshal(data, v)
}

// MarshalIndentJSON 是便捷函数，将值编码为带缩进的 JSON。
//
// 参数:
//   - v: 要编码的值
//   - prefix: 缩进前缀
//   - indent: 缩进值
//
// 返回值:
//   - []byte: JSON 字节数组
//   - error: 如果编码失败，返回错误
//
// 示例:
//
//	data, err := codec.MarshalIndentJSON(order, "", "  ")
func MarshalIndentJSON(v interface{}, prefix, indent string) ([]byte, error) {
	encoder := NewEncoder(WithIndent(prefix, indent))
	return encoder.Marshal(v)
}

