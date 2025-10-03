// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package fulfillment_inbound_v2024_03_20

import (
	"context"
	"fmt"
	"strings"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client fulfillment-inbound API v2024-03-20
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// ListInboundPlans
// Method: GET | Path: /inbound/fba/2024-03-20/inboundPlans
func (c *Client) ListInboundPlans(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/inboundPlans"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("ListInboundPlans: %w", err)
	}
	return result, nil
}

// SetPrepDetails
// Method: POST | Path: /inbound/fba/2024-03-20/items/prepDetails
func (c *Client) SetPrepDetails(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/items/prepDetails"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("SetPrepDetails: %w", err)
	}
	return result, nil
}

// UpdateInboundPlanName
// Method: PUT | Path: /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/name
func (c *Client) UpdateInboundPlanName(ctx context.Context, inboundPlanId string, body interface{}) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/name"
	path = strings.Replace(path, "{inboundPlanId}", inboundPlanId, 1)
	var result interface{}
	err := c.baseClient.Put(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("UpdateInboundPlanName: %w", err)
	}
	return result, nil
}

// ConfirmTransportationOptions
// Method: POST | Path: /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/transportationOptions/confirmation
func (c *Client) ConfirmTransportationOptions(ctx context.Context, inboundPlanId string, body interface{}) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/transportationOptions/confirmation"
	path = strings.Replace(path, "{inboundPlanId}", inboundPlanId, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("ConfirmTransportationOptions: %w", err)
	}
	return result, nil
}

// GetDeliveryChallanDocument
// Method: GET | Path: /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/deliveryChallanDocument
func (c *Client) GetDeliveryChallanDocument(ctx context.Context, inboundPlanId string, shipmentId string, query map[string]string) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/deliveryChallanDocument"
	path = strings.Replace(path, "{inboundPlanId}", inboundPlanId, 1)
	path = strings.Replace(path, "{shipmentId}", shipmentId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetDeliveryChallanDocument: %w", err)
	}
	return result, nil
}

// ListPrepDetails
// Method: GET | Path: /inbound/fba/2024-03-20/items/prepDetails
func (c *Client) ListPrepDetails(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/items/prepDetails"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("ListPrepDetails: %w", err)
	}
	return result, nil
}

// ScheduleSelfShipAppointment
// Method: POST | Path: /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots/{slotId}/schedule
func (c *Client) ScheduleSelfShipAppointment(ctx context.Context, inboundPlanId string, shipmentId string, slotId string, body interface{}) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots/{slotId}/schedule"
	path = strings.Replace(path, "{inboundPlanId}", inboundPlanId, 1)
	path = strings.Replace(path, "{shipmentId}", shipmentId, 1)
	path = strings.Replace(path, "{slotId}", slotId, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("ScheduleSelfShipAppointment: %w", err)
	}
	return result, nil
}

// CancelInboundPlan
// Method: PUT | Path: /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/cancellation
func (c *Client) CancelInboundPlan(ctx context.Context, inboundPlanId string, body interface{}) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/cancellation"
	path = strings.Replace(path, "{inboundPlanId}", inboundPlanId, 1)
	var result interface{}
	err := c.baseClient.Put(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("CancelInboundPlan: %w", err)
	}
	return result, nil
}

// GeneratePlacementOptions
// Method: POST | Path: /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/placementOptions
func (c *Client) GeneratePlacementOptions(ctx context.Context, inboundPlanId string, body interface{}) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/placementOptions"
	path = strings.Replace(path, "{inboundPlanId}", inboundPlanId, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("GeneratePlacementOptions: %w", err)
	}
	return result, nil
}

// ConfirmDeliveryWindowOptions
// Method: POST | Path: /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/deliveryWindowOptions/{deliveryWindowOptionId}/confirmation
func (c *Client) ConfirmDeliveryWindowOptions(ctx context.Context, inboundPlanId string, shipmentId string, deliveryWindowOptionId string, body interface{}) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/deliveryWindowOptions/{deliveryWindowOptionId}/confirmation"
	path = strings.Replace(path, "{inboundPlanId}", inboundPlanId, 1)
	path = strings.Replace(path, "{shipmentId}", shipmentId, 1)
	path = strings.Replace(path, "{deliveryWindowOptionId}", deliveryWindowOptionId, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("ConfirmDeliveryWindowOptions: %w", err)
	}
	return result, nil
}

