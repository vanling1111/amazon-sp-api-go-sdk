// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

// +build integration

package integration

import (
	"context"
	"os"
	"testing"

	catalog "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi/catalog-items-v2022-04-01"
)

func TestCatalogItems_Integration(t *testing.T) {
	if os.Getenv("RUN_INTEGRATION_TESTS") == "" {
		t.Skip("跳过集成测试 (设置 RUN_INTEGRATION_TESTS=1 来运行)")
	}

	baseClient := createTestClient(t)
	defer baseClient.Close()

	catalogClient := catalog.NewClient(baseClient)
	ctx := context.Background()

	t.Run("SearchCatalogItems", func(t *testing.T) {
		params := map[string]string{
			"marketplaceIds": "ATVPDKIKX0DER",
			"keywords":       "book",
			"pageSize":       "10",
		}

		result, err := catalogClient.SearchCatalogItems(ctx, params)
		if err != nil {
			t.Errorf("SearchCatalogItems failed: %v", err)
		}
		if result == nil {
			t.Error("SearchCatalogItems returned nil")
		}
		t.Logf("✓ SearchCatalogItems success")
	})
}

