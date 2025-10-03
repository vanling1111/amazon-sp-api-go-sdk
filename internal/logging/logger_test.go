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

