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
