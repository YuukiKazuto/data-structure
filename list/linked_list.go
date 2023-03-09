package list

import "fmt"

type Node[T comparable] struct {
	Data T
	Next *Node[T]
}

func (n *Node[T]) InsertNextNode(data T) error {
	if n == nil {
		return fmt.Errorf("error of node is empty")
	}
	next := &Node[T]{
		Data: data,
		Next: n.Next,
	}
	n.Next = next
	return nil
}

func (n *Node[T]) InsertPriorNode(data T) error {
	if n == nil {
		return fmt.Errorf("error of node is empty")
	}
	// p复制节点n
	p := &Node[T]{
		Data: n.Data,
		Next: n.Next,
	}
	n.Data = data
	n.Next = p
	return nil
}

type LinkedList[T comparable] struct {
	Head   *Node[T]
	Length int
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{Head: &Node[T]{}}
}

func (l *LinkedList[T]) IsEmpty() bool {
	if l.Head.Next == nil {
		return true
	} else {
		return false
	}
}

func (l *LinkedList[T]) Insert(i int, data T) error {
	if i < 1 {
		return fmt.Errorf("error of index: %d", i)
	}
	if i > l.Length+1 {
		return fmt.Errorf("error of index is greater than the list length+1")
	}
	n := l.GetNode(i - 1)
	if err := n.InsertNextNode(data); err != nil {
		return err
	}
	l.Length++
	return nil
}

func (l *LinkedList[T]) Delete(i int) (T, error) {
	var ret T
	if i < 1 {
		return ret, fmt.Errorf("error of index: %d", i)
	}
	if i > l.Length {
		return ret, fmt.Errorf("error of index is greater than the list length")
	}
	n := l.GetNode(i - 1)
	p := n.Next
	ret = p.Data
	n.Next = p.Next
	l.Length--
	return ret, nil
}

func (l *LinkedList[T]) Add(data T) {
	n := l.GetNode(l.Length)
	_ = n.InsertNextNode(data)
	l.Length++
}
func (l *LinkedList[T]) InsertNextNode(n *Node[T], data T) error {
	if err := n.InsertNextNode(data); err != nil {
		return err
	}
	l.Length++
	return nil
}

func (l *LinkedList[T]) InsertPriorNode(n *Node[T], data T) error {
	if err := n.InsertPriorNode(data); err != nil {
		return err
	}
	l.Length++
	return nil
}

func (l *LinkedList[T]) Get(i int) (T, error) {
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

func (l *LinkedList[T]) GetNode(i int) *Node[T] {
	if i < 0 {
		return nil
	}
	n := l.Head
	for j := 0; n != nil && j < i; j++ {
		n = n.Next
	}
	return n
}
func (l *LinkedList[T]) Locate(data T) *Node[T] {
	n := l.Head.Next
	for n != nil && n.Data != data {
		n = n.Next
	}
	return n
}

func (l *LinkedList[T]) Index(data T) int {
	n := l.Head
	i := 0
	for ; n != nil && n.Data != data; i++ {
		n = n.Next
	}
	return i
}

func (l *LinkedList[T]) Size() int {
	return l.Length
}

func (l *LinkedList[T]) Push(data T) {
	n := &Node[T]{Data: data}
	n.Next = l.Head.Next
	l.Head.Next = n
	l.Length++
}

func (l *LinkedList[T]) Pop() (T, error) {
	var ret T
	if l.IsEmpty() {
		return ret, fmt.Errorf("stack is empty")
	}
	top := l.Head.Next
	l.Head.Next = top.Next
	l.Length--
	ret = top.Data
	return ret, nil
}
