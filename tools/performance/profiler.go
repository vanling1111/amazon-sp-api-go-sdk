// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

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

