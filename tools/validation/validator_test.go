// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

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
