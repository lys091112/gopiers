package leetcode

import "testing"

func TestCanReorderDoubled(t *testing.T) {
	res := canReorderDoubled([]int{-3, -6})
	if res {
		t.Log("success")
	} else {
		t.Errorf("excepted=%t, actual=%t\n", false, res)
	}
}

func TestValidateStackSequences(t *testing.T) {
	pushed := []int{8, 2, 1, 4, 7, 9, 0, 3, 5, 6}
	popped := []int{1, 2, 7, 3, 6, 4, 0, 9, 5, 8}
	res := validateStackSequences(pushed, popped)

	if !res {
		t.Log("success")
	} else {
		t.Error("failed")
	}
}

func TestLargestComponentSize(t *testing.T) {
	s := largestComponentSize([]int{65, 35, 43, 76, 15, 11, 81, 22, 55, 92, 31})
	resultOfInt(9, s, t)
}
