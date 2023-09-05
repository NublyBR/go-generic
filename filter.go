package generic

func Filter[T any](fn func(T) bool, args ...T) []T {
	var ret = make([]T, 0, len(args))

	for _, a := range args {
		if fn(a) {
			ret = append(ret, a)
		}
	}

	return ret
}

func FilterMap[K comparable, V any](fn func(K, V) bool, mp map[K]V) map[K]V {
	var ret = make(map[K]V, len(mp))

	for k, v := range mp {
		if fn(k, v) {
			ret[k] = v
		}
	}

	return ret
}
