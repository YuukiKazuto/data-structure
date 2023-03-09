package search

func Hash(key string, tableSize int) (h uint) {
	for _, k := range key {
		h = (h << 5) + uint(k)
	}
	h %= uint(tableSize)
	return
}
