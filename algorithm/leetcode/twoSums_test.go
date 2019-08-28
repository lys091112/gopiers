package leetcode

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 11, 13, 18,26}
	target := 28
	res := TwoSum(nums, target)
	if len(res) == 2 && nums[res[0]] + nums[res[1]] == target {
		t.Log(res)
		t.Log("success")
	}else {
		t.Error("unexpected resultOfInt")
	}
}

func TestTwoSum2(t *testing.T) {
	nums := []int{3,2,4}
	target := 6
	res := TwoSum(nums, target)
	if len(res) == 2 && nums[res[0]] + nums[res[1]] == target {
		t.Log(res)
		t.Log("success")
	}else {
		t.Error("unexcpted resultOfInt")
	}
}

func TestNull(t *testing.T) {
	target := 9
	res := TwoSum([]int{},target)

	if len(res) <= 0 {
		t.Log("success")
	}else {
		t.Error("unexpected resultOfInt")
	}
}

func TestTwoSumNotExist(t *testing.T) {
	nums := []int{2, 6, 11, 13}
	target := 18
	res := TwoSum(nums, target)
	if len(res) <= 0 {
		t.Log("success")
	}else {
		t.Error("unexpected resultOfInt")
	}
}


