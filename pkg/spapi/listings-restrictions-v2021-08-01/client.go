// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package listings_restrictions_v2021_08_01

import (
	"context"
	"fmt"
	
	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client listings-restrictions API v2021-08-01
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// GetListingsRestrictions 
// Method: GET | Path: /listings/2021-08-01/restrictions
func (c *Client) GetListingsRestrictions(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/listings/2021-08-01/restrictions"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil { return nil, fmt.Errorf("GetListingsRestrictions: %w", err) }
	return result, nil
}
