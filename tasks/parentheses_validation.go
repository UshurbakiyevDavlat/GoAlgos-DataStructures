package tasks

import "main/data_structures"

func IsValidParentheses(s string) bool {
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
