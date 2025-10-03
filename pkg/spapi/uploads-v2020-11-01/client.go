// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package uploads_v2020_11_01

import (
	"context"
	"fmt"
	"strings"
	
	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client uploads API v2020-11-01
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// CreateUploadDestinationForResource 
// Method: POST | Path: /uploads/2020-11-01/uploadDestinations/{resource}
func (c *Client) CreateUploadDestinationForResource(ctx context.Context, resource string, body interface{}) (interface{}, error) {
	path := "/uploads/2020-11-01/uploadDestinations/{resource}"
	path = strings.Replace(path, "{resource}", resource, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil { return nil, fmt.Errorf("CreateUploadDestinationForResource: %w", err) }
	return result, nil
}
