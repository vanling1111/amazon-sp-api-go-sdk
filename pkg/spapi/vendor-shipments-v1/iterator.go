// Copyright 2025 Amazon SP-API Go SDK Authors.
//
// This file is part of Amazon SP-API Go SDK.
//
// Amazon SP-API Go SDK is dual-licensed:
//
// 1. GNU Affero General Public License v3.0 (AGPL-3.0) for open source use
//    - Free for personal, educational, and open source projects
//    - Your project must also be open sourced under AGPL-3.0
//    - See: https://www.gnu.org/licenses/agpl-3.0.html
//
// 2. Commercial License for proprietary/commercial use
//    - Required for any commercial, enterprise, or proprietary use
//    - Allows closed source distribution
//    - Contact: vanling1111@gmail.com
//
// Unless you have obtained a commercial license, this file is licensed
// under AGPL-3.0. By using this software, you agree to comply with the
// terms of the applicable license.
//
// This file is part of Amazon SP-API Go SDK.
//
// Amazon SP-API Go SDK is dual-licensed:
//
// 1. GNU Affero General Public License v3.0 (AGPL-3.0) for open source use
//    - Free for personal, educational, and open source projects
//    - Your project must also be open sourced under AGPL-3.0
//    - See: https://www.gnu.org/licenses/agpl-3.0.html
//
// 2. Commercial License for proprietary/commercial use
//    - Required for any commercial, enterprise, or proprietary use
//    - Allows closed source distribution
//    - Contact: vanling1111@gmail.com
//
// Unless you have obtained a commercial license, this file is licensed
// under AGPL-3.0. By using this software, you agree to comply with the
// terms of the applicable license.

package vendor_shipments_v1

import (
	"context"
	"encoding/json"
	"iter"

	"github.com/pkg/errors"
)

// IterateShipmentDetails 返回货件详情迭代器，自动处理分页。
//
// 使用 Go 1.25 迭代器特性，自动处理分页逻辑。
//
// 示例:
//
//	for shipment, err := range client.IterateShipmentDetails(ctx, query) {
//	    if err != nil { return err }
//	    fmt.Printf("Shipment: %s\n", shipment["shipmentId"])
//	}
func (c *Client) IterateShipmentDetails(ctx context.Context, query map[string]string) iter.Seq2[map[string]interface{}, error] {
	return func(yield func(map[string]interface{}, error) bool) {
		currentQuery := make(map[string]string)
		for k, v := range query {
			currentQuery[k] = v
		}

		for {
			result, err := c.GetShipmentDetails(ctx, currentQuery)
			if err != nil {
				yield(nil, errors.Wrap(err, "failed to get shipment details"))
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

			// 获取 shipments 数组
			payload, ok := response["payload"].(map[string]interface{})
			if !ok {
				break
			}

			items, ok := payload["shipments"].([]interface{})
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
