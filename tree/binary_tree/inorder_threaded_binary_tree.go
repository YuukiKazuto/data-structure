package binary_tree

import "fmt"

func (n *TNode[T]) InorderFirstNode() *TNode[T] {
	res := n
	for res.LTag == 0 {
		res = res.LChild
	}
	return res
}

func (n *TNode[T]) InorderLastNode() *TNode[T] {
	res := n
	for res.RTag == 0 {
		res = res.RChild
	}
	return res
}

func (n *TNode[T]) InorderNextNode() *TNode[T] {
	if n.RTag == 0 {
		return n.RChild.InorderFirstNode()
	}
	return n.RChild
}

func (n *TNode[T]) InorderPreNode() *TNode[T] {
	if n.LTag == 0 {
		return n.LChild.InorderLastNode()
	}
	return n.LChild
}

type InorderThreadTree[T comparable] ThreadTree[T]

func (t *ThreadTree[T]) CreateInThread() *InorderThreadTree[T] {
	var pre *TNode[T]
	if t != nil {
		t.Root.InThread(&pre)
		if pre.RChild == nil {
			pre.RTag = 1
		}
	}
	return (*InorderThreadTree[T])(t)
}

func (t *InorderThreadTree[T]) Inorder() {
	for p := t.Root.InorderFirstNode(); p != nil; p = p.InorderNextNode() {
		fmt.Printf("%v ", p.Data)
	}
}

func (t *InorderThreadTree[T]) RevInorder() {
	for p := t.Root.InorderLastNode(); p != nil; p = p.InorderPreNode() {
		fmt.Printf("%v ", p.Data)
	}
}
