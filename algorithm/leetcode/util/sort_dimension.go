package util

import "sort"

// 二维数组排序
type Int2DSlice [][]int

func (p Int2DSlice) Len() int           { return len(p) }
func (p Int2DSlice) Less(i, j int) bool { return p[i][0] < p[j][0] }
func (p Int2DSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort is a convenience method.
func (p Int2DSlice) Int2DSort() { sort.Sort(p) }
