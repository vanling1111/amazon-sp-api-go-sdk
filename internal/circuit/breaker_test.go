// Copyright 2025 Amazon SP-API Go SDK Authors. All rights reserved.
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

package circuit

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestBreaker_Success tests successful executions
func TestBreaker_Success(t *testing.T) {
	breaker := NewBreaker(&Config{
		MaxFailures: 3,
		Timeout:     time.Second,
	})

	// 执行成功的请求
	for range 10 {
		err := breaker.Execute(func() error {
			return nil
		})
		assert.NoError(t, err)
	}

	// 状态应该保持 Closed
	assert.Equal(t, StateClosed, breaker.State())
	assert.Equal(t, 0, breaker.Failures())
}

// TestBreaker_Failures tests failure handling
func TestBreaker_Failures(t *testing.T) {
	breaker := NewBreaker(&Config{
		MaxFailures: 3,
		Timeout:     100 * time.Millisecond,
	})

	testErr := errors.New("test error")

	// 执行 3 次失败请求
	for i := range 3 {
		err := breaker.Execute(func() error {
			return testErr
		})
		assert.Error(t, err)
		assert.Equal(t, i+1, breaker.Failures())
	}

	// 达到阈值，应该切换到 Open 状态
	assert.Equal(t, StateOpen, breaker.State())

	// 下一个请求应该被拒绝
	err := breaker.Execute(func() error {
		return nil
	})
	assert.Equal(t, ErrCircuitOpen, err)
}

// TestBreaker_Recovery tests recovery from open state
func TestBreaker_Recovery(t *testing.T) {
	breaker := NewBreaker(&Config{
		MaxFailures: 2,
		Timeout:     100 * time.Millisecond,
	})

	// 触发熔断
	for range 2 {
		breaker.Execute(func() error {
			return errors.New("fail")
		})
	}

	assert.Equal(t, StateOpen, breaker.State())

	// 等待超时
	time.Sleep(150 * time.Millisecond)

	// 下一个请求应该被允许（Half-Open）
	err := breaker.Execute(func() error {
		return nil  // 成功
	})
	assert.NoError(t, err)

	// 成功后应该回到 Closed 状态
	assert.Equal(t, StateClosed, breaker.State())
	assert.Equal(t, 0, breaker.Failures())
}

// TestBreaker_HalfOpenFailure tests half-open state failure
func TestBreaker_HalfOpenFailure(t *testing.T) {
	breaker := NewBreaker(&Config{
		MaxFailures: 2,
		Timeout:     50 * time.Millisecond,
	})

	// 触发熔断
	for range 2 {
		breaker.Execute(func() error {
			return errors.New("fail")
		})
	}

	assert.Equal(t, StateOpen, breaker.State())

	// 等待超时
	time.Sleep(60 * time.Millisecond)

	// 半开状态的请求失败
	err := breaker.Execute(func() error {
		return errors.New("still failing")
	})
	assert.Error(t, err)

	// 应该回到 Open 状态
	assert.Equal(t, StateOpen, breaker.State())
}

// TestBreaker_Reset tests manual reset
func TestBreaker_Reset(t *testing.T) {
	breaker := NewBreaker(&Config{
		MaxFailures: 2,
		Timeout:     time.Second,
	})

	// 触发熔断
	for range 2 {
		breaker.Execute(func() error {
			return errors.New("fail")
		})
	}

	assert.Equal(t, StateOpen, breaker.State())

	// 手动重置
	breaker.Reset()

	assert.Equal(t, StateClosed, breaker.State())
	assert.Equal(t, 0, breaker.Failures())

	// 可以正常执行
	err := breaker.Execute(func() error {
		return nil
	})
	assert.NoError(t, err)
}

// TestBreaker_StateChange tests state change callback
func TestBreaker_StateChange(t *testing.T) {
	stateChanges := []State{}

	breaker := NewBreaker(&Config{
		MaxFailures: 2,
		Timeout:     50 * time.Millisecond,
		OnStateChange: func(from, to State) {
			stateChanges = append(stateChanges, to)
		},
	})

	// 触发熔断
	for range 2 {
		breaker.Execute(func() error {
			return errors.New("fail")
		})
	}

	// 等待回调执行
	time.Sleep(10 * time.Millisecond)

	// 应该收到状态变更通知
	assert.Contains(t, stateChanges, StateOpen)
}

// TestBreaker_Concurrent tests concurrent usage
func TestBreaker_Concurrent(t *testing.T) {
	breaker := NewBreaker(&Config{
		MaxFailures: 10,
		Timeout:     time.Second,
	})

	// 并发执行
	done := make(chan bool, 20)

	for range 20 {
		go func() {
			defer func() { done <- true }()
			
			breaker.Execute(func() error {
				time.Sleep(time.Millisecond)
				return nil
			})
		}()
	}

	// 等待所有 goroutine 完成
	for range 20 {
		<-done
	}

	// 应该没有问题
	assert.Equal(t, StateClosed, breaker.State())
}

// BenchmarkBreaker benchmarks breaker performance
func BenchmarkBreaker(b *testing.B) {
	breaker := NewBreaker(&Config{
		MaxFailures: 1000,
		Timeout:     time.Minute,
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		breaker.Execute(func() error {
			return nil
		})
	}
}

