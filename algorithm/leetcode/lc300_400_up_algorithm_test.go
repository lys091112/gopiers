package leetcode

import "testing"

func TestPalindromePairs(t *testing.T) {

	words := []string{"abcd", "dcba", "lls", "s", "sssll"}
	result := palindromePairs(words)
	t.Log(result)
	words = []string{"bat", "tab", "cat"}
	t.Log(palindromePairs(words))

	words = []string{"", "tab"}
	t.Log(palindromePairs(words))

	words = []string{"bb", "bababab", "baab", "abaabaa", "aaba", "", "bbaa", "aba", "baa", "b"}
	t.Log(palindromePairs(words))
	t.Log(isPalindrome("bababab"))
}
