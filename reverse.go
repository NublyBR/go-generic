package generic

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
