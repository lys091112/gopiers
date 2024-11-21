package leetcode

import (
	"testing"
)

func TestFindFirstCompleteIndex(t *testing.T) {
	type args struct {
		arr []int
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "empty input",
		// 	args: args{
		// 		arr: []int{},
		// 		mat: [][]int{},
		// 	},
		// 	want: -1,
		// },
		// {
		// 	name: "different dimensions",
		// 	args: args{
		// 		arr: []int{1},
		// 		mat: [][]int{{1}},
		// 	},
		// 	want: 0,
		// },
		{
			name: "first complete index",
			args: args{
				arr: []int{1, 4, 5, 2, 6, 1},
				mat: [][]int{
					{4, 3, 5},
					{1, 2, 6},
				},
			},
			want: 1,
		},
		// {
		// 	name: "first complete 乱序",
		// 	args: args{
		// 		arr: []int{2, 8, 7, 4, 1, 3, 5, 6, 9},
		// 		mat: [][]int{
		// 			{3, 2, 5},
		// 			{1, 4, 6},
		// 			{8, 7, 9},
		// 		},
		// 	},
		// 	want: 3,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindFirstCompleteIndex(tt.args.arr, tt.args.mat); got != tt.want {
				t.Errorf("findFirstCompleteIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
