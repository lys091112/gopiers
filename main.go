package main

import (
	"fmt"
	"github.com/lys091112/gopiers/container"
)

func main() {
	stack := container.NewStack()
	stack.Push(1)
	fmt.Println(stack.Len())
}
