// Copyright 2025 Amazon SP-API Go SDK Authors.
//
// This file is part of Amazon SP-API Go SDK.
//
// Amazon SP-API Go SDK is dual-licensed:
//
// 1. GNU Affero General Public License v3.0 (AGPL-3.0) for open source use
//   - Free for personal, educational, and open source projects
//   - Your project must also be open sourced under AGPL-3.0
//   - See: https://www.gnu.org/licenses/agpl-3.0.html
//
// 2. Commercial License for proprietary/commercial use
//   - Required for any commercial, enterprise, or proprietary use
//   - Allows closed source distribution
//   - Contact: vanling1111@gmail.com
//
// Unless you have obtained a commercial license, this file is licensed
// under AGPL-3.0. By using this software, you agree to comply with the
// terms of the applicable license. All rights reserved.

package spapi

import "github.com/vanling1111/amazon-sp-api-go-sdk/internal/models"

// Region 表示 Amazon SP-API 的区域。
//
// Amazon SP-API 在全球有三个主要区域，每个区域有不同的端点和市场。
//
// 支持的区域：
//   - RegionNA: 北美区域（美国、加拿大、墨西哥、巴西）
//   - RegionEU: 欧洲区域（英国、德国、法国、意大利、西班牙等）
//   - RegionFE: 远东区域（日本、澳大利亚、新加坡、印度等）
//
// 示例:
//
//	client, err := spapi.NewClient(
//	    spapi.WithRegion(spapi.RegionNA),
//	    spapi.WithCredentials("client-id", "client-secret", "refresh-token"),
//	)
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

// 预定义的区域
var (
	// RegionNA 北美区域
	//
	// 端点: https://sellingpartnerapi-na.amazon.com
	//
	// 支持的市场:
	//   - 美国 (US)
	//   - 加拿大 (CA)
	//   - 墨西哥 (MX)
	//   - 巴西 (BR)
	RegionNA = Region{
		Code:        "na",
		Name:        "North America",
		Endpoint:    "https://sellingpartnerapi-na.amazon.com",
		LWAEndpoint: "https://api.amazon.com/auth/o2/token",
	}

	// RegionEU 欧洲区域
	//
	// 端点: https://sellingpartnerapi-eu.amazon.com
	//
	// 支持的市场:
	//   - 英国 (UK)
	//   - 德国 (DE)
	//   - 法国 (FR)
	//   - 意大利 (IT)
	//   - 西班牙 (ES)
	//   - 荷兰 (NL)
	//   - 波兰 (PL)
	//   - 瑞典 (SE)
	//   - 土耳其 (TR)
	//   - 阿联酋 (AE)
	//   - 印度 (IN)
	RegionEU = Region{
		Code:        "eu",
		Name:        "Europe",
		Endpoint:    "https://sellingpartnerapi-eu.amazon.com",
		LWAEndpoint: "https://api.amazon.com/auth/o2/token",
	}

	// RegionFE 远东区域
	//
	// 端点: https://sellingpartnerapi-fe.amazon.com
	//
	// 支持的市场:
	//   - 日本 (JP)
	//   - 澳大利亚 (AU)
	//   - 新加坡 (SG)
	RegionFE = Region{
		Code:        "fe",
		Name:        "Far East",
		Endpoint:    "https://sellingpartnerapi-fe.amazon.com",
		LWAEndpoint: "https://api.amazon.com/auth/o2/token",
	}
)

// String 返回区域的字符串表示。
func (r Region) String() string {
	return r.Code
}

// IsValid 检查区域是否有效。
func (r Region) IsValid() bool {
	return r.Code != "" && r.Endpoint != "" && r.LWAEndpoint != ""
}

// toInternal 将公开的 Region 转换为内部的 models.Region。
//
// 这是一个内部方法，用于在 SDK 内部进行类型转换。
func (r Region) toInternal() models.Region {
	return models.Region{
		Code:        r.Code,
		Name:        r.Name,
		Endpoint:    r.Endpoint,
		LWAEndpoint: r.LWAEndpoint,
	}
}

// regionFromInternal 将内部的 models.Region 转换为公开的 Region。
//
// 这是一个内部方法，用于在 SDK 内部进行类型转换。
func regionFromInternal(r models.Region) Region {
	return Region{
		Code:        r.Code,
		Name:        r.Name,
		Endpoint:    r.Endpoint,
		LWAEndpoint: r.LWAEndpoint,
	}
}
