package floyd

import (
	"data-structure/graph"
	. "data-structure/graph/application/shortest_path/global"
	"log"
)

// ShorterPath 目前只能使用邻接矩阵存储的图
func ShorterPath[T comparable](g graph.Graph[T]) {
	for i, _ := range PathM {
		for j, _ := range PathM[i] {
			PathM[i][j] = -1
		}
	}
	log.Println(PathM)
	log.Println(A)
	n := len(PathM)
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if A[i][k] == INFINITY || A[k][j] == INFINITY {
					continue
				}
				if A[i][j] > A[i][k]+A[k][j] {
					A[i][j] = A[i][k] + A[k][j]
					PathM[i][j] = k
				}
			}
		}
	}
}
