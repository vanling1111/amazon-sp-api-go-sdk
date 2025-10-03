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

// Package logging 提供结构化日志功能。
//
// 此包封装了 Zap 日志库，提供统一的日志接口和配置。
package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger 日志接口。
//
// SDK 的所有日志都通过此接口输出。
// 用户可以实现此接口来自定义日志行为。
type Logger interface {
	// Debug 记录调试级别日志
	Debug(msg string, fields ...Field)

	// Info 记录信息级别日志
	Info(msg string, fields ...Field)

	// Warn 记录警告级别日志
	Warn(msg string, fields ...Field)

	// Error 记录错误级别日志
	Error(msg string, fields ...Field)

	// With 创建带有附加字段的新 logger
	With(fields ...Field) Logger
}

// Field 日志字段
type Field struct {
	Key   string
	Value interface{}
}

// String 创建字符串字段
func String(key, val string) Field {
	return Field{Key: key, Value: val}
}

// Int 创建整数字段
func Int(key string, val int) Field {
	return Field{Key: key, Value: val}
}

// Duration 创建时间段字段
func Duration(key string, val interface{}) Field {
	return Field{Key: key, Value: val}
}

// Error 创建错误字段
func Error(err error) Field {
	return Field{Key: "error", Value: err}
}

// ZapLogger Zap 日志器实现
type ZapLogger struct {
	logger *zap.Logger
}

// NewZapLogger 创建 Zap 日志器。
//
// 参数:
//   - config: Zap 配置（如果为 nil，使用生产配置）
//
// 返回值:
//   - *ZapLogger: Zap 日志器实例
//   - error: 如果创建失败，返回错误
//
// 示例:
//
//	// 使用生产配置
//	logger, _ := logging.NewZapLogger(nil)
//
//	// 使用开发配置
//	logger, _ := logging.NewZapLogger(zap.NewDevelopmentConfig())
func NewZapLogger(config *zap.Config) (*ZapLogger, error) {
	var logger *zap.Logger
	var err error

	if config == nil {
		// 使用生产配置
		logger, err = zap.NewProduction()
	} else {
		logger, err = config.Build()
	}

	if err != nil {
		return nil, err
	}

	return &ZapLogger{logger: logger}, nil
}

// NewDevelopmentLogger 创建开发环境日志器。
//
// 开发环境日志器特点：
//   - 人类可读格式
//   - 彩色输出
//   - 调用栈信息
//   - Debug 级别
func NewDevelopmentLogger() (*ZapLogger, error) {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return NewZapLogger(&config)
}

// NewProductionLogger 创建生产环境日志器。
//
// 生产环境日志器特点：
//   - JSON 格式
//   - Info 级别
//   - 采样（高频日志采样，避免日志爆炸）
//   - 错误堆栈
func NewProductionLogger() (*ZapLogger, error) {
	return NewZapLogger(nil)
}

// Debug 记录调试日志
func (l *ZapLogger) Debug(msg string, fields ...Field) {
	l.logger.Debug(msg, convertFields(fields)...)
}

// Info 记录信息日志
func (l *ZapLogger) Info(msg string, fields ...Field) {
	l.logger.Info(msg, convertFields(fields)...)
}

// Warn 记录警告日志
func (l *ZapLogger) Warn(msg string, fields ...Field) {
	l.logger.Warn(msg, convertFields(fields)...)
}

// Error 记录错误日志
func (l *ZapLogger) Error(msg string, fields ...Field) {
	l.logger.Error(msg, convertFields(fields)...)
}

// With 创建带有附加字段的新 logger
func (l *ZapLogger) With(fields ...Field) Logger {
	return &ZapLogger{
		logger: l.logger.With(convertFields(fields)...),
	}
}

// Sync 刷新缓冲的日志
func (l *ZapLogger) Sync() error {
	return l.logger.Sync()
}

// convertFields 转换字段格式
func convertFields(fields []Field) []zap.Field {
	zapFields := make([]zap.Field, len(fields))
	for i, f := range fields {
		zapFields[i] = zap.Any(f.Key, f.Value)
	}
	return zapFields
}

// NopLogger 空日志器（不输出任何日志）
type NopLogger struct{}

// NewNopLogger 创建空日志器
func NewNopLogger() Logger {
	return &NopLogger{}
}

func (n *NopLogger) Debug(msg string, fields ...Field) {}
func (n *NopLogger) Info(msg string, fields ...Field)  {}
func (n *NopLogger) Warn(msg string, fields ...Field)  {}
func (n *NopLogger) Error(msg string, fields ...Field) {}
func (n *NopLogger) With(fields ...Field) Logger       { return n }

