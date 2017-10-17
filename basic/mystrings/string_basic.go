package mystrings

import "fmt"

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
