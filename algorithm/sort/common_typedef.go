package sort

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

type Minheap[T Number] []T

func BuildHeap[T Number](slice []T) *Minheap[T] {
	minheap := make([]T, len(slice))
	copy(minheap, slice)
	return (*Minheap[T])(&minheap)
}
func (h *Minheap[T]) Len() int {
	return len(*h)
}

func (h *Minheap[T]) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *Minheap[T]) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Minheap[T]) Push(x any) {
	*h = append(*h, x.(T))
}

func (h *Minheap[T]) Pop() any {
	r := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return r
}

type Maxheap[T Number] []T

func (h *Maxheap[T]) Len() int {
	return len(*h)
}

func (h *Maxheap[T]) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *Maxheap[T]) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Maxheap[T]) Push(x any) {
	*h = append(*h, x.(T))
}

func (h *Maxheap[T]) Pop() any {
	r := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return r
}
