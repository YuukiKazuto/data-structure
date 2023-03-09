package bfs

import (
	"data-structure/graph"
	"data-structure/graph/adjacency_list"
	"data-structure/graph/application/shortest_path/global"
	"fmt"
	"testing"
)

var g graph.Graph[int]

func TestShorterPath(t *testing.T) {
	ShorterPath[int](g, 2)
	fmt.Println(global.Dist)
	fmt.Println(global.Path)
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
	global.InitializeBFS[int](g)
	m.Run()
}
