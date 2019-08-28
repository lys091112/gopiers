package leetcode

import "testing"

func TestRomanToInt(t *testing.T) {

	resultOfInt(3, romanToInt("III"), t)

	resultOfInt(1994, romanToInt("MCMXCIV"), t)
	resultOfInt(58, romanToInt("LVIII"), t)
}
