package tasks

import "testing"

func TestLargestRectangleArea(t *testing.T) {
	tests := []struct {
		heights  []int
		expected int
	}{
		{[]int{2, 1, 5, 6, 2, 3}, 10},    // Standard case
		{[]int{2, 4}, 4},                 // Small array
		{[]int{1, 1, 1, 1}, 4},           // All heights equal
		{[]int{}, 0},                     // Empty array
		{[]int{1}, 1},                    // Single bar
		{[]int{6, 2, 5, 4, 5, 1, 6}, 12}, // Complex case
	}

	for _, test := range tests {
		result := LargestRectangleArea(test.heights)
		if result != test.expected {
			t.Errorf("For heights %v, expected %d, got %d", test.heights, test.expected, result)
		}
	}
}
