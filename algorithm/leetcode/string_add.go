package leetcode

import "bytes"

// TODO{crescent} string[]获取数组元素后的类型 高效的字符串拼接 rune的含义 没有++ --， golang 引用类型

/**
 * 字符串拼接
 * 给定的两个字符串长度小于5000，只包含数字0-9，同时不包含任何的前导0，不允许使用任何的大整数计算库
 **/
const (
	strConst = "0123456789"
)

func AddStrings(str1, str2 string) string {
	num1 := len(str1)
	num2 := len(str2)

	var res bytes.Buffer = bytes.Buffer{}
	flag := 0
	i, j := 0, 0
	for i, j = num1-1, num2-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		sum := int(str1[i]-'0') + int(str2[j]-'0') + flag
		flag = carry_propagation(&res, sum)
	}

	flag = endString(&res, str1, i, flag)
	flag = endString(&res, str2, j, flag)

	if flag == 1 {
		res.WriteString("1")
	}

	return reverse(res.String())
}

func endString(res *bytes.Buffer, str string, index int, flag int) int {
	for ; index >= 0; index-- {
		sum := int(str[index]-'0') + flag
		flag = carry_propagation(res, sum)
	}
	return flag
}

func carry_propagation(res *bytes.Buffer, sum int) int {
	res.WriteString(string(strConst[sum%10]))
	return sum / 10
}

func reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
