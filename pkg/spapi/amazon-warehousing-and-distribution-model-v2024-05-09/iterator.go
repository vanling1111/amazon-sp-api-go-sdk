// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package amazon_warehousing_and_distribution_model_v2024_05_09

import (
	"context"
	"encoding/json"
	"iter"

	"github.com/pkg/errors"
)

// IterateInbound 返回入库货件迭代器，自动处理分页。
func (c *Client) IterateInboundShipments(ctx context.Context, query map[string]string) iter.Seq2[map[string]interface{}, error] {
	return func(yield func(map[string]interface{}, error) bool) {
		currentQuery := make(map[string]string)
		for k, v := range query {
			currentQuery[k] = v
		}

		for {
			result, err := c.ListInboundShipments(ctx, currentQuery)
			if err != nil {
				yield(nil, errors.Wrap(err, "failed to list inbound shipments"))
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

			items, ok := response["shipments"].([]interface{})
			if !ok || items == nil {
				break
			}

			for _, item := range items {
				itemMap, ok := item.(map[string]interface{})
				if !ok {
					continue
				}
				if !yield(itemMap, nil) {
					return
				}
			}

			nextToken, _ := response["nextToken"].(string)
			if nextToken == "" {
				break
			}

			currentQuery["nextToken"] = nextToken
		}
	}
}

