package generic

import "testing"

func TestSortBy(t *testing.T) {
	var (
		slice    = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		expected = []int{9, 7, 5, 3, 1, 2, 4, 6, 8}

		result = SortedBy(slice, func(i int) int {
			if i%2 == 0 {
				return i
			}
			return -i
		})
	)

	if !EqualsSlice(result, expected) {
		t.Errorf("expected SortedBy(%+v) to be %+v, got %+v", slice, expected, result)
	}
}
