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

package product_fees_v0

import (
	"context"
	"fmt"
	"strings"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client product-fees API v0
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// GetMyFeesEstimateForSKU
// Method: POST | Path: /products/fees/v0/listings/{SellerSKU}/feesEstimate
func (c *Client) GetMyFeesEstimateForSKU(ctx context.Context, sellerSKU string, body interface{}) (interface{}, error) {
	path := "/products/fees/v0/listings/{SellerSKU}/feesEstimate"
	path = strings.Replace(path, "{SellerSKU}", sellerSKU, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("GetMyFeesEstimateForSKU: %w", err)
	}
	return result, nil
}

// GetMyFeesEstimateForASIN
// Method: POST | Path: /products/fees/v0/items/{Asin}/feesEstimate
func (c *Client) GetMyFeesEstimateForASIN(ctx context.Context, asin string, body interface{}) (interface{}, error) {
	path := "/products/fees/v0/items/{Asin}/feesEstimate"
	path = strings.Replace(path, "{Asin}", asin, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("GetMyFeesEstimateForASIN: %w", err)
	}
	return result, nil
}

// GetMyFeesEstimates
// Method: POST | Path: /products/fees/v0/feesEstimate
func (c *Client) GetMyFeesEstimates(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/products/fees/v0/feesEstimate"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("GetMyFeesEstimates: %w", err)
	}
	return result, nil
}
