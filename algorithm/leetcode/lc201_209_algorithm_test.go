package leetcode

import "testing"

func TestCanFinish(t *testing.T) {
	res := canFinish(3, [][]int{{0, 2}, {1, 2}, {2, 0}})

	if res {
		t.Log("success")
	} else {
		t.Error("failed")
	}
}
