// Package ratelimit 提供 SP-API 速率限制功能。
package ratelimit

import (
	"fmt"
	"sync"
	"time"
)

// Bucket 表示 Token Bucket（令牌桶）。
//
// Token Bucket 是一种速率限制算法，用于控制请求速率。
// 桶中有一定数量的令牌，每个请求消耗一个令牌。
// 令牌以固定速率补充，直到达到桶的最大容量（突发限制）。
//
// 根据官方 SP-API 文档：
//   - 桶会以固定速率自动补充令牌
//   - 突发限制是桶的最大容量
//   - 每个请求消耗一个令牌
//   - 当桶为空时，请求会被限流
//
// 官方文档:
//   - https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits#rate-limiting-algorithm
type Bucket struct {
	// rate 是每秒补充的令牌数
	rate float64

	// burst 是桶的最大容量（突发限制）
	burst int

	// tokens 是当前可用的令牌数
	tokens float64

	// lastRefill 是上次补充令牌的时间
	lastRefill time.Time

	// mu 保护并发访问
	mu sync.Mutex
}

// NewBucket 创建新的 Token Bucket。
//
// 参数:
//   - rate: 每秒补充的令牌数（速率限制）
//   - burst: 桶的最大容量（突发限制）
//
// 返回值:
//   - *Bucket: Token Bucket 实例
//   - error: 如果参数无效，返回错误
//
// 示例:
//
//	// 创建每秒 0.5 个令牌，最大容量 10 的桶
//	bucket, err := ratelimit.NewBucket(0.5, 10)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// 官方文档:
//   - https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits#rate-limiting-algorithm
func NewBucket(rate float64, burst int) (*Bucket, error) {
	if rate <= 0 {
		return nil, fmt.Errorf("rate must be positive, got: %f", rate)
	}
	if burst <= 0 {
		return nil, fmt.Errorf("burst must be positive, got: %d", burst)
	}

	return &Bucket{
		rate:       rate,
		burst:      burst,
		tokens:     float64(burst), // 初始时桶是满的
		lastRefill: time.Now(),
	}, nil
}

// Take 尝试从桶中取出一个令牌。
//
// 如果桶中有可用令牌，则取出一个令牌并返回 true。
// 否则返回 false 和需要等待的时间。
//
// 返回值:
//   - bool: 如果成功取出令牌返回 true，否则返回 false
//   - time.Duration: 如果失败，返回需要等待的时间
//
// 示例:
//
//	ok, waitTime := bucket.Take()
//	if !ok {
//	    time.Sleep(waitTime)
//	    // 重试
//	}
func (b *Bucket) Take() (bool, time.Duration) {
	b.mu.Lock()
	defer b.mu.Unlock()

	// 先补充令牌
	b.refillLocked()

	// 检查是否有可用令牌
	if b.tokens >= 1.0 {
		b.tokens -= 1.0
		return true, 0
	}

	// 计算需要等待的时间
	waitTime := b.calculateWaitTimeLocked()
	return false, waitTime
}

// TakeN 尝试从桶中取出 n 个令牌。
//
// 参数:
//   - n: 需要取出的令牌数
//
// 返回值:
//   - bool: 如果成功取出令牌返回 true，否则返回 false
//   - time.Duration: 如果失败，返回需要等待的时间
//
// 示例:
//
//	// 批量操作可能需要多个令牌
//	ok, waitTime := bucket.TakeN(5)
func (b *Bucket) TakeN(n int) (bool, time.Duration) {
	if n <= 0 {
		return true, 0
	}

	b.mu.Lock()
	defer b.mu.Unlock()

	// 先补充令牌
	b.refillLocked()

	// 检查是否有足够的令牌
	if b.tokens >= float64(n) {
		b.tokens -= float64(n)
		return true, 0
	}

	// 计算需要生成 n 个令牌的时间
	tokensNeeded := float64(n) - b.tokens
	waitSeconds := tokensNeeded / b.rate
	waitTime := time.Duration(waitSeconds * float64(time.Second))

	if waitTime < time.Millisecond {
		waitTime = time.Millisecond
	}

	return false, waitTime
}

// Available 返回当前可用的令牌数。
//
// 返回值:
//   - float64: 当前可用的令牌数
//
// 示例:
//
//	available := bucket.Available()
//	fmt.Printf("Available tokens: %.2f\n", available)
func (b *Bucket) Available() float64 {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.refillLocked()
	return b.tokens
}

// UpdateRate 更新桶的速率和突发限制。
//
// 此方法用于根据 API 响应头动态调整速率限制。
//
// 参数:
//   - rate: 新的每秒令牌补充速率
//   - burst: 新的桶容量
//
// 返回值:
//   - error: 如果参数无效，返回错误
//
// 示例:
//
//	// 从 API 响应头获取新的速率限制
//	newRate := parseRateLimitHeader(resp.Header.Get("x-amzn-RateLimit-Limit"))
//	bucket.UpdateRate(newRate, newBurst)
//
// 官方文档:
//   - https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits#how-to-find-your-usage-plan
func (b *Bucket) UpdateRate(rate float64, burst int) error {
	if rate <= 0 {
		return fmt.Errorf("rate must be positive, got: %f", rate)
	}
	if burst <= 0 {
		return fmt.Errorf("burst must be positive, got: %d", burst)
	}

	b.mu.Lock()
	defer b.mu.Unlock()

	b.rate = rate
	b.burst = burst

	// 如果当前令牌数超过新的突发限制，调整为突发限制
	if b.tokens > float64(burst) {
		b.tokens = float64(burst)
	}

	return nil
}

// GetRate 获取当前的速率配置。
//
// 返回值:
//   - float64: 每秒令牌补充速率
//   - int: 桶容量（突发限制）
func (b *Bucket) GetRate() (rate float64, burst int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	return b.rate, b.burst
}

// Reset 重置桶到初始状态（桶满）。
//
// 示例:
//
//	bucket.Reset()  // 桶重新装满
func (b *Bucket) Reset() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.tokens = float64(b.burst)
	b.lastRefill = time.Now()
}

// refillLocked 补充令牌（已持锁）。
//
// 根据经过的时间按速率补充令牌。
// 调用此方法前必须持有 b.mu 锁。
func (b *Bucket) refillLocked() {
	now := time.Now()
	elapsed := now.Sub(b.lastRefill)

	// 计算应该补充的令牌数
	tokensToAdd := elapsed.Seconds() * b.rate

	// 更新令牌数，不超过突发限制
	b.tokens += tokensToAdd
	if b.tokens > float64(b.burst) {
		b.tokens = float64(b.burst)
	}

	b.lastRefill = now
}

// calculateWaitTimeLocked 计算需要等待的时间（已持锁）。
//
// 调用此方法前必须持有 b.mu 锁。
//
// 返回值:
//   - time.Duration: 需要等待的时间
func (b *Bucket) calculateWaitTimeLocked() time.Duration {
	// 需要生成 1 个令牌
	tokensNeeded := 1.0 - b.tokens

	// 计算需要等待的时间
	waitSeconds := tokensNeeded / b.rate

	// 转换为 Duration，至少等待 1 毫秒
	waitTime := time.Duration(waitSeconds * float64(time.Second))
	if waitTime < time.Millisecond {
		waitTime = time.Millisecond
	}

	return waitTime
}

