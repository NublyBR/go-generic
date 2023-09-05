package generic

type Stack[T any] []T

// Create new empty or pre-populated stack
func NewStack[T any](elems ...T) Stack[T] {
	return Stack[T](elems)
}

// Push elements to the stack
func (s *Stack[T]) Push(elem ...T) {
	*s = append(*s, elem...)
}

// Pop an element from the stack
//
// Info: Popping from an empty stack will cause a panic
func (s *Stack[T]) Pop() T {
	var (
		ln  = len(*s)
		ret = (*s)[ln-1]
	)

	*s = (*s)[:ln-1]

	return ret
}

// Pop N elements from the stack
//
// Info: Popping from an empty stack will cause a panic
func (s *Stack[T]) PopN(n int) []T {
	var (
		ln  = len(*s)
		ret = make([]T, n)
	)

	copy(ret, (*s)[ln-n:ln])

	*s = (*s)[:ln-n]

	return ret
}

// Get current length of the stack
func (s *Stack[T]) Len() int {
	return len(*s)
}

// Empty the stack
func (s *Stack[T]) Clear() {
	*s = (*s)[:0]
}
