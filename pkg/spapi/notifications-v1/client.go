// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package notifications_v1

import (
	"context"
	"fmt"
	"strings"
	
	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client notifications API v1
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// CreateSubscription 
// Method: POST | Path: /notifications/v1/subscriptions/{notificationType}
func (c *Client) CreateSubscription(ctx context.Context, notificationType string, body interface{}) (interface{}, error) {
	path := "/notifications/v1/subscriptions/{notificationType}"
	path = strings.Replace(path, "{notificationType}", notificationType, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil { return nil, fmt.Errorf("CreateSubscription: %w", err) }
	return result, nil
}

// CreateDestination 
// Method: POST | Path: /notifications/v1/destinations
func (c *Client) CreateDestination(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/notifications/v1/destinations"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil { return nil, fmt.Errorf("CreateDestination: %w", err) }
	return result, nil
}

// GetDestinations 
// Method: GET | Path: /notifications/v1/destinations
func (c *Client) GetDestinations(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/notifications/v1/destinations"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil { return nil, fmt.Errorf("GetDestinations: %w", err) }
	return result, nil
}

// DeleteSubscriptionById 
// Method: DELETE | Path: /notifications/v1/subscriptions/{notificationType}/{subscriptionId}
func (c *Client) DeleteSubscriptionById(ctx context.Context, notificationType string, subscriptionId string) (interface{}, error) {
	path := "/notifications/v1/subscriptions/{notificationType}/{subscriptionId}"
	path = strings.Replace(path, "{notificationType}", notificationType, 1)
	path = strings.Replace(path, "{subscriptionId}", subscriptionId, 1)
	var result interface{}
	err := c.baseClient.Delete(ctx, path, &result)
	if err != nil { return nil, fmt.Errorf("DeleteSubscriptionById: %w", err) }
	return result, nil
}

// GetSubscription 
// Method: GET | Path: /notifications/v1/subscriptions/{notificationType}
func (c *Client) GetSubscription(ctx context.Context, notificationType string, query map[string]string) (interface{}, error) {
	path := "/notifications/v1/subscriptions/{notificationType}"
	path = strings.Replace(path, "{notificationType}", notificationType, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil { return nil, fmt.Errorf("GetSubscription: %w", err) }
	return result, nil
}

// GetDestination 
// Method: GET | Path: /notifications/v1/destinations/{destinationId}
func (c *Client) GetDestination(ctx context.Context, destinationId string, query map[string]string) (interface{}, error) {
	path := "/notifications/v1/destinations/{destinationId}"
	path = strings.Replace(path, "{destinationId}", destinationId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil { return nil, fmt.Errorf("GetDestination: %w", err) }
	return result, nil
}

// DeleteDestination 
// Method: DELETE | Path: /notifications/v1/destinations/{destinationId}
func (c *Client) DeleteDestination(ctx context.Context, destinationId string) (interface{}, error) {
	path := "/notifications/v1/destinations/{destinationId}"
	path = strings.Replace(path, "{destinationId}", destinationId, 1)
	var result interface{}
	err := c.baseClient.Delete(ctx, path, &result)
	if err != nil { return nil, fmt.Errorf("DeleteDestination: %w", err) }
	return result, nil
}

// GetSubscriptionById 
// Method: GET | Path: /notifications/v1/subscriptions/{notificationType}/{subscriptionId}
func (c *Client) GetSubscriptionById(ctx context.Context, notificationType string, subscriptionId string, query map[string]string) (interface{}, error) {
	path := "/notifications/v1/subscriptions/{notificationType}/{subscriptionId}"
	path = strings.Replace(path, "{notificationType}", notificationType, 1)
	path = strings.Replace(path, "{subscriptionId}", subscriptionId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil { return nil, fmt.Errorf("GetSubscriptionById: %w", err) }
	return result, nil
}
