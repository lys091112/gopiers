package mystrings

import "fmt"

/**
  * --- rune,string,type ---
  *1、type 占一个字节
  *2、rune 英文占一个字节，中文占三个字节
  *3、string底层是用byte数组存的，并且是不可以改变的。 
  *4、在 Go 中，字符串是以 UTF-8 为格式进行存储的，在字符串上调用 len 函数，取得的是字符串包含的 byte 的个数
  *例如 s:="你好"  fmt.Println(len(s))  输出结果应该是6，因为中文字符是用3个字节存的。
  *所以用string存储unicode的话，如果有中文，按下标是访问不到，如果想要获得我们想要的情况的话，需要先转换为rune切片再使用内置的len函数

 */

func demo() {
	// nil 不能赋值给string
	var str interface{} = nil
	if str == nil {
		fmt.Println("is nil")
	}

	var str2 = ""
	if len(str2) <= 0 {
		fmt.Println("is empty")
	}

	var str3 = "01234567"
	fmt.Println(str3[2] - '0')
	fmt.Println(str3 + "8")
}

func t1() {
	var str = "Hello world"

	// 字符串截取 ello worl
	fmt.Println(str[1 : len(str)-1])

	for _, v := range str {
		fmt.Println(v)
	}

	r := []rune(str)

	for _, v := range r {
		fmt.Println(v)
	}
}
