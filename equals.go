package generic

func EqualsSlice[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i, elem := range a {
		if elem != b[i] {
			return false
		}
	}

	return true
}

func EqualsSlice2[T comparable](a, b [][]T) bool {
	if len(a) != len(b) {
		return false
	}

	for i, elem := range a {
		if !EqualsSlice(elem, b[i]) {
			return false
		}
	}

	return true
}

func EqualsMap[K, V comparable](a, b map[K]V) bool {
	if len(a) != len(b) {
		return false
	}

	for k, elem := range a {
		if other, ok := b[k]; !ok || elem != other {
			return false
		}
	}

	return true
}
