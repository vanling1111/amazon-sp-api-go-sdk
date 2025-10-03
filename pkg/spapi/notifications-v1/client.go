// Copyright 2025 Amazon SP-API Go SDK Authors.
//
// This file is part of Amazon SP-API Go SDK.
//
// Amazon SP-API Go SDK is dual-licensed:
//
// 1. GNU Affero General Public License v3.0 (AGPL-3.0) for open source use
//    - Free for personal, educational, and open source projects
//    - Your project must also be open sourced under AGPL-3.0
//    - See: https://www.gnu.org/licenses/agpl-3.0.html
//
// 2. Commercial License for proprietary/commercial use
//    - Required for any commercial, enterprise, or proprietary use
//    - Allows closed source distribution
//    - Contact: vanling1111@gmail.com
//
// Unless you have obtained a commercial license, this file is licensed
// under AGPL-3.0. By using this software, you agree to comply with the
// terms of the applicable license.
//
// This file is part of Amazon SP-API Go SDK.
//
// Amazon SP-API Go SDK is dual-licensed:
//
// 1. GNU Affero General Public License v3.0 (AGPL-3.0) for open source use
//    - Free for personal, educational, and open source projects
//    - Your project must also be open sourced under AGPL-3.0
//    - See: https://www.gnu.org/licenses/agpl-3.0.html
//
// 2. Commercial License for proprietary/commercial use
//    - Required for any commercial, enterprise, or proprietary use
//    - Allows closed source distribution
//    - Contact: vanling1111@gmail.com
//
// Unless you have obtained a commercial license, this file is licensed
// under AGPL-3.0. By using this software, you agree to comply with the
// terms of the applicable license.

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
	if err != nil {
		return nil, fmt.Errorf("CreateSubscription: %w", err)
	}
	return result, nil
}

// CreateDestination
// Method: POST | Path: /notifications/v1/destinations
func (c *Client) CreateDestination(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/notifications/v1/destinations"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("CreateDestination: %w", err)
	}
	return result, nil
}

// GetDestinations
// Method: GET | Path: /notifications/v1/destinations
func (c *Client) GetDestinations(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/notifications/v1/destinations"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetDestinations: %w", err)
	}
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
	if err != nil {
		return nil, fmt.Errorf("DeleteSubscriptionById: %w", err)
	}
	return result, nil
}

// GetSubscription
// Method: GET | Path: /notifications/v1/subscriptions/{notificationType}
func (c *Client) GetSubscription(ctx context.Context, notificationType string, query map[string]string) (interface{}, error) {
	path := "/notifications/v1/subscriptions/{notificationType}"
	path = strings.Replace(path, "{notificationType}", notificationType, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetSubscription: %w", err)
	}
	return result, nil
}

// GetDestination
// Method: GET | Path: /notifications/v1/destinations/{destinationId}
func (c *Client) GetDestination(ctx context.Context, destinationId string, query map[string]string) (interface{}, error) {
	path := "/notifications/v1/destinations/{destinationId}"
	path = strings.Replace(path, "{destinationId}", destinationId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetDestination: %w", err)
	}
	return result, nil
}

// DeleteDestination
// Method: DELETE | Path: /notifications/v1/destinations/{destinationId}
func (c *Client) DeleteDestination(ctx context.Context, destinationId string) (interface{}, error) {
	path := "/notifications/v1/destinations/{destinationId}"
	path = strings.Replace(path, "{destinationId}", destinationId, 1)
	var result interface{}
	err := c.baseClient.Delete(ctx, path, &result)
	if err != nil {
		return nil, fmt.Errorf("DeleteDestination: %w", err)
	}
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
	if err != nil {
		return nil, fmt.Errorf("GetSubscriptionById: %w", err)
	}
	return result, nil
}
