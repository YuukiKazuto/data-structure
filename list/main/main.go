package main

import (
	"data-structure/list"
	"fmt"
	"strconv"
)

func main() {
	l := list.NewArrayList[string]()
	for i := 0; i < 10; i++ {
		l.Add(strconv.Itoa(i) + "i")
	}
	fmt.Println(l.Slice)
	l.Insert(2, "aaa")
	fmt.Println(l.Slice)
	s, _ := l.Delete(2)
	fmt.Println(s)
	fmt.Println(l.Slice)
	fmt.Println(l.Index("6i"))
}
