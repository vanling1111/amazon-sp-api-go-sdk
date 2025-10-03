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

package feeds_v2021_06_30

import (
	"context"
	"fmt"
	"strings"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client feeds API v2021-06-30
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// CreateFeed
// Method: POST | Path: /feeds/2021-06-30/feeds
func (c *Client) CreateFeed(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/feeds/2021-06-30/feeds"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("CreateFeed: %w", err)
	}
	return result, nil
}

// CreateFeedDocument
// Method: POST | Path: /feeds/2021-06-30/documents
func (c *Client) CreateFeedDocument(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/feeds/2021-06-30/documents"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("CreateFeedDocument: %w", err)
	}
	return result, nil
}

// GetFeeds
// Method: GET | Path: /feeds/2021-06-30/feeds
func (c *Client) GetFeeds(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/feeds/2021-06-30/feeds"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetFeeds: %w", err)
	}
	return result, nil
}

// CancelFeed
// Method: DELETE | Path: /feeds/2021-06-30/feeds/{feedId}
func (c *Client) CancelFeed(ctx context.Context, feedId string) (interface{}, error) {
	path := "/feeds/2021-06-30/feeds/{feedId}"
	path = strings.Replace(path, "{feedId}", feedId, 1)
	var result interface{}
	err := c.baseClient.Delete(ctx, path, &result)
	if err != nil {
		return nil, fmt.Errorf("CancelFeed: %w", err)
	}
	return result, nil
}

// GetFeedDocument
// Method: GET | Path: /feeds/2021-06-30/documents/{feedDocumentId}
func (c *Client) GetFeedDocument(ctx context.Context, feedDocumentId string, query map[string]string) (interface{}, error) {
	path := "/feeds/2021-06-30/documents/{feedDocumentId}"
	path = strings.Replace(path, "{feedDocumentId}", feedDocumentId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetFeedDocument: %w", err)
	}
	return result, nil
}

// GetFeed
// Method: GET | Path: /feeds/2021-06-30/feeds/{feedId}
func (c *Client) GetFeed(ctx context.Context, feedId string, query map[string]string) (interface{}, error) {
	path := "/feeds/2021-06-30/feeds/{feedId}"
	path = strings.Replace(path, "{feedId}", feedId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetFeed: %w", err)
	}
	return result, nil
}
