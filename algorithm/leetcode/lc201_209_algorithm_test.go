package leetcode

import (
	"fmt"
	"testing"
)

func TestCanFinish(t *testing.T) {
	res := canFinish(3, [][]int{{0, 2}, {1, 2}, {2, 0}})

	if res {
		t.Log("success")
	} else {
		t.Error("failed")
	}
}

func TestRightSideView(t *testing.T) {
	// [1,2,3,null,5,null,4]
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}}
	r := rightSideView(root)
	fmt.Print(r)
}
