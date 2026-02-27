package generic

// Sum returns the sum of all arguments.
func Sum[T Orderable](args ...T) T {
	var sum T

	for _, elem := range args {
		sum += elem
	}

	return sum
}

// Max returns the maximum value among the arguments.
func Max[T Orderable](args ...T) T {
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
func Min[T Orderable](args ...T) T {
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
//
// This function will round implicitly with integer types.
func Average[T Numeric](args ...T) T {
	var sum T

	for _, elem := range args {
		sum += elem
	}

	return sum / T(len(args))
}

// Clamp restricts a value to be within the range [min, max].
func Clamp[T Orderable](value, min, max T) T {
	switch {
	case value < min:
		return min
	case value > max:
		return max
	default:
		return value
	}
}

// MapRange linearly maps a value from the range [inMin, inMax] to the range [outMin, outMax].
func MapRange[T Numeric](value, inMin, inMax, outMin, outMax T) T {
	switch {
	case value <= inMin:
		return outMin
	case value >= inMax:
		return outMax
	default:
		return outMin + (value-inMin)*(outMax-outMin)/(inMax-inMin)
	}
}

func Abs[T Signed](value T) T {
	var zero T

	if value < zero {
		return -value
	}

	return value
}
