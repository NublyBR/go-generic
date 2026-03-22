package generic

func CopySlice[T any](slice []T) []T {
	var ret = make([]T, len(slice))
	copy(ret, slice)

	return ret
}
