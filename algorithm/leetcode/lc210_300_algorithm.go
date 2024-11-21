package leetcode

import (
	"container/list"
	"fmt"
	"sort"
)

// Trie  结构数
type Trie struct {
	isEnd  bool
	childs map[string]*Trie
}

// TireConstructor  Initialize your data structure here.
func TireConstructor() *Trie {
	return &Trie{
		isEnd:  false,
		childs: make(map[string]*Trie, 26)}

}

/** Inserts a word into the trie. */
func (th *Trie) insert(word string) {
	if word == "" {
		return
	}

	ws := []rune(word)

	tmp := th
	for i := 0; i < len(ws); i++ {
		key := string(ws[i])
		if c, ok := tmp.childs[key]; ok {
			tmp = c
		} else {
			newTrie := TireConstructor()
			tmp.childs[key] = newTrie
			tmp = newTrie
		}
	}
	tmp.isEnd = true
}

/* -------------------------------------------- */
// N: 210 课程表  拓扑排序
// BFS 或者经典算法
func findOrder(numCourses int, prerequisites [][]int) []int {
	// 统计各节点的入度
	// 从第一个入度为0的节点出发
	// 查找它的下一个节点，并删除依赖关系
	// 从新找入度为0的节点知道入度为0的节点为空  如果剩余的节点不为空说明有环
	degree := map[int]int{}
	for _, v := range prerequisites {
		degree[v[0]]++
	}
	stack := list.New()
	for i := 0; i < numCourses; i++ {
		if degree[i] == 0 {
			stack.PushFront(i)
		}
	}
	result := make([]int, 0)
	for {
		if stack.Len() <= 0 {
			break
		}
		top := stack.Front()
		stack.Remove(top)
		topValue := top.Value.(int)
		result = append(result, topValue)

		// 查找前驱节点是栈节点的值，入度减一
		for _, v := range prerequisites {
			if v[1] == topValue {
				degree[v[0]]--
				if degree[v[0]] == 0 {
					stack.PushFront(v[0])
				}
			}
		}
	}
	if len(result) != numCourses {
		return []int{}
	}

	return result
}

// DFS 深度遍历
func findOrderV2(numCourses int, prerequisites [][]int) []int {
	// 从任一为访问过的节点开始	，
	// value= 0:未访问 1:临时访问 2:永久访问
	visit := map[int]int{}
	vGroup := map[int][]int{}
	// 找到节点集合
	for _, v := range prerequisites {
		vGroup[v[1]] = append(vGroup[v[1]], v[0])
	}
	res := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if visit[i] == 0 && !visitOrder(i, &res, vGroup, visit) {
			// 有环存在
			return []int{}
		}
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

// 深度遍历未访问的节点
func visitOrder(n int, res *[]int, vGroup map[int][]int, visit map[int]int) bool {
	// 已经访问过，无需再访问
	if visit[n] == 2 {
		return true
	}
	if visit[n] == 1 {
		return false // 有环存在
	}
	visit[n] = 1 // 先标记临时节点
	if nodes, ok := vGroup[n]; ok {
		for _, m := range nodes {
			if !visitOrder(m, res, vGroup, visit) {
				return false

			}
		}
	}
	visit[n] = 2 // 标记为永久访问几点，并加入结果集
	*res = append(*res, n)
	return true
}

/* --------------- 212 --------------------------- */
// N:212 trie + 深度遍历
func findWords(board [][]byte, words []string) []string {

	if len(board) <= 0 {
		return []string{}
	}

	root := TireConstructor()

	for _, v := range words {
		root.insert(v)
	}

	var keys = make(map[string]bool)
	result := make([]string, 0)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			tmp := root
			if t, ok := tmp.childs[string(board[i][j])]; !ok {
				continue
			} else {
				var ways = make(map[int]bool)
				strs := find(t, board, i, j, ways)
				if len(strs) <= 0 {
					continue
				}
				for _, v := range strs {
					if _, ok := keys[v]; !ok {
						result = append(result, v)
						keys[v] = true
					}
				}
			}
		}
	}

	sort.Strings(result)
	return result
}

