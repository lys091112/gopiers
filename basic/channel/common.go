package channel

import (
	"fmt"
	"math/rand"
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

func SelectDemo() {
	select_ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case select_ch <- i:
			// do nothind
		case x := <-select_ch:
			fmt.Println(x)
		}
	}
}

/*
func Select_Demo() {
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		select {
		case ch <- i:
			// do nothind
		case x := <-ch:
			fmt.Println(x)
		}
	}
}
*/

func DeadLockPrac() {
	dead_ch := make(chan string)
	go readFromDeadLockChan(dead_ch)
	dead_ch <- "hello"

	// 不在这里进行chan消费，是因为会造成死锁，chan是一个阻塞chan，
	// 必须首先有其他的其他的线程进行消费, 类比打电话，需要有人事先准备好接听
	// go readFromDeadLockChan(dead_ch)
}

func readFromDeadLockChan(ch chan string) {
	fmt.Println(<-ch)
}

/*
同比与上方法，<- dead_ch 就会造成死锁，因为没有取出任何东西，所以会一直等待
func DeadLockPrac2() {
	dead_ch := make(chan string)
	<-dead_ch
	dead_ch <- "hello"
}
*/

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 100; i++ {
			c <- fmt.Sprintf("%s: %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func fatIn(input1, intpu2 <-chan string, quit <-chan bool) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-intpu2:
				c <- s
			case <-time.After(1 * time.Second):
				fmt.Println("sleep one second")
			case <- quit:
				return // quit
			}
		}
	}()
	return c
}

// daisu chain
// func f(left, right chan int) {
//     left <- 1 + <-right
// }

// func main() {
//     const n = 10000
//     leftmost := make(chan int)
//     right := leftmost
//     left := leftmost
//     for i := 0; i < n; i++ {
//         right = make(chan int)
//         go f(left, right)
//         left = right
//     }
//     go func(c chan int) { c <- 1 }(right)
//     fmt.Println(<-leftmost)
// }
