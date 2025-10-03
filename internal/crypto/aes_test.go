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

package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestDecryptReport 测试报告解密功能
func TestDecryptReport(t *testing.T) {
	// 生成测试密钥和 IV
	key := make([]byte, 32) // AES-256
	iv := make([]byte, aes.BlockSize)
	rand.Read(key)
	rand.Read(iv)

	// 原始数据
	plaintext := []byte("This is a test report content from Amazon SP-API")

	// 加密数据
	block, err := aes.NewCipher(key)
	require.NoError(t, err)

	paddedData := addPKCS7Padding(plaintext, aes.BlockSize)
	encrypted := make([]byte, len(paddedData))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(encrypted, paddedData)

	// Base64 编码
	keyB64 := base64.StdEncoding.EncodeToString(key)
	ivB64 := base64.StdEncoding.EncodeToString(iv)

	// 测试解密
	decrypted, err := DecryptReport(keyB64, ivB64, encrypted)
	require.NoError(t, err)
	assert.Equal(t, plaintext, decrypted)
}

// TestDecryptReport_RealWorldExample 测试真实场景的解密
func TestDecryptReport_RealWorldExample(t *testing.T) {
	// 模拟 Amazon SP-API 返回的加密数据
	// （这是一个示例，实际数据来自 API）
	key := "MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI=" // Base64: 32 bytes
	iv := "MTIzNDU2Nzg5MDEyMzQ1Ng=="                      // Base64: 16 bytes

	// 创建加密数据用于测试
	keyBytes, _ := base64.StdEncoding.DecodeString(key)
	ivBytes, _ := base64.StdEncoding.DecodeString(iv)
	plaintext := []byte("OrderID,CustomerName,Amount\n123,John Doe,99.99\n")

	block, _ := aes.NewCipher(keyBytes)
	paddedData := addPKCS7Padding(plaintext, aes.BlockSize)
	encrypted := make([]byte, len(paddedData))
	mode := cipher.NewCBCEncrypter(block, ivBytes)
	mode.CryptBlocks(encrypted, paddedData)

	// 测试解密
	decrypted, err := DecryptReport(key, iv, encrypted)
	require.NoError(t, err)
	assert.Equal(t, plaintext, decrypted)
	assert.Contains(t, string(decrypted), "OrderID")
	assert.Contains(t, string(decrypted), "John Doe")
}

// TestDecryptReport_EmptyKey 测试空密钥错误
func TestDecryptReport_EmptyKey(t *testing.T) {
	_, err := DecryptReport("", "validIV", []byte("data"))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "key is required")
}

// TestDecryptReport_EmptyIV 测试空 IV 错误
func TestDecryptReport_EmptyIV(t *testing.T) {
	_, err := DecryptReport("validKey", "", []byte("data"))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "initialization vector is required")
}

