package data_structures

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) InsertAtBeginning(value int) {
	newNode := &Node{data: value}
	newNode.next = l.head
	l.head = newNode
}

func (l *LinkedList) InsertAtEnd(value int) {
	newNode := &Node{data: value}
	if l.head == nil {
		l.head = newNode
		return
	}

	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (l *LinkedList) Remove(value int) {
	if l.head == nil {
		fmt.Println("List is empty")
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		return
	}

	current := l.head
	for current.next != nil {
		if current.next.data == value {
			current.next = current.next.next
			return
		}
		current = current.next
	}
	fmt.Println("No element found")
}

func (l *LinkedList) Reverse() {
	var prev *Node
	current := l.head

	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}
	l.head = prev
}

func (l *LinkedList) PrintList() {
	current := l.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}
