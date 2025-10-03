// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package orders_v0

import (
	"context"
	"encoding/json"
	"iter"

	"github.com/pkg/errors"
)

// OrdersResponse 订单列表响应
type OrdersResponse struct {
	Payload struct {
		Orders    []map[string]interface{} `json:"Orders"`
		NextToken string                   `json:"NextToken,omitempty"`
	} `json:"payload"`
}

// IterateOrders 返回订单迭代器，自动处理分页。
//
// 此方法使用 Go 1.25 的迭代器特性，自动处理 NextToken 分页逻辑。
// 用户无需手动管理 NextToken，可以直接遍历所有订单。
//
// 参数:
//   - ctx: 请求上下文
//   - query: 查询参数（不需要设置 NextToken）
//
// 返回值:
//   - iter.Seq2[map[string]interface{}, error]: 订单迭代器
//
// 示例:
//
//	query := map[string]string{
//	    "MarketplaceIds": "ATVPDKIKX0DER",
//	    "CreatedAfter": "2025-01-01T00:00:00Z",
//	}
//
//	for order, err := range client.IterateOrders(ctx, query) {
//	    if err != nil {
//	        log.Fatal(err)
//	    }
//	    fmt.Printf("Order: %s\n", order["AmazonOrderId"])
//	}
//
// 注意:
//   - 自动处理所有分页
//   - 遇到错误时立即返回
//   - 支持提前退出（break）
func (c *Client) IterateOrders(ctx context.Context, query map[string]string) iter.Seq2[map[string]interface{}, error] {
	return func(yield func(map[string]interface{}, error) bool) {
		// 复制 query，避免修改原始参数
		currentQuery := make(map[string]string)
		for k, v := range query {
			currentQuery[k] = v
		}

		for {
			// 调用 GetOrders API
			result, err := c.GetOrders(ctx, currentQuery)
			if err != nil {
				yield(nil, errors.Wrap(err, "failed to get orders"))
				return
			}

			// 解析响应
			resultBytes, err := json.Marshal(result)
			if err != nil {
				yield(nil, errors.Wrap(err, "failed to marshal result"))
				return
			}

			var ordersResp OrdersResponse
			if err := json.Unmarshal(resultBytes, &ordersResp); err != nil {
				yield(nil, errors.Wrap(err, "failed to unmarshal orders response"))
				return
			}

			// 遍历当前页的订单
			for _, order := range ordersResp.Payload.Orders {
				if !yield(order, nil) {
					// 用户提前退出（break）
					return
				}
			}

			// 检查是否还有下一页
			if ordersResp.Payload.NextToken == "" {
				// 没有更多数据
				break
			}

			// 设置 NextToken 继续获取下一页
			currentQuery["NextToken"] = ordersResp.Payload.NextToken
		}
	}
}

// OrderItemsResponse 订单项响应
type OrderItemsResponse struct {
	Payload struct {
		OrderItems []map[string]interface{} `json:"OrderItems"`
		NextToken  string                   `json:"NextToken,omitempty"`
	} `json:"payload"`
}

// IterateOrderItems 返回订单项迭代器，自动处理分页。
//
// 此方法自动获取指定订单的所有订单项，处理分页逻辑。
//
// 参数:
//   - ctx: 请求上下文
//   - orderID: 订单 ID
//   - query: 可选查询参数
//
// 返回值:
//   - iter.Seq2[map[string]interface{}, error]: 订单项迭代器
//
// 示例:
//
//	for item, err := range client.IterateOrderItems(ctx, "123-4567890-1234567", nil) {
//	    if err != nil {
//	        log.Fatal(err)
//	    }
//	    fmt.Printf("Item: %s, Quantity: %v\n",
//	        item["SellerSKU"], item["QuantityOrdered"])
//	}
func (c *Client) IterateOrderItems(ctx context.Context, orderID string, query map[string]string) iter.Seq2[map[string]interface{}, error] {
	return func(yield func(map[string]interface{}, error) bool) {
		// 复制 query
		currentQuery := make(map[string]string)
		if query != nil {
			for k, v := range query {
				currentQuery[k] = v
			}
		}

		for {
			// 调用 GetOrderItems API
			result, err := c.GetOrderItems(ctx, orderID, currentQuery)
			if err != nil {
				yield(nil, errors.Wrap(err, "failed to get order items"))
				return
			}

			// 解析响应
			resultBytes, err := json.Marshal(result)
			if err != nil {
				yield(nil, errors.Wrap(err, "failed to marshal result"))
				return
			}

			var itemsResp OrderItemsResponse
			if err := json.Unmarshal(resultBytes, &itemsResp); err != nil {
				yield(nil, errors.Wrap(err, "failed to unmarshal order items response"))
				return
			}

			// 遍历当前页的订单项
			for _, item := range itemsResp.Payload.OrderItems {
				if !yield(item, nil) {
					// 用户提前退出
					return
				}
			}

			// 检查是否还有下一页
			if itemsResp.Payload.NextToken == "" {
				break
			}

			// 设置 NextToken
			currentQuery["NextToken"] = itemsResp.Payload.NextToken
		}
	}
}
