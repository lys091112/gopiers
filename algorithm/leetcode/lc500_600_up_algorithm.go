package leetcode

/**
 * N:540
 * 还有一种o(n) 的位运算方法：(可以用来应对无序数据）
 * for(int i : nums) num ^= i;
 *
 * 本题的一个限制条件是 输入是有序的数组
 */
func singleNonDuplicate(nums []int) int {

	if len(nums) <= 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	left := 0
	right := len(nums) - 1

	for {
		if left >= right {
			break
		}
		mid := (left + right) / 2

		if isOdd(mid) {
			if nums[mid] == nums[mid+1] {
				right = mid - 1
			} else if nums[mid] == nums[mid-1] {
				left = mid + 1
			} else {
				left, right = mid, mid
			}

		} else {
			if nums[mid] == nums[mid+1] {
				left = mid + 2
			} else if nums[mid] == nums[mid-1] {
				right = mid - 2
			} else {
				left, right = mid, mid
			}
		}

	}

	return nums[right]

}

func isOdd(n int) bool {
	return n%2 == 1
}
