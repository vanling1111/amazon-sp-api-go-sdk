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
	"os"
	"testing"

	"github.com/vanling1111/amazon-sp-api-go-sdk/internal/models"
	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// createTestClient 创建用于集成测试的客户端
func createTestClient(t *testing.T) *spapi.Client {
	t.Helper()

	clientID := os.Getenv("SP_API_CLIENT_ID")
	clientSecret := os.Getenv("SP_API_CLIENT_SECRET")
	refreshToken := os.Getenv("SP_API_REFRESH_TOKEN")

	if clientID == "" || clientSecret == "" || refreshToken == "" {
		t.Fatal("缺少必要的环境变量: SP_API_CLIENT_ID, SP_API_CLIENT_SECRET, SP_API_REFRESH_TOKEN")
	}

	client, err := spapi.NewClient(
		spapi.WithRegion(models.RegionNA),
		spapi.WithCredentials(clientID, clientSecret, refreshToken),
	)
	if err != nil {
		t.Fatalf("创建客户端失败: %v", err)
	}

	return client
}

