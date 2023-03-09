package adjacency_list

import "data-structure/queue"

const VexNotFound = -1

const (
	Undigraph = iota
	DiGraph
)

type (
	ArcNode struct {
		AdjVex int
		Next   *ArcNode
		Weight int
	}
	VNode[T comparable] struct {
		Data  T
		First *ArcNode
	}
	ALGraph[T comparable] struct {
		AdjList []VNode[T]
		Kind    int // 0表示无向图，1表示有向图
	}
)

func NewALGraphAllArgument[T comparable](adjList []VNode[T], kind int) *ALGraph[T] {
	return &ALGraph[T]{
		AdjList: adjList,
		Kind:    kind,
	}
}

func NewALGraph[T comparable](adjList []VNode[T]) *ALGraph[T] {
	return &ALGraph[T]{AdjList: adjList}
}

func (g *ALGraph[T]) LocateVex(v T) int {
	for i, vex := range g.AdjList {
		if v == vex.Data {
			return i
		}
	}
	return VexNotFound
}

func (g *ALGraph[T]) FirstAdjVex(v T) (res T) {
	if vi := g.LocateVex(v); vi != VexNotFound && g.AdjList[vi].First != nil {
		first := g.AdjList[vi].First
		res = g.AdjList[first.AdjVex].Data
		return
	}
	return
}

func (g *ALGraph[T]) NextAdjVex(v, w T) (res T) {
	vi := g.LocateVex(v)
	wi := g.LocateVex(w)
	if vi != VexNotFound && wi != VexNotFound {
		for arc := g.AdjList[vi].First; arc != nil; arc = arc.Next {
			if arc.AdjVex == wi && arc.Next != nil {
				res = g.AdjList[arc.Next.AdjVex].Data
				return
			}
		}
	}
	return
}

func (g *ALGraph[T]) InsertVex(v T) {
	if g.LocateVex(v) != VexNotFound {
		return
	}
	g.AdjList = append(g.AdjList, VNode[T]{Data: v})
}

func (g *ALGraph[T]) DeleteVex(v T) (res T) {
	if vi := g.LocateVex(v); vi != VexNotFound {
		res = g.AdjList[vi].Data
		g.AdjList = append(g.AdjList[:vi], g.AdjList[vi+1:]...)
		for i := range g.AdjList {
			arc := g.AdjList[i].First
			if arc != nil && arc.AdjVex == vi {
				g.AdjList[i].First = arc.Next
			} else {
				pre := arc
				arc = arc.Next
				for ; arc != nil; arc = arc.Next {
					if arc.AdjVex == vi {
						pre.Next = arc.Next
						break
					}
					pre = arc
				}
			}
		}
	}
	return
}

func (g *ALGraph[T]) InsertArc(v, w T, weight int) {
	vi := g.LocateVex(v)
	wi := g.LocateVex(w)
	if vi != VexNotFound && wi != VexNotFound {
		if g.Kind == DiGraph {
			g.digraphInsertArc(vi, wi, weight)
		} else {
			g.undigraphInsertArc(vi, wi, weight)
		}
	}
}

func (g *ALGraph[T]) DeleteArc(v, w T) (weight int) {
	vi := g.LocateVex(v)
	wi := g.LocateVex(w)
	if vi != VexNotFound && wi != VexNotFound {
		if g.Kind == DiGraph {
			weight = g.digraphDeleteArc(vi, wi)
		} else {
			weight = g.undigraphDeleteArc(vi, wi)
		}
	}
	return
}

func (g *ALGraph[T]) BFSTraverse(start T, visit func(v T)) {
	if g.AdjList == nil || len(g.AdjList) == 0 {
		return
	}
	vi := g.LocateVex(start)
	if vi == VexNotFound {
		return
	}
	var zeroT T
	q := queue.NewSequenceQueue[T]()
	visited := make([]bool, len(g.AdjList))
	for i, j := vi, 0; j < len(visited); j++ {
		if i == len(visited) {
			i = 0
		}
		if !visited[i] {
			v := g.AdjList[i].Data
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

func (g *ALGraph[T]) DFSTraverse(start T, visit func(v T)) {
	if g.AdjList == nil || len(g.AdjList) == 0 {
		return
	}
	vi := g.LocateVex(start)
	if vi == VexNotFound {
		return
	}
	visited := make([]bool, len(g.AdjList))
	for i, j := vi, 0; j < len(visited); j++ {
		if i == len(visited) {
			i = 0
		}
		if !visited[i] {
			g.dfs(&visited, g.AdjList[i].Data, visit)
		}
		i++
	}
}

func (g *ALGraph[T]) dfs(visitedP *[]bool, v T, visit func(v T)) {
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

func (g *ALGraph[T]) digraphInsertArc(vi int, wi int, weight int) {
	if g.AdjList[vi].First != nil {
		var last *ArcNode
		for arc := g.AdjList[vi].First; arc != nil; arc = arc.Next {
			if arc.AdjVex == wi {
				return
			}
			last = arc
		}
		last.Next = &ArcNode{
			AdjVex: wi,
			Weight: weight,
		}
	} else {
		g.AdjList[vi].First = &ArcNode{
			AdjVex: wi,
			Weight: weight,
		}
	}
}

func (g *ALGraph[T]) undigraphInsertArc(vi int, wi int, weight int) {
	g.digraphInsertArc(vi, wi, weight)
	if g.AdjList[wi].First != nil {
		var last *ArcNode
		for arc := g.AdjList[wi].First; arc != nil; arc = arc.Next {
			if arc.AdjVex == vi {
				return
			}
			last = arc
		}
		last.Next = &ArcNode{
			AdjVex: vi,
			Weight: weight,
		}
	} else {
		g.AdjList[wi].First = &ArcNode{
			AdjVex: vi,
			Weight: weight,
		}
	}
}

func (g *ALGraph[T]) digraphDeleteArc(vi int, wi int) int {
	arc := g.AdjList[vi].First
	if arc != nil && arc.Next == nil && arc.AdjVex == wi {
		g.AdjList[vi].First = nil
		return arc.Weight
	}
	pre := arc
	for arc = arc.Next; arc != nil; arc = arc.Next {
		if arc.AdjVex == wi {
			pre.Next = arc.Next
			return arc.Weight
		}
		pre = arc
	}
	return 0
}

func (g *ALGraph[T]) undigraphDeleteArc(vi int, wi int) (weight int) {
	weight = g.digraphDeleteArc(vi, wi)
	arc := g.AdjList[wi].First
	if arc != nil && arc.Next == nil && arc.AdjVex == vi {
		g.AdjList[wi].First = nil
		return
	}
	pre := arc
	for arc = arc.Next; arc != nil; arc = arc.Next {
		if arc.AdjVex == vi {
			pre.Next = arc.Next
			return
		}
		pre = arc
	}
	return
}
