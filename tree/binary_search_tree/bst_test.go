package binary_search_tree

import (
	"testing"
)

func TestBST(t *testing.T) {
	tree := new(RBTree[int, int])
	ints := []int{19, 13, 17, 15, 10, 50, 26, 21, 30, 66, 60, 70, 11, 44, 7, 4}
	for _, v := range ints {
		tree.Insert(v, v)
	}
	tree.Search(13)
	//tree.Inorder()
	//fmt.Println()
	//tree.Level()
	//fmt.Println()
}
