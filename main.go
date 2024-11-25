package main

import (
	"main/data_structures"
)

func main() {
	stack := data_structures.Stack{}

	stack.Push(10)
	stack.Push(20)

	stack.Pop()

	stack.Print()
}
