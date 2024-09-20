package fileread

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"syscall"
	// "golang.org/x/exp/mmap" // 仅实现了read操作，如果需要write，还是需要syscall
)

// 参考链接： https://www.jianshu.com/p/509bb77ec103

// ReadFile 一次性加载到内存中
func ReadFile(filePath string) []byte {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println("Read error")
	}
	return content
}

// ReadUseBufio read
// 通过流的方式读取文件（当文件较大时，可以不用将文件一次读入内存中，防止内存溢出）
func ReadUseBufio(filePath string, handle func(string)) error {
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)

	for {
		line, err := buf.ReadString('\n')
		if err == nil {
			line = strings.TrimSpace(line)
			handle(line)
		} else if err == bufio.ErrBufferFull {
			// 代表读取的文件为读完整行缓冲区就满了，需要特殊处理
			// TODO 需要记录中间结果
			return err
		} else if err == io.EOF {
			return nil
		}
	}
}

// ReadUseScanner read
// 通过流的方式读取文件（当文件较大时，可以不用将文件一次读入内存中，防止内存溢出）
func ReadUseScanner(filePath string, handle func(string)) error {
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return err
	}
	buffer := bufio.NewScanner(f)

	for buffer.Scan() {
		handle(buffer.Text())
	}
	if buffer.Err() != nil {
		return err
	}
	return nil
}

// ReadBigFile  读取二进制文件，没有换行符的时候
func ReadBigFile(fileName string, handle func([]byte)) error {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("can't opened this file")
		return err
	}
	defer f.Close()
	s := make([]byte, 4096)
	for {
		switch nr, err := f.Read(s[:]); true {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "cat: error reading: %s\n", err.Error())
			os.Exit(1)
		case nr == 0: // EOF
			return nil
		case nr > 0:
			handle(s[0:nr])
		}
	}
}

// FileName 文件名
const FileName = "/Users/langle/xx.txt"

// readMmap 通过使用mmap的方式操纵文件,修改文件头元素为hello world
// Mmap 参数
// - fd：待映射的文件描述符。
// - offset：映射到内存区域的起始位置，0 表示由内核指定内存地址。
// - length：要映射的内存区域的大小。
// - prot：内存保护标志位，可以通过或运算符`|`组合
//     - PROT_EXEC  // 页内容可以被执行
//     - PROT_READ  // 页内容可以被读取
//     - PROT_WRITE // 页可以被写入
//     - PROT_NONE  // 页不可访问
// - flags：映射对象的类型，常用的是以下两类
//     - MAP_SHARED  // 共享映射，写入数据会复制回文件, 与映射该文件的其他进程共享。
//     - MAP_PRIVATE // 建立一个写入时拷贝的私有映射，写入数据不影响原文件。
func readMmap() {
	pageSize := os.Getpagesize()

	f, _ := os.OpenFile(FileName, os.O_RDWR|os.O_CREATE, 0644)
	stat, _ := f.Stat()
	if stat.Size() == 0 {
		_, _ = f.WriteAt(bytes.Repeat([]byte{'0'}, pageSize), 0)
		stat, _ = f.Stat()
	}

	fmt.Printf("pagesize: %d\n filesize: %d\n", pageSize, stat.Size())

	data, _ := syscall.Mmap(int(f.Fd()), 0, int(stat.Size()), syscall.PROT_WRITE, syscall.MAP_SHARED)

	// 关闭文件，不影响mmap
	f.Close()

	for i, x := range strings.Split("hello world", "") {
		data[i] = []byte(x)[0] // string 转 byte ,用[]byte(*)[index]的方式
	}

	syscall.Munmap(data)
}

func BaseUse() {
	path, _ := os.Getwd()
	fmt.Println(path)

}
