package models

import (
	"testing"
)

func TestGetRegionByCode(t *testing.T) {
	tests := []struct {
		name string
		code string
		want *Region
	}{
		{
			name: "NA region",
			code: "na",
			want: &RegionNA,
		},
		{
			name: "EU region",
			code: "eu",
			want: &RegionEU,
		},
		{
			name: "FE region",
			code: "fe",
			want: &RegionFE,
		},
		{
			name: "invalid region",
			code: "invalid",
			want: nil,
		},
		{
			name: "empty code",
			code: "",
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetRegionByCode(tt.code)
			if got == nil && tt.want == nil {
				return
			}
			if got == nil || tt.want == nil {
				t.Errorf("GetRegionByCode() = %v, want %v", got, tt.want)
				return
			}
			if got.Code != tt.want.Code {
				t.Errorf("GetRegionByCode() Code = %v, want %v", got.Code, tt.want.Code)
			}
		})
	}
}

func TestGetMarketplaceByID(t *testing.T) {
	tests := []struct {
		name string
		id   string
		want *Marketplace
	}{
		{
			name: "US marketplace",
			id:   "ATVPDKIKX0DER",
			want: &MarketplaceUS,
		},
		{
			name: "CA marketplace",
			id:   "A2EUQ1WTGCTBG2",
			want: &MarketplaceCA,
		},
		{
			name: "UK marketplace",
			id:   "A1F83G8C2ARO7P",
			want: &MarketplaceUK,
		},
		{
			name: "JP marketplace",
			id:   "A1VC38T7YXB528",
			want: &MarketplaceJP,
		},
		{
			name: "invalid marketplace",
			id:   "INVALID",
			want: nil,
		},
		{
			name: "empty ID",
			id:   "",
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetMarketplaceByID(tt.id)
			if got == nil && tt.want == nil {
				return
			}
			if got == nil || tt.want == nil {
				t.Errorf("GetMarketplaceByID() = %v, want %v", got, tt.want)
				return
			}
			if got.ID != tt.want.ID {
				t.Errorf("GetMarketplaceByID() ID = %v, want %v", got.ID, tt.want.ID)
			}
		})
	}
}

func TestRegion_Fields(t *testing.T) {
	// 验证 Region 结构的字段
	if RegionNA.Code != "na" {
		t.Errorf("RegionNA.Code = %v, want na", RegionNA.Code)
	}
	if RegionNA.Name != "North America" {
		t.Errorf("RegionNA.Name = %v, want North America", RegionNA.Name)
	}
	if RegionNA.Endpoint != "https://sellingpartnerapi-na.amazon.com" {
		t.Errorf("RegionNA.Endpoint = %v, want https://sellingpartnerapi-na.amazon.com", RegionNA.Endpoint)
	}
	if RegionNA.LWAEndpoint != "https://api.amazon.com/auth/o2/token" {
		t.Errorf("RegionNA.LWAEndpoint = %v, want https://api.amazon.com/auth/o2/token", RegionNA.LWAEndpoint)
	}
}

func TestMarketplace_Fields(t *testing.T) {
	// 验证 Marketplace 结构的字段
	if MarketplaceUS.ID != "ATVPDKIKX0DER" {
		t.Errorf("MarketplaceUS.ID = %v, want ATVPDKIKX0DER", MarketplaceUS.ID)
	}
	if MarketplaceUS.Name != "United States" {
		t.Errorf("MarketplaceUS.Name = %v, want United States", MarketplaceUS.Name)
	}
	if MarketplaceUS.CountryCode != "US" {
		t.Errorf("MarketplaceUS.CountryCode = %v, want US", MarketplaceUS.CountryCode)
	}
	if MarketplaceUS.Endpoint != "https://sellingpartnerapi-na.amazon.com" {
		t.Errorf("MarketplaceUS.Endpoint = %v, want https://sellingpartnerapi-na.amazon.com", MarketplaceUS.Endpoint)
	}
}

func TestAllRegions(t *testing.T) {
	regions := []struct {
		name   string
		region Region
	}{
		{"NA", RegionNA},
		{"EU", RegionEU},
		{"FE", RegionFE},
	}

	for _, tt := range regions {
		t.Run(tt.name, func(t *testing.T) {
			if tt.region.Code == "" {
				t.Error("Region Code should not be empty")
			}
			if tt.region.Name == "" {
				t.Error("Region Name should not be empty")
			}
			if tt.region.Endpoint == "" {
				t.Error("Region Endpoint should not be empty")
			}
			if tt.region.LWAEndpoint == "" {
				t.Error("Region LWAEndpoint should not be empty")
			}
		})
	}
}

func TestAllMarketplaces(t *testing.T) {
	marketplaces := []struct {
		name        string
		marketplace Marketplace
	}{
		{"US", MarketplaceUS},
		{"CA", MarketplaceCA},
		{"MX", MarketplaceMX},
		{"UK", MarketplaceUK},
		{"DE", MarketplaceDE},
		{"FR", MarketplaceFR},
		{"JP", MarketplaceJP},
	}

	for _, tt := range marketplaces {
		t.Run(tt.name, func(t *testing.T) {
			if tt.marketplace.ID == "" {
				t.Error("Marketplace ID should not be empty")
			}
			if tt.marketplace.Name == "" {
				t.Error("Marketplace Name should not be empty")
			}
			if tt.marketplace.CountryCode == "" {
				t.Error("Marketplace CountryCode should not be empty")
			}
			if tt.marketplace.Endpoint == "" {
				t.Error("Marketplace Endpoint should not be empty")
			}
		})
	}
}
