package leetcode

import (
	"testing"
)

func TestReverse(t *testing.T) {

	resultOfInt(321, reverse2(123), t)
	resultOfInt(-321, reverse2(-123), t)
	resultOfInt(21, reverse2(12000), t)
	resultOfInt(0, reverse2(1 << 31 - 1), t)
}

