package leetcode

// 正则匹配
/**
 * 动态规划
 * 初始化
 *

 * TODO 优化 由于每行数据只使用了当前位置的前两行，因此可以进行空间优化
 */

func isMatch(s string, p string) bool {

	if len(s) <= 0 || len(p) <= 0 || p[0] == '*' {
		return false
	}

	s1 := []byte(s)
	p1 := []byte(p)

	m, n := len(s1), len(p1)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true

	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] != '*' {
				if i > 0 && (s1[i-1] == p[j-1] || p[j-1] == '.') {
					dp[i][j] = dp[i-1][j-1]
				}
			} else {

				// repeat 0 times
				dp[i][j] = dp[i][j-2]

				// repeat at least 1 times
				if i > 0 && (s1[i-1] == p[j-2] || p[j-2] == '.') {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}

			}
		}
	}

	return dp[m][n]
}
