package data_structures

import "fmt"

type QueueWithStacks[T any] struct {
	stack1 Stack[T]
	stack2 Stack[T]
}

func (q *QueueWithStacks[T]) Enqueue(value T) {
	q.stack1.Push(value)
}

func (q *QueueWithStacks[T]) Dequeue() (T, error) {
	if q.stack2.IsEmpty() {
		for !q.stack1.IsEmpty() {
			value, _ := q.stack1.Pop()
			q.stack2.Push(value)
		}
	}

	if q.stack2.IsEmpty() {
		var zeroValue T
		return zeroValue, fmt.Errorf("queue is empty")
	}

	return q.stack2.Pop()
}

func (q *QueueWithStacks[T]) Peek() (T, error) {
	if q.stack2.IsEmpty() {
		for !q.stack1.IsEmpty() {
			value, _ := q.stack1.Pop()
			q.stack2.Push(value)
		}
	}

	if q.stack2.IsEmpty() {
		var zeroValue T
		return zeroValue, fmt.Errorf("queue is empty")
	}

	return q.stack2.Peek()
}
