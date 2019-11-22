package util

import (
	"testing"
)

func BenchmarkAbs(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		Abs(-12)
	}
}
