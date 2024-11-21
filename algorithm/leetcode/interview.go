package leetcode

import (
	"container/heap"
	"sort"
)

// 面试题 17.19. 消失的两个数字
// https://leetcode.cn/problems/missing-two-lcci/
func missingTwo(nums []int) []int {
	if len(nums) == 0 {
		return []int{1, 2}
	}

	return mt01(nums)

}

// mt01 等差数列
func mt01(nums []int) []int {
	n := len(nums) + 2
	// 1 求消失的2数之和
	ans := n * (n + 1) / 2
	for _, v := range nums {
		ans -= v
	}

	// 2. 消失的2数一定在min两侧
	mid := ans / 2
	cur := mid * (mid + 1) / 2
	for _, v := range nums {
		if v <= mid {
			cur -= v
		}
	}
	return []int{cur, ans - cur}
}

// 位运算 a^b^a = b
func mt02(nums []int) []int {
	ans, n := 0, len(nums)+2

	// 提取两个数的异或
	for i := 1; i <= n; i++ {
		ans ^= i
	}
	for _, v := range nums {
		ans ^= v
	}

	// 异或运算只有某一位不相同时，结果才为1
	// 该数的从右往左数第一个为 1 的位的权值，根据异或的特性，说明在该位，其中一个为1,一个位0
	c1, c2 := 0, 0
	lb := ans & -ans
	for i := 1; i <= n; i++ {
		if lb&i > 0 {
			c1 ^= i
		} else {
			c2 ^= i
		}
	}
	for _, v := range nums {
		if lb&v > 0 {
			c1 ^= v
		} else {
			c2 ^= v
		}
	}

	return []int{c1, c2}
}

// https://leetcode.cn/problems/get-kth-magic-number-lcci/
// 消失的两个数
func getKthMagicNumberR1719(k int) int {

	return minHeapR1719(k)
	// return dynamicway(k)
}

type MinHeapV1 struct{ sort.IntSlice }

// 继承heap.Interface接口
func (m *MinHeapV1) Push(a interface{}) {
	m.IntSlice = append(m.IntSlice, a.(int))
}

func (m *MinHeapV1) Pop() interface{} {
	a := m.IntSlice
	v := a[m.Len()-1]
	m.IntSlice = a[:m.Len()-1]
	return v
}

var factor []int = []int{3, 5, 7}

// 小顶堆
func minHeapR1719(k int) int {
	existKey := map[int]struct{}{1: {}}
	m, v := &MinHeapV1{}, 1
	heap.Push(m, 1)
	for {
		if k == 0 {
			break
		}
		v = heap.Pop(m).(int)
		// BUG 因为作用域的原因，在这里v在返回结果之前，o使用的外部声明的1
		// v := heap.Pop(m).(int)
		for _, f := range factor {
			if _, ok := existKey[f*v]; !ok {
				heap.Push(m, f*v)
				existKey[f*v] = struct{}{}
			}
		}
		k--
	}

	return v
}

//TODO 动态规划法，三指针 并行归并
