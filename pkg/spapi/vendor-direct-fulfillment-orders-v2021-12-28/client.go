// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package vendor_direct_fulfillment_orders_v2021_12_28

import (
	"context"
	"fmt"
	"strings"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client vendor-direct-fulfillment-orders API v2021-12-28
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// GetOrder
// Method: GET | Path: /vendor/directFulfillment/orders/2021-12-28/purchaseOrders/{purchaseOrderNumber}
func (c *Client) GetOrder(ctx context.Context, purchaseOrderNumber string, query map[string]string) (interface{}, error) {
	path := "/vendor/directFulfillment/orders/2021-12-28/purchaseOrders/{purchaseOrderNumber}"
	path = strings.Replace(path, "{purchaseOrderNumber}", purchaseOrderNumber, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetOrder: %w", err)
	}
	return result, nil
}

// GetOrders
// Method: GET | Path: /vendor/directFulfillment/orders/2021-12-28/purchaseOrders
func (c *Client) GetOrders(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/vendor/directFulfillment/orders/2021-12-28/purchaseOrders"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetOrders: %w", err)
	}
	return result, nil
}

// SubmitAcknowledgement
// Method: POST | Path: /vendor/directFulfillment/orders/2021-12-28/acknowledgements
func (c *Client) SubmitAcknowledgement(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/vendor/directFulfillment/orders/2021-12-28/acknowledgements"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("SubmitAcknowledgement: %w", err)
	}
	return result, nil
}
