package compressed

import "data-structure/queue"

const VexNotFound = -1

type WGraph[T comparable] struct {
	Vex  []T
	Edge []int
}

func NewWGraph[T comparable](vex []T, edge []int) *WGraph[T] {
	return &WGraph[T]{Vex: vex, Edge: edge}
}

func (g *WGraph[T]) LocateVex(v T) int {
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

func (g *WGraph[T]) FirstAdjVex(v T) (res T) {
	if index := g.LocateVex(v); index != VexNotFound {
		for i, j := (index+1)*index/2, 0; i < (index+2)*(index+1)/2; i++ {
			if g.Edge[i] != 0 {
				res = g.Vex[j]
				return
			}
			j++
		}
		for i, vn := (index+2)*(index+1)/2, index+1; i < len(g.Edge); i += 1 + (vn - 1) {
			if g.Edge[i+index] != 0 {
				res = g.Vex[vn]
				return
			}
			vn++
		}
	}
	return
}

func (g *WGraph[T]) NextAdjVex(v, w T) (res T) {
	vi := g.LocateVex(v)
	wi := g.LocateVex(w)
	if vi != VexNotFound && wi != VexNotFound && wi != len(g.Vex)-1 {
		if wi < vi && g.Edge[(vi+1)*vi/2+wi] != 0 {
			for i, j := (vi+1)*vi/2+wi+1, wi+1; i < (vi+2)*(vi+1)/2; i++ {
				if g.Edge[i] != 0 {
					res = g.Vex[j]
					return
				}
				j++
			}
			for i, vn := (vi+2)*(vi+1)/2, vi+1; i < len(g.Edge); i += 1 + (vn - 1) {
				if g.Edge[i+vi] != 0 {
					res = g.Vex[vn]
					return
				}
				vn++
			}
		} else if wi > vi && g.Edge[(wi+1)*wi/2+vi] != 0 {
			for i, vn := (wi+2)*(wi+1)/2, wi+1; i < len(g.Edge); i += 1 + (vn - 1) {
				if g.Edge[i+vi] != 0 {
					res = g.Vex[vn]
					return
				}
				vn++
			}
		}
	}
	return
}

func (g *WGraph[T]) InsertVex(v T) {
	if g.Vex == nil || g.Edge == nil || len(g.Edge) == 0 {
		g.Vex = []T{v}
		g.Edge = []int{0}
		return
	}
	if g.LocateVex(v) != VexNotFound {
		return
	}
	g.Vex = append(g.Vex, v)
	newVexRow := make([]int, len(g.Vex))
	g.Edge = append(g.Edge, newVexRow...)
}

func (g *WGraph[T]) DeleteVex(v T) (res T) {
	if g.Vex == nil || len(g.Vex) == 0 || g.Edge == nil || len(g.Edge) == 0 {
		return
	}
	if index := g.LocateVex(v); index != VexNotFound {
		res = g.Vex[index]
		g.Edge = append(g.Edge[:(index+1)*index/2], g.Edge[(index+2)*(index+1)/2:]...)
		for i, vn := (index+2)*(index+1)/2, index+1; i < len(g.Edge); i += 1 + (vn - 1) {
			g.Edge = append(g.Edge[:i+index], g.Edge[i+index+1:]...)
			vn++
		}
	}
	return
}

func (g *WGraph[T]) InsertArc(v, w T, weight int) {
	vi := g.LocateVex(v)
	wi := g.LocateVex(w)
	if vi != VexNotFound && wi != VexNotFound {
		if vi > wi {
			g.Edge[(vi+1)*vi/2+wi] = weight
		} else {
			g.Edge[(wi+1)*wi/2+vi] = weight
		}
	}
}

func (g *WGraph[T]) DeleteArc(v, w T) (weight int) {
	vi := g.LocateVex(v)
	wi := g.LocateVex(w)
	if vi != VexNotFound && wi != VexNotFound {
		if vi > wi {
			weight = g.Edge[(vi+1)*vi/2+wi]
			g.Edge[(vi+1)*vi/2+wi] = 0
		} else {
			weight = g.Edge[(wi+1)*wi/2+vi]
			g.Edge[(wi+1)*wi/2+vi] = 0
		}
	}
	return
}

func (g *WGraph[T]) BFSTraverse(start T, visit func(v T)) {
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

func (g *WGraph[T]) DFSTraverse(start T, visit func(v T)) {
	if g.Vex == nil || len(g.Vex) == 0 || g.Edge == nil || len(g.Edge) == 0 {
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

func (g *WGraph[T]) dfs(visitedP *[]bool, v T, visit func(v T)) {
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
