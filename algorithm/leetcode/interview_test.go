package leetcode

import (
	"reflect"
	"testing"
)

func Test_missingTwo(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			args: args{nums: []int{2, 3}},
			want: []int{1, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingTwo(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("missingTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minHeapR1719(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "minHeap01",
			args: args{k: 5},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minHeapR1719(tt.args.k); got != tt.want {
				t.Errorf("minHeapR1719() = %v, want %v", got, tt.want)
			}
		})
	}
}
