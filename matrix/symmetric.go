package matrix

type Symmetric[T comparable] []T

func newSymmetric[T comparable](matrix [][]T) Symmetric[T] {
	n := len(matrix)
	if n != len(matrix[0]) {
		return Symmetric[T]{}
	}
	size := (n + 1) * n / 2
	compression := make(Symmetric[T], size, size)
	index := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j <= i; j++ {
			compression[index] = matrix[i][j]
			index++
		}
	}
	return compression
}

func (s Symmetric[T]) Get(i, j int) T {
	if j > i {
		i, j = j, i
	}
	return s[i*(i-1)/2+j-1]
}
