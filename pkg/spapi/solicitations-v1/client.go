// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package solicitations_v1

import (
	"context"
	"fmt"
	"strings"
	
	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client solicitations API v1
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// CreateProductReviewAndSellerFeedbackSolicitation 
// Method: POST | Path: /solicitations/v1/orders/{amazonOrderId}/solicitations/productReviewAndSellerFeedback
func (c *Client) CreateProductReviewAndSellerFeedbackSolicitation(ctx context.Context, amazonOrderId string, body interface{}) (interface{}, error) {
	path := "/solicitations/v1/orders/{amazonOrderId}/solicitations/productReviewAndSellerFeedback"
	path = strings.Replace(path, "{amazonOrderId}", amazonOrderId, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil { return nil, fmt.Errorf("CreateProductReviewAndSellerFeedbackSolicitation: %w", err) }
	return result, nil
}

// GetSolicitationActionsForOrder 
// Method: GET | Path: /solicitations/v1/orders/{amazonOrderId}
func (c *Client) GetSolicitationActionsForOrder(ctx context.Context, amazonOrderId string, query map[string]string) (interface{}, error) {
	path := "/solicitations/v1/orders/{amazonOrderId}"
	path = strings.Replace(path, "{amazonOrderId}", amazonOrderId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil { return nil, fmt.Errorf("GetSolicitationActionsForOrder: %w", err) }
	return result, nil
}
