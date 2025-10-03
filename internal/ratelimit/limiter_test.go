package ratelimit

import (
	"context"
	"net/http"
	"testing"
	"time"
)

func TestNewLimiter(t *testing.T) {
	tests := []struct {
		name      string
		rate      float64
		burst     int
		wantErr   bool
		checkFunc func(*Limiter) bool
	}{
		{
			name:    "valid limiter",
			rate:    10.0,
			burst:   20,
			wantErr: false,
			checkFunc: func(l *Limiter) bool {
				rate, burst := l.GetRate()
				tokens := l.GetTokens()
				return rate == 10.0 && burst == 20 && tokens == 20.0
			},
		},
		{
			name:    "zero rate",
			rate:    0,
			burst:   10,
			wantErr: true,
		},
		{
			name:    "negative rate",
			rate:    -1,
			burst:   10,
			wantErr: true,
		},
		{
			name:    "zero burst",
			rate:    10,
			burst:   0,
			wantErr: true,
		},
		{
			name:    "negative burst",
			rate:    10,
			burst:   -1,
			wantErr: true,
		},
		{
			name:    "fractional rate",
			rate:    0.5,
			burst:   1,
			wantErr: false,
			checkFunc: func(l *Limiter) bool {
				rate, burst := l.GetRate()
				return rate == 0.5 && burst == 1
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			limiter, err := NewLimiter(tt.rate, tt.burst)

			if (err != nil) != tt.wantErr {
				t.Errorf("NewLimiter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && tt.checkFunc != nil {
				if !tt.checkFunc(limiter) {
					t.Error("NewLimiter() limiter configuration check failed")
				}
			}
		})
	}
}

func TestLimiter_Allow(t *testing.T) {
	// 创建一个每秒 10 个请求，突发 5 个的限制器
	limiter, err := NewLimiter(10, 5)
	if err != nil {
		t.Fatalf("NewLimiter() error = %v", err)
	}

	// 前 5 个请求应该立即通过（突发）
	for i := 0; i < 5; i++ {
		if !limiter.Allow() {
			t.Errorf("Request %d should be allowed (burst)", i+1)
		}
	}

	// 第 6 个请求应该被限制
	if limiter.Allow() {
		t.Error("Request 6 should be rate limited")
	}

	// 等待一段时间后应该可以再次请求
	time.Sleep(200 * time.Millisecond) // 等待足够生成至少 2 个令牌
	if !limiter.Allow() {
		t.Error("Request should be allowed after waiting")
	}
}

func TestLimiter_Wait(t *testing.T) {
	// 创建一个每秒 5 个请求，突发 2 个的限制器
	limiter, err := NewLimiter(5, 2)
	if err != nil {
		t.Fatalf("NewLimiter() error = %v", err)
	}

	ctx := context.Background()

	// 前 2 个请求应该立即通过
	start := time.Now()
	for i := 0; i < 2; i++ {
		if err := limiter.Wait(ctx); err != nil {
			t.Errorf("Wait() error = %v", err)
		}
	}
	elapsed := time.Since(start)

	// 前 2 个请求应该几乎不需要等待
	if elapsed > 50*time.Millisecond {
		t.Errorf("First 2 requests took too long: %v", elapsed)
	}

	// 第 3 个请求应该需要等待
	start = time.Now()
	if err := limiter.Wait(ctx); err != nil {
		t.Errorf("Wait() error = %v", err)
	}
	elapsed = time.Since(start)

	// 第 3 个请求应该等待约 200ms（1/5秒）
	if elapsed < 100*time.Millisecond || elapsed > 300*time.Millisecond {
		t.Errorf("Wait time for 3rd request unexpected: %v (expected ~200ms)", elapsed)
	}
}

func TestLimiter_Wait_ContextCancellation(t *testing.T) {
	limiter, err := NewLimiter(1, 1)
	if err != nil {
		t.Fatalf("NewLimiter() error = %v", err)
	}

	// 消耗初始令牌
	limiter.Allow()

	// 创建一个会被取消的 context
	ctx, cancel := context.WithCancel(context.Background())

	// 立即取消 context
	cancel()

	// Wait 应该立即返回错误
	err = limiter.Wait(ctx)
	if err == nil {
		t.Error("Wait() should return error when context is cancelled")
	}
	if err != context.Canceled {
		t.Errorf("Wait() error = %v, want %v", err, context.Canceled)
	}
}

func TestLimiter_Wait_ContextTimeout(t *testing.T) {
	limiter, err := NewLimiter(1, 1)
	if err != nil {
		t.Fatalf("NewLimiter() error = %v", err)
	}

	// 消耗初始令牌
	limiter.Allow()

	// 创建一个会超时的 context
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	// Wait 应该在超时后返回错误
	start := time.Now()
	err = limiter.Wait(ctx)
	elapsed := time.Since(start)

	if err == nil {
		t.Error("Wait() should return error when context times out")
	}
	if err != context.DeadlineExceeded {
		t.Errorf("Wait() error = %v, want %v", err, context.DeadlineExceeded)
	}
	if elapsed < 40*time.Millisecond || elapsed > 100*time.Millisecond {
		t.Errorf("Wait() took %v, expected ~50ms", elapsed)
	}
}

func TestLimiter_Reserve(t *testing.T) {
	limiter, err := NewLimiter(10, 5)
	if err != nil {
		t.Fatalf("NewLimiter() error = %v", err)
	}

	// 前 5 个预留应该不需要等待
	for i := 0; i < 5; i++ {
		waitTime := limiter.Reserve()
		if waitTime != 0 {
			t.Errorf("Reserve %d should have 0 wait time, got %v", i+1, waitTime)
		}
	}

	// 第 6 个预留应该需要等待
	waitTime := limiter.Reserve()
	if waitTime == 0 {
		t.Error("Reserve 6 should require wait time")
	}

	// 等待时间应该约为 100ms（1/10秒）
	if waitTime < 80*time.Millisecond || waitTime > 150*time.Millisecond {
		t.Errorf("Wait time unexpected: %v (expected ~100ms)", waitTime)
	}
}

func TestLimiter_SetRate(t *testing.T) {
	limiter, err := NewLimiter(10, 5)
	if err != nil {
		t.Fatalf("NewLimiter() error = %v", err)
	}

	tests := []struct {
		name    string
		rate    float64
		burst   int
		wantErr bool
	}{
		{
			name:    "valid update",
			rate:    20.0,
			burst:   10,
			wantErr: false,
		},
		{
			name:    "zero rate",
			rate:    0,
			burst:   10,
			wantErr: true,
		},
		{
			name:    "negative rate",
			rate:    -1,
			burst:   10,
			wantErr: true,
		},
		{
			name:    "zero burst",
			rate:    10,
			burst:   0,
			wantErr: true,
		},
		{
			name:    "reduce burst below current tokens",
			rate:    10,
			burst:   3, // 小于初始的 5
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := limiter.SetRate(tt.rate, tt.burst)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetRate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				rate, burst := limiter.GetRate()
				if rate != tt.rate || burst != tt.burst {
					t.Errorf("SetRate() rate = %v, burst = %v, want rate = %v, burst = %v",
						rate, burst, tt.rate, tt.burst)
				}
			}
		})
	}
}