// GetInboundOperationStatus
// Method: GET | Path: /inbound/fba/2024-03-20/operations/{operationId}
func (c *Client) GetInboundOperationStatus(ctx context.Context, operationId string, query map[string]string) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/operations/{operationId}"
	path = strings.Replace(path, "{operationId}", operationId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetInboundOperationStatus: %w", err)
	}
	return result, nil
}

// ListInboundPlanItems
// Method: GET | Path: /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/items
func (c *Client) ListInboundPlanItems(ctx context.Context, inboundPlanId string, query map[string]string) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/items"
	path = strings.Replace(path, "{inboundPlanId}", inboundPlanId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("ListInboundPlanItems: %w", err)
	}
	return result, nil
}

// ListPackingOptions
// Method: GET | Path: /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/packingOptions
func (c *Client) ListPackingOptions(ctx context.Context, inboundPlanId string, query map[string]string) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/packingOptions"
	path = strings.Replace(path, "{inboundPlanId}", inboundPlanId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("ListPackingOptions: %w", err)
	}
	return result, nil
}

// UpdateShipmentSourceAddress
// Method: PUT | Path: /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/sourceAddress
func (c *Client) UpdateShipmentSourceAddress(ctx context.Context, inboundPlanId string, shipmentId string, body interface{}) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/sourceAddress"
	path = strings.Replace(path, "{inboundPlanId}", inboundPlanId, 1)
	path = strings.Replace(path, "{shipmentId}", shipmentId, 1)
	var result interface{}
	err := c.baseClient.Put(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("UpdateShipmentSourceAddress: %w", err)
	}
	return result, nil
}

// UpdateShipmentName
// Method: PUT | Path: /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/name
func (c *Client) UpdateShipmentName(ctx context.Context, inboundPlanId string, shipmentId string, body interface{}) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/name"
	path = strings.Replace(path, "{inboundPlanId}", inboundPlanId, 1)
	path = strings.Replace(path, "{shipmentId}", shipmentId, 1)
	var result interface{}
	err := c.baseClient.Put(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("UpdateShipmentName: %w", err)
	}
	return result, nil
}

// ListItemComplianceDetails
// Method: GET | Path: /inbound/fba/2024-03-20/items/compliance
func (c *Client) ListItemComplianceDetails(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/items/compliance"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("ListItemComplianceDetails: %w", err)
	}
	return result, nil
}

// GetShipmentContentUpdatePreview
// Method: GET | Path: /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/contentUpdatePreviews/{contentUpdatePreviewId}
func (c *Client) GetShipmentContentUpdatePreview(ctx context.Context, inboundPlanId string, shipmentId string, contentUpdatePreviewId string, query map[string]string) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/contentUpdatePreviews/{contentUpdatePreviewId}"
	path = strings.Replace(path, "{inboundPlanId}", inboundPlanId, 1)
	path = strings.Replace(path, "{shipmentId}", shipmentId, 1)
	path = strings.Replace(path, "{contentUpdatePreviewId}", contentUpdatePreviewId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetShipmentContentUpdatePreview: %w", err)
	}
	return result, nil
}

// CreateInboundPlan
// Method: POST | Path: /inbound/fba/2024-03-20/inboundPlans
func (c *Client) CreateInboundPlan(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/inboundPlans"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("CreateInboundPlan: %w", err)
	}
	return result, nil
}

