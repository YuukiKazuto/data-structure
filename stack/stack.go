package stack

type Stack[T comparable] interface {
	IsEmpty() bool
	Push(data T)
	Pop() (T, error)
	GetTop() T
}
