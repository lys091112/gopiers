package leetcode

import "testing"

func Test_minimumXORSum(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "t1", args: args{nums1: []int{2, 4}, nums2: []int{3, 1}}, want: 6},
		{name: "t2", args: args{nums1: []int{1, 2}, nums2: []int{2, 3}}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumXORSum(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("minimumXORSum() = %v, want %v", got, tt.want)
			}
			if got := minimumXORSumV2(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("minimumXORSumV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
