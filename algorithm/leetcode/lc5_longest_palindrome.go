package leetcode

//思路：动态规划;
// 参考： https://www.cnblogs.com/fsmly/p/10048008.html
// 动态规划最优解的子问题同样也是最优解，即最长回文字符串的子串也是回文串，
// 假设p[i][j]是回文字符串，那么p[i+1][j-1]也是回文字符串，这样最长回文字符串就能够分解成为一系列子问题来求解
func longestPalindrome(s string) string {
	length := len(s)
	if length <= 0 {
		return ""
	}

	dep := make([][]int, length)
	for i := 0; i < length; i++ {
		dep[i] = make([]int, length)
	}

	start := 0
	maxLen := 1
	arr := []byte(s)
	for i := range arr {
		dep[i][i] = 1
		if i < length-1 && arr[i] == arr[i+1] {
			dep[i][i+1] = 1
			maxLen = 2
			start = i
		}
	}

	for i := 2; i < length; i++ {
		for j := 0; j < i; j++ {

			m := i - j + 1
			if dep[i-1][j+1] == 1 && (arr[i] == arr[j]) {
				dep[i][j] = 1
				if m >= maxLen {
					maxLen = m
					start = j
				}
			}
		}
	}

	return s[start : start+maxLen]
}
