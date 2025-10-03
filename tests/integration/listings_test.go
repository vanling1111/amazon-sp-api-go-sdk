// Copyright 2025 Amazon SP-API Go SDK Authors.
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

