package list

import "fmt"

type DNode[T comparable] struct {
	Data        T
	Prior, Next *DNode[T]
}

func (d *DNode[T]) InsertNextNode(n *DNode[T]) {
	if d == nil || n == nil {
		return
	}
	n.Next = d.Next
	if d.Next != nil {
		d.Next.Prior = n
	}
	n.Prior = d
	d.Next = n
}

func (d *DNode[T]) DeleteNextNode() {
	if d == nil || d.Next == nil {
		return
	}
	n := d.Next
	d.Next = n.Next
	if n.Next != nil {
		n.Next.Prior = d
	}
}

type DoubleLinkedList[T comparable] struct {
	Head   *DNode[T]
	Length int
}

func NewDoubleLinkedList[T comparable]() *DoubleLinkedList[T] {
	return &DoubleLinkedList[T]{Head: &DNode[T]{}}
}

func Destroy[T comparable](l *DoubleLinkedList[T]) {
	for l.Head.Next != nil {
		l.Head.DeleteNextNode()
	}
	l = nil
}

func (l *DoubleLinkedList[T]) IsEmpty() bool {
	if l.Head.Next == nil {
		return true
	} else {
		return false
	}
}

func (l *DoubleLinkedList[T]) Size() int {
	return l.Length
}

func (l *DoubleLinkedList[T]) GetNode(i int) *DNode[T] {
	if i < 0 {
		return nil
	}
	n := l.Head
	for j := 0; n != nil && j < i; j++ {
		n = n.Next
	}
	return n
}

func (l *DoubleLinkedList[T]) Get(i int) (T, error) {
	var ret T
	if i < 1 {
		return ret, fmt.Errorf("error of index: %d", i)
	}
	if i > l.Length {
		return ret, fmt.Errorf("error of index is greater than the list length")
	}
	n := l.GetNode(i)
	ret = n.Data
	return ret, nil
}

func (l *DoubleLinkedList[T]) Insert(i int, data T) error {
	if i < 1 {
		return fmt.Errorf("error of index: %d", i)
	}
	if i > l.Length+1 {
		return fmt.Errorf("error of index is greater than the list length+1")
	}
	n := l.GetNode(i - 1)
	n.InsertNextNode(&DNode[T]{Data: data})
	l.Length++
	return nil
}

func (l *DoubleLinkedList[T]) Delete(i int) (T, error) {
	var ret T
	if i < 1 {
		return ret, fmt.Errorf("error of index: %d", i)
	}
	if i > l.Length {
		return ret, fmt.Errorf("error of index is greater than the list length")
	}
	n := l.GetNode(i - 1)
	ret = n.Next.Data
	n.DeleteNextNode()
	l.Length--
	return ret, nil
}

func (l *DoubleLinkedList[T]) Add(data T) {
	n := l.GetNode(l.Length)
	n.InsertNextNode(&DNode[T]{Data: data})
	l.Length++
}
