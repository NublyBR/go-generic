package generic

import "cmp"

type orderable cmp.Ordered

type numeric interface {
	signed | unsigned
}

type signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64
}

type unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}
