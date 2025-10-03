// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package vendor_direct_fulfillment_shipping_v1

import (
	"context"
	"fmt"
	"strings"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client vendor-direct-fulfillment-shipping API v1
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// SubmitShippingLabelRequest
// Method: POST | Path: /vendor/directFulfillment/shipping/v1/shippingLabels
func (c *Client) SubmitShippingLabelRequest(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/vendor/directFulfillment/shipping/v1/shippingLabels"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("SubmitShippingLabelRequest: %w", err)
	}
	return result, nil
}

// SubmitShipmentConfirmations
// Method: POST | Path: /vendor/directFulfillment/shipping/v1/shipmentConfirmations
func (c *Client) SubmitShipmentConfirmations(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/vendor/directFulfillment/shipping/v1/shipmentConfirmations"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("SubmitShipmentConfirmations: %w", err)
	}
	return result, nil
}

// GetShippingLabels
// Method: GET | Path: /vendor/directFulfillment/shipping/v1/shippingLabels
func (c *Client) GetShippingLabels(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/vendor/directFulfillment/shipping/v1/shippingLabels"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetShippingLabels: %w", err)
	}
	return result, nil
}

// GetPackingSlips
// Method: GET | Path: /vendor/directFulfillment/shipping/v1/packingSlips
func (c *Client) GetPackingSlips(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/vendor/directFulfillment/shipping/v1/packingSlips"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetPackingSlips: %w", err)
	}
	return result, nil
}

// GetShippingLabel
// Method: GET | Path: /vendor/directFulfillment/shipping/v1/shippingLabels/{purchaseOrderNumber}
func (c *Client) GetShippingLabel(ctx context.Context, purchaseOrderNumber string, query map[string]string) (interface{}, error) {
	path := "/vendor/directFulfillment/shipping/v1/shippingLabels/{purchaseOrderNumber}"
	path = strings.Replace(path, "{purchaseOrderNumber}", purchaseOrderNumber, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetShippingLabel: %w", err)
	}
	return result, nil
}

// GetCustomerInvoices
// Method: GET | Path: /vendor/directFulfillment/shipping/v1/customerInvoices
func (c *Client) GetCustomerInvoices(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/vendor/directFulfillment/shipping/v1/customerInvoices"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetCustomerInvoices: %w", err)
	}
	return result, nil
}

// GetCustomerInvoice
// Method: GET | Path: /vendor/directFulfillment/shipping/v1/customerInvoices/{purchaseOrderNumber}
func (c *Client) GetCustomerInvoice(ctx context.Context, purchaseOrderNumber string, query map[string]string) (interface{}, error) {
	path := "/vendor/directFulfillment/shipping/v1/customerInvoices/{purchaseOrderNumber}"
	path = strings.Replace(path, "{purchaseOrderNumber}", purchaseOrderNumber, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetCustomerInvoice: %w", err)
	}
	return result, nil
}

// GetPackingSlip
// Method: GET | Path: /vendor/directFulfillment/shipping/v1/packingSlips/{purchaseOrderNumber}
func (c *Client) GetPackingSlip(ctx context.Context, purchaseOrderNumber string, query map[string]string) (interface{}, error) {
	path := "/vendor/directFulfillment/shipping/v1/packingSlips/{purchaseOrderNumber}"
	path = strings.Replace(path, "{purchaseOrderNumber}", purchaseOrderNumber, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetPackingSlip: %w", err)
	}
	return result, nil
}

// SubmitShipmentStatusUpdates
// Method: POST | Path: /vendor/directFulfillment/shipping/v1/shipmentStatusUpdates
func (c *Client) SubmitShipmentStatusUpdates(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/vendor/directFulfillment/shipping/v1/shipmentStatusUpdates"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("SubmitShipmentStatusUpdates: %w", err)
	}
	return result, nil
}
