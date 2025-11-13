package spapi_test

import (
	"testing"
	"time"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// 测试自定义Logger实现
type testLogger struct {
	messages []string
}

func (t *testLogger) Debug(msg string, fields ...spapi.Field) {
	t.messages = append(t.messages, "DEBUG: "+msg)
}

func (t *testLogger) Info(msg string, fields ...spapi.Field) {
	t.messages = append(t.messages, "INFO: "+msg)
}

func (t *testLogger) Warn(msg string, fields ...spapi.Field) {
	t.messages = append(t.messages, "WARN: "+msg)
}

func (t *testLogger) Error(msg string, fields ...spapi.Field) {
	t.messages = append(t.messages, "ERROR: "+msg)
}

func (t *testLogger) With(fields ...spapi.Field) spapi.Logger {
	return t
}

// 测试自定义Metrics实现
type testMetrics struct {
	requestCount int
	errorCount   int
}

func (t *testMetrics) RecordRequest(api, method string, duration time.Duration, statusCode int) {
	t.requestCount++
}

func (t *testMetrics) RecordError(api, errorType string) {
	t.errorCount++
}

func (t *testMetrics) RecordRateLimitWait(api string, duration time.Duration) {}

// TestWithLogger 测试自定义Logger
func TestWithLogger(t *testing.T) {
	logger := &testLogger{}

	client, err := spapi.NewClient(
		spapi.WithRegion(spapi.RegionNA),
		spapi.WithCredentials("test-id", "test-secret", "test-token"),
		spapi.WithLogger(logger),
	)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	defer client.Close()

	// 验证logger被设置
	config := client.Config()
	if config.Logger == nil {
		t.Error("Logger should not be nil")
	}
}

// TestWithMetrics 测试自定义Metrics
func TestWithMetrics(t *testing.T) {
	metrics := &testMetrics{}

	client, err := spapi.NewClient(
		spapi.WithRegion(spapi.RegionNA),
		spapi.WithCredentials("test-id", "test-secret", "test-token"),
		spapi.WithMetrics(metrics),
	)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	defer client.Close()

	// 验证metrics被设置
	config := client.Config()
	if config.Metrics == nil {
		t.Error("Metrics should not be nil")
	}
}

// TestDefaultNoOpImplementations 测试默认no-op实现
func TestDefaultNoOpImplementations(t *testing.T) {
	client, err := spapi.NewClient(
		spapi.WithRegion(spapi.RegionNA),
		spapi.WithCredentials("test-id", "test-secret", "test-token"),
		// 不提供Logger、Metrics、Tracer
	)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	defer client.Close()

	config := client.Config()

	// 验证默认实现被设置
	if config.Logger == nil {
		t.Error("Logger should have default no-op implementation")
	}
	if config.Metrics == nil {
		t.Error("Metrics should have default no-op implementation")
	}
	if config.Tracer == nil {
		t.Error("Tracer should have default no-op implementation")
	}
}

// ExampleWithLogger 演示如何使用自定义Logger
func ExampleWithLogger() {
	logger := &testLogger{}

	client, err := spapi.NewClient(
		spapi.WithRegion(spapi.RegionNA),
		spapi.WithCredentials("your-client-id", "your-client-secret", "your-refresh-token"),
		spapi.WithLogger(logger),
	)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// 使用客户端...
}

// ExampleWithMetrics 演示如何使用自定义Metrics
func ExampleWithMetrics() {
	metrics := &testMetrics{}

	client, err := spapi.NewClient(
		spapi.WithRegion(spapi.RegionNA),
		spapi.WithCredentials("your-client-id", "your-client-secret", "your-refresh-token"),
		spapi.WithMetrics(metrics),
	)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// 使用客户端...
}
