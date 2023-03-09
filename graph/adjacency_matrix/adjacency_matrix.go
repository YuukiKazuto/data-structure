package adjacency_matrix

import (
	"data-structure/queue"
	"math"
)

const (
	VexNotFound = -1
	INFINITY    = math.MaxInt
)

// UWGraph 无向带权图
type UWGraph[T comparable] struct {
	Vex  []T
	Edge [][]int
}

func NewUWGraph[T comparable](vex []T, edge [][]int) *UWGraph[T] {
	return &UWGraph[T]{Vex: vex, Edge: edge}
}

func (g *UWGraph[T]) LocateVex(v T) int {
	if g.Vex == nil {
		return VexNotFound
	}
	for i, vex := range g.Vex {
		if v == vex {
			return i
		}
	}
	return VexNotFound
}

func (g *UWGraph[T]) FirstAdjVex(v T) (res T) {
	if index := g.LocateVex(v); index != VexNotFound {
		for i, edgeVex := range g.Edge[index] {
			if edgeVex != 0 && edgeVex != INFINITY {
				res = g.Vex[i]
				return
			}
		}
	}
	return
}

func (g *UWGraph[T]) NextAdjVex(v, w T) (res T) {
	vi := g.LocateVex(v)
	wi := g.LocateVex(w)
	if vi != VexNotFound && wi != VexNotFound && g.Edge[vi][wi] != 0 && wi != len(g.Vex)-1 {
		for i := wi + 1; i < len(g.Edge[vi]); i++ {
			if g.Edge[vi][i] != 0 && g.Edge[vi][i] != INFINITY {
				res = g.Vex[i]
				return
			}
		}
	}
	return
}

func (g *UWGraph[T]) InsertVex(v T) {
	if g.Vex == nil || g.Edge == nil || len(g.Edge) == 0 {
		g.Vex = []T{v}
		g.Edge = [][]int{{0}}
		return
	}
	if g.LocateVex(v) != VexNotFound {
		return
	}
	g.Vex = append(g.Vex, v)
	for i, _ := range g.Edge {
		g.Edge[i] = append(g.Edge[i], 0)
	}
	newRow := make([]int, len(g.Edge[0]))
	g.Edge = append(g.Edge, newRow)
}

func (g *UWGraph[T]) DeleteVex(v T) (res T) {
	if g.Vex == nil || len(g.Vex) == 0 || g.Edge == nil || len(g.Edge) == 0 {
		return
	}
	if index := g.LocateVex(v); index != VexNotFound {
		res = g.Vex[index]
		g.Vex = append(g.Vex[:index-1], g.Vex[index+1:]...)
		g.Edge = append(g.Edge[:index-1], g.Edge[index+1:]...)
		for i, _ := range g.Edge {
			g.Edge[i] = append(g.Edge[i][:index-1], g.Edge[i][index+1:]...)
		}
	}
	return
}

func (g *UWGraph[T]) InsertArc(v, w T, weight int) {
	vi := g.LocateVex(v)
	wi := g.LocateVex(w)
	if vi != VexNotFound && wi != VexNotFound {
		g.Edge[vi][wi] = weight
		g.Edge[wi][vi] = weight
	}
}

func (g *UWGraph[T]) DeleteArc(v, w T) (weight int) {
	vi := g.LocateVex(v)
	wi := g.LocateVex(w)
	if vi == VexNotFound && wi == VexNotFound && vi == wi {
		return
	}
	weight = g.Edge[vi][wi]
	g.Edge[vi][wi] = INFINITY
	g.Edge[wi][vi] = INFINITY
	return
}

func (g *UWGraph[T]) BFSTraverse(start T, visit func(v T)) {
	if g.Vex == nil || len(g.Vex) == 0 || g.Edge == nil || len(g.Edge) == 0 {
		return
	}
	vi := g.LocateVex(start)
	if vi == VexNotFound {
		return
	}
	var zeroT T
	q := queue.NewSequenceQueue[T]()
	visited := make([]bool, len(g.Vex))
	for i, j := vi, 0; j < len(visited); j++ {
		if i == len(visited) {
			i = 0
		}
		if !visited[i] {
			v := g.Vex[i]
			visit(v)
			visited[i] = true
			q.EnQueue(v)
			for !q.IsEmpty() {
				v, _ = q.DeQueue()
				for w := g.FirstAdjVex(v); w != zeroT; w = g.NextAdjVex(v, w) {
					wi := g.LocateVex(w)
					if !visited[wi] {
						visit(w)
						visited[wi] = true
						q.EnQueue(w)
					}
				}
			}
		}
		i++
	}
}

func (g *UWGraph[T]) DFSTraverse(start T, visit func(v T)) {
	if g.Vex == nil || len(g.Vex) == 0 || g.Edge == nil || len(g.Edge) == 0 {
		return
	}
	vi := g.LocateVex(start)
	if vi == VexNotFound {
		return
	}
	visited := make([]bool, len(g.Vex))
	for i, j := vi, 0; j < len(visited); i, j = i+1, j+1 {
		if i == len(visited) {
			i = 0
		}
		if !visited[i] {
			g.dfs(&visited, g.Vex[i], visit)
		}
	}
}

func (g *UWGraph[T]) dfs(visitedP *[]bool, v T, visit func(v T)) {
	visit(v)
	vi := g.LocateVex(v)
	(*visitedP)[vi] = true
	var zeroT T
	for w := g.FirstAdjVex(v); w != zeroT; w = g.NextAdjVex(v, w) {
		wi := g.LocateVex(w)
		if !(*visitedP)[wi] {
			g.dfs(visitedP, w, visit)
		}
	}
}

