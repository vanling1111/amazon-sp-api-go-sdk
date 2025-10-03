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

package product_pricing_v2022_05_01

import (
	"context"
	"fmt"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client product-pricing API v2022-05-01
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// GetCompetitiveSummary
// Method: POST | Path: /batches/products/pricing/2022-05-01/items/competitiveSummary
func (c *Client) GetCompetitiveSummary(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/batches/products/pricing/2022-05-01/items/competitiveSummary"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("GetCompetitiveSummary: %w", err)
	}
	return result, nil
}

// GetFeaturedOfferExpectedPriceBatch
// Method: POST | Path: /batches/products/pricing/2022-05-01/offer/featuredOfferExpectedPrice
func (c *Client) GetFeaturedOfferExpectedPriceBatch(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/batches/products/pricing/2022-05-01/offer/featuredOfferExpectedPrice"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("GetFeaturedOfferExpectedPriceBatch: %w", err)
	}
	return result, nil
}
