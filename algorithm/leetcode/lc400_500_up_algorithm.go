package leetcode

import "github.com/lys091112/gopiers/algorithm/util"

/*
*

  - N:474 1 和 0

  - 动态规划问题

  - count[i][m][n] = max(count[i-1][m][n], 1 + count[i-1][m-len(i)][n-len(i)])

  - dp[i][j] = max(dp(i,j), 1 + dp[i-zero[i][j-one[j]]] )

  - 回溯的算法思路

  - d[i][j][k]表示从[0,...,i]中选用字符串，0的个数最多为j个,1的个数最多为k个时，字符串最大个数。

  - dfs(i,j,k){

  - if(i<0 && j>=0 && k>=0)}{

  - return 0;

  - }

  - if(i<0 || j<0 || k<0){

  - return 负无穷;

  - }

  - a = dfs(i-1,j,k);//不选字符串i

  - b = dfs(i-1,j - 字符串i中0的个数, k - 字符串i中1的个数) + 1;

  - return max(a,b);

  - }
    *
*/
func findMaxForm(strs []string, m int, n int) int {

	dp := util.InitArray(m+1, n+1)
	dp[0][0] = 0
	for _, v := range strs {
		zeroCount, oneCount := countOfZeroAndOne(v)
		for i := m; i >= zeroCount; i-- {
			for j := n; j >= oneCount; j-- {
				dp[i][j] = util.Max(dp[i][j], 1+dp[i-zeroCount][j-oneCount])
			}
		}
	}

	return dp[m][n]
}

func countOfZeroAndOne(s string) (zeroCount, oneCount int) {
	if len(s) <= 0 {
		return 0, 0
	}
	zeroCount, oneCount = 0, 0
	for _, v := range s {
		if v == '0' {
			zeroCount++
		} else {
			oneCount++
		}
	}
	return
}
