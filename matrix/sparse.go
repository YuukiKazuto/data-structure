package matrix

type Node[T comparable] struct {
	row, col int
	val      T
}
type SeqSparse[T comparable] struct {
	sparse []Node[T]
}

func NewSeqSparse[T comparable](matrix [][]T) *SeqSparse[T] {
	row := len(matrix)
	col := len(matrix[0])
	s := []Node[T]{
		Node[T]{
			row: row,
			col: col,
		},
	}
	var zero T
	for i, rows := range matrix {
		for j, v := range rows {
			if v != zero {
				s = append(s, Node[T]{
					row: i,
					col: j,
					val: v,
				})
			}
		}
	}
	return &SeqSparse[T]{s}
}

func (s *SeqSparse[T]) Get(i, j int) (res T) {
	for k := 1; k < len(s.sparse); k++ {
		if i == s.sparse[k].row && j == s.sparse[k].col {
			res = s.sparse[k].val
			return
		}
	}
	return
}

func (s *SeqSparse[T]) Restore() [][]T {
	row := s.sparse[0].row
	col := s.sparse[0].col
	matrix := make([][]T, row)
	for i := 0; i < row; i++ {
		rows := make([]T, col)
		matrix[i] = rows
	}
	for i := 1; i < len(s.sparse); i++ {
		matrix[s.sparse[i].row][s.sparse[i].col] = s.sparse[i].val
	}
	return matrix
}

type ListNode[T comparable] struct {
	Node[T]
	Down, Right *ListNode[T]
}
type OrthogonalListSpase[T comparable] struct {
	down  []*ListNode[T]
	right []*ListNode[T]
}

func NewOrthogonalListSpase[T comparable](matrix [][]T) *OrthogonalListSpase[T] {
	row := len(matrix)
	col := len(matrix[0])
	down := make([]*ListNode[T], row)
	right := make([]*ListNode[T], col)
	var zero T
	for i, rows := range matrix {
		for j, v := range rows {
			if v != zero {
				l := &ListNode[T]{
					Node: Node[T]{
						row: i,
						col: j,
						val: v,
					},
				}
				if down[i] == nil {
					down[i] = l
				} else {
					d := down[i]
					for d.Down != nil {
						d = d.Down
					}
					d.Down = l
				}
				if right[j] == nil {
					right[j] = l
				} else {
					r := right[j]
					for r.Right != nil {
						r = r.Right
					}
					r.Right = l
				}
			}
		}
	}
	return &OrthogonalListSpase[T]{
		down:  down,
		right: right,
	}
}

func (s *OrthogonalListSpase[T]) Get(i, j int) (res T) {
	down := s.down[i-1]
	for down != nil {
		if down.col == j-1 {
			res = down.val
			return
		}
		down = down.Down
	}
	return
}

func (s *OrthogonalListSpase[T]) Restore() [][]T {
	row := len(s.down)
	col := len(s.right)
	matrix := make([][]T, row)
	for i := 0; i < row; i++ {
		rows := make([]T, col)
		matrix[i] = rows
	}
	for _, v := range s.down {
		for v != nil {
			matrix[v.row][v.col] = v.val
			v = v.Down
		}
	}
	return matrix
}
