package data_structures

import "fmt"

type DoublyNode struct {
	data int
	next *DoublyNode
	prev *DoublyNode
}

type DoublyLinkedList struct {
	head *DoublyNode
	tail *DoublyNode
}

func (l *DoublyLinkedList) InsertAtEnd(value int) {
	newNode := &DoublyNode{data: value}
	if l.tail == nil {
		l.head = newNode
		l.tail = newNode
		return
	}

	newNode.prev = l.tail
	l.tail.next = newNode
	l.tail = newNode
}

func (l *DoublyLinkedList) InsertAtBeginning(value int) {
	newNode := &DoublyNode{data: value}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}

	newNode.next = l.head
	l.head.prev = newNode
	l.head = newNode
}

func (l *DoublyLinkedList) PrintForward() {
	current := l.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func (l *DoublyLinkedList) PrintBackward() {
	current := l.tail
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.prev
	}
	fmt.Println("nil")
}
