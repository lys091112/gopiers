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
