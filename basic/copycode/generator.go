package main

import (
	"fmt"
)

func xrange() chan int {
	var ch chan int = make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()

	return ch
}

func get_notification(user string) chan string {
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

	ch := make(chan int)

	go func() {
		for {
			select {
			case v1 := <-c1:
				ch <- v1
			case v2 := <-c2:
				ch <- v2
			case v3 := <-c3:
				ch <- v3
			}
		}
	}()

	for i := 0; i < 3; i++ {
		fmt.Println(<-ch)
	}
}

func main() {
	generator := xrange()

	for i := 0; i < 10; i++ {
		fmt.Println(<-generator)
	}

	joy := get_notification("joy")
	daisy := get_notification("daisy")

	fmt.Println(<-daisy)
	fmt.Println(<-joy)

	selectfoo()
}
