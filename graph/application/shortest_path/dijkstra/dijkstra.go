package dijkstra

import (
	"data-structure/graph"
	"data-structure/graph/adjacency_matrix"
	. "data-structure/graph/application/shortest_path/global"
)

// ShorterPath 目前只能使用邻接矩阵存储的图
func ShorterPath[T comparable](g graph.Graph[T]) {
	for i, _ := range Final {
		Dist[i] = INFINITY
		Path[i] = -1
	}
	var wdg *adjacency_matrix.WDigraph[T]
	if temp, ok := g.(*adjacency_matrix.WDigraph[T]); ok {
		wdg = temp
	}
	if wdg == nil {
		return
	}
	var zeroT T
	Dist[0] = 0
	Final[0] = true
	for w := g.FirstAdjVex(wdg.Vex[0]); w != zeroT; w = g.NextAdjVex(wdg.Vex[0], w) {
		wi := g.LocateVex(w)
		Dist[wi] = wdg.Arc[0][wi]
		if wdg.Arc[0][wi] != INFINITY {
			Path[wi] = 0
		}
	}
	for i := 0; i < len(Final)-1; i++ {
		minLoc := max(Dist)
		for k, v := range Final {
			if v {
				continue
			}
			if Dist[k] < Dist[minLoc] {
				minLoc = k
			}
		}
		Final[minLoc] = true
		for w := g.FirstAdjVex(wdg.Vex[minLoc]); w != zeroT; w = g.NextAdjVex(wdg.Vex[minLoc], w) {
			wi := g.LocateVex(w)
			if !Final[wi] && Dist[minLoc]+wdg.Arc[minLoc][wi] < Dist[wi] {
				Dist[wi] = Dist[minLoc] + wdg.Arc[minLoc][wi]
				Path[wi] = minLoc
			}
		}
	}
}

func max(dist []int) int {
	maxLoc := 0
	for i, v := range dist {
		if v > dist[maxLoc] {
			maxLoc = i
		}
	}
	return maxLoc
}
