package algorithm

import (
	"testing"
)

func BenchmarkAddStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddStrings("11111111111111111111", "2232322222222222222222222222")
	}
}

func TestAddStrings(t *testing.T) {
	res := AddStrings("111", "22323")
	if res == "22434" {
		t.Log("success")
	} else {
		t.Error("failed")
	}
}

func TestAddNilStrings(t *testing.T) {
	res := AddStrings("", "22323")
	if res == "22323" {
		t.Log("success")
	} else {
		t.Error("failed")
	}
}
func TestAddNilStrings2(t *testing.T) {
	res := AddStrings("22323", "")
	if res == "22323" {
		t.Log("success")
	} else {
		t.Error("failed")
	}
}
func TestAddBothNilStrings2(t *testing.T) {
	res := AddStrings("", "")
	if res == "" {
		t.Log("success")
	} else {
		t.Error("failed")
	}
}
