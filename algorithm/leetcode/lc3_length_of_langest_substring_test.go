package leetcode

import "testing"

func TestSameStringLengestSubString(t *testing.T) {
	s := "bbbbb"

	var res = lengthOfLongestSubstring(s)

	if res == 1 {
		t.Log("success")
	} else {
		t.Errorf("test failed! expected size=1, actual=%d", res)
	}
}

func TestLengestSubString02(t *testing.T) {
	s := "abcabcbb"

	var res = lengthOfLongestSubstring(s)

	if res == 3 {
		t.Log("success")
	} else {
		t.Errorf("test failed! expected size=3, actual=%d\n", res)
	}
}

func TestLengestSubString03(t *testing.T) {
	s := "abcdsscsaeiknksuinkmopqrsittlsieng"

	var res = lengthOfLongestSubstring(s)

	if res == 10 {
		t.Log("success")
	} else {
		t.Errorf("test failed! expected size=10, actual=%d\n", res)
	}
}
