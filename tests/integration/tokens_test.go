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

