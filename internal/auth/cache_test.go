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

package auth

import (
	"testing"
	"time"
)

func TestMemoryCache_SetAndGet(t *testing.T) {
	cache := NewMemoryCache()

	key := "test-key"
	token := &Token{
		AccessToken: "test-access-token",
		TokenType:   "bearer",
		ExpiresIn:   3600,
		ExpiresAt:   time.Now().Add(1 * time.Hour),
	}

	// 设置令牌
	cache.Set(key, token)

	// 获取令牌
	got, ok := cache.Get(key)
	if !ok {
		t.Error("Get() returned false, want true")
		return
	}

	if got == nil {
		t.Error("Get() returned nil token")
		return
	}

	if got.AccessToken != token.AccessToken {
		t.Errorf("AccessToken = %v, want %v", got.AccessToken, token.AccessToken)
	}
}

func TestMemoryCache_GetNonExistent(t *testing.T) {
	cache := NewMemoryCache()

	_, ok := cache.Get("non-existent-key")
	if ok {
		t.Error("Get() returned true for non-existent key, want false")
	}
}

func TestMemoryCache_GetExpiredToken(t *testing.T) {
	cache := NewMemoryCache()

	key := "test-key"
	token := &Token{
		AccessToken: "test-access-token",
		TokenType:   "bearer",
		ExpiresIn:   1,
		ExpiresAt:   time.Now().Add(-1 * time.Hour), // 已过期
	}

	cache.Set(key, token)

	// 获取已过期的令牌应该返回 false
	_, ok := cache.Get(key)
	if ok {
		t.Error("Get() returned true for expired token, want false")
	}
}

func TestMemoryCache_Delete(t *testing.T) {
	cache := NewMemoryCache()

	key := "test-key"
	token := &Token{
		AccessToken: "test-access-token",
		TokenType:   "bearer",
		ExpiresIn:   3600,
		ExpiresAt:   time.Now().Add(1 * time.Hour),
	}

	// 设置并删除
	cache.Set(key, token)
	cache.Delete(key)

	// 应该无法获取
	_, ok := cache.Get(key)
	if ok {
		t.Error("Get() returned true after Delete(), want false")
	}
}

func TestMemoryCache_Clear(t *testing.T) {
	cache := NewMemoryCache()

	// 添加多个令牌
	for i := 0; i < 5; i++ {
		key := "test-key-" + string(rune(i))
		token := &Token{
			AccessToken: "test-access-token",
			ExpiresAt:   time.Now().Add(1 * time.Hour),
		}
		cache.Set(key, token)
	}

	// 清空
	cache.Clear()

	// 检查大小
	if cache.Size() != 0 {
		t.Errorf("Size() = %v after Clear(), want 0", cache.Size())
	}
}

func TestMemoryCache_Size(t *testing.T) {
	cache := NewMemoryCache()

	// 初始大小应该是 0
	if cache.Size() != 0 {
		t.Errorf("Size() = %v, want 0", cache.Size())
	}

	// 添加令牌
	cache.Set("key1", &Token{ExpiresAt: time.Now().Add(1 * time.Hour)})
	if cache.Size() != 1 {
		t.Errorf("Size() = %v after adding 1 token, want 1", cache.Size())
	}

	cache.Set("key2", &Token{ExpiresAt: time.Now().Add(1 * time.Hour)})
	if cache.Size() != 2 {
		t.Errorf("Size() = %v after adding 2 tokens, want 2", cache.Size())
	}
}

func TestMemoryCache_ConcurrentAccess(t *testing.T) {
	cache := NewMemoryCache()

	// 并发写入
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(n int) {
			key := "key-" + string(rune(n))
			token := &Token{
				AccessToken: "token-" + string(rune(n)),
				ExpiresAt:   time.Now().Add(1 * time.Hour),
			}
			cache.Set(key, token)
			done <- true
		}(i)
	}

	// 等待所有 goroutine 完成
	for i := 0; i < 10; i++ {
		<-done
	}

	// 检查大小
	if cache.Size() != 10 {
		t.Errorf("Size() = %v after concurrent writes, want 10", cache.Size())
	}
}

func TestToken_IsExpired(t *testing.T) {
	tests := []struct {
		name        string
		token       *Token
		wantExpired bool
	}{
		{
			name:        "nil token",
			token:       nil,
			wantExpired: true,
		},
		{
			name: "valid token",
			token: &Token{
				ExpiresAt: time.Now().Add(2 * time.Minute),
			},
			wantExpired: false,
		},
		{
			name: "expired token",
			token: &Token{
				ExpiresAt: time.Now().Add(-1 * time.Hour),
			},
			wantExpired: true,
		},
		{
			name: "token expiring soon (within 60 seconds)",
			token: &Token{
				ExpiresAt: time.Now().Add(30 * time.Second),
			},
			wantExpired: true, // 提前 60 秒判定为过期
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.token.IsExpired(); got != tt.wantExpired {
				t.Errorf("IsExpired() = %v, want %v", got, tt.wantExpired)
			}
		})
	}
}
