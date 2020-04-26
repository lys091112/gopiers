package fileread

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadUseBufio(path string) {
	f, err := os.Open(path)
	defer f.Close()
	if nil == err {
		buff := bufio.NewReader(f)
		for {
			line, err := buff.ReadString('\n')
			if err != nil || io.EOF == err {
				break
			}
			fmt.Println(line)
		}
	}
}

func BaseUse() {
	path, _ := os.Getwd()
	fmt.Println(path)

}
