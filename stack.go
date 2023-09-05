package generic

type Stack[T any] []T

func NewStack[T any](elems ...T) Stack[T] {
	return Stack[T](elems)
}

func (s *Stack[T]) Push(elem ...T) {
	*s = append(*s, elem...)
}

func (s *Stack[T]) Pop() T {
	var (
		ln  = len(*s)
		ret = (*s)[ln-1]
	)

	*s = (*s)[:ln-1]

	return ret
}

func (s *Stack[T]) PopN(n int) []T {
	var (
		ln  = len(*s)
		ret = make([]T, n)
	)

	copy(ret, (*s)[ln-n:ln])

	*s = (*s)[:ln-n]

	return ret
}

func (s *Stack[T]) Len() int {
	return len(*s)
}

func (s *Stack[T]) Clear() {
	*s = (*s)[:0]
}
