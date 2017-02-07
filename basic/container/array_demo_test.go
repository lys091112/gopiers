package container

import (
	"testing"
)

/**
* golang 方法的测试为TestXxx或者test_xxx
* 对于公共方法使用TestXxx
* 对于私有方法使用Test_xxx
**/
func Test_array(test *testing.T) {
	array()
	test.Log("result is success")
}

func Test_listdemo(test *testing.T) {
	listdemo()
	test.Log("result is success")
}
