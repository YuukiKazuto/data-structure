package main

import (
	"data-structure/queue"
	"fmt"
)

func main() {
	q := queue.NewSequenceQueue[int]()
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	for i := 0; i < 3; i++ {
		data, err := q.DeQueue()
		if err != nil {
			panic(err)
		}
		fmt.Println(data)
	}
}

func linked() {
	q := queue.NewLinkedQueue[int]()
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	q.EnQueue(6)
	q.EnQueue(7)
	q.EnQueue(8)
	q.EnQueue(9)
	fmt.Println(q.Length)
	length := q.Length
	for i := 0; i < length; i++ {
		data, err := q.DeQueue()
		if err != nil {
			panic(err)
		}
		fmt.Println(data)
	}
}
