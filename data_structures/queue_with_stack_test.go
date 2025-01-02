package data_structures

import (
	"testing"
)

func TestQueueWithStacks(t *testing.T) {
	queue := QueueWithStacks[int]{}

	// Test Dequeue on an empty queue
	_, err := queue.Dequeue()
	if err == nil {
		t.Errorf("Expected error when dequeuing from an empty queue, got nil")
	}

	// Test Peek on an empty queue
	_, err = queue.Peek()
	if err == nil {
		t.Errorf("Expected error when peeking into an empty queue, got nil")
	}

	// Test Enqueue
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	// Test Peek
	value, err := queue.Peek()
	if err != nil {
		t.Fatalf("Unexpected error when peeking: %v", err)
	}
	if value != 10 {
		t.Errorf("Expected front to be 10, got %d", value)
	}

	// Test Dequeue
	value, err = queue.Dequeue()
	if err != nil {
		t.Fatalf("Unexpected error when dequeuing: %v", err)
	}
	if value != 10 {
		t.Errorf("Expected dequeued value to be 10, got %d", value)
	}

	// Test remaining elements
	value, err = queue.Dequeue()
	if err != nil {
		t.Fatalf("Unexpected error when dequeuing: %v", err)
	}
	if value != 20 {
		t.Errorf("Expected dequeued value to be 20, got %d", value)
	}

	value, err = queue.Dequeue()
	if err != nil {
		t.Fatalf("Unexpected error when dequeuing: %v", err)
	}
	if value != 30 {
		t.Errorf("Expected dequeued value to be 30, got %d", value)
	}

	// Test Dequeue on an empty queue again
	_, err = queue.Dequeue()
	if err == nil {
		t.Errorf("Expected error when dequeuing from an empty queue, got nil")
	}

	// Test Peek on an empty queue again
	_, err = queue.Peek()
	if err == nil {
		t.Errorf("Expected error when peeking into an empty queue, got nil")
	}

	// Test Enqueue after Dequeue
	queue.Enqueue(40)
	value, err = queue.Peek()
	if err != nil {
		t.Fatalf("Unexpected error when peeking: %v", err)
	}
	if value != 40 {
		t.Errorf("Expected front to be 40, got %d", value)
	}
	value, err = queue.Dequeue()
	if err != nil {
		t.Fatalf("Unexpected error when dequeuing: %v", err)
	}
	if value != 40 {
		t.Errorf("Expected dequeued value to be 40, got %d", value)
	}
}
