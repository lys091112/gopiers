package leetcode

import "github.com/lys091112/gopiers/basic/container"

type MinStack struct {

	// 这个类应该怎样定义和使用
	minst *container.Stack
	maxst *container.Stack
	data  *container.Stack
}

func NewMinStack() *MinStack {
	return &MinStack{container.NewStack(), container.NewStack(), container.NewStack()}
}

/*
func NewMinStack() *MinStack {
	data := container.NewStack()
	maxst = container.NewStack()
	minst = container.NewStack()
	return &MinStack{}
}*/

func (minstack *MinStack) Push(value int) {
	minstack.data.Push(value)
	if minstack.minst.Peek() == nil || minstack.minst.Peek().(int) >= value {
		minstack.minst.Push(value)
	}
	if minstack.maxst.Peek() == nil || minstack.maxst.Peek().(int) <= value {
		minstack.maxst.Push(value)
	}
}

func (minstack *MinStack) Pop() interface{} {
	pop := minstack.data.Pop()
	if pop == nil {
		return nil
	}
	if pop.(int) == minstack.minst.Peek().(int) {
		return minstack.minst.Pop()
	}
	if pop.(int) == minstack.maxst.Peek().(int) {
		return minstack.maxst.Pop()
	}
	return nil
}

func (minstack *MinStack) Max() interface{} {
	return minstack.maxst.Peek()
}
func (minstack *MinStack) Min() interface{} {
	return minstack.minst.Peek()
}

/*
func main() {
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

} */
