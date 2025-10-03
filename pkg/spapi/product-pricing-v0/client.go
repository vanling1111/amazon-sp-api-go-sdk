// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package product_pricing_v0

import (
	"context"
	"fmt"
	"strings"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client product-pricing API v0
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// GetPricing
// Method: GET | Path: /products/pricing/v0/price
func (c *Client) GetPricing(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/products/pricing/v0/price"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetPricing: %w", err)
	}
	return result, nil
}

// GetListingOffersBatch
// Method: POST | Path: /batches/products/pricing/v0/listingOffers
func (c *Client) GetListingOffersBatch(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/batches/products/pricing/v0/listingOffers"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("GetListingOffersBatch: %w", err)
	}
	return result, nil
}

// GetItemOffersBatch
// Method: POST | Path: /batches/products/pricing/v0/itemOffers
func (c *Client) GetItemOffersBatch(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/batches/products/pricing/v0/itemOffers"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("GetItemOffersBatch: %w", err)
	}
	return result, nil
}

// GetItemOffers
// Method: GET | Path: /products/pricing/v0/items/{Asin}/offers
func (c *Client) GetItemOffers(ctx context.Context, asin string, query map[string]string) (interface{}, error) {
	path := "/products/pricing/v0/items/{Asin}/offers"
	path = strings.Replace(path, "{Asin}", asin, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetItemOffers: %w", err)
	}
	return result, nil
}

// GetListingOffers
// Method: GET | Path: /products/pricing/v0/listings/{SellerSKU}/offers
func (c *Client) GetListingOffers(ctx context.Context, sellerSKU string, query map[string]string) (interface{}, error) {
	path := "/products/pricing/v0/listings/{SellerSKU}/offers"
	path = strings.Replace(path, "{SellerSKU}", sellerSKU, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetListingOffers: %w", err)
	}
	return result, nil
}

// GetCompetitivePricing
// Method: GET | Path: /products/pricing/v0/competitivePrice
func (c *Client) GetCompetitivePricing(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/products/pricing/v0/competitivePrice"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetCompetitivePricing: %w", err)
	}
	return result, nil
}
