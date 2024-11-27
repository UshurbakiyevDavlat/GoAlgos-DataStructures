package data_structures

import "fmt"

type QueueNode struct {
	data int
	next *QueueNode
}

type Queue struct {
	front *QueueNode
	rear  *QueueNode
}

func (q *Queue) Enqueue(value int) {
	newNode := &QueueNode{data: value}

	if q.rear == nil {
		q.front = newNode
		q.rear = newNode
		return
	}

	q.rear.next = newNode
	q.rear = newNode
}

func (q *Queue) Dequeue() (int, error) {
	if q.front == nil {
		return 0, fmt.Errorf("queue is empty")
	}
	value := q.front.data
	q.front = q.front.next

	if q.front == nil {
		q.rear = nil
	}

	return value, nil
}

func (q *Queue) Peek() (int, error) {
	if q.front == nil {
		return 0, fmt.Errorf("queue is empty")
	}

	return q.front.data, nil
}

func (q *Queue) Print() {
	current := q.front

	for current != nil {
		fmt.Printf("%d ->", current.data)
		current = current.next
	}
}
