// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package sellers_v1

import (
	"context"
	"fmt"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client sellers API v1
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// GetAccount
// Method: GET | Path: /sellers/v1/account
func (c *Client) GetAccount(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/sellers/v1/account"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetAccount: %w", err)
	}
	return result, nil
}

// GetMarketplaceParticipations
// Method: GET | Path: /sellers/v1/marketplaceParticipations
func (c *Client) GetMarketplaceParticipations(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/sellers/v1/marketplaceParticipations"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetMarketplaceParticipations: %w", err)
	}
	return result, nil
}
