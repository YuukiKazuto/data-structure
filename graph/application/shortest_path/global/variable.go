package global

import "math"

var (
	Visited []bool
	Q       any
	Dist    []int
	Path    []int
	PathM   [][]int
	A       [][]int
	Final   []bool
)

const INFINITY = math.MaxInt
