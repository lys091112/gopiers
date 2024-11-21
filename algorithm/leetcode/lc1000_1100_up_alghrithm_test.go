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
			name: "test01",
			args: args{s: "cbacdcbc"},
			want: "acdb",
		},
		{
			name: "test02",
			args: args{s: "leetcode"},
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

func Test_shortestCommonSupersequence(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			args: args{str1: "abac", str2: "cab"},
			want: "cabac",
		},
		{
			name: "test02",
			args: args{str1: "abaxlisfc", str2: "1l12cab"},
			want: "1l12cabaxlisfc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestCommonSupersequence(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("shortestCommonSupersequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCarPolling(t *testing.T) {
	type args struct {
		trips    [][]int
		capacity int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "c1",
			args: args{trips: [][]int{{2, 1, 5}, {3, 3, 7}}, capacity: 5},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CarPolling(tt.args.trips, tt.args.capacity); got != tt.want {
				t.Errorf("CarPolling() = %v, want %v", got, tt.want)
			}
		})
	}
}
