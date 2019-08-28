package leetcode

/**
 * N :675 高尔夫砍树
 *
 * 输入二维数组， 限制条件： 树木要从低到高砍
 */
func cutOffTree(forest [][]int) int {

	// 根据树高进行排序

	// 从（0，0） 出发 到下个点

	// 计算下个点到下下个点之间的距离 dist() 计算两个坐标间的最短距离（考虑A* 或者其他方法)

	return -1
}

/**
 * N: 689  三个无重叠子数组的最大和
 * 思想： 动态规划
 * 状态转移方程：
 *    dp[i][n]=max(dp[i-1][n], dp[i-k][n-1]+sumRange(i-k+1, i));
 */
func maxSumOfThreeSubarrays(nums []int, k int) []int {

	if len(nums) < 3*k {
		return []int{}
	}

	prefixSum := make([]int, len(nums))
	prefixSum[0] = nums[0]

	// 前缀和
	for i := 1; i < len(nums); i ++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
	}

	// 动态规划 dp 记录
	dp := initArray(len(nums), 3)
	for n := 0; n < 3; n++ {
		for i := 0; i < len(nums); i ++ {
			prevL := 0
			if i-1 >= 0 {
				prevL = dp[i-1][n]
			}
			prevK := 0
			tailSum := 0
			if i-k >= 0 {
				if n > 0 {
					prevK = dp[i-k][n-1]
				}
			}
			if i-k >= 0 {
				tailSum = prefixSum[i] - prefixSum[i-k]
			} else {
				tailSum = prefixSum[i]
			}
			dp[i][n] = maxInt(prevL, prevK+tailSum)
		}
	}

	//
	result := make([]int, 0)
	curr := len(nums) - 1
	n := 2

	for {
		if n < 0 {
			break
		}
		v := dp[curr][n]
		// 因为要找最小序
		for {
			if curr > 0 && dp[curr-1][n] == v {
				curr -= 1
			} else {
				break
			}

		}

		result = append(result, curr-k+1)
		n -= 1
		curr -= k
	}

	return Reverse(result)
}
func maxInt(left, right int) int {
	if left > right {
		return left
	}
	return right
}
