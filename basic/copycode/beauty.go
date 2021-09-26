package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("vim-go")
}

func Append(slice, data []byte) []byte {
	l := len(slice)
	if l+len(data) > cap(slice) { // reallocate
		// Allocate double what's needed, for future growth.
		newSlice := make([]byte, (l+len(data))*2)
		// The copy function is predeclared and works for any slice type.
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : l+len(data)]
	for i, c := range data {
		slice[l+i] = c
	}
	return slice
}

var sem = make(chan int, 10)
var maxOutStanding = 10

func handle(r *http.Request) {
	<-sem
	// Wait for active queue to drain.
	process(r)
	// May take a long time.
	sem <- 1
	// Done; enable next request to run.
}
func init() {
	for i := 0; i < maxOutStanding; i++ {
		sem <- 1
	}
}
func Serve(queue chan *http.Request) {
	for {
		req := <-queue
		go handle(req) // Don't wait for handle to finish.
	}
}

// 使用闭包,让req对每个go程都是唯一的
func Serve2(queue chan *http.Request) {
	for req := range queue {
		<-sem
		go func(req *http.Request) {
			process(req)
			sem <- 1
		}(req)
	}
}

//同方法3,建立一个变量
func Serve3(queue chan *http.Request) {
	for req := range queue {
		<-sem
		req := req
		go func(req *http.Request) {
			process(req)
			sem <- 1
		}(req)
	}
}

//另一个途径是 启动 固定数量的 handle
//Goroutine , 每个 Goroutine 都直接从 channel 中 读 取 请 求 。 这 个固定的数 值 就是同 时执 行 process 的最大并 发 数
func handle4(queue chan *http.Request) {
	for r := range queue {
		process(r)
	}
}

// Serve4 xx
func Serve4(clientRequests chan *http.Request, quit chan bool) {
	// Start handlers
	for i := 0; i < maxOutStanding; i++ {
		go handle(<-clientRequests)
	}
	<-quit // Wait to be told to exit.
}

func process(r *http.Request) {
	fmt.Println(r.Method)
}
