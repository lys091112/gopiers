package leetcode

import "strconv"

// 考虑清楚先决条件
// 考虑各种可能出现的特殊情况和边界情况
func myAtoi(str string) int {
	if len(str) <= 0 {
		return 0
	}

	ss := []byte(str)

	maxStr := "2147483648"
	nagetive := 1
	startNotPlug := true
	start, end := -1, -1
	for idx, v := range ss {
		if v == ' ' && startNotPlug {
			start = idx
		} else if v == '+' && startNotPlug {
			start = idx
			startNotPlug = false
		} else if v == '-' && startNotPlug {
			start = idx
			startNotPlug = false
			nagetive = -1
		} else {
			if v < 48 || v > 57 {
				// 说明没有数字
				if start > end {
					return 0
				} else {
					break
				}
			} else if v <= 57 && v >= 48 {
				if v == 48 && start >= end {
					start = idx
				}
				end = idx
				startNotPlug = false
			}
		}
	}

	if start > end {
		return 0
	}

	targetStr := str[start+1 : end+1]
	if len(targetStr) > len(maxStr) || (len(targetStr) == len(maxStr) && targetStr >= maxStr) {
		if nagetive > 0 {
			return 2147483647
		} else {
			return -2147483648
		}
	}

	i, err := strconv.Atoi(targetStr)

	if err != nil {
		return 0
	}

	return nagetive * i

}

/**
func myAtoi(str string) int {
	if len(str) <= 0 {
		return 0
	}

	ss := []byte(str)

	maxStr := "2147483648"
	nagetive := 1
	plug := false
	startNotZero := false
	start, end := -1, -1
	for idx, v := range ss {
		if (v == ' ' || v == '+') && start >= end {
			start = idx
			if v == '+' {
				if plug {
					return 0
				}
				plug = true
			}
		} else if v == '-' {
			if plug {
				return 0
			}
			start = idx
			nagetive = -1
			plug = true
		} else if v < 48 || v > 57 {
			// 说明没有数字
			if start > end {
				return 0
			} else {
				break
			}
		} else if v <= 57 && v >= 48 {
			if v == 48 && !startNotZero {
				start = idx
			} else {
				startNotZero = true
			}
			end = idx
		}
	}

	if start > end {
		return 0
	}

	targetStr := str[start+1 : end+1]
	if len(targetStr) > len(maxStr) || (len(targetStr) == len(maxStr) && targetStr >= maxStr) {
		if nagetive > 0 {
			return 2147483647
		} else {
			return -2147483648
		}
	}

	i, err := strconv.Atoi(targetStr)

	if err != nil {
		return 0
	}

	return nagetive * i

}
*/
