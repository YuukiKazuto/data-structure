package sort

func merge[T Number](arr, tmpArr []T, l, r, rEnd int, isRecursion bool) {
	lEnd := r - 1
	tmp := l
	num := rEnd - l + 1
	for ; l <= lEnd && r <= rEnd; tmp++ {
		if arr[l] <= arr[r] {
			tmpArr[tmp] = arr[l]
			l++
		} else {
			tmpArr[tmp] = arr[r]
			r++
		}
	}
	for ; l <= lEnd; tmp, l = tmp+1, l+1 {
		tmpArr[tmp] = arr[l]
	}
	for ; r <= rEnd; tmp, r = tmp+1, r+1 {
		tmpArr[tmp] = arr[r]
	}
	if isRecursion {
		for i := 0; i < num; i, rEnd = i+1, rEnd-1 {
			arr[rEnd] = tmpArr[rEnd]
		}
	}
}

func mSort[T Number](arr, tmpArr []T, l, rEnd int) {
	center := -1
	if l < rEnd {
		center = (l + rEnd) / 2
		mSort(arr, tmpArr, l, center)
		mSort(arr, tmpArr, center+1, rEnd)
		merge(arr, tmpArr, l, center+1, rEnd, true)
	}
}

func MergeSortRecursion[T Number](arr []T, n int) {
	tmpArr := make([]T, n)
	mSort(arr, tmpArr, 0, n-1)
	tmpArr = nil
}

func MergeSort[T Number](arr []T, n int) {
	length := 1
	tmpArr := make([]T, n)
	for length < n {
		mergePass(arr, tmpArr, length)
		length *= 2
		mergePass(tmpArr, arr, length)
		length *= 2
	}
	tmpArr = nil
}

func mergePass[T Number](arr, tmpArr []T, length int) {
	n := len(arr)
	var i int
	for i = 0; i <= n-2*length; i += 2 * length {
		merge(arr, tmpArr, i, i+length, i+2*length-1, false)
	}
	if i+length < n {
		merge(arr, tmpArr, i, i+length, n-1, false)
	} else {
		for j := i; j < n; j++ {
			tmpArr[j] = arr[j]
		}
	}
}
