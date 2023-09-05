package generic

import "testing"

func TestCopySlice(t *testing.T) {
	var (
		slice = []int{1, 2, 3, 4, 5}

		copied = CopySlice(slice)
	)

	if !EqualsSlice(slice, copied) {
		t.Errorf("expected CopySlice(%+v) to be %+v, got %+v", slice, slice, copied)
	}
}

func TestCopyMap(t *testing.T) {
	var (
		mp = map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}

		copied = CopyMap(mp)
	)

	if !EqualsMap(mp, copied) {
		t.Errorf("expected EqualsMap(%+v) to be %+v, got %+v", mp, mp, copied)
	}
}
