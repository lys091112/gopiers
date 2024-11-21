package base

import "github.com/lys091112/gopiers/algorithm/leetcode/util"

// MonostoneStack 获取下一个比他小的元素的数组下标
func MonostoneStack(q []int) []int {
	arr := make([]int, len(q))
	util.ArrayFills(arr, -1)
	stack := []int{}
	for i := 0; i < len(q); i++ {
		for len(stack) > 0 && q[stack[len(stack)-1]] > q[i] {
			peek := stack[len(stack)-1]
			arr[peek] = i
			stack = stack[0 : len(stack)-1]
		}
		stack = append(stack, i)
	}

	return arr
}
