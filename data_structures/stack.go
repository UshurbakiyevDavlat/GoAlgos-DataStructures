package data_structures

import "fmt"

type StackNode struct {
	data int
	next *StackNode
}

type Stack struct {
	top *StackNode
}

func (s *Stack) Push(value int) {
	newNode := &StackNode{data: value}
	newNode.next = s.top
	s.top = newNode
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}

	value := s.top.data
	s.top = s.top.next
	return value, nil
}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	return s.top.data, nil
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack) Print() {
	current := s.top

	for current != nil {
		fmt.Printf("%d ->", current.data)
		current = current.next
	}
}
