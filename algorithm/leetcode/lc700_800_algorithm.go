package leetcode

import (
	"bytes"
	"fmt"
	"math"
	"sort"
	"strconv"
)

/**
*  N: 785
* 即无向图G为二分图的充要条件是，G至少有两个顶点，且其所有回路的长度均为偶数
*  判定方法：根据定义比根据定理更好判定，只需要从第一个点出发，将其染成颜色1，则与其所有相连的所有终点都染成颜色2，如果与其相连的点已经被染成颜色1，则图G不是二分图
*
* TODO 可以尝试用队列，而非递归
**/
func isBipartite(graph [][]int) bool {
	if len(graph) <= 0 {
		return false
	}

	mp := make([]int, len(graph))
	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[i]); j++ {
			node := graph[i][j]
			// 遍历临接点，如果访问过，那么continue
			if mp[node] == 0 {
				if !dfs(node, 1, &graph, mp) {
					return false
				}
			}
		}
	}

	return true

}

func dfs(u, f int, graph *[][]int, mp []int) bool {
	mp[u] = f
	for i := 0; i < len((*graph)[u]); i++ {
		node := (*graph)[u][i]
		if mp[node] == f {
			return false
		}
		if mp[node] == 0 && !dfs(node, -f, graph, mp) {
			return false
		}
	}
	return true
}

/**
 *  N: 720
 *
 * 如果单词的长度为1， 那么必定为 空字符加上本身
 */
func longestWord(words []string) string {
	if len(words) <= 0 {
		return ""
	}

	var kvs = make(map[string]int)
	for k, v := range words {
		kvs[v] = k
	}

	// 按字符串长度排序
	p := StringSlice2(words)
	sort.Sort(p)

	maxLen := math.MaxInt64
	result := make([]string, 0)
	for _, v := range p {
		if checkWords(v, kvs) {
			if maxLen == math.MaxInt64 {
				maxLen = len(v)
			} else if maxLen > len(v) {
				break
			}
			result = append(result, v)
		}
	}

	if len(result) <= 0 {
		return ""
	}
	if len(result) == 1 {
		return result[0]
	}

	sort.Sort(sort.StringSlice(result))
	return result[0]

}

// 检测单词是否存在
func checkWords(s string, kvs map[string]int) bool {
	if len(s) == 1 {
		return true
	}
	var tmp = s[0 : len(s)-1]

	for {
		if len(tmp) <= 0 {
			return false
		}
		if _, ok := kvs[tmp]; !ok {
			return false
		}
		if len(tmp) == 1 {
			return true
		}
		tmp = tmp[0 : len(tmp)-1]
	}
}

/**
* N: 753
* 因为每个位置都有k总选择，因为在这个欧拉路中每个点的度都为偶数，即存在欧拉回路，因此无论从哪个点开始，都可以最终回到该点
* Hierholzer算法 复杂度是O(n * k^(n))
* 算法思想：
* 建一个栈cur_path保存当前遍历路径，一个vector circuit保存结果路径。
* 先按DFS遍历，如果当前结点还存在未访问过的边，把结点压入cur_path中，如果当前结点所有边都访问过了，把当前结点push_back进circuit，然后从cur_path中pop一个点出来作为新的当前结点
 */
func crackSafe(n int, k int) string {

	if n <= 0 || k <= 0 {
		return ""
	}

	var buf bytes.Buffer
	if n == 1 {
		for i := 0; i < k; i++ {
			buf.WriteString(strconv.Itoa(i))
		}
		return buf.String()
	}

	// 记录接点是否被访问过
	visit := make(map[string]bool)

	// 从 0 * (n-1) 开始 遍历
	s := ""
	for i := 0; i < n-1; i++ {
		s += "0"
	}
	dfs753(&buf, visit, s[:], k)
	buf.WriteString(s)
	fmt.Println("count=", count)
	return buf.String()
}

var count = 0

func dfs753(buf *bytes.Buffer, visit map[string]bool, s string, k int) {
	for i := 0; i < k; i++ {
		nei := s + strconv.Itoa(i)
		// 如果已经被访问过，那么继续
		if visit[nei] {
			continue
		}
		count++
		fmt.Println(s, "-------", nei)
		visit[nei] = true
		dfs753(buf, visit, nei[1:], k)
		(*buf).WriteString(strconv.Itoa(i))
		// fmt.Println("xx", buf.String())
	}
}

type StringSlice2 []string

func (p StringSlice2) Len() int           { return len(p) }
func (p StringSlice2) Less(i, j int) bool { return len(p[i]) > len(p[j]) }
func (p StringSlice2) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
