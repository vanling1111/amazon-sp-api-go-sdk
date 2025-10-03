// Package ratelimit 提供 SP-API 速率限制功能。
//
// 此包实现了 Token Bucket 算法，用于控制 API 请求速率，
// 确保符合 Amazon SP-API 的速率限制要求。
//
// 官方文档:
//   - https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits
package ratelimit

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Limiter 表示速率限制器。
//
// 基于 Token Bucket 实现，提供高层次的速率限制接口。
// 支持阻塞等待、非阻塞检查和从 API 响应头动态更新速率。
//
// 速率限制器是并发安全的，可以在多个 goroutine 中共享使用。
//
// 官方文档:
//   - https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits
type Limiter struct {
	// bucket 是底层的 Token Bucket
	bucket *Bucket
}

// NewLimiter 创建新的速率限制器。
//
// 参数:
//   - rate: 允许的请求速率（请求数/秒）
//   - burst: 允许的突发请求数
//
// 返回值:
//   - *Limiter: 速率限制器实例
//   - error: 如果参数无效，返回错误
//
// 示例:
//
//	// 创建每秒 0.5 个请求，突发 10 个的限制器
//	limiter, err := ratelimit.NewLimiter(0.5, 10)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// 官方文档:
//   - https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits
func NewLimiter(rate float64, burst int) (*Limiter, error) {
	bucket, err := NewBucket(rate, burst)
	if err != nil {
		return nil, err
	}

	return &Limiter{
		bucket: bucket,
	}, nil
}

// Wait 等待直到可以发送请求。
//
// 此方法会阻塞当前 goroutine，直到有可用的令牌。
// 如果 context 被取消，方法会立即返回错误。
//
// 参数:
//   - ctx: 请求上下文
//
// 返回值:
//   - error: 如果 context 被取消，返回错误；否则返回 nil
//
// 示例:
//
//	ctx := context.Background()
//	if err := limiter.Wait(ctx); err != nil {
//	    log.Printf("rate limit wait failed: %v", err)
//	    return err
//	}
//	// 发送 API 请求
func (l *Limiter) Wait(ctx context.Context) error {
	for {
		// 检查 context 是否已取消
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// 尝试获取令牌
		ok, waitTime := l.bucket.Take()
		if ok {
			return nil
		}

		// 等待一段时间后重试
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(waitTime):
			// 继续循环
		}
	}
}

// Allow 检查是否允许当前请求。
//
// 此方法是非阻塞的，如果当前有可用令牌则返回 true 并消耗一个令牌，
// 否则返回 false。
//
// 返回值:
//   - bool: 如果允许请求返回 true，否则返回 false
//
// 示例:
//
//	if limiter.Allow() {
//	    // 发送 API 请求
//	} else {
//	    // 等待或稍后重试
//	}
func (l *Limiter) Allow() bool {
	ok, _ := l.bucket.Take()
	return ok
}

// Reserve 预留令牌并返回等待时间。
//
// 此方法不会阻塞，而是返回需要等待的时间。
// 如果当前有可用令牌，等待时间为 0。
//
// 返回值:
//   - time.Duration: 需要等待的时间
//
// 示例:
//
//	waitTime := limiter.Reserve()
//	if waitTime > 0 {
//	    time.Sleep(waitTime)
//	}
//	// 发送 API 请求
func (l *Limiter) Reserve() time.Duration {
	ok, waitTime := l.bucket.Take()
	if ok {
		return 0
	}
	return waitTime
}

// UpdateFromResponse 从 HTTP 响应头更新速率限制。
//
// 此方法解析 x-amzn-RateLimit-Limit 响应头并更新速率配置。
//
// 根据官方 SP-API 文档：
//   - 响应头格式：x-amzn-RateLimit-Limit: rate
//   - 例如：x-amzn-RateLimit-Limit: 0.5 表示每秒 0.5 个请求
//   - 不是所有响应都包含此头部
//
// 参数:
//   - resp: HTTP 响应
//
// 返回值:
//   - error: 如果解析失败，返回错误；如果头部不存在，返回 nil
//
// 示例:
//
//	resp, err := client.Do(req)
//	if err != nil {
//	    return err
//	}
//	defer resp.Body.Close()
//
//	// 从响应头动态更新速率
//	if err := limiter.UpdateFromResponse(resp); err != nil {
//	    log.Printf("failed to update rate limit: %v", err)
//	}
//
// 官方文档:
//   - https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits#how-to-find-your-usage-plan
func (l *Limiter) UpdateFromResponse(resp *http.Response) error {
	// 获取 x-amzn-RateLimit-Limit 头部
	rateLimitHeader := resp.Header.Get("x-amzn-RateLimit-Limit")
	if rateLimitHeader == "" {
		// 头部不存在，不是错误
		return nil
	}

	// 解析速率值
	rate, err := strconv.ParseFloat(strings.TrimSpace(rateLimitHeader), 64)
	if err != nil {
		return fmt.Errorf("parse rate limit header: %w", err)
	}

	// 获取当前的突发限制
	_, burst := l.bucket.GetRate()

	// 更新速率（保持原有的突发限制）
	return l.bucket.UpdateRate(rate, burst)
}

// SetRate 动态更新速率限制。
//
// 此方法允许在运行时更改速率限制配置。
//
// 参数:
//   - rate: 新的请求速率（请求数/秒）
//   - burst: 新的突发请求数
//
// 返回值:
//   - error: 如果参数无效，返回错误
//
// 示例:
//
//	// 根据 API 响应头动态调整速率
//	if err := limiter.SetRate(newRate, newBurst); err != nil {
//	    log.Printf("failed to update rate limit: %v", err)
//	}
func (l *Limiter) SetRate(rate float64, burst int) error {
	return l.bucket.UpdateRate(rate, burst)
}

// GetTokens 获取当前可用的令牌数。
//
// 返回值:
//   - float64: 当前可用的令牌数
//
// 示例:
//
//	tokens := limiter.GetTokens()
//	log.Printf("Available tokens: %.2f", tokens)
func (l *Limiter) GetTokens() float64 {
	return l.bucket.Available()
}

// GetRate 获取当前的速率配置。
//
// 返回值:
//   - float64: 请求速率（请求数/秒）
//   - int: 突发请求数
//
// 示例:
//
//	rate, burst := limiter.GetRate()
//	log.Printf("Rate: %.2f req/s, Burst: %d", rate, burst)
func (l *Limiter) GetRate() (rate float64, burst int) {
	return l.bucket.GetRate()
}
