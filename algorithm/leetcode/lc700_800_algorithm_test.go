package leetcode

import (
	"fmt"
	"testing"
)

func TestLongestWord(t *testing.T) {

	words := []string{"a", "banana", "app", "appl", "ap", "apply", "apple"}

	res := longestWord(words)

	if res == "apple" {
		t.Log("success")
	} else {
		t.Error("failed")
	}

}

func TestIsBipartite(t *testing.T) {
	// graph := [][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}
	graph := [][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}
	r := isBipartite(graph)
	t.Log(r)
}

func TestCrackSafe(t *testing.T) {
	res := crackSafe(3, 4)
	fmt.Println(res)
}

func TestCustomerSortString(t *testing.T) {
	type args struct {
		order string
		s     string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{order: "cba", s: "abcd"},
			want: "cbad",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CustomerSortString(tt.args.order, tt.args.s); got != tt.want {
				t.Errorf("CustomerSortString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reachNumber(t *testing.T) {
	type args struct {
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "t1",
			args: args{target: 2},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reachNumber(tt.args.target); got != tt.want {
				t.Errorf("reachNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
