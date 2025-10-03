// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

// +build integration

package integration

import (
	"context"
	"os"
	"testing"

	sellers "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi/sellers-v1"
)

func TestSellers_Integration(t *testing.T) {
	if os.Getenv("RUN_INTEGRATION_TESTS") == "" {
		t.Skip("跳过集成测试 (设置 RUN_INTEGRATION_TESTS=1 来运行)")
	}

	baseClient := createTestClient(t)
	defer baseClient.Close()

	sellersClient := sellers.NewClient(baseClient)
	ctx := context.Background()

	t.Run("GetMarketplaceParticipations", func(t *testing.T) {
		result, err := sellersClient.GetMarketplaceParticipations(ctx, nil)
		if err != nil {
			t.Errorf("GetMarketplaceParticipations failed: %v", err)
		}
		if result == nil {
			t.Error("GetMarketplaceParticipations returned nil")
		}
		t.Logf("✓ GetMarketplaceParticipations success")
	})
}

