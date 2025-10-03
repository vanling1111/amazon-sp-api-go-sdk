// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package seller_wallet_v2024_03_01

import (
	"context"
	"fmt"
	"strings"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client seller-wallet API v2024-03-01
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// ListAccounts Get all Amazon SW accounts for the seller
// Method: GET | Path: /finances/transfers/wallet/2024-03-01/accounts
func (c *Client) ListAccounts(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/finances/transfers/wallet/2024-03-01/accounts"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("ListAccounts: %w", err)
	}
	return result, nil
}

// ListAccountTransactions The API will return all the transactions for a given Amazon SW account sorted by the transaction request date
// Method: GET | Path: /finances/transfers/wallet/2024-03-01/transactions
func (c *Client) ListAccountTransactions(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/finances/transfers/wallet/2024-03-01/transactions"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("ListAccountTransactions: %w", err)
	}
	return result, nil
}

// GetAccount Find particular Amazon SW account by Amazon account identifier
// Method: GET | Path: /finances/transfers/wallet/2024-03-01/accounts/{accountId}
func (c *Client) GetAccount(ctx context.Context, accountId string, query map[string]string) (interface{}, error) {
	path := "/finances/transfers/wallet/2024-03-01/accounts/{accountId}"
	path = strings.Replace(path, "{accountId}", accountId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetAccount: %w", err)
	}
	return result, nil
}

// CreateTransaction Create a transaction request from Amazon SW account to another customer provided account
// Method: POST | Path: /finances/transfers/wallet/2024-03-01/transactions
func (c *Client) CreateTransaction(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/finances/transfers/wallet/2024-03-01/transactions"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("CreateTransaction: %w", err)
	}
	return result, nil
}

// ListAccountBalances Find balance in particular Amazon SW account by Amazon account identifier
// Method: GET | Path: /finances/transfers/wallet/2024-03-01/accounts/{accountId}/balance
func (c *Client) ListAccountBalances(ctx context.Context, accountId string, query map[string]string) (interface{}, error) {
	path := "/finances/transfers/wallet/2024-03-01/accounts/{accountId}/balance"
	path = strings.Replace(path, "{accountId}", accountId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("ListAccountBalances: %w", err)
	}
	return result, nil
}

// GetTransferSchedule Find particular Amazon Seller Wallet account transfer schedule by Amazon transfer schedule identifier
// Method: GET | Path: /finances/transfers/wallet/2024-03-01/transferSchedules/{transferScheduleId}
func (c *Client) GetTransferSchedule(ctx context.Context, transferScheduleId string, query map[string]string) (interface{}, error) {
	path := "/finances/transfers/wallet/2024-03-01/transferSchedules/{transferScheduleId}"
	path = strings.Replace(path, "{transferScheduleId}", transferScheduleId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetTransferSchedule: %w", err)
	}
	return result, nil
}

// GetTransferPreview Fetch potential fees that could be applied on a transaction on the basis of the source and destination country currency code
// Method: GET | Path: /finances/transfers/wallet/2024-03-01/transferPreview
func (c *Client) GetTransferPreview(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/finances/transfers/wallet/2024-03-01/transferPreview"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetTransferPreview: %w", err)
	}
	return result, nil
}

// UpdateTransferSchedule Update a transfer schedule information. Only fields (i.e; transferScheduleInformation, paymentPreference, transferScheduleStatus) in the request body can be updated.
// Method: PUT | Path: /finances/transfers/wallet/2024-03-01/transferSchedules
func (c *Client) UpdateTransferSchedule(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/finances/transfers/wallet/2024-03-01/transferSchedules"
	var result interface{}
	err := c.baseClient.Put(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("UpdateTransferSchedule: %w", err)
	}
	return result, nil
}

// ListTransferSchedules The API will return all the transfer schedules for a given Amazon SW account
// Method: GET | Path: /finances/transfers/wallet/2024-03-01/transferSchedules
func (c *Client) ListTransferSchedules(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/finances/transfers/wallet/2024-03-01/transferSchedules"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("ListTransferSchedules: %w", err)
	}
	return result, nil
}

// DeleteScheduleTransaction Delete a transaction request that is scheduled from Amazon Seller Wallet account to another customer-provided account
// Method: DELETE | Path: /finances/transfers/wallet/2024-03-01/transferSchedules/{transferScheduleId}
func (c *Client) DeleteScheduleTransaction(ctx context.Context, transferScheduleId string) (interface{}, error) {
	path := "/finances/transfers/wallet/2024-03-01/transferSchedules/{transferScheduleId}"
	path = strings.Replace(path, "{transferScheduleId}", transferScheduleId, 1)
	var result interface{}
	err := c.baseClient.Delete(ctx, path, &result)
	if err != nil {
		return nil, fmt.Errorf("DeleteScheduleTransaction: %w", err)
	}
	return result, nil
}

// GetTransaction Find particular Amazon SW account transaction by Amazon transaction identifier
// Method: GET | Path: /finances/transfers/wallet/2024-03-01/transactions/{transactionId}
func (c *Client) GetTransaction(ctx context.Context, transactionId string, query map[string]string) (interface{}, error) {
	path := "/finances/transfers/wallet/2024-03-01/transactions/{transactionId}"
	path = strings.Replace(path, "{transactionId}", transactionId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetTransaction: %w", err)
	}
	return result, nil
}

// CreateTransferSchedule Create a transfer schedule request from Amazon SW account to another customer provided account
// Method: POST | Path: /finances/transfers/wallet/2024-03-01/transferSchedules
func (c *Client) CreateTransferSchedule(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/finances/transfers/wallet/2024-03-01/transferSchedules"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("CreateTransferSchedule: %w", err)
	}
	return result, nil
}
