package data_structures

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	list := LinkedList{}

	// Test InsertAtBeginning
	list.InsertAtBeginning(10)
	list.InsertAtBeginning(20)
	list.InsertAtBeginning(30)
	expectedHead := 30
	if list.head.data != expectedHead {
		t.Errorf("Expected head to be %d, got %d", expectedHead, list.head.data)
	}

	// Test InsertAtEnd
	list.InsertAtEnd(40)
	list.InsertAtEnd(50)

	current := list.head
	values := []int{30, 20, 10, 40, 50}
	for _, val := range values {
		if current == nil || current.data != val {
			t.Errorf("Expected value %d in the list, got %d", val, current.data)
		}
		current = current.next
	}

	// Test Remove (head removal)
	list.Remove(30)
	if list.head.data != 20 {
		t.Errorf("Expected head to be 20 after removal, got %d", list.head.data)
	}

	// Test Remove (middle element)
	list.Remove(40)
	current = list.head
	values = []int{20, 10, 50}
	for _, val := range values {
		if current == nil || current.data != val {
			t.Errorf("Expected value %d in the list after removal, got %d", val, current.data)
		}
		current = current.next
	}

	// Test Remove (non-existent element)
	list.Remove(60)
	current = list.head
	values = []int{20, 10, 50}
	for _, val := range values {
		if current == nil || current.data != val {
			t.Errorf("Expected value %d in the list after attempting to remove non-existent element, got %d", val, current.data)
		}
		current = current.next
	}

	// Test Reverse
	list.Reverse()
	current = list.head
	values = []int{50, 10, 20} // List should be reversed
	for _, val := range values {
		if current == nil || current.data != val {
			t.Errorf("Expected value %d in the reversed list, got %d", val, current.data)
		}
		current = current.next
	}
}
