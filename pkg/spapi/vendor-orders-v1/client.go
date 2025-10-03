// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package vendor_orders_v1

import (
	"context"
	"fmt"
	"strings"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client vendor-orders API v1
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// GetPurchaseOrder
// Method: GET | Path: /vendor/orders/v1/purchaseOrders/{purchaseOrderNumber}
func (c *Client) GetPurchaseOrder(ctx context.Context, purchaseOrderNumber string, query map[string]string) (interface{}, error) {
	path := "/vendor/orders/v1/purchaseOrders/{purchaseOrderNumber}"
	path = strings.Replace(path, "{purchaseOrderNumber}", purchaseOrderNumber, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetPurchaseOrder: %w", err)
	}
	return result, nil
}

// GetPurchaseOrders
// Method: GET | Path: /vendor/orders/v1/purchaseOrders
func (c *Client) GetPurchaseOrders(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/vendor/orders/v1/purchaseOrders"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetPurchaseOrders: %w", err)
	}
	return result, nil
}

// GetPurchaseOrdersStatus
// Method: GET | Path: /vendor/orders/v1/purchaseOrdersStatus
func (c *Client) GetPurchaseOrdersStatus(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/vendor/orders/v1/purchaseOrdersStatus"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetPurchaseOrdersStatus: %w", err)
	}
	return result, nil
}

// SubmitAcknowledgement
// Method: POST | Path: /vendor/orders/v1/acknowledgements
func (c *Client) SubmitAcknowledgement(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/vendor/orders/v1/acknowledgements"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("SubmitAcknowledgement: %w", err)
	}
	return result, nil
}
