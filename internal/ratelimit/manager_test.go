package ratelimit

import (
	"context"
	"testing"
	"time"
)

func TestNewManager(t *testing.T) {
	tests := []struct {
		name string
		opts []ManagerOption
		want func(*Manager) bool
	}{
		{
			name: "default manager",
			opts: nil,
			want: func(m *Manager) bool {
				return m.defaultRate == 1.0 && m.defaultBurst == 5
			},
		},
		{
			name: "custom default rate",
			opts: []ManagerOption{WithDefaultRate(2.0, 10)},
			want: func(m *Manager) bool {
				return m.defaultRate == 2.0 && m.defaultBurst == 10
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			manager := NewManager(tt.opts...)

			if manager == nil {
				t.Error("NewManager() returned nil")
				return
			}

			if tt.want != nil && !tt.want(manager) {
				t.Error("NewManager() configuration check failed")
			}
		})
	}
}

func TestManager_GetOrCreateLimiter(t *testing.T) {
	manager := NewManager()

	// 首次获取应该创建
	limiter1 := manager.GetOrCreateLimiter("seller1", "app1", "market1", "op1")
	if limiter1 == nil {
		t.Fatal("GetOrCreateLimiter() returned nil")
	}

	// 再次获取应该返回相同实例
	limiter2 := manager.GetOrCreateLimiter("seller1", "app1", "market1", "op1")
	if limiter1 != limiter2 {
		t.Error("GetOrCreateLimiter() should return same instance")
	}

	// 不同维度应该返回不同实例
	limiter3 := manager.GetOrCreateLimiter("seller2", "app1", "market1", "op1")
	if limiter1 == limiter3 {
		t.Error("GetOrCreateLimiter() should return different instance for different seller")
	}
}

func TestManager_Allow(t *testing.T) {
	manager := NewManager(WithDefaultRate(10, 5))

	// 前 5 次应该允许（突发）
	for i := 0; i < 5; i++ {
		if !manager.Allow("seller1", "app1", "market1", "op1") {
			t.Errorf("Allow() attempt %d should be allowed (burst)", i+1)
		}
	}

	// 第 6 次应该被限制
	if manager.Allow("seller1", "app1", "market1", "op1") {
		t.Error("Allow() attempt 6 should be denied")
	}

	// 不同维度应该有独立的限制
	if !manager.Allow("seller2", "app1", "market1", "op1") {
		t.Error("Allow() for different seller should be allowed")
	}
}

func TestManager_Wait(t *testing.T) {
	manager := NewManager(WithDefaultRate(5, 2))
	ctx := context.Background()

	// 前 2 次应该立即通过
	start := time.Now()
	for i := 0; i < 2; i++ {
		if err := manager.Wait(ctx, "seller1", "app1", "market1", "op1"); err != nil {
			t.Errorf("Wait() attempt %d error = %v", i+1, err)
		}
	}
	elapsed := time.Since(start)

	if elapsed > 50*time.Millisecond {
		t.Errorf("First 2 waits took too long: %v", elapsed)
	}

	// 第 3 次应该等待
	start = time.Now()
	if err := manager.Wait(ctx, "seller1", "app1", "market1", "op1"); err != nil {
		t.Errorf("Wait() attempt 3 error = %v", err)
	}
	elapsed = time.Since(start)

	// 应该等待约 200ms（1/5秒）
	if elapsed < 100*time.Millisecond || elapsed > 300*time.Millisecond {
		t.Errorf("Wait time unexpected: %v (expected ~200ms)", elapsed)
	}
}

func TestManager_Wait_ContextCancellation(t *testing.T) {
	manager := NewManager(WithDefaultRate(1, 1))

	// 消耗初始令牌
	manager.Allow("seller1", "app1", "market1", "op1")

	// 创建会被取消的 context
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	// Wait 应该立即返回错误
	err := manager.Wait(ctx, "seller1", "app1", "market1", "op1")
	if err == nil {
		t.Error("Wait() should return error when context is cancelled")
	}
}

func TestManager_UpdateRate(t *testing.T) {
	manager := NewManager()

	// 更新速率
	err := manager.UpdateRate("seller1", "app1", "market1", "op1", 10.0, 20)
	if err != nil {
		t.Errorf("UpdateRate() error = %v", err)
	}

	// 验证速率已更新
	limiter := manager.GetOrCreateLimiter("seller1", "app1", "market1", "op1")
	rate, burst := limiter.GetRate()

	if rate != 10.0 {
		t.Errorf("Rate = %v, want 10.0", rate)
	}

	if burst != 20 {
		t.Errorf("Burst = %v, want 20", burst)
	}
}

