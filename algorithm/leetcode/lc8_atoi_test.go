package leetcode

import "testing"

func TestMyAtoi(t *testing.T) {

	resultOfInt(123, myAtoi("123o34"), t)
	resultOfInt(123, myAtoi("  123o34"), t)
	resultOfInt(0, myAtoi("  o12334"), t)
	resultOfInt(-123, myAtoi("   -123xx34"), t)
	resultOfInt(-2147483648, myAtoi("-91283472332"), t)
	resultOfInt(1, myAtoi("+1"), t)
	resultOfInt(0, myAtoi("+-2"), t)
	resultOfInt(0, myAtoi("-+2"), t)
	resultOfInt(2147483647, myAtoi("20000000000000000000"), t)
	resultOfInt(2, myAtoi("00000000000000000002"), t)
	resultOfInt(0, myAtoi(" -0 123"), t)
	resultOfInt(4139, myAtoi("4139 test"), t)

}
