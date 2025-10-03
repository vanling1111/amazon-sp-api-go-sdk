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
package transfer

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestNewUploader tests creating uploader
func TestNewUploader(t *testing.T) {
	uploader := NewUploader(nil)
	assert.NotNil(t, uploader)
	assert.Equal(t, int64(10*1024*1024), uploader.chunkSize)
	assert.Equal(t, 3, uploader.concurrency)
}

// TestUpload tests basic upload
func TestUpload(t *testing.T) {
	// 创建测试服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PUT", r.Method)
		assert.Equal(t, "text/plain", r.Header.Get("Content-Type"))

		body, _ := io.ReadAll(r.Body)
		assert.Equal(t, "test data", string(body))

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	uploader := NewUploader(&UploaderConfig{
		HTTPClient: server.Client(),
	})

	data := strings.NewReader("test data")
	err := uploader.Upload(context.Background(), server.URL, data, "text/plain", 9)
	assert.NoError(t, err)
}

// TestNewDownloader tests creating downloader
func TestNewDownloader(t *testing.T) {
	downloader := NewDownloader(nil)
	assert.NotNil(t, downloader)
	assert.Equal(t, int64(32*1024), downloader.bufferSize)
}

// TestDownload tests basic download
func TestDownload(t *testing.T) {
	testData := "test download data"

	// 创建测试服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(testData))
	}))
	defer server.Close()

	downloader := NewDownloader(&DownloaderConfig{
		HTTPClient: server.Client(),
	})

	var buf bytes.Buffer
	size, err := downloader.Download(context.Background(), server.URL, &buf)

	require.NoError(t, err)
	assert.Equal(t, int64(len(testData)), size)
	assert.Equal(t, testData, buf.String())
}

// TestDownloadWithProgress tests download with progress callback
func TestDownloadWithProgress(t *testing.T) {
	testData := "test progress download"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Length", fmt.Sprintf("%d", len(testData)))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(testData))
	}))
	defer server.Close()

	downloader := NewDownloader(&DownloaderConfig{
		HTTPClient: server.Client(),
	})

	progressCalls := 0
	var buf bytes.Buffer

	size, err := downloader.DownloadWithProgress(context.Background(), server.URL, &buf,
		func(downloaded, total int64, percent float64) {
			progressCalls++
			assert.LessOrEqual(t, downloaded, total)
			assert.LessOrEqual(t, percent, 100.0)
		},
	)

	require.NoError(t, err)
	assert.Equal(t, int64(len(testData)), size)
	assert.Greater(t, progressCalls, 0)
}
