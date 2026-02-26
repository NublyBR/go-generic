package generic

// All checks if a given function fn returns true for all elements in the args slice of type T.
func All[T any](fn func(T) bool, args ...T) bool {
	for k := range args {
		if !fn(args[k]) {
			return false
		}
	}

	return true
}

// Any checks if a given function fn returns true for at least one element in the args slice of type T.
func Any[T any](fn func(T) bool, args ...T) bool {
	for k := range args {
		if fn(args[k]) {
			return true
		}
	}

	return false
}

// If evaluates a condition and returns either the trueVal or falseVal based on the condition.
func If[T any](cond bool, trueVal, falseVal T) T {
	if cond {
		return trueVal
	}
	return falseVal
}
