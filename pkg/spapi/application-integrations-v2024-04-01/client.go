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
	if err != nil {
		return nil, fmt.Errorf("RecordActionFeedback: %w", err)
	}
	return result, nil
}

// CreateNotification
// Method: POST | Path: /appIntegrations/2024-04-01/notifications
func (c *Client) CreateNotification(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/appIntegrations/2024-04-01/notifications"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("CreateNotification: %w", err)
	}
	return result, nil
}

// DeleteNotifications
// Method: POST | Path: /appIntegrations/2024-04-01/notifications/deletion
func (c *Client) DeleteNotifications(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/appIntegrations/2024-04-01/notifications/deletion"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("DeleteNotifications: %w", err)
	}
	return result, nil
}
