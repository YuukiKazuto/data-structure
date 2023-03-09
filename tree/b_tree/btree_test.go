package b_tree

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

type user struct {
	id   int
	name string
}

func (u *user) GetKey() int {
	return u.id
}

func TestBTree_Insert(t *testing.T) {
	m = 5
	b := &BTree{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rint := make([]int, 0)
	for i := 1; i <= 20; i++ {
		rint = append(rint, r.Intn(100))
	}
	fmt.Println(rint)
	for _, v := range rint {
		record := &user{
			id:   v,
			name: "s" + strconv.Itoa(v),
		}
		b.Insert(record)
	}
	fmt.Printf("%v\n", b)
}

func TestBTree_Delete(t *testing.T) {
	m = 5
	b := &BTree{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rint := make([]int, 0)
	for i := 1; i <= 20; i++ {
		rint = append(rint, r.Intn(100))
	}
	fmt.Println(rint)
	for _, v := range rint {
		record := &user{
			id:   v,
			name: "s" + strconv.Itoa(v),
		}
		b.Insert(record)
	}
	fmt.Printf("%v\n", b)
	fmt.Println("-----------------------------------------------------------")
	record := b.Delete(rint[r.Intn(20)])
	fmt.Printf("%v\n", record)
	fmt.Printf("%v\n", b)
	record = b.Delete(rint[r.Intn(20)])
	fmt.Printf("%v\n", record)
	fmt.Printf("%v\n", b)
	record = b.Delete(rint[r.Intn(20)])
	fmt.Printf("%v\n", record)
	fmt.Printf("%v\n", b)
	record = b.Delete(rint[r.Intn(20)])
	fmt.Printf("%v\n", record)
	fmt.Printf("%v\n", b)
}
