package leetcode

import "testing"

func TestPalindromePairs(t *testing.T) {

	words := []string{"abcd", "dcba", "lls", "s", "sssll"}
	result := PalindromePairs(words)
	t.Log(result)
	words = []string{"bat", "tab", "cat"}
	t.Log(PalindromePairs(words))

	words = []string{"", "tab"}
	t.Log(PalindromePairs(words))

	words = []string{"bb", "bababab", "baab", "abaabaa", "aaba", "", "bbaa", "aba", "baa", "b"}
	t.Log(PalindromePairs(words))
	t.Log(isPalindrome("bababab"))
}