// ListPackingGroupBoxes
// Method: GET | Path: /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/packingGroups/{packingGroupId}/boxes
func (c *Client) ListPackingGroupBoxes(ctx context.Context, inboundPlanId string, packingGroupId string, query map[string]string) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/packingGroups/{packingGroupId}/boxes"
	path = strings.Replace(path, "{inboundPlanId}", inboundPlanId, 1)
	path = strings.Replace(path, "{packingGroupId}", packingGroupId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("ListPackingGroupBoxes: %w", err)
	}
	return result, nil
}

// UpdateItemComplianceDetails
// Method: PUT | Path: /inbound/fba/2024-03-20/items/compliance
func (c *Client) UpdateItemComplianceDetails(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/items/compliance"
	var result interface{}
	err := c.baseClient.Put(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("UpdateItemComplianceDetails: %w", err)
	}
	return result, nil
}

// GeneratePackingOptions
// Method: POST | Path: /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/packingOptions
func (c *Client) GeneratePackingOptions(ctx context.Context, inboundPlanId string, body interface{}) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/packingOptions"
	path = strings.Replace(path, "{inboundPlanId}", inboundPlanId, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("GeneratePackingOptions: %w", err)
	}
	return result, nil
}

// CancelSelfShipAppointment
// Method: PUT | Path: /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentCancellation
func (c *Client) CancelSelfShipAppointment(ctx context.Context, inboundPlanId string, shipmentId string, body interface{}) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentCancellation"
	path = strings.Replace(path, "{inboundPlanId}", inboundPlanId, 1)
	path = strings.Replace(path, "{shipmentId}", shipmentId, 1)
	var result interface{}
	err := c.baseClient.Put(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("CancelSelfShipAppointment: %w", err)
	}
	return result, nil
}

// SetPackingInformation
// Method: POST | Path: /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/packingInformation
func (c *Client) SetPackingInformation(ctx context.Context, inboundPlanId string, body interface{}) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/packingInformation"
	path = strings.Replace(path, "{inboundPlanId}", inboundPlanId, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("SetPackingInformation: %w", err)
	}
	return result, nil
}

// UpdateShipmentTrackingDetails
// Method: PUT | Path: /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/trackingDetails
func (c *Client) UpdateShipmentTrackingDetails(ctx context.Context, inboundPlanId string, shipmentId string, body interface{}) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/trackingDetails"
	path = strings.Replace(path, "{inboundPlanId}", inboundPlanId, 1)
	path = strings.Replace(path, "{shipmentId}", shipmentId, 1)
	var result interface{}
	err := c.baseClient.Put(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("UpdateShipmentTrackingDetails: %w", err)
	}
	return result, nil
}

// GenerateTransportationOptions
// Method: POST | Path: /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/transportationOptions
func (c *Client) GenerateTransportationOptions(ctx context.Context, inboundPlanId string, body interface{}) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/transportationOptions"
	path = strings.Replace(path, "{inboundPlanId}", inboundPlanId, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("GenerateTransportationOptions: %w", err)
	}
	return result, nil
}

// GetInboundPlan
// Method: GET | Path: /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}
func (c *Client) GetInboundPlan(ctx context.Context, inboundPlanId string, query map[string]string) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}"
	path = strings.Replace(path, "{inboundPlanId}", inboundPlanId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetInboundPlan: %w", err)
	}
	return result, nil
}

// ConfirmPlacementOption
// Method: POST | Path: /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/placementOptions/{placementOptionId}/confirmation
func (c *Client) ConfirmPlacementOption(ctx context.Context, inboundPlanId string, placementOptionId string, body interface{}) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/placementOptions/{placementOptionId}/confirmation"
	path = strings.Replace(path, "{inboundPlanId}", inboundPlanId, 1)
	path = strings.Replace(path, "{placementOptionId}", placementOptionId, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("ConfirmPlacementOption: %w", err)
	}
	return result, nil
}

// ListShipmentPallets
// Method: GET | Path: /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/pallets
func (c *Client) ListShipmentPallets(ctx context.Context, inboundPlanId string, shipmentId string, query map[string]string) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/pallets"
	path = strings.Replace(path, "{inboundPlanId}", inboundPlanId, 1)
	path = strings.Replace(path, "{shipmentId}", shipmentId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("ListShipmentPallets: %w", err)
	}
	return result, nil
}

