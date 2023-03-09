package sort

import "container/heap"

func HeapSortDoubleSpace[T Number](arr []T, n int) {
	h := BuildHeap(arr)
	heap.Init(h)
	for i := 0; i < n; i++ {
		arr[i] = heap.Pop(h).(T)
	}
	h = nil
}

func HeapSort[T Number](arr Maxheap[T], n int) {
	heap.Init(&arr)
	for i := n - 1; i > 0; i-- {
		arr.Swap(0, i)
		down(&arr, 0, i)
	}
}

func down(h heap.Interface, i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
	return i > i0
}
