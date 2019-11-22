package leetcode

import (
	"container/heap"
	"math"

	"github.com/lys091112/gopiers/algorithm/base/queue"
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

/**
 * N: 1046 最后一块石头的重量

 * 考虑策略 为优先队列，
 * 之所以使用su的指针，是为了保证队列修改了su的内容后，可以在外部也体现出来
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
