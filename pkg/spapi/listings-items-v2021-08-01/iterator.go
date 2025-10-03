// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package listings_items_v2021_08_01

import (
	"context"
	"encoding/json"
	"iter"

	"github.com/pkg/errors"
)

// IterateListingsItems 返回 Listings 商品迭代器，自动处理分页。
//
// 使用 Go 1.25 迭代器特性，自动处理分页逻辑。
//
// 参数:
//   - ctx: 请求上下文
//   - sellerId: 卖家 ID
//   - query: 查询参数（包含过滤条件等）
//
// 示例:
//
//	for item, err := range client.IterateListingsItems(ctx, sellerID, query) {
//	    if err != nil { return err }
//	    fmt.Printf("SKU: %s\n", item["sku"])
//	}
func (c *Client) IterateListingsItems(ctx context.Context, sellerId string, query map[string]string) iter.Seq2[map[string]interface{}, error] {
	return func(yield func(map[string]interface{}, error) bool) {
		currentQuery := make(map[string]string)
		for k, v := range query {
			currentQuery[k] = v
		}

		for {
			result, err := c.SearchListingsItems(ctx, sellerId, currentQuery)
			if err != nil {
				yield(nil, errors.Wrap(err, "failed to search listings items"))
				return
			}

			resultBytes, err := json.Marshal(result)
			if err != nil {
				yield(nil, errors.Wrap(err, "failed to marshal result"))
				return
			}

			var response map[string]interface{}
			if err := json.Unmarshal(resultBytes, &response); err != nil {
				yield(nil, errors.Wrap(err, "failed to unmarshal response"))
				return
			}

			// 获取 items 数组
			items, ok := response["items"].([]interface{})
			if !ok || items == nil {
				break
			}

			// 遍历当前页
			for _, item := range items {
				itemMap, ok := item.(map[string]interface{})
				if !ok {
					continue
				}
				if !yield(itemMap, nil) {
					return
				}
			}

			// 检查下一页
			nextToken, _ := response["nextToken"].(string)
			if nextToken == "" {
				break
			}

			currentQuery["pageToken"] = nextToken
		}
	}
}
