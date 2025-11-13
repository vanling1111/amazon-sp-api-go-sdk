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

	// RegionNASandbox 北美Sandbox区域（测试环境）
	//
	// 端点: https://sandbox.sellingpartnerapi-na.amazon.com
	//
	// 用于测试和开发，不会影响生产数据
	RegionNASandbox = Region{
		Code:        "na-sandbox",
		Name:        "North America Sandbox",
		Endpoint:    "https://sandbox.sellingpartnerapi-na.amazon.com",
		LWAEndpoint: "https://api.amazon.com/auth/o2/token",
	}

	// RegionEUSandbox 欧洲Sandbox区域（测试环境）
	//
	// 端点: https://sandbox.sellingpartnerapi-eu.amazon.com
	//
	// 用于测试和开发，不会影响生产数据
	RegionEUSandbox = Region{
		Code:        "eu-sandbox",
		Name:        "Europe Sandbox",
		Endpoint:    "https://sandbox.sellingpartnerapi-eu.amazon.com",
		LWAEndpoint: "https://api.amazon.com/auth/o2/token",
	}

	// RegionFESandbox 远东Sandbox区域（测试环境）
	//
	// 端点: https://sandbox.sellingpartnerapi-fe.amazon.com
	//
	// 用于测试和开发，不会影响生产数据
	RegionFESandbox = Region{
		Code:        "fe-sandbox",
		Name:        "Far East Sandbox",
		Endpoint:    "https://sandbox.sellingpartnerapi-fe.amazon.com",
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

// IsSandbox 检查是否为Sandbox区域。
func (r Region) IsSandbox() bool {
	return r.Code == "na-sandbox" || r.Code == "eu-sandbox" || r.Code == "fe-sandbox"
}

// ToSandbox 将生产区域转换为对应的Sandbox区域。
//
// 如果已经是Sandbox区域，返回自身。
func (r Region) ToSandbox() Region {
	switch r.Code {
	case "na":
		return RegionNASandbox
	case "eu":
		return RegionEUSandbox
	case "fe":
		return RegionFESandbox
	default:
		return r // 已经是sandbox或未知区域
	}
}

// ToProduction 将Sandbox区域转换为对应的生产区域。
//
// 如果已经是生产区域，返回自身。
func (r Region) ToProduction() Region {
	switch r.Code {
	case "na-sandbox":
		return RegionNA
	case "eu-sandbox":
		return RegionEU
	case "fe-sandbox":
		return RegionFE
	default:
		return r // 已经是生产区域或未知区域
	}
}
