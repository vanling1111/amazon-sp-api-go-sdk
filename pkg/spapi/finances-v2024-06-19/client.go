// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package finances_v2024_06_19

import (
	"context"
	"fmt"
	
	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client finances API v2024-06-19
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// ListTransactions 
// Method: GET | Path: /finances/2024-06-19/transactions
func (c *Client) ListTransactions(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/finances/2024-06-19/transactions"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil { return nil, fmt.Errorf("ListTransactions: %w", err) }
	return result, nil
}
