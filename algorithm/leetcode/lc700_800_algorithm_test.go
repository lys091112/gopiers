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
