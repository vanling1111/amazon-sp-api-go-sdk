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

package listings_items_v2020_09_01

import (
	"context"
	"fmt"
	"strings"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client listings-items API v2020-09-01
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// PutListingsItem
// Method: PUT | Path: /listings/2020-09-01/items/{sellerId}/{sku}
func (c *Client) PutListingsItem(ctx context.Context, sellerId string, sku string, body interface{}) (interface{}, error) {
	path := "/listings/2020-09-01/items/{sellerId}/{sku}"
	path = strings.Replace(path, "{sellerId}", sellerId, 1)
	path = strings.Replace(path, "{sku}", sku, 1)
	var result interface{}
	err := c.baseClient.Put(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("PutListingsItem: %w", err)
	}
	return result, nil
}

// PatchListingsItem
// Method: PATCH | Path: /listings/2020-09-01/items/{sellerId}/{sku}
func (c *Client) PatchListingsItem(ctx context.Context, sellerId string, sku string, body interface{}) (interface{}, error) {
	path := "/listings/2020-09-01/items/{sellerId}/{sku}"
	path = strings.Replace(path, "{sellerId}", sellerId, 1)
	path = strings.Replace(path, "{sku}", sku, 1)
	var result interface{}
	err := c.baseClient.DoRequest(ctx, "PATCH", path, nil, body, &result)
	if err != nil {
		return nil, fmt.Errorf("PatchListingsItem: %w", err)
	}
	return result, nil
}

// DeleteListingsItem
// Method: DELETE | Path: /listings/2020-09-01/items/{sellerId}/{sku}
func (c *Client) DeleteListingsItem(ctx context.Context, sellerId string, sku string) (interface{}, error) {
	path := "/listings/2020-09-01/items/{sellerId}/{sku}"
	path = strings.Replace(path, "{sellerId}", sellerId, 1)
	path = strings.Replace(path, "{sku}", sku, 1)
	var result interface{}
	err := c.baseClient.Delete(ctx, path, &result)
	if err != nil {
		return nil, fmt.Errorf("DeleteListingsItem: %w", err)
	}
	return result, nil
}
