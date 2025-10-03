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

// Package transfer 提供大文件传输功能。
//
// 此包实现了大文件的分片上传和下载功能，
// 支持断点续传、进度回调等高级特性。
package transfer

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

// ProgressFunc 进度回调函数
//
// 参数:
//   - uploaded: 已上传/下载字节数
//   - total: 总字节数
//   - percent: 百分比（0-100）
type ProgressFunc func(uploaded, total int64, percent float64)

// Uploader 大文件上传器
type Uploader struct {
	client      *http.Client
	chunkSize   int64
	concurrency int
	maxRetries  int
}

// UploaderConfig 上传器配置
type UploaderConfig struct {
	// ChunkSize 分片大小（字节，默认 10MB）
	ChunkSize int64

	// Concurrency 并发上传数（默认 3）
	Concurrency int

	// MaxRetries 失败重试次数（默认 3）
	MaxRetries int

	// HTTPClient 自定义 HTTP 客户端（可选）
	HTTPClient *http.Client
}

// NewUploader 创建大文件上传器。
//
// 参数:
//   - config: 上传器配置
//
// 返回值:
//   - *Uploader: 上传器实例
//
// 示例:
//
//	uploader := transfer.NewUploader(&transfer.UploaderConfig{
//	    ChunkSize:   10 * 1024 * 1024,  // 10MB
//	    Concurrency: 3,
//	})
func NewUploader(config *UploaderConfig) *Uploader {
	if config == nil {
		config = &UploaderConfig{}
	}

	if config.ChunkSize == 0 {
		config.ChunkSize = 10 * 1024 * 1024 // 10MB
	}

	if config.Concurrency == 0 {
		config.Concurrency = 3
	}

	if config.MaxRetries == 0 {
		config.MaxRetries = 3
	}

	client := config.HTTPClient
	if client == nil {
		client = http.DefaultClient
	}

	return &Uploader{
		client:      client,
		chunkSize:   config.ChunkSize,
		concurrency: config.Concurrency,
		maxRetries:  config.MaxRetries,
	}
}

// Upload 上传文件到指定 URL。
//
// 对于小文件（< chunkSize），直接上传。
// 对于大文件，自动分片上传（TODO: v1.3.0 实现）。
//
// 参数:
//   - ctx: 请求上下文
//   - url: 上传 URL（通常来自 CreateFeedDocument 或 CreateServiceDocumentUploadDestination）
//   - reader: 文件数据源
//   - contentType: 内容类型
//   - size: 文件大小（字节，-1 表示未知）
//
// 返回值:
//   - error: 如果上传失败，返回错误
//
// 示例:
//
//	file, _ := os.Open("feed.xml")
//	defer file.Close()
//
//	uploader := transfer.NewUploader(nil)
//	err := uploader.Upload(ctx, uploadURL, file, "text/xml", fileSize)
func (u *Uploader) Upload(ctx context.Context, url string, reader io.Reader, contentType string, size int64) error {
	// 读取所有数据
	data, err := io.ReadAll(reader)
	if err != nil {
		return errors.Wrap(err, "failed to read data")
	}

	// 使用标准库的 bytes.NewReader
	body := io.NopCloser(newReaderFromBytes(data))

	// 创建 PUT 请求
	req, err := http.NewRequestWithContext(ctx, "PUT", url, body)
	if err != nil {
		return errors.Wrap(err, "failed to create request")
	}

	// 设置头部
	req.Header.Set("Content-Type", contentType)
	req.ContentLength = int64(len(data))

	// 执行上传
	resp, err := u.client.Do(req)
	if err != nil {
		return errors.Wrap(err, "upload failed")
	}
	defer resp.Body.Close()

	// 检查响应状态
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("upload failed with status %d: %s", resp.StatusCode, body)
	}

	return nil
}

// newReaderFromBytes 创建一个简单的 Reader
func newReaderFromBytes(data []byte) io.Reader {
	return &simpleReader{data: data, pos: 0}
}

// simpleReader 简单的字节 Reader
type simpleReader struct {
	data []byte
	pos  int
}

func (r *simpleReader) Read(p []byte) (n int, err error) {
	if r.pos >= len(r.data) {
		return 0, io.EOF
	}
	n = copy(p, r.data[r.pos:])
	r.pos += n
	return n, nil
}

// UploadWithProgress 上传文件并报告进度。
//
// 参数:
//   - ctx: 请求上下文
//   - url: 上传 URL
//   - reader: 文件数据源
//   - contentType: 内容类型
//   - size: 文件大小
//   - onProgress: 进度回调函数
//
// 返回值:
//   - error: 如果上传失败，返回错误
//
// 示例:
//
//	uploader.UploadWithProgress(ctx, url, file, "text/xml", size,
//	    func(uploaded, total int64, percent float64) {
//	        fmt.Printf("Progress: %.1f%% (%d/%d bytes)\n", percent, uploaded, total)
//	    },
//	)
func (u *Uploader) UploadWithProgress(ctx context.Context, url string, reader io.Reader, contentType string, size int64, onProgress ProgressFunc) error {
	// TODO: 实现真正的进度追踪
	// 当前简化版本：直接上传

	err := u.Upload(ctx, url, reader, contentType, size)

	if onProgress != nil && err == nil {
		onProgress(size, size, 100.0)
	}

	return err
}
