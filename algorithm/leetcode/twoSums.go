package leetcode

// 遍历数组，然后查看减去数组元素剩余的值是否在数组中
// 使用map查找
func TwoSum(nums []int, target int) []int {
	if nums == nil || len(nums) <= 1 {
		return []int{}
	}

	store := make(map[int]int)

	for middle, k := range nums {
		left := target - k
		if value, ok := store[left]; ok {
			return []int{value, middle}
		}else {
			store[k] = middle
		}
	}
	return []int{}
}

//func find01(s []int, k int) (int, bool) {
//	lo, hi := 0, len(s)-1
//	for lo <= hi {
//		m := (lo + hi) >> 1
//		if s[m] < k {
//			lo = m + 1
//		} else if s[m] > k {
//			hi = m - 1
//		} else {
//			return m, true
//		}
//	}
//	return -1, false
//}
