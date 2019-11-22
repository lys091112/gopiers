package leetcode

import "testing"

func TestClimbStairs_70(t *testing.T) {
	c := climbStairs_70(5)
	if c == 8 {
		t.Log("success")
	} else {
		t.Errorf("failed,except=%d actual=%d", 8, c)
	}

}

func TestIsValidBST(t *testing.T) {
	// [3,1,5,0,2,4,6]
	//root := &TreeNode{Val: 3, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 2}},
	//	Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 6}}}

	//[-2147483648,null,2147483647]
	root := &TreeNode{Val: -2147483648, Right: &TreeNode{Val: 2147483647}}
	res := isValidBST(root)

	if res {
		t.Log("success")
	} else {
		t.Error("test failed!")
	}
}

func TestSolveNQueens(t *testing.T) {
	solveNQueens(4)
}
