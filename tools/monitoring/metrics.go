// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

// Package monitoring provides monitoring and metrics utilities for SP-API SDK
package monitoring

import (
	"fmt"
	"time"
)

// Metric represents a single metric measurement
type Metric struct {
	Name      string
	Value     float64
	Timestamp time.Time
	Labels    map[string]string
}

// Collector collects and exports metrics
type Collector struct {
	metrics []Metric
}

// NewCollector creates a new metrics collector
func NewCollector() *Collector {
	return &Collector{
		metrics: make([]Metric, 0),
	}
}

// Record records a new metric
func (c *Collector) Record(name string, value float64, labels map[string]string) {
	c.metrics = append(c.metrics, Metric{
		Name:      name,
		Value:     value,
		Timestamp: time.Now(),
		Labels:    labels,
	})
}

// Export exports all collected metrics
func (c *Collector) Export() []Metric {
	return c.metrics
}

// Print prints all metrics to console
func (c *Collector) Print() {
	fmt.Println("=== Metrics ===")
	for _, m := range c.metrics {
		fmt.Printf("[%s] %s: %.2f %v\n", m.Timestamp.Format(time.RFC3339), m.Name, m.Value, m.Labels)
	}
}
