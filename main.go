package main

import (
	"fmt"
	"main/data_structures"
	"strconv"
	"strings"
)

func main() {
	examples := []string{
		"({[]})", // true
		"([)]",   // false
		"(]",     // false
		"()",     // true
		"",       // true
	}

	for _, example := range examples {
		fmt.Printf("Input: %s, Valid: %v\n", example, isValidParentheses(example))
	}
}

func isValidParentheses(s string) bool {
	stack := data_structures.Stack[rune]{}
	matchingBrackets := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, char := range s {
		if open, ok := matchingBrackets[char]; ok {
			top, err := stack.Pop()
			if err != nil || top != open {
				return false
			}
		} else {
			stack.Push(char)
		}
	}
	return stack.IsEmpty()
}
func evaluateRPN(expression string) (int, error) {
	stack := data_structures.Stack[int]{}
	tokens := strings.Split(expression, " ")

	for _, token := range tokens {
		if num, err := strconv.Atoi(token); err == nil {
			stack.Push(num)
		} else {
			if token == "+" || token == "-" || token == "*" || token == "/" {
				right, err := stack.Pop()
				if err != nil {
					return 0, fmt.Errorf("invalid expression: %s", err)
				}
				left, err := stack.Pop()
				if err != nil {
					return 0, fmt.Errorf("invalid expression: %s", err)
				}

				var result int
				switch token {
				case "+":
					result = left + right
				case "-":
					result = left - right
				case "*":
					result = left * right
				case "/":
					if right == 0 {
						return 0, fmt.Errorf("division by zero")
					}
					result = left / right
				default:
					return 0, fmt.Errorf("unknown operator: %s", token)
				}

				stack.Push(result)
			} else {
				return 0, fmt.Errorf("invalid token: %s", token)
			}
		}
	}

	if result, err := stack.Pop(); err == nil {
		if _, err := stack.Pop(); err == nil {
			return 0, fmt.Errorf("invalid expression: too many agruments")
		}
		return result, nil
	} else {
		return 0, fmt.Errorf("invalid expression: %s", err)
	}
}
