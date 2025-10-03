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
