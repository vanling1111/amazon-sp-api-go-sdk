package errors

import (
	"net/http"
	"testing"
)

func TestNewSPAPIError(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		message    string
		wantType   ErrorType
		retryable  bool
	}{
		{
			name:       "rate limit error",
			statusCode: 429,
			message:    "Too many requests",
			wantType:   ErrorTypeRateLimit,
			retryable:  true,
		},
		{
			name:       "auth error - 401",
			statusCode: 401,
			message:    "Unauthorized",
			wantType:   ErrorTypeAuth,
			retryable:  false,
		},
		{
			name:       "auth error - 403",
			statusCode: 403,
			message:    "Forbidden",
			wantType:   ErrorTypeAuth,
			retryable:  false,
		},
		{
			name:       "validation error",
			statusCode: 400,
			message:    "Bad request",
			wantType:   ErrorTypeValidation,
			retryable:  false,
		},
		{
			name:       "not found error",
			statusCode: 404,
			message:    "Not found",
			wantType:   ErrorTypeNotFound,
			retryable:  false,
		},
		{
			name:       "server error - 500",
			statusCode: 500,
			message:    "Internal server error",
			wantType:   ErrorTypeServer,
			retryable:  true,
		},
		{
			name:       "server error - 503",
			statusCode: 503,
			message:    "Service unavailable",
			wantType:   ErrorTypeServer,
			retryable:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NewSPAPIError(tt.statusCode, tt.message)

			if err.StatusCode != tt.statusCode {
				t.Errorf("StatusCode = %v, want %v", err.StatusCode, tt.statusCode)
			}

			if err.Message != tt.message {
				t.Errorf("Message = %v, want %v", err.Message, tt.message)
			}

			if err.Type != tt.wantType {
				t.Errorf("Type = %v, want %v", err.Type, tt.wantType)
			}

			if err.Retryable != tt.retryable {
				t.Errorf("Retryable = %v, want %v", err.Retryable, tt.retryable)
			}
		})
	}
}

func TestNewSPAPIErrorFromResponse(t *testing.T) {
	resp := &http.Response{
		StatusCode: 429,
		Header:     http.Header{},
	}
	resp.Header.Set("x-amzn-requestid", "test-request-id")
	resp.Header.Set("x-amzn-ratelimit-limit", "0.5")

	err := NewSPAPIErrorFromResponse(resp, "Rate limited")

	if err.StatusCode != 429 {
		t.Errorf("StatusCode = %v, want 429", err.StatusCode)
	}

	if err.RequestID != "test-request-id" {
		t.Errorf("RequestID = %v, want 'test-request-id'", err.RequestID)
	}

	if err.Type != ErrorTypeRateLimit {
		t.Errorf("Type = %v, want ErrorTypeRateLimit", err.Type)
	}

	if rateLimit, ok := err.Details["rateLimit"]; !ok || rateLimit != "0.5" {
		t.Errorf("Details[rateLimit] = %v, want '0.5'", rateLimit)
	}
}

func TestSPAPIError_Error(t *testing.T) {
	tests := []struct {
		name      string
		err       *SPAPIError
		wantMatch string
	}{
		{
			name: "with request ID",
			err: &SPAPIError{
				Type:       ErrorTypeRateLimit,
				Message:    "Too many requests",
				StatusCode: 429,
				RequestID:  "req-123",
			},
			wantMatch: "request ID: req-123",
		},
		{
			name: "without request ID",
			err: &SPAPIError{
				Type:       ErrorTypeValidation,
				Message:    "Invalid input",
				StatusCode: 400,
			},
			wantMatch: "status: 400",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errMsg := tt.err.Error()
			if errMsg == "" {
				t.Error("Error() returned empty string")
			}
		})
	}
}

func TestSPAPIError_WithMethods(t *testing.T) {
	err := NewSPAPIError(400, "Test error")

	// Test WithErrorCode
	err.WithErrorCode("TEST_ERROR")
	if err.ErrorCode != "TEST_ERROR" {
		t.Errorf("ErrorCode = %v, want 'TEST_ERROR'", err.ErrorCode)
	}

	// Test WithRequestID
	err.WithRequestID("req-456")
	if err.RequestID != "req-456" {
		t.Errorf("RequestID = %v, want 'req-456'", err.RequestID)
	}

	// Test WithDetail
	err.WithDetail("key", "value")
	if val, ok := err.Details["key"]; !ok || val != "value" {
		t.Errorf("Details[key] = %v, want 'value'", val)
	}
}

func TestIsRateLimitError(t *testing.T) {
	rateLimitErr := NewSPAPIError(429, "Rate limited")
	otherErr := NewSPAPIError(400, "Bad request")

	if !IsRateLimitError(rateLimitErr) {
		t.Error("IsRateLimitError() should return true for rate limit error")
	}

	if IsRateLimitError(otherErr) {
		t.Error("IsRateLimitError() should return false for non-rate-limit error")
	}

	// Test with non-SPAPIError
	if IsRateLimitError(nil) {
		t.Error("IsRateLimitError() should return false for nil")
	}
}

func TestIsAuthError(t *testing.T) {
	authErr := NewSPAPIError(401, "Unauthorized")
	otherErr := NewSPAPIError(400, "Bad request")

	if !IsAuthError(authErr) {
		t.Error("IsAuthError() should return true for auth error")
	}

	if IsAuthError(otherErr) {
		t.Error("IsAuthError() should return false for non-auth error")
	}
}

func TestIsValidationError(t *testing.T) {
	validationErr := NewSPAPIError(400, "Validation failed")
	otherErr := NewSPAPIError(500, "Server error")

	if !IsValidationError(validationErr) {
		t.Error("IsValidationError() should return true for validation error")
	}

	if IsValidationError(otherErr) {
		t.Error("IsValidationError() should return false for non-validation error")
	}
}

func TestIsServerError(t *testing.T) {
	serverErr := NewSPAPIError(500, "Internal server error")
	otherErr := NewSPAPIError(400, "Bad request")

	if !IsServerError(serverErr) {
		t.Error("IsServerError() should return true for server error")
	}

	if IsServerError(otherErr) {
		t.Error("IsServerError() should return false for non-server error")
	}
}

func TestIsRetryable(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		want       bool
	}{
		{
			name:       "429 - retryable",
			statusCode: 429,
			want:       true,
		},
		{
			name:       "500 - retryable",
			statusCode: 500,
			want:       true,
		},
		{
			name:       "503 - retryable",
			statusCode: 503,
			want:       true,
		},
		{
			name:       "400 - not retryable",
			statusCode: 400,
			want:       false,
		},
		{
			name:       "401 - not retryable",
			statusCode: 401,
			want:       false,
		},
		{
			name:       "404 - not retryable",
			statusCode: 404,
			want:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NewSPAPIError(tt.statusCode, "Test error")
			if got := IsRetryable(err); got != tt.want {
				t.Errorf("IsRetryable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsRetryable_NilError(t *testing.T) {
	if IsRetryable(nil) {
		t.Error("IsRetryable() should return false for nil error")
	}
}

func TestSPAPIError_IsRetryable(t *testing.T) {
	err := NewSPAPIError(429, "Rate limited")
	if !err.IsRetryable() {
		t.Error("IsRetryable() method should return true for rate limit error")
	}

	err = NewSPAPIError(400, "Bad request")
	if err.IsRetryable() {
		t.Error("IsRetryable() method should return false for bad request error")
	}
}
