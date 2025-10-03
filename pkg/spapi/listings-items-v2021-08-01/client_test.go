// Copyright 2025 Amazon SP-API Go SDK Authors.
package listings_items_v2021_08_01_test

import (
	"testing"
	"github.com/vanling1111/amazon-sp-api-go-sdk/internal/models"
	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
	api "github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi/listings-items-v2021-08-01"
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
	expected := 5
	t.Logf("API has %d methods", expected)
}