// GenerateShipmentContentUpdatePreviews
// Method: POST | Path: /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/contentUpdatePreviews
func (c *Client) GenerateShipmentContentUpdatePreviews(ctx context.Context, inboundPlanId string, shipmentId string, body interface{}) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/contentUpdatePreviews"
	path = strings.Replace(path, "{inboundPlanId}", inboundPlanId, 1)
	path = strings.Replace(path, "{shipmentId}", shipmentId, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("GenerateShipmentContentUpdatePreviews: %w", err)
	}
	return result, nil
}

// ListPlacementOptions
// Method: GET | Path: /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/placementOptions
func (c *Client) ListPlacementOptions(ctx context.Context, inboundPlanId string, query map[string]string) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/placementOptions"
	path = strings.Replace(path, "{inboundPlanId}", inboundPlanId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("ListPlacementOptions: %w", err)
	}
	return result, nil
}

// ConfirmPackingOption
// Method: POST | Path: /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/packingOptions/{packingOptionId}/confirmation
func (c *Client) ConfirmPackingOption(ctx context.Context, inboundPlanId string, packingOptionId string, body interface{}) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/packingOptions/{packingOptionId}/confirmation"
	path = strings.Replace(path, "{inboundPlanId}", inboundPlanId, 1)
	path = strings.Replace(path, "{packingOptionId}", packingOptionId, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("ConfirmPackingOption: %w", err)
	}
	return result, nil
}

// GetSelfShipAppointmentSlots
// Method: GET | Path: /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots
func (c *Client) GetSelfShipAppointmentSlots(ctx context.Context, inboundPlanId string, shipmentId string, query map[string]string) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots"
	path = strings.Replace(path, "{inboundPlanId}", inboundPlanId, 1)
	path = strings.Replace(path, "{shipmentId}", shipmentId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetSelfShipAppointmentSlots: %w", err)
	}
	return result, nil
}

// GenerateDeliveryWindowOptions
// Method: POST | Path: /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/deliveryWindowOptions
func (c *Client) GenerateDeliveryWindowOptions(ctx context.Context, inboundPlanId string, shipmentId string, body interface{}) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/deliveryWindowOptions"
	path = strings.Replace(path, "{inboundPlanId}", inboundPlanId, 1)
	path = strings.Replace(path, "{shipmentId}", shipmentId, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("GenerateDeliveryWindowOptions: %w", err)
	}
	return result, nil
}

// CreateMarketplaceItemLabels
// Method: POST | Path: /inbound/fba/2024-03-20/items/labels
func (c *Client) CreateMarketplaceItemLabels(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/items/labels"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("CreateMarketplaceItemLabels: %w", err)
	}
	return result, nil
}

// ListDeliveryWindowOptions
// Method: GET | Path: /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/deliveryWindowOptions
func (c *Client) ListDeliveryWindowOptions(ctx context.Context, inboundPlanId string, shipmentId string, query map[string]string) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/deliveryWindowOptions"
	path = strings.Replace(path, "{inboundPlanId}", inboundPlanId, 1)
	path = strings.Replace(path, "{shipmentId}", shipmentId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("ListDeliveryWindowOptions: %w", err)
	}
	return result, nil
}

// ListInboundPlanPallets
// Method: GET | Path: /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/pallets
func (c *Client) ListInboundPlanPallets(ctx context.Context, inboundPlanId string, query map[string]string) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/pallets"
	path = strings.Replace(path, "{inboundPlanId}", inboundPlanId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("ListInboundPlanPallets: %w", err)
	}
	return result, nil
}

// ListInboundPlanBoxes
// Method: GET | Path: /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/boxes
func (c *Client) ListInboundPlanBoxes(ctx context.Context, inboundPlanId string, query map[string]string) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/boxes"
	path = strings.Replace(path, "{inboundPlanId}", inboundPlanId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("ListInboundPlanBoxes: %w", err)
	}
	return result, nil
}

