package leetcode

import (
	"testing"
)

func TestThreeSumClosest(t *testing.T) {

	resultOfInt(2, threeSumClosest([]int{-1, 2, 1, -4}, 1), t)
	resultOfInt(0, threeSumClosest([]int{1, 2, 0, -3}, 1), t)
}

func TestRemoveNthFromEnd(t *testing.T) {

	head := buildListNode([]int{1, 2, 3, 4, 5})
	head = removeNthFromEnd(head, 4)
	//printListNode(head, t)

	head = buildListNode([]int{1, 2, 3, 4, 5})
	head = removeNthFromEnd(head, 1)
	//printListNode(head, t)

	head = buildListNode([]int{1, 2, 3, 4, 5})
	head = removeNthFromEnd(head, 5)
	//printListNode(head, t)
}

func TestIsValid(t *testing.T) {
	resultOfBool(true, isValid("(())"), t)
	resultOfBool(true, isValid("([])"), t)
	resultOfBool(false, isValid("((})"), t)
	resultOfBool(false, isValid("({))"), t)
	resultOfBool(false, isValid("((())"), t)
	resultOfBool(false, isValid("(())}"), t)
	resultOfBool(true, isValid("{(([]))}"), t)
	resultOfBool(true, isValid("()[]{}"), t)
	resultOfBool(true, isValid(""), t)
	resultOfBool(false, isValid("(]"), t)
	resultOfBool(true, isValid("(([]){})"), t)
}
