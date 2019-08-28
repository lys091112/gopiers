package leetcode

import "testing"

func TestGenerateParenthesis(t *testing.T) {

	res := generateParenthesis(4)

	for _, v := range res {
		t.Log("elem: " + v)
	}

}
