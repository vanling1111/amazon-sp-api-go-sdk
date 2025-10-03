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

// Package main provides a CLI tool for generating API clients and models
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "models":
		fmt.Println("Generating API models...")
		fmt.Println("Run: scripts/generate-apis-versioned.ps1")
	case "clients":
		fmt.Println("Generating API clients...")
		fmt.Println("Run: scripts/generate-all-api-clients.ps1")
	case "tests":
		fmt.Println("Generating API tests...")
		fmt.Println("Run: scripts/generate-api-client-tests.ps1")
	case "all":
		fmt.Println("Generating all code...")
		fmt.Println("1. Run: scripts/generate-apis-versioned.ps1")
		fmt.Println("2. Run: scripts/generate-all-api-clients.ps1")
		fmt.Println("3. Run: scripts/generate-api-client-tests.ps1")
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Amazon SP-API Go SDK Code Generator")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  generator <command>")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  models   - Generate API models from OpenAPI specs")
	fmt.Println("  clients  - Generate API client methods")
	fmt.Println("  tests    - Generate API client tests")
	fmt.Println("  all      - Generate everything")
	fmt.Println()
	fmt.Println("Note: This tool provides guidance. Actual generation is done via PowerShell scripts.")
}
