package metrics

import (
	"testing"
	"time"
)

// MockRecorder 是用于测试的 mock 记录器。
type MockRecorder struct {
	Counters   map[string]float64
	Gauges     map[string]float64
	Histograms map[string]float64
	Timings    map[string]time.Duration
}

func NewMockRecorder() *MockRecorder {
	return &MockRecorder{
		Counters:   make(map[string]float64),
		Gauges:     make(map[string]float64),
		Histograms: make(map[string]float64),
		Timings:    make(map[string]time.Duration),
	}
}

func (m *MockRecorder) RecordCounter(name string, value float64, labels map[string]string) {
	m.Counters[name] += value
}

func (m *MockRecorder) RecordGauge(name string, value float64, labels map[string]string) {
	m.Gauges[name] = value
}

func (m *MockRecorder) RecordHistogram(name string, value float64, labels map[string]string) {
	m.Histograms[name] = value
}

func (m *MockRecorder) RecordTiming(name string, duration time.Duration, labels map[string]string) {
	m.Timings[name] = duration
}

func TestNoOpRecorder(t *testing.T) {
	recorder := &NoOpRecorder{}

	// 测试所有方法都不会 panic
	recorder.RecordCounter("test", 1.0, nil)
	recorder.RecordGauge("test", 1.0, nil)
	recorder.RecordHistogram("test", 1.0, nil)
	recorder.RecordTiming("test", time.Second, nil)
}

func TestMockRecorder(t *testing.T) {
	recorder := NewMockRecorder()

	// 测试 Counter
	recorder.RecordCounter("requests", 1.0, map[string]string{"operation": "getOrders"})
	recorder.RecordCounter("requests", 2.0, map[string]string{"operation": "getOrders"})
	if count := recorder.Counters["requests"]; count != 3.0 {
		t.Errorf("Counter value = %v, want 3.0", count)
	}

	// 测试 Gauge
	recorder.RecordGauge("active_limiters", 5.0, nil)
	recorder.RecordGauge("active_limiters", 10.0, nil) // 覆盖
	if value := recorder.Gauges["active_limiters"]; value != 10.0 {
		t.Errorf("Gauge value = %v, want 10.0", value)
	}

	// 测试 Histogram
	recorder.RecordHistogram("request_duration", 0.123, nil)
	if value := recorder.Histograms["request_duration"]; value != 0.123 {
		t.Errorf("Histogram value = %v, want 0.123", value)
	}

	// 测试 Timing
	duration := 500 * time.Millisecond
	recorder.RecordTiming("api_call", duration, nil)
	if d := recorder.Timings["api_call"]; d != duration {
		t.Errorf("Timing value = %v, want %v", d, duration)
	}
}

func TestDefaultRecorder(t *testing.T) {
	// 确保默认记录器不是 nil
	if DefaultRecorder == nil {
		t.Error("DefaultRecorder should not be nil")
	}

	// 测试默认记录器是 NoOpRecorder
	if _, ok := DefaultRecorder.(*NoOpRecorder); !ok {
		t.Error("DefaultRecorder should be NoOpRecorder")
	}
}

func TestMetricConstants(t *testing.T) {
	// 测试指标名称常量
	tests := []struct {
		name     string
		constant string
	}{
		{"Request Total", MetricRequestTotal},
		{"Request Duration", MetricRequestDuration},
		{"Request Errors", MetricRequestErrors},
		{"Auth Token Refresh", MetricAuthTokenRefresh},
		{"Rate Limit Wait", MetricRateLimitWait},
		{"Rate Limit Active", MetricRateLimitActive},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant == "" {
				t.Errorf("%s constant should not be empty", tt.name)
			}
		})
	}
}

func TestLabelConstants(t *testing.T) {
	// 测试标签键常量
	tests := []struct {
		name     string
		constant string
	}{
		{"Operation", LabelOperation},
		{"Region", LabelRegion},
		{"Marketplace", LabelMarketplace},
		{"Status Code", LabelStatusCode},
		{"Error Type", LabelErrorType},
		{"Grant Type", LabelGrantType},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant == "" {
				t.Errorf("%s constant should not be empty", tt.name)
			}
		})
	}
}
