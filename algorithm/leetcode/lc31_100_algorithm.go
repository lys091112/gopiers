package leetcode

import (
	"container/list"
	"fmt"
	"math"
	"sort"

	"github.com/lys091112/gopiers/algorithm/leetcode/util"
)

/**
 * N: 32 最长有效括号
 * dp[i] 代表以下标为结尾的字符串的有效括号长度
 * 其他方法包括，压栈或者计数（'('代表1 ')' 代表-1)
 */
func longestValidParentheses(s string) int {
	if len(s) < 2 {
		return 0
	}

	dp := make([]int, len(s))
	dp[0] = 0
	cs := []byte(s)

	for i := 1; i < len(cs); i++ {
		if cs[i] == ')' {
			if cs[i-1] == '(' {
				dp[i] = 2
				if i >= 2 && dp[i-2] > 0 {
					dp[i] += dp[i-2]
				}
				fmt.Printf("i=%d,dp=%d\n", i, dp[i])
			} else if i-dp[i-1]-1 >= 0 && cs[i-dp[i-1]-1] == '(' {
				dp[i] = dp[i-1] + 2
				if i-dp[i-1]-2 >= 0 && dp[i-dp[i-1]-2] > 0 {
					dp[i] += dp[i-dp[i-1]-2]
				}
				fmt.Printf("i=%d,dp=%d\n", i, dp[i])
			}
		}
	}

	max := -1
	for _, v := range dp {
		if v > max {
			max = v
		}
	}

	return max
}

/**
 * N: 39
 * 数组中出现次数超过一半的数字
 * 1. 排序在中间 2. hash记录数量
 */
func majorityElement(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	sort.IntSlice(nums).Sort()
	return nums[len(nums)/2]
}

/**
 * N:42 接雨水
 * 双指针移动 时间复杂度: O(N) 空间： O(1)
 */
func trap(height []int) int {
	if len(height) <= 1 {
		return 0
	}

	area := 0
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0

	for {
		if left > right {
			break
		}

		// 低山峰被高山峰遮挡，那么可以计算了
		if height[left] <= height[right] {
			leftMax = util.Max(leftMax, height[left])
			area += leftMax - height[left]
			left++
		} else {
			rightMax = util.Max(rightMax, height[right])
			area += rightMax - height[right]
			right--
		}
	}
	return area
}

func trapV2(height []int) int {
	if len(height) <= 0 {
		return 0
	}

	area := 0
	stack := list.New()
	for i := range height {
		for {
			top := stack.Front()
			if top == nil || height[i] < height[top.Value.(int)] {
				break
			}

			stack.Remove(top)
			bottom := top.Value.(int)

			for {

				if stack.Front() == nil || height[stack.Front().Value.(int)] != height[bottom] {
					break
				}
				stack.Remove(stack.Front())
			}

			// 将要入栈的元素大于栈顶元素 出栈并计算面积
			if stack.Len() > 0 {
				top = stack.Front()
				area += (util.Min(height[i], height[top.Value.(int)]) - height[bottom]) * (i - top.Value.(int) - 1)
			}
		}
		fmt.Printf("i=%d,area=%d\n", i, area)
		stack.PushFront(i)
	}

	return area
}

// N: 46 全排列
// 回溯算法(DFS)
func permute(nums []int) [][]int {
	if len(nums) <= 1 {
		return [][]int{nums}
	}

	t := make([][]int, 0)
	record := make(map[int]bool, len(nums))
	path := make([]int, 0)
	permuteDfs(nums, path, &t, record)

	return t

}

// 传递切片指针，避免修改切片不能作用到主方法中
func permuteDfs(nums, path []int, res *[][]int, record map[int]bool) {
	if len(nums) == len(path) {
		w := make([]int, len(path))
		copy(w, path[:])
		*res = append(*res, w)
		return
	}

	for i, v := range nums {
		if _, ok := record[i]; ok {
			continue
		}
		path = append(path, v)
		record[i] = true
		permuteDfs(nums, path, res, record)

		delete(record, i)
		path = path[0 : len(path)-1]
	}

}

