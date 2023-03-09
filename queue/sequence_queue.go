package queue

import "fmt"

type SequenceQueue[T comparable] struct {
	Slice []T
}

func NewSequenceQueue[T comparable]() *SequenceQueue[T] {
	return &SequenceQueue[T]{Slice: make([]T, 0, 0)}
}

func (q *SequenceQueue[T]) IsEmpty() bool {
	if len(q.Slice) == 0 {
		return true
	} else {
		return false
	}
}

func (q *SequenceQueue[T]) EnQueue(data T) {
	q.Slice = append(q.Slice, data)
}

func (q *SequenceQueue[T]) DeQueue() (T, error) {
	var ret T
	if q.IsEmpty() {
		return ret, fmt.Errorf("queue is empty")
	}
	ret = q.Slice[0]
	q.Slice = q.Slice[1:]
	return ret, nil
}
