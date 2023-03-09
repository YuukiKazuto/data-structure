package binary_tree

import (
	"data-structure/stack"
	"fmt"
)

type Node[T comparable] struct {
	Data   T        `json:"data"`
	LChild *Node[T] `json:"l_child"`
	RChild *Node[T] `json:"r_child"`
}

func (n *Node[T]) preorder() {
	if n != nil {
		fmt.Printf("%v ", n.Visit())
		n.LChild.preorder()
		n.RChild.preorder()
	}
}

func (n *Node[T]) inorder() {
	if n != nil {
		n.LChild.inorder()
		fmt.Printf("%v ", n.Visit())
		//n.Visit()
		n.RChild.inorder()
	}
}

func (n *Node[T]) postorder() {
	if n != nil {
		n.LChild.postorder()
		n.RChild.postorder()
		fmt.Printf("%v ", n.Visit())
	}
}

func (n *Node[T]) Visit() T {
	return n.Data
}

func (n *Node[T]) depth() int {
	if n == nil {
		return 0
	} else {
		l := n.LChild.depth()
		r := n.RChild.depth()
		if l > r {
			return l + 1
		} else {
			return r + 1
		}
	}
}

type BiTree[T comparable] struct {
	Root *Node[T]
}

func NewBiTree[T comparable]() *BiTree[T] {
	return &BiTree[T]{}
}

func (t *BiTree[T]) AddNode(data T) {
	if t.Root == nil {
		t.Root = &Node[T]{Data: data}
		return
	}
	queue := []*Node[T]{t.Root}
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur.LChild == nil {
			cur.LChild = &Node[T]{Data: data}
			return
		} else if cur.RChild == nil {
			cur.RChild = &Node[T]{Data: data}
			return
		} else {
			queue = append(queue, cur.LChild)
			queue = append(queue, cur.RChild)
		}
	}
}

func (t *BiTree[T]) Level() {
	if t.Root != nil {
		queue := []*Node[T]{t.Root}
		for len(queue) != 0 {
			cur := queue[0]
			queue = queue[1:]
			fmt.Printf("%v ", cur.Visit())
			if cur.LChild != nil {
				queue = append(queue, cur.LChild)
			}
			if cur.RChild != nil {
				queue = append(queue, cur.RChild)
			}
		}
	}
}

func (t *BiTree[T]) PreorderRecursion() {
	t.Root.preorder()
}

// Preorder 非递归方法实现先序遍历
func (t *BiTree[T]) Preorder() {
	s := stack.NewLinkedStack[*Node[T]]()
	node := t.Root
	for node != nil || !s.IsEmpty() {
		if node != nil {
			fmt.Printf("%v ", node.Visit())
			s.Push(node)
			node = node.LChild
		} else {
			top, _ := s.Pop()
			node = top.RChild
		}
	}
}

// Preorder1 非递归方法实现先序遍历（切片模拟栈）
//func (t *BiTree[T]) Preorder1() {
//	sliceStack := make([]*Node[T], 0, 10)
//	top := 0 // 栈顶指针
//	node := t.Root
//	for node != nil || len(sliceStack) != 0 {
//		if node != nil {
//			fmt.Printf("%v ", node.Visit())
//			sliceStack = append(sliceStack, node)
//			top++
//			node = node.LChild
//		} else {
//			topNode := sliceStack[top-1]    // 获取栈顶元素
//			sliceStack = sliceStack[:top-1] // 出栈
//			top--
//			node = topNode.RChild
//		}
//	}
//}

func (t *BiTree[T]) InorderRecursion() {
	t.Root.inorder()
}

// Inorder 非递归方法实现中序遍历
func (t *BiTree[T]) Inorder() {
	s := stack.NewSequenceStack[*Node[T]]()
	node := t.Root
	for node != nil || !s.IsEmpty() {
		if node != nil {
			s.Push(node)
			node = node.LChild
		} else {
			top, _ := s.Pop()
			fmt.Printf("%v ", top.Visit())
			//top.Visit()
			node = top.RChild
		}
	}
}

func (t *BiTree[T]) PostorderRecursion() {
	t.Root.postorder()
}

func (t *BiTree[T]) Postorder() {
	s := stack.NewLinkedStack[*Node[T]]()
	node := t.Root
	var r *Node[T]
	for node != nil || !s.IsEmpty() {
		if node != nil {
			s.Push(node)
			node = node.LChild
		} else {
			node = s.GetTop()
			if node.RChild != nil && node.RChild != r {
				node = node.RChild
				s.Push(node)
				node = node.LChild
			} else {
				node, _ = s.Pop()
				fmt.Printf("%v ", node.Visit())
				r = node
				node = nil
			}
		}
	}
}

func (t *BiTree[T]) Depth() int {
	return t.Root.depth()
}
