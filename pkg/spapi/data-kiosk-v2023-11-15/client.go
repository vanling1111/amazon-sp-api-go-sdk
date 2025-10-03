// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package data_kiosk_v2023_11_15

import (
	"context"
	"fmt"
	"strings"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client data-kiosk API v2023-11-15
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// GetQuery
// Method: GET | Path: /dataKiosk/2023-11-15/queries/{queryId}
func (c *Client) GetQuery(ctx context.Context, queryId string, query map[string]string) (interface{}, error) {
	path := "/dataKiosk/2023-11-15/queries/{queryId}"
	path = strings.Replace(path, "{queryId}", queryId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetQuery: %w", err)
	}
	return result, nil
}

// GetDocument
// Method: GET | Path: /dataKiosk/2023-11-15/documents/{documentId}
func (c *Client) GetDocument(ctx context.Context, documentId string, query map[string]string) (interface{}, error) {
	path := "/dataKiosk/2023-11-15/documents/{documentId}"
	path = strings.Replace(path, "{documentId}", documentId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetDocument: %w", err)
	}
	return result, nil
}

// CreateQuery
// Method: POST | Path: /dataKiosk/2023-11-15/queries
func (c *Client) CreateQuery(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/dataKiosk/2023-11-15/queries"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("CreateQuery: %w", err)
	}
	return result, nil
}

// GetQueries
// Method: GET | Path: /dataKiosk/2023-11-15/queries
func (c *Client) GetQueries(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/dataKiosk/2023-11-15/queries"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetQueries: %w", err)
	}
	return result, nil
}

// CancelQuery
// Method: DELETE | Path: /dataKiosk/2023-11-15/queries/{queryId}
func (c *Client) CancelQuery(ctx context.Context, queryId string) (interface{}, error) {
	path := "/dataKiosk/2023-11-15/queries/{queryId}"
	path = strings.Replace(path, "{queryId}", queryId, 1)
	var result interface{}
	err := c.baseClient.Delete(ctx, path, &result)
	if err != nil {
		return nil, fmt.Errorf("CancelQuery: %w", err)
	}
	return result, nil
}
