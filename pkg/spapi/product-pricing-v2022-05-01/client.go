// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package product_pricing_v2022_05_01

import (
	"context"
	"fmt"
	
	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client product-pricing API v2022-05-01
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// GetCompetitiveSummary 
// Method: POST | Path: /batches/products/pricing/2022-05-01/items/competitiveSummary
func (c *Client) GetCompetitiveSummary(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/batches/products/pricing/2022-05-01/items/competitiveSummary"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil { return nil, fmt.Errorf("GetCompetitiveSummary: %w", err) }
	return result, nil
}

// GetFeaturedOfferExpectedPriceBatch 
// Method: POST | Path: /batches/products/pricing/2022-05-01/offer/featuredOfferExpectedPrice
func (c *Client) GetFeaturedOfferExpectedPriceBatch(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/batches/products/pricing/2022-05-01/offer/featuredOfferExpectedPrice"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil { return nil, fmt.Errorf("GetFeaturedOfferExpectedPriceBatch: %w", err) }
	return result, nil
}
