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

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test01",
			args: args{buildListNodeWithCycle([]int{3, 2, 0, -4}, 1)},
			want: true,
		},
		{
			name: "test02",
			args: args{buildListNodeWithCycle([]int{1, 2}, 0)},
			want: true,
		},
		{
			name: "test03",
			args: args{buildListNodeWithCycle([]int{1}, -1)},
			want: false,
		},
		{
			name: "test04",
			args: args{buildListNodeWithCycle([]int{1}, 0)},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_preorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test01",
			args: args{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_majorityElementImprove(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test01",
			args:args{nums:[]int{2,2,1,3,1,1,4,1,1,5,1,1,6}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElementImprove(tt.args.nums); got != tt.want {
				t.Errorf("majorityElementImprove() = %v, want %v", got, tt.want)
			}
		})
	}
}
