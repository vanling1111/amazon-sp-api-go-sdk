package codec

import (
	"strings"
	"testing"
)

func TestNewValidator(t *testing.T) {
	validator := NewValidator()
	if validator == nil {
		t.Error("NewValidator() returned nil")
	}

	if validator.HasErrors() {
		t.Error("New validator should not have errors")
	}
}

func TestValidator_Required(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{
			name:    "non-empty value",
			value:   "test",
			wantErr: false,
		},
		{
			name:    "empty value",
			value:   "",
			wantErr: true,
		},
		{
			name:    "whitespace only",
			value:   "   ",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := NewValidator()
			validator.Required("testField", tt.value)

			if validator.HasErrors() != tt.wantErr {
				t.Errorf("Required() hasError = %v, want %v", validator.HasErrors(), tt.wantErr)
			}
		})
	}
}

func TestValidator_MinLength(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		min     int
		wantErr bool
	}{
		{
			name:    "sufficient length",
			value:   "hello",
			min:     3,
			wantErr: false,
		},
		{
			name:    "exact length",
			value:   "hello",
			min:     5,
			wantErr: false,
		},
		{
			name:    "insufficient length",
			value:   "hi",
			min:     5,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := NewValidator()
			validator.MinLength("testField", tt.value, tt.min)

			if validator.HasErrors() != tt.wantErr {
				t.Errorf("MinLength() hasError = %v, want %v", validator.HasErrors(), tt.wantErr)
			}
		})
	}
}

func TestValidator_MaxLength(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		max     int
		wantErr bool
	}{
		{
			name:    "within limit",
			value:   "hello",
			max:     10,
			wantErr: false,
		},
		{
			name:    "exact limit",
			value:   "hello",
			max:     5,
			wantErr: false,
		},
		{
			name:    "exceeds limit",
			value:   "hello world",
			max:     5,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := NewValidator()
			validator.MaxLength("testField", tt.value, tt.max)

			if validator.HasErrors() != tt.wantErr {
				t.Errorf("MaxLength() hasError = %v, want %v", validator.HasErrors(), tt.wantErr)
			}
		})
	}
}

func TestValidator_Min(t *testing.T) {
	tests := []struct {
		name    string
		value   int
		min     int
		wantErr bool
	}{
		{
			name:    "above minimum",
			value:   10,
			min:     5,
			wantErr: false,
		},
		{
			name:    "at minimum",
			value:   5,
			min:     5,
			wantErr: false,
		},
		{
			name:    "below minimum",
			value:   3,
			min:     5,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := NewValidator()
			validator.Min("testField", tt.value, tt.min)

			if validator.HasErrors() != tt.wantErr {
				t.Errorf("Min() hasError = %v, want %v", validator.HasErrors(), tt.wantErr)
			}
		})
	}
}

func TestValidator_Max(t *testing.T) {
	tests := []struct {
		name    string
		value   int
		max     int
		wantErr bool
	}{
		{
			name:    "below maximum",
			value:   5,
			max:     10,
			wantErr: false,
		},
		{
			name:    "at maximum",
			value:   10,
			max:     10,
			wantErr: false,
		},
		{
			name:    "above maximum",
			value:   15,
			max:     10,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := NewValidator()
			validator.Max("testField", tt.value, tt.max)

			if validator.HasErrors() != tt.wantErr {
				t.Errorf("Max() hasError = %v, want %v", validator.HasErrors(), tt.wantErr)
			}
		})
	}
}

func TestValidator_Range(t *testing.T) {
	tests := []struct {
		name    string
		value   int
		min     int
		max     int
		wantErr bool
	}{
		{
			name:    "within range",
			value:   5,
			min:     1,
			max:     10,
			wantErr: false,
		},
		{
			name:    "at minimum",
			value:   1,
			min:     1,
			max:     10,
			wantErr: false,
		},
		{
			name:    "at maximum",
			value:   10,
			min:     1,
			max:     10,
			wantErr: false,
		},
		{
			name:    "below minimum",
			value:   0,
			min:     1,
			max:     10,
			wantErr: true,
		},
		{
			name:    "above maximum",
			value:   11,
			min:     1,
			max:     10,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := NewValidator()
			validator.Range("testField", tt.value, tt.min, tt.max)

			if validator.HasErrors() != tt.wantErr {
				t.Errorf("Range() hasError = %v, want %v", validator.HasErrors(), tt.wantErr)
			}
		})
	}
}

func TestValidator_Email(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{
			name:    "valid email",
			value:   "test@example.com",
			wantErr: false,
		},
		{
			name:    "valid email with subdomain",
			value:   "test@mail.example.com",
			wantErr: false,
		},
		{
			name:    "empty email",
			value:   "",
			wantErr: false, // Empty is allowed, use Required() to enforce non-empty
		},
		{
			name:    "invalid email - missing @",
			value:   "testexample.com",
			wantErr: true,
		},
		{
			name:    "invalid email - missing domain",
			value:   "test@",
			wantErr: true,
		},
		{
			name:    "invalid email - missing user",
			value:   "@example.com",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := NewValidator()
			validator.Email("testField", tt.value)

			if validator.HasErrors() != tt.wantErr {
				t.Errorf("Email() hasError = %v, want %v", validator.HasErrors(), tt.wantErr)
			}
		})
	}
}

