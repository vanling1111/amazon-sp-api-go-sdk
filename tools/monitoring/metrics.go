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
// terms of the applicable license.
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
// terms of the applicable license.

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