// WDigraph 有向带权图
type WDigraph[T comparable] struct {
	Vex []T
	Arc [][]int
}

func NewWDigraph[T comparable](vex []T, arc [][]int) *WDigraph[T] {
	return &WDigraph[T]{Vex: vex, Arc: arc}
}

func (g *WDigraph[T]) LocateVex(v T) int {
	if g.Vex == nil {
		return VexNotFound
	}
	for i, vertex := range g.Vex {
		if v == vertex {
			return i
		}
	}
	return VexNotFound
}

func (g *WDigraph[T]) FirstAdjVex(v T) (res T) {
	if index := g.LocateVex(v); index != VexNotFound {
		for i, arcVex := range g.Arc[index] {
			if arcVex != 0 && arcVex != INFINITY {
				res = g.Vex[i]
				return
			}
		}
	}
	return
}

func (g *WDigraph[T]) NextAdjVex(v, w T) (res T) {
	vi := g.LocateVex(v)
	wi := g.LocateVex(w)
	if vi != VexNotFound && wi != VexNotFound && g.Arc[vi][wi] != 0 && wi != len(g.Vex)-1 {
		for i := wi + 1; i < len(g.Arc[vi]); i++ {
			if g.Arc[vi][i] != 0 && g.Arc[vi][i] != INFINITY {
				res = g.Vex[i]
				return
			}
		}
	}
	return
}

func (g *WDigraph[T]) InsertVex(v T) {
	if g.Vex == nil || g.Arc == nil || len(g.Arc) == 0 {
		g.Vex = []T{v}
		g.Arc = [][]int{{0}}
		return
	}
	if g.LocateVex(v) != VexNotFound {
		return
	}
	g.Vex = append(g.Vex, v)
	for i, _ := range g.Arc {
		g.Arc[i] = append(g.Arc[i], 0)
	}
	newRow := make([]int, len(g.Arc[0]))
	g.Arc = append(g.Arc, newRow)
}

func (g *WDigraph[T]) DeleteVex(v T) (res T) {
	if g.Vex == nil || len(g.Vex) == 0 || g.Arc == nil || len(g.Arc) == 0 {
		return
	}
	if index := g.LocateVex(v); index != VexNotFound {
		res = g.Vex[index]
		g.Vex = append(g.Vex[:index-1], g.Vex[index+1:]...)
		g.Arc = append(g.Arc[:index-1], g.Arc[index+1:]...)
		for i, _ := range g.Arc {
			g.Arc[i] = append(g.Arc[i][:index-1], g.Arc[i][index+1:]...)
		}
	}
	return
}

func (g *WDigraph[T]) InsertArc(v, w T, weight int) {
	vi := g.LocateVex(v)
	wi := g.LocateVex(w)
	if vi != VexNotFound && wi != VexNotFound {
		g.Arc[vi][wi] = weight
	}
}

func (g *WDigraph[T]) DeleteArc(v, w T) (weight int) {
	vi := g.LocateVex(v)
	wi := g.LocateVex(w)
	if vi == VexNotFound && wi == VexNotFound && vi == wi {
		return
	}
	weight = g.Arc[vi][wi]
	g.Arc[vi][wi] = INFINITY
	return
}

func (g *WDigraph[T]) BFSTraverse(start T, visit func(v T)) {
	if g.Vex == nil || len(g.Vex) == 0 || g.Arc == nil || len(g.Arc) == 0 {
		return
	}
	vi := g.LocateVex(start)
	if vi == VexNotFound {
		return
	}
	var zeroT T
	q := queue.NewSequenceQueue[T]()
	visited := make([]bool, len(g.Vex))
	for i, j := vi, 0; j < len(visited); j++ {
		if i == len(visited) {
			i = 0
		}
		if !visited[i] {
			v := g.Vex[i]
			visit(v)
			visited[i] = true
			q.EnQueue(v)
			for !q.IsEmpty() {
				v, _ = q.DeQueue()
				for w := g.FirstAdjVex(v); w != zeroT; w = g.NextAdjVex(v, w) {
					wi := g.LocateVex(w)
					if !visited[wi] {
						visit(w)
						visited[wi] = true
						q.EnQueue(w)
					}
				}
			}
		}
		i++
	}
}

func (g *WDigraph[T]) DFSTraverse(start T, visit func(v T)) {
	if g.Vex == nil || len(g.Vex) == 0 || g.Arc == nil || len(g.Arc) == 0 {
		return
	}
	vi := g.LocateVex(start)
	if vi == VexNotFound {
		return
	}
	visited := make([]bool, len(g.Vex))
	for i, j := vi, 0; j < len(visited); j++ {
		if i == len(visited) {
			i = 0
		}
		if !visited[i] {
			g.dfs(&visited, g.Vex[i], visit)
		}
		i++
	}
}

func (g *WDigraph[T]) dfs(visitedP *[]bool, v T, visit func(v T)) {
	visit(v)
	vi := g.LocateVex(v)
	(*visitedP)[vi] = true
	var zeroT T
	for w := g.FirstAdjVex(v); w != zeroT; w = g.NextAdjVex(v, w) {
		wi := g.LocateVex(w)
		if !(*visitedP)[wi] {
			g.dfs(visitedP, w, visit)
		}
	}

}
