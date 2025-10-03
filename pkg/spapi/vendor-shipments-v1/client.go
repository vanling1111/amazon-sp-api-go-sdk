// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package vendor_shipments_v1

import (
	"context"
	"fmt"
	
	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client vendor-shipments API v1
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// SubmitShipments SubmitShipments
// Method: POST | Path: /vendor/shipping/v1/shipments
func (c *Client) SubmitShipments(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/vendor/shipping/v1/shipments"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil { return nil, fmt.Errorf("SubmitShipments: %w", err) }
	return result, nil
}

// GetShipmentLabels 
// Method: GET | Path: /vendor/shipping/v1/transportLabels
func (c *Client) GetShipmentLabels(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/vendor/shipping/v1/transportLabels"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil { return nil, fmt.Errorf("GetShipmentLabels: %w", err) }
	return result, nil
}

// SubmitShipmentConfirmations SubmitShipmentConfirmations
// Method: POST | Path: /vendor/shipping/v1/shipmentConfirmations
func (c *Client) SubmitShipmentConfirmations(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/vendor/shipping/v1/shipmentConfirmations"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil { return nil, fmt.Errorf("SubmitShipmentConfirmations: %w", err) }
	return result, nil
}

// GetShipmentDetails GetShipmentDetails
// Method: GET | Path: /vendor/shipping/v1/shipments
func (c *Client) GetShipmentDetails(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/vendor/shipping/v1/shipments"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil { return nil, fmt.Errorf("GetShipmentDetails: %w", err) }
	return result, nil
}
