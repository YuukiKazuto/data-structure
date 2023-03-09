package sort

func BubbleSort[T Number](arr []T, n int) {
	for i := n - 1; i >= 0; i-- {
		flag := false
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}
