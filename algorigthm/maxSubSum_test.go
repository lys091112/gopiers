package leetcode

import (
	"testing"
)

func TestMaxSubSum(test *testing.T) {
	numbers := []int{1, 2, 4, -5, 3, 4, 2, -7, 6, 2}
	max := MaxSubSum(numbers[0:])
	if max != 12 {
		test.Errorf("result is %d,except is 12")
	}
	test.Log("result is %d", max)
}

func TestMaxSubSumNIL(test *testing.T) {
	max := MaxSubSum(nil)
	if max != 0 {
		test.Errorf("result is %d,except is 0")
	}
	test.Log("result is %d", max)
}

func TestMaxSubSumEmpty(test *testing.T) {
	numbers := []int{}
	max := MaxSubSum(numbers)
	if max != 0 {
		test.Errorf("result is %d,except is 0")
	}
	test.Log("result is %d", max)
}

func TestDivideSubSum(test *testing.T) {
	numbers := []int{1, 2, 4, -5, 3, 4, 2, -7, 6, 2}
	max := DivideAndConquerSum(numbers[0:], 0, len(numbers)-1)
	if max != 12 {
		test.Errorf("result is %d,except is 12", max)
	}
}
