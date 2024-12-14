package tasks

func LargestRectangleArea(heights []int) int {
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
