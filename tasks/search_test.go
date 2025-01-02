package tasks

import "testing"

func TestBinarySearch(t *testing.T) {
	arr := []int{2, 3, 5, 6, 8} // Assumes the array is already sorted

	tests := []struct {
		target   int
		expected int
	}{
		{5, 2},
		{2, 0},
		{8, 4},
		{10, -1}, // Target not found
	}

	for _, test := range tests {
		result := BinarySearch(arr, test.target)
		if result != test.expected {
			t.Errorf("BinarySearch failed for target %d. Expected %d, got %d", test.target, test.expected, result)
		}
	}
}
