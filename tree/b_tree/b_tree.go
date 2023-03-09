package b_tree

import (
	"data-structure/algorithm/search"
	"fmt"
	"math"
)

var m = 3

type Record search.Searchable

type BTNode struct {
	parent *BTNode
	key    []int
	ptr    []*BTNode
	recPtr []Record
}

func (n *BTNode) String() string {
	return fmt.Sprintf(
		`{"key": %d, "ptr": %++v}`, n.key, n.ptr,
	)
}

func (n *BTNode) search(k int) int {
	keyNum := len(n.key)
	for i := 1; i < keyNum-1; i++ {
		if k >= n.key[i] && k < n.key[i+1] {
			return i
		}
	}
	if k >= n.key[keyNum-1] {
		return keyNum - 1
	}
	return 0
}

func (n *BTNode) insertKey(r Record) {
	keyNum := len(n.key)
	i := 1
	for ; i <= keyNum; i++ {
		if i == keyNum {
			n.key = append(n.key[:i], r.GetKey())
			n.ptr = append(n.ptr[:i], nil)
			n.recPtr = append(n.recPtr[:i], r)
		} else if n.key[i] > r.GetKey() {
			n.key = append(n.key, 0)
			copy(n.key[i+1:], n.key[i:])
			n.key[i] = r.GetKey()

			n.ptr = append(n.ptr, nil)
			copy(n.ptr[i+1:], n.ptr[i:])
			n.ptr[i] = nil

			n.recPtr = append(n.recPtr, nil)
			copy(n.recPtr[i+1:], n.recPtr[i:])
			n.recPtr[i] = r
			break
		}
	}
}

func (n *BTNode) split(s int) {
	k := n.key[s]
	r := n.recPtr[s]
	right := &BTNode{
		key:    append([]int{0}, n.key[s+1:]...),
		ptr:    n.ptr[s:],
		recPtr: append([]Record{nil}, n.recPtr[s+1:]...),
	}
	for i := range right.ptr {
		if right.ptr[i] != nil {
			right.ptr[i].parent = right
		}
	}
	n.key = n.key[:s]
	n.ptr = n.ptr[:s]
	n.recPtr = n.recPtr[:s]
	if n.parent != nil {
		right.parent = n.parent
		n.parent.insertKey(r)
		i := n.parent.search(r.GetKey())
		n.parent.ptr[i] = right
	} else {
		n.parent = &BTNode{
			key:    []int{0, k},
			ptr:    []*BTNode{n, right},
			recPtr: []Record{nil, r},
		}
		right.parent = n.parent
	}
	return
}

func (n *BTNode) delNode(i, pi int, p *BTNode) (parent *BTNode, pki int) {
	//k := n.key[i]
	keyNum := len(n.key) - 1
	if keyNum-1 >= int(math.Ceil(float64(m)/2))-1 || p == nil {
		n.key = append(n.key[:i], n.key[i+1:]...)
		n.ptr = append(n.ptr[:i], n.ptr[i+1:]...)
		n.recPtr = append(n.recPtr[:i], n.recPtr[i+1:]...)
		return
	}
	//pi := p.search(k)
	pKNum := len(p.key) - 1
	switch pi {
	case 0:
		nRB := p.ptr[pi+1]
		rbKNum := len(nRB.key) - 1
		if rbKNum-1 >= int(math.Ceil(float64(m)/2))-1 {
			n.borrowRight(i, pi, nRB)
		} else {
			n.mergeRight(i, pi, nRB)
			parent = p
			pki = pi + 1
		}
	case pKNum:
		nLB := p.ptr[pi-1]
		lbKNum := len(nLB.key) - 1
		if lbKNum-1 >= int(math.Ceil(float64(m)/2))-1 {
			n.borrowLeft(i, pi, nLB)
		} else {
			n.mergeLeft(i, pi, nLB)
			parent = p
			pki = pi
		}
	default:
		nLB := p.ptr[pi-1]
		lbKNum := len(nLB.key) - 1
		nRB := p.ptr[pi+1]
		rbKNum := len(nRB.key) - 1
		if lbKNum-1 >= int(math.Ceil(float64(m)/2))-1 {
			n.borrowLeft(i, pi, nLB)
		} else if rbKNum-1 >= int(math.Ceil(float64(m)/2))-1 {
			n.borrowRight(i, pi, nRB)
		} else {
			if lbKNum < rbKNum {
				n.mergeLeft(i, pi, nLB)
				pki = pi
			} else {
				n.mergeRight(i, pi, nRB)
				pki = pi + 1
			}
			parent = p
		}
	}
	return
}

