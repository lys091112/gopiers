package mystrings

import (
	"bytes"
	"fmt"
	"strings"
	"time"
	"unsafe"
)

const (
	Sprintf = iota
	StringAdd
	Join
	BufferAppender
)

var base = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var baseSlice []string

func init() {
	for i := 0; i < 200; i++ {
		baseSlice = append(baseSlice, base)
	}
}

/**** 字符串拼接 ****/

func AddString() string {
	res := ""
	for _, v := range baseSlice {
		res += v
	}
	return res
}

func JoinString() string {
	return strings.Join(baseSlice, "")
}

func SprintfString() string {
	res := ""
	for _, v := range baseSlice {
		res = fmt.Sprintf("%s%s", res, v)
	}
	return res
}

func BuilderString() string {
	var builder strings.Builder
	builder.Grow(200 * len(base)) // 提前分配好，避免重新分配
	for _, v := range baseSlice {
		builder.WriteString(v)
	}
	return builder.String()
}

func BufferString() string {
	var buffer bytes.Buffer
	for _, v := range baseSlice {
		buffer.WriteString(v)
	}
	return buffer.String()
}

// 提前预分配好,byteSlice 和 strings.Buidler 性能差不多，
func ByteSliceString() string {
	buf := make([]byte, 0, 200*len(base))
	// buf := make([]byte, 200*len(base))
	for _, v := range baseSlice {
		buf = append(buf, v...)
	}
	return *(*string)(unsafe.Pointer(&buf))
	// return string(buf)
}

func MakeTest() {
	buf := make([]byte, len(base)) // 等价于分配了长度62的空值
	buf = append(buf, base...)
	fmt.Printf("-----------%d:%s\n", len(buf), string(buf[62:]))
	// fmt.Printf("-----------%d:%s\n", len(buf), string(buf)) // 不出结果，前62个空值无法被转为字符串

	buf2 := make([]byte, 0, len(base))
	buf2 = append(buf2, base...)
	fmt.Printf("========%d:%s\n", len(buf2), string(buf2))
	time.Sleep(1 * time.Second)
}

// https://zhuanlan.zhihu.com/p/48865454 字符串性能对比
