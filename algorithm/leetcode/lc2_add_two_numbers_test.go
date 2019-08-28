package leetcode

import (
	"testing"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func TestAddTwoNumbers(t *testing.T) {
	r1 := buildListNode([]int{5})
	r2 := buildListNode([]int{5})

	res := addTwoNumbers(r1, r2)
	printListNode(res,t)

}

