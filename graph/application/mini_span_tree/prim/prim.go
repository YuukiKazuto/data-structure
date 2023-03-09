package prim

import (
	"data-structure/graph"
	"data-structure/graph/adjacency_matrix"
	"fmt"
	"math"
)

const INFINITY = math.MaxInt

func MiniSpanTree[T comparable](g graph.Graph[T], u T) {
	uwg, ok := g.(*adjacency_matrix.UWGraph[T])
	if !ok {
		return
	}
	closedge := make([]struct {
		adjvex  T
		lowcost int
	}, len(uwg.Vex))
	k := g.LocateVex(u)
	for j := range uwg.Vex {
		if j != k {
			closedge[j].adjvex = u
			closedge[j].lowcost = uwg.Edge[k][j]
		}
	}
	closedge[k].lowcost = 0
	for i := 1; i < len(uwg.Vex); i++ {
		k = mininum[T](closedge)
		fmt.Println(closedge[k].adjvex, uwg.Vex[k])
		closedge[k].lowcost = 0
		for j := range uwg.Vex {
			if uwg.Edge[k][j] < closedge[j].lowcost {
				closedge[j].adjvex = uwg.Vex[k]
				closedge[j].lowcost = uwg.Edge[k][j]
			}
		}
	}
}

func mininum[T comparable](closedge []struct {
	adjvex  T
	lowcost int
}) (num int) {
	min := INFINITY
	for i, val := range closedge {
		if val.lowcost != 0 && val.lowcost < min {
			min = val.lowcost
			num = i
		}
	}
	return
}
