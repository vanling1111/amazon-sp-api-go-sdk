// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package invoices_v2024_06_19

import (
	"context"
	"fmt"
	"strings"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client invoices API v2024-06-19
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// GetInvoicesExports
// Method: GET | Path: /tax/invoices/2024-06-19/exports
func (c *Client) GetInvoicesExports(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/tax/invoices/2024-06-19/exports"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetInvoicesExports: %w", err)
	}
	return result, nil
}

// GetInvoicesExport
// Method: GET | Path: /tax/invoices/2024-06-19/exports/{exportId}
func (c *Client) GetInvoicesExport(ctx context.Context, exportId string, query map[string]string) (interface{}, error) {
	path := "/tax/invoices/2024-06-19/exports/{exportId}"
	path = strings.Replace(path, "{exportId}", exportId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetInvoicesExport: %w", err)
	}
	return result, nil
}

// GetInvoicesAttributes
// Method: GET | Path: /tax/invoices/2024-06-19/attributes
func (c *Client) GetInvoicesAttributes(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/tax/invoices/2024-06-19/attributes"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetInvoicesAttributes: %w", err)
	}
	return result, nil
}

// GetInvoicesDocument
// Method: GET | Path: /tax/invoices/2024-06-19/documents/{invoicesDocumentId}
func (c *Client) GetInvoicesDocument(ctx context.Context, invoicesDocumentId string, query map[string]string) (interface{}, error) {
	path := "/tax/invoices/2024-06-19/documents/{invoicesDocumentId}"
	path = strings.Replace(path, "{invoicesDocumentId}", invoicesDocumentId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetInvoicesDocument: %w", err)
	}
	return result, nil
}

// GetInvoice
// Method: GET | Path: /tax/invoices/2024-06-19/invoices/{invoiceId}
func (c *Client) GetInvoice(ctx context.Context, invoiceId string, query map[string]string) (interface{}, error) {
	path := "/tax/invoices/2024-06-19/invoices/{invoiceId}"
	path = strings.Replace(path, "{invoiceId}", invoiceId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetInvoice: %w", err)
	}
	return result, nil
}

// CreateInvoicesExport
// Method: POST | Path: /tax/invoices/2024-06-19/exports
func (c *Client) CreateInvoicesExport(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/tax/invoices/2024-06-19/exports"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("CreateInvoicesExport: %w", err)
	}
	return result, nil
}

// GetInvoices
// Method: GET | Path: /tax/invoices/2024-06-19/invoices
func (c *Client) GetInvoices(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/tax/invoices/2024-06-19/invoices"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetInvoices: %w", err)
	}
	return result, nil
}
