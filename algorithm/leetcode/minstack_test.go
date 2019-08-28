package leetcode

import (
	"testing"
)

func TestMax(t *testing.T) {
	minstack := NewMinStack()
	minstack.Push(3)
	minstack.Push(4)
	minstack.Push(2)
	minstack.Push(4)
	minstack.Push(1)
	minstack.Push(6)
	minstack.Push(8)
	minstack.Pop()
	minstack.Pop()
	minstack.Pop()
	minstack.Pop()
	t.Log("hello")
}
