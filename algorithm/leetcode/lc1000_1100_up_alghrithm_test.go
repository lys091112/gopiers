package leetcode

import "testing"

func TestBitwiseComplement(t *testing.T) {
	res := bitwiseComplement(5)
	t.Log(res)
	t.Log(convertToBinaryStr(5))
}

func TestLastStoneWeight(t *testing.T) {
	var keys = []int{2, 7, 4, 1, 8, 1}
	t.Log(lastStoneWeight(keys))
}
