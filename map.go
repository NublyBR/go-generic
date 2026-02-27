package generic

func Map[A any, B any](slice []A, fn func(A) B) []B {
	var ret = make([]B, len(slice))

	for i := range slice {
		ret[i] = fn(slice[i])
	}

	return ret
}
