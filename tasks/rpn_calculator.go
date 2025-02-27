package tasks

import (
	"fmt"
	"strconv"
	"strings"

	"main/data_structures"
)

func EvaluateRPN(expression string) (int, error) {
	if expression == "" {
		return 0, fmt.Errorf("invalid expression: invalid token: ")
	}

	stack := data_structures.Stack[int]{}
	tokens := strings.Split(expression, " ")

	for _, token := range tokens {
		if num, err := strconv.Atoi(token); err == nil {
			stack.Push(num)
		} else {
			if token == "+" || token == "-" || token == "*" || token == "/" {
				right, err := stack.Pop()
				if err != nil {
					return 0, fmt.Errorf("invalid expression: too many arguments")
				}
				left, err := stack.Pop()
				if err != nil {
					return 0, fmt.Errorf("invalid expression: too many arguments")
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
				}

				stack.Push(result)
			} else {
				return 0, fmt.Errorf("invalid token: %s", token)
			}
		}
	}

	if result, err := stack.Pop(); err == nil {
		if _, err := stack.Pop(); err == nil {
			return 0, fmt.Errorf("invalid expression: too many arguments")
		}
		return result, nil
	} else {
		return 0, fmt.Errorf("invalid expression: %s", err)
	}
}
