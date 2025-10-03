// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package sales_v1

import (
	"context"
	"fmt"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client sales API v1
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// GetOrderMetrics
// Method: GET | Path: /sales/v1/orderMetrics
func (c *Client) GetOrderMetrics(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/sales/v1/orderMetrics"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetOrderMetrics: %w", err)
	}
	return result, nil
}