func TestLimiter_GetTokens(t *testing.T) {
	limiter, err := NewLimiter(10, 5)
	if err != nil {
		t.Fatalf("NewLimiter() error = %v", err)
	}

	// 初始令牌应该等于突发限制（使用容差）
	tokens := limiter.GetTokens()
	if tokens < 4.9 || tokens > 5.1 {
		t.Errorf("GetTokens() = %v, want ~5.0", tokens)
	}

	// 消耗 2 个令牌
	limiter.Allow()
	limiter.Allow()

	// 剩余令牌应该为 3（使用容差）
	tokens = limiter.GetTokens()
	if tokens < 2.9 || tokens > 3.1 {
		t.Errorf("GetTokens() = %v, want ~3.0", tokens)
	}

	// 等待一段时间后令牌应该增加
	time.Sleep(200 * time.Millisecond)
	tokens = limiter.GetTokens()
	// 200ms * 10 req/s = 2 tokens，3 + 2 = 5
	if tokens < 4.5 || tokens > 5.0 {
		t.Errorf("GetTokens() after wait = %v, expected ~5.0", tokens)
	}
}

func TestLimiter_Refill(t *testing.T) {
	limiter, err := NewLimiter(10, 5)
	if err != nil {
		t.Fatalf("NewLimiter() error = %v", err)
	}

	// 消耗所有令牌
	for i := 0; i < 5; i++ {
		limiter.Allow()
	}

	// 验证令牌为 0
	tokens := limiter.GetTokens()
	if tokens >= 1.0 {
		t.Errorf("Tokens should be < 1.0 after consuming all, got %v", tokens)
	}

	// 等待 500ms（应该生成 5 个令牌）
	time.Sleep(500 * time.Millisecond)

	// 验证令牌已经补充
	tokens = limiter.GetTokens()
	if tokens < 4.5 || tokens > 5.0 {
		t.Errorf("Tokens after refill = %v, expected ~5.0", tokens)
	}

	// 令牌不应该超过突发限制
	time.Sleep(1 * time.Second) // 等待足够长的时间
	tokens = limiter.GetTokens()
	if tokens > 5.0 {
		t.Errorf("Tokens should not exceed burst limit, got %v", tokens)
	}
}

