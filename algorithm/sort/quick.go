package sort

const Cutoff = 100

func median3[T Number](arr []T, l, r int) T {
	center := (l + r) / 2
	if arr[l] > arr[center] {
		arr[l], arr[center] = arr[center], arr[l]
	}
	if arr[l] > arr[r] {
		arr[l], arr[r] = arr[r], arr[l]
	}
	if arr[center] > arr[r] {
		arr[center], arr[r] = arr[r], arr[center]
	}
	arr[center], arr[r-1] = arr[r-1], arr[center]
	return arr[r-1]
}

func quicksort[T Number](arr []T, l, r int) {
	if Cutoff <= r-l {
		pivot := median3(arr, l, r)
		i, j := l+1, r-2
		for {
			for arr[i] < pivot {
				i++
			}
			for arr[j] > pivot {
				j--
			}
			if i < j {
				arr[i], arr[j] = arr[j], arr[i]
			} else {
				break
			}
		}
		arr[i], arr[r-1] = arr[r-1], arr[i]
		quicksort(arr, l, i-1)
		quicksort(arr, i, r)
	} else {
		InsertionSort(arr[l:], r-l+1)
	}
}

func QuickSort[T Number](arr []T, n int) {
	quicksort(arr, 0, n-1)
}
