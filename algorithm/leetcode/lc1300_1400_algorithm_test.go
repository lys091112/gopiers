package leetcode

import (
	"reflect"
	"testing"
)

func Test_arrayRankTransform(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "t1",
			args: args{[]int{40, 10, 20, 30}},
			want: []int{4, 1, 2, 3},
		},
		{
			name: "t2",
			args: args{[]int{10, 10, 10}},
			want: []int{1, 1, 1},
		},
		{
			name: "t3",
			args: args{[]int{37, 12, 28, 9, 100, 56, 80, 5, 12}},
			want: []int{5, 3, 4, 2, 8, 6, 7, 1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrayRankTransform(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("arrayRankTransform() = %v, want %v", got, tt.want)
			}
		})
	}
}
