package mystrings

import (
	"fmt"
	"testing"
	"time"
)

func TestStringAppender(test *testing.T) {
	k := 4
	d := [4]time.Duration{}
	for i := 0; i < k; i++ {
		d[i] = StringAppender(i, 10000)
	}

	for i := 0; i < k-1; i++ {
		fmt.Printf("way %d is %6.1f times of way %d\n", i, float32(d[i])/float32(d[k-1]), k-1)
	}

	test.Log("sucdess")

}
