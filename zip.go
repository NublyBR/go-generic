package generic

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

	for i := 0; i < smallest; i++ {
		ret[i] = make([]T, len(args))

		for j, elem := range args {
			ret[i][j] = elem[i]
		}
	}

	return ret
}
