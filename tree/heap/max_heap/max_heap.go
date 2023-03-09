package max_heap

import "math"

type MaxHeap struct {
	Elements []int // 下标0不放值
}

func NewMaxHeap() *MaxHeap {
	maxData := math.MaxInt // 哨兵 放在Elements[0]
	return &MaxHeap{[]int{maxData}}
}

func (h *MaxHeap) Insert(data int) {
	h.Elements = append(h.Elements, 0)
	var i int
	for i = len(h.Elements) - 1; h.Elements[i/2] < data; i /= 2 {
		h.Elements[i] = h.Elements[i/2]
	}
	h.Elements[i] = data
}

func (h *MaxHeap) IsEmpty() bool {
	return len(h.Elements) == 1
}

func (h *MaxHeap) DeleteMax() int {
	if h.IsEmpty() {
		return 0
	}
	var parent, child int
	size := len(h.Elements) - 1
	maxItem := h.Elements[1]
	temp := h.Elements[size]
	for parent = 1; parent*2 <= size; parent = child {
		child = parent * 2
		if child != size && h.Elements[child] < h.Elements[child+1] {
			child++
		}
		if temp >= h.Elements[child] {
			break
		} else {
			h.Elements[parent] = h.Elements[child]
		}
	}
	h.Elements[parent] = temp
	h.Elements = h.Elements[:size]
	return maxItem
}

func (h *MaxHeap) Build(arr []int) {
	h.Elements = append(h.Elements, arr...)
	size := len(h.Elements) - 1
	for i := size / 2; i > 0; i-- {
		var parent, child int
		ei := h.Elements[i]
		for parent = i; parent*2 <= size; parent = child {
			child = 2 * parent
			if child != size && h.Elements[child] < h.Elements[child+1] {
				child++
			}
			if ei >= h.Elements[child] {
				break
			} else {
				h.Elements[parent] = h.Elements[child]
			}
		}
	}
}
