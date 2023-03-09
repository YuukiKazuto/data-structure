package global

import (
	"data-structure/graph"
	"data-structure/graph/adjacency_list"
	"data-structure/graph/adjacency_matrix"
	"data-structure/graph/adjacency_matrix/compressed"
	"data-structure/queue"
)

func InitializeBFS[T comparable](g graph.Graph[T]) {
	Q = queue.NewSequenceQueue[T]()
	if alG, ok := g.(*adjacency_list.ALGraph[T]); ok {
		Visited = make([]bool, len(alG.AdjList))
		Dist = make([]int, len(alG.AdjList))
		Path = make([]int, len(alG.AdjList))
	} else if uwG, ok := g.(*adjacency_matrix.UWGraph[T]); ok {
		Visited = make([]bool, len(uwG.Vex))
		Dist = make([]int, len(uwG.Vex))
		Path = make([]int, len(uwG.Vex))
	} else if wG, ok := g.(*compressed.WGraph[T]); ok {
		Visited = make([]bool, len(wG.Vex))
		Dist = make([]int, len(wG.Vex))
		Path = make([]int, len(wG.Vex))
	} else {
		return
	}
}

func InitializeDijkstra[T comparable](g graph.Graph[T]) {
	if alG, ok := g.(*adjacency_list.ALGraph[T]); ok {
		Final = make([]bool, len(alG.AdjList))
		Dist = make([]int, len(alG.AdjList))
		Path = make([]int, len(alG.AdjList))
	} else if uwG, ok := g.(*adjacency_matrix.UWGraph[T]); ok {
		Final = make([]bool, len(uwG.Vex))
		Dist = make([]int, len(uwG.Vex))
		Path = make([]int, len(uwG.Vex))
	} else if wG, ok := g.(*compressed.WGraph[T]); ok {
		Final = make([]bool, len(wG.Vex))
		Dist = make([]int, len(wG.Vex))
		Path = make([]int, len(wG.Vex))
	} else if wD, ok := g.(*adjacency_matrix.WDigraph[T]); ok {
		Final = make([]bool, len(wD.Vex))
		Dist = make([]int, len(wD.Vex))
		Path = make([]int, len(wD.Vex))
	} else {
		return
	}
}

func InitializeFloyd[T comparable](g graph.Graph[T]) {
	if wD, ok := g.(*adjacency_matrix.WDigraph[T]); ok {
		PathM = make([][]int, len(wD.Vex))
		for i, _ := range PathM {
			PathM[i] = make([]int, len(wD.Vex))
		}
		A = make([][]int, len(wD.Arc))
		for i, _ := range A {
			A[i] = make([]int, len(wD.Arc[i]))
			copy(A[i], wD.Arc[i])
		}
	} else {
		return
	}
}
