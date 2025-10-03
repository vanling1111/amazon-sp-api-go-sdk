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

package performance

import (
	"testing"
	"time"
)

func TestNewProfiler(t *testing.T) {
	profiler := NewProfiler()
	if profiler == nil {
		t.Error("NewProfiler returned nil")
	}
	if profiler.profiles == nil {
		t.Error("Profiler profiles not initialized")
	}
}

func TestProfiler_StartStop(t *testing.T) {
	profiler := NewProfiler()

	profile := profiler.Start("test-operation")
	if profile == nil {
		t.Fatal("Start returned nil")
	}

	// Simulate some work
	time.Sleep(10 * time.Millisecond)

	profiler.Stop(profile)

	if profile.Duration == 0 {
		t.Error("Duration not recorded")
	}

	if len(profiler.profiles) != 1 {
		t.Errorf("Expected 1 profile, got %d", len(profiler.profiles))
	}
}

func TestProfiler_Report(t *testing.T) {
	profiler := NewProfiler()

	profile := profiler.Start("operation1")
	time.Sleep(5 * time.Millisecond)
	profiler.Stop(profile)

	profile = profiler.Start("operation2")
	time.Sleep(5 * time.Millisecond)
	profiler.Stop(profile)

	// Should not panic
	profiler.Report()
}
