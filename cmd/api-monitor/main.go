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

// Package main monitors Amazon SP-API OpenAPI specifications for changes
package main

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	baseURL = "https://raw.githubusercontent.com/amzn/selling-partner-api-models/main/models"
)

// APISpec represents an API specification
type APISpec struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	File    string `json:"file"`
	Hash    string `json:"hash"`
	Updated string `json:"updated"`
}

// Monitor monitors API specifications for changes
type Monitor struct {
	specs       map[string]*APISpec
	stateFile   string
	githubToken string
}

// NewMonitor creates a new API monitor
func NewMonitor(stateFile, githubToken string) *Monitor {
	return &Monitor{
		specs:       make(map[string]*APISpec),
		stateFile:   stateFile,
		githubToken: githubToken,
	}
}

// LoadState loads previous state from file
func (m *Monitor) LoadState() error {
	data, err := os.ReadFile(m.stateFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // First run
		}
		return err
	}

	var specs []APISpec
	if err := json.Unmarshal(data, &specs); err != nil {
		return err
	}

	for i := range specs {
		key := specs[i].Name + "-" + specs[i].Version
		m.specs[key] = &specs[i]
	}

	return nil
}

// SaveState saves current state to file
func (m *Monitor) SaveState() error {
	var specs []APISpec
	for _, spec := range m.specs {
		specs = append(specs, *spec)
	}

	data, err := json.MarshalIndent(specs, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(m.stateFile, data, 0644)
}

// FetchSpec fetches an API specification from GitHub
func (m *Monitor) FetchSpec(ctx context.Context, apiName, version, fileName string) ([]byte, error) {
	// Build URL
	modelDir := apiName
	if modelDir != "amazon-warehousing-and-distribution-model" &&
		modelDir != "easy-ship-model" &&
		!contains(modelDir, "-model") {
		modelDir += "-api-model"
	}

	url := fmt.Sprintf("%s/%s/%s", baseURL, modelDir, fileName)

	// Create request
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	if m.githubToken != "" {
		req.Header.Set("Authorization", "token "+m.githubToken)
	}

	// Send request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, resp.Status)
	}

	return io.ReadAll(resp.Body)
}

// CalculateHash calculates SHA256 hash of content
func CalculateHash(content []byte) string {
	hash := sha256.Sum256(content)
	return fmt.Sprintf("%x", hash)
}

// CheckForUpdates checks all API specifications for updates
func (m *Monitor) CheckForUpdates(ctx context.Context, apis []APIConfig) ([]Change, error) {
	var changes []Change

	for _, api := range apis {
		key := api.Name + "-" + api.Version

		// Fetch current spec
		content, err := m.FetchSpec(ctx, api.Name, api.Version, api.File)
		if err != nil {
			log.Printf("Failed to fetch %s: %v", key, err)
			continue
		}

		// Calculate hash
		currentHash := CalculateHash(content)

		// Compare with stored hash
		if stored, exists := m.specs[key]; exists {
			if stored.Hash != currentHash {
				changes = append(changes, Change{
					API:     key,
					OldHash: stored.Hash,
					NewHash: currentHash,
					File:    api.File,
				})
			}
		}

		// Update state
		m.specs[key] = &APISpec{
			Name:    api.Name,
			Version: api.Version,
			File:    api.File,
			Hash:    currentHash,
			Updated: time.Now().Format(time.RFC3339),
		}
	}

	return changes, nil
}

// Change represents a detected change
type Change struct {
	API     string
	OldHash string
	NewHash string
	File    string
}

// APIConfig represents API configuration
type APIConfig struct {
	Name    string
	Version string
	File    string
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && s[len(s)-len(substr):] == substr
}

func main() {
	ctx := context.Background()

	// Get GitHub token from environment
	githubToken := os.Getenv("GITHUB_TOKEN")

	// Create monitor
	monitor := NewMonitor("api-state.json", githubToken)

	// Load previous state
	if err := monitor.LoadState(); err != nil {
		log.Fatalf("Failed to load state: %v", err)
	}

	// Load API list from configuration file
	apiListData, err := os.ReadFile("api-list.json")
	if err != nil {
		log.Fatalf("Failed to load api-list.json: %v", err)
	}

	var apis []APIConfig
	if err := json.Unmarshal(apiListData, &apis); err != nil {
		log.Fatalf("Failed to parse api-list.json: %v", err)
	}

	fmt.Printf("Checking %d API specifications for updates...\n", len(apis))

	// Check for updates
	changes, err := monitor.CheckForUpdates(ctx, apis)
	if err != nil {
		log.Fatalf("Failed to check updates: %v", err)
	}

	// Save state
	if err := monitor.SaveState(); err != nil {
		log.Fatalf("Failed to save state: %v", err)
	}

	// Report changes
	if len(changes) > 0 {
		fmt.Printf("\n⚠️  Found %d API specification changes:\n\n", len(changes))
		for _, change := range changes {
			fmt.Printf("  - %s\n", change.API)
			fmt.Printf("    Old: %s\n", change.OldHash[:16])
			fmt.Printf("    New: %s\n", change.NewHash[:16])
		}
		os.Exit(1) // Exit with error to trigger CI notification
	} else {
		fmt.Println("✓ No changes detected. All APIs are up to date.")
	}
}
