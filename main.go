package main

import (
	"fmt"
	"main/data_structures"
)

func main() {
	list := data_structures.DoublyLinkedList{}
	list.InsertAtEnd(20)
	list.InsertAtEnd(30)
	list.InsertAtBeginning(10)

	fmt.Println("Forward:")
	list.PrintForward()

	fmt.Println("Backward:")
	list.PrintBackward()
}
