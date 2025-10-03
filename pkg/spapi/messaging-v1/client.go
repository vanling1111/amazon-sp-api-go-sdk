// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package messaging_v1

import (
	"context"
	"fmt"
	"strings"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client messaging API v1
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// SendInvoice
// Method: POST | Path: /messaging/v1/orders/{amazonOrderId}/messages/invoice
func (c *Client) SendInvoice(ctx context.Context, amazonOrderId string, body interface{}) (interface{}, error) {
	path := "/messaging/v1/orders/{amazonOrderId}/messages/invoice"
	path = strings.Replace(path, "{amazonOrderId}", amazonOrderId, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("SendInvoice: %w", err)
	}
	return result, nil
}

// GetAttributes
// Method: GET | Path: /messaging/v1/orders/{amazonOrderId}/attributes
func (c *Client) GetAttributes(ctx context.Context, amazonOrderId string, query map[string]string) (interface{}, error) {
	path := "/messaging/v1/orders/{amazonOrderId}/attributes"
	path = strings.Replace(path, "{amazonOrderId}", amazonOrderId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetAttributes: %w", err)
	}
	return result, nil
}

// CreateAmazonMotors
// Method: POST | Path: /messaging/v1/orders/{amazonOrderId}/messages/amazonMotors
func (c *Client) CreateAmazonMotors(ctx context.Context, amazonOrderId string, body interface{}) (interface{}, error) {
	path := "/messaging/v1/orders/{amazonOrderId}/messages/amazonMotors"
	path = strings.Replace(path, "{amazonOrderId}", amazonOrderId, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("CreateAmazonMotors: %w", err)
	}
	return result, nil
}

// CreateLegalDisclosure
// Method: POST | Path: /messaging/v1/orders/{amazonOrderId}/messages/legalDisclosure
func (c *Client) CreateLegalDisclosure(ctx context.Context, amazonOrderId string, body interface{}) (interface{}, error) {
	path := "/messaging/v1/orders/{amazonOrderId}/messages/legalDisclosure"
	path = strings.Replace(path, "{amazonOrderId}", amazonOrderId, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("CreateLegalDisclosure: %w", err)
	}
	return result, nil
}

// ConfirmCustomizationDetails
// Method: POST | Path: /messaging/v1/orders/{amazonOrderId}/messages/confirmCustomizationDetails
func (c *Client) ConfirmCustomizationDetails(ctx context.Context, amazonOrderId string, body interface{}) (interface{}, error) {
	path := "/messaging/v1/orders/{amazonOrderId}/messages/confirmCustomizationDetails"
	path = strings.Replace(path, "{amazonOrderId}", amazonOrderId, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("ConfirmCustomizationDetails: %w", err)
	}
	return result, nil
}

// CreateWarranty
// Method: POST | Path: /messaging/v1/orders/{amazonOrderId}/messages/warranty
func (c *Client) CreateWarranty(ctx context.Context, amazonOrderId string, body interface{}) (interface{}, error) {
	path := "/messaging/v1/orders/{amazonOrderId}/messages/warranty"
	path = strings.Replace(path, "{amazonOrderId}", amazonOrderId, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("CreateWarranty: %w", err)
	}
	return result, nil
}

// CreateUnexpectedProblem
// Method: POST | Path: /messaging/v1/orders/{amazonOrderId}/messages/unexpectedProblem
func (c *Client) CreateUnexpectedProblem(ctx context.Context, amazonOrderId string, body interface{}) (interface{}, error) {
	path := "/messaging/v1/orders/{amazonOrderId}/messages/unexpectedProblem"
	path = strings.Replace(path, "{amazonOrderId}", amazonOrderId, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("CreateUnexpectedProblem: %w", err)
	}
	return result, nil
}

// CreateDigitalAccessKey
// Method: POST | Path: /messaging/v1/orders/{amazonOrderId}/messages/digitalAccessKey
func (c *Client) CreateDigitalAccessKey(ctx context.Context, amazonOrderId string, body interface{}) (interface{}, error) {
	path := "/messaging/v1/orders/{amazonOrderId}/messages/digitalAccessKey"
	path = strings.Replace(path, "{amazonOrderId}", amazonOrderId, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("CreateDigitalAccessKey: %w", err)
	}
	return result, nil
}

// CreateConfirmServiceDetails
// Method: POST | Path: /messaging/v1/orders/{amazonOrderId}/messages/confirmServiceDetails
func (c *Client) CreateConfirmServiceDetails(ctx context.Context, amazonOrderId string, body interface{}) (interface{}, error) {
	path := "/messaging/v1/orders/{amazonOrderId}/messages/confirmServiceDetails"
	path = strings.Replace(path, "{amazonOrderId}", amazonOrderId, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("CreateConfirmServiceDetails: %w", err)
	}
	return result, nil
}

// CreateConfirmOrderDetails
// Method: POST | Path: /messaging/v1/orders/{amazonOrderId}/messages/confirmOrderDetails
func (c *Client) CreateConfirmOrderDetails(ctx context.Context, amazonOrderId string, body interface{}) (interface{}, error) {
	path := "/messaging/v1/orders/{amazonOrderId}/messages/confirmOrderDetails"
	path = strings.Replace(path, "{amazonOrderId}", amazonOrderId, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("CreateConfirmOrderDetails: %w", err)
	}
	return result, nil
}

// CreateConfirmDeliveryDetails
// Method: POST | Path: /messaging/v1/orders/{amazonOrderId}/messages/confirmDeliveryDetails
func (c *Client) CreateConfirmDeliveryDetails(ctx context.Context, amazonOrderId string, body interface{}) (interface{}, error) {
	path := "/messaging/v1/orders/{amazonOrderId}/messages/confirmDeliveryDetails"
	path = strings.Replace(path, "{amazonOrderId}", amazonOrderId, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("CreateConfirmDeliveryDetails: %w", err)
	}
	return result, nil
}

// GetMessagingActionsForOrder
// Method: GET | Path: /messaging/v1/orders/{amazonOrderId}
func (c *Client) GetMessagingActionsForOrder(ctx context.Context, amazonOrderId string, query map[string]string) (interface{}, error) {
	path := "/messaging/v1/orders/{amazonOrderId}"
	path = strings.Replace(path, "{amazonOrderId}", amazonOrderId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetMessagingActionsForOrder: %w", err)
	}
	return result, nil
}
