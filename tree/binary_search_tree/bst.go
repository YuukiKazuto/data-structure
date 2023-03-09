package binary_search_tree

import (
	"data-structure/stack"
	"fmt"
)

type BSTNode struct {
	Key            int
	LChild, RChild *BSTNode
}

func (n *BSTNode) inorder() {
	if n != nil {
		n.LChild.inorder()
		fmt.Printf("%d ", n.Key)
		n.RChild.inorder()
	}
}

// BSTree 二叉排序树
// ASL: 平均查找长度
type BSTree struct {
	Root *BSTNode
}

func (t *BSTree) Search(key int) *BSTNode {
	node := t.Root
	for node != nil && key != node.Key {
		if key < node.Key {
			node = node.LChild
		} else {
			node = node.RChild
		}
	}
	return node
}

func (t *BSTree) InsertNode(key int) bool {
	if t.Root == nil {
		t.Root = &BSTNode{Key: key}
		return true
	}
	node := t.Root
	for node != nil && key != node.Key {
		if key < node.Key {
			if node.LChild == nil {
				node.LChild = &BSTNode{Key: key}
				return true
			}
			node = node.LChild
		} else {
			if node.RChild == nil {
				node.RChild = &BSTNode{Key: key}
				return true
			}
			node = node.RChild
		}
	}
	return false
}

// DeleteNode 删除
func (t *BSTree) DeleteNode(key int) bool {
	node := t.Root
	var pre *BSTNode
	for node != nil {
		if key < node.Key {
			pre = node
			node = node.LChild
		} else if key > node.Key {
			pre = node
			node = node.RChild
		} else {
			if node.RChild == nil {
				if pre == nil {
					t.Root = node.LChild
					return true
				}
				if node == pre.LChild {
					pre.LChild = node.LChild
					return true
				}
				if node == pre.RChild {
					pre.RChild = node.LChild
					return true
				}
			} else if node.LChild == nil {
				if pre == nil {
					t.Root = node.RChild
					return true
				}
				if node == pre.LChild {
					pre.LChild = node.RChild
					return true
				}
				if node == pre.RChild {
					pre.RChild = node.RChild
					return true
				}
			} else {
				if min, minPre := FindRChildMinNode(node.RChild); minPre != nil {
					minPre.LChild = min.RChild
					min.LChild = node.LChild
					min.RChild = node.RChild
					if pre == nil {
						t.Root = min
						return true
					}
					if node == pre.LChild {
						pre.LChild = min
						return true
					}
					if node == pre.RChild {
						pre.RChild = min
						return true
					}
				} else {
					if pre == nil {
						min.LChild = node.LChild
						min.RChild = node.RChild
						t.Root = min
						return true
					}
					if node == pre.LChild {
						min.LChild = node.LChild
						pre.LChild = min
						return true
					}
					if node == pre.RChild {
						min.LChild = node.LChild
						pre.RChild = min
						return true
					}
				}
			}
		}
	}
	return false
}

func (t *BSTree) Create(arr []int) {
	t.Root = nil
	for _, v := range arr {
		t.InsertNode(v)
	}
}

func (t *BSTree) Inorder() {
	t.Root.inorder()
}

func FindRChildMinNode(n *BSTNode) (min *BSTNode, pre *BSTNode) {
	s := stack.NewSequenceStack[*BSTNode]()
	for n != nil || !s.IsEmpty() {
		if n != nil {
			s.Push(n)
			n = n.LChild
		} else {
			min, _ = s.Pop()
			pre, _ = s.Pop()
			return
		}
	}
	min = n
	return
}
