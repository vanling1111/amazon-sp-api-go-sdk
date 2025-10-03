// Copyright 2025 Amazon SP-API Go SDK Authors.
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

