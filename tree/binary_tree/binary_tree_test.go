package binary_tree

import (
	"fmt"
	"testing"
)

func BenchmarkInorderThreadTree_Inorder(b *testing.B) {
	bt := NewBiTree[int]()
	for i := 0; i < 10; i++ {
		bt.AddNode(i)
	}
	t := NewThreadTree(bt)
	in := t.CreateInThread()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		in.Inorder()
		fmt.Println()
	}

}

func BenchmarkPreorderThreadTree_Preorder(b *testing.B) {
	bt := NewBiTree[int]()
	for i := 0; i < 10; i++ {
		bt.AddNode(i)
	}
	t := NewThreadTree(bt)
	pre := t.CreatePreThread()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pre.Preorder()
		fmt.Println()
	}

}

//func TestMain(m *testing.M) {
//	m.Run()
//}

//pre := t.CreatePreThread()
//post := t.CreatePostThread()
