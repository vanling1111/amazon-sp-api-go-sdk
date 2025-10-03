// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package finances_v0

import (
	"context"
	"fmt"
	"strings"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client finances API v0
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// ListFinancialEventGroups
// Method: GET | Path: /finances/v0/financialEventGroups
func (c *Client) ListFinancialEventGroups(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/finances/v0/financialEventGroups"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("ListFinancialEventGroups: %w", err)
	}
	return result, nil
}

// ListFinancialEventsByGroupId
// Method: GET | Path: /finances/v0/financialEventGroups/{eventGroupId}/financialEvents
func (c *Client) ListFinancialEventsByGroupId(ctx context.Context, eventGroupId string, query map[string]string) (interface{}, error) {
	path := "/finances/v0/financialEventGroups/{eventGroupId}/financialEvents"
	path = strings.Replace(path, "{eventGroupId}", eventGroupId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("ListFinancialEventsByGroupId: %w", err)
	}
	return result, nil
}

// ListFinancialEventsByOrderId
// Method: GET | Path: /finances/v0/orders/{orderId}/financialEvents
func (c *Client) ListFinancialEventsByOrderId(ctx context.Context, orderId string, query map[string]string) (interface{}, error) {
	path := "/finances/v0/orders/{orderId}/financialEvents"
	path = strings.Replace(path, "{orderId}", orderId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("ListFinancialEventsByOrderId: %w", err)
	}
	return result, nil
}

// ListFinancialEvents
// Method: GET | Path: /finances/v0/financialEvents
func (c *Client) ListFinancialEvents(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/finances/v0/financialEvents"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("ListFinancialEvents: %w", err)
	}
	return result, nil
}
