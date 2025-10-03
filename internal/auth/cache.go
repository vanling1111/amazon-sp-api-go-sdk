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

package auth

import (
	"sync"
)

// MemoryCache 是基于内存的令牌缓存实现。
//
// 此实现使用 sync.RWMutex 保证并发安全。
// 适合单机部署的场景。对于分布式部署，建议使用 Redis 等外部缓存。
type MemoryCache struct {
	mu    sync.RWMutex
	cache map[string]*Token
}

// NewMemoryCache 创建新的内存缓存实例。
//
// 返回值:
//   - *MemoryCache: 内存缓存实例
//
// 示例:
//   cache := auth.NewMemoryCache()
//   client.SetCache(cache)
func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		cache: make(map[string]*Token),
	}
}

// Get 获取缓存的令牌。
//
// 参数:
//   - key: 缓存键
//
// 返回值:
//   - *Token: 令牌对象，如果不存在返回 nil
//   - bool: 是否存在
func (c *MemoryCache) Get(key string) (*Token, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	token, ok := c.cache[key]
	if !ok {
		return nil, false
	}

	// 检查是否过期
	if token.IsExpired() {
		return nil, false
	}

	return token, true
}

// Set 设置令牌到缓存。
//
// 参数:
//   - key: 缓存键
//   - token: 令牌对象
func (c *MemoryCache) Set(key string, token *Token) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = token
}

// Delete 删除缓存的令牌。
//
// 参数:
//   - key: 缓存键
func (c *MemoryCache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.cache, key)
}

// Clear 清空所有缓存。
//
// 此方法主要用于测试。
func (c *MemoryCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache = make(map[string]*Token)
}

// Size 返回缓存中令牌的数量。
//
// 此方法主要用于监控和调试。
func (c *MemoryCache) Size() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return len(c.cache)
}