// 1. [i][j] [i-1][j] [i+1][j] [i][j-1] [i][j+1]
// 2. 走过的路不能在走回来
// 3. 如果未走完，不能停止,可能有多重匹配
// 4. 子递归走过的路，在其他分支仍可以重新走
// 通过递归遍历查找
func find(trie *Trie, board [][]byte, i int, j int, ways map[int]bool) []string {
	col := len(board[0])
	ways[i*col+j] = true

	result := make([]string, 0)

	if i > 0 && !ways[(i-1)*col+j] && trie.childs[string(board[i-1][j])] != nil {
		r1 := find(trie.childs[string(board[i-1][j])], board, i-1, j, ways)
		if len(r1) > 0 {
			for _, v := range r1 {
				result = append(result, string(board[i][j])+v)
			}
		}
		delete(ways, (i-1)*col+j) // 为了让下一重递归重新走
	}

	if i < len(board)-1 && !ways[(i+1)*col+j] && trie.childs[string(board[i+1][j])] != nil {
		r1 := find(trie.childs[string(board[i+1][j])], board, i+1, j, ways)
		if len(r1) > 0 {
			for _, v := range r1 {
				result = append(result, string(board[i][j])+v)
			}
		}
		delete(ways, (i+1)*col+j)
	}
	if j > 0 && !ways[i*col+j-1] && trie.childs[string(board[i][j-1])] != nil {
		r1 := find(trie.childs[string(board[i][j-1])], board, i, j-1, ways)
		if len(r1) > 0 {
			for _, v := range r1 {
				result = append(result, string(board[i][j])+v)
			}
		}
		delete(ways, i*col+j-1)
	}

	if j < len(board[0])-1 && !ways[i*col+j+1] && trie.childs[string(board[i][j+1])] != nil {
		r1 := find(trie.childs[string(board[i][j+1])], board, i, j+1, ways)
		if len(r1) > 0 {
			for _, v := range r1 {
				result = append(result, string(board[i][j])+v)
			}
		}
		delete(ways, i*col+j+1)
	}

	if trie.isEnd {
		return append(result, string(board[i][j]))
	}
	return result
}

// N:221 最大正方形
// 动态规划 多分配一个单位的空间，是为了避免边界判断
func maximalSquare(matrix [][]byte) int {
	if len(matrix) <= 0 {
		return 0
	}

	// dp代表的是右下角i，j为点的正方形的最长边
	dp := make([][]int, len(matrix)+1)
	for i := 0; i < len(matrix)+1; i++ {
		dp[i] = make([]int, len(matrix[0])+1)
		dp[i][0] = 0
	}
	// if matrix[i][j] == 1 : dp[i][j] = 1 + min(dp[i-1][j],dp[i-1][j-1],dp[i][j-1])
	// if matrix[i][j] == 1 : dp[i][j] = 0
	max := 0
	for i := 1; i <= len(matrix); i++ {
		for j := 1; j <= len(matrix[i-1]); j++ {
			if matrix[i-1][j-1] == 1 {
				dp[i][j] = 1 + min221(dp[i-1][j], dp[i-1][j-1], dp[i][j-1])
				if max < dp[i][j] {
					max = dp[i][j]
					fmt.Printf("i=%d,j=%d,area=%d\n", i, j, max*max)
				}
			}
		}
	}

	return max * max
}

func min221(a, b, c int) int {
	min := a
	if min > b {
		min = b
	}
	if min > c {
		min = c
	}
	return min
}

// N:242
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counter := [26]byte{}

	for i := 0; i < len(s); i++ {
		counter[s[i]-'a']++
		counter[t[i]-'a']--
	}

	for _, v := range counter {
		if v != 0 {
			return false
		}
	}
	return true
}

// N: 274 寻找h数
// 实现思路是：h的含义是至少发布h个文档，且这个h个文档每个文档的引用次数不少于h，其他的非h个文档内的引用数少于h
// 所以本题可以理解为至少发布了h次，且每篇最少被引用h的最大h值。 当某个数量的文档数不满足时，比他大的文档数更不满足
// k篇出现的次数不足k次 ==》 意味着 k+1篇出现的 次数一定不足 k+1次
// https://zhuanlan.zhihu.com/p/355503124
func HIndex(citations []int) int {
	if len(citations) == 0 {
		return 0
	}
	l, r := 0, len(citations)
	for l < r {
		mid := (l + r + 1) >> 1 // 这里的mid代指的是文档数，即看一半以上的文档数是否满足h
		if checkCs(citations, mid) {
			l = mid // 因为下次查询的结果可能包含mid，所以需要将mid包含进下一次的迭代里
		} else {
			r = mid - 1
		}
	}
	return r
}

// 检查是否有h个文档的引用数大于h
func checkCs(cs []int, h int) bool {
	cnt := 0
	for _, v := range cs {
		if v >= h {
			cnt += 1
		}
	}
	return cnt >= h
}
