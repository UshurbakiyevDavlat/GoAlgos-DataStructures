package sorting

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{5, 3, 8, 6, 2}
	expected := []int{2, 3, 5, 6, 8}

	BubleSort(arr) // Note: Modifies the array in-place

	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("BubbleSort failed. Expected %v, got %v", expected, arr)
	}
}

func TestSelectionSort(t *testing.T) {
	arr := []int{5, 3, 8, 6, 2}
	expected := []int{2, 3, 5, 6, 8}

	SelectionSort(arr) // Note: Modifies the array in-place

	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("SelectionSort failed. Expected %v, got %v", expected, arr)
	}
}

func TestInsertionSort(t *testing.T) {
	arr := []int{5, 3, 8, 6, 2}
	expected := []int{2, 3, 5, 6, 8}

	InsertionSort(arr) // Note: Modifies the array in-place

	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("InsertionSort failed. Expected %v, got %v", expected, arr)
	}
}

func TestMergeSort(t *testing.T) {
	arr := []int{5, 3, 8, 6, 2}
	expected := []int{2, 3, 5, 6, 8}

	result := MergeSort(arr) // Returns a new sorted array

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("MergeSort failed. Expected %v, got %v", expected, result)
	}
}

func TestQuickSort(t *testing.T) {
	arr := []int{5, 3, 8, 6, 2}
	expected := []int{2, 3, 5, 6, 8}

	result := QuickSort(arr) // Returns a new sorted array

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("QuickSort failed. Expected %v, got %v", expected, result)
	}
}
