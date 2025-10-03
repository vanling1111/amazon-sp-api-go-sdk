// Copyright 2025 Amazon SP-API Go SDK Authors.
//
// This file is part of Amazon SP-API Go SDK.
//
// Amazon SP-API Go SDK is dual-licensed:
//
// 1. GNU Affero General Public License v3.0 (AGPL-3.0) for open source use
//   - Free for personal, educational, and open source projects
//   - Your project must also be open sourced under AGPL-3.0
//   - See: https://www.gnu.org/licenses/agpl-3.0.html
//
// 2. Commercial License for proprietary/commercial use
//   - Required for any commercial, enterprise, or proprietary use
//   - Allows closed source distribution
//   - Contact: vanling1111@gmail.com
//
// Unless you have obtained a commercial license, this file is licensed
// under AGPL-3.0. By using this software, you agree to comply with the
// terms of the applicable license. All rights reserved.
//
// Package crypto 提供加密和解密功能。
//
// 此包实现了 Amazon SP-API 报告和文档的加密/解密功能。
// Amazon SP-API 使用 AES-256-CBC 算法加密敏感报告数据。
//
// 基于官方 SP-API 文档:
//   - https://developer-docs.amazon.com/sp-api/docs/reports-api-v2021-06-30-reference
package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/pkg/errors"
)

// EncryptionDetails 包含加密所需的详细信息。
//
// 这些信息由 Amazon SP-API 在报告文档响应中提供。
type EncryptionDetails struct {
	// Standard 加密标准（通常是 "AES"）
	Standard string

	// InitializationVector CBC 模式的初始化向量（Base64 编码）
	InitializationVector string

	// Key AES 加密密钥（Base64 编码）
	Key string
}

// DecryptReport 解密 Amazon SP-API 报告。
//
// Amazon SP-API 使用 AES-256-CBC 算法加密报告数据。
// 此函数接收加密密钥、初始化向量和加密数据，返回解密后的原始数据。
//
// 参数:
//   - key: Base64 编码的 AES-256 密钥
//   - iv: Base64 编码的初始化向量（IV）
//   - encryptedData: 加密的报告数据
//
// 返回值:
//   - []byte: 解密后的原始数据
//   - error: 如果解密失败，返回错误
//
// 示例:
//
//	reportDoc, _ := client.Reports.GetReportDocument(ctx, reportDocumentID)
//	encryptedData, _ := downloadReportContent(reportDoc.URL)
//	decrypted, err := crypto.DecryptReport(
//	    reportDoc.EncryptionDetails.Key,
//	    reportDoc.EncryptionDetails.InitializationVector,
//	    encryptedData,
//	)
func DecryptReport(key, iv string, encryptedData []byte) ([]byte, error) {
	if key == "" {
		return nil, errors.New("encryption key is required")
	}
	if iv == "" {
		return nil, errors.New("initialization vector is required")
	}
	if len(encryptedData) == 0 {
		return nil, errors.New("encrypted data is empty")
	}

	// 解码 Base64 编码的 IV
	ivBytes, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode initialization vector")
	}

	// 解码 Base64 编码的密钥
	keyBytes, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode encryption key")
	}

	// 验证密钥长度（AES-256 需要 32 字节）
	if len(keyBytes) != 32 {
		return nil, fmt.Errorf("invalid key length: expected 32 bytes, got %d", len(keyBytes))
	}

	// 验证 IV 长度（AES 需要 16 字节）
	if len(ivBytes) != aes.BlockSize {
		return nil, fmt.Errorf("invalid IV length: expected %d bytes, got %d", aes.BlockSize, len(ivBytes))
	}

	// 验证加密数据长度（必须是块大小的倍数）
	if len(encryptedData)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("encrypted data length must be multiple of %d", aes.BlockSize)
	}

	// 创建 AES cipher
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create AES cipher")
	}

	// 创建 CBC 解密器
	mode := cipher.NewCBCDecrypter(block, ivBytes)

	// 解密数据
	decrypted := make([]byte, len(encryptedData))
	mode.CryptBlocks(decrypted, encryptedData)

	// 移除 PKCS7 填充
	decrypted, err = removePKCS7Padding(decrypted)
	if err != nil {
		return nil, errors.Wrap(err, "failed to remove padding")
	}

	return decrypted, nil
}

