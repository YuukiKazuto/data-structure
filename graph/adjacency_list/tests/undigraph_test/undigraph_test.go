package undigraph_test

import (
	"data-structure/graph"
	"data-structure/graph/adjacency_list"
	"fmt"
	"testing"
)

var g graph.Graph[int]

func TestALGraph_DFSTraverse(t *testing.T) {
	traverse := make([]int, 0)
	g.DFSTraverse(2, func(v int) {
		traverse = append(traverse, v)
	})
	fmt.Println(traverse)
}

func TestALGraph_BFSTraverse(t *testing.T) {
	traverse := make([]int, 0)
	g.BFSTraverse(2, func(v int) {
		traverse = append(traverse, v)
	})
	fmt.Println(traverse)
}
func TestMain(m *testing.M) {
	vex := []int{1, 2, 3, 4, 5, 6, 7, 8}
	adjList := make([]adjacency_list.VNode[int], 0)
	for _, v := range vex {
		adjList = append(adjList, adjacency_list.VNode[int]{Data: v})
	}
	g = adjacency_list.NewALGraph[int](adjList)
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
