package searchable

import "strconv"

type Int int

func (i Int) GetKey() int {
	return int(i)
}

type Float32 float32

func (f Float32) GetKey() int {
	sf := strconv.FormatFloat(float64(f), 'f', 2, 32)
	k := 0
	for _, c := range sf {
		k += int(c)
	}
	return k
}

type Float64 float64

func (f Float64) GetKey() int {
	sf := strconv.FormatFloat(float64(f), 'f', 2, 32)
	k := 0
	for _, c := range sf {
		k += int(c)
	}
	return k
}
