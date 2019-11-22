package basic

import (
	"testing"
)

func TestFloat64toBits(t *testing.T) {
	f := 62.0
	t.Logf("%#016x\n", Float64toBits(f))
}

func TestSizeof_demo(t *testing.T) {
	sizeof_demo()
}
