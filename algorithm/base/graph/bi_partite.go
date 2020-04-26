package graph

// 无向图
type BipartiteGraph struct {
	// 配对
	// index代表B分组 value代表A分组 用来表示B的配对前缀A
	Pair []int

	//邻接表
	// 第一维代表A分组 第二维代表B分组
	Graph [][]int
	n     int
	m     int
}

// 交错路  增广路

func NewBipartite(n, m int, graph [][]int) *BipartiteGraph {
	if len(graph) == 0 {
		return nil
	}
	p := make([]int, m)
	for i := range p {
		p[i] = -1
	}
	return &BipartiteGraph{
		Pair:  p,
		Graph: graph,
		n:     n,
		m:     m,
	}
}

// 匈牙利算法
// for loop node
// if another node not visit then update pair; continue
// else 寻找增广路径，修改节点匹配
// end
// 时间复杂度O(m*n)
func (g *BipartiteGraph) Match() int {

	num := 0
	for i, _ := range g.Graph {
		// 访问记录
		visit := make([]int, g.m)
		if g.dfs(i, visit) {
			num += 1
		}
	}
	return num

}

// 如果本次迭代没有访问过，那么访问，如果没有前置节点，则进行节点配对
// 如果进行过节点配对，则寻找增广路径，然后依次
// 把这条增广路中的匹配边改为未匹配边，把未匹配边改为匹配边，
// 这样就可以使总匹配边数+1了
func (g *BipartiteGraph) dfs(u int, visit []int) bool {
	for j := 0; j < len(g.Graph[u]); j++ {
		v := g.Graph[u][j]
		if visit[v] == 0 {
			visit[v] = 1
			if g.Pair[v] == -1 || g.dfs(g.Pair[v], visit) {
				g.Pair[v] = u
				return true
			}
		}
	}
	return false
}

