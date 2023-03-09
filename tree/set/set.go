package set

const ElementNotFound = -1

type element[T comparable] struct {
	Data   T
	Parent int
}
type Set[T comparable] struct {
	E []element[T]
}

func NewSet[T comparable](els []T) *Set[T] {
	e := make([]element[T], 0)
	for _, el := range els {
		e = append(e, element[T]{
			Data:   el,
			Parent: -1,
		})
	}
	return &Set[T]{E: e}
}

func (s *Set[T]) Find(x T) int {
	for i, e := range s.E {
		if e.Data == x {
			for s.E[i].Parent >= 0 {
				i = s.E[i].Parent
			}
			return i
		}
	}
	return ElementNotFound
}

func (s *Set[T]) Union(s1e, s2e T) {
	s1 := s.Find(s1e)
	s2 := s.Find(s2e)
	if (s1 != ElementNotFound && s2 != ElementNotFound) && s1 != s2 {
		s.E[s2].Parent = s1
	}
}
