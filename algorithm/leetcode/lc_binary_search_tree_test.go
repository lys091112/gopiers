package leetcode

import (
	"fmt"
	"testing"
)

func TestCountSmaller(t *testing.T) {

	res := countSmaller([]int{26, 78, 27, 100, 33, 67, 90, 23, 66, 5, 38, 7, 35})

	fmt.Println(res)

	res2 := countSmaller2([]int{26, 78, 27, 100, 33, 67, 90, 23, 66, 5, 38, 7, 35})

	fmt.Println(res2)
}
