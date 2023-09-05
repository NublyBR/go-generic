package generic

import "testing"

func TestReverse(t *testing.T) {
	var inp = []int{1, 2, 3, 4, 5}
	var exp = []int{5, 4, 3, 2, 1}

	var ret = Reverse(inp)

	var equals = true

	if len(ret) == len(exp) {
		for i, e := range exp {
			if ret[i] != e {
				equals = false
				break
			}
		}
	} else {
		equals = false
	}

	if !equals {
		t.Errorf("expected Reverse(%+v) to be %+v, got %+v", inp, exp, ret)
	}
}

func TestReverseInPlace(t *testing.T) {
	var inp = []int{1, 2, 3, 4, 5}
	var exp = []int{5, 4, 3, 2, 1}

	var ret = make([]int, len(inp))
	copy(ret, inp)
	ReverseInPlace(ret)

	var equals = true

	if len(ret) == len(exp) {
		for i, e := range exp {
			if ret[i] != e {
				equals = false
				break
			}
		}
	} else {
		equals = false
	}

	if !equals {
		t.Errorf("expected ReverseInPlace(%+v) to be %+v, got %+v", inp, exp, ret)
	}
}
