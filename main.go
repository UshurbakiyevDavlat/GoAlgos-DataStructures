package main

import (
	"fmt"
	"main/tasks"
)

func main() {
	// Задача 1: Проверка скобок
	fmt.Println("Valid Parentheses:", tasks.IsValidParentheses("({[]})")) // true

	// Задача 2: Обратная польская запись
	rpnResult, err := tasks.EvaluateRPN("3 4 + 2 * 7 /")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("RPN Result:", rpnResult) // 2
	}

	// Задача 3: Прямоугольник в гистограмме
	heights := []int{2, 1, 5, 6, 2, 3}
	fmt.Println("Largest Rectangle Area:", tasks.LargestRectangleArea(heights)) // 10
}
