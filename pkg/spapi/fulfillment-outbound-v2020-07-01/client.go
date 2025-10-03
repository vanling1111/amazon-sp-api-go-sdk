// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package fulfillment_outbound_v2020_07_01

import (
	"context"
	"fmt"
	"strings"
	
	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client fulfillment-outbound API v2020-07-01
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// CreateFulfillmentReturn 
// Method: PUT | Path: /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/return
func (c *Client) CreateFulfillmentReturn(ctx context.Context, sellerFulfillmentOrderId string, body interface{}) (interface{}, error) {
	path := "/fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/return"
	path = strings.Replace(path, "{sellerFulfillmentOrderId}", sellerFulfillmentOrderId, 1)
	var result interface{}
	err := c.baseClient.Put(ctx, path, body, &result)
	if err != nil { return nil, fmt.Errorf("CreateFulfillmentReturn: %w", err) }
	return result, nil
}

// CancelFulfillmentOrder 
// Method: PUT | Path: /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/cancel
func (c *Client) CancelFulfillmentOrder(ctx context.Context, sellerFulfillmentOrderId string, body interface{}) (interface{}, error) {
	path := "/fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/cancel"
	path = strings.Replace(path, "{sellerFulfillmentOrderId}", sellerFulfillmentOrderId, 1)
	var result interface{}
	err := c.baseClient.Put(ctx, path, body, &result)
	if err != nil { return nil, fmt.Errorf("CancelFulfillmentOrder: %w", err) }
	return result, nil
}

// GetFeatures 
// Method: GET | Path: /fba/outbound/2020-07-01/features
func (c *Client) GetFeatures(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/fba/outbound/2020-07-01/features"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil { return nil, fmt.Errorf("GetFeatures: %w", err) }
	return result, nil
}

// CreateFulfillmentOrder 
// Method: POST | Path: /fba/outbound/2020-07-01/fulfillmentOrders
func (c *Client) CreateFulfillmentOrder(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/fba/outbound/2020-07-01/fulfillmentOrders"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil { return nil, fmt.Errorf("CreateFulfillmentOrder: %w", err) }
	return result, nil
}

// ListAllFulfillmentOrders 
// Method: GET | Path: /fba/outbound/2020-07-01/fulfillmentOrders
func (c *Client) ListAllFulfillmentOrders(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/fba/outbound/2020-07-01/fulfillmentOrders"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil { return nil, fmt.Errorf("ListAllFulfillmentOrders: %w", err) }
	return result, nil
}

// SubmitFulfillmentOrderStatusUpdate 
// Method: PUT | Path: /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/status
func (c *Client) SubmitFulfillmentOrderStatusUpdate(ctx context.Context, sellerFulfillmentOrderId string, body interface{}) (interface{}, error) {
	path := "/fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/status"
	path = strings.Replace(path, "{sellerFulfillmentOrderId}", sellerFulfillmentOrderId, 1)
	var result interface{}
	err := c.baseClient.Put(ctx, path, body, &result)
	if err != nil { return nil, fmt.Errorf("SubmitFulfillmentOrderStatusUpdate: %w", err) }
	return result, nil
}

// GetFulfillmentPreview 
// Method: POST | Path: /fba/outbound/2020-07-01/fulfillmentOrders/preview
func (c *Client) GetFulfillmentPreview(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/fba/outbound/2020-07-01/fulfillmentOrders/preview"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil { return nil, fmt.Errorf("GetFulfillmentPreview: %w", err) }
	return result, nil
}

// GetFeatureSKU 
// Method: GET | Path: /fba/outbound/2020-07-01/features/inventory/{featureName}/{sellerSku}
func (c *Client) GetFeatureSKU(ctx context.Context, featureName string, sellerSku string, query map[string]string) (interface{}, error) {
	path := "/fba/outbound/2020-07-01/features/inventory/{featureName}/{sellerSku}"
	path = strings.Replace(path, "{featureName}", featureName, 1)
	path = strings.Replace(path, "{sellerSku}", sellerSku, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil { return nil, fmt.Errorf("GetFeatureSKU: %w", err) }
	return result, nil
}

// GetFulfillmentOrder 
// Method: GET | Path: /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}
func (c *Client) GetFulfillmentOrder(ctx context.Context, sellerFulfillmentOrderId string, query map[string]string) (interface{}, error) {
	path := "/fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}"
	path = strings.Replace(path, "{sellerFulfillmentOrderId}", sellerFulfillmentOrderId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil { return nil, fmt.Errorf("GetFulfillmentOrder: %w", err) }
	return result, nil
}

// ListReturnReasonCodes 
// Method: GET | Path: /fba/outbound/2020-07-01/returnReasonCodes
func (c *Client) ListReturnReasonCodes(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/fba/outbound/2020-07-01/returnReasonCodes"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil { return nil, fmt.Errorf("ListReturnReasonCodes: %w", err) }
	return result, nil
}

// GetFeatureInventory 
// Method: GET | Path: /fba/outbound/2020-07-01/features/inventory/{featureName}
func (c *Client) GetFeatureInventory(ctx context.Context, featureName string, query map[string]string) (interface{}, error) {
	path := "/fba/outbound/2020-07-01/features/inventory/{featureName}"
	path = strings.Replace(path, "{featureName}", featureName, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil { return nil, fmt.Errorf("GetFeatureInventory: %w", err) }
	return result, nil
}

// DeliveryOffers 
// Method: POST | Path: /fba/outbound/2020-07-01/deliveryOffers
func (c *Client) DeliveryOffers(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/fba/outbound/2020-07-01/deliveryOffers"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil { return nil, fmt.Errorf("DeliveryOffers: %w", err) }
	return result, nil
}

// GetPackageTrackingDetails 
// Method: GET | Path: /fba/outbound/2020-07-01/tracking
func (c *Client) GetPackageTrackingDetails(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/fba/outbound/2020-07-01/tracking"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil { return nil, fmt.Errorf("GetPackageTrackingDetails: %w", err) }
	return result, nil
}

// UpdateFulfillmentOrder 
// Method: PUT | Path: /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}
func (c *Client) UpdateFulfillmentOrder(ctx context.Context, sellerFulfillmentOrderId string, body interface{}) (interface{}, error) {
	path := "/fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}"
	path = strings.Replace(path, "{sellerFulfillmentOrderId}", sellerFulfillmentOrderId, 1)
	var result interface{}
	err := c.baseClient.Put(ctx, path, body, &result)
	if err != nil { return nil, fmt.Errorf("UpdateFulfillmentOrder: %w", err) }
	return result, nil
}
