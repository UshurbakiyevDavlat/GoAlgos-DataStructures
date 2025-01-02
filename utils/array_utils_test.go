package utils

import (
	"testing"
)

func TestFindMinMax(t *testing.T) {
	// Test with a normal array
	arr := []int{3, 5, 1, 9, 7}
	min, max := FindMinMax(arr)
	if min != 1 || max != 9 {
		t.Errorf("For array %v, expected min=1, max=9, got min=%d, max=%d", arr, min, max)
	}

	// Test with a single-element array
	arr = []int{4}
	min, max = FindMinMax(arr)
	if min != 4 || max != 4 {
		t.Errorf("For array %v, expected min=4, max=4, got min=%d, max=%d", arr, min, max)
	}

	// Test with all equal elements
	arr = []int{2, 2, 2, 2}
	min, max = FindMinMax(arr)
	if min != 2 || max != 2 {
		t.Errorf("For array %v, expected min=2, max=2, got min=%d, max=%d", arr, min, max)
	}

	// Test with a large range
	arr = []int{-100, 0, 50, 100}
	min, max = FindMinMax(arr)
	if min != -100 || max != 100 {
		t.Errorf("For array %v, expected min=-100, max=100, got min=%d, max=%d", arr, min, max)
	}

	// Test with an empty array (should panic)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for empty array, but function did not panic")
		}
	}()
	FindMinMax([]int{})
}
