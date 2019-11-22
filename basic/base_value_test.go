package basic

import (
	"fmt"
	"testing"
	"unsafe"
)

func Test_float_demo(t *testing.T) {
	float_demo()
	t.Log("success")
}

// 变量占用的内存大小
func TestWordSize(t *testing.T) {
	var a1 int = 4
	var a2 int64 = 4
	var a3 int32 = 4
	var a4 int16 = 4

	fmt.Println(unsafe.Sizeof(a1))
	fmt.Println(unsafe.Sizeof(a2))
	fmt.Println(unsafe.Sizeof(a3))
	fmt.Println(unsafe.Sizeof(a4))
}

func TestFmt_formet_demo(t *testing.T) {
	fmt_formet_demo()
}