//func (n *BTNode) delNodeMiddleOpt(i int, p *BTNode, k int) (parent *BTNode, pki int) {
//	pi := p.search(k)
//	pKNum := len(p.key) - 1
//	switch pi {
//	case 0:
//		nRB := p.ptr[pi+1]
//		rbKNum := len(nRB.key) - 1
//		if rbKNum >= int(math.Ceil(float64(m)/2))-1 {
//			n.borrowRight(i, pi, nRB)
//		} else {
//			n.mergeRight(i, pi, nRB)
//			parent = p
//			pki = pi + 1
//		}
//	case pKNum:
//		nLB := p.ptr[pi+1]
//		lbKNum := len(nLB.key) - 1
//		if lbKNum >= int(math.Ceil(float64(m)/2))-1 {
//			n.borrowLeft(i, pi, nLB)
//		} else {
//			n.mergeLeft(i, pi, nLB)
//			parent = p
//			pki = pi
//		}
//	default:
//		nLB := p.ptr[pi+1]
//		lbKNum := len(nLB.key) - 1
//		nRB := p.ptr[pi+1]
//		rbKNum := len(nRB.key) - 1
//		if lbKNum >= int(math.Ceil(float64(m)/2))-1 {
//			n.borrowLeft(i, pi, nLB)
//		} else if rbKNum >= int(math.Ceil(float64(m)/2))-1 {
//			n.borrowRight(i, pi, nRB)
//		} else {
//			if lbKNum < rbKNum {
//				n.mergeLeft(i, pi, nLB)
//				pki = pi
//			} else {
//				n.mergeRight(i, pi, nRB)
//				pki = pi + 1
//			}
//			parent = p
//		}
//	}
//	return
//}

func (n *BTNode) borrowRight(i, pi int, sib *BTNode) {
	keyNum := len(n.key) - 1
	k1 := sib.key[1]
	record1 := sib.recPtr[1]
	for j := i; j < keyNum; j++ {
		n.key[j] = n.key[j+1]
		n.recPtr[j] = n.recPtr[j+1]
	}
	n.key[keyNum] = n.parent.key[pi+1]
	n.recPtr[keyNum] = n.parent.recPtr[pi+1]
	n.parent.key[pi+1] = k1
	n.parent.recPtr[pi+1] = record1

	sib.key = append(sib.key[:1], sib.key[2:]...)
	sib.ptr = append(sib.ptr[:1], sib.ptr[2:]...)
	sib.recPtr = append(sib.recPtr[:1], sib.recPtr[2:]...)
}

func (n *BTNode) mergeRight(i, pi int, sib *BTNode) {
	keyNum := len(n.key) - 1
	for j := i; j < keyNum; j++ {
		n.key[j] = n.key[j+1]
		n.recPtr[j] = n.recPtr[j+1]
	}
	n.key[keyNum] = n.parent.key[pi+1]
	n.recPtr[keyNum] = n.parent.recPtr[pi+1]
	n.key = append(n.key, sib.key[1:]...)
	n.ptr = append(n.ptr[:keyNum], sib.ptr...)
	n.recPtr = append(n.recPtr, sib.recPtr[1:]...)
	for j := range n.ptr {
		if n.ptr[j] != nil {
			n.ptr[j].parent = n
		}
	}
}

func (n *BTNode) borrowLeft(i, pi int, sib *BTNode) {
	SKNum := len(sib.key) - 1
	k := sib.key[SKNum]
	record := sib.recPtr[SKNum]
	for j := i; j > 1; j-- {
		n.key[j] = n.key[j-1]
		n.recPtr[j] = n.recPtr[j-1]
	}
	n.key[1] = n.parent.key[pi]
	n.recPtr[1] = n.parent.recPtr[pi]
	n.parent.key[pi] = k
	n.parent.recPtr[pi] = record

	sib.key = sib.key[:SKNum]
	sib.ptr = sib.ptr[:SKNum]
	sib.recPtr = sib.recPtr[:SKNum]
}

