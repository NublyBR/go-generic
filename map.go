package generic

func Map[A any, B any](fn func(A) B, args ...A) []B {
	var ret = make([]B, len(args))

	for i, a := range args {
		ret[i] = fn(a)
	}

	return ret
}
