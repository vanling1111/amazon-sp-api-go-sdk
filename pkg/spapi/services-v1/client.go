// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package services_v1

import (
	"context"
	"fmt"
	"strings"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// Client services API v1
type Client struct {
	baseClient *spapi.Client
}

// NewClient creates API client
func NewClient(baseClient *spapi.Client) *Client {
	return &Client{baseClient: baseClient}
}

// CancelReservation
// Method: DELETE | Path: /service/v1/reservation/{reservationId}
func (c *Client) CancelReservation(ctx context.Context, reservationId string) (interface{}, error) {
	path := "/service/v1/reservation/{reservationId}"
	path = strings.Replace(path, "{reservationId}", reservationId, 1)
	var result interface{}
	err := c.baseClient.Delete(ctx, path, &result)
	if err != nil {
		return nil, fmt.Errorf("CancelReservation: %w", err)
	}
	return result, nil
}

// AssignAppointmentResources
// Method: PUT | Path: /service/v1/serviceJobs/{serviceJobId}/appointments/{appointmentId}/resources
func (c *Client) AssignAppointmentResources(ctx context.Context, serviceJobId string, appointmentId string, body interface{}) (interface{}, error) {
	path := "/service/v1/serviceJobs/{serviceJobId}/appointments/{appointmentId}/resources"
	path = strings.Replace(path, "{serviceJobId}", serviceJobId, 1)
	path = strings.Replace(path, "{appointmentId}", appointmentId, 1)
	var result interface{}
	err := c.baseClient.Put(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("AssignAppointmentResources: %w", err)
	}
	return result, nil
}

// CompleteServiceJobByServiceJobId
// Method: PUT | Path: /service/v1/serviceJobs/{serviceJobId}/completions
func (c *Client) CompleteServiceJobByServiceJobId(ctx context.Context, serviceJobId string, body interface{}) (interface{}, error) {
	path := "/service/v1/serviceJobs/{serviceJobId}/completions"
	path = strings.Replace(path, "{serviceJobId}", serviceJobId, 1)
	var result interface{}
	err := c.baseClient.Put(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("CompleteServiceJobByServiceJobId: %w", err)
	}
	return result, nil
}

// GetRangeSlotCapacity
// Method: POST | Path: /service/v1/serviceResources/{resourceId}/capacity/range
func (c *Client) GetRangeSlotCapacity(ctx context.Context, resourceId string, body interface{}) (interface{}, error) {
	path := "/service/v1/serviceResources/{resourceId}/capacity/range"
	path = strings.Replace(path, "{resourceId}", resourceId, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("GetRangeSlotCapacity: %w", err)
	}
	return result, nil
}

// CreateReservation
// Method: POST | Path: /service/v1/reservation
func (c *Client) CreateReservation(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/service/v1/reservation"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("CreateReservation: %w", err)
	}
	return result, nil
}

// GetServiceJobByServiceJobId
// Method: GET | Path: /service/v1/serviceJobs/{serviceJobId}
func (c *Client) GetServiceJobByServiceJobId(ctx context.Context, serviceJobId string, query map[string]string) (interface{}, error) {
	path := "/service/v1/serviceJobs/{serviceJobId}"
	path = strings.Replace(path, "{serviceJobId}", serviceJobId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetServiceJobByServiceJobId: %w", err)
	}
	return result, nil
}

// GetAppointmmentSlotsByJobId
// Method: GET | Path: /service/v1/serviceJobs/{serviceJobId}/appointmentSlots
func (c *Client) GetAppointmmentSlotsByJobId(ctx context.Context, serviceJobId string, query map[string]string) (interface{}, error) {
	path := "/service/v1/serviceJobs/{serviceJobId}/appointmentSlots"
	path = strings.Replace(path, "{serviceJobId}", serviceJobId, 1)
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetAppointmmentSlotsByJobId: %w", err)
	}
	return result, nil
}

// RescheduleAppointmentForServiceJobByServiceJobId
// Method: POST | Path: /service/v1/serviceJobs/{serviceJobId}/appointments/{appointmentId}
func (c *Client) RescheduleAppointmentForServiceJobByServiceJobId(ctx context.Context, serviceJobId string, appointmentId string, body interface{}) (interface{}, error) {
	path := "/service/v1/serviceJobs/{serviceJobId}/appointments/{appointmentId}"
	path = strings.Replace(path, "{serviceJobId}", serviceJobId, 1)
	path = strings.Replace(path, "{appointmentId}", appointmentId, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("RescheduleAppointmentForServiceJobByServiceJobId: %w", err)
	}
	return result, nil
}

