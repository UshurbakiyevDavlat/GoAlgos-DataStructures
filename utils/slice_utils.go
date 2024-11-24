package utils

func ReverseSlice(slice []int) []int {
	reversed := make([]int, len(slice))
	for i := 0; i < len(slice); i++ {
		reversed[i] = slice[len(slice)-1-i]
	}
	return reversed
}

func SplitEvenOdd(slice []int) ([]int, []int) {
	oddSlice := []int{}
	evenSlice := []int{}

	for _, v := range slice {
		if v%2 != 0 {
			oddSlice = append(oddSlice, v)
		} else {
			evenSlice = append(evenSlice, v)
		}
	}

	return oddSlice, evenSlice
}

func RemoveElement(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		panic("Index out of range")
	}
	copy(slice[index:], slice[index+1:])
	return slice[:len(slice)-1]
}
