// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package application_integrations_v2024_04_01

import (
	"context"
	"fmt"
	"strings"
	
	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client application-integrations API v2024-04-01
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// RecordActionFeedback 
// Method: POST | Path: /appIntegrations/2024-04-01/notifications/{notificationId}/feedback
func (c *Client) RecordActionFeedback(ctx context.Context, notificationId string, body interface{}) (interface{}, error) {
	path := "/appIntegrations/2024-04-01/notifications/{notificationId}/feedback"
	path = strings.Replace(path, "{notificationId}", notificationId, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil { return nil, fmt.Errorf("RecordActionFeedback: %w", err) }
	return result, nil
}

// CreateNotification 
// Method: POST | Path: /appIntegrations/2024-04-01/notifications
func (c *Client) CreateNotification(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/appIntegrations/2024-04-01/notifications"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil { return nil, fmt.Errorf("CreateNotification: %w", err) }
	return result, nil
}

// DeleteNotifications 
// Method: POST | Path: /appIntegrations/2024-04-01/notifications/deletion
func (c *Client) DeleteNotifications(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/appIntegrations/2024-04-01/notifications/deletion"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil { return nil, fmt.Errorf("DeleteNotifications: %w", err) }
	return result, nil
}
