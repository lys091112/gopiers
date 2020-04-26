package main

import (
	"fmt"
	"sync"
	"time"
)

func xrange() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()

	return ch
}

func getNotification(user string) chan string {
	notify := make(chan string)

	go func() {
		notify <- fmt.Sprintf("welcome %s, do your own logic", user)
	}()

	return notify
}

func foo(i int) chan int {
	ch := make(chan int)
	go func() {
		ch <- 10 * i
	}()
	return ch
}

func selectfoo() {
	c1, c2, c3 := foo(1), foo(2), foo(3)
	var wait = sync.WaitGroup{}
	wait.Add(1)

	// time 是一个计时信道，时间一到，会收到数据
	timeout := time.After(1 * time.Second)
	ch := make(chan int)

	go func() {
		for isTimeout := false; !isTimeout; {
			select {
			case v1 := <-c1:
				ch <- v1
			case v2 := <-c2:
				ch <- v2
			case v3 := <-c3:
				ch <- v3
			case <-timeout:
				fmt.Println("timeout")
				isTimeout = true
				wait.Done()
			}
		}
	}()

	for i := 0; i < 3; i++ {
		fmt.Println(<-ch)
	}
	wait.Wait()
}

//虽然c1先受到数据，是否存在timeout先被select的情况
//这个方法要稍后再次验证
func timeout1() {
	c1, timeout := make(chan string), make(chan int)

	go func() {
		c1 <- "joyce"
		timeout <- 1
	}()

	for quit := false; !quit; {
		select {
		case v1 := <-c1:
			fmt.Printf("timeout demo comming. %s \r\n", v1)
		case <-timeout:
			fmt.Println("timeout demo end.")
			quit = true
		}
	}
}

func rand01() {
	ch := make(chan int)

	go func() {
		for {
			select {
			case ch <- 0:
			case ch <- 1:
			}
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Print(<-ch)
	}
	fmt.Println("")
	fmt.Println("rand01 end")
}

func txst() {
	generator := xrange()

	for i := 0; i < 10; i++ {
		fmt.Println(<-generator)
	}

	joy := getNotification("joy")
	daisy := getNotification("daisy")

	fmt.Println(<-daisy)
	fmt.Println(<-joy)

	selectfoo()

	timeout1()

	rand01()
}
