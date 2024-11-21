package leetcode

/**
 * N: 336 回文对
 * 分情况考虑各种可能性
 */
func PalindromePairs(words []string) [][]int {
	if len(words) < 2 {
		return nil
	}

	wordMap := make(map[string]int, len(words))
	for i, v := range words {
		wordMap[v] = i
	}

	result := make([][]int, 0)
	for i, v := range words {
		if v == "" {
			continue
		}
		// 查询是否包含空串 并且自身是回文串
		if j, ok := wordMap[""]; ok {
			if isPalindrome(v) {
				if i < j {
					result = append(result, []int{i, j})
					result = append(result, []int{j, i})
				} else {
					result = append(result, []int{j, i})
					result = append(result, []int{i, j})
				}
			}
		}

		// 查询自身的翻转串是否存在且非自身
		reverseWord := reverseString(v)
		if j, ok := wordMap[reverseWord]; ok {
			if i < j {
				result = append(result, []int{i, j})
				result = append(result, []int{j, i})
			}
		}

		// 查询右半部分，找出回文串切点 如 abacd 以ac之间为切点
		for k := 1; k < len(v); k++ {
			if isPalindrome(v[0:k]) {
				if j, ok := wordMap[reverseString(v[k:])]; ok {
					result = append(result, []int{j, i})
				}
			}
		}

		// 查询右半部分，找出回文串切点
		for k := 1; k < len(v); k++ {
			if isPalindrome(v[len(v)-k:]) {
				if j, ok := wordMap[reverseString(v[:len(v)-k])]; ok {
					result = append(result, []int{i, j})
				}
			}
		}

	}
	return result
}
