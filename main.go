package main

import (
	"fmt"
	"os"

	"github.com/lys091112/gopiers/common/websocket"
)

func main() {
	path, err := os.Getwd()
	if nil != err {
		panic(err)
	}
	fmt.Println(path)
	// loura.Start()
	websocket.Get()
}
