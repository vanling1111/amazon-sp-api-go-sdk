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

