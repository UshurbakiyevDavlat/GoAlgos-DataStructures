package main

import (
	"fmt"
	"main/sorting"
	"main/tasks"
)

func main() {
	// Проверка скобок
	fmt.Println("Valid Parentheses:", tasks.IsValidParentheses("({[]})")) // true

	// Обратная польская запись
	rpnResult, err := tasks.EvaluateRPN("3 4 + 2 * 7 /")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("RPN Result:", rpnResult) // 2
	}

	// Прямоугольник в гистограмме
	heights := []int{2, 1, 5, 6, 2, 3}
	fmt.Println("Largest Rectangle Area:", tasks.LargestRectangleArea(heights)) // 10

	// Сортировка
	arr := []int{3, 8, 2, 1, 7}
	fmt.Println("Before sorting:", arr)

	sortedArr := sorting.QuickSort(arr)

	fmt.Println("After sorting:", sortedArr)

	target := 2

	index := sorting.BinarySearch(sortedArr, target)

	if index != -1 {
		fmt.Printf("Element %d found at index %d\n", target, index)
	} else {
		fmt.Printf("Element %d not found\n", target)
	}
}
