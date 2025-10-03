// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package vendor_direct_fulfillment_payments_v1

import (
	"context"
	"fmt"
	
	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client vendor-direct-fulfillment-payments API v1
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// SubmitInvoice 
// Method: POST | Path: /vendor/directFulfillment/payments/v1/invoices
func (c *Client) SubmitInvoice(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/vendor/directFulfillment/payments/v1/invoices"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil { return nil, fmt.Errorf("SubmitInvoice: %w", err) }
	return result, nil
}
