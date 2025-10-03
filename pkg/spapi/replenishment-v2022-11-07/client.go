// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package replenishment_v2022_11_07

import (
	"context"
	"fmt"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client replenishment API v2022-11-07
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// GetSellingPartnerMetrics
// Method: POST | Path: /replenishment/2022-11-07/sellingPartners/metrics/search
func (c *Client) GetSellingPartnerMetrics(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/replenishment/2022-11-07/sellingPartners/metrics/search"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("GetSellingPartnerMetrics: %w", err)
	}
	return result, nil
}

// ListOfferMetrics
// Method: POST | Path: /replenishment/2022-11-07/offers/metrics/search
func (c *Client) ListOfferMetrics(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/replenishment/2022-11-07/offers/metrics/search"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("ListOfferMetrics: %w", err)
	}
	return result, nil
}

// ListOffers
// Method: POST | Path: /replenishment/2022-11-07/offers/search
func (c *Client) ListOffers(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/replenishment/2022-11-07/offers/search"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("ListOffers: %w", err)
	}
	return result, nil
}
