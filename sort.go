package generic

import "sort"

type sortAB[A orderable, B any] struct {
	a []A
	b []B
}

func (s sortAB[A, B]) Len() int           { return len(s.a) }
func (s sortAB[A, B]) Less(i, j int) bool { return s.a[i] < s.a[j] }
func (s sortAB[A, B]) Swap(i, j int) {
	s.a[i], s.a[j] = s.a[j], s.a[i]
	s.b[i], s.b[j] = s.b[j], s.b[i]
}

func Sort[T orderable](slice []T) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
}

func SortFunc[T any](slice []T, less func(a, b T) bool) {
	sort.Slice(slice, func(i, j int) bool {
		return less(slice[i], slice[j])
	})
}

func SortBy[T any, K orderable](slice []T, fn func(T) K) {
	var order = &sortAB[K, T]{
		a: Map(fn, slice...),
		b: slice,
	}

	sort.Sort(order)
}

func Sorted[T orderable](slice []T) []T {
	var ret = make([]T, len(slice))
	copy(ret, slice)

	Sort(ret)

	return ret
}

func SortedFunc[T any](slice []T, less func(a, b T) bool) []T {
	var ret = make([]T, len(slice))
	copy(ret, slice)

	SortFunc(ret, less)

	return ret
}

func SortedBy[T any, K orderable](slice []T, fn func(T) K) []T {
	var ret = make([]T, len(slice))
	copy(ret, slice)

	SortBy(ret, fn)

	return ret
}
