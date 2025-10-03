// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package product_type_definitions_v2020_09_01

import (
	"context"
	"fmt"
	"strings"
	
	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client product-type-definitions API v2020-09-01
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// GetDefinitionsProductType 
// Method: GET | Path: /definitions/2020-09-01/productTypes/{productType}
func (c *Client) GetDefinitionsProductType(ctx context.Context, productType string, query map[string]string) (interface{}, error) {
	path := "/definitions/2020-09-01/productTypes/{productType}"
	path = strings.Replace(path, "{productType}", productType, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil { return nil, fmt.Errorf("GetDefinitionsProductType: %w", err) }
	return result, nil
}

// SearchDefinitionsProductTypes 
// Method: GET | Path: /definitions/2020-09-01/productTypes
func (c *Client) SearchDefinitionsProductTypes(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/definitions/2020-09-01/productTypes"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil { return nil, fmt.Errorf("SearchDefinitionsProductTypes: %w", err) }
	return result, nil
}
