package leetcode

// 反转字符串
func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}


// 是否回文字符串
func isPalindrome(word string) bool {
	if len(word) <= 1 {
		return true
	}

	for i := 0; i < len(word)/2; i ++ {
		if word[i] != word[len(word)-i-1] {
			return false
		}
	}
	return true
}
