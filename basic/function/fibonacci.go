package function

import "fmt"

// 斐波纳契 数列
// 函数闭包 临时变量不用声明，在闭包函数中循环使用
// 可以用于数字发射器
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		fmt.Printf("a=%d,b=%d", a, b)
		fmt.Println()
		return a
	}
}
