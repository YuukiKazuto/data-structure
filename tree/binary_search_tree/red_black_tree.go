package binary_search_tree

import "golang.org/x/exp/constraints"

const (
	RED   = true
	BLACK = false
)

type RBTNode[K constraints.Ordered, V comparable] struct {
	parent         *RBTNode[K, V]
	key            K
	value          V
	lChild, rChild *RBTNode[K, V]
	color          bool
}

func (n *RBTNode[K, V]) solveDoubleRed() {
	p, g := n.parent, n.parent.parent
	if g == nil || !p.color {
		return
	}
	if g.parent == nil && p.color {
		n.color = BLACK
		return
	}
	gpl := g == g.parent.lChild
	npl := n == p.lChild
	if p == g.lChild {
		u := g.rChild
		g.color = RED
		if u == nil || !u.color {
			if npl {
				p.color = BLACK
				if gpl {
					g.parent.lChild = p
				} else {
					g.parent.rChild = p
				}
				p.parent = g.parent
				g.lChild = nil
				p.rChild = g
				g.parent = p
			} else {
				n.color = BLACK
				if gpl {
					g.parent.lChild = n
				} else {
					g.parent.rChild = n
				}
				n.parent = g.parent
				n.lChild = p
				n.rChild = g
				g.lChild = nil
				p.rChild = nil
				p.parent = n
				g.parent = n
			}
		} else {
			p.color, u.color = BLACK, BLACK
		}
	} else {
		u := g.lChild
		g.color = RED
		if u == nil || !u.color {
			if npl {
				n.color = BLACK
				if gpl {
					g.parent.lChild = n
				} else {
					g.parent.rChild = n
				}
				n.parent = g.parent
				n.lChild = g
				n.rChild = p
				p.lChild = nil
				g.rChild = nil
				p.parent = n
				g.parent = n
			} else {
				p.color = BLACK
				if gpl {
					g.parent.lChild = p
				} else {
					g.parent.rChild = p
				}
				p.parent = g.parent
				g.rChild = nil
				p.lChild = g
				g.parent = p
			}
		} else {
			p.color, u.color = BLACK, BLACK
		}
	}
}

type RBTree[K constraints.Ordered, V comparable] struct {
	root *RBTNode[K, V]
}

func (t *RBTree[K, V]) Search(key K) *RBTNode[K, V] {
	node := t.root
	for node != nil && key != node.key {
		if key < node.key {
			node = node.lChild
		} else {
			node = node.rChild
		}
	}
	return node
}

// Insert TODO: 重新完善该算法
func (t *RBTree[K, V]) Insert(k K, v V) {
	if x := t.Search(k); x != nil {
		return
	}
	newNode := &RBTNode[K, V]{
		key:   k,
		value: v,
		color: RED,
	}
	if t.root == nil {
		newNode.color = BLACK
		t.root = newNode
	} else {
		var p *RBTNode[K, V]
		for node := t.root; node != nil; {
			p = node
			if k < node.key {
				node = node.lChild
			} else {
				node = node.rChild
			}
		}
		newNode.parent = p
		if k < p.key {
			p.lChild = newNode
		} else {
			p.rChild = newNode
		}
		newNode.solveDoubleRed()
		node := newNode
		for node.parent != nil {
			node = node.parent
		}
		t.root = node
	}
}

// Delete TODO: 实现该算法
func (t *RBTree[K, V]) Delete(key K) {

}