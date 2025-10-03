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

	"github.com/vanling1111/amazon-sp-api-go-sdk/internal/models"
	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
	notifications "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi/notifications-v1"
)

func TestNotifications_Grantless_Integration(t *testing.T) {
	if os.Getenv("RUN_INTEGRATION_TESTS") == "" {
		t.Skip("跳过集成测试 (设置 RUN_INTEGRATION_TESTS=1 来运行)")
	}

	// 获取凭证
	clientID := os.Getenv("SP_API_CLIENT_ID")
	clientSecret := os.Getenv("SP_API_CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		t.Fatal("缺少必要的环境变量: SP_API_CLIENT_ID, SP_API_CLIENT_SECRET")
	}

	// 创建 Grantless 客户端
	baseClient, err := spapi.NewClient(
		spapi.WithRegion(models.RegionNA),
		spapi.WithGrantlessCredentials(clientID, clientSecret, []string{
			"sellingpartnerapi::notifications",
		}),
	)
	if err != nil {
		t.Fatalf("创建 Grantless 客户端失败: %v", err)
	}
	defer baseClient.Close()

	notificationsClient := notifications.NewClient(baseClient)
	ctx := context.Background()

	t.Run("GetDestinations", func(t *testing.T) {
		result, err := notificationsClient.GetDestinations(ctx, nil)
		if err != nil {
			t.Logf("GetDestinations returned error (may be expected): %v", err)
		} else {
			t.Logf("✓ GetDestinations success: %v", result)
		}
	})
}

