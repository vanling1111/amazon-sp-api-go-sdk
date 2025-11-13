// Copyright 2025 Amazon SP-API Go SDK Authors.
//
// This file is part of Amazon SP-API Go SDK.
//
// Amazon SP-API Go SDK is dual-licensed:
//
// 1. GNU Affero General Public License v3.0 (AGPL-3.0) for open source use
//    - Free for personal, educational, and open source projects
//    - Your project must also be open sourced under AGPL-3.0
//    - See: https://www.gnu.org/licenses/agpl-3.0.html
//
// 2. Commercial License for proprietary/commercial use
//    - Required for any commercial, enterprise, or proprietary use
//    - Allows closed source distribution
//    - Contact: vanling1111@gmail.com
//
// Unless you have obtained a commercial license, this file is licensed
// under AGPL-3.0. By using this software, you agree to comply with the
// terms of the applicable license.
//
// This file is part of Amazon SP-API Go SDK.
//
// Amazon SP-API Go SDK is dual-licensed:
//
// 1. GNU Affero General Public License v3.0 (AGPL-3.0) for open source use
//    - Free for personal, educational, and open source projects
//    - Your project must also be open sourced under AGPL-3.0
//    - See: https://www.gnu.org/licenses/agpl-3.0.html
//
// 2. Commercial License for proprietary/commercial use
//    - Required for any commercial, enterprise, or proprietary use
//    - Allows closed source distribution
//    - Contact: vanling1111@gmail.com
//
// Unless you have obtained a commercial license, this file is licensed
// under AGPL-3.0. By using this software, you agree to comply with the
// terms of the applicable license.

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi"
)

func main() {
	// 创建 SP-API 客户端
	// 需要从 Amazon Seller Central 获取以下凭证：
	// - Client ID
	// - Client Secret
	// - Refresh Token
	client, err := spapi.NewClient(
		spapi.WithRegion(spapi.RegionNA),
		spapi.WithCredentials(
			"your-client-id",
			"your-client-secret",
			"your-refresh-token",
		),
		spapi.WithHTTPTimeout(30*time.Second),
		spapi.WithMaxRetries(3),
	)
	if err != nil {
		log.Fatalf("创建客户端失败: %v", err)
	}
	defer client.Close()

	// 获取访问令牌（SDK 会自动管理）
	ctx := context.Background()
	token, err := client.GetAccessToken(ctx)
	if err != nil {
		log.Fatalf("获取访问令牌失败: %v", err)
	}

	fmt.Printf("✓ 成功获取访问令牌: %s...\n", token[:20])
	fmt.Println("✓ SDK 初始化完成，可以开始调用 API")
	fmt.Println()
	fmt.Println("提示: 请查看 examples/ 目录下的其他示例：")
	fmt.Println("  - examples/orders/      订单 API 示例")
	fmt.Println("  - examples/feeds/       Feeds API 示例")
	fmt.Println("  - examples/reports/     Reports API 示例")
	fmt.Println("  - examples/listings/    Listings API 示例")
	fmt.Println("  - examples/grantless/   Grantless 操作示例")
}
