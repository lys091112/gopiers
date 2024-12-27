package base

import (
	"math"

	"github.com/lys091112/gopiers/algorithm/util"
)

// 创建ST表
func BuildST(nums []int) [][]int {
	n := len(nums)
	w := int(math.Log(float64(n)))
	dp := util.InitArray(n, w+1)
	for i, v := range nums {
		dp[i][0] = v
	}
	for j := 1; j <= w; j++ {
		// 保证可以覆盖到[l,n]的范围
		for i := 0; i+(1<<j)-1 < n; i++ {
			dp[i][j] = util.Max(dp[i][j-1], dp[i+(1<<(j-1))][j-1])
		}
	}
	return dp
}

// 查询ST表元素
func QueryST(dp [][]int, l, r int) int {
	// 2^k <= r-l+1 < 2^(k+1) k是小于区间长度，且满足2的k
	k := int(math.Log(float64(r - l + 1)))
	return util.Max(dp[l][k], dp[r-(1<<k)+1][k])
}
