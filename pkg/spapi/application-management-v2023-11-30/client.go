// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package application_management_v2023_11_30

import (
	"context"
	"fmt"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client application-management API v2023-11-30
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// RotateApplicationClientSecret
// Method: POST | Path: /applications/2023-11-30/clientSecret
func (c *Client) RotateApplicationClientSecret(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/applications/2023-11-30/clientSecret"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("RotateApplicationClientSecret: %w", err)
	}
	return result, nil
}
