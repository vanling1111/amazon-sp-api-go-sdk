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

package vendor_direct_fulfillment_orders_v1

import (
	"context"
	"fmt"
	"strings"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client vendor-direct-fulfillment-orders API v1
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// GetOrder
// Method: GET | Path: /vendor/directFulfillment/orders/v1/purchaseOrders/{purchaseOrderNumber}
func (c *Client) GetOrder(ctx context.Context, purchaseOrderNumber string, query map[string]string) (interface{}, error) {
	path := "/vendor/directFulfillment/orders/v1/purchaseOrders/{purchaseOrderNumber}"
	path = strings.Replace(path, "{purchaseOrderNumber}", purchaseOrderNumber, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetOrder: %w", err)
	}
	return result, nil
}

// GetOrders
// Method: GET | Path: /vendor/directFulfillment/orders/v1/purchaseOrders
func (c *Client) GetOrders(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/vendor/directFulfillment/orders/v1/purchaseOrders"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetOrders: %w", err)
	}
	return result, nil
}

// SubmitAcknowledgement
// Method: POST | Path: /vendor/directFulfillment/orders/v1/acknowledgements
func (c *Client) SubmitAcknowledgement(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/vendor/directFulfillment/orders/v1/acknowledgements"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("SubmitAcknowledgement: %w", err)
	}
	return result, nil
}
