package leetcode

import (
	"container/list"
	"fmt"
	"math"
)

/**
* N:70 爬楼梯
* d[n] = d[n-1] + d[n-2] 动态规划法

* 示例： d[2] = ((2)(1,1))
		 d[3] = d[2]+d[1] ((2,1),(1,1,1)(1,2))
*		 d[4] = d[3]+d[2] ((2,1,1),(1,1,1,1),(1,2,1),(2,2),(2,1,1))
*/
func climbStairs_70(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	d := make([]int, n+1)
	d[1] = 1
	d[2] = 2
	for i := 3; i <= n; i++ {
		d[i] = d[i-1] + d[i-2]
	}

	return d[n]
}

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

/**
 * N51: n皇后问题
 */
func solveNQueens(n int) [][]string {
	if n <= 0 {
		return nil
	}

	res := make([][]string, 0)
	a := make([]int, n)
	queen(res, a, n, 0)

	return res
}

func queen(res [][]string, a []int, n, rowN int) {
	if n == rowN {
		fmt.Println(a)
		return
	}

	for i := 0; i < n; i++ {
		a[rowN] = i

		// 检验该点是否和已经选择的点满足匹配
		f := true
		for j := 0; j < rowN; j++ {
			if !isFit51(a, j, rowN) {
				f = false
				break
			}
		}

		if f {
			queen(res, a, n, rowN+1)
		}
	}
}

func isFit51(a []int, rowM, rowN int) bool {
	// 查看是否是同列（遍历的时候，用下标代表了行号，所以不用检测行)
	if a[rowN] == a[rowM] {
		return false
	}

	t := a[rowM] - a[rowN]
	if t < 0 {
		t = -t
	}
	// 看是否是对角线，即 行列的比例是1：1，即行增加m行，列增加m列，
	if t == rowN-rowM {
		return false
	}

	return true
}
