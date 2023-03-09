package critical_path

import (
	"data-structure/graph"
	"data-structure/graph/adjacency_list"
	"strconv"
	"testing"
)

var g graph.Graph[string]

func TestCriticalPath(t *testing.T) {
	CriticalPath(g)
}
func TestMain(m *testing.M) {
	var vex [6]string
	for i := 1; i < 7; i++ {
		vex[i-1] = "C" + strconv.Itoa(i)
	}
	adjList := make([]adjacency_list.VNode[string], len(vex))
	for i, v := range vex {
		adjList[i].Data = v
	}
	g = adjacency_list.NewALGraphAllArgument[string](adjList, adjacency_list.DiGraph)
	g.InsertArc("C1", "C2", 3)
	g.InsertArc("C1", "C3", 2)
	g.InsertArc("C2", "C4", 2)
	g.InsertArc("C2", "C5", 3)
	g.InsertArc("C3", "C4", 4)
	g.InsertArc("C3", "C6", 3)
	g.InsertArc("C4", "C6", 2)
	g.InsertArc("C5", "C6", 1)
	m.Run()
}
