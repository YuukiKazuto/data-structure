package stack

import "fmt"

type Node[T comparable] struct {
	Data T
	Next *Node[T]
}

type LinkedStack[T comparable] struct {
	Head *Node[T]
}

func NewLinkedStack[T comparable]() *LinkedStack[T] {
	return &LinkedStack[T]{Head: &Node[T]{}}
}

func (s *LinkedStack[T]) Push(data T) {
	n := &Node[T]{Data: data}
	n.Next = s.Head.Next
	s.Head.Next = n
}

func (s *LinkedStack[T]) Pop() (T, error) {
	var ret T
	if s.IsEmpty() {
		return ret, fmt.Errorf("stack is empty")
	}
	top := s.Head.Next
	s.Head.Next = top.Next
	ret = top.Data
	return ret, nil
}

func (s *LinkedStack[T]) IsEmpty() bool {
	if s.Head.Next == nil {
		return true
	} else {
		return false
	}
}

func (s *LinkedStack[T]) GetTop() T {
	return s.Head.Next.Data
}
