// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package vehicles_v2024_11_01

import (
	"context"
	"fmt"
	
	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client vehicles API v2024-11-01
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// GetVehicles 
// Method: GET | Path: /catalog/2024-11-01/automotive/vehicles
func (c *Client) GetVehicles(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/catalog/2024-11-01/automotive/vehicles"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil { return nil, fmt.Errorf("GetVehicles: %w", err) }
	return result, nil
}
