package redisson

import (
	"fmt"
	"testing"
)

func TestStrings(t *testing.T) {
	r := NewRedisson("127.0.0.1", 6379,0)
	err := r.Set("rName","xx")
	if err != nil {
		fmt.Print(err)
		t.Error(err)
		return
	}
	name,err := r.Get("rName")
	fmt.Println(name)
	if name == "xx" {
		t.Log("success")
	}else {
		t.Error("failed")
	}

	n,err := r.Append("rName", "yy")
	t.Logf("n=%d",n)
	t.Log(err)

	n,err = r.BitCount("rName")
	t.Logf("bitCount=%d",n)
	n, _ = r.BitCountWithRange("rName",0,0)
	t.Logf("bitCountWithRange=%d",n)
	n, _ = r.BitCountWithRange("rName",4,4)
	t.Logf("bitCountWithRange=%d",n)

	// 修改x变为y
	n ,_ = r.SetBit("rName",7, 1)
	t.Logf("setBit=%d",n)
	
	v ,_ := r.Get("rName")
	t.Logf("rName=%s",v)

	n ,err = r.SetBit("rName",7, 0)
	if err != nil {
		t.Logf("setBit error=%v",err)
	}

	v ,_ = r.Get("rName")
	t.Logf("rName=%s",v)
}

func TestOther(t *testing.T) {
	// 1111000 1111000 1111001 1111001
	v := []byte("xxyy")
	t.Logf("v=%v\n",v)
	t.Logf("v=%b\n",v)
}

