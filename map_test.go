package generic

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	var inp = []int{1, 2, 3}
	var exp = []string{"1", "2", "3"}

	var ret = Map(func(i int) string { return fmt.Sprint(i) }, inp...)

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
		t.Errorf("expected Map(%+v, fn) to be %q, got %q", inp, exp, ret)
	}
}
