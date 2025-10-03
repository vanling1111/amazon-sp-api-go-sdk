// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package reports_v2021_06_30

import (
	"context"
	"fmt"
	"strings"
	
	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client reports API v2021-06-30
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// CreateReport 
// Method: POST | Path: /reports/2021-06-30/reports
func (c *Client) CreateReport(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/reports/2021-06-30/reports"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil { return nil, fmt.Errorf("CreateReport: %w", err) }
	return result, nil
}

// GetReportSchedule 
// Method: GET | Path: /reports/2021-06-30/schedules/{reportScheduleId}
func (c *Client) GetReportSchedule(ctx context.Context, reportScheduleId string, query map[string]string) (interface{}, error) {
	path := "/reports/2021-06-30/schedules/{reportScheduleId}"
	path = strings.Replace(path, "{reportScheduleId}", reportScheduleId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil { return nil, fmt.Errorf("GetReportSchedule: %w", err) }
	return result, nil
}

// GetReportSchedules 
// Method: GET | Path: /reports/2021-06-30/schedules
func (c *Client) GetReportSchedules(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/reports/2021-06-30/schedules"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil { return nil, fmt.Errorf("GetReportSchedules: %w", err) }
	return result, nil
}

// GetReports 
// Method: GET | Path: /reports/2021-06-30/reports
func (c *Client) GetReports(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/reports/2021-06-30/reports"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil { return nil, fmt.Errorf("GetReports: %w", err) }
	return result, nil
}

// GetReport 
// Method: GET | Path: /reports/2021-06-30/reports/{reportId}
func (c *Client) GetReport(ctx context.Context, reportId string, query map[string]string) (interface{}, error) {
	path := "/reports/2021-06-30/reports/{reportId}"
	path = strings.Replace(path, "{reportId}", reportId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil { return nil, fmt.Errorf("GetReport: %w", err) }
	return result, nil
}

// CancelReportSchedule 
// Method: DELETE | Path: /reports/2021-06-30/schedules/{reportScheduleId}
func (c *Client) CancelReportSchedule(ctx context.Context, reportScheduleId string) (interface{}, error) {
	path := "/reports/2021-06-30/schedules/{reportScheduleId}"
	path = strings.Replace(path, "{reportScheduleId}", reportScheduleId, 1)
	var result interface{}
	err := c.baseClient.Delete(ctx, path, &result)
	if err != nil { return nil, fmt.Errorf("CancelReportSchedule: %w", err) }
	return result, nil
}

// GetReportDocument 
// Method: GET | Path: /reports/2021-06-30/documents/{reportDocumentId}
func (c *Client) GetReportDocument(ctx context.Context, reportDocumentId string, query map[string]string) (interface{}, error) {
	path := "/reports/2021-06-30/documents/{reportDocumentId}"
	path = strings.Replace(path, "{reportDocumentId}", reportDocumentId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil { return nil, fmt.Errorf("GetReportDocument: %w", err) }
	return result, nil
}

// CancelReport 
// Method: DELETE | Path: /reports/2021-06-30/reports/{reportId}
func (c *Client) CancelReport(ctx context.Context, reportId string) (interface{}, error) {
	path := "/reports/2021-06-30/reports/{reportId}"
	path = strings.Replace(path, "{reportId}", reportId, 1)
	var result interface{}
	err := c.baseClient.Delete(ctx, path, &result)
	if err != nil { return nil, fmt.Errorf("CancelReport: %w", err) }
	return result, nil
}

// CreateReportSchedule 
// Method: POST | Path: /reports/2021-06-30/schedules
func (c *Client) CreateReportSchedule(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/reports/2021-06-30/schedules"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil { return nil, fmt.Errorf("CreateReportSchedule: %w", err) }
	return result, nil
}
