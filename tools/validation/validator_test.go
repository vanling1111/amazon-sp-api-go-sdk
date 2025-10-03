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
// terms of the applicable license.
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
// terms of the applicable license.

package validation

import (
	"testing"
)

func TestNewValidator(t *testing.T) {
	validator := NewValidator()
	if validator == nil {
		t.Error("NewValidator returned nil")
	}
	if validator.rules == nil {
		t.Error("Validator rules not initialized")
	}
}

func TestValidator_AddRule(t *testing.T) {
	validator := NewValidator()
	validator.AddRule("email", "required", "Email is required")
	validator.AddRule("email", "email", "Invalid email format")

	if len(validator.rules) != 2 {
		t.Errorf("Expected 2 rules, got %d", len(validator.rules))
	}
}

func TestValidator_Validate_Required(t *testing.T) {
	validator := NewValidator()
	validator.AddRule("name", "required", "Name is required")

	tests := []struct {
		name    string
		data    map[string]interface{}
		wantErr bool
	}{
		{
			name:    "missing field",
			data:    map[string]interface{}{},
			wantErr: true,
		},
		{
			name:    "empty value",
			data:    map[string]interface{}{"name": ""},
			wantErr: true,
		},
		{
			name:    "valid value",
			data:    map[string]interface{}{"name": "John"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs := validator.Validate(tt.data)
			if (len(errs) > 0) != tt.wantErr {
				t.Errorf("Validate() errors = %v, wantErr %v", errs, tt.wantErr)
			}
		})
	}
}

func TestValidator_Validate_Email(t *testing.T) {
	validator := NewValidator()
	validator.AddRule("email", "email", "Invalid email")

	tests := []struct {
		name    string
		email   string
		wantErr bool
	}{
		{"valid email", "test@example.com", false},
		{"invalid email", "invalid", true},
		{"missing @", "testexample.com", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := map[string]interface{}{"email": tt.email}
			errs := validator.Validate(data)
			if (len(errs) > 0) != tt.wantErr {
				t.Errorf("Validate() errors = %v, wantErr %v", errs, tt.wantErr)
			}
		})
	}
}
