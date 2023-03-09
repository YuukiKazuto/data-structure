package stack

import "fmt"

type SequenceStack[T comparable] struct {
	Slice []T
}

func NewSequenceStack[T comparable]() *SequenceStack[T] {
	return &SequenceStack[T]{Slice: make([]T, 0, 0)}
}

func (s *SequenceStack[T]) IsEmpty() bool {
	if len(s.Slice) == 0 {
		return true
	} else {
		return false
	}
}

func (s *SequenceStack[T]) Push(data T) {
	s.Slice = append(s.Slice, data)
}

func (s *SequenceStack[T]) Pop() (T, error) {
	var res T
	if s.IsEmpty() {
		return res, fmt.Errorf("stack is empty")
	}
	sLen := len(s.Slice)
	res = s.Slice[sLen-1]
	s.Slice = s.Slice[:sLen-1]
	return res, nil
}

func (s *SequenceStack[T]) GetTop() T {
	var res T
	if s.IsEmpty() {
		return res
	}
	sLen := len(s.Slice)
	res = s.Slice[sLen-1]
	return res
}

func (s *SequenceStack[T]) FindFromTop(data T) bool {
	for i := len(s.Slice) - 1; i >= 0; i-- {
		if s.Slice[i] == data {
			return true
		}
	}
	return false
}
