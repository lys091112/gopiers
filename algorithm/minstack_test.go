package algorithm

import (
	"fmt"
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
	fmt.Println(minstack.Max())
	fmt.Println(minstack.Min())
	minstack.Pop()
	minstack.Pop()
	minstack.Pop()
	minstack.Pop()
	fmt.Println(minstack.Max())
	fmt.Println(minstack.Min())
	t.Log("hello")
}
