package spapi_test

import (
	"testing"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

// TestWithSandbox 测试Sandbox模式
func TestWithSandbox(t *testing.T) {
	client, err := spapi.NewClient(
		spapi.WithRegion(spapi.RegionNA),
		spapi.WithSandbox(), // 应该自动切换到 RegionNASandbox
		spapi.WithCredentials("test-id", "test-secret", "test-token"),
	)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	defer client.Close()

	config := client.Config()
	if !config.Region.IsSandbox() {
		t.Errorf("Expected sandbox region, got %v", config.Region)
	}
	if config.Region.Code != "na-sandbox" {
		t.Errorf("Expected na-sandbox, got %v", config.Region.Code)
	}
}

// TestRegionSandboxMethods 测试Region的Sandbox方法
func TestRegionSandboxMethods(t *testing.T) {
	tests := []struct {
		name       string
		region     spapi.Region
		isSandbox  bool
		toSandbox  spapi.Region
		toProduction spapi.Region
	}{
		{
			name:       "NA production",
			region:     spapi.RegionNA,
			isSandbox:  false,
			toSandbox:  spapi.RegionNASandbox,
			toProduction: spapi.RegionNA,
		},
		{
			name:       "NA sandbox",
			region:     spapi.RegionNASandbox,
			isSandbox:  true,
			toSandbox:  spapi.RegionNASandbox,
			toProduction: spapi.RegionNA,
		},
		{
			name:       "EU production",
			region:     spapi.RegionEU,
			isSandbox:  false,
			toSandbox:  spapi.RegionEUSandbox,
			toProduction: spapi.RegionEU,
		},
		{
			name:       "FE production",
			region:     spapi.RegionFE,
			isSandbox:  false,
			toSandbox:  spapi.RegionFESandbox,
			toProduction: spapi.RegionFE,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.region.IsSandbox(); got != tt.isSandbox {
				t.Errorf("IsSandbox() = %v, want %v", got, tt.isSandbox)
			}
			if got := tt.region.ToSandbox(); got != tt.toSandbox {
				t.Errorf("ToSandbox() = %v, want %v", got, tt.toSandbox)
			}
			if got := tt.region.ToProduction(); got != tt.toProduction {
				t.Errorf("ToProduction() = %v, want %v", got, tt.toProduction)
			}
		})
	}
}

// ExampleWithSandbox 演示如何使用Sandbox模式
func ExampleWithSandbox() {
	client, err := spapi.NewClient(
		spapi.WithRegion(spapi.RegionNA),
		spapi.WithSandbox(), // 自动切换到测试环境
		spapi.WithCredentials("your-client-id", "your-client-secret", "your-refresh-token"),
	)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// 使用客户端进行测试...
}
