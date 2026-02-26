package generic

import "testing"

func TestIterator(t *testing.T) {
	var (
		original = []int{1, 2, 3, 4, 5}
		result   = make([]int, 0, len(original))
	)

	iterator := NewIterator(original)

	for {
		cur, ok := iterator()
		if !ok {
			break
		}
		result = append(result, cur)
	}

	if !EqualsSlice(result, original) {
		t.Errorf("expected Iterator(%+v) to be %+v, got %+v", original, original, result)
	}
}
