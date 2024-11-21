package weekly

import (
	"fmt"
	"sort"
)

// 5653. 可以形成最大正方形的矩形数目
func countGoodRectangles(rectangles [][]int) int {
	if len(rectangles) <= 0 {
		return 0
	}
	count := make(map[int]int)
	maxLen := -1

	for _, v := range rectangles {
		edge := v[1]
		if v[0] < v[1] {
			edge = v[0]
		}
		count[edge]++
		if maxLen < edge {
			maxLen = edge
		}
	}

	return count[maxLen]

}

// 5243. 同积元组
// 无非是a*b == c * d,记录下两两乘积，然后计算C(n,2) 两两组合
func tupleSameProduct(nums []int) int {
	record := make(map[int]int)

	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			record[v*nums[j]]++
		}
	}

	ret := 0
	for _, v := range record {
		ret += v * (v - 1) / 2
	}
	return ret * 8
}

// 全局遍历，超时
func TupleSameProductV2(nums []int) int {
	if len(nums) < 4 {
		return 0
	}

	count := 0

	road := make(map[string]bool)
	record := make(map[int]int)

	sort.Sort(sort.IntSlice(nums))

	for i, v := range nums {
		record[v] = i
	}

	for i := range nums {
		for j := i + 1; j < len(nums); j++ {

			n := nums[i] * nums[j]
			for k := i + 1; k < j; k++ {
				// 已经遍历过
				if n%nums[k] != 0 || road[uniqueKey(nums[i], nums[k], 0, nums[j])] {
					continue
				}
				if idx, ok := record[n/nums[k]]; ok && idx > k {
					count += 8
					road[uniqueKey(nums[i], nums[idx], 0, nums[j])] = true
				}

			}
		}
	}

	return count
}

func uniqueKey(i, j, k, z int) string {
	return fmt.Sprintf("%d_%d_%d_%d\n", i, j, k, z)
}

// N 5655. 重新排列后的最大子矩阵
//
//	类似于85，对数据预处理，算以这个点为结尾，上面有多少个连续的1，就是这一列以这个点为结尾的最大高度
//
// 这样就将二维问题转成一维, 即求圆柱可以组成的最大矩形面积
func largestSubmatrix(matrix [][]int) int {

	if len(matrix) == 0 {
		return 0
	}
	for i, v := range matrix {
		if i == 0 {
			continue
		}
		for j, v1 := range v {
			if v1 == 1 {
				matrix[i][j] = 1 + matrix[i-1][j]
			}
		}
	}

	n := len(matrix[0])
	ret, area, height := 0, 0, 0
	for _, v := range matrix {
		sort.Sort(sort.IntSlice(v))
		for j := 0; j < n; j++ {
			height = v[j]
			if 0 == height {
				continue
			}
			area = height * (n - j)
			if ret < area {
				ret = area
			}
		}
	}

	return ret
}

// N: 1728 1723
