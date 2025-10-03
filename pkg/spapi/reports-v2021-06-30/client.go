// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package reports_v2021_06_30

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/vanling1111/amazon-sp-api-go-sdk/internal/crypto"
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

// GetReportDocumentDecrypted 获取并自动解密报告文档。
//
// 此方法封装了完整的报告下载和解密流程：
// 1. 调用 GetReportDocument API 获取报告元数据
// 2. 从返回的 URL 下载加密的报告内容
// 3. 如果报告是加密的，自动解密
// 4. 返回解密后的原始报告数据
//
// 参数:
//   - ctx: 请求上下文
//   - reportDocumentID: 报告文档 ID
//
// 返回值:
//   - []byte: 解密后的报告内容
//   - error: 如果获取或解密失败，返回错误
//
// 示例:
//
//	// 创建报告
//	createResp, _ := client.Reports.CreateReport(ctx, &CreateReportRequest{
//	    ReportType: "GET_FLAT_FILE_ALL_ORDERS_DATA_BY_ORDER_DATE",
//	    MarketplaceIds: []string{"ATVPDKIKX0DER"},
//	})
//
//	// 等待报告生成...
//	reportID := createResp["reportId"].(string)
//	report, _ := client.Reports.GetReport(ctx, reportID, nil)
//	reportDocumentID := report["reportDocumentId"].(string)
//
//	// 自动下载并解密报告
//	decryptedData, err := client.Reports.GetReportDocumentDecrypted(ctx, reportDocumentID)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// 使用解密后的数据
//	fmt.Println(string(decryptedData))
func (c *Client) GetReportDocumentDecrypted(ctx context.Context, reportDocumentID string) ([]byte, error) {
	// 1. 获取报告文档元数据
	docResult, err := c.GetReportDocument(ctx, reportDocumentID, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get report document metadata")
	}

	// 解析响应
	docBytes, err := json.Marshal(docResult)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal document result")
	}

	var docResp struct {
		ReportDocumentID string `json:"reportDocumentId"`
		URL              string `json:"url"`
		EncryptionDetails *struct {
			Standard             string `json:"standard"`
			InitializationVector string `json:"initializationVector"`
			Key                  string `json:"key"`
		} `json:"encryptionDetails,omitempty"`
		CompressionAlgorithm string `json:"compressionAlgorithm,omitempty"`
	}

	if err := json.Unmarshal(docBytes, &docResp); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal document response")
	}

	if docResp.URL == "" {
		return nil, errors.New("document URL is empty")
	}

	// 2. 下载报告内容
	httpResp, err := http.Get(docResp.URL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to download report content")
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("download failed with status: %d", httpResp.StatusCode)
	}

	encryptedData, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read report content")
	}

	// 3. 解密（如果需要）
	if docResp.EncryptionDetails != nil {
		// 验证加密详情
		details := &crypto.EncryptionDetails{
			Standard:             docResp.EncryptionDetails.Standard,
			InitializationVector: docResp.EncryptionDetails.InitializationVector,
			Key:                  docResp.EncryptionDetails.Key,
		}

		if err := crypto.ValidateEncryptionDetails(details); err != nil {
			return nil, errors.Wrap(err, "invalid encryption details")
		}

		// 解密
		decrypted, err := crypto.DecryptReport(details.Key, details.InitializationVector, encryptedData)
		if err != nil {
			return nil, errors.Wrap(err, "failed to decrypt report")
		}

		return decrypted, nil
	}

	// 未加密，直接返回
	return encryptedData, nil
}
