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
package auth

import (
	"errors"
	"testing"
)

func TestNewCredentials(t *testing.T) {
	tests := []struct {
		name         string
		clientID     string
		clientSecret string
		refreshToken string
		endpoint     string
		wantErr      error
	}{
		{
			name:         "valid credentials",
			clientID:     "test-client-id",
			clientSecret: "test-client-secret",
			refreshToken: "test-refresh-token",
			endpoint:     EndpointNA,
			wantErr:      nil,
		},
		{
			name:         "missing client ID",
			clientID:     "",
			clientSecret: "test-client-secret",
			refreshToken: "test-refresh-token",
			endpoint:     EndpointNA,
			wantErr:      ErrMissingClientID,
		},
		{
			name:         "missing client secret",
			clientID:     "test-client-id",
			clientSecret: "",
			refreshToken: "test-refresh-token",
			endpoint:     EndpointNA,
			wantErr:      ErrMissingClientSecret,
		},
		{
			name:         "missing refresh token",
			clientID:     "test-client-id",
			clientSecret: "test-client-secret",
			refreshToken: "",
			endpoint:     EndpointNA,
			wantErr:      ErrInvalidCredentials,
		},
		{
			name:         "missing endpoint",
			clientID:     "test-client-id",
			clientSecret: "test-client-secret",
			refreshToken: "test-refresh-token",
			endpoint:     "",
			wantErr:      ErrMissingEndpoint,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			creds, err := NewCredentials(
				tt.clientID,
				tt.clientSecret,
				tt.refreshToken,
				tt.endpoint,
			)

			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("NewCredentials() error = nil, want %v", tt.wantErr)
					return
				}
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("NewCredentials() error = %v, want %v", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Errorf("NewCredentials() unexpected error = %v", err)
				return
			}

			if creds == nil {
				t.Error("NewCredentials() returned nil credentials")
				return
			}

			if creds.ClientID != tt.clientID {
				t.Errorf("ClientID = %v, want %v", creds.ClientID, tt.clientID)
			}
			if creds.ClientSecret != tt.clientSecret {
				t.Errorf("ClientSecret = %v, want %v", creds.ClientSecret, tt.clientSecret)
			}
			if creds.RefreshToken != tt.refreshToken {
				t.Errorf("RefreshToken = %v, want %v", creds.RefreshToken, tt.refreshToken)
			}
			if creds.Endpoint != tt.endpoint {
				t.Errorf("Endpoint = %v, want %v", creds.Endpoint, tt.endpoint)
			}
		})
	}
}

func TestCredentials_Validate(t *testing.T) {
	tests := []struct {
		name    string
		creds   *Credentials
		wantErr error
	}{
		{
			name: "valid credentials",
			creds: &Credentials{
				ClientID:     "test-id",
				ClientSecret: "test-secret",
				RefreshToken: "test-token",
				Endpoint:     EndpointNA,
			},
			wantErr: nil,
		},
		{
			name: "missing client ID",
			creds: &Credentials{
				ClientID:     "",
				ClientSecret: "test-secret",
				RefreshToken: "test-token",
				Endpoint:     EndpointNA,
			},
			wantErr: ErrMissingClientID,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.creds.Validate()

			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("Validate() error = nil, want %v", tt.wantErr)
					return
				}
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("Validate() error = %v, want %v", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Errorf("Validate() unexpected error = %v", err)
			}
		})
	}
}

func TestCredentials_String(t *testing.T) {
	creds := &Credentials{
		ClientID:     "test-client-id-1234567890",
		ClientSecret: "test-client-secret",
		RefreshToken: "test-refresh-token",
		Endpoint:     EndpointNA,
	}

	str := creds.String()

	// 应该包含端点
	if str == "" {
		t.Error("String() returned empty string")
	}

	// 不应该包含完整的 client secret 或 refresh token
	if len(str) > 0 && !contains(str, "****") {
		t.Error("String() should mask sensitive information")
	}
}

func TestMaskString(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "short string",
			input: "short",
			want:  "****",
		},
		{
			name:  "long string",
			input: "1234567890abcdef",
			want:  "1234****cdef",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maskString(tt.input)
			if got != tt.want {
				t.Errorf("maskString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewGrantlessCredentials(t *testing.T) {
	tests := []struct {
		name         string
		clientID     string
		clientSecret string
		scopes       []string
		endpoint     string
		wantErr      error
	}{
		{
			name:         "valid grantless credentials",
			clientID:     "test-client-id",
			clientSecret: "test-client-secret",
			scopes:       []string{ScopeNotifications},
			endpoint:     EndpointNA,
			wantErr:      nil,
		},
		{
			name:         "multiple scopes",
			clientID:     "test-client-id",
			clientSecret: "test-client-secret",
			scopes:       []string{ScopeNotifications, ScopeCredentialRotation},
			endpoint:     EndpointNA,
			wantErr:      nil,
		},
		{
			name:         "missing scopes",
			clientID:     "test-client-id",
			clientSecret: "test-client-secret",
			scopes:       []string{},
			endpoint:     EndpointNA,
			wantErr:      ErrInvalidCredentials,
		},
		{
			name:         "missing client ID",
			clientID:     "",
			clientSecret: "test-client-secret",
			scopes:       []string{ScopeNotifications},
			endpoint:     EndpointNA,
			wantErr:      ErrMissingClientID,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			creds, err := NewGrantlessCredentials(
				tt.clientID,
				tt.clientSecret,
				tt.scopes,
				tt.endpoint,
			)

			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("NewGrantlessCredentials() error = nil, want %v", tt.wantErr)
					return
				}
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("NewGrantlessCredentials() error = %v, want %v", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Errorf("NewGrantlessCredentials() unexpected error = %v", err)
				return
			}

			if creds == nil {
				t.Error("NewGrantlessCredentials() returned nil credentials")
				return
			}

			if !creds.IsGrantless() {
				t.Error("NewGrantlessCredentials() IsGrantless() = false, want true")
			}
		})
	}
}

func TestCredentials_IsGrantless(t *testing.T) {
	tests := []struct {
		name  string
		creds *Credentials
		want  bool
	}{
		{
			name: "regular credentials",
			creds: &Credentials{
				ClientID:     "test-client-id",
				ClientSecret: "test-client-secret",
				RefreshToken: "test-refresh-token",
				Endpoint:     EndpointNA,
			},
			want: false,
		},
		{
			name: "grantless credentials",
			creds: &Credentials{
				ClientID:     "test-client-id",
				ClientSecret: "test-client-secret",
				Scopes:       []string{ScopeNotifications},
				Endpoint:     EndpointNA,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.creds.IsGrantless(); got != tt.want {
				t.Errorf("IsGrantless() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCredentials_Validate_BothRefreshTokenAndScopes(t *testing.T) {
	creds := &Credentials{
		ClientID:     "test-client-id",
		ClientSecret: "test-client-secret",
		RefreshToken: "test-refresh-token",
		Scopes:       []string{ScopeNotifications},
		Endpoint:     EndpointNA,
	}

	err := creds.Validate()
	if !errors.Is(err, ErrBothRefreshTokenAndScopes) {
		t.Errorf("Validate() error = %v, want %v", err, ErrBothRefreshTokenAndScopes)
	}
}

// 辅助函数
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && (s[:len(substr)] == substr || s[len(s)-len(substr):] == substr || containsSubstring(s, substr)))
}

func containsSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
