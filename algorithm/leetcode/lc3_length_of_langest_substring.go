package leetcode

import (
	"strings"
)

// 最长不重复子串
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	return len(longestSubString(s))
}

func longestSubString(s string) string {
	store := make(map[rune]int)
	startIdx := 0 // 记录字符串遍历的当前点
	maxLength := 0 // 记录字符串最大长度
	maxStartIdx := 0 //记录字符串最大长度的起始点

	st := []rune(s)
	for idx, v := range st {

		// 如果出现重复元素，则从第一个重复元素的下一位开始，动态更新起始位置
		if value, ok := store[v]; ok && startIdx <= value {
			startIdx = value + 1
		}

		// 动态更新当前字符串的最大长度，以及最大长度的起始位置
		if idx+1-startIdx > maxLength {
			maxLength = idx + 1 - startIdx
			maxStartIdx = startIdx
		}

		// 记录元素的位置
		store[v] = idx
	}

	var builder strings.Builder
	for _, v := range st[maxStartIdx : maxStartIdx+maxLength] {
		builder.WriteRune(v)
	}

	return builder.String()
}
