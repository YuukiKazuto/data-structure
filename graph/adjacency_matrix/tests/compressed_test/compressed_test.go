package compressed_test

import (
	"data-structure/graph"
	"data-structure/graph/adjacency_matrix/compressed"
	"fmt"
	"testing"
)

var g graph.Graph[int]

// 0                  0
// 1 0                2
// 0 0 0              5
// 0 0 1 0            9
// 1 0 0 0 0          14
// 0 1 1 0 0 0        20
// 0 0 1 1 0 1 0      27
// 0 0 0 1 0 0 1 0    35
func TestWGraph_BFSTraverse(t *testing.T) {
	traverse := make([]int, 0)
	g.BFSTraverse(2, func(v int) {
		traverse = append(traverse, v)
	})
	fmt.Println(traverse)
}

func TestWGraph_DFSTraverse(t *testing.T) {
	traverse := make([]int, 0)
	g.DFSTraverse(2, func(v int) {
		traverse = append(traverse, v)
	})
	fmt.Println(traverse)
}

func TestMain(m *testing.M) {
	vex := []int{1, 2, 3, 4, 5, 6, 7, 8}
	edge := make([]int, (len(vex)+1)*len(vex)/2)
	g = compressed.NewWGraph[int](vex, edge)
	g.InsertArc(1, 2, 1)
	g.InsertArc(1, 5, 1)
	g.InsertArc(2, 6, 1)
	g.InsertArc(3, 6, 1)
	g.InsertArc(3, 4, 1)
	g.InsertArc(3, 7, 1)
	g.InsertArc(4, 7, 1)
	g.InsertArc(4, 8, 1)
	g.InsertArc(6, 7, 1)
	g.InsertArc(7, 8, 1)
	m.Run()
}
