package sort

func ShellSort[T Number](arr []T, n int) {
	for d := n / 2; d > 0; d = d / 2 {
		for i := d; i < n; i++ {
			t := arr[i]
			j := 0
			for j = i; j >= d && arr[j-d] > t; j -= d {
				arr[j] = arr[j-d]
			}
			arr[j] = t
		}
	}
}
