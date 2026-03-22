package generic

func Map[A any, B any](slice []A, fn func(A) B) []B {
	var ret = make([]B, len(slice))

	for i := range slice {
		ret[i] = fn(slice[i])
	}

	return ret
}

func Filter[T any](slice []T, fn func(T) bool) []T {
	var ret = make([]T, 0, len(slice))

	for _, a := range slice {
		if fn(a) {
			ret = append(ret, a)
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

func Reverse[T any](slice []T) []T {
	var (
		ret = make([]T, len(slice))
		ln1 = len(slice) - 1
	)

	for i, elem := range slice {
		ret[ln1-i] = elem
	}

	return ret
}

func ReverseInPlace[T any](slice []T) {
	var (
		ln  = len(slice)
		ln1 = ln - 1
	)

	for i := 0; i < ln/2; i++ {
		slice[i], slice[ln1-i] = slice[ln1-i], slice[i]
	}
}

func Zip[T any](args ...[]T) [][]T {
	if len(args) == 0 {
		return nil
	}

	var (
		smallest = len(args[0])
	)

	for _, elem := range args[1:] {
		if ln := len(elem); ln < smallest {
			smallest = ln
		}
	}

	var ret = make([][]T, smallest)

	for i := range ret {
		ret[i] = make([]T, len(args))

		for j, elem := range args {
			ret[i][j] = elem[i]
		}
	}

	return ret
}

func Zip2[A, B any](a []A, b []B) []Zipped[A, B] {
	var (
		smallest = min(len(a), len(b))
		ret      = make([]Zipped[A, B], smallest)
	)

	for i := range ret {
		ret[i] = Zipped[A, B]{
			A: a[i],
			B: b[i],
		}
	}

	return ret
}
