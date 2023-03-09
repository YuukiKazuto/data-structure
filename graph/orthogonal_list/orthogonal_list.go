package orthogonal_list

type ArcNode struct {
	Tail         int
	Head         int
	Weight       int
	HLink, TLink *ArcNode
}

type OLVNode[T comparable] struct {
	Data              T
	FirstIn, FirstOut *ArcNode
}

type OLGraph[T comparable] struct {
	Vertexes OLVNode[T]
}
