package channel

import (
	"fmt"
	"math/rand"
	"time"
)

// 关闭原则
// 1. 永远不要尝试在读取端关闭 channel ，写入端无法知道 channel 是否已经关闭，往已关闭的 channel 写数据会 panic ；
// 2. 一个写入端，在这个写入端可以放心关闭 channel；
// 3. 多个写入端时，不要在写入端关闭 channel ，其他写入端无法知道 channel 是否已经关闭，关闭已经关闭的 channel 会发生 panic （你要想个办法保证只有一个人调用 close）；
// 4. channel 作为函数参数的时候，最好带方向；

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

// 	下面方法等价于
// while(true) {
// 	if (hasValue(input1)) {
// 		do something	
// 	}else if (hasValue(input2)) {
// 		do something	
// 	}else if(hasAfterOneMinute) {
// 		do something	
// 	}else if hasValue(quit) {
// 		return 
// 	}
// }
func fatIn(input1, intpu2 <-chan string, quit <-chan bool) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
				// 被翻译为 runtime.go ==> selectnbrecv1(...)
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

/**
	Flag: close 
	Return: MainGoRuntine
	Reason: 
 **/
func colse_demo() {
	ch := make(chan bool)

	go func(){
		// v:false  ok : false 
		v, ok := <- ch
		fmt.Print("GoRuntine ", v, ok)

		// 从关闭的channel可以一直获取到值（类型的默认值）
		// for {
		// 	v := <- ch
		// 	fmt.Println("tt --> " , v)
		// }
	}()
	time.After(time.Second * 2)
	close(ch)
	time.After(time.Second * 1)
	fmt.Println("Main")
}

// 非阻塞示例，需要有default来退出
func select_cycle() {
	ch := make(chan int)
	timeout := make(chan bool)
	finish := false
	go func() {
		time.Sleep(2*time.Second)
		fmt.Println("---------- 2s ")
		timeout <- true
		finish = true
	}()

	go func() {
		for!finish{
			select {
			case <-ch:
				fmt.Println("receive")
			case <-timeout:
				fmt.Println("timeout")
			default:
				// 非阻塞类型的
			}
		}
	}()
	go func() {
		time.Sleep(4 * time.Second)
		fmt.Println("---------- 4s ")
		ch <- 1
	}()
	time.Sleep(5 * time.Second)
}

// select_sync 阻塞等待超时
func select_sync() {
	ch := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- 1
	}()
	select {
	case <- ch :
		fmt.Print("receive01")
	case <- time.After(1 * time.Second):
		fmt.Println("timeout")
	}

	c2 := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- 1
	}()

	select {
	case <- c2 :
		fmt.Println("receive02")
	case <- time.After(3 * time.Second):
		fmt.Println("timeout02")

	}

}
