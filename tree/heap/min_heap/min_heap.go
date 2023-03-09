package min_heap

import "math"

type MinHeap struct {
	Elements []int // 下标0不放值
}

func NewMinHeap() *MinHeap {
	minData := math.MinInt // 哨兵 放在Elements[0]
	return &MinHeap{Elements: []int{minData}}
}

func (h *MinHeap) IsEmpty() bool {
	return len(h.Elements) == 1
}

func (h *MinHeap) Insert(data int) {
	h.Elements = append(h.Elements, 0)
	var i int
	for i = len(h.Elements) - 1; h.Elements[i/2] > data; i /= 2 {
		h.Elements[i] = h.Elements[i/2]
	}
	h.Elements[i] = data
}

func (h *MinHeap) Build(arr []int) {
	//TODO implement me
	panic("implement me")
}
