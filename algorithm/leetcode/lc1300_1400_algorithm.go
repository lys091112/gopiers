package leetcode

import "sort"

// arrayRankTransform N:1331 数组序号转换 离散化模板（排序+hash）
func arrayRankTransform(arr []int) []int {

	if len(arr) <= 0 {
		return []int{}
	}

	// SORT会改变切片，因此需要对数组做备份
	tmp := append([]int{}, arr...)

	sort.Ints(tmp)
	m := make(map[int]int, len(tmp))
	for _, v := range tmp {
		if _, ok := m[v]; !ok {
			// 已经存在的就u不需要更新下标,如果有set结构，会更方便一点
			m[v] = len(m) + 1
		}
	}

	res := make([]int, len(tmp))
	for i, v := range arr {
		res[i] = m[v]
	}

	return res
}
