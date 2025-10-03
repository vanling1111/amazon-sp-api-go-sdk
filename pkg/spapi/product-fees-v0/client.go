// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package product_fees_v0

import (
	"context"
	"fmt"
	"strings"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client product-fees API v0
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// GetMyFeesEstimateForSKU
// Method: POST | Path: /products/fees/v0/listings/{SellerSKU}/feesEstimate
func (c *Client) GetMyFeesEstimateForSKU(ctx context.Context, sellerSKU string, body interface{}) (interface{}, error) {
	path := "/products/fees/v0/listings/{SellerSKU}/feesEstimate"
	path = strings.Replace(path, "{SellerSKU}", sellerSKU, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("GetMyFeesEstimateForSKU: %w", err)
	}
	return result, nil
}

// GetMyFeesEstimateForASIN
// Method: POST | Path: /products/fees/v0/items/{Asin}/feesEstimate
func (c *Client) GetMyFeesEstimateForASIN(ctx context.Context, asin string, body interface{}) (interface{}, error) {
	path := "/products/fees/v0/items/{Asin}/feesEstimate"
	path = strings.Replace(path, "{Asin}", asin, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("GetMyFeesEstimateForASIN: %w", err)
	}
	return result, nil
}

// GetMyFeesEstimates
// Method: POST | Path: /products/fees/v0/feesEstimate
func (c *Client) GetMyFeesEstimates(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/products/fees/v0/feesEstimate"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("GetMyFeesEstimates: %w", err)
	}
	return result, nil
}
