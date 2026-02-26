package generic

type Iterator[T any] func() (T, bool)

func NewIterator[T any](slice []T) Iterator[T] {
	var (
		idx int
		ln  = len(slice)
	)

	return func() (ret T, ok bool) {
		if idx >= ln {
			return
		}

		ret = slice[idx]
		ok = true
		idx++

		return
	}
}
