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

	// Сортировка выбором
	arr := []int{64, 25, 12, 22, 11}
	fmt.Println("Before sorting:", arr)

	sorting.SelectionSort(arr)

	fmt.Println("After sorting:", arr)
}
