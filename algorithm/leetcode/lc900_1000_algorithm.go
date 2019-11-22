package leetcode

import (
	"container/list"
	"sort"

	"github.com/lys091112/gopiers/algorithm/base"
	"github.com/lys091112/gopiers/algorithm/leetcode/util"
)

// N:946
// 判断三种情况
// 	1. 栈顶等于当前值，remove(stack.top) -> next one
//	2. 栈顶的索引大于当前值的索引，即该值已经被放入到栈内，无法被取出，直接返回
//	3. 遍历pushed栈，继续放入还未放入的元素
func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) <= 0 && len(popped) <= 0 {
		return true
	}
	m := make(map[int]int, len(pushed))
	for i, v := range pushed {
		m[v] = i
	}

	topIndex := 0
	stack := list.New()
	for i := 0; i < len(popped); i++ {
		top := stack.Front()
		v := popped[i]

		if top != nil && v == top.Value.(int) {
			stack.Remove(top)
		} else if top != nil && m[v] < m[top.Value.(int)] {
			return false
		} else {
			for {
				if topIndex < len(pushed) {
					if v != pushed[topIndex] {
						stack.PushFront(pushed[topIndex])
						topIndex++
					} else {
						topIndex++
						break
					}
				} else {
					break
				}
			}
		}
	}
	return true
}

// N:952 按公因数计算最大组件大小
// 对于较大数据集，如果两两比对公约数，可能会超时
// 考虑埃氏筛法，进行筛选
// TODO
func largestComponentSize(A []int) int {
	if len(A) <= 0 {
		return 0
	}

	find := base.UnionFind{}
	find.MakeSet(len(A))

	// 是否有优化的空间
	for i := 0; i < len(A)-1; i++ {
		for j := i + 1; j < len(A); j++ {
			// 如果这俩是公约数大于1的，则进行连接
			if util.Divisor(A[i], A[j]) > 1 {
				find.Union(i, j)
			}
		}
	}

	max := 0

	m := find.Collections()

	for _, v := range m {
		if len(v) > max {
			max = len(v)
		}
	}
	return max
}

// 954. 二倍数对数组
// 按照一定的顺序，当前值一定有对应的值匹配，否则返回失败
// 注意在等比缩小时，奇偶数的问题
func canReorderDoubled(A []int) bool {
	if len(A) <= 0 {
		return false
	}

	m := make(map[int]int, len(A))

	for _, v := range A {
		m[v] += 1
	}

	sort.Ints(A)
	for _, d := range A {
		if m[d] <= 0 {
			continue
		}

		var n int
		if d < 0 {
			if (-d)%2 != 0 {
				return false
			}
			n = d / 2
		} else {
			n = d * 2
		}

		if v, ok := m[n]; ok && v > 0 {
			m[d] -= 1
			m[n] -= 1
		} else {
			return false
		}
	}

	return true
}
