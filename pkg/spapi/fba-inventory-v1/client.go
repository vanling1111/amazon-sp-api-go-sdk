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

package fba_inventory_v1

import (
	"context"
	"fmt"
	"strings"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client fba-inventory API v1
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// CreateInventoryItem
// Method: POST | Path: /fba/inventory/v1/items
func (c *Client) CreateInventoryItem(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/fba/inventory/v1/items"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("CreateInventoryItem: %w", err)
	}
	return result, nil
}

// AddInventory
// Method: POST | Path: /fba/inventory/v1/items/inventory
func (c *Client) AddInventory(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/fba/inventory/v1/items/inventory"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("AddInventory: %w", err)
	}
	return result, nil
}

// GetInventorySummaries
// Method: GET | Path: /fba/inventory/v1/summaries
func (c *Client) GetInventorySummaries(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/fba/inventory/v1/summaries"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetInventorySummaries: %w", err)
	}
	return result, nil
}

// DeleteInventoryItem
// Method: DELETE | Path: /fba/inventory/v1/items/{sellerSku}
func (c *Client) DeleteInventoryItem(ctx context.Context, sellerSku string) (interface{}, error) {
	path := "/fba/inventory/v1/items/{sellerSku}"
	path = strings.Replace(path, "{sellerSku}", sellerSku, 1)
	var result interface{}
	err := c.baseClient.Delete(ctx, path, &result)
	if err != nil {
		return nil, fmt.Errorf("DeleteInventoryItem: %w", err)
	}
	return result, nil
}
