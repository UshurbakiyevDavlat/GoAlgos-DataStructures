package sorting

func SelectionSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		minIndex := i

		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ { // 3 8 2 1 7
		key := arr[i] //arr[2] -> 2
		j := i - 1    // arr[1] -> 8

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}
}
