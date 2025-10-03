// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package finances_v2024_06_01_transfers

import (
	"context"
	"fmt"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client finances API v2024-06-01-transfers
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// InitiatePayout
// Method: POST | Path: /finances/transfers/2024-06-01/payouts
func (c *Client) InitiatePayout(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/finances/transfers/2024-06-01/payouts"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("InitiatePayout: %w", err)
	}
	return result, nil
}

// GetPaymentMethods
// Method: GET | Path: /finances/transfers/2024-06-01/paymentMethods
func (c *Client) GetPaymentMethods(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/finances/transfers/2024-06-01/paymentMethods"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetPaymentMethods: %w", err)
	}
	return result, nil
}
