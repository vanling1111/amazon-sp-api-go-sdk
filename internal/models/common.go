// Package models 提供内部通用模型和数据结构。
//
// 此包包含在多个内部模块之间共享的通用数据结构，
// 不对外暴露。公开的 API 模型定义在 api/ 目录下。
package models

import (
	"time"
)

// Marketplace 表示 Amazon 市场。
type Marketplace struct {
	// ID 是市场的唯一标识符
	ID string

	// Name 是市场的名称
	Name string

	// CountryCode 是市场所在国家的 ISO 代码
	CountryCode string

	// Endpoint 是市场的 API 端点
	Endpoint string
}

// Region 表示 Amazon SP-API 区域。
type Region struct {
	// Code 是区域代码 (na, eu, fe)
	Code string

	// Name 是区域名称
	Name string

	// Endpoint 是区域的 API 端点
	Endpoint string

	// LWAEndpoint 是 LWA 令牌端点
	LWAEndpoint string
}

// RateLimitInfo 表示速率限制信息。
type RateLimitInfo struct {
	// Rate 是允许的请求速率（请求数/时间单位）
	Rate float64

	// Burst 是允许的突发请求数
	Burst int

	// RefillInterval 是令牌桶的补充间隔
	RefillInterval time.Duration
}

// RequestMetadata 表示请求的元数据。
type RequestMetadata struct {
	// RequestID 是请求的唯一标识符
	RequestID string

	// Timestamp 是请求时间戳
	Timestamp time.Time

	// Marketplace 是请求的市场
	Marketplace string

	// Endpoint 是请求的端点
	Endpoint string

	// Method 是 HTTP 方法
	Method string
}

// ErrorDetail 表示错误详情。
type ErrorDetail struct {
	// Code 是错误代码
	Code string

	// Message 是错误消息
	Message string

	// Details 是额外的错误详情
	Details map[string]interface{}
}

// 预定义的区域
var (
	// RegionNA 是北美区域
	RegionNA = Region{
		Code:        "na",
		Name:        "North America",
		Endpoint:    "https://sellingpartnerapi-na.amazon.com",
		LWAEndpoint: "https://api.amazon.com/auth/o2/token",
	}

	// RegionEU 是欧洲区域
	RegionEU = Region{
		Code:        "eu",
		Name:        "Europe",
		Endpoint:    "https://sellingpartnerapi-eu.amazon.com",
		LWAEndpoint: "https://api.amazon.co.uk/auth/token",
	}

	// RegionFE 是远东区域
	RegionFE = Region{
		Code:        "fe",
		Name:        "Far East",
		Endpoint:    "https://sellingpartnerapi-fe.amazon.com",
		LWAEndpoint: "https://api.amazon.co.jp/auth/token",
	}
)

// 预定义的市场
var (
	// MarketplaceUS 是美国市场
	MarketplaceUS = Marketplace{
		ID:          "ATVPDKIKX0DER",
		Name:        "United States",
		CountryCode: "US",
		Endpoint:    "https://sellingpartnerapi-na.amazon.com",
	}

	// MarketplaceCA 是加拿大市场
	MarketplaceCA = Marketplace{
		ID:          "A2EUQ1WTGCTBG2",
		Name:        "Canada",
		CountryCode: "CA",
		Endpoint:    "https://sellingpartnerapi-na.amazon.com",
	}

	// MarketplaceMX 是墨西哥市场
	MarketplaceMX = Marketplace{
		ID:          "A1AM78C64UM0Y8",
		Name:        "Mexico",
		CountryCode: "MX",
		Endpoint:    "https://sellingpartnerapi-na.amazon.com",
	}

	// MarketplaceUK 是英国市场
	MarketplaceUK = Marketplace{
		ID:          "A1F83G8C2ARO7P",
		Name:        "United Kingdom",
		CountryCode: "GB",
		Endpoint:    "https://sellingpartnerapi-eu.amazon.com",
	}

	// MarketplaceDE 是德国市场
	MarketplaceDE = Marketplace{
		ID:          "A1PA6795UKMFR9",
		Name:        "Germany",
		CountryCode: "DE",
		Endpoint:    "https://sellingpartnerapi-eu.amazon.com",
	}

	// MarketplaceFR 是法国市场
	MarketplaceFR = Marketplace{
		ID:          "A13V1IB3VIYZZH",
		Name:        "France",
		CountryCode: "FR",
		Endpoint:    "https://sellingpartnerapi-eu.amazon.com",
	}

	// MarketplaceJP 是日本市场
	MarketplaceJP = Marketplace{
		ID:          "A1VC38T7YXB528",
		Name:        "Japan",
		CountryCode: "JP",
		Endpoint:    "https://sellingpartnerapi-fe.amazon.com",
	}
)

// GetRegionByCode 根据区域代码获取区域信息。
//
// 参数:
//   - code: 区域代码 (na, eu, fe)
//
// 返回值:
//   - *Region: 区域信息，如果未找到返回 nil
func GetRegionByCode(code string) *Region {
	switch code {
	case "na":
		return &RegionNA
	case "eu":
		return &RegionEU
	case "fe":
		return &RegionFE
	default:
		return nil
	}
}

// GetMarketplaceByID 根据市场 ID 获取市场信息。
//
// 参数:
//   - id: 市场 ID
//
// 返回值:
//   - *Marketplace: 市场信息，如果未找到返回 nil
func GetMarketplaceByID(id string) *Marketplace {
	marketplaces := map[string]Marketplace{
		"ATVPDKIKX0DER":  MarketplaceUS,
		"A2EUQ1WTGCTBG2": MarketplaceCA,
		"A1AM78C64UM0Y8": MarketplaceMX,
		"A1F83G8C2ARO7P": MarketplaceUK,
		"A1PA6795UKMFR9": MarketplaceDE,
		"A13V1IB3VIYZZH": MarketplaceFR,
		"A1VC38T7YXB528": MarketplaceJP,
	}

	if mp, ok := marketplaces[id]; ok {
		return &mp
	}
	return nil
}
