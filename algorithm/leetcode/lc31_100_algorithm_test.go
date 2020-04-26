package leetcode

import (
	"fmt"
	"testing"
)

func TestClimbStairs_70(t *testing.T) {
	c := climbStairs_70(5)
	if c == 8 {
		t.Log("success")
	} else {
		t.Errorf("failed,except=%d actual=%d", 8, c)
	}

}

func TestMinPathSum(t *testing.T) {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	way := minPathSum(grid)
	t.Log(way)

}

func TestLongestValidParentheses(t *testing.T) {
	s := "(()())"
	t.Log(longestValidParentheses(s))
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

// N: 46
func TestPermute(t *testing.T) {
	r := permute([]int{1, 2, 3})

	fmt.Printf("%v\n", r)
}

func TestTrap(t *testing.T) {
	// height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	// height := []int{5, 2, 1, 2, 1, 5}
	height := []int{2, 0, 2}

	if trapV2(height) == 6 {
		t.Log("success")
	} else {
		t.Error("error")
	}
}

func TestSolveNQueens(t *testing.T) {
	solveNQueens(4)
}

func TestMinDistance(t *testing.T) {
	l := minDistance("horse", "ros")
	t.Log(l)
}

func TestLargestRectangleArea(t *testing.T) {
	area := largestRectangleArea([]int{4, 2, 0, 3, 2, 5})
	if area == 6 {
		t.Log("success")
	} else {
		t.Error("failed")
	}

	area = largestRectangleArea([]int{3, 6, 5, 7, 4, 8, 1, 0})
	if area == 20 {
		t.Log("success")
	} else {
		t.Error("failed")
	}

}
