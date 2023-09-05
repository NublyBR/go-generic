package generic

import "testing"

func TestZip(t *testing.T) {
	var inp = [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 0},
	}
	var exp = [][]int{
		{1, 6},
		{2, 7},
		{3, 8},
		{4, 9},
		{5, 0},
	}

	var ret = Zip(inp...)

	if !EqualsSlice2(ret, exp) {
		t.Errorf("Expected zip(%+v) to equal %+v, got %+v", inp, exp, ret)
	}
}
