package ratelimit

import (
	"testing"
	"time"
)

func TestNewBucket(t *testing.T) {
	tests := []struct {
		name    string
		rate    float64
		burst   int
		wantErr bool
	}{
		{
			name:    "valid bucket",
			rate:    0.5,
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
			rate:    1.0,
			burst:   0,
			wantErr: true,
		},
		{
			name:    "negative burst",
			rate:    1.0,
			burst:   -1,
			wantErr: true,
		},
		{
			name:    "fractional rate",
			rate:    0.1,
			burst:   1,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bucket, err := NewBucket(tt.rate, tt.burst)

			if (err != nil) != tt.wantErr {
				t.Errorf("NewBucket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if bucket == nil {
					t.Error("NewBucket() returned nil bucket")
					return
				}

				// 验证初始状态
				if bucket.tokens != float64(tt.burst) {
					t.Errorf("Initial tokens = %v, want %v", bucket.tokens, float64(tt.burst))
				}

				if bucket.rate != tt.rate {
					t.Errorf("Rate = %v, want %v", bucket.rate, tt.rate)
				}

				if bucket.burst != tt.burst {
					t.Errorf("Burst = %v, want %v", bucket.burst, tt.burst)
				}
			}
		})
	}
}

func TestBucket_Take(t *testing.T) {
	bucket, err := NewBucket(10, 5)
	if err != nil {
		t.Fatalf("NewBucket() error = %v", err)
	}

	// 前 5 次应该成功（突发）
	for i := 0; i < 5; i++ {
		ok, waitTime := bucket.Take()
		if !ok {
			t.Errorf("Take %d failed, should succeed (burst)", i+1)
		}
		if waitTime != 0 {
			t.Errorf("Take %d wait time = %v, want 0", i+1, waitTime)
		}
	}

	// 第 6 次应该失败
	ok, waitTime := bucket.Take()
	if ok {
		t.Error("Take 6 should fail (bucket empty)")
	}
	if waitTime == 0 {
		t.Error("Take 6 wait time should be > 0")
	}

	// 等待一段时间后应该可以再次取出
	time.Sleep(200 * time.Millisecond) // 200ms * 10 req/s = 2 tokens
	ok, _ = bucket.Take()
	if !ok {
		t.Error("Take after wait should succeed")
	}
}

func TestBucket_TakeN(t *testing.T) {
	bucket, err := NewBucket(10, 10)
	if err != nil {
		t.Fatalf("NewBucket() error = %v", err)
	}

	tests := []struct {
		name      string
		n         int
		wantOk    bool
		checkWait bool
	}{
		{
			name:      "take 5 tokens",
			n:         5,
			wantOk:    true,
			checkWait: false,
		},
		{
			name:      "take 5 more tokens",
			n:         5,
			wantOk:    true,
			checkWait: false,
		},
		{
			name:      "take 1 token (should fail)",
			n:         1,
			wantOk:    false,
			checkWait: true,
		},
		{
			name:      "take 0 tokens",
			n:         0,
			wantOk:    true,
			checkWait: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok, waitTime := bucket.TakeN(tt.n)

			if ok != tt.wantOk {
				t.Errorf("TakeN(%d) ok = %v, want %v", tt.n, ok, tt.wantOk)
			}

			if tt.checkWait && waitTime == 0 {
				t.Errorf("TakeN(%d) wait time should be > 0", tt.n)
			}
		})
	}
}

func TestBucket_Available(t *testing.T) {
	bucket, err := NewBucket(10, 5)
	if err != nil {
		t.Fatalf("NewBucket() error = %v", err)
	}

	// 初始应该有 5 个令牌（使用容差）
	available := bucket.Available()
	if available < 4.9 || available > 5.1 {
		t.Errorf("Initial available = %v, want ~5.0", available)
	}

	// 取出 2 个令牌
	bucket.Take()
	bucket.Take()

	// 应该剩余 3 个令牌（使用容差比较）
	available = bucket.Available()
	if available < 2.9 || available > 3.1 {
		t.Errorf("After taking 2, available = %v, want ~3.0", available)
	}

	// 等待一段时间后令牌应该增加
	time.Sleep(200 * time.Millisecond) // 200ms * 10 req/s = 2 tokens
	available = bucket.Available()
	// 3 + 2 = 5
	if available < 4.5 || available > 5.0 {
		t.Errorf("After wait, available = %v, expected ~5.0", available)
	}
}

