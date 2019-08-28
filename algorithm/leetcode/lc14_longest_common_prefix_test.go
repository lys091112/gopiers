package leetcode

import "testing"

func TestLongestCommonPrefix(t *testing.T) {

	resultOfString("fl", longestCommonPrefix([]string{"flower", "flow", "flight"}), t)
	resultOfString("", longestCommonPrefix([]string{"dog","racecar","car"}), t)
	resultOfString("a", longestCommonPrefix([]string{"a"}), t)

}
