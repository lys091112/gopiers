package leetcode

import "math"

func longestCommonPrefix(strs []string) string {

	if len(strs) < 1 {
		return ""
	}

	min := math.MaxInt32
	var tmp string
	for i := range strs {
		if min > len(strs[i]) {
			min = len(strs[i])
			tmp = strs[i]
		}
	}

	targetIdx := -1

Loop:
	for i := range tmp {
		for _, v := range strs {
			if tmp[i] != v[i] {
				break Loop
			}
		}
		targetIdx = i
	}

	return tmp[0 : targetIdx+1]
}
