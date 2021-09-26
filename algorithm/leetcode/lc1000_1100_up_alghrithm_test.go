package leetcode

import (
	"testing"
)

func TestBitwiseComplement(t *testing.T) {
	res := bitwiseComplement(5)
	t.Log(res)
	t.Log(convertToBinaryStr(5))
}

func TestLastStoneWeight(t *testing.T) {
	var keys = []int{2, 7, 4, 1, 8, 1}
	t.Log(lastStoneWeight(keys))
}

func Test_smallestSubsequence(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name:"test01",
			args: args{s:"cbacdcbc"},
			want: "acdb",
		},
		{
			name:"test02",
			args: args{s:"leetcode"},
			want: "letcod",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestSubsequence(tt.args.s); got != tt.want {
				t.Errorf("smallestSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
