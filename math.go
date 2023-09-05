package generic

// Sum returns the sum of all arguments.
func Sum[T orderable](args ...T) T {
	var sum T

	for _, elem := range args {
		sum += elem
	}

	return sum
}

// Max returns the maximum value among the arguments.
func Max[T orderable](args ...T) T {
	var max T

	if len(args) == 0 {
		return max
	}

	max = args[0]

	for _, elem := range args[1:] {
		if elem > max {
			max = elem
		}
	}

	return max
}

// Min returns the minimum value among the arguments.
func Min[T orderable](args ...T) T {
	var min T

	if len(args) == 0 {
		return min
	}

	min = args[0]

	for _, elem := range args[1:] {
		if elem < min {
			min = elem
		}
	}

	return min
}

// Average calculates the average value of the arguments.
func Average[T numeric](args ...T) T {
	var sum T

	for _, elem := range args {
		sum += elem
	}

	return sum / T(len(args))
}

// Clamp restricts a value to be within the range [min, max].
func Clamp[T orderable](value, min, max T) T {
	switch {
	case value < min:
		return min
	case value > max:
		return max
	default:
		return value
	}
}
