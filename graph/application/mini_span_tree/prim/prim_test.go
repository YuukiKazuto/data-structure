package prim

import (
	"data-structure/graph"
	"data-structure/graph/adjacency_matrix"
	"data-structure/graph/application/shortest_path/global"
	"testing"
)

var g graph.Graph[int]

func TestMiniSpanTree(t *testing.T) {
	MiniSpanTree[int](g, 0)
}
func TestMain(m *testing.M) {
	vex := []int{0, 1, 2, 3, 4, 5}
	arc := make([][]int, len(vex))
	for i, _ := range arc {
		arc[i] = make([]int, len(vex))
	}
	for i, row := range arc {
		for j, _ := range row {
			if i == j {
				arc[i][j] = 0
			} else {
				arc[i][j] = global.INFINITY
			}
		}
	}
	g = adjacency_matrix.NewUWGraph[int](vex, arc)
	g.InsertArc(0, 1, 6)
	g.InsertArc(0, 2, 5)
	g.InsertArc(0, 3, 1)
	g.InsertArc(1, 4, 3)
	g.InsertArc(1, 3, 5)
	g.InsertArc(2, 5, 2)
	g.InsertArc(2, 3, 4)
	g.InsertArc(3, 4, 6)
	g.InsertArc(3, 5, 4)
	g.InsertArc(4, 5, 6)
	m.Run()
}
