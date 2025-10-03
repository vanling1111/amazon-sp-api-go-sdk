package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBuildURL(t *testing.T) {
	tests := []struct {
		name    string
		baseURL string
		path    string
		want    string
	}{
		{
			name:    "normal case",
			baseURL: "https://sellingpartnerapi-na.amazon.com",
			path:    "/orders/v0/orders",
			want:    "https://sellingpartnerapi-na.amazon.com/orders/v0/orders",
		},
		{
			name:    "base URL with trailing slash",
			baseURL: "https://sellingpartnerapi-na.amazon.com/",
			path:    "/orders/v0/orders",
			want:    "https://sellingpartnerapi-na.amazon.com/orders/v0/orders",
		},
		{
			name:    "path without leading slash",
			baseURL: "https://sellingpartnerapi-na.amazon.com",
			path:    "orders/v0/orders",
			want:    "https://sellingpartnerapi-na.amazon.com/orders/v0/orders",
		},
		{
			name:    "both with slashes",
			baseURL: "https://sellingpartnerapi-na.amazon.com/",
			path:    "/orders/v0/orders",
			want:    "https://sellingpartnerapi-na.amazon.com/orders/v0/orders",
		},
		{
			name:    "root path",
			baseURL: "https://sellingpartnerapi-na.amazon.com",
			path:    "/",
			want:    "https://sellingpartnerapi-na.amazon.com/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BuildURL(tt.baseURL, tt.path)
			if got != tt.want {
				t.Errorf("BuildURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetContentType(t *testing.T) {
	tests := []struct {
		name        string
		contentType string
		want        string
	}{
		{
			name:        "JSON content type",
			contentType: "application/json",
			want:        "application/json",
		},
		{
			name:        "JSON with charset",
			contentType: "application/json; charset=utf-8",
			want:        "application/json; charset=utf-8",
		},
		{
			name:        "no content type",
			contentType: "",
			want:        "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &http.Response{
				Header: http.Header{},
			}
			if tt.contentType != "" {
				resp.Header.Set("Content-Type", tt.contentType)
			}

			got := GetContentType(resp)
			if got != tt.want {
				t.Errorf("GetContentType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsJSONResponse(t *testing.T) {
	tests := []struct {
		name        string
		contentType string
		want        bool
	}{
		{
			name:        "JSON content type",
			contentType: "application/json",
			want:        true,
		},
		{
			name:        "JSON with charset",
			contentType: "application/json; charset=utf-8",
			want:        true,
		},
		{
			name:        "uppercase JSON",
			contentType: "APPLICATION/JSON",
			want:        true,
		},
		{
			name:        "text/plain",
			contentType: "text/plain",
			want:        false,
		},
		{
			name:        "no content type",
			contentType: "",
			want:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &http.Response{
				Header: http.Header{},
			}
			if tt.contentType != "" {
				resp.Header.Set("Content-Type", tt.contentType)
			}

			got := IsJSONResponse(resp)
			if got != tt.want {
				t.Errorf("IsJSONResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRequestID(t *testing.T) {
	tests := []struct {
		name      string
		requestID string
		want      string
	}{
		{
			name:      "with request ID",
			requestID: "123e4567-e89b-12d3-a456-426614174000",
			want:      "123e4567-e89b-12d3-a456-426614174000",
		},
		{
			name:      "no request ID",
			requestID: "",
			want:      "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &http.Response{
				Header: http.Header{},
			}
			if tt.requestID != "" {
				resp.Header.Set("x-amzn-requestid", tt.requestID)
			}

			got := GetRequestID(resp)
			if got != tt.want {
				t.Errorf("GetRequestID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRateLimitHeader(t *testing.T) {
	tests := []struct {
		name      string
		rateLimit string
		want      string
	}{
		{
			name:      "with rate limit",
			rateLimit: "0.5",
			want:      "0.5",
		},
		{
			name:      "no rate limit",
			rateLimit: "",
			want:      "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &http.Response{
				Header: http.Header{},
			}
			if tt.rateLimit != "" {
				resp.Header.Set("x-amzn-ratelimit-limit", tt.rateLimit)
			}

			got := GetRateLimitHeader(resp)
			if got != tt.want {
				t.Errorf("GetRateLimitHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatHTTPError(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		status     string
		requestID  string
		want       string
	}{
		{
			name:       "with request ID",
			statusCode: 404,
			status:     "404 Not Found",
			requestID:  "test-request-id",
			want:       "status 404 (404 Not Found), request ID: test-request-id",
		},
		{
			name:       "without request ID",
			statusCode: 500,
			status:     "500 Internal Server Error",
			requestID:  "",
			want:       "status 500 (500 Internal Server Error)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &http.Response{
				StatusCode: tt.statusCode,
				Status:     tt.status,
				Header:     http.Header{},
			}
			if tt.requestID != "" {
				resp.Header.Set("x-amzn-requestid", tt.requestID)
			}

			got := FormatHTTPError(resp)
			if got != tt.want {
				t.Errorf("FormatHTTPError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsRetryableStatusCode(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		want       bool
	}{
		{
			name:       "429 Too Many Requests",
			statusCode: 429,
			want:       true,
		},
		{
			name:       "500 Internal Server Error",
			statusCode: 500,
			want:       true,
		},
		{
			name:       "502 Bad Gateway",
			statusCode: 502,
			want:       true,
		},
		{
			name:       "503 Service Unavailable",
			statusCode: 503,
			want:       true,
		},
		{
			name:       "504 Gateway Timeout",
			statusCode: 504,
			want:       true,
		},
		{
			name:       "200 OK",
			statusCode: 200,
			want:       false,
		},
		{
			name:       "400 Bad Request",
			statusCode: 400,
			want:       false,
		},
		{
			name:       "401 Unauthorized",
			statusCode: 401,
			want:       false,
		},
		{
			name:       "404 Not Found",
			statusCode: 404,
			want:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsRetryableStatusCode(tt.statusCode)
			if got != tt.want {
				t.Errorf("IsRetryableStatusCode(%d) = %v, want %v", tt.statusCode, got, tt.want)
			}
		})
	}
}

func TestGetRequestID_Integration(t *testing.T) {
	// 创建一个测试服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("x-amzn-requestid", "integration-test-id")
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// 发送请求
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("http.Get() error = %v", err)
	}
	defer resp.Body.Close()

	// 验证请求 ID
	requestID := GetRequestID(resp)
	if requestID != "integration-test-id" {
		t.Errorf("GetRequestID() = %v, want 'integration-test-id'", requestID)
	}
}
