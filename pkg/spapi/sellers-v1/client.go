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

package sellers_v1

import (
	"context"
	"fmt"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client sellers API v1
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// GetAccount
// Method: GET | Path: /sellers/v1/account
func (c *Client) GetAccount(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/sellers/v1/account"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetAccount: %w", err)
	}
	return result, nil
}

// GetMarketplaceParticipations
// Method: GET | Path: /sellers/v1/marketplaceParticipations
func (c *Client) GetMarketplaceParticipations(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/sellers/v1/marketplaceParticipations"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetMarketplaceParticipations: %w", err)
	}
	return result, nil
}
