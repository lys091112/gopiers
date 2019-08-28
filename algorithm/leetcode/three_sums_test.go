package leetcode

import "testing"

func TestThreeSums(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}

	res := threeSum(nums)
	t.Log(res)
}

func TestThreeSums2(t *testing.T) {
	nums := []int{1,2,-2,-1}

	res := threeSum(nums)
	t.Log(res)
}

func TestThreeSums3(t *testing.T) {
	nums := []int{-1,-2,-3,4,1,3,0,3,-2,1,-2,2,-1,1,-5,4,-3}

	res := threeSum(nums)
	t.Log(res)
}

