package tasks

import (
	"testing"
)

func TestEvaluateRPN(t *testing.T) {
	tests := []struct {
		expression string
		expected   int
		err        string
	}{
		{"3 4 + 2 * 7 /", 2, ""},                                 // Valid expression
		{"5 1 2 + 4 * + 3 -", 14, ""},                            // Valid expression with multiple operations
		{"2 3 +", 5, ""},                                         // Simple addition
		{"10 5 /", 2, ""},                                        // Simple division
		{"10 0 /", 0, "division by zero"},                        // Division by zero
		{"3 4 + +", 0, "invalid expression: too many arguments"}, // Extra operator
		{"3 4", 0, "invalid expression: too many arguments"},     // Missing operator
		{"3 4 + a", 0, "invalid token: a"},                       // Invalid token
		{"", 0, "invalid expression: invalid token: "},           // Empty input
	}

	for _, test := range tests {
		result, err := EvaluateRPN(test.expression)

		if test.err != "" {
			if err == nil || err.Error() != test.err {
				t.Errorf("For expression '%s', expected error '%s', got '%v'", test.expression, test.err, err)
			}
		} else {
			if err != nil {
				t.Errorf("For expression '%s', unexpected error: %v", test.expression, err)
			} else if result != test.expected {
				t.Errorf("For expression '%s', expected result %d, got %d", test.expression, test.expected, result)
			}
		}
	}
}
