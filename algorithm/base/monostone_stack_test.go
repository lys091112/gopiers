package base

import (
	"reflect"
	"testing"
)

func TestMonostoneStack(t *testing.T) {
	type args struct {
		q []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "下一个小的元素",
			args: args{q: []int{1, 3, 4, 5, 2, 9, 6}},
			want: []int{-1, 4, 4, 4, -1, 6, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MonostoneStack(tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MonostoneStack() = %v, want %v", got, tt.want)
			}
		})
	}
}
