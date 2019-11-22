package basic

import (
	"fmt"
	"strconv"
)

func float_demo() {
	fmt.Println(strconv.Atoi("9"))
	fmt.Println(strconv.FormatInt(10, 2))
	var a = complex(3, 4)
	var b = real(a)
	var c = imag(a)
	fmt.Println(b)
	fmt.Println(c)

	// 通过var声明，默认会给予初始值
	var i int
	fmt.Printf(`i is %d`, i)
}

// 声明切片二维数组
func two_array(length int) [][]int {
	dep := make([][]int, length)
	for i := 0; i < length; i++ {
		dep[i] = make([]int, length)
	}

	return dep
}

type BaseHuman struct {
	Name string
	Sex  string
}

// fmt format
func fmt_formet_demo() {
	var human = BaseHuman{Name: "lily", Sex: "man"}
	fmt.Printf("name is %v\n", human)  // 相应值的默认格式
	fmt.Printf("name is %+v\n", human) // 打印结构体时，会添加字段名
	fmt.Printf("name is %#v\n", human) // 相应值的Go语法表示
	fmt.Printf("name is %T\n", human)  // 相应值的类型的Go语法表示

	fmt.Printf("bool format %%t is %t\n", true) // bool 值标识

	// %b	二进制表示	Printf(“%b”, 5)	101
	// %c	相应Unicode码点所表示的字符	Printf(“%c”, 0x4E2D)	中
	// %d	十进制表示	Printf(“%d”, 0x12)	18
	// %o	八进制表示	Printf(“%d”, 10)	12
	// %q	单引号围绕的字符字面值，由Go语法安全地转义	Printf(“%q”, 0x4E2D)	‘中’
	// %x	十六进制表示，字母形式为小写 a-f	Printf(“%x”, 13)	d
	// %X	十六进制表示，字母形式为大写 A-F	Printf(“%X”, 13)	D
	// %U	Unicode格式：U+1234，等同于 “U+%04X”	Printf(“%U”, 0x4E2D)	U+4E2D

	fmt.Printf("binary is %%b -->  %b\n", 5)
	fmt.Printf("char is %%c -->  %c\n", 0x4E2D)
	fmt.Printf("int is %%d -->  %d\n", 0x12)
	fmt.Printf("八进制 is %%o -->  %o\n", 10)
	fmt.Printf("单引号围绕 is %%q -->  %q\n", 0x4E2D)
	fmt.Printf("十六进制 is %%x -->  %x\n", 13)
	fmt.Printf("十六进制大写 is %%X -->  %X\n", 13)
	fmt.Printf("Unicode is %%U -->  %U\n", 0x4E2D)

	// 浮点数
	// %b	无小数部分的，指数为二的幂的科学计数法，与 strconv.FormatFloat 的 ‘b’ 转换格式一致。例如 -123456p-78
	// %e	科学计数法，例如 -1234.456e+78	Printf(“%e”, 10.2)	1.020000e+01
	// %E	科学计数法，例如 -1234.456E+78	Printf(“%e”, 10.2)	1.020000E+01
	// %f	有小数点而无指数，例如 123.456	Printf(“%f”, 10.2)	10.200000
	// %g	根据情况选择 %e 或 %f 以产生更紧凑的（无末尾的0）	输出 Printf(“%g”, 10.20)	10.2
	// %G	根据情况选择 %E 或 %f 以产生更紧凑的（无末尾的0）	输出 Printf(“%G”, 10.20+2i)	(10.2+2i)
	fmt.Printf("%%e,  %e\n", 10.2)
	fmt.Printf("%%E,  %E\n", 10.2)
	fmt.Printf("%%f,  %f\n", 10.2)
	fmt.Printf("%%G,  %G\n", 10.20+2i)
	fmt.Printf("%%g,  %g\n", 10.20)

	// 字符串与字节切片
	// %s	输出字符串表示（string类型或[]byte)	Printf(“%s”, []byte(“Go语言”))	Go语言
	// %q	双引号围绕的字符串，由Go语法安全地转义	Printf(“%q”, “Go语言”)	“Go语言”
	// %x	十六进制，小写字母，每字节两个字符	Printf(“%x”, “golang”)	676f6c616e67
	// %X	十六进制，大写字母，每字节两个字符	Printf(“%X”, “golang”)	676F6C616E67
	fmt.Printf("%%s --> %s\n", []byte("Go language"))
	fmt.Printf("%%q --> %q\n", []byte("Go language"))
	fmt.Printf("%%x --> %x\n", "golang")
	fmt.Printf("%%X --> %X\n", "golang")

	// 指针
	// %p	十六进制表示，前缀 0x	Printf(“%p”, &people)	0x4f57f0
	fmt.Printf("%%p --> %p\n", &human)

	// %w 在errors.Unwrap 时使用
}
