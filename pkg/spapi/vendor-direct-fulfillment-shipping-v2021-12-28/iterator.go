// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package vendor_direct_fulfillment_shipping_v2021_12_28

import (
	"context"
	"encoding/json"
	"iter"

	"github.com/pkg/errors"
)

// IterateShippingLabels 返回物流标签迭代器，自动处理分页。
func (c *Client) IterateShippingLabels(ctx context.Context, query map[string]string) iter.Seq2[map[string]interface{}, error] {
	return func(yield func(map[string]interface{}, error) bool) {
		currentQuery := make(map[string]string)
		for k, v := range query {
			currentQuery[k] = v
		}

		for {
			result, err := c.GetShippingLabels(ctx, currentQuery)
			if err != nil {
				yield(nil, errors.Wrap(err, "failed to get shipping labels"))
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

			payload, ok := response["payload"].(map[string]interface{})
			if !ok {
				break
			}

			items, ok := payload["shippingLabels"].([]interface{})
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

			pagination, ok := payload["pagination"].(map[string]interface{})
			if !ok {
				break
			}

			nextToken, _ := pagination["nextToken"].(string)
			if nextToken == "" {
				break
			}

			currentQuery["nextToken"] = nextToken
		}
	}
}
