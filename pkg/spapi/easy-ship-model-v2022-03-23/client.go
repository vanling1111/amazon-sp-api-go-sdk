// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package easy_ship_model_v2022_03_23

import (
	"context"
	"fmt"
	
	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client easy-ship-model API v2022-03-23
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// UpdateScheduledPackages 
// Method: PATCH | Path: /easyShip/2022-03-23/package
func (c *Client) UpdateScheduledPackages(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/easyShip/2022-03-23/package"
	var result interface{}
	err := c.baseClient.DoRequest(ctx, "PATCH", path, nil, body, &result)
	if err != nil { return nil, fmt.Errorf("UpdateScheduledPackages: %w", err) }
	return result, nil
}

// CreateScheduledPackageBulk 
// Method: POST | Path: /easyShip/2022-03-23/packages/bulk
func (c *Client) CreateScheduledPackageBulk(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/easyShip/2022-03-23/packages/bulk"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil { return nil, fmt.Errorf("CreateScheduledPackageBulk: %w", err) }
	return result, nil
}

// CreateScheduledPackage 
// Method: POST | Path: /easyShip/2022-03-23/package
func (c *Client) CreateScheduledPackage(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/easyShip/2022-03-23/package"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil { return nil, fmt.Errorf("CreateScheduledPackage: %w", err) }
	return result, nil
}

// ListHandoverSlots 
// Method: POST | Path: /easyShip/2022-03-23/timeSlot
func (c *Client) ListHandoverSlots(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/easyShip/2022-03-23/timeSlot"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil { return nil, fmt.Errorf("ListHandoverSlots: %w", err) }
	return result, nil
}

// GetScheduledPackage 
// Method: GET | Path: /easyShip/2022-03-23/package
func (c *Client) GetScheduledPackage(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/easyShip/2022-03-23/package"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil { return nil, fmt.Errorf("GetScheduledPackage: %w", err) }
	return result, nil
}
