// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package fulfillment_inbound_v0

import (
	"context"
	"fmt"
	"strings"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client fulfillment-inbound API v0
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// GetShipmentItems
// Method: GET | Path: /fba/inbound/v0/shipmentItems
func (c *Client) GetShipmentItems(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/fba/inbound/v0/shipmentItems"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetShipmentItems: %w", err)
	}
	return result, nil
}

// GetLabels
// Method: GET | Path: /fba/inbound/v0/shipments/{shipmentId}/labels
func (c *Client) GetLabels(ctx context.Context, shipmentId string, query map[string]string) (interface{}, error) {
	path := "/fba/inbound/v0/shipments/{shipmentId}/labels"
	path = strings.Replace(path, "{shipmentId}", shipmentId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetLabels: %w", err)
	}
	return result, nil
}

// GetPrepInstructions
// Method: GET | Path: /fba/inbound/v0/prepInstructions
func (c *Client) GetPrepInstructions(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/fba/inbound/v0/prepInstructions"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetPrepInstructions: %w", err)
	}
	return result, nil
}

// GetShipments
// Method: GET | Path: /fba/inbound/v0/shipments
func (c *Client) GetShipments(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/fba/inbound/v0/shipments"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetShipments: %w", err)
	}
	return result, nil
}

// GetBillOfLading
// Method: GET | Path: /fba/inbound/v0/shipments/{shipmentId}/billOfLading
func (c *Client) GetBillOfLading(ctx context.Context, shipmentId string, query map[string]string) (interface{}, error) {
	path := "/fba/inbound/v0/shipments/{shipmentId}/billOfLading"
	path = strings.Replace(path, "{shipmentId}", shipmentId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetBillOfLading: %w", err)
	}
	return result, nil
}

// GetShipmentItemsByShipmentId
// Method: GET | Path: /fba/inbound/v0/shipments/{shipmentId}/items
func (c *Client) GetShipmentItemsByShipmentId(ctx context.Context, shipmentId string, query map[string]string) (interface{}, error) {
	path := "/fba/inbound/v0/shipments/{shipmentId}/items"
	path = strings.Replace(path, "{shipmentId}", shipmentId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetShipmentItemsByShipmentId: %w", err)
	}
	return result, nil
}
