package search

import (
	"math"
	"sort"
)

const NotFound = -1

type Searchable interface {
	GetKey() int
}

type Table[T Searchable] []T

func (s *Table[T]) Less(i, j int) bool {
	return s.Get(i).(Searchable).GetKey() < s.Get(j).(Searchable).GetKey()
}

func (s *Table[T]) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func (s *Table[T]) Len() int {
	return len(*s)
}

func (s *Table[T]) Get(i int) any {
	return (*s)[i]
}

func (s *Table[T]) Sort() {
	sort.Sort(s)
}

func (s *Table[T]) Create(e ...any) {
	for i, a := range e {
		if i < s.Len() {
			(*s)[i] = a.(T)
		}
	}
}

func (s *Table[T]) Search(key int) int {
	for i, v := range *s {
		if key == v.GetKey() {
			return i
		}
	}
	return NotFound
}

func (s *Table[T]) Traverse(visit func(v any)) {
	for _, v := range *s {
		visit(v)
	}
}

type BinTable[T Searchable] struct {
	Table[T]
}

func (s *BinTable[T]) Create(e ...any) {
	s.Table.Create(e...)
	s.Sort()
}
func (s *BinTable[T]) Search(key int) int {
	low, high := 0, (*s).Len()-1
	for low <= high {
		mid := (low + high) / 2
		if key == s.Get(mid).(Searchable).GetKey() {
			return mid
		} else if key < s.Get(mid).(Searchable).GetKey() {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return NotFound
}

type Index struct {
	maxKey int
	start  int
}

type IndexTable[T Searchable] struct {
	index []Index
	Table[T]
}

func (s *IndexTable[T]) Create(e ...any) {
	s.Table.Create(e...)
	s.Sort()
	block := int(math.Ceil(math.Sqrt(float64(s.Len()))))
	s.index = make([]Index, block)
	lenBlock := int(math.Ceil(float64(s.Len()) / float64(block)))
	for i := 0; i < block-1; i++ {
		s.index[i].start = i * lenBlock
		s.index[i].maxKey = s.Get((i+1)*lenBlock - 1).(Searchable).GetKey()
	}
	s.index[block-1].start = (block - 1) * lenBlock
	s.index[block-1].maxKey = s.Get(s.Len() - 1).(Searchable).GetKey()
}
func (s *IndexTable[T]) Search(key int) int {
	start, end := s.searchIndexBySeq(key)
	if end < 0 {
		return NotFound
	}
	for i := start; i < end; i++ {
		if key == s.Get(i).(Searchable).GetKey() {
			return i
		}
	}
	return NotFound
}

func (s *IndexTable[T]) searchIndexBySeq(key int) (int, int) {
	start, end := 0, -1
	n := len(s.index)
	for i, index := range s.index {
		if key <= index.maxKey {
			start = index.start
			if i < n-1 {
				end = s.index[i+1].start
			} else {
				end = n
			}
			break
		}
	}
	return start, end
}

func (s *IndexTable[T]) searchIndexByBin(key int) (int, int) { //78.28
	start, end := 0, -1
	n := len(s.index)
	low, high := 0, n-1
	for low <= high {
		mid := (low + high) / 2
		if key <= s.index[mid].maxKey {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	if low < n {
		start = s.index[low].start
		if low < n-1 {
			end = s.index[low+1].start
		} else {
			end = n
		}
	}
	return start, end
}
