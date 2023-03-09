package heap

type Heap interface {
	IsEmpty() bool
	Insert(data int)
	Build(arr []int)
}
