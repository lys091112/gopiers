package skiplist

import (
	"fmt"
	"testing"
)

func TestSkipNode(t *testing.T) {
	head := New()
	fmt.Println(&head)
	fmt.Println(*head)
	for i := 0 ;i < 20; i++ {
		head.Insert(i,99)
	}
	fmt.Println(&head)
	fmt.Println(*head)
	head.print()
}
