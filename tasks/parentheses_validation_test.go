package tasks

import (
	"testing"
)

func TestIsValidParentheses(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"({[]})", true},    // Balanced parentheses
		{"({[})", false},    // Mismatched parentheses
		{"", true},          // Empty string
		{"(", false},        // Single open parenthesis
		{")", false},        // Single close parenthesis
		{"((()))", true},    // Nested parentheses
		{"()[]{}", true},    // Different types of parentheses
		{"([)]", false},     // Incorrect nesting
		{"{[()()]}", true},  // Complex balanced case
		{"{[(()]}]", false}, // Complex unbalanced case
	}

	for _, test := range tests {
		result := IsValidParentheses(test.input)
		if result != test.expected {
			t.Errorf("For input '%s', expected %v, got %v", test.input, test.expected, result)
		}
	}
}
