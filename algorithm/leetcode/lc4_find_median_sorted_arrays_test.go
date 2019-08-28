package leetcode

import (
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {

	nums1 := []int{1}
	nums2 := []int{2, 3, 4}
	res := findMedianSortedArrays(nums1, nums2)

	if res == 2.5 {
		t.Log("success")
	} else {
		t.Errorf("failed! excepted=%f, actual=%f", 2.5, res)
	}

}

func TestFindMedianSortedArrays02(t *testing.T) {

	nums1 := []int{1, 2, 3, 4, 5, 6}
	nums2 := []int{2, 3, 4, 5}
	res := findMedianSortedArrays(nums1, nums2)

	if res == 3.5 {
		t.Log("success")
	} else {
		t.Errorf("failed! excepted=%f, actual=%f", 3.5, res)
	}

}
