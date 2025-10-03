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
// terms of the applicable license. All rights reserved.
//
package logging

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

// TestNewZapLogger tests creating Zap logger
func TestNewZapLogger(t *testing.T) {
	t.Run("production_config", func(t *testing.T) {
		logger, err := NewProductionLogger()
		require.NoError(t, err)
		assert.NotNil(t, logger)
		defer logger.Sync()
	})

	t.Run("development_config", func(t *testing.T) {
		logger, err := NewDevelopmentLogger()
		require.NoError(t, err)
		assert.NotNil(t, logger)
		defer logger.Sync()
	})

	t.Run("custom_config", func(t *testing.T) {
		config := zap.NewProductionConfig()
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)

		logger, err := NewZapLogger(&config)
		require.NoError(t, err)
		assert.NotNil(t, logger)
		defer logger.Sync()
	})
}

// TestZapLogger_LogLevels tests all log levels
func TestZapLogger_LogLevels(t *testing.T) {
	logger, err := NewDevelopmentLogger()
	require.NoError(t, err)
	defer logger.Sync()

	// 测试所有级别
	logger.Debug("debug message", String("key", "value"))
	logger.Info("info message", Int("count", 42))
	logger.Warn("warn message", String("reason", "test"))
	logger.Error("error message", Error(assert.AnError))
}

// TestZapLogger_With tests logger with fields
func TestZapLogger_With(t *testing.T) {
	logger, err := NewDevelopmentLogger()
	require.NoError(t, err)
	defer logger.Sync()

	// 创建带有固定字段的 logger
	requestLogger := logger.With(
		String("request_id", "123"),
		String("api", "orders"),
	)

	// 所有日志都会包含这些字段
	requestLogger.Info("processing request")
	requestLogger.Info("request complete")
}

// TestNopLogger tests no-op logger
func TestNopLogger(t *testing.T) {
	logger := NewNopLogger()

	// 不应该 panic
	logger.Debug("test")
	logger.Info("test")
	logger.Warn("test")
	logger.Error("test")

	newLogger := logger.With(String("key", "value"))
	assert.NotNil(t, newLogger)
}

// TestField_Helpers tests field helper functions
func TestField_Helpers(t *testing.T) {
	tests := []struct {
		name  string
		field Field
	}{
		{
			name:  "string_field",
			field: String("name", "value"),
		},
		{
			name:  "int_field",
			field: Int("count", 42),
		},
		{
			name:  "error_field",
			field: Error(assert.AnError),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotEmpty(t, tt.field.Key)
			assert.NotNil(t, tt.field.Value)
		})
	}
}

// BenchmarkZapLogger benchmarks Zap logger performance
func BenchmarkZapLogger(b *testing.B) {
	logger, _ := NewProductionLogger()
	defer logger.Sync()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("benchmark message",
			String("key1", "value1"),
			Int("key2", 123),
		)
	}
}

// BenchmarkNopLogger benchmarks no-op logger performance
func BenchmarkNopLogger(b *testing.B) {
	logger := NewNopLogger()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("benchmark message",
			String("key1", "value1"),
			Int("key2", 123),
		)
	}
}
