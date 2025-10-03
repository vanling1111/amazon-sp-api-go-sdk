// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package catalog_items_v2020_12_01

import (
	"context"
	"fmt"
	"strings"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client catalog-items API v2020-12-01
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// SearchCatalogItems
// Method: GET | Path: /catalog/2020-12-01/items
func (c *Client) SearchCatalogItems(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/catalog/2020-12-01/items"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("SearchCatalogItems: %w", err)
	}
	return result, nil
}

// GetCatalogItem
// Method: GET | Path: /catalog/2020-12-01/items/{asin}
func (c *Client) GetCatalogItem(ctx context.Context, asin string, query map[string]string) (interface{}, error) {
	path := "/catalog/2020-12-01/items/{asin}"
	path = strings.Replace(path, "{asin}", asin, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetCatalogItem: %w", err)
	}
	return result, nil
}
