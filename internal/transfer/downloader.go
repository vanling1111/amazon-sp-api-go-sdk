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
package transfer

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

// Downloader 大文件下载器
type Downloader struct {
	client     *http.Client
	bufferSize int64
}

// DownloaderConfig 下载器配置
type DownloaderConfig struct {
	// BufferSize 缓冲区大小（字节，默认 32KB）
	BufferSize int64

	// HTTPClient 自定义 HTTP 客户端（可选）
	HTTPClient *http.Client
}

// NewDownloader 创建大文件下载器。
//
// 参数:
//   - config: 下载器配置
//
// 返回值:
//   - *Downloader: 下载器实例
func NewDownloader(config *DownloaderConfig) *Downloader {
	if config == nil {
		config = &DownloaderConfig{}
	}

	if config.BufferSize == 0 {
		config.BufferSize = 32 * 1024 // 32KB
	}

	client := config.HTTPClient
	if client == nil {
		client = http.DefaultClient
	}

	return &Downloader{
		client:     client,
		bufferSize: config.BufferSize,
	}
}

// Download 下载文件到 Writer。
//
// 参数:
//   - ctx: 请求上下文
//   - url: 下载 URL
//   - writer: 数据写入目标
//
// 返回值:
//   - int64: 下载的字节数
//   - error: 如果下载失败，返回错误
//
// 示例:
//
//	file, _ := os.Create("report.csv")
//	defer file.Close()
//
//	downloader := transfer.NewDownloader(nil)
//	size, err := downloader.Download(ctx, reportURL, file)
func (d *Downloader) Download(ctx context.Context, url string, writer io.Writer) (int64, error) {
	// 创建 GET 请求
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return 0, errors.Wrap(err, "failed to create request")
	}

	// 执行请求
	resp, err := d.client.Do(req)
	if err != nil {
		return 0, errors.Wrap(err, "download failed")
	}
	defer resp.Body.Close()

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return 0, fmt.Errorf("download failed with status %d: %s", resp.StatusCode, body)
	}

	// 复制数据到 writer
	written, err := io.Copy(writer, resp.Body)
	if err != nil {
		return written, errors.Wrap(err, "failed to write data")
	}

	return written, nil
}

// DownloadWithProgress 下载文件并报告进度。
//
// 参数:
//   - ctx: 请求上下文
//   - url: 下载 URL
//   - writer: 数据写入目标
//   - onProgress: 进度回调函数
//
// 返回值:
//   - int64: 下载的字节数
//   - error: 如果下载失败，返回错误
func (d *Downloader) DownloadWithProgress(ctx context.Context, url string, writer io.Writer, onProgress ProgressFunc) (int64, error) {
	// 创建请求
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return 0, errors.Wrap(err, "failed to create request")
	}

	// 执行请求
	resp, err := d.client.Do(req)
	if err != nil {
		return 0, errors.Wrap(err, "download failed")
	}
	defer resp.Body.Close()

	// 检查状态
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("download failed with status: %d", resp.StatusCode)
	}

	// 获取总大小
	totalSize := resp.ContentLength

	// 使用进度 Reader
	progressReader := &progressReader{
		reader:     resp.Body,
		totalSize:  totalSize,
		onProgress: onProgress,
	}

	// 复制数据
	written, err := io.Copy(writer, progressReader)
	if err != nil {
		return written, errors.Wrap(err, "failed to write data")
	}

	return written, nil
}

// progressReader 带进度追踪的 Reader
type progressReader struct {
	reader     io.Reader
	totalSize  int64
	downloaded int64
	onProgress ProgressFunc
}

func (p *progressReader) Read(buf []byte) (int, error) {
	n, err := p.reader.Read(buf)

	p.downloaded += int64(n)

	// 调用进度回调
	if p.onProgress != nil && p.totalSize > 0 {
		percent := float64(p.downloaded) / float64(p.totalSize) * 100.0
		p.onProgress(p.downloaded, p.totalSize, percent)
	}

	return n, err
}
