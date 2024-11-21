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
**/

// 初始化
func InitDemo() {
	// nil 不能赋值给string
	// var str interface{} = nil
	// // if str == nil {
	// // 	fmt.Println("is nil")
	// // }

	var str2 = ""
	if len(str2) <= 0 {
		fmt.Println("is empty")
	}

	var str3 = "01234567"
	fmt.Println(str3[2] - '0')
	fmt.Println(str3 + "8")
}

// 遍历和截取
// 字符串存储的是类型为byte的字节切片，而byte实际为uint8类型的值
// TIP: 两种不同的方式遍历，得到的中间值不同
func LoopAndSubString() {
	var str = "Hello,世界"

	// 字符串截取 ello worl
	fmt.Println(str[1 : len(str)-1])

	// utf-8字符遍历
	for i := 0; i < len(str); i++ {
		fmt.Printf("%b,%s\n", str[i], string(str[i]))
	}

	// unicode字符遍历 v为 rune类型
	for i, v := range str {
		fmt.Printf("%b,%s\n", str[i], string(v))
	}

	str2 := "截取数据"
	fmt.Println(str2[:3]) // 截
	r := []rune(str2)
	fmt.Println(r[:3]) //截取数
}
