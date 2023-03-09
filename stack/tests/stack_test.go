package tests

import (
	"data-structure/stack"
	"testing"
)

var s stack.Stack[int]

func init() {
	s = stack.NewSLinkedStack[int]()
	//s = stack.NewLinkedStack[int]()
}
func BenchmarkPush(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
}

func BenchmarkPop(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}
