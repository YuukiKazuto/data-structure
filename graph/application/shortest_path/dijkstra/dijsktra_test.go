package dijkstra

import (
	"data-structure/graph"
	"data-structure/graph/adjacency_matrix"
	"data-structure/graph/application/shortest_path/global"
	"fmt"
	"testing"
)

var g graph.Graph[int]

func TestShorterPath(t *testing.T) {
	ShorterPath[int](g)
	fmt.Println(global.Path)
	fmt.Println(global.Dist)
}

func TestMain(m *testing.M) {
	vex := []int{0, 1, 2, 3, 4}
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
	g = adjacency_matrix.NewWDigraph[int](vex, arc)
	g.InsertArc(0, 1, 10)
	g.InsertArc(0, 4, 5)
	g.InsertArc(1, 4, 2)
	g.InsertArc(4, 1, 3)
	g.InsertArc(1, 2, 1)
	g.InsertArc(2, 3, 4)
	g.InsertArc(4, 2, 9)
	g.InsertArc(4, 3, 2)
	g.InsertArc(3, 2, 6)
	g.InsertArc(3, 1, 7)
	global.InitializeDijkstra[int](g)
	m.Run()
}
