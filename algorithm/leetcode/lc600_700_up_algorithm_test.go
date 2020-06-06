package leetcode

import (
	"fmt"
	"testing"
)

func TestMaxSumOfThreeSubarrays(t *testing.T) {

	//w := []int{1, 2, 1, 2, 6, 7, 5, 1}
	//k := 2
	//result := maxSumOfThreeSubarrays(w, k)
	//t.Log(result)

	t.Log(maxSumOfThreeSubarrays([]int{1, 2, 1, 2, 1, 2, 1, 2, 1}, 2))

}

func TestCutOffTree(t *testing.T) {
	matrix := [][]int{
		{54581641, 64080174, 24346381, 69107959},
		{86374198, 61363882, 68783324, 79706116},
		{668150, 92178815, 89819108, 94701471},
		{83920491, 22724204, 46281641, 47531096},
		{89078499, 18904913, 25462145, 60813308},
	}

	distance := cutOffTree(matrix)
	fmt.Println(distance)
}

func Test_checkValidString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "validString01",
			args: args{s: "(*)"},
			want: true,
		},
		{
			name: "validString02",
			args: args{s: "(*))"},
			want: true,
		},
		{
			name: "validString03",
			args: args{s: "(*)))"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkValidString(tt.args.s); got != tt.want {
				t.Errorf("checkValidString() = %v, want %v", got, tt.want)
			}
		})
	}
}
