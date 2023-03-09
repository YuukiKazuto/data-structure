package queue

import "fmt"

type Node[T comparable] struct {
	Data T
	Next *Node[T]
}

type LinkedQueue[T comparable] struct {
	Front, Rear *Node[T]
	Length      int
}

func NewLinkedQueue[T comparable]() *LinkedQueue[T] {
	front := &Node[T]{}
	rear := front
	return &LinkedQueue[T]{
		Front: front,
		Rear:  rear,
	}
}

func (q *LinkedQueue[T]) IsEmpty() bool {
	if q.Front == q.Rear {
		return true
	} else {
		return false
	}
}

func (q *LinkedQueue[T]) EnQueue(data T) {
	n := &Node[T]{Data: data}
	q.Rear.Next = n
	q.Rear = n
	q.Length++
}

func (q *LinkedQueue[T]) DeQueue() (T, error) {
	var ret T
	if q.IsEmpty() {
		return ret, fmt.Errorf("queue is empty")
	}
	p := q.Front.Next
	ret = p.Data
	q.Front.Next = p.Next
	q.Length--
	if q.Rear == p {
		q.Rear = q.Front
	}
	return ret, nil
}
