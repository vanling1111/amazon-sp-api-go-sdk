// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package monitoring

import (
	"testing"
	"time"
)

func TestNewCollector(t *testing.T) {
	collector := NewCollector()
	if collector == nil {
		t.Error("NewCollector returned nil")
	}
	if collector.metrics == nil {
		t.Error("Collector metrics not initialized")
	}
}

func TestCollector_Record(t *testing.T) {
	collector := NewCollector()
	
	labels := map[string]string{
		"operation": "GetOrders",
		"region":    "NA",
	}
	
	collector.Record("request_count", 1.0, labels)
	
	metrics := collector.Export()
	if len(metrics) != 1 {
		t.Errorf("Expected 1 metric, got %d", len(metrics))
	}
	
	if metrics[0].Name != "request_count" {
		t.Errorf("Expected name 'request_count', got '%s'", metrics[0].Name)
	}
	
	if metrics[0].Value != 1.0 {
		t.Errorf("Expected value 1.0, got %f", metrics[0].Value)
	}
}

func TestCollector_Multiple(t *testing.T) {
	collector := NewCollector()
	
	collector.Record("metric1", 10.0, nil)
	collector.Record("metric2", 20.0, nil)
	collector.Record("metric3", 30.0, nil)
	
	metrics := collector.Export()
	if len(metrics) != 3 {
		t.Errorf("Expected 3 metrics, got %d", len(metrics))
	}
}

func TestCollector_Print(t *testing.T) {
	collector := NewCollector()
	collector.Record("test_metric", 42.0, map[string]string{"tag": "value"})
	
	// Should not panic
	collector.Print()
}

func TestMetric_Timestamp(t *testing.T) {
	collector := NewCollector()
	before := time.Now()
	collector.Record("metric", 1.0, nil)
	after := time.Now()
	
	metrics := collector.Export()
	if metrics[0].Timestamp.Before(before) || metrics[0].Timestamp.After(after) {
		t.Error("Metric timestamp out of expected range")
	}
}