// SetAppointmentFulfillmentData
// Method: PUT | Path: /service/v1/serviceJobs/{serviceJobId}/appointments/{appointmentId}/fulfillment
func (c *Client) SetAppointmentFulfillmentData(ctx context.Context, serviceJobId string, appointmentId string, body interface{}) (interface{}, error) {
	path := "/service/v1/serviceJobs/{serviceJobId}/appointments/{appointmentId}/fulfillment"
	path = strings.Replace(path, "{serviceJobId}", serviceJobId, 1)
	path = strings.Replace(path, "{appointmentId}", appointmentId, 1)
	var result interface{}
	err := c.baseClient.Put(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("SetAppointmentFulfillmentData: %w", err)
	}
	return result, nil
}

// CreateServiceDocumentUploadDestination
// Method: POST | Path: /service/v1/documents
func (c *Client) CreateServiceDocumentUploadDestination(ctx context.Context, body interface{}) (interface{}, error) {
	path := "/service/v1/documents"
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("CreateServiceDocumentUploadDestination: %w", err)
	}
	return result, nil
}

// GetServiceJobs
// Method: GET | Path: /service/v1/serviceJobs
func (c *Client) GetServiceJobs(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/service/v1/serviceJobs"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetServiceJobs: %w", err)
	}
	return result, nil
}

// AddAppointmentForServiceJobByServiceJobId
// Method: POST | Path: /service/v1/serviceJobs/{serviceJobId}/appointments
func (c *Client) AddAppointmentForServiceJobByServiceJobId(ctx context.Context, serviceJobId string, body interface{}) (interface{}, error) {
	path := "/service/v1/serviceJobs/{serviceJobId}/appointments"
	path = strings.Replace(path, "{serviceJobId}", serviceJobId, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("AddAppointmentForServiceJobByServiceJobId: %w", err)
	}
	return result, nil
}

// UpdateReservation
// Method: PUT | Path: /service/v1/reservation/{reservationId}
func (c *Client) UpdateReservation(ctx context.Context, reservationId string, body interface{}) (interface{}, error) {
	path := "/service/v1/reservation/{reservationId}"
	path = strings.Replace(path, "{reservationId}", reservationId, 1)
	var result interface{}
	err := c.baseClient.Put(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("UpdateReservation: %w", err)
	}
	return result, nil
}

// GetAppointmentSlots
// Method: GET | Path: /service/v1/appointmentSlots
func (c *Client) GetAppointmentSlots(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/service/v1/appointmentSlots"
	var result interface{}
	err := c.baseClient.Get(ctx, path, query, &result)
	if err != nil {
		return nil, fmt.Errorf("GetAppointmentSlots: %w", err)
	}
	return result, nil
}

// UpdateSchedule
// Method: PUT | Path: /service/v1/serviceResources/{resourceId}/schedules
func (c *Client) UpdateSchedule(ctx context.Context, resourceId string, body interface{}) (interface{}, error) {
	path := "/service/v1/serviceResources/{resourceId}/schedules"
	path = strings.Replace(path, "{resourceId}", resourceId, 1)
	var result interface{}
	err := c.baseClient.Put(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("UpdateSchedule: %w", err)
	}
	return result, nil
}

// CancelServiceJobByServiceJobId
// Method: PUT | Path: /service/v1/serviceJobs/{serviceJobId}/cancellations
func (c *Client) CancelServiceJobByServiceJobId(ctx context.Context, serviceJobId string, body interface{}) (interface{}, error) {
	path := "/service/v1/serviceJobs/{serviceJobId}/cancellations"
	path = strings.Replace(path, "{serviceJobId}", serviceJobId, 1)
	var result interface{}
	err := c.baseClient.Put(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("CancelServiceJobByServiceJobId: %w", err)
	}
	return result, nil
}

// GetFixedSlotCapacity
// Method: POST | Path: /service/v1/serviceResources/{resourceId}/capacity/fixed
func (c *Client) GetFixedSlotCapacity(ctx context.Context, resourceId string, body interface{}) (interface{}, error) {
	path := "/service/v1/serviceResources/{resourceId}/capacity/fixed"
	path = strings.Replace(path, "{resourceId}", resourceId, 1)
	var result interface{}
	err := c.baseClient.Post(ctx, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("GetFixedSlotCapacity: %w", err)
	}
	return result, nil
}
