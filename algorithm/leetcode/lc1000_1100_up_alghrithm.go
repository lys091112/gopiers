package leetcode

import "math"

/**
 * N: 1002
 */
func bitwiseComplement(N int) int {
	// 这里的边界考虑
	if N <= 0 {
		return 1
	}

	b := convertToBinaryStr(N)

	res := 0
	// 找到第一个停止的位置
	for i, v := range b {
		if v == '0' {
			res = res + int(math.Pow(float64(2), float64(len(b)-1-i)))
		}

	}
	return res
}
