// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package orders_v0

import (
	"context"
	"fmt"
	"strings"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client orders API v0
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// GetOrderItemsBuyerInfo
// Method: GET | Path: /orders/v0/orders/{orderId}/orderItems/buyerInfo
func (c *Client) GetOrderItemsBuyerInfo(ctx context.Context, orderId string, query map[string]string) (interface{}, error) {
	path := "/orders/v0/orders/{orderId}/orderItems/buyerInfo"
	path = strings.Replace(path, "{orderId}", orderId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetOrderItemsBuyerInfo: %w", err)
	}
	return result, nil
}

// GetOrder
// Method: GET | Path: /orders/v0/orders/{orderId}
func (c *Client) GetOrder(ctx context.Context, orderId string, query map[string]string) (interface{}, error) {
	path := "/orders/v0/orders/{orderId}"
	path = strings.Replace(path, "{orderId}", orderId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetOrder: %w", err)
	}
	return result, nil
}

// GetOrders
// Method: GET | Path: /orders/v0/orders
func (c *Client) GetOrders(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/orders/v0/orders"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetOrders: %w", err)
	}
	return result, nil
}

// GetOrderBuyerInfo
// Method: GET | Path: /orders/v0/orders/{orderId}/buyerInfo
func (c *Client) GetOrderBuyerInfo(ctx context.Context, orderId string, query map[string]string) (interface{}, error) {
	path := "/orders/v0/orders/{orderId}/buyerInfo"
	path = strings.Replace(path, "{orderId}", orderId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetOrderBuyerInfo: %w", err)
	}
	return result, nil
}

// GetOrderItems
// Method: GET | Path: /orders/v0/orders/{orderId}/orderItems
func (c *Client) GetOrderItems(ctx context.Context, orderId string, query map[string]string) (interface{}, error) {
	path := "/orders/v0/orders/{orderId}/orderItems"
	path = strings.Replace(path, "{orderId}", orderId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetOrderItems: %w", err)
	}
	return result, nil
}

// GetOrderRegulatedInfo
// Method: GET | Path: /orders/v0/orders/{orderId}/regulatedInfo
func (c *Client) GetOrderRegulatedInfo(ctx context.Context, orderId string, query map[string]string) (interface{}, error) {
	path := "/orders/v0/orders/{orderId}/regulatedInfo"
	path = strings.Replace(path, "{orderId}", orderId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetOrderRegulatedInfo: %w", err)
	}
	return result, nil
}

// GetOrderAddress
// Method: GET | Path: /orders/v0/orders/{orderId}/address
func (c *Client) GetOrderAddress(ctx context.Context, orderId string, query map[string]string) (interface{}, error) {
	path := "/orders/v0/orders/{orderId}/address"
	path = strings.Replace(path, "{orderId}", orderId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetOrderAddress: %w", err)
	}
	return result, nil
}

// ConfirmShipment
// Method: POST | Path: /orders/v0/orders/{orderId}/shipmentConfirmation
func (c *Client) ConfirmShipment(ctx context.Context, orderId string, body interface{}) (interface{}, error) {
	path := "/orders/v0/orders/{orderId}/shipmentConfirmation"
	path = strings.Replace(path, "{orderId}", orderId, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("ConfirmShipment: %w", err)
	}
	return result, nil
}

// UpdateVerificationStatus
// Method: PATCH | Path: /orders/v0/orders/{orderId}/regulatedInfo
func (c *Client) UpdateVerificationStatus(ctx context.Context, orderId string, body interface{}) (interface{}, error) {
	path := "/orders/v0/orders/{orderId}/regulatedInfo"
	path = strings.Replace(path, "{orderId}", orderId, 1)
	var result interface{}
	err := c.baseClient.DoRequest(ctx, "PATCH", path, nil, body, &result)
	if err != nil {
		return nil, fmt.Errorf("UpdateVerificationStatus: %w", err)
	}
	return result, nil
}

// UpdateShipmentStatus
// Method: POST | Path: /orders/v0/orders/{orderId}/shipment
func (c *Client) UpdateShipmentStatus(ctx context.Context, orderId string, body interface{}) (interface{}, error) {
	path := "/orders/v0/orders/{orderId}/shipment"
	path = strings.Replace(path, "{orderId}", orderId, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("UpdateShipmentStatus: %w", err)
	}
	return result, nil
}
