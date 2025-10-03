// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package benchmarks

import (
	"context"
	"testing"

	"github.com/vanling1111/amazon-sp-api-go-sdk/internal/models"
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
			spapi.WithRegion(models.RegionNA),
			spapi.WithCredentials("test", "test", "test"),
		)
		client.Close()
	}
}

// BenchmarkAPIClientCreation benchmarks API client creation
func BenchmarkAPIClientCreation(b *testing.B) {
	baseClient, _ := spapi.NewClient(
		spapi.WithRegion(models.RegionNA),
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
