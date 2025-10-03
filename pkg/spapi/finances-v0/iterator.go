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

package finances_v0

import (
	"context"
	"encoding/json"
	"iter"

	"github.com/pkg/errors"
)

// IterateFinancialEvents 杩斿洖杩唬鍣紝鑷姩澶勭悊鍒嗛〉銆?
//
// 浣跨敤 Go 1.25 杩唬鍣ㄧ壒鎬э紝鑷姩澶勭悊 NextToken 鍒嗛〉閫昏緫銆?
//
// 绀轰緥:
//
//	for item, err := range client.IterateFinancialEvents(ctx, query) {
//	    if err != nil { return err }
//	    process(item)
//	}
func (c *Client) IterateFinancialEvents(ctx context.Context, query map[string]string) iter.Seq2[map[string]interface{}, error] {
	return func(yield func(map[string]interface{}, error) bool) {
		currentQuery := make(map[string]string)
		for k, v := range query {
			currentQuery[k] = v
		}

		for {
			result, err := c.ListFinancialEvents(ctx, currentQuery)
			if err != nil {
				yield(nil, errors.Wrap(err, "failed to call ListFinancialEvents"))
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

			// 鑾峰彇鏁版嵁鏁扮粍
			items, ok := response["FinancialEvents"].([]interface{})
			if !ok || items == nil {
				break
			}

			// 閬嶅巻褰撳墠椤?
			for _, item := range items {
				itemMap, ok := item.(map[string]interface{})
				if !ok {
					continue
				}
				if !yield(itemMap, nil) {
					return
				}
			}

			// 妫€鏌ヤ笅涓€椤?
			nextToken, _ := response["NextToken"].(string)
			if nextToken == "" {
				break
			}

			currentQuery["NextToken"] = nextToken
		}
	}
}

// IterateFinancialEventGroups 杩斿洖杩唬鍣紝鑷姩澶勭悊鍒嗛〉銆?
//
// 浣跨敤 Go 1.25 杩唬鍣ㄧ壒鎬э紝鑷姩澶勭悊 NextToken 鍒嗛〉閫昏緫銆?
//
// 绀轰緥:
//
//	for item, err := range client.IterateFinancialEventGroups(ctx, query) {
//	    if err != nil { return err }
//	    process(item)
//	}
func (c *Client) IterateFinancialEventGroups(ctx context.Context, query map[string]string) iter.Seq2[map[string]interface{}, error] {
	return func(yield func(map[string]interface{}, error) bool) {
		currentQuery := make(map[string]string)
		for k, v := range query {
			currentQuery[k] = v
		}

		for {
			result, err := c.ListFinancialEventGroups(ctx, currentQuery)
			if err != nil {
				yield(nil, errors.Wrap(err, "failed to call ListFinancialEventGroups"))
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

			// 鑾峰彇鏁版嵁鏁扮粍
			items, ok := response["FinancialEventGroupList"].([]interface{})
			if !ok || items == nil {
				break
			}

			// 閬嶅巻褰撳墠椤?
			for _, item := range items {
				itemMap, ok := item.(map[string]interface{})
				if !ok {
					continue
				}
				if !yield(itemMap, nil) {
					return
				}
			}

			// 妫€鏌ヤ笅涓€椤?
			nextToken, _ := response["NextToken"].(string)
			if nextToken == "" {
				break
			}

			currentQuery["NextToken"] = nextToken
		}
	}
}
