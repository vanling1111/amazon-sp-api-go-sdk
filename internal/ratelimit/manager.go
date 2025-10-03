// Package ratelimit 提供 SP-API 速率限制功能。
package ratelimit

import (
	"context"
	"net/http"
	"strings"
	"sync"
)

// Manager 管理多个速率限制器。
//
// 根据官方 SP-API 文档，速率限制是基于多个维度的：
//   - Selling Partner (卖家账号)
//   - Application (应用)
//   - Marketplace (市场)
//   - Operation (API 操作)
//
// Manager 为每个组合维护独立的 Limiter。
//
// 官方文档:
//   - https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits#factors-that-determine-usage-plans
type Manager struct {
	// limiters 存储所有限制器
	// key 格式: "sellerID:appID:marketplace:operation"
	limiters map[string]*Limiter

	// mu 保护并发访问
	mu sync.RWMutex

	// defaultRate 默认速率
	defaultRate float64

	// defaultBurst 默认突发限制
	defaultBurst int
}

// ManagerOption 表示管理器选项。
type ManagerOption func(*Manager)

// WithDefaultRate 设置默认速率。
//
// 参数:
//   - rate: 默认速率（请求数/秒）
//   - burst: 默认突发限制
//
// 示例:
//
//	manager := ratelimit.NewManager(
//	    ratelimit.WithDefaultRate(1.0, 5),
//	)
func WithDefaultRate(rate float64, burst int) ManagerOption {
	return func(m *Manager) {
		m.defaultRate = rate
		m.defaultBurst = burst
	}
}

// NewManager 创建新的速率限制管理器。
//
// 参数:
//   - opts: 管理器选项
//
// 返回值:
//   - *Manager: 管理器实例
//
// 示例:
//
//	// 创建默认管理器
//	manager := ratelimit.NewManager()
//
//	// 创建自定义默认速率的管理器
//	manager := ratelimit.NewManager(
//	    ratelimit.WithDefaultRate(1.0, 5),
//	)
//
// 官方文档:
//   - https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits
func NewManager(opts ...ManagerOption) *Manager {
	m := &Manager{
		limiters:     make(map[string]*Limiter),
		defaultRate:  1.0, // 默认每秒 1 个请求
		defaultBurst: 5,   // 默认突发 5 个
	}

	for _, opt := range opts {
		opt(m)
	}

	return m
}

// Wait 等待直到可以发送请求。
//
// 根据维度参数选择或创建对应的限制器。
//
// 参数:
//   - ctx: 请求上下文
//   - sellerID: 卖家 ID
//   - appID: 应用 ID
//   - marketplace: 市场 ID
//   - operation: 操作名称
//
// 返回值:
//   - error: 如果等待失败，返回错误
//
// 示例:
//
//	err := manager.Wait(ctx, "seller123", "app456", "ATVPDKIKX0DER", "getOrders")
//	if err != nil {
//	    return err
//	}
func (m *Manager) Wait(ctx context.Context, sellerID, appID, marketplace, operation string) error {
	limiter := m.GetOrCreateLimiter(sellerID, appID, marketplace, operation)
	return limiter.Wait(ctx)
}

// Allow 检查是否允许当前请求。
//
// 参数:
//   - sellerID: 卖家 ID
//   - appID: 应用 ID
//   - marketplace: 市场 ID
//   - operation: 操作名称
//
// 返回值:
//   - bool: 如果允许请求返回 true，否则返回 false
//
// 示例:
//
//	if manager.Allow("seller123", "app456", "ATVPDKIKX0DER", "getOrders") {
//	    // 发送请求
//	}
func (m *Manager) Allow(sellerID, appID, marketplace, operation string) bool {
	limiter := m.GetOrCreateLimiter(sellerID, appID, marketplace, operation)
	return limiter.Allow()
}

