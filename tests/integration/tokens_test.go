// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

// +build integration

package integration

import (
	"context"
	"os"
	"testing"

	tokens "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi/tokens-v2021-03-01"
)

func TestTokens_Integration(t *testing.T) {
	if os.Getenv("RUN_INTEGRATION_TESTS") == "" {
		t.Skip("跳过集成测试 (设置 RUN_INTEGRATION_TESTS=1 来运行)")
	}

	baseClient := createTestClient(t)
	defer baseClient.Close()

	tokensClient := tokens.NewClient(baseClient)
	ctx := context.Background()

	t.Run("CreateRestrictedDataToken", func(t *testing.T) {
		request := map[string]interface{}{
			"restrictedResources": []map[string]interface{}{
				{
					"method": "GET",
					"path":   "/orders/v0/orders",
				},
			},
		}

		result, err := tokensClient.CreateRestrictedDataToken(ctx, request)
		if err != nil {
			t.Errorf("CreateRestrictedDataToken failed: %v", err)
		}
		if result == nil {
			t.Error("CreateRestrictedDataToken returned nil")
		}
		t.Logf("✓ CreateRestrictedDataToken success")
	})
}

