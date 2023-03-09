package critical_path

import (
	"data-structure/graph"
	"data-structure/graph/adjacency_list"
	"data-structure/graph/application/topological_sort"
	"data-structure/stack"
	"fmt"
)

var ve []int

func CriticalPath[T comparable](g graph.Graph[T]) {
	alg, ok := g.(*adjacency_list.ALGraph[T])
	if !ok || alg.Kind != adjacency_list.DiGraph {
		panic("graph is not a digraph")
	}
	t := TopologicalOrder[T](alg)
	if t == nil {
		panic("There are loops in the diagram")
	}
	vl := make([]int, len(ve))
	for i := range vl {
		vl[i] = ve[len(ve)-1]
	}
	for !t.IsEmpty() {
		v, _ := t.Pop()
		j := g.LocateVex(v)
		for p := alg.AdjList[j].First; p != nil; p = p.Next {
			k := p.AdjVex
			if vl[k]-p.Weight < vl[j] {
				vl[j] = vl[k] - p.Weight
			}
		}
	}
	fmt.Println(vl)
	for j := 0; j < len(alg.AdjList); j++ {
		for p := alg.AdjList[j].First; p != nil; p = p.Next {
			k, dut := p.AdjVex, p.Weight
			ee, el := ve[j], vl[k]-dut
			var tag string
			if ee == el {
				tag = "*"
			}
			fmt.Print("arc<", j, ",", k, "> ", dut, ee, el, " ", tag)
			fmt.Println()
		}
	}
}

func TopologicalOrder[T comparable](g *adjacency_list.ALGraph[T]) stack.Stack[T] {
	indegree := topological_sort.FindInDegree(g)
	s := stack.NewSequenceStack[int]()
	for _, in := range indegree {
		if in == 0 {
			s.Push(in)
		}
	}
	count := 0
	ve = make([]int, len(g.AdjList))
	t := stack.NewSequenceStack[T]()
	for !s.IsEmpty() {
		j, _ := s.Pop()
		t.Push(g.AdjList[j].Data)
		count++
		for p := g.AdjList[j].First; p != nil; p = p.Next {
			k := p.AdjVex
			indegree[k]--
			if indegree[k] == 0 {
				s.Push(k)
			}
			if ve[j]+p.Weight > ve[k] {
				ve[k] = ve[j] + p.Weight
			}
		}
	}
	if count < len(g.AdjList) {
		return nil
	}
	return t
}
