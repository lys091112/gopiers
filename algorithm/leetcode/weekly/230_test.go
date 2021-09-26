package weekly

import (
	"testing"
)

func Test_minOperations(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "test01",
		// 	args: args{nums1:[]int{1,2,3,4,5,6},nums2:[]int{1,1,2,2,2,2}},
		// 	want:3,
		// },
		// {
		// 	name: "test02",
		// 	args: args{nums1:[]int{1,1,1,1,1,1,1},nums2:[]int{6}},
		// 	want:-1,
		// },
		{
			name: "test03",
			args: args{nums1: []int{6, 6}, nums2: []int{1}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_closestCost(t *testing.T) {
	type args struct {
		baseCosts    []int
		toppingCosts []int
		target       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name:"test01",
			args: args{baseCosts:[]int{2,3},toppingCosts:[]int{4,5,100},target:18},
			want: 17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closestCost(tt.args.baseCosts, tt.args.toppingCosts, tt.args.target); got != tt.want {
				t.Errorf("closestCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
