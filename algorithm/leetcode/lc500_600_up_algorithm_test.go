package leetcode

import (
	"testing"
)

func TestSingleNonDuplicate(t *testing.T) {
	//res := singleNonDuplicate([]int{3, 3, 7, 7, 10, 11, 11})
	//t.Log("target: ", res)

	res := singleNonDuplicate([]int{3, 3, 7, 7, 10, 11, 11, 5, 5})
	t.Log("target: ", res)

}

func TestFindMaximizedCapital(t *testing.T) {
	res := findMaximizedCapital(2, 0, []int{1, 2, 3}, []int{0, 1, 1})
	t.Logf("res=%d", res)
}

func Test_subarraySum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "subarraySum01",
			args: args{nums: []int{1, 1, 1}, k: 2},
			want: 2,
		},
		{
			name: "subarraySum02",
			args: args{nums: []int{1, 2, 4, 6, 3}, k: 10},
			want: 1,
		},
		{
			name: "subarraySum03",
			args: args{nums: []int{-1,-1, 1}, k: 0},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraySum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("subarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
