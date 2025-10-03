// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package tokens_v2021_03_01

import (
	"context"
	"fmt"
	
	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client tokens API v2021-03-01
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// CreateRestrictedDataToken 
// Method: POST | Path: /tokens/2021-03-01/restrictedDataToken
func (c *Client) CreateRestrictedDataToken(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/tokens/2021-03-01/restrictedDataToken"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil { return nil, fmt.Errorf("CreateRestrictedDataToken: %w", err) }
	return result, nil
}
