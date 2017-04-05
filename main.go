package main

import (
	"fmt"
	"os"

	"github.com/lys091112/gopiers/insect"
)

func main() {
	path, err := os.Getwd()
	if nil != err {
		panic(err)
	}
	fmt.Print(path)
	//moon.Start()

	insect.Start()
}

//func path() {
//	filePath := strings.Join([]string{path, "README.md"}, "/")
//	fmt.Println(filePath)
//	fileread.ReadUseBufio(filePath)
//	network.GetUrl("https://baidu.com")
//
//	//	network.StartServer()
//	basic.TestSha()
//}
