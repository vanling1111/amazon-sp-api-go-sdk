// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package aplus_content_v2020_11_01

import (
	"context"
	"fmt"
	"strings"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client aplus-content API v2020-11-01
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// PostContentDocumentAsinRelations
// Method: POST | Path: /aplus/2020-11-01/contentDocuments/{contentReferenceKey}/asins
func (c *Client) PostContentDocumentAsinRelations(ctx context.Context, contentReferenceKey string, body interface{}) (interface{}, error) {
	path := "/aplus/2020-11-01/contentDocuments/{contentReferenceKey}/asins"
	path = strings.Replace(path, "{contentReferenceKey}", contentReferenceKey, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("PostContentDocumentAsinRelations: %w", err)
	}
	return result, nil
}

// SearchContentPublishRecords
// Method: GET | Path: /aplus/2020-11-01/contentPublishRecords
func (c *Client) SearchContentPublishRecords(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/aplus/2020-11-01/contentPublishRecords"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("SearchContentPublishRecords: %w", err)
	}
	return result, nil
}

// GetContentDocument
// Method: GET | Path: /aplus/2020-11-01/contentDocuments/{contentReferenceKey}
func (c *Client) GetContentDocument(ctx context.Context, contentReferenceKey string, query map[string]string) (interface{}, error) {
	path := "/aplus/2020-11-01/contentDocuments/{contentReferenceKey}"
	path = strings.Replace(path, "{contentReferenceKey}", contentReferenceKey, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetContentDocument: %w", err)
	}
	return result, nil
}

// CreateContentDocument
// Method: POST | Path: /aplus/2020-11-01/contentDocuments
func (c *Client) CreateContentDocument(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/aplus/2020-11-01/contentDocuments"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("CreateContentDocument: %w", err)
	}
	return result, nil
}

// ValidateContentDocumentAsinRelations
// Method: POST | Path: /aplus/2020-11-01/contentAsinValidations
func (c *Client) ValidateContentDocumentAsinRelations(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/aplus/2020-11-01/contentAsinValidations"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("ValidateContentDocumentAsinRelations: %w", err)
	}
	return result, nil
}

// PostContentDocumentSuspendSubmission
// Method: POST | Path: /aplus/2020-11-01/contentDocuments/{contentReferenceKey}/suspendSubmissions
func (c *Client) PostContentDocumentSuspendSubmission(ctx context.Context, contentReferenceKey string, body interface{}) (interface{}, error) {
	path := "/aplus/2020-11-01/contentDocuments/{contentReferenceKey}/suspendSubmissions"
	path = strings.Replace(path, "{contentReferenceKey}", contentReferenceKey, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("PostContentDocumentSuspendSubmission: %w", err)
	}
	return result, nil
}

// ListContentDocumentAsinRelations
// Method: GET | Path: /aplus/2020-11-01/contentDocuments/{contentReferenceKey}/asins
func (c *Client) ListContentDocumentAsinRelations(ctx context.Context, contentReferenceKey string, query map[string]string) (interface{}, error) {
	path := "/aplus/2020-11-01/contentDocuments/{contentReferenceKey}/asins"
	path = strings.Replace(path, "{contentReferenceKey}", contentReferenceKey, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("ListContentDocumentAsinRelations: %w", err)
	}
	return result, nil
}

// SearchContentDocuments
// Method: GET | Path: /aplus/2020-11-01/contentDocuments
func (c *Client) SearchContentDocuments(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/aplus/2020-11-01/contentDocuments"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("SearchContentDocuments: %w", err)
	}
	return result, nil
}

// PostContentDocumentApprovalSubmission
// Method: POST | Path: /aplus/2020-11-01/contentDocuments/{contentReferenceKey}/approvalSubmissions
func (c *Client) PostContentDocumentApprovalSubmission(ctx context.Context, contentReferenceKey string, body interface{}) (interface{}, error) {
	path := "/aplus/2020-11-01/contentDocuments/{contentReferenceKey}/approvalSubmissions"
	path = strings.Replace(path, "{contentReferenceKey}", contentReferenceKey, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("PostContentDocumentApprovalSubmission: %w", err)
	}
	return result, nil
}

// UpdateContentDocument
// Method: POST | Path: /aplus/2020-11-01/contentDocuments/{contentReferenceKey}
func (c *Client) UpdateContentDocument(ctx context.Context, contentReferenceKey string, body interface{}) (interface{}, error) {
	path := "/aplus/2020-11-01/contentDocuments/{contentReferenceKey}"
	path = strings.Replace(path, "{contentReferenceKey}", contentReferenceKey, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("UpdateContentDocument: %w", err)
	}
	return result, nil
}
