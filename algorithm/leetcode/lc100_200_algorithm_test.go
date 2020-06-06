package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCalculateMinimumHP(t *testing.T) {
	dungeon := [][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}
	// dungeon := [][]int{{2, 0, -1}}
	res := calculateMinimumHP(dungeon)
	t.Log(res)

}

func TestMinStackV2_Pop(t *testing.T) {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Printf("%d\n", minStack.GetMin())
	minStack.Pop()
	fmt.Printf("%d\n", minStack.Top())
	fmt.Printf("%d\n", minStack.GetMin())

}

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test01",
			args: args{root: &TreeNode{Val: 3, Left: &TreeNode{Val: 9},
				Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}},
			want: [][]int{{3}, {9, 20}, {15, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("value=%v", levelOrder(tt.args.root))
			if got := levelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