func TestLimiter_ConcurrentAccess(t *testing.T) {
	limiter, err := NewLimiter(100, 50)
	if err != nil {
		t.Fatalf("NewLimiter() error = %v", err)
	}

	// 并发访问测试
	const goroutines = 10
	const requestsPerGoroutine = 10
	done := make(chan bool, goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			for j := 0; j < requestsPerGoroutine; j++ {
				ctx := context.Background()
				if err := limiter.Wait(ctx); err != nil {
					t.Errorf("Wait() error = %v", err)
				}
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

func BenchmarkLimiter_Allow(b *testing.B) {
	limiter, _ := NewLimiter(1000, 500)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		limiter.Allow()
	}
}

func BenchmarkLimiter_Reserve(b *testing.B) {
	limiter, _ := NewLimiter(1000, 500)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		limiter.Reserve()
	}
}

func BenchmarkLimiter_GetTokens(b *testing.B) {
	limiter, _ := NewLimiter(1000, 500)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		limiter.GetTokens()
	}
}

func TestLimiter_UpdateFromResponse(t *testing.T) {
	limiter, err := NewLimiter(10, 5)
	if err != nil {
		t.Fatalf("NewLimiter() error = %v", err)
	}

	tests := []struct {
		name       string
		rateHeader string
		wantErr    bool
		wantRate   float64
	}{
		{
			name:       "valid rate header",
			rateHeader: "0.5",
			wantErr:    false,
			wantRate:   0.5,
		},
		{
			name:       "integer rate",
			rateHeader: "5",
			wantErr:    false,
			wantRate:   5.0,
		},
		{
			name:       "no rate header",
			rateHeader: "",
			wantErr:    false,
			wantRate:   5.0, // 保持原值
		},
		{
			name:       "invalid rate header",
			rateHeader: "invalid",
			wantErr:    true,
		},
		{
			name:       "rate with whitespace",
			rateHeader: "  2.5  ",
			wantErr:    false,
			wantRate:   2.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建模拟响应
			resp := &http.Response{
				Header: http.Header{},
			}
			if tt.rateHeader != "" {
				resp.Header.Set("x-amzn-RateLimit-Limit", tt.rateHeader)
			}

			// 更新速率
			err := limiter.UpdateFromResponse(resp)

			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateFromResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				rate, _ := limiter.GetRate()
				if rate != tt.wantRate {
					t.Errorf("After UpdateFromResponse(), rate = %v, want %v", rate, tt.wantRate)
				}
			}
		})
	}
}

