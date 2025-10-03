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

package replenishment_v2022_11_07

import (
	"context"
	"fmt"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client replenishment API v2022-11-07
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// GetSellingPartnerMetrics
// Method: POST | Path: /replenishment/2022-11-07/sellingPartners/metrics/search
func (c *Client) GetSellingPartnerMetrics(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/replenishment/2022-11-07/sellingPartners/metrics/search"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("GetSellingPartnerMetrics: %w", err)
	}
	return result, nil
}

// ListOfferMetrics
// Method: POST | Path: /replenishment/2022-11-07/offers/metrics/search
func (c *Client) ListOfferMetrics(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/replenishment/2022-11-07/offers/metrics/search"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("ListOfferMetrics: %w", err)
	}
	return result, nil
}

// ListOffers
// Method: POST | Path: /replenishment/2022-11-07/offers/search
func (c *Client) ListOffers(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/replenishment/2022-11-07/offers/search"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("ListOffers: %w", err)
	}
	return result, nil
}
