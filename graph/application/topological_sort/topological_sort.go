package topological_sort

import (
	"data-structure/graph"
	"data-structure/graph/adjacency_list"
	"data-structure/queue"
)

func TopSort[T comparable](g graph.Graph[T]) ([]T, []struct{ Pre []int }) {
	alg, ok := g.(*adjacency_list.ALGraph[T])
	if !ok || alg.Kind != adjacency_list.DiGraph {
		panic("graph is not a digraph")
	}
	q := queue.NewLinkedQueue[T]()
	indegree := FindInDegree(alg)
	for i, v := range alg.AdjList {
		if indegree[i] == 0 {
			q.EnQueue(v.Data)
		}
	}
	vertices := make([]T, 0)
	path := make([]struct{ Pre []int }, len(alg.AdjList))
	count := 0
	for !q.IsEmpty() { // O(|V|+|E|)
		v, _ := q.DeQueue()
		vi := alg.LocateVex(v)
		vertices = append(vertices, v)
		count++
		var zeroT T
		for w := alg.FirstAdjVex(v); w != zeroT; w = alg.NextAdjVex(v, w) {
			wi := alg.LocateVex(w)
			path[wi].Pre = append(path[wi].Pre, vi)
			indegree[wi]--
			if indegree[wi] == 0 {
				q.EnQueue(w)
			}
		}
	}
	if count != len(alg.AdjList) {
		panic("There are loops in the diagram")
	}
	return vertices, path
}

func FindInDegree[T comparable](alg *adjacency_list.ALGraph[T]) []int {
	indegree := make([]int, len(alg.AdjList))
	for _, v := range alg.AdjList { // O(|V|+|E|)
		var zeroT T
		for w := alg.FirstAdjVex(v.Data); w != zeroT; w = alg.NextAdjVex(v.Data, w) {
			wi := alg.LocateVex(w)
			indegree[wi]++
		}
	}
	return indegree
}
