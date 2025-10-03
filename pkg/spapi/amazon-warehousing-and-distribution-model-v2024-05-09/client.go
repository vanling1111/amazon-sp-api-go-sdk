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

package amazon_warehousing_and_distribution_model_v2024_05_09

import (
	"context"
	"fmt"
	"strings"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client amazon-warehousing-and-distribution-model API v2024-05-09
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// UpdateInbound
// Method: PUT | Path: /awd/2024-05-09/inboundOrders/{orderId}
func (c *Client) UpdateInbound(ctx context.Context, orderId string, body interface{}) (interface{}, error) {
	path := "/awd/2024-05-09/inboundOrders/{orderId}"
	path = strings.Replace(path, "{orderId}", orderId, 1)
	var result interface{}
	err := c.baseClient.Put(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("UpdateInbound: %w", err)
	}
	return result, nil
}

// CheckInboundEligibility
// Method: POST | Path: /awd/2024-05-09/inboundEligibility
func (c *Client) CheckInboundEligibility(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/awd/2024-05-09/inboundEligibility"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("CheckInboundEligibility: %w", err)
	}
	return result, nil
}

// UpdateInboundShipmentTransportDetails
// Method: PUT | Path: /awd/2024-05-09/inboundShipments/{shipmentId}/transport
func (c *Client) UpdateInboundShipmentTransportDetails(ctx context.Context, shipmentId string, body interface{}) (interface{}, error) {
	path := "/awd/2024-05-09/inboundShipments/{shipmentId}/transport"
	path = strings.Replace(path, "{shipmentId}", shipmentId, 1)
	var result interface{}
	err := c.baseClient.Put(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("UpdateInboundShipmentTransportDetails: %w", err)
	}
	return result, nil
}

// ConfirmInbound
// Method: POST | Path: /awd/2024-05-09/inboundOrders/{orderId}/confirmation
func (c *Client) ConfirmInbound(ctx context.Context, orderId string, body interface{}) (interface{}, error) {
	path := "/awd/2024-05-09/inboundOrders/{orderId}/confirmation"
	path = strings.Replace(path, "{orderId}", orderId, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("ConfirmInbound: %w", err)
	}
	return result, nil
}

// GetInbound
// Method: GET | Path: /awd/2024-05-09/inboundOrders/{orderId}
func (c *Client) GetInbound(ctx context.Context, orderId string, query map[string]string) (interface{}, error) {
	path := "/awd/2024-05-09/inboundOrders/{orderId}"
	path = strings.Replace(path, "{orderId}", orderId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetInbound: %w", err)
	}
	return result, nil
}

// GetInboundShipment
// Method: GET | Path: /awd/2024-05-09/inboundShipments/{shipmentId}
func (c *Client) GetInboundShipment(ctx context.Context, shipmentId string, query map[string]string) (interface{}, error) {
	path := "/awd/2024-05-09/inboundShipments/{shipmentId}"
	path = strings.Replace(path, "{shipmentId}", shipmentId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetInboundShipment: %w", err)
	}
	return result, nil
}

// CreateInbound
// Method: POST | Path: /awd/2024-05-09/inboundOrders
func (c *Client) CreateInbound(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/awd/2024-05-09/inboundOrders"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("CreateInbound: %w", err)
	}
	return result, nil
}

// CancelInbound
// Method: POST | Path: /awd/2024-05-09/inboundOrders/{orderId}/cancellation
func (c *Client) CancelInbound(ctx context.Context, orderId string, body interface{}) (interface{}, error) {
	path := "/awd/2024-05-09/inboundOrders/{orderId}/cancellation"
	path = strings.Replace(path, "{orderId}", orderId, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("CancelInbound: %w", err)
	}
	return result, nil
}

// ListInboundShipments
// Method: GET | Path: /awd/2024-05-09/inboundShipments
func (c *Client) ListInboundShipments(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/awd/2024-05-09/inboundShipments"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("ListInboundShipments: %w", err)
	}
	return result, nil
}

// GetInboundShipmentLabels
// Method: GET | Path: /awd/2024-05-09/inboundShipments/{shipmentId}/labels
func (c *Client) GetInboundShipmentLabels(ctx context.Context, shipmentId string, query map[string]string) (interface{}, error) {
	path := "/awd/2024-05-09/inboundShipments/{shipmentId}/labels"
	path = strings.Replace(path, "{shipmentId}", shipmentId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetInboundShipmentLabels: %w", err)
	}
	return result, nil
}

// ListInventory
// Method: GET | Path: /awd/2024-05-09/inventory
func (c *Client) ListInventory(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/awd/2024-05-09/inventory"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("ListInventory: %w", err)
	}
	return result, nil
}
