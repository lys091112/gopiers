package leetcode

import (
	"container/list"
	"fmt"
	"strconv"
)

// dp[i][0] 代表左侧等于0 的个数 dp[i][1] 代表右侧等于1的个数
// dp[i][0] = (dp[i-1][0], dp[i-1][0] + 1 {if s[i] == '0'})
// dp[i][1] = (dp[i+1][1], dp[i+1][1] + 1 {if s[i+1] == '1'})
// 主要是要区分分割的临界值   1） 两个非空字符串 因此要求一定要有分割  2） 动态规划 一定要有明确的初始值
func maxScore(s string) int {

	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = []int{0, 0}
	}

	if s[0] == '0' {
		dp[0][0] = 1
	}
	dp[len(s)-1][1] = 0

	for i := 1; i < len(s)-1; i++ {
		if s[i] == '0' {
			dp[i][0] = dp[i-1][0] + 1
		} else {
			dp[i][0] = dp[i-1][0]
		}
	}

	for i := len(s) - 2; i >= 0; i-- {
		if s[i+1] == '1' {
			dp[i][1] = dp[i+1][1] + 1
		} else {
			dp[i][1] = dp[i+1][1]
		}
	}

	max := 0
	for _, v := range dp {
		fmt.Printf("%v\n", v)
		if max < v[0]+v[1] {
			max = v[0] + v[1]
		}
	}
	return max
}

// N: 5394 对角线遍历 II
// 中序遍历 + BFS
// left = (row+1,col) left = (row,col+1)
// 由于同一个节点可能被多次访问，因此需要记录
func findDiagonalOrder(nums [][]int) []int {
	if len(nums) <= 1 {
		return nums[0]
	}

	result := make([]int, 0)
	result = append(result, nums[0][0])

	stack := list.New()
	stack.PushBack([]int{0, 0})
	row := len(nums)

	way := make(map[string]bool, 0)
	way["0_0"] = true

	for {
		if stack.Len() <= 0 {
			break
		}
		top := stack.Front()
		stack.Remove(top)
		tv := top.Value.([]int)

		// 找左节点
		r, c := tv[0]+1, tv[1]
		key := ""
		if r < row && c < len(nums[r]) {
			key = strconv.Itoa(r) + "_" + strconv.Itoa(c)
			if ok := way[key]; !ok {
				way[key] = true
				result = append(result, nums[r][c])
				stack.PushBack([]int{r, c})
			}
		}

		// 找右节点
		r, c = tv[0], tv[1]+1
		if r < row && c < len(nums[r]) {
			key = strconv.Itoa(r) + "_" + strconv.Itoa(c)
			if ok := way[key]; !ok {
				way[key] = true
				result = append(result, nums[r][c])
				stack.PushBack([]int{r, c})
			}
		}
	}
	return result
}