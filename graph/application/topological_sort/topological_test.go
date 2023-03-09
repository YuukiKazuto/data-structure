package topological_sort

import (
	"data-structure/graph"
	"data-structure/graph/adjacency_list"
	"fmt"
	"strconv"
	"testing"
)

var g graph.Graph[string]

func TestTopSort(t *testing.T) {
	v, p := TopSort[string](g)
	fmt.Println(v, p)
}

func TestMain(m *testing.M) {
	var vex [12]string
	for i := 1; i < 13; i++ {
		vex[i-1] = "C" + strconv.Itoa(i)
	}
	adjList := make([]adjacency_list.VNode[string], len(vex))
	for i, v := range vex {
		adjList[i].Data = v
	}
	g = adjacency_list.NewALGraphAllArgument[string](adjList, adjacency_list.DiGraph)
	g.InsertArc("C1", "C4", 1)
	g.InsertArc("C1", "C2", 1)
	g.InsertArc("C1", "C3", 1)
	g.InsertArc("C1", "C12", 1)
	g.InsertArc("C9", "C12", 1)
	g.InsertArc("C9", "C10", 1)
	g.InsertArc("C9", "C11", 1)
	g.InsertArc("C4", "C5", 1)
	g.InsertArc("C2", "C3", 1)
	g.InsertArc("C10", "C12", 1)
	g.InsertArc("C11", "C6", 1)
	g.InsertArc("C3", "C5", 1)
	g.InsertArc("C3", "C7", 1)
	g.InsertArc("C3", "C8", 1)
	g.InsertArc("C6", "C8", 1)
	g.InsertArc("C5", "C7", 1)
	m.Run()
}
