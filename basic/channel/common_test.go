package channel

import (
	"fmt"
	"testing"
	"time"
)

func Test_g(test *testing.T) {
	g()
	test.Log("common test success")
}

func TestChannelSelect(test *testing.T) {
	result := ChannelSelect()
	if result != "timeout" {
		test.Errorf("result is %s,except is 'timeout'", result)
	}
}

// TODO 奇怪的执行结果
func TestO(test *testing.T) {
	ch := make(chan interface{})
	for i := 0; i < 2; i++ {
		go func(i int) {
			select {
			case ch <- time.After(time.Second * time.Duration(4+i)):
				time.Sleep(time.Second)
				fmt.Printf("xxxxx%d\n", i)
			default:
				fmt.Printf("default=%d\n", i)
			}
		}(i)
	}
	fmt.Println("come")
	v := <-ch
	fmt.Println("", <-v.(<-chan time.Time))
}

// 即便 select 退出以后，没有执行完的协程依然在背后执行
// 如果退出背后的协程{声明周期内}
func TestO1(test *testing.T) {
	ch := make(chan string)
	go func() {
		fmt.Println("come in")
		time.Sleep(4 * time.Second)
		fmt.Println("come out")
		ch <- "hello"
	}()
	r := make(chan string)
	// select结束也无法中断正在执行的线程
	go func(i int) {
		select {
		case <-ch:
			fmt.Printf("xxxxx%d\n", i)
			r <- "world"
		default:
			fmt.Printf("default\n")
		}
	}(4)
	fmt.Println("come")
	// v := <-r
	// fmt.Println(v)
	time.Sleep(3 * time.Second)
	fmt.Println("end")
	close(ch)
	time.Sleep(2 * time.Second)
}

func TestO2(test *testing.T) {
	// context.WithTimeout()
	ch := make(chan int, 1)
	ch <- 1
	select {
	case ch <- 2:
	default:
		fmt.Println("channel is full !")
	}
	fmt.Println("finish")
}
