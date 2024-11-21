package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func TestAddTwoNumbers(t *testing.T) {
	r1 := buildListNode([]int{5})
	r2 := buildListNode([]int{5})

	res := addTwoNumbers(r1, r2)
	printListNode(res, t)

}

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

func Test_search33(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test01",
			args: args{nums: []int{5, 1, 3}, target: 3},
			want: 2,
		},
		{
			name: "test02",
			args: args{nums: []int{1, 3, 5}, target: 3},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search33(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_myPow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test01",
			args: args{x: 2.0000, n: 10},
			want: 1024.0000,
		},
		{
			name: "test02",
			args: args{x: 2.0000, n: -2},
			want: 0.2500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPow(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("myPow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_spiralOrder(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// {
		// 	name: "spiralOrder01",
		// 	args: args{matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
		// 	want: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		// },
		{
			name: "spiralOrder02",
			args: args{matrix: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}},
			want: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiralOrderV2(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "jump01",
			args: args{nums: []int{2, 3, 1, 1, 4}},
			want: 2,
		},
		{
			name: "jump02",
			args: args{nums: []int{1, 2, 3}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jumpV2(tt.args.nums); got != tt.want {
				t.Errorf("jump() = %v, want %v", got, tt.want)
			}
		})
	}
}
