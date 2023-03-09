package sort

func SelectionSort[T Number](arr []T, n int) {
	for i := range arr {
		ml := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[ml] {
				ml = j
			}
		}
		arr[i], arr[ml] = arr[ml], arr[i]
	}
}
