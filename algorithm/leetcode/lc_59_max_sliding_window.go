package leetcode

import (
	st "github.com/lys091112/gopiers/algorithm/base/basestruct"
)

// tag: ST表

// 59. 滑动窗口最大值,
// ST表(倍增思想)
// d[i][j] 代表从 i开始长度为 2^(j-1) 的最大值
// 即拆成2个长度为2^(j-1)的子区，取较大值
// 转移公式为：d[i][j] = max(d[i][j-1], d[i+2^(j-1)][j-1])
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k <= 0 {
		return []int{}
	}
	dp := st.BuildST(nums)
	res := make([]int, 0)
	for i := 0; i <= len(nums)-k; i++ {
		res = append(res, st.QueryST(dp, i, i+k-1))
	}
	return res
}

// 优先队列

// 线段树

// 单调栈
// 双端队列，简单来讲就是找下一个比当前元素更大的元素，保证队列是个递减队列
// （动态将非改窗口的元素移除）此时的队列底部一定是该滑动窗口内的最大元素
// 数学理解：如果下标 i 小于 j，且 nums[i] < nums[j]，那么 nums[i] 一定不是滑动窗口内的最大值，直接将其移除队列
func maxSlidingWindowStack(nums []int, k int) []int {
	if len(nums) == 0 || k <= 0 {
		return []int{}
	}

	res := []int{}
	stack := []int{}
	for i := 0; i < len(nums); i++ {
		// 移除栈顶非最大元素
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
			stack = stack[1:]
		}

		// 添加元素到队列
		stack = append(stack, i)

		if i >= k-1 {
			// 移除非该窗口内的无效元素以及非最大元素
			for len(stack) > 0 && stack[0] <= i-k {
				stack = stack[1:]
			}
			res = append(res, nums[stack[0]])
		}
	}

	return res
}

func sLog(n int) (k int) {
	for 1<<(k+1) <= n {
		k += 1
	}
	return k
}
