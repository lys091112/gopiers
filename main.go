package main

import (
	"fmt"
	"os"

	"github.com/lys091112/gopiers/basic/channel"
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