// TestDecryptReport_EmptyData 测试空数据错误
func TestDecryptReport_EmptyData(t *testing.T) {
	_, err := DecryptReport("validKey", "validIV", []byte{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "data is empty")
}

// TestDecryptReport_InvalidKeyEncoding 测试无效密钥编码
func TestDecryptReport_InvalidKeyEncoding(t *testing.T) {
	_, err := DecryptReport("invalid-base64!", "MTIzNDU2Nzg5MDEyMzQ1Ng==", []byte("data"))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to decode encryption key")
}

// TestDecryptReport_InvalidIVEncoding 测试无效 IV 编码
func TestDecryptReport_InvalidIVEncoding(t *testing.T) {
	_, err := DecryptReport("MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI=", "invalid!", []byte("data"))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to decode initialization vector")
}

// TestDecryptReport_InvalidKeyLength 测试无效密钥长度
func TestDecryptReport_InvalidKeyLength(t *testing.T) {
	shortKey := base64.StdEncoding.EncodeToString([]byte("short"))
	iv := base64.StdEncoding.EncodeToString(make([]byte, 16))

	_, err := DecryptReport(shortKey, iv, make([]byte, 16))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid key length")
}

// TestEncryptDocument 测试文档加密功能
func TestEncryptDocument(t *testing.T) {
	// 原始数据
	plaintext := []byte("This is a test document to upload")

	// 加密
	details, encrypted, err := EncryptDocument(plaintext)
	require.NoError(t, err)
	assert.NotNil(t, details)
	assert.NotEmpty(t, encrypted)
	assert.Equal(t, "AES", details.Standard)
	assert.NotEmpty(t, details.Key)
	assert.NotEmpty(t, details.InitializationVector)

	// 验证可以解密回来
	decrypted, err := DecryptReport(details.Key, details.InitializationVector, encrypted)
	require.NoError(t, err)
	assert.Equal(t, plaintext, decrypted)
}

// TestEncryptDocument_EmptyData 测试空数据加密
func TestEncryptDocument_EmptyData(t *testing.T) {
	_, _, err := EncryptDocument([]byte{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "data is empty")
}

// TestValidateEncryptionDetails 测试加密详情验证
func TestValidateEncryptionDetails(t *testing.T) {
	tests := []struct {
		name    string
		details *EncryptionDetails
		wantErr bool
		errMsg  string
	}{
		{
			name:    "nil details",
			details: nil,
			wantErr: true,
			errMsg:  "encryption details is nil",
		},
		{
			name: "empty standard",
			details: &EncryptionDetails{
				Standard:             "",
				Key:                  "validKey",
				InitializationVector: "validIV",
			},
			wantErr: true,
			errMsg:  "encryption standard is required",
		},
		{
			name: "unsupported standard",
			details: &EncryptionDetails{
				Standard:             "DES",
				Key:                  "validKey",
				InitializationVector: "validIV",
			},
			wantErr: true,
			errMsg:  "unsupported encryption standard",
		},
		{
			name: "empty key",
			details: &EncryptionDetails{
				Standard:             "AES",
				Key:                  "",
				InitializationVector: "validIV",
			},
			wantErr: true,
			errMsg:  "encryption key is required",
		},
		{
			name: "empty IV",
			details: &EncryptionDetails{
				Standard:             "AES",
				Key:                  "validKey",
				InitializationVector: "",
			},
			wantErr: true,
			errMsg:  "initialization vector is required",
		},
		{
			name: "valid details",
			details: &EncryptionDetails{
				Standard:             "AES",
				Key:                  base64.StdEncoding.EncodeToString(make([]byte, 32)),
				InitializationVector: base64.StdEncoding.EncodeToString(make([]byte, 16)),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateEncryptionDetails(tt.details)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// TestPKCS7Padding 测试 PKCS7 填充和移除
func TestPKCS7Padding(t *testing.T) {
	tests := []struct {
		name      string
		data      []byte
		blockSize int
	}{
		{
			name:      "empty data",
			data:      []byte{},
			blockSize: 16,
		},
		{
			name:      "data smaller than block",
			data:      []byte("hello"),
			blockSize: 16,
		},
		{
			name:      "data equal to block",
			data:      make([]byte, 16),
			blockSize: 16,
		},
		{
			name:      "data larger than block",
			data:      []byte("this is a longer test data string"),
			blockSize: 16,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 添加填充
			padded := addPKCS7Padding(tt.data, tt.blockSize)

			// 验证填充后长度是块大小的倍数
			assert.Equal(t, 0, len(padded)%tt.blockSize)

			// 移除填充
			unpadded, err := removePKCS7Padding(padded)
			require.NoError(t, err)

			// 验证数据一致
			assert.Equal(t, tt.data, unpadded)
		})
	}
}

// TestRemovePKCS7Padding_InvalidPadding 测试无效填充
func TestRemovePKCS7Padding_InvalidPadding(t *testing.T) {
	tests := []struct {
		name string
		data []byte
	}{
		{
			name: "empty data",
			data: []byte{},
		},
		{
			name: "padding too large",
			data: []byte{1, 2, 3, 4, 17}, // 填充长度 17 > 16
		},
		{
			name: "inconsistent padding",
			data: []byte{1, 2, 3, 4, 5, 4, 4, 3}, // 填充不一致
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := removePKCS7Padding(tt.data)
			assert.Error(t, err)
		})
	}
}

// BenchmarkDecryptReport 性能测试
func BenchmarkDecryptReport(b *testing.B) {
	// 准备测试数据
	key := make([]byte, 32)
	iv := make([]byte, aes.BlockSize)
	rand.Read(key)
	rand.Read(iv)

	plaintext := make([]byte, 1024*1024) // 1MB
	rand.Read(plaintext)

	block, _ := aes.NewCipher(key)
	paddedData := addPKCS7Padding(plaintext, aes.BlockSize)
	encrypted := make([]byte, len(paddedData))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(encrypted, paddedData)

	keyB64 := base64.StdEncoding.EncodeToString(key)
	ivB64 := base64.StdEncoding.EncodeToString(iv)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = DecryptReport(keyB64, ivB64, encrypted)
	}
}

// BenchmarkEncryptDocument 加密性能测试
func BenchmarkEncryptDocument(b *testing.B) {
	data := make([]byte, 1024*1024) // 1MB
	rand.Read(data)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = EncryptDocument(data)
	}
}