func (n *BTNode) mergeLeft(i, pi int, sib *BTNode) {
	for j := i; j > 1; j-- {
		n.key[j] = n.key[j-1]
		n.recPtr[j] = n.recPtr[j-1]
	}
	n.key[1] = n.parent.key[pi]
	n.recPtr[1] = n.parent.recPtr[pi]
	n.key = append(sib.key, n.key[1:]...)
	n.ptr = append(sib.ptr, n.ptr[1:]...)
	n.recPtr = append(sib.recPtr, n.recPtr[1:]...)
	for j := range n.ptr {
		if n.ptr[j] != nil {
			n.ptr[j].parent = n
		}
	}
}

func (n *BTNode) searchPre(i int) (res *BTNode) {
	node := n.ptr[i-1]
	for node != nil {
		res = node
		kn := len(node.key) - 1
		node = node.ptr[kn]
	}
	return
}

func (n *BTNode) searchNext(i int) (res *BTNode) {
	node := n.ptr[i]
	for node != nil {
		res = node
		node = node.ptr[1]
	}
	return
}

type Result struct {
	pt  *BTNode
	i   int
	tag bool
}

type BTree struct {
	root *BTNode
}

func (bt *BTree) String() string {
	return fmt.Sprintf(`{"root": %v}`, bt.root)
}

func (bt *BTree) Search(k int) Result {
	node, found := bt.root, false
	i := 0
	var q *BTNode
	for node != nil && !found {
		i = node.search(k)
		if i > 0 && k == node.key[i] {
			found = true
			return Result{
				pt:  node,
				i:   i,
				tag: true,
			}
		} else {
			q = node
			node = node.ptr[i]
		}
	}
	return Result{
		pt: q,
		i:  i,
	}
}

func (bt *BTree) Insert(r Record) {
	if bt.root == nil {
		bt.root = &BTNode{
			key:    []int{0, r.GetKey()},
			ptr:    []*BTNode{nil, nil},
			recPtr: []Record{nil, r},
		}
		return
	}
	res := bt.Search(r.GetKey())
	if res.tag {
		return
	}
	p := res.pt
	parentInserted := false
	for p != nil {
		if !parentInserted {
			p.insertKey(r)
			parentInserted = true
		}
		if len(p.key)-1 < m {
			for p.parent != nil {
				p = p.parent
			}
			bt.root = p
			return
		} else {
			s := math.Ceil(float64(m) / 2)
			p.split(int(s))
			p = p.parent
		}
	}
}

func (bt *BTree) Delete(k int) Record {
	res := bt.Search(k)
	if !res.tag {
		return nil
	}
	node := res.pt
	i := res.i
	record := res.pt.recPtr[i]
	pi := 0
	p := node.parent
	if p != nil {
		pi = p.search(k)
	}
	if node.ptr[i-1] != nil {
		if pre, next := node.searchPre(i), node.searchNext(i); len(pre.key) > len(next.key) {
			i2 := len(pre.key) - 1
			p := pre.parent
			if p != nil {
				pi = p.search(pre.key[i2])
			}
			node.key[i], pre.key[i2] = pre.key[i2], node.key[i]
			node.recPtr[i], pre.recPtr[i2] = pre.recPtr[i2], node.recPtr[i]
			node = pre
			i = i2
		} else {
			i2 := 1
			p := next.parent
			if p != nil {
				pi = p.search(next.key[i2])
			}
			node.key[i], next.key[i2] = next.key[i2], node.key[i]
			node.recPtr[i], next.recPtr[i2] = next.recPtr[i2], node.recPtr[i]
			node = next
			i = i2
		}
	}
	parent, pki := node.delNode(i, pi, p)
	for parent != nil {
		p = parent.parent
		if p != nil {
			pi = p.search(k)
		}
		parent, pki = parent.delNode(pki, pi, p)
	}
	root := bt.root
	if len(root.key) == 1 {
		root.ptr[0].parent = nil
		bt.root = root.ptr[0]
		root = nil
	}
	return record
}
