package main

import (
	"data-structure/tree/binary_tree"
	"fmt"
)

func main() {
	//bt := binary_tree.NewBiTree[int]()
	//for i := 0; i < 10; i++ {
	//	bt.AddNode(i)
	//}
	//tt := binary_tree.NewThreadTree(bt)
	//in := tt.CreateInThread()
	//in.Inorder()
	//fmt.Println()
	//in.RevInorder()
	//fmt.Println()
	//tt = binary_tree.NewThreadTree(bt)
	//tt.CreatePreThread().Preorder()
	//fmt.Println()
	//tt = binary_tree.NewThreadTree(bt)
	//tt.CreatePostThread().RevPostorder()
	//fmt.Println()
}

func intTree() {
	t := binary_tree.NewBiTree[int]()
	t.AddNode(1)
	t.AddNode(2)
	t.AddNode(3)
	t.AddNode(4)
	t.AddNode(5)
	t.AddNode(6)
	t.AddNode(7)

	fmt.Print("层序遍历：")
	t.Level()
	fmt.Println()
	fmt.Print("先序遍历：")
	t.Preorder()
	fmt.Println()
	fmt.Print("中序遍历：")
	t.Inorder()
	fmt.Println()
	fmt.Print("后序遍历：")
	t.Postorder()
	fmt.Println()
	fmt.Print("t的深度：")
	fmt.Println(t.Depth())
}
func stringTree() {
	t := binary_tree.NewBiTree[string]()
	t.AddNode("A")
	t.AddNode("B")
	t.AddNode("C")
	t.AddNode("D")
	t.AddNode("E")
	t.AddNode("F")
	t.AddNode("G")

	fmt.Print("层序遍历：")
	t.Level()
	fmt.Println()
	fmt.Print("先序遍历：")
	t.Preorder()
	fmt.Println()
	fmt.Print("中序遍历：")
	t.Inorder()
	fmt.Println()
	fmt.Print("后序遍历：")
	t.Postorder()
	fmt.Println()
	fmt.Print("t的深度：")
	fmt.Println(t.Depth())
}
