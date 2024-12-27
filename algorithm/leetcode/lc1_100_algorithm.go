package leetcode

import (
	"container/list"
	"fmt"
	"math"
	"sort"

	"github.com/lys091112/gopiers/algorithm/util"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// N: 2
// 主要是学习指针的使用
// 通过遍历指针以及移动指针来获取数值
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var result *ListNode
	var end *ListNode
	p1 := l1
	p2 := l2

	n := 0
	for {
		if p1 == nil && p2 == nil && n == 0 {
			break
		}
		r1 := fetchValue(p1)
		r2 := fetchValue(p2)

		t := r1 + r2 + n
		if t >= 10 {
			n = 1
			t = t % 10
		} else {
			n = 0
		}

		if result == nil {
			result = &ListNode{Val: t, Next: nil}
			end = result
		} else {
			newNode := &ListNode{Val: t, Next: nil}
			end.Next = newNode
			end = newNode
		}

		if p1 != nil {
			p1 = p1.Next
		}
		if p2 != nil {
			p2 = p2.Next
		}

	}

	// 考虑遍历完成后，仍有进位的情况
	if n != 0 {
		newNode := &ListNode{Val: n, Next: nil}
		end.Next = newNode
		end = newNode
	}

	return result

}

func fetchValue(p *ListNode) int {
	if p == nil {
		return 0
	}
	return p.Val
}

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

// N: 33 搜索旋转排序数组
// 题目的前提是有序数组以未知位进行了旋转
// 说明 这个是两段有序的数组 进行二分查找时，
// 考虑临界条件，到底应不应该等于边界值
func search33(nums []int, target int) int {
	if len(nums) <= 0 {
		return -1
	}

	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}

	l, r := 0, len(nums)-1
	for {
		if l > r {
			return -1
		}

		middle := (l + r) / 2
		if nums[middle] == target {
			return middle
		}

		if nums[0] <= nums[middle] { // 代表 0-middle 有序
			if nums[0] <= target && target < nums[middle] { // 是否在有序队列中
				r = middle - 1
			} else {
				l = middle + 1
			}

		} else {
			// 代表 middle+1 len(nums-1) 有序
			if nums[middle] <= target && target <= nums[r] { // 是否在有序队列中
				l = middle + 1
			} else {
				r = middle - 1
			}
		}
	}
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

// 方法实现描述 根据单调递减栈 寻找左侧第一个比当前元素a大的值
// 找到该元素b后，将当前栈中比该元素（b）小的元素依次出栈，并记录 以a为槽高度能存储的面积（桶的最短板原理，能承的水依赖于水桶的最低的板）
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

/* ------------------------ */
// N: 45 跳跃游戏 II
// 动态规划
func jump(nums []int) int {

	// dp[m] = min(dp[m-1], dp[m-2], ... , dp[0]) + 1  which dp_i + nums[i] >= m
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	dp[0] = 0
	for i := 1; i < len(nums); i++ {
		dp[i] = len(nums) + 1
		for j := i - 1; j >= 0; j-- {
			if j+nums[j] >= i && dp[i] > dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
		fmt.Printf("dp[%d]=%d\n", i, dp[i])
	}

	return dp[len(nums)-1]
}

// 贪心算法
// 记录每个节点能跳的最远的距离，满足一定的条件才更新step
func jumpV2(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	step := 0
	maxDistance := 0
	maxInx := 0

	// 记录当前能跳的最远距离
	// 如果当前节点等于记录的最远节点，那么step+1，并且更新最远距离
	for i := 0; i < len(nums)-1; i++ {
		// 记录的是从上一个节点到它能跳到的节点之间，能到达的最大距离。从而更新step，并置maxDistance为最远距离
		if maxDistance < nums[i]+i {
			maxDistance = nums[i] + i
		}
		if i == maxInx {
			step++
			maxInx = maxDistance
		}
	}
	return step
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

// N: 50 Pow(x, n) 二分查找
// 迭代法
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	k := n
	if n < 0 {
		k = -n
	}
	res := 1.0
	for i := k; i > 0; i = i / 2 {
		// 偶数直接double
		// 基数 乘x然后在double
		if i%2 != 0 {
			res *= x
		}
		x *= x
	}

	if n < 0 {
		return 1 / res
	}
	return res
}

// 递归 方式  （方法还是使用二分的方法）
func myPow50(x float64, n int) float64 {
	if n > 0 {
		return p50(x, n)
	}

	return 1.0 / p50(x, n)
}
func p50(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if n == 1 {
		return x
	}

	res := p50(x, n/2)
	if n%2 == 0 {
		return res * res
	}

	return res * res * x
}

// N: 54  螺旋矩阵
// 可以理解为一圈一圈的遍历，每次遍历完减少两行 减少两列，然后重新开始
// 另一种解法是使用方向数组[[0, 1](列+1), [1, 0], [0, -1], [-1, 0]]
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	if len(matrix) == 1 {
		return matrix[0]
	}

	return spiralOrder54(&matrix, 0, 0, len(matrix)-1, len(matrix[0])-1)
}

// 利用方向数
func spiralOrderV2(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	if len(matrix) == 1 {
		return matrix[0]
	}
	res := make([]int, 0)
	row, col := len(matrix), len(matrix[0])
	fx := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	x, y, dt := 0, 0, 0
	edge := []int{0, col - 1, row - 1, 0}
	for i := 0; i < row*col; i++ {
		res = append(res, matrix[x][y])
		x2, y2 := x+fx[dt][0], y+fx[dt][1]
		// 如果已经越过边界了，则需要转向
		if x2 < edge[0] || x2 > edge[2] || y2 < edge[3] || y2 > edge[1] {
			if dt == 1 || dt == 2 {
				edge[dt]--
			} else {
				edge[dt]++
			}
			dt = (dt + 1) % 4
			x2, y2 = x+fx[dt][0], y+fx[dt][1]
		}
		x, y = x2, y2
	}

	return spiralOrder54(&matrix, 0, 0, len(matrix)-1, len(matrix[0])-1)
}

func spiralOrder54(matrix *[][]int, x, y, row, col int) []int {
	res := make([]int, 0)
	if x == row {
		return append(res, (*matrix)[x][y:col+1]...)
	}
	if y == col {
		for i := x; i <= row; i++ {
			res = append(res, (*matrix)[i][y])
		}
		return res
	}

	// 循环获取数据 行 列 行列
	//[x,(x...col-1)]
	for c := y; c < col; c++ {
		res = append(res, (*matrix)[x][c])
	}
	// [(x ... row),col]
	for r := x; r < row; r++ {
		res = append(res, (*matrix)[r][col])
	}
	// [row,(col-1,y)]
	for c := col; c > y; c-- {
		res = append(res, (*matrix)[row][c])
	}
	// 遍历起始列
	// [(col ... x+1), y]
	for r := row; r > x; r-- {
		res = append(res, (*matrix)[r][y])
	}

	if row-x > 1 && col-y > 1 {
		res = append(res, spiralOrder54(matrix, x+1, y+1, row-1, col-1)...)
	}
	return res
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

/*
*
* N:70 爬楼梯
* d[n] = d[n-1] + d[n-2] 动态规划法

  - 示例： d[2] = ((2)(1,1))
    d[3] = d[2]+d[1] ((2,1),(1,1,1)(1,2))
  - d[4] = d[3]+d[2] ((2,1,1),(1,1,1,1),(1,2,1),(2,2),(2,1,1))
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
