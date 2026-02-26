package generic

func EqualsSlice[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func EqualsSlice2[T comparable](a, b [][]T) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if !EqualsSlice(a[i], b[i]) {
			return false
		}
	}

	return true
}

func EqualsMap[K, V comparable](a, b map[K]V) bool {
	if len(a) != len(b) {
		return false
	}

	for k := range a {
		if a[k] != b[k] {
			return false
		}
	}

	return true
}

func EqualApprox[T numeric](lhs, rhs, delta T) bool {
	if lhs > rhs {
		return lhs-rhs < delta
	}

	return rhs-lhs < delta
}
