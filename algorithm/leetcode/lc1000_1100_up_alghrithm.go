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

// letaeta
// N:1081 最小子序列
func smallestSubsequence(s string) string {
	m := make(map[byte]int,len(s))
	cs := []byte(s)
	for _,v := range cs {
		m[v]++
	}

	stack := []byte{}
	existKey := make(map[byte]bool,0)

	for _,v := range cs {
		if !existKey[v] {
			// 这里使用循环，是为了递归检查栈顶元素，将满足条件的依次出队
			for len(stack) > 0 && stack[len(stack)-1] > v  {
				topV := stack[len(stack)-1]
				// 为了保证每个元素最少存在依次，如果后边未检查的数据不包含该元素，那么该元素就不能出队
				if m[topV]== 0 {
					break
				}

				stack = stack[:len(stack)-1]
				existKey[topV] = false
			}
			stack = append(stack,v)
			existKey[v] = true
		}	
		m[v]--
	}

	return string(stack) 
}


