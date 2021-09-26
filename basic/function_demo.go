package basic

import (
	"fmt"
	"runtime"
	"sync"
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

// 闭包的理解

// Increase n成为了闭包方法的共用变量
func Increase() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

// 现象的原因在于闭包共享外部的变量i，注意到，每次调用go就会启动一个goroutine，这需要一定时间；
// 但是，启动的goroutine与循环变量递增不是在同一个goroutine，可以把i认为处于主goroutine中。
// 启动一个goroutine的速度远小于循环执行的速度，所以即使是第一个goroutine刚起启动时，外层的循环也执行到了最后一步了。
// 由于所有的goroutine共享i，而且这个i会在最后一个使用它的goroutine结束后被销毁，所以最后的输出结果都是最后一步的i==5。
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
		// 等待与不等待的区别是很大的
		// time.Sleep(1 * time.Second) // 设置时间延时1秒
	}
	wg.Wait()
}
