// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package fba_inbound_eligibility_v1

import (
	"context"
	"fmt"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client fba-inbound-eligibility API v1
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// GetItemEligibilityPreview
// Method: GET | Path: /fba/inbound/v1/eligibility/itemPreview
func (c *Client) GetItemEligibilityPreview(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/fba/inbound/v1/eligibility/itemPreview"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetItemEligibilityPreview: %w", err)
	}
	return result, nil
}
