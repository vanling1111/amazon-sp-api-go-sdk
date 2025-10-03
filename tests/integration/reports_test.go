// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

// +build integration

package integration

import (
	"context"
	"os"
	"testing"
	"time"

	reports "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi/reports-v2021-06-30"
)

func TestReports_Integration(t *testing.T) {
	if os.Getenv("RUN_INTEGRATION_TESTS") == "" {
		t.Skip("跳过集成测试 (设置 RUN_INTEGRATION_TESTS=1 来运行)")
	}

	baseClient := createTestClient(t)
	defer baseClient.Close()

	reportsClient := reports.NewClient(baseClient)
	ctx := context.Background()

	t.Run("CreateReport", func(t *testing.T) {
		request := map[string]interface{}{
			"reportType":     "GET_MERCHANT_LISTINGS_ALL_DATA",
			"marketplaceIds": []string{"ATVPDKIKX0DER"},
		}

		result, err := reportsClient.CreateReport(ctx, request)
		if err != nil {
			t.Errorf("CreateReport failed: %v", err)
		}
		if result == nil {
			t.Error("CreateReport returned nil")
		}
		t.Logf("✓ CreateReport success")
	})

	t.Run("GetReports", func(t *testing.T) {
		params := map[string]string{
			"reportTypes":  "GET_MERCHANT_LISTINGS_ALL_DATA",
			"createdSince": time.Now().Add(-30 * 24 * time.Hour).Format(time.RFC3339),
			"pageSize":     "10",
		}

		result, err := reportsClient.GetReports(ctx, params)
		if err != nil {
			t.Errorf("GetReports failed: %v", err)
		}
		if result == nil {
			t.Error("GetReports returned nil")
		}
		t.Logf("✓ GetReports success")
	})
}

