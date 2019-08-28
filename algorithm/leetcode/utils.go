package leetcode

import (
	"testing"
)

func resultOfInt(except int, actual int, t *testing.T) {

	if except == actual {
		t.Log("success")
	} else {
		t.Errorf("failed! except=%d, actual=%d", except, actual)
	}

}

func resultOfString(except string, actual string, t *testing.T) {

	if except == actual {
		t.Log("success")
	} else {
		t.Errorf("failed! except=%s, actual=%s", except, actual)
	}

}

func resultOfBool(except bool, actual bool, t *testing.T) {

	if except == actual {
		t.Log("success")
	} else {
		t.Errorf("failed! except=%t, actual=%t", except, actual)
	}

}

func If(b bool, v1 interface{}, v2 interface{}) interface{} {
	if b {
		return v1
	} else {
		return v2
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func printListNode(l *ListNode, t *testing.T) {
	for {
		if l != nil {
			t.Log(l.Val, "-->")
			l = l.Next
		} else {
			break
		}
	}
	t.Log("-------------------")
}

func buildListNode(nums []int) *ListNode {
	var result *ListNode
	var end *ListNode

	for _, v := range nums {
		if result == nil {
			result = &ListNode{Val: v, Next: nil}
			end = result
		} else {
			newNode := &ListNode{Val: v, Next: nil}
			end.Next = newNode
			end = newNode
		}
	}

	return result
}

