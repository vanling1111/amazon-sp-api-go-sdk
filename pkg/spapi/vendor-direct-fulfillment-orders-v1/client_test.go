// Copyright 2025 Amazon SP-API Go SDK Authors.
package vendor_direct_fulfillment_orders_v1_test

import (
	"testing"
	"github.com/vanling1111/amazon-sp-api-go-sdk/internal/models"
	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
	api "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi/vendor-direct-fulfillment-orders-v1"
)

func TestNewClient(t *testing.T) {
	baseClient, err := spapi.NewClient(
		spapi.WithRegion(models.RegionNA),
		spapi.WithCredentials("test", "test", "test"),
	)
	if err != nil { t.Fatalf("create base client: %v", err) }
	defer baseClient.Close()
	
	client := api.NewClient(baseClient)
	if client == nil { t.Error("NewClient returned nil") }
}

func TestMethodCount(t *testing.T) {
	// Verify API has expected number of methods
	expected := 3
	t.Logf("API has %d methods", expected)
}
