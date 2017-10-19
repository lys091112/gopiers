package main

import (
	"fmt"
	"os"

	"github.com/lys091112/gopiers/basic/channel"
	"github.com/lys091112/gopiers/support/database"
)

func main() {
	path, err := os.Getwd()
	if nil != err {
		panic(err)
	}
	fmt.Println(path)

	//moon.Start()

	//	insect.Start()

	channel.SelectDemo()

	event := database.New(11, "hello")
	fmt.Println(event)
}

/**
func path() {
	filePath := strings.Join([]string{path, "README.md"}, "/")
	fmt.Println(filePath)
	fileread.ReadUseBufio(filePath)
	network.GetUrl("https://baidu.com")

	//	network.StartServer()
	basic.TestSha()
}
*/
