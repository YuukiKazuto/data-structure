package list

import "fmt"

type ArrayList[T comparable] struct {
	Slice []T
}

func NewArrayList[T comparable]() *ArrayList[T] {
	return &ArrayList[T]{Slice: make([]T, 0, 0)}
}

func (l *ArrayList[T]) IsEmpty() bool {
	if len(l.Slice) == 0 {
		return true
	} else {
		return false
	}
}

func (l *ArrayList[T]) Size() int {
	return len(l.Slice)
}

func (l *ArrayList[T]) Add(data T) {
	l.Slice = append(l.Slice, data)
}

func (l *ArrayList[T]) Insert(i int, data T) error {
	if i < 0 {
		return fmt.Errorf("error of index: %d", i)
	}
	if i > l.Size() {
		return fmt.Errorf("error of index is greater than the list size")
	}
	l.Slice = append(l.Slice, data)
	for j := l.Size() - 1; j > i; j-- {
		l.Slice[j] = l.Slice[j-1]
	}
	l.Slice[i] = data
	return nil
}

func (l *ArrayList[T]) Delete(i int) (T, error) {
	var ret T
	if i < 0 {
		return ret, fmt.Errorf("error of index: %d", i)
	}
	if i >= l.Size() {
		return ret, fmt.Errorf("error of index is greater than the list size")
	}
	ret = l.Slice[i]
	l.Slice = append(l.Slice[:i], l.Slice[i+1:]...)
	return ret, nil
}
func (l *ArrayList[T]) Get(i int) (T, error) {
	var ret T
	if i < 0 {
		return ret, fmt.Errorf("error of index: %d", i)
	}
	if i >= l.Size() {
		return ret, fmt.Errorf("error of index is greater than the list size")
	}
	ret = l.Slice[i]
	return ret, nil
}

func (l *ArrayList[T]) Index(data T) int {
	i := -1
	for j := 0; j < l.Size(); j++ {
		if l.Slice[j] == data {
			i = j
			break
		}
	}
	return i
}
