package kruskal

import (
	"container/heap"
	"data-structure/graph"
	"data-structure/graph/adjacency_matrix"
	"data-structure/tree/set"
	"fmt"
)

type edgeDetail struct {
	weight int
	i, j   int
}

type edgeHeap []edgeDetail

func (h *edgeHeap) Len() int {
	return len(*h)
}

func (h *edgeHeap) Less(i, j int) bool {
	return (*h)[i].weight < (*h)[j].weight
}

func (h *edgeHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *edgeHeap) Push(x any) {
	*h = append(*h, x.(edgeDetail))
}

func (h *edgeHeap) Pop() any {
	x := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return x
}

func MiniSpanTree[T comparable](g graph.Graph[T]) {
	uwg, ok := g.(*adjacency_matrix.UWGraph[T])
	if !ok {
		return
	}
	h := &edgeHeap{}
	heap.Init(h)
	s := set.NewSet[T](uwg.Vex)
	for i := range uwg.Edge {
		for j := 0; j < i; j++ {
			if uwg.Edge[i][j] != adjacency_matrix.INFINITY {
				heap.Push(h, edgeDetail{
					weight: uwg.Edge[i][j],
					i:      i,
					j:      j,
				})
			}
		}
	}
	mst := make([]edgeDetail, 0)
	for len(mst) < len(uwg.Vex)-1 && h.Len() != 0 {
		min := heap.Pop(h)
		minD := min.(edgeDetail)
		if s.Find(uwg.Vex[minD.i]) != s.Find(uwg.Vex[minD.j]) {
			s.Union(uwg.Vex[minD.i], uwg.Vex[minD.j])
			mst = append(mst, minD)
		}
	}
	fmt.Println(mst)
}
