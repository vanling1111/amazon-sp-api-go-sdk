// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package catalog_items_v0

import (
	"context"
	"fmt"
	
	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client catalog-items API v0
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// ListCatalogCategories 
// Method: GET | Path: /catalog/v0/categories
func (c *Client) ListCatalogCategories(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/catalog/v0/categories"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil { return nil, fmt.Errorf("ListCatalogCategories: %w", err) }
	return result, nil
}
