package sort

func InsertionSort[T Number](arr []T, n int) {
	for i := 1; i < n; i++ {
		t := arr[i]
		j := 0
		for j = i; j > 0 && arr[j-1] > t; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = t
	}
}
