package basic

import (
	"fmt"
)

/**
*每个go程序最先执行的是init方法,其次才是main
	Println的签名: func Println(a ...interface{})(n int, err error)
	go语言不能在函数内声明另一个函数,但因为有函数变量,所以
	可以在一个函数中声明一个函数类型的变量,此时的函数成为闭包
	而go编译器汇报闭包使用的外围变量分配到堆上,堆不会随着函数的返回而
	自动消失,它在用完后才会被垃圾回收,因此,上层框架消失后,外围变量
	不会跟着消失(云集算--98)
**/

func f() (result int) {
	defer func() {
		result++
	}()
	return 1
}
func g() {
	a := f()
	fmt.Println(a) //result a = 2
}
