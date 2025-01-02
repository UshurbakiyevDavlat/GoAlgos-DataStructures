package data_structures

import (
	"testing"
)

func TestQueue(t *testing.T) {
	queue := Queue{}

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

	// Test Reverse
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	queue.Reverse()

	value, err = queue.Dequeue()
	if value != 3 {
		t.Errorf("Expected dequeued value to be 3 after reverse, got %d", value)
	}
	value, err = queue.Dequeue()
	if value != 2 {
		t.Errorf("Expected dequeued value to be 2 after reverse, got %d", value)
	}
	value, err = queue.Dequeue()
	if value != 1 {
		t.Errorf("Expected dequeued value to be 1 after reverse, got %d", value)
	}

	// Test CompareQueues
	q1 := &Queue{}
	q2 := &Queue{}

	if !CompareQueues(q1, q2) {
		t.Errorf("Expected two empty queues to be equal")
	}

	q1.Enqueue(10)
	q1.Enqueue(20)
	q2.Enqueue(10)
	q2.Enqueue(20)

	if !CompareQueues(q1, q2) {
		t.Errorf("Expected queues with same elements to be equal")
	}

	q2.Enqueue(30)
	if CompareQueues(q1, q2) {
		t.Errorf("Expected queues with different elements to be unequal")
	}
}
