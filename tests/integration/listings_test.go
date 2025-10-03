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
// Licensed under the Apache License, Version 2.0.

// +build integration

package integration

import (
	"context"
	"os"
	"testing"

	listings "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi/listings-items-v2021-08-01"
)

func TestListingsItems_Integration(t *testing.T) {
	if os.Getenv("RUN_INTEGRATION_TESTS") == "" {
		t.Skip("跳过集成测试 (设置 RUN_INTEGRATION_TESTS=1 来运行)")
	}

	baseClient := createTestClient(t)
	defer baseClient.Close()

	listingsClient := listings.NewClient(baseClient)
	ctx := context.Background()

	// 需要从环境变量获取 Seller ID
	sellerId := os.Getenv("SP_API_SELLER_ID")
	if sellerId == "" {
		t.Skip("跳过测试: 缺少 SP_API_SELLER_ID 环境变量")
	}

	t.Run("SearchListingsItems", func(t *testing.T) {
		params := map[string]string{
			"marketplaceIds": "ATVPDKIKX0DER",
			"pageSize":       "5",
		}

		result, err := listingsClient.SearchListingsItems(ctx, sellerId, params)
		if err != nil {
			t.Errorf("SearchListingsItems failed: %v", err)
		}
		if result == nil {
			t.Error("SearchListingsItems returned nil")
		}
		t.Logf("✓ SearchListingsItems success")
	})

	t.Run("GetListingsItem", func(t *testing.T) {
		sku := os.Getenv("SP_API_TEST_SKU")
		if sku == "" {
			t.Skip("跳过测试: 缺少 SP_API_TEST_SKU 环境变量")
		}

		params := map[string]string{
			"marketplaceIds": "ATVPDKIKX0DER",
			"includedData":   "summaries",
		}

		result, err := listingsClient.GetListingsItem(ctx, sellerId, sku, params)
		if err != nil {
			t.Logf("GetListingsItem returned error: %v", err)
		} else {
			t.Logf("✓ GetListingsItem success: %v", result)
		}
	})
}

