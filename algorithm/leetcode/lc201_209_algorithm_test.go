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

func Test_isHappy(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test01",
			args: args{n: 18},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHappy(tt.args.n); got != tt.want {
				t.Errorf("isHappy() = %v, want %v", got, tt.want)
			}
		})
	}
}
