// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package catalog_items_v2022_04_01

import (
	"context"
	"encoding/json"
	"iter"

	"github.com/pkg/errors"
)

// SearchResponse 商品搜索响应
type SearchResponse struct {
	Items           []map[string]interface{} `json:"items"`
	NextToken       string                   `json:"nextToken,omitempty"`
	Refinements     interface{}              `json:"refinements,omitempty"`
	NumberOfResults int                      `json:"numberOfResults,omitempty"`
}

// IterateCatalogItems 返回商品目录迭代器，自动处理分页。
//
// 此方法使用 Go 1.25 的迭代器特性，自动处理分页逻辑。
//
// 参数:
//   - ctx: 请求上下文
//   - query: 查询参数（keywords, marketplaceIds 等）
//
// 返回值:
//   - iter.Seq2[map[string]interface{}, error]: 商品迭代器
//
// 示例:
//
//	query := map[string]string{
//	    "keywords":       "laptop",
//	    "marketplaceIds": "ATVPDKIKX0DER",
//	}
//
//	for item, err := range client.IterateCatalogItems(ctx, query) {
//	    if err != nil {
//	        log.Fatal(err)
//	    }
//	    fmt.Printf("ASIN: %s, Title: %s\n",
//	        item["asin"], item["summaries"].([]interface{})[0].(map[string]interface{})["itemName"])
//	}
func (c *Client) IterateCatalogItems(ctx context.Context, query map[string]string) iter.Seq2[map[string]interface{}, error] {
	return func(yield func(map[string]interface{}, error) bool) {
		currentQuery := make(map[string]string)
		for k, v := range query {
			currentQuery[k] = v
		}

		for {
			// 调用 SearchCatalogItems API
			result, err := c.SearchCatalogItems(ctx, currentQuery)
			if err != nil {
				yield(nil, errors.Wrap(err, "failed to search catalog items"))
				return
			}

			// 解析响应
			resultBytes, err := json.Marshal(result)
			if err != nil {
				yield(nil, errors.Wrap(err, "failed to marshal result"))
				return
			}

			var searchResp SearchResponse
			if err := json.Unmarshal(resultBytes, &searchResp); err != nil {
				yield(nil, errors.Wrap(err, "failed to unmarshal search response"))
				return
			}

			// 遍历当前页的商品
			for _, item := range searchResp.Items {
				if !yield(item, nil) {
					return
				}
			}

			// 检查是否还有下一页
			if searchResp.NextToken == "" {
				break
			}

			currentQuery["pageToken"] = searchResp.NextToken
		}
	}
}
