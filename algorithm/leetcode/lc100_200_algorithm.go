package leetcode

import (
	"fmt"
	"math"

	"github.com/lys091112/gopiers/algorithm/leetcode/util"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
	func dfs( root ) {
		append(root)
		dfs(root.left)
		dfs(root.right)
	}
*/

// N: 102 DFS
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := make([][]int, 0)
	// dfs
	dfs102(root, &result, 0)
	return result
}

func dfs102(node *TreeNode, result *[][]int, level int) {
	if node == nil {
		return
	}
	if len(*result) <= level {
		*result = append(*result, make([]int, 0))
	}
	(*result)[level] = append((*result)[level], node.Val)

	dfs102(node.Left, result, level+1)
	dfs102(node.Right, result, level+1)
}

// MinStackV2 N: 155 最小栈
// min(stack) = min(curr, min(stack-1))
// 最小栈栈顶始终记录的是当前栈的最小值，因为栈没有乱序出栈的
type MinStackV2 struct {
	data   []int
	minSet []int
}

// Constructor 初始化
func Constructor() MinStackV2 {
	m2 := MinStackV2{
		data:   make([]int, 0),
		minSet: make([]int, 0),
	}
	return m2
}

// Push  push
func (th *MinStackV2) Push(x int) {

	th.data = append(th.data, x)

	// 最关键的地方，如果当前值不大于当前最小栈的栈顶，那么入栈，记录加入x后的栈最小值
	if len(th.minSet) <= 0 || x <= th.minSet[len(th.minSet)-1] {
		th.minSet = append(th.minSet, x)
	}
}

// Pop pop
func (th *MinStackV2) Pop() {
	dataLen := len(th.data)
	if dataLen > 0 {
		top := th.Top()
		th.data = th.data[0 : dataLen-1]
		if th.GetMin() == top {
			th.minSet = th.minSet[0 : len(th.minSet)-1]
		}
	}
}

// Top top
func (th *MinStackV2) Top() int {
	return th.data[len(th.data)-1]
}

// GetMin d
func (th *MinStackV2) GetMin() int {
	return th.minSet[len(th.minSet)-1]
}

// N: 136
// 异或运算 任何数与自身异或都得0
func singleNumber(nums []int) int {
	r := nums[0]
	for i := 1; i < len(nums); i++ {
		r = r ^ nums[i]
	}
	return r
}

// N: 141 环形链表
// 1. 快慢指针来判断是否有重合
// 2. 通过hash来记录某个节点是否已经访问过了
func hasCycle(head *ListNode) bool {
	if nil == head {
		return false
	}
	if head.Next == nil {
		return false
	}

	slow, fast := head, head

	for {
		if slow == nil {
			return false
		} else if nil == fast || nil == fast.Next {
			return false
		} else {
			slow = slow.Next
			fast = fast.Next.Next
			if slow == fast {
				return true
			}
		}
	}
}

/**
 * N: 144
 * 方法的使用
 */
 func preorderTraversal(root *TreeNode) (vals []int) {
	 if root == nil {
		 return []int{}
	 }
	 // 通过函数闭包的使用
	 var preorder func(*TreeNode)
	 preorder = func(node *TreeNode) {
		if node == nil {
			return
		} 
		vals = append(vals, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	 }

	 preorder(root)
	 return 

	//  result := make([]int,0)
	//  result = traversal(root, result)
	// return result 
}

// 迭代方式
 func preorderTraversalV2(root *TreeNode) (vals []int) {
	 if root == nil {
		 return []int{}
	 }

	 stack := []*TreeNode{}
	 node := root
	 for node != nil || len(stack) > 0 {
		 // 遍历所有的左节点并依次入栈
		for node != nil {
			stack = append(stack, node)
			vals = append(vals, node.Val)
			node = node.Left
		}
		// 获取右节点
		node = stack[len(stack) - 1].Right
		stack = stack[:len(stack)-1] // 最后一个节点出栈
	 }

	 return
 }

func traversal(node *TreeNode, result []int) []int {
	if node == nil {
		return result
	}
	result = append(result, node.Val)
	result = traversal(node.Left, result)
	result = traversal(node.Right,result)
	return result
}

/**
 * N: 169 
 */
func majorityElementImprove(nums []int) int {

	var k int = nums[0]
	count := 0

	for _, v := range nums {
		if count == 0 {
			k = v
			count++
		}else {
			if k == v {
				count++
			}else {
				count--
			}
		}
	}

	return k 
}

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
	dfs199(l.Right, deep+1, arr)
	dfs199(l.Left, deep+1, arr)
}
