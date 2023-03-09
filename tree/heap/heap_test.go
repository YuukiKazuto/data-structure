package heap

import (
	"data-structure/tree/heap/max_heap"
	"data-structure/tree/heap/min_heap"
	"fmt"
	"testing"
)

func TestMaxHeap_Insert(t *testing.T) {
	heap := max_heap.NewMaxHeap()
	for i := 1; i < 10; i++ {
		heap.Insert(i*2 + 1)
	}
	fmt.Println(heap.Elements[1:])
	heap.DeleteMax()
	fmt.Println(heap.Elements[1:])
}

func TestMinHeap_Insert(t *testing.T) {
	heap := min_heap.NewMinHeap()
	for i := 1; i < 10; i++ {
		heap.Insert(i*2 + 1)
	}
	fmt.Println(heap.Elements[1:])
}
