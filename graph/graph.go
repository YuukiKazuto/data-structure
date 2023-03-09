package graph

type Graph[T comparable] interface {
	LocateVex(v T) int
	FirstAdjVex(v T) (res T)
	NextAdjVex(v, w T) (res T)
	InsertVex(v T)
	DeleteVex(v T) (res T)
	InsertArc(v, w T, weight int)
	DeleteArc(v, w T) (weight int)
	BFSTraverse(start T, visit func(v T))
	DFSTraverse(start T, visit func(v T))
}
