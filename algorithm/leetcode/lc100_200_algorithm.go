package leetcode

import (
	"fmt"
	"math"

	"github.com/lys091112/gopiers/algorithm/leetcode/util"
)

/**
 * N:174 地下城与勇士
 * 正着推不动可以考虑存活的最基本条件，保证存活的情况下来推导需要满足的最低生命力
 */
func calculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) <= 0 {
		return 1
	}

	m := len(dungeon)
	n := len(dungeon[0])
	dp := util.InitArray(m, n)

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				if dungeon[i][j] > 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = 1 - dungeon[i][j]
				}
			} else {
				max := math.MaxInt32
				if i+1 < m {
					if dungeon[i][j] >= dp[i+1][j] {
						max = 1
					} else if max > dp[i+1][j]-dungeon[i][j] {
						max = dp[i+1][j] - dungeon[i][j]
					}
				}
				if j+1 < n {
					if dungeon[i][j] >= dp[i][j+1] {
						max = 1
					} else if max > dp[i][j+1]-dungeon[i][j] {
						max = dp[i][j+1] - dungeon[i][j]
					}
				}
				dp[i][j] = max
			}
			fmt.Printf("%d,%d, %d\n", i, j, dp[i][j])
		}
	}
	fmt.Printf("%v\n", dp)
	return dp[0][0]
}

/**
 * N: 199
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 * BFS
 */
func rightSideView(root *TreeNode) []int {
	if nil == root {
		return []int{}
	}

	seen := make([]int, 0)
	nodes := []*TreeNode{root}
	for {
		if len(nodes) <= 0 {
			break
		}

		seen = append(seen, nodes[0].Val)
		nodes = findN199(nodes)
	}
	return seen
}

func findN199(nodes []*TreeNode) []*TreeNode {
	if len(nodes) <= 0 {
		return []*TreeNode{}
	}

	r := make([]*TreeNode, 0)
	for _, node := range nodes {
		if node.Right != nil {
			r = append(r, node.Right)
		}
		if node.Left != nil {
			r = append(r, node.Left)
		}
	}
	return r
}

// DFS 遍历 利用数组长度变化来记录中间值
func rightSideView199Dfs(root *TreeNode) []int {
	if nil == root {
		return []int{}
	}
	arr := make([]int, 0)
	dfs199(root, 0, arr)
	return arr
}

func dfs199(l *TreeNode, deep int, arr []int) {
	if l == nil {
		return
	}
	if len(arr) == deep {
		// 因为如果遍历到第一个节点，长度就会改变，后边的遍历不会在影响arr
		arr = append(arr, l.Val)
	}
	dfs199(l.Right,deep+1,arr)
	dfs199(l.Left,deep+1,arr)
}
