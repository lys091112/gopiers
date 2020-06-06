package main

import (
	"fmt"
	"os"

	"github.com/lys091112/gopiers/moon"
)

func main() {
	fmt.Println("hello world")
	path, err := os.Getwd()
	if nil != err {
		panic(err)
	}
	fmt.Println(path)
	moon.Start()
}
