package generic

import "cmp"

type Orderable cmp.Ordered

type Numeric interface {
	Signed | Unsigned
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Item[K comparable, V any] struct {
	Key   K
	Value V
}

type Items[K comparable, V any] []Item[K, V]

type Zipped[A, B any] struct {
	A A
	B B
}
