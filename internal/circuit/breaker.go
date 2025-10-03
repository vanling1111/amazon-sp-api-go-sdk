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
// terms of the applicable license. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package circuit 提供熔断器（Circuit Breaker）功能。
//
// 熔断器用于防止级联失败和雪崩效应。当检测到大量失败时，
// 自动"熔断"（停止发送请求），给下游服务恢复时间。
package circuit

import (
	"fmt"
	"sync"
	"time"
)

// State 熔断器状态
type State int

const (
	// StateClosed 关闭状态（正常工作）
	StateClosed State = iota

	// StateOpen 打开状态（熔断中，拒绝所有请求）
	StateOpen

	// StateHalfOpen 半开状态（尝试恢复）
	StateHalfOpen
)

// String 返回状态字符串表示
func (s State) String() string {
	switch s {
	case StateClosed:
		return "closed"
	case StateOpen:
		return "open"
	case StateHalfOpen:
		return "half-open"
	default:
		return "unknown"
	}
}

// Breaker 熔断器
type Breaker struct {
	// maxFailures 触发熔断的最大失败次数
	maxFailures int

	// timeout 熔断超时时间（打开状态持续时间）
	timeout time.Duration

	// state 当前状态
	state State

	// failures 当前失败计数
	failures int

	// lastFailTime 最后一次失败的时间
	lastFailTime time.Time

	// mu 保护并发访问
	mu sync.RWMutex

	// onStateChange 状态变更回调
	onStateChange func(from, to State)
}

// Config 熔断器配置
type Config struct {
	// MaxFailures 触发熔断的最大连续失败次数（默认 5）
	MaxFailures int

	// Timeout 熔断超时时间（默认 60 秒）
	Timeout time.Duration

	// OnStateChange 状态变更回调（可选）
	OnStateChange func(from, to State)
}

// NewBreaker 创建新的熔断器。
//
// 参数:
//   - config: 熔断器配置
//
// 返回值:
//   - *Breaker: 熔断器实例
//
// 示例:
//
//	breaker := circuit.NewBreaker(&circuit.Config{
//	    MaxFailures: 5,
//	    Timeout:     60 * time.Second,
//	})
func NewBreaker(config *Config) *Breaker {
	if config.MaxFailures == 0 {
		config.MaxFailures = 5
	}
	if config.Timeout == 0 {
		config.Timeout = 60 * time.Second
	}

	return &Breaker{
		maxFailures:   config.MaxFailures,
		timeout:       config.Timeout,
		state:         StateClosed,
		onStateChange: config.OnStateChange,
	}
}

// Execute 执行函数，带熔断保护。
//
// 参数:
//   - fn: 要执行的函数
//
// 返回值:
//   - error: 如果熔断或执行失败，返回错误
//
// 示例:
//
//	err := breaker.Execute(func() error {
//	    return doAPICall()
//	})
func (b *Breaker) Execute(fn func() error) error {
	// 检查当前状态
	if !b.allowRequest() {
		return ErrCircuitOpen
	}

	// 执行函数
	err := fn()

	// 记录结果
	b.recordResult(err)

	return err
}

// allowRequest 检查是否允许请求
func (b *Breaker) allowRequest() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	switch b.state {
	case StateClosed:
		// 关闭状态，允许所有请求
		return true

	case StateOpen:
		// 检查是否应该尝试恢复
		if time.Since(b.lastFailTime) > b.timeout {
			// 超时时间已过，切换到半开状态
			b.setState(StateHalfOpen)
			return true
		}
		// 仍在熔断中
		return false

	case StateHalfOpen:
		// 半开状态，允许少量请求测试
		return true

	default:
		return false
	}
}

// recordResult 记录执行结果
func (b *Breaker) recordResult(err error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if err != nil {
		// 失败
		b.failures++
		b.lastFailTime = time.Now()

		switch b.state {
		case StateClosed:
			// 检查是否达到熔断阈值
			if b.failures >= b.maxFailures {
				b.setState(StateOpen)
			}

		case StateHalfOpen:
			// 半开状态失败，立即回到打开状态
			b.setState(StateOpen)
		}
	} else {
		// 成功
		switch b.state {
		case StateClosed:
			// 重置失败计数
			b.failures = 0

		case StateHalfOpen:
			// 半开状态成功，回到关闭状态
			b.failures = 0
			b.setState(StateClosed)
		}
	}
}

// setState 设置状态并触发回调
func (b *Breaker) setState(newState State) {
	oldState := b.state
	b.state = newState

	if b.onStateChange != nil {
		// 异步调用回调，避免阻塞
		go b.onStateChange(oldState, newState)
	}
}

// State 获取当前状态
func (b *Breaker) State() State {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.state
}

// Failures 获取当前失败计数
func (b *Breaker) Failures() int {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.failures
}

// Reset 重置熔断器到关闭状态
func (b *Breaker) Reset() {
	b.mu.Lock()
	defer b.mu.Unlock()

	oldState := b.state
	b.state = StateClosed
	b.failures = 0

	if b.onStateChange != nil && oldState != StateClosed {
		go b.onStateChange(oldState, StateClosed)
	}
}

// ErrCircuitOpen 表示熔断器处于打开状态
var ErrCircuitOpen = fmt.Errorf("circuit breaker is open")