func TestBucket_UpdateRate(t *testing.T) {
	bucket, err := NewBucket(10, 5)
	if err != nil {
		t.Fatalf("NewBucket() error = %v", err)
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
			name:    "negative burst",
			rate:    10,
			burst:   -1,
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
			err := bucket.UpdateRate(tt.rate, tt.burst)

			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateRate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				rate, burst := bucket.GetRate()
				if rate != tt.rate || burst != tt.burst {
					t.Errorf("UpdateRate() rate = %v, burst = %v, want rate = %v, burst = %v",
						rate, burst, tt.rate, tt.burst)
				}
			}
		})
	}
}

func TestBucket_Reset(t *testing.T) {
	bucket, err := NewBucket(10, 5)
	if err != nil {
		t.Fatalf("NewBucket() error = %v", err)
	}

	// 消耗一些令牌
	bucket.Take()
	bucket.Take()
	bucket.Take()

	// 验证令牌已消耗（使用容差）
	avail := bucket.Available()
	if avail < 1.9 || avail > 2.1 {
		t.Errorf("After taking 3, available = %v, want ~2.0", avail)
	}

	// 重置
	bucket.Reset()

	// 验证桶已满（使用容差）
	available := bucket.Available()
	if available < 4.9 || available > 5.1 {
		t.Errorf("After reset, available = %v, want ~5.0", available)
	}
}

func TestBucket_Refill(t *testing.T) {
	bucket, err := NewBucket(10, 5)
	if err != nil {
		t.Fatalf("NewBucket() error = %v", err)
	}

	// 消耗所有令牌
	for i := 0; i < 5; i++ {
		bucket.Take()
	}

	// 验证桶为空
	available := bucket.Available()
	if available >= 1.0 {
		t.Errorf("After consuming all, available = %v, want < 1.0", available)
	}

	// 等待 500ms（应该生成 5 个令牌）
	time.Sleep(500 * time.Millisecond)

	// 验证令牌已补充
	available = bucket.Available()
	if available < 4.5 || available > 5.0 {
		t.Errorf("After refill, available = %v, expected ~5.0", available)
	}

	// 等待更长时间，验证不超过突发限制
	time.Sleep(1 * time.Second)
	available = bucket.Available()
	if available > 5.0 {
		t.Errorf("Available = %v, should not exceed burst limit 5.0", available)
	}
}

func TestBucket_ConcurrentAccess(t *testing.T) {
	bucket, err := NewBucket(100, 50)
	if err != nil {
		t.Fatalf("NewBucket() error = %v", err)
	}

	// 并发访问测试
	const goroutines = 10
	const takesPerGoroutine = 10
	done := make(chan bool, goroutines)
	successCount := make(chan int, goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			count := 0
			for j := 0; j < takesPerGoroutine; j++ {
				ok, _ := bucket.Take()
				if ok {
					count++
				}
			}
			successCount <- count
			done <- true
		}()
	}

	// 等待所有 goroutine 完成
	totalSuccess := 0
	for i := 0; i < goroutines; i++ {
		<-done
		totalSuccess += <-successCount
	}

	// 总成功次数应该约等于初始令牌数（可能稍多，因为有补充）
	// 但不应该超过太多
	if totalSuccess < 50 {
		t.Errorf("Total success = %d, expected at least 50 (initial burst)", totalSuccess)
	}

	// 如果没有死锁或 panic，测试通过
}

func BenchmarkBucket_Take(b *testing.B) {
	bucket, _ := NewBucket(1000, 500)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bucket.Take()
	}
}

func BenchmarkBucket_TakeN(b *testing.B) {
	bucket, _ := NewBucket(1000, 500)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bucket.TakeN(5)
	}
}

func BenchmarkBucket_Available(b *testing.B) {
	bucket, _ := NewBucket(1000, 500)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bucket.Available()
	}
}
