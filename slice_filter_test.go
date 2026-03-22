package generic

import "testing"

func TestUnique(t *testing.T) {
	var (
		slice    = []int{1, 1, 2, 3, 2, 4, 3, 3, 5, 5}
		expected = []int{1, 2, 3, 4, 5}

		result = Unique(slice)
	)

	if !EqualsSlice(result, expected) {
		t.Errorf("expected Unique(%+v) to be %+v, got %+v", slice, expected, result)
	}
}

func TestUniqueBy(t *testing.T) {
	var (
		slice    = []int{1, -1, 2, -3, -2, 4, 3, -3, 5, 5}
		expected = []int{1, 2, -3, 4, 5}

		result = UniqueBy(slice, func(i int) int {
			return Abs(i)
		})
	)

	if !EqualsSlice(result, expected) {
		t.Errorf("expected UniqueBy(%+v, ...) to be %+v, got %+v", slice, expected, result)
	}
}

func TestOnce(t *testing.T) {
	var (
		slice    = []int{1, 1, 2, 3, 2, 4, 3, 3, 5, 5}
		expected = []int{1, 2, 3, 4, 5}

		once   Once[int]
		result []int
	)

	for _, item := range slice {
		if once.First(item) {
			result = append(result, item)
		}
	}

	if !EqualsSlice(result, expected) {
		t.Errorf("expected result to be %+v, got %+v", expected, result)
	}
}
