package leetcode

import (
	"strconv"
	"strings"
)

// 给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转
func reverse2(x int) int {
	var flag = 1
	str := strconv.Itoa(x)
	if x < 0 {
		flag = -1
		str = str[1:]
	}

	return flag * convert(str)

}

func convert(str string) int {

	maxStr := "2147483648"

	arr := []byte(str)

	for i := range arr {
		j := len(str) - i - 1
		if i >= j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
	}

	k := 0
	for i := range arr {
		if arr[i] == 0 && i < (len(arr)-1) {
			k = i + 1
		}
		break
	}
	str = string(arr[k:])

	if len(str) >= len(maxStr) && strings.Compare(str, maxStr) > 0 {
		return 0
	} else {
		res, err := strconv.Atoi(str)
		if err != nil {
			return 0
		}
		return res
	}

}
