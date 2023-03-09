package main

import (
	"data-structure/matrix"
	"fmt"
)

func main() {
	//m := [][]int{
	//	{1, 2, 3, 7},
	//	{0, 4, 5, 8},
	//	{0, 0, 6, 9},
	//	{0, 0, 0, 10},
	//}
	//ut := matrix.NewUpperTriangular[int](m)
	//fmt.Println(ut.Get(1, 2))
	//fmt.Println(ut.Get(3, 4))
	//fmt.Println(ut.Get(2, 1))
	//fmt.Println(ut.Get(2, 2))
	//fmt.Println(ut.Get(3, 3))

	m := [][]int{
		{1, 0, 3, 7},
		{0, 4, 0, 8},
		{0, 0, 6, 0},
		{0, 10, 0, 0},
	}

	os := matrix.NewOrthogonalListSpase[int](m)
	fmt.Println(os.Get(1, 3))
	fmt.Println(os.Get(2, 2))
	fmt.Println(os.Get(2, 1))
	fmt.Println(os.Get(3, 3))
	fmt.Println(os.Get(4, 2))
	fmt.Println(os.Restore())
}
