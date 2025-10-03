// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package vendor_direct_fulfillment_inventory_v1

import (
	"context"
	"fmt"
	"strings"
	
	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client vendor-direct-fulfillment-inventory API v1
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// SubmitInventoryUpdate 
// Method: POST | Path: /vendor/directFulfillment/inventory/v1/warehouses/{warehouseId}/items
func (c *Client) SubmitInventoryUpdate(ctx context.Context, warehouseId string, body interface{}) (interface{}, error) {
	path := "/vendor/directFulfillment/inventory/v1/warehouses/{warehouseId}/items"
	path = strings.Replace(path, "{warehouseId}", warehouseId, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil { return nil, fmt.Errorf("SubmitInventoryUpdate: %w", err) }
	return result, nil
}
