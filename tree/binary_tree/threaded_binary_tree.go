package binary_tree

// TNode 线索二叉树节点
type TNode[T comparable] struct {
	Data           T
	LChild, RChild *TNode[T]
	LTag, RTag     int
}

func (n *TNode[T]) visit(p **TNode[T]) {
	if n.LChild == nil {
		n.LChild = *p
		n.LTag = 1
	}
	if *p != nil && (*p).RChild == nil {
		(*p).RChild = n
		(*p).RTag = 1
	}
	*p = n
}

func (n *TNode[T]) InThread(p **TNode[T]) {
	if n != nil {
		n.LChild.InThread(p)
		n.visit(p)
		n.RChild.InThread(p)
	}
}

func (n *TNode[T]) PreThread(p **TNode[T]) {
	if n != nil {
		n.visit(p)
		if n.LTag == 0 {
			n.LChild.PreThread(p)
		}
		if n.RTag == 0 {
			n.RChild.PreThread(p)
		}
	}
}

func (n *TNode[T]) PostThread(p **TNode[T]) {
	if n != nil {
		n.LChild.PostThread(p)
		n.RChild.PostThread(p)
		n.visit(p)
	}
}

type ThreadTree[T comparable] struct {
	Root *TNode[T]
}

func NewThreadTree[T comparable](bt *BiTree[T]) *ThreadTree[T] {
	if bt.Root != nil {
		t := &ThreadTree[T]{}
		t.Root = &TNode[T]{Data: bt.Root.Data}
		queue := []*Node[T]{bt.Root}
		queue1 := []*TNode[T]{t.Root}
		for len(queue) != 0 {
			cur := queue[0]
			tn := queue1[0]
			queue = queue[1:]
			queue1 = queue1[1:]
			if cur.LChild != nil {
				tn.LChild = &TNode[T]{Data: cur.LChild.Data}
				queue = append(queue, cur.LChild)
				queue1 = append(queue1, tn.LChild)
			}
			if cur.RChild != nil {
				tn.RChild = &TNode[T]{Data: cur.RChild.Data}
				queue = append(queue, cur.RChild)
				queue1 = append(queue1, tn.RChild)
			}
		}
		return t
	}
	return nil
}
