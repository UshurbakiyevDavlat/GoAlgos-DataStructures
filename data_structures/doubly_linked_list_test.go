package data_structures

import (
	"testing"
)

func TestDoublyLinkedList(t *testing.T) {
	list := DoublyLinkedList{}

	// Test InsertAtEnd
	list.InsertAtEnd(10)
	list.InsertAtEnd(20)
	list.InsertAtEnd(30)

	// Validate forward traversal
	expectedForward := []int{10, 20, 30}
	current := list.head
	for i, val := range expectedForward {
		if current == nil || current.data != val {
			t.Errorf("Forward traversal: Expected value %d at position %d, got %d", val, i, current.data)
		}
		current = current.next
	}

	// Validate backward traversal
	expectedBackward := []int{30, 20, 10}
	current = list.tail
	for i, val := range expectedBackward {
		if current == nil || current.data != val {
			t.Errorf("Backward traversal: Expected value %d at position %d, got %d", val, i, current.data)
		}
		current = current.prev
	}

	// Test InsertAtBeginning
	list.InsertAtBeginning(5)
	list.InsertAtBeginning(0)

	// Validate forward traversal after InsertAtBeginning
	expectedForward = []int{0, 5, 10, 20, 30}
	current = list.head
	for i, val := range expectedForward {
		if current == nil || current.data != val {
			t.Errorf("Forward traversal after InsertAtBeginning: Expected value %d at position %d, got %d", val, i, current.data)
		}
		current = current.next
	}

	// Validate backward traversal after InsertAtBeginning
	expectedBackward = []int{30, 20, 10, 5, 0}
	current = list.tail
	for i, val := range expectedBackward {
		if current == nil || current.data != val {
			t.Errorf("Backward traversal after InsertAtBeginning: Expected value %d at position %d, got %d", val, i, current.data)
		}
		current = current.prev
	}
}
