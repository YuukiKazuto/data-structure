package binary_tree

import "fmt"

func (n *TNode[T]) PostorderPreNode() *TNode[T] {
	if n.LTag == 0 {
		if n.RTag == 0 {
			return n.RChild
		}
		return n.LChild
	}
	return n.LChild
}

type PostorderThreadTree[T comparable] ThreadTree[T]

func (t *ThreadTree[T]) CreatePostThread() *PostorderThreadTree[T] {
	var pre *TNode[T]
	if t != nil {
		t.Root.PostThread(&pre)
		if pre.RChild == nil {
			pre.RTag = 1
		}
	}
	return (*PostorderThreadTree[T])(t)
}

func (t *PostorderThreadTree[T]) RevPostorder() {
	for p := t.Root; p != nil; p = p.PostorderPreNode() {
		fmt.Printf("%v ", p.Data)
	}
}
