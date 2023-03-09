package bfs

import (
	"data-structure/graph"
	. "data-structure/graph/application/shortest_path/global"
	"data-structure/queue"
)

func ShorterPath[T comparable](g graph.Graph[T], v T) {
	q := Q.(*queue.SequenceQueue[T])
	for i, _ := range Visited {
		Dist[i] = INFINITY
		Path[i] = -1
	}
	var zeroT T
	vi := g.LocateVex(v)
	Dist[vi] = 0
	Visited[vi] = true
	q.EnQueue(v)
	for !q.IsEmpty() {
		x, _ := q.DeQueue()
		xi := g.LocateVex(x)
		for w := g.FirstAdjVex(x); w != zeroT; w = g.NextAdjVex(x, w) {
			wi := g.LocateVex(w)
			if !Visited[wi] {
				Dist[wi] = Dist[xi] + 1
				Path[wi] = xi
				Visited[wi] = true
				q.EnQueue(w)
			}
		}
	}
}
