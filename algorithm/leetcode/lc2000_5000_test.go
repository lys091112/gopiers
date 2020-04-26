package leetcode

import (
	"reflect"
	"testing"
)

func Test_maxScore(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test01",
			args: args{s: "011101"},
			want: 5,
		},
		{
			name: "test02",
			args: args{s: "00111"},
			want: 5,
		},
		{
			name: "test03",
			args: args{s: "00"},
			want: 1,
		},
		{
			name: "test04",
			args: args{s: "1111"},
			want: 3,
		},
		{
			name: "test05",
			args: args{s: "010"},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScore(tt.args.s); got != tt.want {
				t.Errorf("maxScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findDiagonalOrder(t *testing.T) {
	type args struct {
		nums [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			args: args{nums: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
		},
		{
			name: "test02",
			args: args{nums: [][]int{{1, 2, 3, 4, 5}, {6, 7}, {8}, {9, 10, 11}, {12, 13, 14, 15, 16}}},
			want: []int{1, 6, 2, 8, 7, 3, 9, 4, 12, 10, 5, 13, 11, 14, 15, 16},
		},
		{
			name: "test03",
			args: args{nums: [][]int{{1, 2, 3}, {4}, {5, 6, 7}, {8}, {9, 10, 11}}},
			want: []int{1, 4, 2, 5, 3, 8, 6, 9, 7, 10, 11},
		},
		{
			name: "test04",
			args: args{nums: [][]int{{20, 17, 13, 14}, {12, 6}, {3, 4}}},
			want: []int{20, 12, 17, 3, 6, 13, 4, 14},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDiagonalOrder(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDiagonalOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
