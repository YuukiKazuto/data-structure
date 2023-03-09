package stack

import (
	"fmt"
	"sync"
)

type SNode[T comparable] struct {
	Data T
	Next *Node[T]
}

type SLinkedStack[T comparable] struct {
	Head   *Node[T]
	Locker sync.Mutex
}

func NewSLinkedStack[T comparable]() *LinkedStack[T] {
	return &LinkedStack[T]{Head: &Node[T]{}}
}

func (s *SLinkedStack[T]) Push(data T) {
	s.Locker.Lock()
	defer s.Locker.Unlock()
	n := &Node[T]{Data: data}
	n.Next = s.Head.Next
	s.Head.Next = n
}

func (s *SLinkedStack[T]) Pop() (T, error) {
	s.Locker.Lock()
	defer s.Locker.Unlock()
	var ret T
	if s.IsEmpty() {
		return ret, fmt.Errorf("stack is empty")
	}
	top := s.Head.Next
	s.Head.Next = top.Next
	ret = top.Data
	top = nil
	return ret, nil
}

func (s *SLinkedStack[T]) IsEmpty() bool {
	if s.Head.Next == nil {
		return true
	} else {
		return false
	}
}

func (s *SLinkedStack[T]) GetTop() T {
	return s.Head.Next.Data
}
