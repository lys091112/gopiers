package leetcode

import (
	"math"
)

/**
 *  按照中位数的定义，问题求解变成:
	1.len(leftPart)=len(rightPart)
	2.max(leftPart) ≤ min(rightPart)
	3.median = [max(leftPart) + min(rightPart)] / 2
	为了简化分析，我们先不考虑 i=0，i=m，j=0,j=n 这样的临界条件。
	1. len(leftPart)=len(rightPart)条件转化为:
	i+j=m−i+n−j+1
	i+j=(m+n+1)/2
	j = (m+n+1)/2 - i
	看一下约束条件:
	但是j不能是负数
	(m+n+1)/2 > i
	i的取值范围是[0,m]
	(m+n+1)/2 > m
	n+1 > m
	所以,n ≥ m才能使用上面的关系.

	2.max(leftPart)≤min(rightPart)条件转化为：
	B[j−1]≤A[i] 且 A[i-1] ≤ B[j]

	3.条件三分一下奇数和偶数:
	a. 当 m + n为奇数时,中位数为：max(A[i−1],B[j−1])
	​​b. 当 m + n为偶数时,中位数为：max(A[i−1],B[j−1])+min(A[i],B[j]) / 2

	i=0,i=m,j=0,j=n，此时A[i−1],B[j−1],A[i],B[j] 可能不存在。
	我们需要做的是确保max(leftPart)≤min(rightPart)。 如果A[i−1],B[j−1],A[i],B[j] 中部分不存在，
	那么我们只需要检查这两个条件中的一个（或不需要检查）

TODO error 待纠正
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if isEmpty(nums1) && isEmpty(nums2) {
		return 0.0
	}
	if isEmpty(nums1) {
		return findMiddle(nums2)
	}
	if isEmpty(nums2) {
		return findMiddle(nums1)
	}

	minArray := nums1
	maxArray := nums2
	var minLen, maxLen = len(minArray), len(maxArray)
	if minLen > maxLen {
		minArray, maxArray = maxArray, minArray
		minLen, maxLen = maxLen, minLen
	}

	low := 0
	high := minLen

	maxLeft, minRight := 0.0, 0.0
	i, j := 0, 0

	for {
		if low >= high {
			break
		}

		i = (low + high) / 2
		j = (minLen+maxLen+1)/2 - i

		if i > 0 && j < maxLen && minArray[i-1] > maxArray[j] {
			high = i - 1
		} else if j > 0 && i < minLen && maxArray[j-1] > minArray[i] {
			low = i + 1
		} else {

			if i == 0 {
				maxLeft = float64(maxArray[j-1])
			} else if j == 0 {
				maxLeft = float64(minArray[i-1])
			} else {
				maxLeft = math.Max(float64(minArray[i-1]), float64(maxArray[j-1]))
			}
			if (minLen+maxLen)%2 == 1 {
				return (maxLeft + minRight) / 2.0
			}
			break
		}
	}

	if j == maxLen {
		minRight = float64(minArray[i])
	} else if i == minLen {
		minRight = float64(maxArray[j])
	} else {
		minRight = math.Max(float64(minArray[i]), float64(maxArray[j]))
	}

	return (maxLeft + minRight) / 2.0
}

func isEmpty(nums []int) bool {
	return nums == nil || len(nums) <= 0

}

func maxNum(nums1 []int, nums2 []int) int {
	len1 := len(nums1) - 1
	len2 := len(nums2) - 1
	return If(nums1[len1] > nums2[len2], nums1[len1], nums2[len2]).(int)
}

func findMiddle(nums []int) float64 {

	var middle = len(nums) / 2

	if len(nums)%2 == 0 {
		return float64((nums[middle] + nums[middle+1]) / 2)
	} else {
		return float64(nums[middle])
	}

}
