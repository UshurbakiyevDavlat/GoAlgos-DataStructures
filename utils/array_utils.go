package utils

func FindMinMax(arr []int) (int, int) {
	if len(arr) == 0 {
		panic("Array shouldn't be empty!")
	}

	min := arr[0]
	max := arr[0]

	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	return min, max
}
