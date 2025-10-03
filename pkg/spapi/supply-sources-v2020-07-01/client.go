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

package supply_sources_v2020_07_01

import (
	"context"
	"fmt"
	"strings"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client supply-sources API v2020-07-01
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// GetSupplySources
// Method: GET | Path: /supplySources/2020-07-01/supplySources
func (c *Client) GetSupplySources(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/supplySources/2020-07-01/supplySources"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetSupplySources: %w", err)
	}
	return result, nil
}

// UpdateSupplySourceStatus
// Method: PUT | Path: /supplySources/2020-07-01/supplySources/{supplySourceId}/status
func (c *Client) UpdateSupplySourceStatus(ctx context.Context, supplySourceId string, body interface{}) (interface{}, error) {
	path := "/supplySources/2020-07-01/supplySources/{supplySourceId}/status"
	path = strings.Replace(path, "{supplySourceId}", supplySourceId, 1)
	var result interface{}
	err := c.baseClient.Put(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("UpdateSupplySourceStatus: %w", err)
	}
	return result, nil
}

// UpdateSupplySource
// Method: PUT | Path: /supplySources/2020-07-01/supplySources/{supplySourceId}
func (c *Client) UpdateSupplySource(ctx context.Context, supplySourceId string, body interface{}) (interface{}, error) {
	path := "/supplySources/2020-07-01/supplySources/{supplySourceId}"
	path = strings.Replace(path, "{supplySourceId}", supplySourceId, 1)
	var result interface{}
	err := c.baseClient.Put(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("UpdateSupplySource: %w", err)
	}
	return result, nil
}

// GetSupplySource
// Method: GET | Path: /supplySources/2020-07-01/supplySources/{supplySourceId}
func (c *Client) GetSupplySource(ctx context.Context, supplySourceId string, query map[string]string) (interface{}, error) {
	path := "/supplySources/2020-07-01/supplySources/{supplySourceId}"
	path = strings.Replace(path, "{supplySourceId}", supplySourceId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetSupplySource: %w", err)
	}
	return result, nil
}

// CreateSupplySource
// Method: POST | Path: /supplySources/2020-07-01/supplySources
func (c *Client) CreateSupplySource(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/supplySources/2020-07-01/supplySources"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("CreateSupplySource: %w", err)
	}
	return result, nil
}

// ArchiveSupplySource
// Method: DELETE | Path: /supplySources/2020-07-01/supplySources/{supplySourceId}
func (c *Client) ArchiveSupplySource(ctx context.Context, supplySourceId string) (interface{}, error) {
	path := "/supplySources/2020-07-01/supplySources/{supplySourceId}"
	path = strings.Replace(path, "{supplySourceId}", supplySourceId, 1)
	var result interface{}
	err := c.baseClient.Delete(ctx, path, &result)
	if err != nil {
		return nil, fmt.Errorf("ArchiveSupplySource: %w", err)
	}
	return result, nil
}
