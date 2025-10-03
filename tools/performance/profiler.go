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

// Package performance provides performance profiling utilities
package performance

import (
	"fmt"
	"runtime"
	"time"
)

// Profile represents a performance profile
type Profile struct {
	Name      string
	StartTime time.Time
	EndTime   time.Time
	Duration  time.Duration
	MemBefore runtime.MemStats
	MemAfter  runtime.MemStats
}

// Profiler profiles performance of operations
type Profiler struct {
	profiles []Profile
}

// NewProfiler creates a new profiler
func NewProfiler() *Profiler {
	return &Profiler{
		profiles: make([]Profile, 0),
	}
}

// Start starts profiling an operation
func (p *Profiler) Start(name string) *Profile {
	profile := &Profile{
		Name:      name,
		StartTime: time.Now(),
	}
	runtime.ReadMemStats(&profile.MemBefore)
	return profile
}

// Stop stops profiling
func (p *Profiler) Stop(profile *Profile) {
	profile.EndTime = time.Now()
	profile.Duration = profile.EndTime.Sub(profile.StartTime)
	runtime.ReadMemStats(&profile.MemAfter)
	p.profiles = append(p.profiles, *profile)
}

// Report prints profiling report
func (p *Profiler) Report() {
	fmt.Println("=== Performance Report ===")
	for _, prof := range p.profiles {
		fmt.Printf("%s: %v\n", prof.Name, prof.Duration)
		memDiff := prof.MemAfter.Alloc - prof.MemBefore.Alloc
		fmt.Printf("  Memory: %d bytes\n", memDiff)
	}
}
