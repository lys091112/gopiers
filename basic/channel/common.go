package channel

import (
	"fmt"
	"time"
)

/**
*一旦main函数结束,那么go函数也会结束
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	go sheep(1)

	time.Sleep(time.Millisecond)
	fmt.Println("end", a)
}

func sheep(i int) {
	for ; ; i += 1 {
		fmt.Println(i, " ka cha")
		a += 1
	}
}
**/
var ch = make(chan int, 10) // channal 初始化
var a string

func f() {
	fmt.Println("i am f")
	a = "hello, world"
	ch <- 0
}

func g() {
	fmt.Println("i am g")
	go f()
	<-ch
	fmt.Println("i am got a value", a)
}

func ChannelSelect() string {
	var timeout = make(chan bool, 1)
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("sleep")
		timeout <- true
	}()

	select {
	case <-ch:
		return "read"
	case <-timeout:
		return "timeout"
	}
}
