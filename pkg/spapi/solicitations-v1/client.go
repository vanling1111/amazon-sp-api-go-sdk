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

package solicitations_v1

import (
	"context"
	"fmt"
	"strings"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client solicitations API v1
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// CreateProductReviewAndSellerFeedbackSolicitation
// Method: POST | Path: /solicitations/v1/orders/{amazonOrderId}/solicitations/productReviewAndSellerFeedback
func (c *Client) CreateProductReviewAndSellerFeedbackSolicitation(ctx context.Context, amazonOrderId string, body interface{}) (interface{}, error) {
	path := "/solicitations/v1/orders/{amazonOrderId}/solicitations/productReviewAndSellerFeedback"
	path = strings.Replace(path, "{amazonOrderId}", amazonOrderId, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("CreateProductReviewAndSellerFeedbackSolicitation: %w", err)
	}
	return result, nil
}

// GetSolicitationActionsForOrder
// Method: GET | Path: /solicitations/v1/orders/{amazonOrderId}
func (c *Client) GetSolicitationActionsForOrder(ctx context.Context, amazonOrderId string, query map[string]string) (interface{}, error) {
	path := "/solicitations/v1/orders/{amazonOrderId}"
	path = strings.Replace(path, "{amazonOrderId}", amazonOrderId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetSolicitationActionsForOrder: %w", err)
	}
	return result, nil
}
