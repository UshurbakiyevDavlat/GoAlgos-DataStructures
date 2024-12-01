package data_structures

import "fmt"

type StackNode[T any] struct {
	data T
	next *StackNode[T]
}

type Stack[T any] struct {
	top *StackNode[T]
}

func (s *Stack[T]) Push(value T) {
	newNode := &StackNode[T]{data: value}
	newNode.next = s.top
	s.top = newNode
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue, fmt.Errorf("stack is empty")
	}

	value := s.top.data
	s.top = s.top.next
	return value, nil
}

func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue, fmt.Errorf("stack is empty")
	}
	return s.top.data, nil
}

func (s *Stack[T]) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack[T]) Print() {
	current := s.top

	for current != nil {
		fmt.Printf("%v ->", current.data)
		current = current.next
	}
}
