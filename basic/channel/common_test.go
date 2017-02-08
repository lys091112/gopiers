package channel

import (
	"testing"
)

func Test_g(test *testing.T) {
	g()
	test.Log("common test success")
}

func TestChannelSelect(test *testing.T) {
	result := ChannelSelect()
	if result != "timeout" {
		test.Errorf("result is %s,except is 'timeout'", result)
	}
}
