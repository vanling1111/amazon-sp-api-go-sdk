// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package shipment_invoicing_v0

import (
	"context"
	"fmt"
	"strings"
	
	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client shipment-invoicing API v0
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// GetShipmentDetails 
// Method: GET | Path: /fba/outbound/brazil/v0/shipments/{shipmentId}
func (c *Client) GetShipmentDetails(ctx context.Context, shipmentId string, query map[string]string) (interface{}, error) {
	path := "/fba/outbound/brazil/v0/shipments/{shipmentId}"
	path = strings.Replace(path, "{shipmentId}", shipmentId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil { return nil, fmt.Errorf("GetShipmentDetails: %w", err) }
	return result, nil
}

// GetInvoiceStatus 
// Method: GET | Path: /fba/outbound/brazil/v0/shipments/{shipmentId}/invoice/status
func (c *Client) GetInvoiceStatus(ctx context.Context, shipmentId string, query map[string]string) (interface{}, error) {
	path := "/fba/outbound/brazil/v0/shipments/{shipmentId}/invoice/status"
	path = strings.Replace(path, "{shipmentId}", shipmentId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil { return nil, fmt.Errorf("GetInvoiceStatus: %w", err) }
	return result, nil
}

// SubmitInvoice 
// Method: POST | Path: /fba/outbound/brazil/v0/shipments/{shipmentId}/invoice
func (c *Client) SubmitInvoice(ctx context.Context, shipmentId string, body interface{}) (interface{}, error) {
	path := "/fba/outbound/brazil/v0/shipments/{shipmentId}/invoice"
	path = strings.Replace(path, "{shipmentId}", shipmentId, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil { return nil, fmt.Errorf("SubmitInvoice: %w", err) }
	return result, nil
}
