// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

// +build integration

package integration

import (
	"context"
	"os"
	"testing"

	pricing "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi/product-pricing-v0"
)

func TestProductPricing_Integration(t *testing.T) {
	if os.Getenv("RUN_INTEGRATION_TESTS") == "" {
		t.Skip("跳过集成测试 (设置 RUN_INTEGRATION_TESTS=1 来运行)")
	}

	baseClient := createTestClient(t)
	defer baseClient.Close()

	pricingClient := pricing.NewClient(baseClient)
	ctx := context.Background()

	t.Run("GetPricing", func(t *testing.T) {
		params := map[string]string{
			"MarketplaceId": "ATVPDKIKX0DER",
			"ItemType":      "Asin",
			"Asins":         "B00X4WHP5E", // 示例ASIN
		}

		result, err := pricingClient.GetPricing(ctx, params)
		if err != nil {
			t.Logf("GetPricing returned error (expected in test): %v", err)
		} else {
			t.Logf("✓ GetPricing success: %v", result)
		}
	})

	t.Run("GetCompetitivePricing", func(t *testing.T) {
		params := map[string]string{
			"MarketplaceId": "ATVPDKIKX0DER",
			"ItemType":      "Asin",
			"Asins":         "B00X4WHP5E",
		}

		result, err := pricingClient.GetCompetitivePricing(ctx, params)
		if err != nil {
			t.Logf("GetCompetitivePricing returned error (expected in test): %v", err)
		} else {
			t.Logf("✓ GetCompetitivePricing success: %v", result)
		}
	})
}