// EncryptDocument 加密文档用于上传。
//
// 某些 SP-API 操作（如 Uploads API）要求客户端加密上传的文档。
// 此函数使用 AES-256-CBC 算法加密数据，并返回加密详情。
//
// 参数:
//   - data: 要加密的原始数据
//
// 返回值:
//   - *EncryptionDetails: 加密详情（包含密钥和 IV）
//   - []byte: 加密后的数据
//   - error: 如果加密失败，返回错误
//
// 示例:
//
//	details, encrypted, err := crypto.EncryptDocument(documentData)
//	// 上传 encrypted 数据和 details
func EncryptDocument(data []byte) (*EncryptionDetails, []byte, error) {
	if len(data) == 0 {
		return nil, nil, errors.New("data is empty")
	}

	// 生成随机密钥（32 字节 = 256 位）
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return nil, nil, errors.Wrap(err, "failed to generate encryption key")
	}

	// 生成随机 IV（16 字节 = 128 位）
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return nil, nil, errors.Wrap(err, "failed to generate IV")
	}

	// 创建 AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to create AES cipher")
	}

	// 添加 PKCS7 填充
	paddedData := addPKCS7Padding(data, aes.BlockSize)

	// 创建 CBC 加密器
	mode := cipher.NewCBCEncrypter(block, iv)

	// 加密数据
	encrypted := make([]byte, len(paddedData))
	mode.CryptBlocks(encrypted, paddedData)

	// Base64 编码密钥和 IV
	details := &EncryptionDetails{
		Standard:             "AES",
		Key:                  base64.StdEncoding.EncodeToString(key),
		InitializationVector: base64.StdEncoding.EncodeToString(iv),
	}

	return details, encrypted, nil
}

// ValidateEncryptionDetails 验证加密详情的有效性。
//
// 此函数检查加密详情是否包含所有必需的字段，
// 以及密钥和 IV 是否可以正确解码。
//
// 参数:
//   - details: 要验证的加密详情
//
// 返回值:
//   - error: 如果验证失败，返回错误
func ValidateEncryptionDetails(details *EncryptionDetails) error {
	if details == nil {
		return errors.New("encryption details is nil")
	}

	if details.Standard == "" {
		return errors.New("encryption standard is required")
	}

	if details.Standard != "AES" {
		return fmt.Errorf("unsupported encryption standard: %s", details.Standard)
	}

	if details.Key == "" {
		return errors.New("encryption key is required")
	}

	if details.InitializationVector == "" {
		return errors.New("initialization vector is required")
	}

	// 验证 Key 可以解码
	keyBytes, err := base64.StdEncoding.DecodeString(details.Key)
	if err != nil {
		return errors.Wrap(err, "invalid encryption key encoding")
	}
	if len(keyBytes) != 32 {
		return fmt.Errorf("invalid key length: expected 32 bytes, got %d", len(keyBytes))
	}

	// 验证 IV 可以解码
	ivBytes, err := base64.StdEncoding.DecodeString(details.InitializationVector)
	if err != nil {
		return errors.Wrap(err, "invalid initialization vector encoding")
	}
	if len(ivBytes) != aes.BlockSize {
		return fmt.Errorf("invalid IV length: expected %d bytes, got %d", aes.BlockSize, len(ivBytes))
	}

	return nil
}

// addPKCS7Padding 添加 PKCS7 填充。
//
// PKCS7 填充用于确保数据长度是块大小的倍数。
func addPKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padtext := make([]byte, padding)
	for i := range padtext {
		padtext[i] = byte(padding)
	}
	return append(data, padtext...)
}

// removePKCS7Padding 移除 PKCS7 填充。
//
// 从解密后的数据中移除填充字节。
func removePKCS7Padding(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, errors.New("data is empty")
	}

	// 获取填充长度（最后一个字节）
	padding := int(data[len(data)-1])

	// 验证填充长度
	if padding == 0 || padding > aes.BlockSize {
		return nil, fmt.Errorf("invalid padding length: %d", padding)
	}

	if len(data) < padding {
		return nil, fmt.Errorf("data length (%d) < padding length (%d)", len(data), padding)
	}

	// 验证填充内容（所有填充字节应该相同）
	for i := len(data) - padding; i < len(data); i++ {
		if data[i] != byte(padding) {
			return nil, fmt.Errorf("invalid padding at position %d", i)
		}
	}

	// 移除填充
	return data[:len(data)-padding], nil
}