// BFS 广度有限遍历解
func permute01(nums []int) [][]int {
	if len(nums) <= 1 {
		return [][]int{nums}
	}
	res := make([][]int, 0)
	for i, v := range nums {
		t := permute(subNums(nums, i))
		for _, v2 := range t {
			res = append(res, append([]int{v}, v2...))
		}
	}
	return res
}
func subNums(nums []int, i int) []int {
	if i == 0 {
		return nums[1:]
	}
	if i == len(nums)-1 {
		return nums[0 : len(nums)-1]
	}
	r := make([]int, 0)
	r = append(r, nums[0:i]...)
	r = append(r, nums[i+1:]...)
	return r
}

/**
 * N51: n皇后问题
 * 回溯
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

/**
*  N: 64  最小路径和
 *  dp[i][j] = grid[i][j] + min(dp[i+1][j], dp[i][j+1])
*/
func minPathSum(grid [][]int) int {

	if len(grid) <= 0 {
		return 0
	}

	fmt.Printf("grid=%v\n", grid)

	row := len(grid)
	col := len(grid[0])

	dp := util.InitArray(row, col)
	dp[row-1][col-1] = grid[row-1][col-1]
	for i := row - 1; i >= 0; i-- {
		for j := col - 1; j >= 0; j-- {
			if i != row-1 || j != col-1 {
				way := int(math.MaxInt32)
				if j < col-1 {
					way = dp[i][j+1]
				}
				if i < row-1 && dp[i+1][j] < way {
					way = dp[i+1][j]
				}
				dp[i][j] = grid[i][j] + way
				fmt.Printf("i=%d,j=%d,way=%d,dp=%d\n", i, j, way, dp[i][j])
			}
		}
	}
	fmt.Printf("dp=%v\n", dp)
	return dp[0][0]
}

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
* N:72 编辑距离
* TODO 使用递归实现
 */

func minDistance(a string, b string) int {

	if len(a) <= 0 {
		return len(b)
	}

	dp := util.InitArray(len(a)+1, len(b)+1)
	// 初始条件的初始化很重要
	for i := 0; i <= len(a); i++ {
		dp[i][0] = i
	}
	for j := 0; j <= len(b); j++ {
		dp[0][j] = j
	}

	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			// 插入
			m1 := dp[i][j-1]
			// 删除
			m2 := dp[i-1][j]
			// 替换
			m3 := dp[i-1][j-1]
			if a[i-1] == b[j-1] {
				m3--
			}
			fmt.Printf("(i,j)=%d,%d,<m1,m2,m3>=%d,%d,%d,choice=%d\n", i, j, m1, m2, m3, 1+Min(m1, m2, m3))
			dp[i][j] = 1 + Min(m1, m2, m3)
		}
	}
	fmt.Printf("%s,%s, %v\n", a, b, dp)

	return dp[len(a)][len(b)]
}

// Min 比较大小返回最小值
func Min(a, b, c int) int {
	if a > b {
		a = b
	}
	if a > c {
		a = c
	}
	return a
}

// N: 84 矩形面积
// 单链表 链表中维持当前矩形的有效高度
// TODO
func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}

	if len(heights) == 1 {
		return heights[0]
	}

	heights = append(heights, 0)

	area := 0
	stack := list.New()
	for i := range heights {
		for {
			// 保持栈顶小于当前元素
			// 如果不小于，那么移除栈顶，计算当前高度的最大面积，否则直接入栈，因为添加了末尾元素0，所以即便是连续的递增，最终也会被计算
			f := stack.Front()
			if f == nil || heights[f.Value.(int)] < heights[i] {
				break
			}

			tArea := 0
			stack.Remove(f)
			top := f.Value.(int)
			if stack.Len() == 0 {
				tArea = heights[top] * i
			} else {
				tArea = heights[top] * (i - stack.Front().Value.(int) - 1)
			}
			if tArea > area {
				area = tArea
			}
		}
		stack.PushFront(i)
	}
	return area
}
