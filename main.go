package main

import (
	"fmt"
)

func main() {
	heights := []int{2, 1, 5, 6, 2, 3}
	fmt.Println("Largest Rectangle Area:", largestRectangleArea(heights))
}

// func isValidParentheses(s string) bool {
// 	stack := data_structures.Stack[rune]{}
// 	matchingBrackets := map[rune]rune{
// 		')': '(',
// 		'}': '{',
// 		']': '[',
// 	}
// 	for _, char := range s {
// 		if open, ok := matchingBrackets[char]; ok {
// 			top, err := stack.Pop()
// 			if err != nil || top != open {
// 				return false
// 			}
// 		} else {
// 			stack.Push(char)
// 		}
// 	}
// 	return stack.IsEmpty()
// }
// func evaluateRPN(expression string) (int, error) {
// 	stack := data_structures.Stack[int]{}
// 	tokens := strings.Split(expression, " ")

// 	for _, token := range tokens {
// 		if num, err := strconv.Atoi(token); err == nil {
// 			stack.Push(num)
// 		} else {
// 			if token == "+" || token == "-" || token == "*" || token == "/" {
// 				right, err := stack.Pop()
// 				if err != nil {
// 					return 0, fmt.Errorf("invalid expression: %s", err)
// 				}
// 				left, err := stack.Pop()
// 				if err != nil {
// 					return 0, fmt.Errorf("invalid expression: %s", err)
// 				}

// 				var result int
// 				switch token {
// 				case "+":
// 					result = left + right
// 				case "-":
// 					result = left - right
// 				case "*":
// 					result = left * right
// 				case "/":
// 					if right == 0 {
// 						return 0, fmt.Errorf("division by zero")
// 					}
// 					result = left / right
// 				default:
// 					return 0, fmt.Errorf("unknown operator: %s", token)
// 				}

// 				stack.Push(result)
// 			} else {
// 				return 0, fmt.Errorf("invalid token: %s", token)
// 			}
// 		}
// 	}

// 	if result, err := stack.Pop(); err == nil {
// 		if _, err := stack.Pop(); err == nil {
// 			return 0, fmt.Errorf("invalid expression: too many agruments")
// 		}
// 		return result, nil
// 	} else {
// 		return 0, fmt.Errorf("invalid expression: %s", err)
// 	}
// }

func largestRectangleArea(heights []int) int {
	stack := []int{}
	maxArea := 0
	index := 0

	for index < len(heights) {
		if len(stack) == 0 || heights[index] >= heights[stack[len(stack)-1]] {
			stack = append(stack, index)
			index++
		} else {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			width := index
			if len(stack) > 0 {
				width = index - stack[len(stack)-1] - 1
			}

			maxArea = max(maxArea, heights[top]*width)
		}
	}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		width := index
		if len(stack) > 0 {
			width = index - stack[len(stack)-1] - 1
		}
		maxArea = max(maxArea, heights[top]*width)
	}

	return maxArea
}
