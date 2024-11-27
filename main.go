package main

import (
	"main/data_structures"
)

func main() {
	queue := data_structures.Queue{}

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	queue.Dequeue()

	queue.Print()
}
