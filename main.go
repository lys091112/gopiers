package main

import (
	"fmt"
	"github.com/lys091112/gopiers/container"
	"github.com/lys091112/gopiers/leetcode"
)

func main() {
	stack := container.NewStack()
	stack.Push(1)
	fmt.Println(stack.Len())

	minstack := leetcode.NewMinStack()
	minstack.Push(1)

}
