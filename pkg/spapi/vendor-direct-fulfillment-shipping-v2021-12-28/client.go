// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package vendor_direct_fulfillment_shipping_v2021_12_28

import (
	"context"
	"fmt"
	"strings"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client vendor-direct-fulfillment-shipping API v2021-12-28
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// SubmitShippingLabelRequest submitShippingLabelRequest
// Method: POST | Path: /vendor/directFulfillment/shipping/2021-12-28/shippingLabels
func (c *Client) SubmitShippingLabelRequest(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/vendor/directFulfillment/shipping/2021-12-28/shippingLabels"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("SubmitShippingLabelRequest: %w", err)
	}
	return result, nil
}

// SubmitShipmentConfirmations submitShipmentConfirmations
// Method: POST | Path: /vendor/directFulfillment/shipping/2021-12-28/shipmentConfirmations
func (c *Client) SubmitShipmentConfirmations(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/vendor/directFulfillment/shipping/2021-12-28/shipmentConfirmations"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("SubmitShipmentConfirmations: %w", err)
	}
	return result, nil
}

// GetShippingLabels getShippingLabels
// Method: GET | Path: /vendor/directFulfillment/shipping/2021-12-28/shippingLabels
func (c *Client) GetShippingLabels(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/vendor/directFulfillment/shipping/2021-12-28/shippingLabels"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetShippingLabels: %w", err)
	}
	return result, nil
}

// GetPackingSlips getPackingSlips
// Method: GET | Path: /vendor/directFulfillment/shipping/2021-12-28/packingSlips
func (c *Client) GetPackingSlips(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/vendor/directFulfillment/shipping/2021-12-28/packingSlips"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetPackingSlips: %w", err)
	}
	return result, nil
}

// GetShippingLabel getShippingLabel
// Method: GET | Path: /vendor/directFulfillment/shipping/2021-12-28/shippingLabels/{purchaseOrderNumber}
func (c *Client) GetShippingLabel(ctx context.Context, purchaseOrderNumber string, query map[string]string) (interface{}, error) {
	path := "/vendor/directFulfillment/shipping/2021-12-28/shippingLabels/{purchaseOrderNumber}"
	path = strings.Replace(path, "{purchaseOrderNumber}", purchaseOrderNumber, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetShippingLabel: %w", err)
	}
	return result, nil
}

// GetCustomerInvoices getCustomerInvoices
// Method: GET | Path: /vendor/directFulfillment/shipping/2021-12-28/customerInvoices
func (c *Client) GetCustomerInvoices(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/vendor/directFulfillment/shipping/2021-12-28/customerInvoices"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetCustomerInvoices: %w", err)
	}
	return result, nil
}

// GetCustomerInvoice getCustomerInvoice
// Method: GET | Path: /vendor/directFulfillment/shipping/2021-12-28/customerInvoices/{purchaseOrderNumber}
func (c *Client) GetCustomerInvoice(ctx context.Context, purchaseOrderNumber string, query map[string]string) (interface{}, error) {
	path := "/vendor/directFulfillment/shipping/2021-12-28/customerInvoices/{purchaseOrderNumber}"
	path = strings.Replace(path, "{purchaseOrderNumber}", purchaseOrderNumber, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetCustomerInvoice: %w", err)
	}
	return result, nil
}

// CreateShippingLabels createShippingLabels
// Method: POST | Path: /vendor/directFulfillment/shipping/2021-12-28/shippingLabels/{purchaseOrderNumber}
func (c *Client) CreateShippingLabels(ctx context.Context, purchaseOrderNumber string, body interface{}) (interface{}, error) {
	path := "/vendor/directFulfillment/shipping/2021-12-28/shippingLabels/{purchaseOrderNumber}"
	path = strings.Replace(path, "{purchaseOrderNumber}", purchaseOrderNumber, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("CreateShippingLabels: %w", err)
	}
	return result, nil
}

// GetPackingSlip getPackingSlip
// Method: GET | Path: /vendor/directFulfillment/shipping/2021-12-28/packingSlips/{purchaseOrderNumber}
func (c *Client) GetPackingSlip(ctx context.Context, purchaseOrderNumber string, query map[string]string) (interface{}, error) {
	path := "/vendor/directFulfillment/shipping/2021-12-28/packingSlips/{purchaseOrderNumber}"
	path = strings.Replace(path, "{purchaseOrderNumber}", purchaseOrderNumber, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetPackingSlip: %w", err)
	}
	return result, nil
}

// SubmitShipmentStatusUpdates submitShipmentStatusUpdates
// Method: POST | Path: /vendor/directFulfillment/shipping/2021-12-28/shipmentStatusUpdates
func (c *Client) SubmitShipmentStatusUpdates(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/vendor/directFulfillment/shipping/2021-12-28/shipmentStatusUpdates"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("SubmitShipmentStatusUpdates: %w", err)
	}
	return result, nil
}

// CreateContainerLabel createContainerLabel
// Method: POST | Path: /vendor/directFulfillment/shipping/2021-12-28/containerLabel
func (c *Client) CreateContainerLabel(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/vendor/directFulfillment/shipping/2021-12-28/containerLabel"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("CreateContainerLabel: %w", err)
	}
	return result, nil
}
