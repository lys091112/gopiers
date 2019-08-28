package leetcode

import (
	"bytes"
	"strconv"
)

/**
 * int to binary string
 * n > 0
 */
func convertToBinaryStr(n int) string {
	if n == 0 {
		return "0"
	}

	var res bytes.Buffer
	for {
		if n <= 0 {
			break
		}

		res.WriteString(strconv.Itoa(n % 2))
		n = n >> 1
	}
	return reverseString(res.String())
}
