package binary_tree

import "fmt"

func (n *TNode[T]) PreorderFirstNode() *TNode[T] {
	res := n
	for res.RTag == 0 {
		if res.LTag == 0 {
			return res.LChild
		}
		res = res.RChild
	}
	return res
}

func (n *TNode[T]) PreorderLastNode() *TNode[T] {
	res := n
	for res.RTag == 0 {
		res = res.RChild
	}
	return res
}

func (n *TNode[T]) PreorderNextNode() *TNode[T] {
	if n.LTag == 0 {
		return n.LChild
	}
	if n.RTag == 0 {
		return n.RChild.PreorderFirstNode()
	}
	return n.RChild
}

type PreorderThreadTree[T comparable] ThreadTree[T]

func (t *ThreadTree[T]) CreatePreThread() *PreorderThreadTree[T] {
	var pre *TNode[T]
	if t != nil {
		t.Root.PreThread(&pre)
		if pre.RChild == nil {
			pre.RTag = 1
		}
	}
	return (*PreorderThreadTree[T])(t)
}

func (t *PreorderThreadTree[T]) Preorder() {
	for p := t.Root; p != nil; p = p.PreorderNextNode() {
		fmt.Printf("%v ", p.Data)
	}
}
