package generic

import "testing"

func TestMax(t *testing.T) {
	if max := Max(1, 2, 3, 4); max != 4 {
		t.Errorf("expected for Max(1, 2, 3, 4) to be 4, got %d", max)
	}
}

func TestMin(t *testing.T) {
	if min := Min(1, 2, 3, 4); min != 1 {
		t.Errorf("expected for Min(1, 2, 3, 4) to be 1, got %d", min)
	}
}

func TestAverage(t *testing.T) {
	if avg := Average(0, 10, 20, 30); avg != 15 {
		t.Errorf("expected for Average(0, 10, 20, 30) to be 15, got %d", avg)
	}
}
