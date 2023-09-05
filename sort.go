package generic

import "sort"

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
