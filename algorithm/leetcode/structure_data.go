package leetcode

type Pair struct {
	Key   interface{}
	Value interface{}
}

func (Pair) of(key interface{}, value interface{}) *Pair {
	return &Pair{Key: key, Value: value}
}

type TreeNode struct {
	Val   int
	Count int
	Right  *TreeNode
	Left   *TreeNode
}