// ConfirmShipmentContentUpdatePreview
// Method: POST | Path: /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/contentUpdatePreviews/{contentUpdatePreviewId}/confirmation
func (c *Client) ConfirmShipmentContentUpdatePreview(ctx context.Context, inboundPlanId string, shipmentId string, contentUpdatePreviewId string, body interface{}) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/contentUpdatePreviews/{contentUpdatePreviewId}/confirmation"
	path = strings.Replace(path, "{inboundPlanId}", inboundPlanId, 1)
	path = strings.Replace(path, "{shipmentId}", shipmentId, 1)
	path = strings.Replace(path, "{contentUpdatePreviewId}", contentUpdatePreviewId, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("ConfirmShipmentContentUpdatePreview: %w", err)
	}
	return result, nil
}

// GetShipment
// Method: GET | Path: /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}
func (c *Client) GetShipment(ctx context.Context, inboundPlanId string, shipmentId string, query map[string]string) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}"
	path = strings.Replace(path, "{inboundPlanId}", inboundPlanId, 1)
	path = strings.Replace(path, "{shipmentId}", shipmentId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetShipment: %w", err)
	}
	return result, nil
}

// ListShipmentContentUpdatePreviews
// Method: GET | Path: /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/contentUpdatePreviews
func (c *Client) ListShipmentContentUpdatePreviews(ctx context.Context, inboundPlanId string, shipmentId string, query map[string]string) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/contentUpdatePreviews"
	path = strings.Replace(path, "{inboundPlanId}", inboundPlanId, 1)
	path = strings.Replace(path, "{shipmentId}", shipmentId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("ListShipmentContentUpdatePreviews: %w", err)
	}
	return result, nil
}

// GenerateSelfShipAppointmentSlots
// Method: POST | Path: /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots
func (c *Client) GenerateSelfShipAppointmentSlots(ctx context.Context, inboundPlanId string, shipmentId string, body interface{}) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots"
	path = strings.Replace(path, "{inboundPlanId}", inboundPlanId, 1)
	path = strings.Replace(path, "{shipmentId}", shipmentId, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("GenerateSelfShipAppointmentSlots: %w", err)
	}
	return result, nil
}

// ListPackingGroupItems
// Method: GET | Path: /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/packingGroups/{packingGroupId}/items
func (c *Client) ListPackingGroupItems(ctx context.Context, inboundPlanId string, packingGroupId string, query map[string]string) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/packingGroups/{packingGroupId}/items"
	path = strings.Replace(path, "{inboundPlanId}", inboundPlanId, 1)
	path = strings.Replace(path, "{packingGroupId}", packingGroupId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("ListPackingGroupItems: %w", err)
	}
	return result, nil
}

// ListTransportationOptions
// Method: GET | Path: /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/transportationOptions
func (c *Client) ListTransportationOptions(ctx context.Context, inboundPlanId string, query map[string]string) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/transportationOptions"
	path = strings.Replace(path, "{inboundPlanId}", inboundPlanId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("ListTransportationOptions: %w", err)
	}
	return result, nil
}

// ListShipmentBoxes
// Method: GET | Path: /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/boxes
func (c *Client) ListShipmentBoxes(ctx context.Context, inboundPlanId string, shipmentId string, query map[string]string) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/boxes"
	path = strings.Replace(path, "{inboundPlanId}", inboundPlanId, 1)
	path = strings.Replace(path, "{shipmentId}", shipmentId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("ListShipmentBoxes: %w", err)
	}
	return result, nil
}

// ListShipmentItems
// Method: GET | Path: /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/items
func (c *Client) ListShipmentItems(ctx context.Context, inboundPlanId string, shipmentId string, query map[string]string) (interface{}, error) {
	path := "/inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/items"
	path = strings.Replace(path, "{inboundPlanId}", inboundPlanId, 1)
	path = strings.Replace(path, "{shipmentId}", shipmentId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("ListShipmentItems: %w", err)
	}
	return result, nil
}
