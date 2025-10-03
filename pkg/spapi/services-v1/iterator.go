// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package services_v1

import (
	"context"
	"encoding/json"
	"iter"

	"github.com/pkg/errors"
)

// IterateServiceJobs 杩斿洖杩唬鍣紝鑷姩澶勭悊鍒嗛〉銆?
//
// 浣跨敤 Go 1.25 杩唬鍣ㄧ壒鎬э紝鑷姩澶勭悊 nextToken 鍒嗛〉閫昏緫銆?
//
// 绀轰緥:
//   for item, err := range client.IterateServiceJobs(ctx, query) {
//       if err != nil { return err }
//       process(item)
//   }
func (c *Client) IterateServiceJobs(ctx context.Context, query map[string]string) iter.Seq2[map[string]interface{}, error] {
	return func(yield func(map[string]interface{}, error) bool) {
		currentQuery := make(map[string]string)
		for k, v := range query {
			currentQuery[k] = v
		}

		for {
			result, err := c.GetServiceJobs(ctx, currentQuery)
			if err != nil {
				yield(nil, errors.Wrap(err, "failed to call GetServiceJobs"))
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
			items, ok := response["jobs"].([]interface{})
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
			nextToken, _ := response["nextToken"].(string)
			if nextToken == "" {
				break
			}

			currentQuery["nextToken"] = nextToken
		}
	}
}

