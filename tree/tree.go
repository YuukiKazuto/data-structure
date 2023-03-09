package tree

import (
	"data-structure/queue"
	"fmt"
)

type PNode[T comparable] struct {
	Data   T
	Parent int
}

// PTree 树的双亲表示法
type PTree[T comparable] struct {
	Nodes []PNode[T]
}

func NewPTree[T comparable](rData T) *PTree[T] {
	return &PTree[T]{
		Nodes: []PNode[T]{
			{rData, -1},
		},
	}
}

func (t *PTree[T]) IsFound(data T, parent int) (bool, int) {
	for i, node := range t.Nodes {
		if data == node.Data && parent == node.Parent {
			return true, i
		}
	}
	return false, -1
}

func (t *PTree[T]) AddNode(data T, parent int) {
	if parent == -1 {
		return
	}
	if found, _ := t.IsFound(data, parent); !found {
		t.Nodes = append(
			t.Nodes,
			PNode[T]{
				Data:   data,
				Parent: parent,
			},
		)
	}
}

func (t *PTree[T]) FindParent(index int) PNode[T] {
	return t.Nodes[t.Nodes[index].Parent]
}

type CTNode[T comparable] struct {
	Child int
	Next  *CTNode[T]
}

type CTBox[T comparable] struct {
	Data       T
	FirstChild *CTNode[T]
}

type CTree[T comparable] struct {
	Nodes []CTBox[T]
}

func NewCTree[T comparable](rData T) *CTree[T] {
	return &CTree[T]{
		Nodes: []CTBox[T]{
			{rData, nil},
		},
	}
}

func (t *CTree[T]) IsFound(data T, parent int) (bool, int) {
	if parent == -1 {
		return false, -1
	}
	p := t.Nodes[parent].FirstChild
	for p != nil {
		if data == t.Nodes[p.Child].Data {
			return true, p.Child
		}
		p = p.Next
	}
	return false, -1
}

func (t *CTree[T]) AddNode(data T, parent int) {
	if parent == -1 {
		return
	}
	if found, _ := t.IsFound(data, parent); !found {
		size := len(t.Nodes)
		t.Nodes = append(
			t.Nodes,
			CTBox[T]{
				Data: data,
			},
		)
		c := &CTNode[T]{
			Child: size,
		}
		p := t.Nodes[parent].FirstChild
		if p == nil {
			t.Nodes[parent].FirstChild = c
			return
		}
		for p.Next != nil {
			p = p.Next
		}
		p.Next = c
	}
}

// DeleteNode
//TODO: Fix bugs Or abandon
func (t *CTree[T]) DeleteNode(index int) {
	size := len(t.Nodes)
	if index >= size || index < 0 {
		return
	}
	p := t.Nodes[index].FirstChild
	if p != nil {
		for p != nil {
			t.DeleteNode(p.Child)
			p = p.Next
		}
	} else {
		t.Nodes[index] = t.Nodes[size-1]
		t.Nodes = t.Nodes[:size-1]
	}
}

type CSNode[T comparable] struct {
	Data                    T
	FirstChild, NextSibling *CSNode[T]
}

type CSTree[T comparable] struct {
	Root *CSNode[T]
}

// Preorder 树的先根遍历=森林的先序遍历=二叉树的先序遍历
func (n *CSNode[T]) Preorder() {
	if n != nil {
		fmt.Printf("%v ", n.Data)
		n.FirstChild.Preorder()
		n.NextSibling.Preorder()
	}
}
func (t *CSTree[T]) Preorder() {
	t.Root.Preorder()
}

// Postorder 树的后根遍历=森林的中序遍历=二叉树的中序遍历
func (n *CSNode[T]) Postorder() {
	if n != nil {
		n.FirstChild.Postorder()
		fmt.Printf("%v ", n.Data)
		n.NextSibling.Postorder()
	}
}
func (t *CSTree[T]) Postorder() {
	t.Root.Postorder()
}

func (t *CSTree[T]) Level() {
	q := queue.NewSequenceQueue[*CSNode[T]]()
	q.EnQueue(t.Root)
	for !q.IsEmpty() {
		node, _ := q.DeQueue()
		fmt.Printf("%v ", node.Data)
		if node.FirstChild != nil {
			q.EnQueue(node.FirstChild)
			node = node.FirstChild
			for node.NextSibling != nil {
				q.EnQueue(node.NextSibling)
				node = node.NextSibling
			}
		}
	}
}
