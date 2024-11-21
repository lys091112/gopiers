package leetcode

import (
	"container/heap"
	"math"
	"strings"

	"github.com/lys091112/gopiers/algorithm/base/queue"
	"github.com/lys091112/gopiers/algorithm/leetcode/util"
)

/**
 * N: 1002
 */
func bitwiseComplement(N int) int {
	// 这里的边界考虑
	if N <= 0 {
		return 1
	}

	b := convertToBinaryStr(N)

	res := 0
	// 找到第一个停止的位置
	for i, v := range b {
		if v == '0' {
			res = res + int(math.Pow(float64(2), float64(len(b)-1-i)))
		}

	}
	return res
}

/*
*

  - N: 1046 最后一块石头的重量

  - 考虑策略 为优先队列，

  - 之所以使用su的指针，是为了保证队列修改了su的内容后，可以在外部也体现出来
*/
func lastStoneWeight(stones []int) int {
	var su queue.PriorityQueue = stones[:]
	heap.Init(&su)

	for {
		if su.Len() <= 1 {
			break
		}
		x := heap.Pop(&su).(int)
		y := heap.Pop(&su).(int)

		if x == y {
			continue
		}

		if x < y {
			heap.Push(&su, y-x)
		} else {
			heap.Push(&su, x-y)
		}
	}

	if su.Len() == 1 {
		return su[0]
	}
	return 0
}

// letaeta
// N:1081 最小子序列
func smallestSubsequence(s string) string {
	m := make(map[byte]int, len(s))
	cs := []byte(s)
	for _, v := range cs {
		m[v]++
	}

	stack := []byte{}
	existKey := make(map[byte]bool, 0)

	for _, v := range cs {
		if !existKey[v] {
			// 这里使用循环，是为了递归检查栈顶元素，将满足条件的依次出队
			for len(stack) > 0 && stack[len(stack)-1] > v {
				topV := stack[len(stack)-1]
				// 为了保证每个元素最少存在依次，如果后边未检查的数据不包含该元素，那么该元素就不能出队
				if m[topV] == 0 {
					break
				}

				stack = stack[:len(stack)-1]
				existKey[topV] = false
			}
			stack = append(stack, v)
			existKey[v] = true
		}
		m[v]--
	}

	return string(stack)
}

// shortestCommonSupersequence
// N:1092 最短公共超序列
// https://blog.csdn.net/c2434525400/article/details/123573955
// LCS
// if a[i] == b[j]  f[i][j] = f[i-1][j-1] + 1
// if a[i] != b[j]  f[i][j] = max(f[i-1][j],f[i][j-1])
func shortestCommonSupersequence(str1 string, str2 string) string {
	n, m := len(str1), len(str2)
	f := util.InitArray(n+1, m+1, 0)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if str1[i-1] == str2[j-1] {
				f[i][j] = f[i-1][j-1] + 1
			} else {
				f[i][j] = util.Max(f[i-1][j], f[i][j-1])
			}
		}
	}

	var res strings.Builder
	i, j := n, m
	for {
		if i == 0 && j == 0 {
			break
		}

		if i == 0 {
			res.WriteByte(str2[j-1])
			j--
		} else if j == 0 {
			res.WriteByte(str1[i-1])
			i--
		} else {
			if str1[i-1] == str2[j-1] {
				res.WriteByte(str1[i-1])
				i--
				j--
			} else if f[i][j] == f[i-1][j] {
				res.WriteByte(str1[i-1])
				i--
			} else if f[i][j] == f[i][j-1] {
				res.WriteByte(str2[j-1])
				j--

			}
		}
	}

	return reverseString(res.String())
}

// N:1094
// 差分数组
func CarPolling(trips [][]int, capacity int) bool {

	if len(trips) == 0 {
		return true
	}

	// 因为to <= 1000,初始化差分数组
	var nums [1010]int

	for _, v := range trips {
		p, from, to := v[0], v[1], v[2]

		//更新[from to)之间的差分数组,左开右闭，更新nums[from]、nums[to],为从1开始，可以统一加1,nums[0]作为空哨兵
		// 因为原生数组的获取会利用到规则： c[i] = c[i-1] + nums[i], c 为期望的源数组
		nums[from+1] += p
		nums[to+1] -= p
	}

	for i := 1; i < 1000; i++ {
		// 差分数组的利用
		nums[i] += nums[i-1]
		if nums[i] > capacity {
			return false
		}
	}

	return true
}
