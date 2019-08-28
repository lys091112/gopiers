package leetcode

import (
	"container/list"
	"math"
)

/**
 * N: 98
 * 考虑临界 为空 ，或者 为临界值的情况
 * 每次比较，都是以当前点为中心，将比较切为两部分
 * left ==》 (prev.min, node.val)
 * right ==> (node.val, prev.max)  prev 为父节点的取值范围
 */
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return false
	}
	var treeRange = list.New()
	var nodes = list.New()

	treeRange.PushBack(Pair{math.MinInt64, math.MaxInt64})
	nodes.PushBack(root)

	for {
		if nodes.Len() <= 0 {
			break
		}

		node := nodes.Front().Value.(*TreeNode)
		p := treeRange.Front().Value.(Pair)
		// check left
		if node.Left != nil {
			if node.Left.Val > p.Key.(int) && node.Left.Val < node.Val {
				treeRange.PushBack(Pair{p.Key, node.Val})
				nodes.PushBack(node.Left)
			} else {
				return false
			}
		}

		// check right
		if node.Right != nil {
			if node.Right.Val > node.Val && node.Right.Val < p.Value.(int) {
				treeRange.PushBack(Pair{node.Val, p.Value})
				nodes.PushBack(node.Right)
			} else {
				return false
			}
		}
		nodes.Remove(nodes.Front())
		treeRange.Remove(treeRange.Front())
	}

	return true
}
