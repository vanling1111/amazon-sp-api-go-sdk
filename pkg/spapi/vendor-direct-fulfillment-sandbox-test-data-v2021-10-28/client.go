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

package vendor_direct_fulfillment_sandbox_test_data_v2021_10_28

import (
	"context"
	"fmt"
	"strings"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client vendor-direct-fulfillment-sandbox-test-data API v2021-10-28
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// GenerateOrderScenarios
// Method: POST | Path: /vendor/directFulfillment/sandbox/2021-10-28/orders
func (c *Client) GenerateOrderScenarios(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/vendor/directFulfillment/sandbox/2021-10-28/orders"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("GenerateOrderScenarios: %w", err)
	}
	return result, nil
}

// GetOrderScenarios
// Method: GET | Path: /vendor/directFulfillment/sandbox/2021-10-28/transactions/{transactionId}
func (c *Client) GetOrderScenarios(ctx context.Context, transactionId string, query map[string]string) (interface{}, error) {
	path := "/vendor/directFulfillment/sandbox/2021-10-28/transactions/{transactionId}"
	path = strings.Replace(path, "{transactionId}", transactionId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetOrderScenarios: %w", err)
	}
	return result, nil
}
