package leetcode

import "testing"

func TestLongestWord(t *testing.T) {

	words := []string{"a", "banana", "app", "appl", "ap", "apply", "apple"}

	res := longestWord(words)

	if res == "apple" {
		t.Log("success")
	} else {
		t.Error("failed")
	}

}
