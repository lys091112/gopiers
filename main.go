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

	var values = [...]int{1, 2, 3, 4, -5, 6}
	max1 := leetcode.MaxSubSum(values[0:])
	fmt.Println("the max is %d", max1)
	max2 := leetcode.MaxSubSum(values[3:])
	fmt.Println("the max is %d", max2)
	max3 := leetcode.MaxSubSum(values[0:5])
	fmt.Println("the max is %d", max3)

	//minStack := leetcode.NewMinStack()
	//iminStack.Max()
}
