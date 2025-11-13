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

package benchmarks

import (
	"context"
	"testing"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
	orders "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi/orders-v0"
)

// TestBenchmarks is a placeholder test to ensure the package is tested
func TestBenchmarks(t *testing.T) {
	t.Log("Benchmark tests require -bench flag to run")
}

// BenchmarkClientCreation benchmarks client creation performance
func BenchmarkClientCreation(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		client, _ := spapi.NewClient(
			spapi.WithRegion(spapi.RegionNA),
			spapi.WithCredentials("test", "test", "test"),
		)
		client.Close()
	}
}

// BenchmarkAPIClientCreation benchmarks API client creation
func BenchmarkAPIClientCreation(b *testing.B) {
	baseClient, _ := spapi.NewClient(
		spapi.WithRegion(spapi.RegionNA),
		spapi.WithCredentials("test", "test", "test"),
	)
	defer baseClient.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = orders.NewClient(baseClient)
	}
}

// BenchmarkContextCreation benchmarks context creation overhead
func BenchmarkContextCreation(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = context.Background()
	}
}
