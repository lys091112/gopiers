package leetcode

import "testing"

func TestIsMatch(t *testing.T) {
	resultOfBool(false, isMatch("aaa", "aaaa"), t)

	resultOfBool(false, isMatch("aa", "a"), t)

	resultOfBool(true, isMatch("aab", "a*b"), t)

	resultOfBool(false, isMatch("aa", "*ab"), t)
	resultOfBool(false, isMatch("aa", ".*b"), t)
	resultOfBool(true, isMatch("aab", ".*ab"), t)
	resultOfBool(true, isMatch("aba", ".*"), t)
	resultOfBool(true, isMatch("aab", "c*a*b"), t)
	resultOfBool(false, isMatch("mississippi", "mis*is*p*."), t)
}
