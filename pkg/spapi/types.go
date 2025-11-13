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

// MarketplaceID 表示 Amazon 市场 ID。
//
// 每个 Amazon 市场都有一个唯一的 ID，用于标识不同国家/地区的市场。
//
// 示例:
//
//	// 使用预定义的市场 ID
//	marketplaceID := spapi.MarketplaceUS
//
//	// 或使用自定义市场 ID
//	marketplaceID := spapi.MarketplaceID("ATVPDKIKX0DER")
type MarketplaceID string

// 北美区域市场
const (
	// MarketplaceUS 美国市场
	MarketplaceUS MarketplaceID = "ATVPDKIKX0DER"

	// MarketplaceCA 加拿大市场
	MarketplaceCA MarketplaceID = "A2EUQ1WTGCTBG2"

	// MarketplaceMX 墨西哥市场
	MarketplaceMX MarketplaceID = "A1AM78C64UM0Y8"

	// MarketplaceBR 巴西市场
	MarketplaceBR MarketplaceID = "A2Q3Y263D00KWC"
)

// 欧洲区域市场
const (
	// MarketplaceUK 英国市场
	MarketplaceUK MarketplaceID = "A1F83G8C2ARO7P"

	// MarketplaceDE 德国市场
	MarketplaceDE MarketplaceID = "A1PA6795UKMFR9"

	// MarketplaceFR 法国市场
	MarketplaceFR MarketplaceID = "A13V1IB3VIYZZH"

	// MarketplaceIT 意大利市场
	MarketplaceIT MarketplaceID = "APJ6JRA9NG5V4"

	// MarketplaceES 西班牙市场
	MarketplaceES MarketplaceID = "A1RKKUPIHCS9HS"

	// MarketplaceNL 荷兰市场
	MarketplaceNL MarketplaceID = "A1805IZSGTT6HS"

	// MarketplacePL 波兰市场
	MarketplacePL MarketplaceID = "A1C3SOZRARQ6R3"

	// MarketplaceSE 瑞典市场
	MarketplaceSE MarketplaceID = "A2NODRKZP88ZB9"

	// MarketplaceTR 土耳其市场
	MarketplaceTR MarketplaceID = "A33AVAJ2PDY3EV"

	// MarketplaceAE 阿联酋市场
	MarketplaceAE MarketplaceID = "A2VIGQ35RCS4UG"

	// MarketplaceIN 印度市场
	MarketplaceIN MarketplaceID = "A21TJRUUN4KGV"
)

// 远东区域市场
const (
	// MarketplaceJP 日本市场
	MarketplaceJP MarketplaceID = "A1VC38T7YXB528"

	// MarketplaceAU 澳大利亚市场
	MarketplaceAU MarketplaceID = "A39IBJ37TRP1C6"

	// MarketplaceSG 新加坡市场
	MarketplaceSG MarketplaceID = "A19VAU5U5O7RUS"
)

// String 返回市场 ID 的字符串表示。
func (m MarketplaceID) String() string {
	return string(m)
}

// Region 返回市场所属的区域。
func (m MarketplaceID) Region() Region {
	switch m {
	case MarketplaceUS, MarketplaceCA, MarketplaceMX, MarketplaceBR:
		return RegionNA
	case MarketplaceUK, MarketplaceDE, MarketplaceFR, MarketplaceIT, MarketplaceES,
		MarketplaceNL, MarketplacePL, MarketplaceSE, MarketplaceTR, MarketplaceAE, MarketplaceIN:
		return RegionEU
	case MarketplaceJP, MarketplaceAU, MarketplaceSG:
		return RegionFE
	default:
		return Region{} // 返回空的Region结构体
	}
}
