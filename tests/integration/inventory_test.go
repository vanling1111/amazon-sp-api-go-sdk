// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

// +build integration

package integration

import (
	"context"
	"os"
	"testing"

	inventory "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi/fba-inventory-v1"
)

func TestFBAInventory_Integration(t *testing.T) {
	if os.Getenv("RUN_INTEGRATION_TESTS") == "" {
		t.Skip("跳过集成测试 (设置 RUN_INTEGRATION_TESTS=1 来运行)")
	}

	baseClient := createTestClient(t)
	defer baseClient.Close()

	inventoryClient := inventory.NewClient(baseClient)
	ctx := context.Background()

	t.Run("GetInventorySummaries", func(t *testing.T) {
		params := map[string]string{
			"granularityType": "Marketplace",
			"granularityId":   "ATVPDKIKX0DER",
			"marketplaceIds":  "ATVPDKIKX0DER",
		}

		result, err := inventoryClient.GetInventorySummaries(ctx, params)
		if err != nil {
			t.Errorf("GetInventorySummaries failed: %v", err)
		}
		if result == nil {
			t.Error("GetInventorySummaries returned nil")
		}
		t.Logf("✓ GetInventorySummaries success")
	})
}

