package generic

import (
	"testing"
)

func TestStack(t *testing.T) {
	var stack = NewStack[int](1, 2, 3)

	if exp := []int{1, 2, 3}; !EqualsSlice(stack, exp) {
		t.Errorf("Expected stack to equal %+v, got %+v", exp, stack)
	}

	stack.Push(4, 5)

	if exp := []int{1, 2, 3, 4, 5}; !EqualsSlice(stack, exp) {
		t.Errorf("Expected stack to equal %+v, got %+v", exp, stack)
	}

	pop := stack.Pop()

	if exp := 5; pop != exp {
		t.Errorf("Expected stack.Pop() to equal %v, got %v", exp, pop)
	}

	if exp := []int{1, 2, 3, 4}; !EqualsSlice(stack, exp) {
		t.Errorf("Expected stack to equal %+v, got %+v", exp, stack)
	}

	pops := stack.PopN(2)

	if exp := []int{3, 4}; !EqualsSlice(pops, exp) {
		t.Errorf("Expected stack.PopN(2) to equal %+v, got %+v", exp, pops)
	}

	if exp := []int{1, 2}; !EqualsSlice(stack, exp) {
		t.Errorf("Expected stack to equal %+v, got %+v", exp, stack)
	}

	stack.Clear()

	if exp := []int{}; !EqualsSlice(stack, exp) {
		t.Errorf("Expected stack to equal %+v, got %+v", exp, stack)
	}
}
