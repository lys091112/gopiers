package algorithm

import (
	"testing"
)

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
