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

func Partition[T any](fn func(T) bool, args ...T) (match []T, rest []T) {
	match = make([]T, 0, len(args))
	rest = make([]T, 0, len(args))

	for _, v := range args {
		if fn(v) {
			match = append(match, v)
		} else {
			rest = append(rest, v)
		}
	}

	return
}

func Unique[T comparable](slice []T) []T {
	var ret = make([]T, 0, len(slice))

outer:
	for i, v := range slice {
		for _, cmp := range slice[:i] {
			if v == cmp {
				continue outer
			}
		}

		ret = append(ret, v)
	}

	return ret
}