func TestValidator_URL(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{
			name:    "valid URL",
			value:   "https://example.com",
			wantErr: false,
		},
		{
			name:    "valid URL with path",
			value:   "https://example.com/path/to/resource",
			wantErr: false,
		},
		{
			name:    "empty URL",
			value:   "",
			wantErr: false, // Empty is allowed
		},
		{
			name:    "invalid URL",
			value:   "not a url",
			wantErr: false, // url.Parse will parse this, just not as expected
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := NewValidator()
			validator.URL("testField", tt.value)

			if validator.HasErrors() != tt.wantErr {
				t.Errorf("URL() hasError = %v, want %v", validator.HasErrors(), tt.wantErr)
			}
		})
	}
}

func TestValidator_Pattern(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		pattern string
		wantErr bool
	}{
		{
			name:    "matching pattern",
			value:   "123-456",
			pattern: `^\d{3}-\d{3}$`,
			wantErr: false,
		},
		{
			name:    "non-matching pattern",
			value:   "abc-def",
			pattern: `^\d{3}-\d{3}$`,
			wantErr: true,
		},
		{
			name:    "empty value",
			value:   "",
			pattern: `^\d{3}-\d{3}$`,
			wantErr: false, // Empty is allowed
		},
		{
			name:    "invalid pattern",
			value:   "test",
			pattern: `[`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := NewValidator()
			validator.Pattern("testField", tt.value, tt.pattern, "must match pattern")

			if validator.HasErrors() != tt.wantErr {
				t.Errorf("Pattern() hasError = %v, want %v", validator.HasErrors(), tt.wantErr)
			}
		})
	}
}

func TestValidator_OneOf(t *testing.T) {
	allowed := []string{"Pending", "Shipped", "Delivered"}

	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{
			name:    "valid value",
			value:   "Shipped",
			wantErr: false,
		},
		{
			name:    "invalid value",
			value:   "Unknown",
			wantErr: true,
		},
		{
			name:    "empty value",
			value:   "",
			wantErr: false, // Empty is allowed
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := NewValidator()
			validator.OneOf("testField", tt.value, allowed)

			if validator.HasErrors() != tt.wantErr {
				t.Errorf("OneOf() hasError = %v, want %v", validator.HasErrors(), tt.wantErr)
			}
		})
	}
}

func TestValidator_Custom(t *testing.T) {
	validator := NewValidator()
	validator.Custom("testField", "custom error message")

	if !validator.HasErrors() {
		t.Error("Custom() should create an error")
	}

	err := validator.Error()
	if err == nil {
		t.Error("Error() should return an error")
	}

	if !strings.Contains(err.Error(), "custom error message") {
		t.Errorf("Error message should contain custom message, got: %v", err)
	}
}

func TestValidator_MultipleErrors(t *testing.T) {
	validator := NewValidator()

	validator.Required("field1", "")
	validator.MinLength("field2", "hi", 5)
	validator.Max("field3", 100, 50)

	if !validator.HasErrors() {
		t.Error("Validator should have errors")
	}

	errors := validator.Errors()
	if len(errors) != 3 {
		t.Errorf("Expected 3 errors, got %d", len(errors))
	}
}

func TestValidator_Clear(t *testing.T) {
	validator := NewValidator()
	validator.Required("testField", "")

	if !validator.HasErrors() {
		t.Error("Validator should have errors before clear")
	}

	validator.Clear()

	if validator.HasErrors() {
		t.Error("Validator should not have errors after clear")
	}
}

func TestValidationError_Error(t *testing.T) {
	tests := []struct {
		name     string
		err      ValidationError
		contains string
	}{
		{
			name: "error with field",
			err: ValidationError{
				Field:   "testField",
				Message: "is required",
			},
			contains: "testField",
		},
		{
			name: "error without field",
			err: ValidationError{
				Field:   "",
				Message: "general error",
			},
			contains: "general error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errMsg := tt.err.Error()
			if !strings.Contains(errMsg, tt.contains) {
				t.Errorf("Error message should contain '%s', got: %v", tt.contains, errMsg)
			}
		})
	}
}

func TestValidationErrors_Error(t *testing.T) {
	tests := []struct {
		name     string
		errors   ValidationErrors
		contains string
	}{
		{
			name:     "no errors",
			errors:   ValidationErrors{},
			contains: "no validation errors",
		},
		{
			name: "single error",
			errors: ValidationErrors{
				{Field: "field1", Message: "is required"},
			},
			contains: "field1",
		},
		{
			name: "multiple errors",
			errors: ValidationErrors{
				{Field: "field1", Message: "is required"},
				{Field: "field2", Message: "is too short"},
			},
			contains: "field1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errMsg := tt.errors.Error()
			if !strings.Contains(errMsg, tt.contains) {
				t.Errorf("Error message should contain '%s', got: %v", tt.contains, errMsg)
			}
		})
	}
}

func TestValidator_CompleteExample(t *testing.T) {
	// 模拟订单验证
	type Order struct {
		OrderID  string
		Status   string
		Quantity int
		Email    string
	}

	order := Order{
		OrderID:  "",
		Status:   "InvalidStatus",
		Quantity: 1001,
		Email:    "invalid-email",
	}

	validator := NewValidator()

	// 验证订单
	validator.Required("orderId", order.OrderID)
	validator.OneOf("status", order.Status, []string{"Pending", "Shipped", "Delivered"})
	validator.Range("quantity", order.Quantity, 1, 1000)
	validator.Email("email", order.Email)

	// 应该有 4 个错误
	if !validator.HasErrors() {
		t.Error("Validator should have errors")
	}

	errors := validator.Errors()
	if len(errors) != 4 {
		t.Errorf("Expected 4 errors, got %d", len(errors))
	}

	// 检查错误消息
	err := validator.Error()
	if err == nil {
		t.Error("Error() should return an error")
	}
}

