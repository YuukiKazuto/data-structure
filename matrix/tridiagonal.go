package matrix

import (
	"math"
)

type Tridiagonal[T comparable] []T

func NewTridiagonal[T comparable](matrix [][]T) Tridiagonal[T] {
	n := len(matrix)
	if n != len(matrix[0]) {
		return Tridiagonal[T]{}
	}
	size := 3*n - 2
	compression := make(Tridiagonal[T], size, size)
	index := 0
	for j := 0; j < 2; j++ {
		compression[index] = matrix[0][j]
		index++
	}
	for i := 1; i < n-1; i++ {
		for j := i - 1; j <= i+1; j++ {
			compression[index] = matrix[i][j]
			index++
		}
	}
	{
		compression[index] = matrix[n-1][n-2]
		index++
		compression[index] = matrix[n-1][n-1]
	}
	return compression
}

func (t Tridiagonal[T]) Get(i, j int) (res T) {
	if math.Abs(float64(i-j)) <= 1 {
		res = t[2*i+j-3]
		return
	}
	return
}
