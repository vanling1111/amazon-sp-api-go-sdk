// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package reports_v2021_06_30

import (
	"context"
	"encoding/json"
	"iter"

	"github.com/pkg/errors"
)

// ReportsResponse 报告列表响应
type ReportsResponse struct {
	Reports   []map[string]interface{} `json:"reports"`
	NextToken string                   `json:"nextToken,omitempty"`
}

// IterateReports 返回报告迭代器，自动处理分页。
//
// 此方法使用 Go 1.25 的迭代器特性，自动处理 nextToken 分页逻辑。
//
// 参数:
//   - ctx: 请求上下文
//   - query: 查询参数（reportTypes, marketplaceIds 等）
//
// 返回值:
//   - iter.Seq2[map[string]interface{}, error]: 报告迭代器
//
// 示例:
//
//	query := map[string]string{
//	    "reportTypes":    "GET_FLAT_FILE_ALL_ORDERS_DATA_BY_ORDER_DATE",
//	    "marketplaceIds": "ATVPDKIKX0DER",
//	}
//
//	for report, err := range client.IterateReports(ctx, query) {
//	    if err != nil {
//	        log.Fatal(err)
//	    }
//	    fmt.Printf("Report: %s, Status: %s\n",
//	        report["reportId"], report["processingStatus"])
//	}
func (c *Client) IterateReports(ctx context.Context, query map[string]string) iter.Seq2[map[string]interface{}, error] {
	return func(yield func(map[string]interface{}, error) bool) {
		currentQuery := make(map[string]string)
		for k, v := range query {
			currentQuery[k] = v
		}

		for {
			// 调用 GetReports API
			result, err := c.GetReports(ctx, currentQuery)
			if err != nil {
				yield(nil, errors.Wrap(err, "failed to get reports"))
				return
			}

			// 解析响应
			resultBytes, err := json.Marshal(result)
			if err != nil {
				yield(nil, errors.Wrap(err, "failed to marshal result"))
				return
			}

			var reportsResp ReportsResponse
			if err := json.Unmarshal(resultBytes, &reportsResp); err != nil {
				yield(nil, errors.Wrap(err, "failed to unmarshal reports response"))
				return
			}

			// 遍历当前页的报告
			for _, report := range reportsResp.Reports {
				if !yield(report, nil) {
					return
				}
			}

			// 检查是否还有下一页
			if reportsResp.NextToken == "" {
				break
			}

			currentQuery["nextToken"] = reportsResp.NextToken
		}
	}
}

