package adjacency_multilist

type ENode struct {
	I      int
	J      int
	weight int
	ILink  *ENode
	JLink  *ENode
}

type AMVNode[T comparable] struct {
	Data  T
	First *ENode
}
