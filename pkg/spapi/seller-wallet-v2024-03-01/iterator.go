// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package seller_wallet_v2024_03_01

import (
	"context"
	"encoding/json"
	"iter"

	"github.com/pkg/errors"
)

// IterateAccountTransactions 返回账户交易迭代器，自动处理分页。
//
// 使用 Go 1.25 迭代器特性，自动处理分页逻辑。
//
// 示例:
//
//	for transaction, err := range client.IterateAccountTransactions(ctx, query) {
//	    if err != nil { return err }
//	    fmt.Printf("Transaction: %s\n", transaction["transactionId"])
//	}
func (c *Client) IterateAccountTransactions(ctx context.Context, query map[string]string) iter.Seq2[map[string]interface{}, error] {
	return func(yield func(map[string]interface{}, error) bool) {
		currentQuery := make(map[string]string)
		for k, v := range query {
			currentQuery[k] = v
		}

		for {
			result, err := c.ListAccountTransactions(ctx, currentQuery)
			if err != nil {
				yield(nil, errors.Wrap(err, "failed to list account transactions"))
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

			// 获取 transactions 数组
			items, ok := response["transactions"].([]interface{})
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

			currentQuery["nextToken"] = nextToken
		}
	}
}

