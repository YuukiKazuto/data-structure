package sort

const (
	MaxDigit = 4
	Radix    = 10
)

type node struct {
	key  int
	next *node
}

type bucketNode struct {
	head, tail *node
}

type bucket []bucketNode

func getDigit(x, d int) (res int) {
	for i := 1; i <= d; i++ {
		res = x % Radix
		x /= Radix
	}
	return
}

func LSDRadixSort[T Number](arr []T, n int) {
	b := make(bucket, Radix)
	var l *node
	for _, v := range arr {
		l = &node{
			key:  int(v),
			next: l,
		}
	}
	for d := 1; d <= MaxDigit; d++ {
		for l != nil {
			di := getDigit(l.key, d)
			tmp := l
			l = l.next
			tmp.next = nil
			if b[di].head == nil {
				b[di].head = tmp
				b[di].tail = b[di].head
			} else {
				b[di].tail.next = tmp
				b[di].tail = tmp
			}
		}
		l = nil
		for di := Radix - 1; di >= 0; di-- {
			if b[di].head != nil {
				b[di].tail.next = l
				l = b[di].head
				b[di].head = nil
				b[di].tail = nil
			}
		}
	}
	for i := 0; i < n; i++ {
		tmp := l
		l = l.next
		arr[i] = T(tmp.key)
		tmp = nil
	}
}

func msd[T Number](arr []T, l, r, d int) {
	b := make(bucket, Radix)
	var li *node
	if d == 0 {
		return
	}
	for i := range b {
		b[i].head = b[i].tail
	}
	for i := l; i <= r; i++ {
		li = &node{
			key:  int(arr[i]),
			next: li,
		}

	}
	for li != nil {
		di := getDigit(li.key, d)
		tmp := li
		li = li.next
		if b[di].head == nil {
			b[di].tail = tmp
		}
		tmp.next = b[di].head
		b[di].head = tmp
	}
	i, j := l, l
	for di := 0; di < Radix; di++ {
		if b[di].head != nil {
			p := b[di].head
			for p != nil {
				arr[j], p = T(p.key), p.next
				j++
			}
			msd(arr, i, j-1, d-1)
			i = j
		}
	}
}

func MSDRadixSort[T Number](arr []T, n int) {
	msd(arr, 0, n-1, MaxDigit)
}
