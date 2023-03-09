package binary_search_tree

import "fmt"

const (
	RH = iota - 1 // 右高
	EH            // 等高
	LH            // 左高
)

type AVLNode struct {
	Key            int
	Bf             int // 平衡因子
	LChild, RChild *AVLNode
}

func (n *AVLNode) LRotate() *AVLNode {
	rc := n.RChild
	n.RChild = rc.LChild
	rc.LChild = n
	return rc
}

func (n *AVLNode) RRotate() *AVLNode {
	lc := n.LChild
	n.LChild = lc.RChild
	lc.RChild = n
	return lc
}

func (n *AVLNode) inorder() {
	if n != nil {
		n.LChild.inorder()
		fmt.Printf("%d ", n.Key)
		n.RChild.inorder()
	}
}
func (n *AVLNode) LeftBalance() *AVLNode {
	res := n
	lc := res.LChild
	switch lc.Bf {
	case LH:
		res.Bf = EH
		lc.Bf = EH
		res = res.RRotate()
	case RH:
		rd := lc.RChild
		switch rd.Bf {
		case LH:
			res.Bf = RH
			lc.Bf = EH
		case EH:
			res.Bf = EH
			lc.Bf = EH
		case RH:
			res.Bf = EH
			lc.Bf = LH
		}
		rd.Bf = EH
		res.LChild = res.LChild.LRotate()
		res = res.RRotate()
	}
	return res
}

func (n *AVLNode) RightBalance() *AVLNode {
	res := n
	rc := res.RChild
	switch rc.Bf {
	case LH:
		ld := rc.LChild
		switch ld.Bf {
		case LH:
			res.Bf = EH
			rc.Bf = RH
		case EH:
			res.Bf = EH
			rc.Bf = EH
		case RH:
			res.Bf = LH
			rc.Bf = EH
		}
		ld.Bf = EH
		res.RChild = res.RChild.RRotate()
		res = res.LRotate()
	case RH:
		res.Bf = EH
		rc.Bf = EH
		res = res.LRotate()
	}
	return res
}

type AVLTree struct {
	Root *AVLNode
}

func (t *AVLTree) Insert(n **AVLNode, key int, taller *bool) bool {
	if *n == nil {
		*n = &AVLNode{Key: key}
		*taller = true
	} else {
		if key == (*n).Key {
			*taller = false
			return false
		}
		if key < (*n).Key {
			if !t.Insert(&((*n).LChild), key, taller) {
				return false
			}
			if *taller {
				switch (*n).Bf {
				case LH:
					*n = (*n).LeftBalance()
					*taller = false
				case EH:
					(*n).Bf = LH
					*taller = true
				case RH:
					(*n).Bf = EH
					*taller = false
				}
			}
		} else {
			if !t.Insert(&((*n).RChild), key, taller) {
				return false
			}
			if *taller {
				switch (*n).Bf {
				case LH:
					(*n).Bf = EH
					*taller = false
				case EH:
					(*n).Bf = RH
					*taller = true
				case RH:
					*n = (*n).RightBalance()
					*taller = false
				}
			}
		}
	}
	return true
}

func (t *AVLTree) Create(arr []int) {
	for _, v := range arr {
		var taller bool
		t.Insert(&(t.Root), v, &taller)
	}
}

func (t *AVLTree) Inorder() {
	t.Root.inorder()
}

func (t *AVLTree) Level() {
	if t.Root != nil {
		queue := []*AVLNode{t.Root}
		for len(queue) != 0 {
			cur := queue[0]
			queue = queue[1:]
			fmt.Printf("%v ", cur.Key)
			if cur.LChild != nil {
				queue = append(queue, cur.LChild)
			}
			if cur.RChild != nil {
				queue = append(queue, cur.RChild)
			}
		}
	}
}
