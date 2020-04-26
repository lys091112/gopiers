package base

import (
	"fmt"
	"testing"
)

// todo learn panix recover的使用
func TestUnion(t *testing.T) {

	//defer func() {
	//	recover()
	//}()
	unionFind := UnionFind{}
	unionFind.MakeSet(6)
	unionFind.Union(1, 2)
	unionFind.Union(1, 3)
	unionFind.Union(2, 4)
	unionFind.Union(0, 5)
	unionFind.Union(1, 5)
	//unionFind.Union(0, 6)

	res := unionFind.Collections()

	for _, v := range res {
		fmt.Println(v)
	}
}
