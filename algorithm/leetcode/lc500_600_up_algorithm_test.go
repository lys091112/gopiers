package leetcode

import "testing"

func TestSingleNonDuplicate(t *testing.T) {
	//res := singleNonDuplicate([]int{3, 3, 7, 7, 10, 11, 11})
	//t.Log("target: ", res)

	res := singleNonDuplicate([]int{3, 3, 7, 7, 10, 11, 11, 5, 5})
	t.Log("target: ", res)

}

func TestFindMaximizedCapital(t *testing.T) {
	res := findMaximizedCapital(2, 0, []int{1, 2, 3}, []int{0, 1, 1})
	t.Logf("res=%d", res)
}