// GetOrCreateLimiter 获取或创建限制器。
//
// 参数:
//   - sellerID: 卖家 ID
//   - appID: 应用 ID
//   - marketplace: 市场 ID
//   - operation: 操作名称
//
// 返回值:
//   - *Limiter: 限制器实例
//
// 示例:
//
//	limiter := manager.GetOrCreateLimiter("seller123", "app456", "ATVPDKIKX0DER", "getOrders")
func (m *Manager) GetOrCreateLimiter(sellerID, appID, marketplace, operation string) *Limiter {
	key := buildLimiterKey(sellerID, appID, marketplace, operation)

	// 首先尝试读锁
	m.mu.RLock()
	limiter, exists := m.limiters[key]
	m.mu.RUnlock()

	if exists {
		return limiter
	}

	// 需要创建，使用写锁
	m.mu.Lock()
	defer m.mu.Unlock()

	// 双重检查（可能在等待写锁期间被其他 goroutine 创建）
	if limiter, exists := m.limiters[key]; exists {
		return limiter
	}

	// 创建新的限制器
	limiter, _ = NewLimiter(m.defaultRate, m.defaultBurst)
	m.limiters[key] = limiter

	return limiter
}

// UpdateRate 更新指定维度的速率限制。
//
// 参数:
//   - sellerID: 卖家 ID
//   - appID: 应用 ID
//   - marketplace: 市场 ID
//   - operation: 操作名称
//   - rate: 新的速率
//   - burst: 新的突发限制
//
// 返回值:
//   - error: 如果更新失败，返回错误
//
// 示例:
//
//	// 从 API 响应头获取实际速率后更新
//	manager.UpdateRate("seller123", "app456", "ATVPDKIKX0DER", "getOrders", 2.0, 10)
func (m *Manager) UpdateRate(sellerID, appID, marketplace, operation string, rate float64, burst int) error {
	limiter := m.GetOrCreateLimiter(sellerID, appID, marketplace, operation)
	return limiter.SetRate(rate, burst)
}

// UpdateFromResponse 从 HTTP 响应头更新指定维度的速率限制。
//
// SP-API 在响应头中返回 `x-amzn-RateLimit-Limit`，指示当前操作的速率限制。
// 此方法会解析该头部并自动更新对应的限制器。
//
// 参数:
//   - sellerID: 卖家 ID
//   - appID: 应用 ID
//   - marketplace: 市场 ID
//   - operation: 操作名称
//   - resp: HTTP 响应
//
// 返回值:
//   - error: 如果解析失败，返回错误
//
// 示例:
//
//	resp, err := httpClient.Do(ctx, req)
//	if err != nil {
//	    return err
//	}
//	// 自动更新速率限制
//	manager.UpdateFromResponse("seller123", "app456", "ATVPDKIKX0DER", "getOrders", resp)
//
// 官方文档:
//   - https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits#how-to-find-your-usage-plan
func (m *Manager) UpdateFromResponse(sellerID, appID, marketplace, operation string, resp *http.Response) error {
	limiter := m.GetOrCreateLimiter(sellerID, appID, marketplace, operation)
	return limiter.UpdateFromResponse(resp)
}

// RemoveLimiter 移除指定维度的限制器。
//
// 用于释放不再使用的限制器，避免内存泄漏。
//
// 参数:
//   - sellerID: 卖家 ID
//   - appID: 应用 ID
//   - marketplace: 市场 ID
//   - operation: 操作名称
//
// 示例:
//
//	manager.RemoveLimiter("seller123", "app456", "ATVPDKIKX0DER", "getOrders")
func (m *Manager) RemoveLimiter(sellerID, appID, marketplace, operation string) {
	key := buildLimiterKey(sellerID, appID, marketplace, operation)

	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.limiters, key)
}

// Clear 清空所有限制器。
//
// 示例:
//
//	manager.Clear()
func (m *Manager) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.limiters = make(map[string]*Limiter)
}

// Count 返回当前管理的限制器数量。
//
// 返回值:
//   - int: 限制器数量
//
// 示例:
//
//	count := manager.Count()
//	fmt.Printf("Managing %d limiters\n", count)
func (m *Manager) Count() int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return len(m.limiters)
}

// buildLimiterKey 构建限制器键。
//
// 格式: "sellerID:appID:marketplace:operation"
//
// 使用 strings.Builder 以提高性能，避免多次内存分配。
func buildLimiterKey(sellerID, appID, marketplace, operation string) string {
	var builder strings.Builder
	// 预分配足够容量，避免扩容
	builder.Grow(len(sellerID) + len(appID) + len(marketplace) + len(operation) + 3) // +3 for colons

	builder.WriteString(sellerID)
	builder.WriteByte(':')
	builder.WriteString(appID)
	builder.WriteByte(':')
	builder.WriteString(marketplace)
	builder.WriteByte(':')
	builder.WriteString(operation)

	return builder.String()
}
