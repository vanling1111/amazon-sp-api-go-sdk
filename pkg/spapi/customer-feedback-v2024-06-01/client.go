// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package customer_feedback_v2024_06_01

import (
	"context"
	"fmt"
	"strings"
	
	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client customer-feedback API v2024-06-01
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// GetBrowseNodeReturnTopics 
// Method: GET | Path: /customerFeedback/2024-06-01/browseNodes/{browseNodeId}/returns/topics
func (c *Client) GetBrowseNodeReturnTopics(ctx context.Context, browseNodeId string, query map[string]string) (interface{}, error) {
	path := "/customerFeedback/2024-06-01/browseNodes/{browseNodeId}/returns/topics"
	path = strings.Replace(path, "{browseNodeId}", browseNodeId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil { return nil, fmt.Errorf("GetBrowseNodeReturnTopics: %w", err) }
	return result, nil
}

// GetBrowseNodeReviewTopics 
// Method: GET | Path: /customerFeedback/2024-06-01/browseNodes/{browseNodeId}/reviews/topics
func (c *Client) GetBrowseNodeReviewTopics(ctx context.Context, browseNodeId string, query map[string]string) (interface{}, error) {
	path := "/customerFeedback/2024-06-01/browseNodes/{browseNodeId}/reviews/topics"
	path = strings.Replace(path, "{browseNodeId}", browseNodeId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil { return nil, fmt.Errorf("GetBrowseNodeReviewTopics: %w", err) }
	return result, nil
}

// GetItemBrowseNode 
// Method: GET | Path: /customerFeedback/2024-06-01/items/{asin}/browseNode
func (c *Client) GetItemBrowseNode(ctx context.Context, asin string, query map[string]string) (interface{}, error) {
	path := "/customerFeedback/2024-06-01/items/{asin}/browseNode"
	path = strings.Replace(path, "{asin}", asin, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil { return nil, fmt.Errorf("GetItemBrowseNode: %w", err) }
	return result, nil
}

// GetBrowseNodeReviewTrends 
// Method: GET | Path: /customerFeedback/2024-06-01/browseNodes/{browseNodeId}/reviews/trends
func (c *Client) GetBrowseNodeReviewTrends(ctx context.Context, browseNodeId string, query map[string]string) (interface{}, error) {
	path := "/customerFeedback/2024-06-01/browseNodes/{browseNodeId}/reviews/trends"
	path = strings.Replace(path, "{browseNodeId}", browseNodeId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil { return nil, fmt.Errorf("GetBrowseNodeReviewTrends: %w", err) }
	return result, nil
}

// GetBrowseNodeReturnTrends 
// Method: GET | Path: /customerFeedback/2024-06-01/browseNodes/{browseNodeId}/returns/trends
func (c *Client) GetBrowseNodeReturnTrends(ctx context.Context, browseNodeId string, query map[string]string) (interface{}, error) {
	path := "/customerFeedback/2024-06-01/browseNodes/{browseNodeId}/returns/trends"
	path = strings.Replace(path, "{browseNodeId}", browseNodeId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil { return nil, fmt.Errorf("GetBrowseNodeReturnTrends: %w", err) }
	return result, nil
}

// GetItemReviewTrends 
// Method: GET | Path: /customerFeedback/2024-06-01/items/{asin}/reviews/trends
func (c *Client) GetItemReviewTrends(ctx context.Context, asin string, query map[string]string) (interface{}, error) {
	path := "/customerFeedback/2024-06-01/items/{asin}/reviews/trends"
	path = strings.Replace(path, "{asin}", asin, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil { return nil, fmt.Errorf("GetItemReviewTrends: %w", err) }
	return result, nil
}

// GetItemReviewTopics 
// Method: GET | Path: /customerFeedback/2024-06-01/items/{asin}/reviews/topics
func (c *Client) GetItemReviewTopics(ctx context.Context, asin string, query map[string]string) (interface{}, error) {
	path := "/customerFeedback/2024-06-01/items/{asin}/reviews/topics"
	path = strings.Replace(path, "{asin}", asin, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil { return nil, fmt.Errorf("GetItemReviewTopics: %w", err) }
	return result, nil
}
