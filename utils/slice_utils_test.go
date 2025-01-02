package utils

import (
	"testing"
)

func TestReverseSlice(t *testing.T) {
	// Test with a normal slice
	slice := []int{1, 2, 3, 4, 5}
	reversed := ReverseSlice(slice)
	expected := []int{5, 4, 3, 2, 1}
	for i, v := range reversed {
		if v != expected[i] {
			t.Errorf("Expected reversed slice to be %v, got %v", expected, reversed)
			break
		}
	}

	// Test with an empty slice
	emptySlice := []int{}
	reversed = ReverseSlice(emptySlice)
	if len(reversed) != 0 {
		t.Errorf("Expected reversed empty slice to be empty, got %v", reversed)
	}

	// Test with a single-element slice
	singleElement := []int{10}
	reversed = ReverseSlice(singleElement)
	if reversed[0] != 10 {
		t.Errorf("Expected reversed single-element slice to be %v, got %v", singleElement, reversed)
	}
}

func TestSplitEvenOdd(t *testing.T) {
	// Test with mixed even and odd numbers
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	oddSlice, evenSlice := SplitEvenOdd(slice)

	expectedOdd := []int{1, 3, 5, 7, 9}
	expectedEven := []int{2, 4, 6, 8}

	// Check odd slice
	for i, v := range oddSlice {
		if v != expectedOdd[i] {
			t.Errorf("Expected odd slice %v, got %v", expectedOdd, oddSlice)
			break
		}
	}

	// Check even slice
	for i, v := range evenSlice {
		if v != expectedEven[i] {
			t.Errorf("Expected even slice %v, got %v", expectedEven, evenSlice)
			break
		}
	}

	// Test with an empty slice
	emptySlice := []int{}
	oddSlice, evenSlice = SplitEvenOdd(emptySlice)
	if len(oddSlice) != 0 || len(evenSlice) != 0 {
		t.Errorf("Expected empty slices for empty input, got odd: %v, even: %v", oddSlice, evenSlice)
	}

	// Test with only odd numbers
	oddOnly := []int{1, 3, 5, 7}
	oddSlice, evenSlice = SplitEvenOdd(oddOnly)
	if len(evenSlice) != 0 {
		t.Errorf("Expected empty even slice, got %v", evenSlice)
	}
	if len(oddSlice) != len(oddOnly) {
		t.Errorf("Expected odd slice %v, got %v", oddOnly, oddSlice)
	}

	// Test with only even numbers
	evenOnly := []int{2, 4, 6, 8}
	oddSlice, evenSlice = SplitEvenOdd(evenOnly)
	if len(oddSlice) != 0 {
		t.Errorf("Expected empty odd slice, got %v", oddSlice)
	}
	if len(evenSlice) != len(evenOnly) {
		t.Errorf("Expected even slice %v, got %v", evenOnly, evenSlice)
	}
}

func TestRemoveElement(t *testing.T) {
	// Test removal at the beginning
	slice := []int{1, 2, 3, 4, 5}
	result := RemoveElement(slice, 0) // Removes 1
	expected := []int{2, 3, 4, 5}
	if !slicesEqual(result, expected) {
		t.Errorf("Expected slice %v after removing first element, got %v", expected, result)
	}
	slice = result

	// Test removal at the end
	result = RemoveElement(slice, len(slice)-1) // Removes 5
	expected = []int{2, 3, 4}
	if !slicesEqual(result, expected) {
		t.Errorf("Expected slice %v after removing last element, got %v", expected, result)
	}
	slice = result

	// Test removal from the middle
	result = RemoveElement(slice, 1) // Removes 3
	expected = []int{2, 4}
	if !slicesEqual(result, expected) {
		t.Errorf("Expected slice %v after removing middle element, got %v", expected, result)
	}
	slice = result

	// Test out-of-bounds index
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for out-of-bounds index")
		}
	}()
	RemoveElement(slice, len(slice)) // Out of bounds
}
