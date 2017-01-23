package main

import (
	"fmt"
	"github.com/lys091112/gopiers/algorithm"
	"github.com/lys091112/gopiers/basic/fileread"
	"github.com/lys091112/gopiers/container"
	"os"
	"strings"
)

func main() {
	stack := container.NewStack()
	stack.Push(1)
	fmt.Println(stack.Len())

	var values = [...]int{1, 2, 3, 4, -5, 6}
	max1 := algorithm.MaxSubSum(values[0:])
	fmt.Println("the max is %d", max1)
	max2 := algorithm.MaxSubSum(values[3:])
	fmt.Println("the max is %d", max2)
	max3 := algorithm.MaxSubSum(values[0:5])
	fmt.Println("the max is %d", max3)

	//minStack := leetcode.NewMinStack()
	//iminStack.Max()
	path, err := os.Getwd()
	if nil != err {
		panic(err)
	}
	filePath := strings.Join([]string{path, "README.md"}, "/")
	fmt.Println(filePath)
	fileread.ReadUseBufio(filePath)
}
