package generic

func Filter[T any](slice []T, fn func(T) bool) []T {
	var ret = make([]T, 0, len(slice))

	for _, a := range slice {
		if fn(a) {
			ret = append(ret, a)
		}
	}

	return ret
}

func FilterMap[K comparable, V any](mp map[K]V, fn func(K, V) bool) map[K]V {
	var ret = make(map[K]V, len(mp))

	for k, v := range mp {
		if fn(k, v) {
			ret[k] = v
		}
	}

	return ret
}

func Partition[T any](slice []T, fn func(T) bool) (match []T, rest []T) {
	match = make([]T, 0, len(slice))
	rest = make([]T, 0, len(slice))

	for _, v := range slice {
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

func UniqueBy[K comparable, T any](slice []T, fn func(T) K) []T {
	var (
		ret = make([]T, 0, len(slice))
		key = make([]K, 0, len(slice))
	)

outer:
	for _, v := range slice {
		k := fn(v)

		for _, cmp := range key {
			if k == cmp {
				continue outer
			}
		}

		key = append(key, k)
		ret = append(ret, v)
	}

	return ret
}

type Once[K comparable] struct {
	Keys []K
}

func (o *Once[K]) First(key K) bool {
	for _, cmp := range o.Keys {
		if key == cmp {
			return false
		}
	}

	o.Keys = append(o.Keys, key)

	return true
}
