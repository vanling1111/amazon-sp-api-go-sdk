// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package merchant_fulfillment_v0

import (
	"context"
	"fmt"
	"strings"
	
	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client merchant-fulfillment API v0
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// GetShipment 
// Method: GET | Path: /mfn/v0/shipments/{shipmentId}
func (c *Client) GetShipment(ctx context.Context, shipmentId string, query map[string]string) (interface{}, error) {
	path := "/mfn/v0/shipments/{shipmentId}"
	path = strings.Replace(path, "{shipmentId}", shipmentId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil { return nil, fmt.Errorf("GetShipment: %w", err) }
	return result, nil
}

// GetAdditionalSellerInputs 
// Method: POST | Path: /mfn/v0/additionalSellerInputs
func (c *Client) GetAdditionalSellerInputs(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/mfn/v0/additionalSellerInputs"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil { return nil, fmt.Errorf("GetAdditionalSellerInputs: %w", err) }
	return result, nil
}

// CancelShipment 
// Method: DELETE | Path: /mfn/v0/shipments/{shipmentId}
func (c *Client) CancelShipment(ctx context.Context, shipmentId string) (interface{}, error) {
	path := "/mfn/v0/shipments/{shipmentId}"
	path = strings.Replace(path, "{shipmentId}", shipmentId, 1)
	var result interface{}
	err := c.baseClient.Delete(ctx, path, &result)
	if err != nil { return nil, fmt.Errorf("CancelShipment: %w", err) }
	return result, nil
}

// GetEligibleShipmentServices 
// Method: POST | Path: /mfn/v0/eligibleShippingServices
func (c *Client) GetEligibleShipmentServices(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/mfn/v0/eligibleShippingServices"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil { return nil, fmt.Errorf("GetEligibleShipmentServices: %w", err) }
	return result, nil
}

// CreateShipment 
// Method: POST | Path: /mfn/v0/shipments
func (c *Client) CreateShipment(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/mfn/v0/shipments"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil { return nil, fmt.Errorf("CreateShipment: %w", err) }
	return result, nil
}
