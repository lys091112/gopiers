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
}
