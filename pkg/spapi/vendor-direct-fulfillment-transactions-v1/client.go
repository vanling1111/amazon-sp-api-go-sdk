// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package vendor_direct_fulfillment_transactions_v1

import (
	"context"
	"fmt"
	"strings"
	
	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client vendor-direct-fulfillment-transactions API v1
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// GetTransactionStatus 
// Method: GET | Path: /vendor/directFulfillment/transactions/v1/transactions/{transactionId}
func (c *Client) GetTransactionStatus(ctx context.Context, transactionId string, query map[string]string) (interface{}, error) {
	path := "/vendor/directFulfillment/transactions/v1/transactions/{transactionId}"
	path = strings.Replace(path, "{transactionId}", transactionId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil { return nil, fmt.Errorf("GetTransactionStatus: %w", err) }
	return result, nil
}
