package leetcode

import (
	"fmt"
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

func Test_maxCoins(t *testing.T) {
	type args struct {
		piles []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "coins01",
			args: args{[]int{1, 2, 9, 5, 3, 7, 8, 4, 6}},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCoins(tt.args.piles); got != tt.want {
				t.Errorf("maxCoins() = %v, want %v", got, tt.want)
			}
			fmt.Print("xxx")
		})
	}
}

func Test_decode(t *testing.T) {
	type args struct {
		encoded []int
		first   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test01",
			args: args{encoded: []int{1, 2, 3}, first: 1},
			want: []int{1, 0, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decode(tt.args.encoded, tt.args.first); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_swapNodes(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// {
		// 	name: "test01",
		// 	args: args{head :buildListNode([]int{1,2,3,4,5}),k:3},
		// 	want: buildListNode([]int{1,4,3,2,5}),
		// },
		{
			name: "test02",
			args: args{head: buildListNode([]int{100, 90}), k: 2},
			want: buildListNode([]int{1, 4, 3, 2, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapNodes(tt.args.head, tt.args.k); true {
				printListNode(got, t)
			}
		})
	}
}

func Test_minimumTimeRequired(t *testing.T) {
	type args struct {
		jobs []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test01",
			args: args{jobs: []int{38, 49, 91, 59, 14, 76, 84}, k: 3},
			want: 140,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTimeRequired(tt.args.jobs, tt.args.k); got != tt.want {
				t.Errorf("minimumTimeRequired() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxConsecutiveAnswers(t *testing.T) {
	type args struct {
		answerKey string
		k         int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name:"test01",
			args: args{answerKey: "TTFFTT", k: 1},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxConsecutiveAnswers(tt.args.answerKey, tt.args.k); got != tt.want {
				t.Errorf("maxConsecutiveAnswers() = %v, want %v", got, tt.want)
			}
		})
	}
}
