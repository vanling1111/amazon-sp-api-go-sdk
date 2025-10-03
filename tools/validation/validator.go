// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

// Package validation provides request/response validation utilities
package validation

import (
	"fmt"
	"regexp"
)

// Rule represents a validation rule
type Rule struct {
	Field   string
	Rule    string
	Message string
}

// Validator validates requests and responses
type Validator struct {
	rules []Rule
}

// NewValidator creates a new validator
func NewValidator() *Validator {
	return &Validator{
		rules: make([]Rule, 0),
	}
}

// AddRule adds a validation rule
func (v *Validator) AddRule(field, rule, message string) {
	v.rules = append(v.rules, Rule{
		Field:   field,
		Rule:    rule,
		Message: message,
	})
}

// Validate validates a value against all rules
func (v *Validator) Validate(data map[string]interface{}) []error {
	var errors []error

	for _, rule := range v.rules {
		value, exists := data[rule.Field]
		if !exists {
			if rule.Rule == "required" {
				errors = append(errors, fmt.Errorf("%s: %s", rule.Field, rule.Message))
			}
			continue
		}

		switch rule.Rule {
		case "required":
			if value == nil || value == "" {
				errors = append(errors, fmt.Errorf("%s: %s", rule.Field, rule.Message))
			}
		case "email":
			if str, ok := value.(string); ok {
				matched, _ := regexp.MatchString(`^[^\s@]+@[^\s@]+\.[^\s@]+$`, str)
				if !matched {
					errors = append(errors, fmt.Errorf("%s: %s", rule.Field, rule.Message))
				}
			}
		}
	}

	return errors
}

