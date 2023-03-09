package search

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
	~float32 | ~float64
}

type StaticSearchTable interface {
	Create(e ...any)
	// Search 令 keyType 为 int。 若原始的 key 是浮点数，则将其转换成 int (1)。
	// 注：(1)这里的转换不是单纯的类型转换，是将其通过某种方式计算得到的 int。
	Search(key int) int
	Traverse(visit func(v any))
}

type DynamicSearchTable interface {
	Init()
	Search(key any) int
	Insert(v any)
	Delete(key any) any
	Traverse(visit func(v any))
}
