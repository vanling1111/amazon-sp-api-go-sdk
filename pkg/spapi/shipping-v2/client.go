// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package shipping_v2

import (
	"context"
	"fmt"
	"strings"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client shipping API v2
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// GetCollectionForm
// Method: GET | Path: /shipping/v2/collectionForms/{collectionFormId}
func (c *Client) GetCollectionForm(ctx context.Context, collectionFormId string, query map[string]string) (interface{}, error) {
	path := "/shipping/v2/collectionForms/{collectionFormId}"
	path = strings.Replace(path, "{collectionFormId}", collectionFormId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetCollectionForm: %w", err)
	}
	return result, nil
}

// GetAccessPoints
// Method: GET | Path: /shipping/v2/accessPoints
func (c *Client) GetAccessPoints(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/shipping/v2/accessPoints"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetAccessPoints: %w", err)
	}
	return result, nil
}

// GetCollectionFormHistory
// Method: PUT | Path: /shipping/v2/collectionForms/history
func (c *Client) GetCollectionFormHistory(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/shipping/v2/collectionForms/history"
	var result interface{}
	err := c.baseClient.Put(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("GetCollectionFormHistory: %w", err)
	}
	return result, nil
}

// CreateClaim
// Method: POST | Path: /shipping/v2/claims
func (c *Client) CreateClaim(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/shipping/v2/claims"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("CreateClaim: %w", err)
	}
	return result, nil
}

// SubmitNdrFeedback
// Method: POST | Path: /shipping/v2/ndrFeedback
func (c *Client) SubmitNdrFeedback(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/shipping/v2/ndrFeedback"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("SubmitNdrFeedback: %w", err)
	}
	return result, nil
}

// GetShipmentDocuments
// Method: GET | Path: /shipping/v2/shipments/{shipmentId}/documents
func (c *Client) GetShipmentDocuments(ctx context.Context, shipmentId string, query map[string]string) (interface{}, error) {
	path := "/shipping/v2/shipments/{shipmentId}/documents"
	path = strings.Replace(path, "{shipmentId}", shipmentId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetShipmentDocuments: %w", err)
	}
	return result, nil
}

// GetCarrierAccounts
// Method: PUT | Path: /shipping/v2/carrierAccounts
func (c *Client) GetCarrierAccounts(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/shipping/v2/carrierAccounts"
	var result interface{}
	err := c.baseClient.Put(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("GetCarrierAccounts: %w", err)
	}
	return result, nil
}

// CancelShipment
// Method: PUT | Path: /shipping/v2/shipments/{shipmentId}/cancel
func (c *Client) CancelShipment(ctx context.Context, shipmentId string, body interface{}) (interface{}, error) {
	path := "/shipping/v2/shipments/{shipmentId}/cancel"
	path = strings.Replace(path, "{shipmentId}", shipmentId, 1)
	var result interface{}
	err := c.baseClient.Put(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("CancelShipment: %w", err)
	}
	return result, nil
}

// GetAdditionalInputs
// Method: GET | Path: /shipping/v2/shipments/additionalInputs/schema
func (c *Client) GetAdditionalInputs(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/shipping/v2/shipments/additionalInputs/schema"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetAdditionalInputs: %w", err)
	}
	return result, nil
}

// GenerateCollectionForm
// Method: POST | Path: /shipping/v2/collectionForms
func (c *Client) GenerateCollectionForm(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/shipping/v2/collectionForms"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("GenerateCollectionForm: %w", err)
	}
	return result, nil
}

// OneClickShipment
// Method: POST | Path: /shipping/v2/oneClickShipment
func (c *Client) OneClickShipment(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/shipping/v2/oneClickShipment"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("OneClickShipment: %w", err)
	}
	return result, nil
}

// PurchaseShipment
// Method: POST | Path: /shipping/v2/shipments
func (c *Client) PurchaseShipment(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/shipping/v2/shipments"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("PurchaseShipment: %w", err)
	}
	return result, nil
}

// UnlinkCarrierAccount
// Method: PUT | Path: /shipping/v2/carrierAccounts/{carrierId}/unlink
func (c *Client) UnlinkCarrierAccount(ctx context.Context, carrierId string, body interface{}) (interface{}, error) {
	path := "/shipping/v2/carrierAccounts/{carrierId}/unlink"
	path = strings.Replace(path, "{carrierId}", carrierId, 1)
	var result interface{}
	err := c.baseClient.Put(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("UnlinkCarrierAccount: %w", err)
	}
	return result, nil
}

// GetTracking
// Method: GET | Path: /shipping/v2/tracking
func (c *Client) GetTracking(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/shipping/v2/tracking"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetTracking: %w", err)
	}
	return result, nil
}

// LinkCarrierAccount
// Method: PUT | Path: /shipping/v2/carrierAccounts/{carrierId}
func (c *Client) LinkCarrierAccount(ctx context.Context, carrierId string, body interface{}) (interface{}, error) {
	path := "/shipping/v2/carrierAccounts/{carrierId}"
	path = strings.Replace(path, "{carrierId}", carrierId, 1)
	var result interface{}
	err := c.baseClient.Put(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("LinkCarrierAccount: %w", err)
	}
	return result, nil
}

// GetUnmanifestedShipments
// Method: PUT | Path: /shipping/v2/unmanifestedShipments
func (c *Client) GetUnmanifestedShipments(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/shipping/v2/unmanifestedShipments"
	var result interface{}
	err := c.baseClient.Put(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("GetUnmanifestedShipments: %w", err)
	}
	return result, nil
}

// DirectPurchaseShipment
// Method: POST | Path: /shipping/v2/shipments/directPurchase
func (c *Client) DirectPurchaseShipment(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/shipping/v2/shipments/directPurchase"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("DirectPurchaseShipment: %w", err)
	}
	return result, nil
}

// GetRates
// Method: POST | Path: /shipping/v2/shipments/rates
func (c *Client) GetRates(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/shipping/v2/shipments/rates"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("GetRates: %w", err)
	}
	return result, nil
}

// GetCarrierAccountFormInputs
// Method: GET | Path: /shipping/v2/carrierAccountFormInputs
func (c *Client) GetCarrierAccountFormInputs(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/shipping/v2/carrierAccountFormInputs"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetCarrierAccountFormInputs: %w", err)
	}
	return result, nil
}
