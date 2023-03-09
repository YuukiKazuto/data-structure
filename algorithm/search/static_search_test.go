package search

import (
	"data-structure/algorithm/search/searchable"
	"fmt"
	"math/rand"
	"testing"
)

var ss StaticSearchTable

func TestSearch(t *testing.T) {
	i := ss.Search(50)
	fmt.Println(i)

}

func BenchmarkSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ss.Search(50000)
	}
}

func TestMain(m *testing.M) {
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := 10000
	s := make([]searchable.Int, n)
	//sp := Table[searchable.Int](s) //23362 ns/op
	sp := BinTable[searchable.Int]{Table: s} // 623.8 ns/op
	//sp := IndexTable[searchable.Int]{Table: s}
	ss = &sp
	e := make([]any, n)
	for i := 0; i < n; i++ {
		e[i] = searchable.Int(rand.Intn(100000))
	}
	ss.Create(e...)
	m.Run()
}
