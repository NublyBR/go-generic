package generic

import (
	"testing"
)

func TestMapTopN(t *testing.T) {
	var (
		mp = map[string]int{
			"a": 1,
			"b": 3,
			"c": 2,
			"d": 5,
			"e": 4,
		}

		expectMap = map[string]int{
			"d": 5,
			"e": 4,
			"b": 3,
		}

		expectItems = Items[string, int]{
			{"d", 5},
			{"e", 4},
			{"b", 3},
		}

		topN = AsItems(mp).TopN(3, func(lhs, rhs Item[string, int]) bool {
			return lhs.Value > rhs.Value
		})
	)

	if mp := topN.Map(); !EqualsMap(mp, expectMap) {
		t.Errorf("expected result to be %v, got %v", expectMap, mp)
	}

	if !EqualsSlice(expectItems, topN) {
		t.Errorf("expected result to be %v, got %v", expectItems, topN)
	}
}
