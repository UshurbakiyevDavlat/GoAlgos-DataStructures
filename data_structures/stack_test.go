package data_structures

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := Stack[int]{}

	// Test IsEmpty on empty stack
	if !stack.IsEmpty() {
		t.Errorf("Expected stack to be empty, but it is not")
	}

	// Test Push and Peek
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	top, err := stack.Peek()
	if err != nil {
		t.Fatalf("Unexpected error when peeking: %v", err)
	}
	if top != 30 {
		t.Errorf("Expected top to be 30, got %v", top)
	}

	// Test Pop
	value, err := stack.Pop()
	if err != nil {
		t.Fatalf("Unexpected error when popping: %v", err)
	}
	if value != 30 {
		t.Errorf("Expected popped value to be 30, got %v", value)
	}

	// Test Pop again
	value, err = stack.Pop()
	if err != nil {
		t.Fatalf("Unexpected error when popping: %v", err)
	}
	if value != 20 {
		t.Errorf("Expected popped value to be 20, got %v", value)
	}

	// Test Pop for last element
	value, err = stack.Pop()
	if err != nil {
		t.Fatalf("Unexpected error when popping: %v", err)
	}
	if value != 10 {
		t.Errorf("Expected popped value to be 10, got %v", value)
	}

	// Test Pop on empty stack
	_, err = stack.Pop()
	if err == nil {
		t.Errorf("Expected error when popping from an empty stack, got nil")
	}

	// Test Peek on empty stack
	_, err = stack.Peek()
	if err == nil {
		t.Errorf("Expected error when peeking into an empty stack, got nil")
	}

	// Test IsEmpty after all pops
	if !stack.IsEmpty() {
		t.Errorf("Expected stack to be empty after popping all elements, but it is not")
	}
}
