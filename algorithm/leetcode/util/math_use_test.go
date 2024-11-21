package util

import (
	"testing"
)

func BenchmarkAbs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Abs(-12)
	}
}

func TestIf(t *testing.T) {
	a, b := 3, 5
	r := If(a > b, a, b).(int)
	if r == 5 {
		t.Log("success")
	} else {
		t.Fatalf("want %d,but %d", 5, r)
	}
}
