// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

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
