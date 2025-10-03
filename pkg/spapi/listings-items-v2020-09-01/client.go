// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package listings_items_v2020_09_01

import (
	"context"
	"fmt"
	"strings"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client listings-items API v2020-09-01
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// PutListingsItem
// Method: PUT | Path: /listings/2020-09-01/items/{sellerId}/{sku}
func (c *Client) PutListingsItem(ctx context.Context, sellerId string, sku string, body interface{}) (interface{}, error) {
	path := "/listings/2020-09-01/items/{sellerId}/{sku}"
	path = strings.Replace(path, "{sellerId}", sellerId, 1)
	path = strings.Replace(path, "{sku}", sku, 1)
	var result interface{}
	err := c.baseClient.Put(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("PutListingsItem: %w", err)
	}
	return result, nil
}

// PatchListingsItem
// Method: PATCH | Path: /listings/2020-09-01/items/{sellerId}/{sku}
func (c *Client) PatchListingsItem(ctx context.Context, sellerId string, sku string, body interface{}) (interface{}, error) {
	path := "/listings/2020-09-01/items/{sellerId}/{sku}"
	path = strings.Replace(path, "{sellerId}", sellerId, 1)
	path = strings.Replace(path, "{sku}", sku, 1)
	var result interface{}
	err := c.baseClient.DoRequest(ctx, "PATCH", path, nil, body, &result)
	if err != nil {
		return nil, fmt.Errorf("PatchListingsItem: %w", err)
	}
	return result, nil
}

// DeleteListingsItem
// Method: DELETE | Path: /listings/2020-09-01/items/{sellerId}/{sku}
func (c *Client) DeleteListingsItem(ctx context.Context, sellerId string, sku string) (interface{}, error) {
	path := "/listings/2020-09-01/items/{sellerId}/{sku}"
	path = strings.Replace(path, "{sellerId}", sellerId, 1)
	path = strings.Replace(path, "{sku}", sku, 1)
	var result interface{}
	err := c.baseClient.Delete(ctx, path, &result)
	if err != nil {
		return nil, fmt.Errorf("DeleteListingsItem: %w", err)
	}
	return result, nil
}
