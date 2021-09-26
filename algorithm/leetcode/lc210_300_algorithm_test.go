package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindWords(t *testing.T) {
	words := []string{"oath", "pea", "eat", "rain"}
	board := [][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}

	res := findWords(board, words)

	for _, v := range res {
		println(v)
	}
}

func TestFindWords2(t *testing.T) {
	words := []string{"bbaabaabaaaaabaababaaaaababb", "aabbaaabaaabaabaaaaaabbaaaba", "babaababbbbbbbaabaababaabaaa", "bbbaaabaabbaaababababbbbbaaa", "babbabbbbaabbabaaaaaabbbaaab", "bbbababbbbbbbababbabbbbbabaa", "babababbababaabbbbabbbbabbba", "abbbbbbaabaaabaaababaabbabba", "aabaabababbbbbbababbbababbaa", "aabbbbabbaababaaaabababbaaba", "ababaababaaabbabbaabbaabbaba", "abaabbbaaaaababbbaaaaabbbaab", "aabbabaabaabbabababaaabbbaab", "baaabaaaabbabaaabaabababaaaa", "aaabbabaaaababbabbaabbaabbaa", "aaabaaaaabaabbabaabbbbaabaaa", "abbaabbaaaabbaababababbaabbb", "baabaababbbbaaaabaaabbababbb", "aabaababbaababbaaabaabababab", "abbaaabbaabaabaabbbbaabbbbbb", "aaababaabbaaabbbaaabbabbabab", "bbababbbabbbbabbbbabbbbbabaa", "abbbaabbbaaababbbababbababba", "bbbbbbbabbbababbabaabababaab", "aaaababaabbbbabaaaaabaaaaabb", "bbaaabbbbabbaaabbaabbabbaaba", "aabaabbbbaabaabbabaabababaaa", "abbababbbaababaabbababababbb", "aabbbabbaaaababbbbabbababbbb", "babbbaabababbbbbbbbbaabbabaa"}
	board := [][]byte{{'b', 'a', 'a', 'b', 'a', 'b'}, {'a', 'b', 'a', 'a', 'a', 'a'}, {'a', 'b', 'a', 'a', 'a', 'b'}, {'a', 'b', 'a', 'b', 'b', 'a'},
		{'a', 'a', 'b', 'b', 'a', 'b'}, {'a', 'a', 'b', 'b', 'b', 'a'}, {'a', 'a', 'b', 'a', 'a', 'b'}}

	res := findWords(board, words)

	for _, v := range res {
		println(v)
	}
}
func TestFindWords3(t *testing.T) {
	words := []string{"aaa"}
	board := [][]byte{{'a', 'a'}}

	res := findWords(board, words)

	for _, v := range res {
		println(v)
	}
}
func TestFindWords4(t *testing.T) {
	words := []string{"a"}
	board := [][]byte{{'a', 'a'}}

	res := findWords(board, words)

	for _, v := range res {
		println(v)
	}
}

func Test_maximalSquare(t *testing.T) {
	type args struct {
		matrix [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{matrix: [][]byte{{1, 0, 1, 0, 0},
				{1, 0, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 0, 0, 1, 0}},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalSquare(tt.args.matrix); got != tt.want {
				t.Errorf("maximalSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findOrder(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "findOrder01",
			args: args{numCourses: 4,
				prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}},
			want: []int{0, 2, 1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOrderV2(tt.args.numCourses, tt.args.prerequisites); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isAnagram(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test01",
			args: args{s: "anagram", t: "nagaram"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
	fmt.Print(int('N' - 'a'))
}