func TestManager_RemoveLimiter(t *testing.T) {
	manager := NewManager()

	// 创建限制器
	manager.GetOrCreateLimiter("seller1", "app1", "market1", "op1")

	// 验证存在
	if manager.Count() != 1 {
		t.Errorf("Count() = %v, want 1", manager.Count())
	}

	// 移除限制器
	manager.RemoveLimiter("seller1", "app1", "market1", "op1")

	// 验证已移除
	if manager.Count() != 0 {
		t.Errorf("Count() after remove = %v, want 0", manager.Count())
	}
}

func TestManager_Clear(t *testing.T) {
	manager := NewManager()

	// 创建多个限制器
	manager.GetOrCreateLimiter("seller1", "app1", "market1", "op1")
	manager.GetOrCreateLimiter("seller2", "app1", "market1", "op1")
	manager.GetOrCreateLimiter("seller1", "app2", "market1", "op1")

	// 验证存在
	if manager.Count() != 3 {
		t.Errorf("Count() = %v, want 3", manager.Count())
	}

	// 清空
	manager.Clear()

	// 验证已清空
	if manager.Count() != 0 {
		t.Errorf("Count() after clear = %v, want 0", manager.Count())
	}
}

func TestManager_Count(t *testing.T) {
	manager := NewManager()

	if manager.Count() != 0 {
		t.Errorf("Initial count = %v, want 0", manager.Count())
	}

	// 添加限制器
	manager.GetOrCreateLimiter("seller1", "app1", "market1", "op1")
	manager.GetOrCreateLimiter("seller1", "app1", "market1", "op2")
	manager.GetOrCreateLimiter("seller2", "app1", "market1", "op1")

	if manager.Count() != 3 {
		t.Errorf("Count() = %v, want 3", manager.Count())
	}
}

func TestManager_ConcurrentAccess(t *testing.T) {
	manager := NewManager(WithDefaultRate(100, 50))

	const goroutines = 10
	const requestsPerGoroutine = 10
	done := make(chan bool, goroutines)

	// 并发访问同一维度
	for i := 0; i < goroutines; i++ {
		go func() {
			for j := 0; j < requestsPerGoroutine; j++ {
				ctx := context.Background()
				manager.Wait(ctx, "seller1", "app1", "market1", "op1")
			}
			done <- true
		}()
	}

	// 等待所有 goroutine 完成
	for i := 0; i < goroutines; i++ {
		<-done
	}

	// 如果没有死锁或 panic，测试通过
}

func TestManager_MultipleDimensions(t *testing.T) {
	manager := NewManager(WithDefaultRate(10, 5))

	// 不同卖家应该有独立的限制
	for i := 0; i < 5; i++ {
		if !manager.Allow("seller1", "app1", "market1", "op1") {
			t.Errorf("seller1 request %d should be allowed", i+1)
		}
		if !manager.Allow("seller2", "app1", "market1", "op1") {
			t.Errorf("seller2 request %d should be allowed", i+1)
		}
	}

	// seller1 第 6 次应该被限制
	if manager.Allow("seller1", "app1", "market1", "op1") {
		t.Error("seller1 request 6 should be denied")
	}

	// seller2 第 6 次应该被限制
	if manager.Allow("seller2", "app1", "market1", "op1") {
		t.Error("seller2 request 6 should be denied")
	}

	// 应该管理 2 个独立的限制器
	if manager.Count() != 2 {
		t.Errorf("Count() = %v, want 2", manager.Count())
	}
}

func TestBuildLimiterKey(t *testing.T) {
	tests := []struct {
		name        string
		sellerID    string
		appID       string
		marketplace string
		operation   string
		want        string
	}{
		{
			name:        "normal case",
			sellerID:    "seller1",
			appID:       "app1",
			marketplace: "market1",
			operation:   "op1",
			want:        "seller1:app1:market1:op1",
		},
		{
			name:        "with empty strings",
			sellerID:    "",
			appID:       "app1",
			marketplace: "",
			operation:   "op1",
			want:        ":app1::op1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildLimiterKey(tt.sellerID, tt.appID, tt.marketplace, tt.operation)
			if got != tt.want {
				t.Errorf("buildLimiterKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkManager_Allow(b *testing.B) {
	manager := NewManager(WithDefaultRate(1000, 500))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		manager.Allow("seller1", "app1", "market1", "op1")
	}
}

func BenchmarkManager_GetOrCreateLimiter(b *testing.B) {
	manager := NewManager()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		manager.GetOrCreateLimiter("seller1", "app1", "market1", "op1")
	}
}

