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

	feeds "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi/feeds-v2021-06-30"
)

func TestFeeds_Integration(t *testing.T) {
	if os.Getenv("RUN_INTEGRATION_TESTS") == "" {
		t.Skip("跳过集成测试 (设置 RUN_INTEGRATION_TESTS=1 来运行)")
	}

	baseClient := createTestClient(t)
	defer baseClient.Close()

	feedsClient := feeds.NewClient(baseClient)
	ctx := context.Background()

	t.Run("CreateFeedDocument", func(t *testing.T) {
		request := map[string]interface{}{
			"contentType": "text/xml; charset=UTF-8",
		}

		result, err := feedsClient.CreateFeedDocument(ctx, request)
		if err != nil {
			t.Errorf("CreateFeedDocument failed: %v", err)
		}
		if result == nil {
			t.Error("CreateFeedDocument returned nil")
		}
		t.Logf("✓ CreateFeedDocument success")
	})

	t.Run("GetFeeds", func(t *testing.T) {
		params := map[string]string{
			"feedTypes": "POST_PRODUCT_DATA",
			"pageSize":  "10",
		}

		result, err := feedsClient.GetFeeds(ctx, params)
		if err != nil {
			t.Errorf("GetFeeds failed: %v", err)
		}
		if result == nil {
			t.Error("GetFeeds returned nil")
		}
		t.Logf("✓ GetFeeds success")
	})
}

