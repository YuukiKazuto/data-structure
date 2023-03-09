package matrix

import "math"

type LowerTriangular[T comparable] []T

func NewLowerTriangular[T comparable](matrix [][]T) LowerTriangular[T] {
	n := len(matrix)
	if n != len(matrix[0]) {
		return LowerTriangular[T]{}
	}
	size := (n+1)*n/2 + 1
	compression := make(LowerTriangular[T], size, size)
	index := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j <= i; j++ {
			compression[index] = matrix[i][j]
			index++
		}
	}
	compression[index] = matrix[0][1]
	return compression
}

func (lt LowerTriangular[T]) Get(i, j int) T {
	if j > i {
		return lt[len(lt)-1]
	}
	return lt[i*(i-1)/2+j-1]
}

type UpperTriangular[T comparable] []T

func NewUpperTriangular[T comparable](matrix [][]T) UpperTriangular[T] {
	n := len(matrix)
	if n != len(matrix[0]) {
		return UpperTriangular[T]{}
	}
	size := (n+1)*n/2 + 1
	compression := make(UpperTriangular[T], size, size)
	index := 0
	for i := 0; i < len(matrix); i++ {
		for j := i; j < n; j++ {
			compression[index] = matrix[i][j]
			index++
		}
	}
	compression[index] = matrix[1][0]
	return compression
}

func (ut UpperTriangular[T]) Get(i, j int) T {
	if i > j {
		return ut[len(ut)-1]
	}
	n := int((-1 + math.Sqrt(float64(1+8*(len(ut)-1)))) / 2)
	return ut[((i-1)*(2*n-i+2))/2+j-i]
}
