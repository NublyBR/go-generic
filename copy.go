package generic

func CopySlice[T any](slice []T) []T {
	var ret = make([]T, len(slice))
	copy(ret, slice)

	return ret
}

func CopyMap[K comparable, V any](mp map[K]V) map[K]V {
	var ret = make(map[K]V, len(mp))
	for k := range mp {
		ret[k] = mp[k]
	}

	return ret
}
