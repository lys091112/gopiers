package graph

import (
	"math"
)
import "container/list"

// TODO 待分析解释
// 无向图
type BiHopcorftpartiteGraph struct {
	//邻接表
	// 第一维代表A分组 第二维代表B分组
	dis   int
	n     int
	m     int
	Graph [][]int
	Dx    []int
	Dy    []int
	Mx    []int
	My    []int
}

// 交错路  增广路

func NewBiHopcorftpartite(n, m int, graph [][]int) *BiHopcorftpartiteGraph {
	if len(graph) == 0 {
		return nil
	}
	return &BiHopcorftpartiteGraph{
		Dx:    memset(-1, n),
		Dy:    memset(-1, m),
		Mx:    memset(-1, n),
		My:    memset(-1, m),
		Graph: graph,
		n:     n,
		m:     m,
	}
}

func memset(intial, size int) []int {
	p := make([]int, size)
	for i := range p {
		p[i] = intial
	}
	return p

}

//  每次使用调用BFS查找到多条增广路的路径长度都是相等的，而且都以第一次得到的dis为该次查找增广路径的最大长度
func (g *BiHopcorftpartiteGraph) Match() int {

	num := 0
	for {
		if !g.bfs() {
			break
		}
		for i := 0; i < g.n; i++ {
			// 没有被访问过 或者
			visit := memset(0, g.m)
			if g.Mx[i] == -1 && g.dfs(i, visit) {
				num += 1
			}
		}
	}
	return num
}

// 查找可匹配的长度
func (g *BiHopcorftpartiteGraph) bfs() bool {
	g.dis = math.MaxInt32
	queue := list.New()
	g.Dx = memset(-1, g.n)
	g.Dy = memset(-1, g.m)
	for i := 0; i < g.n; i++ {
		if g.Mx[i] == -1 {
			queue.PushBack(i)
			g.Dx[i] = 0
		}
	}

	for {
		if queue.Len() <= 0 {
			break
		}
		e := queue.Front()
		queue.Remove(e)
		u := e.Value.(int)

		// TODO
		if g.Dx[u] > g.dis {
			break
		}

		for _, v := range g.Graph[u] {
			if g.Dy[v] == -1 {
				g.Dy[v] = g.Dx[u] + 1

				if g.My[v] == -1 {
					g.dis = g.Dy[v]
				} else {
					g.Dx[g.My[v]] = g.Dy[v] + 1
					queue.PushBack(g.My[v])
				}
			}
		}
	}

	return g.dis != math.MaxInt32
}

func (g *BiHopcorftpartiteGraph) dfs(u int, visit []int) bool {
	for j := 0; j < len(g.Graph[u]); j++ {
		v := g.Graph[u][j]
		if visit[v] == 0 && g.Dy[v] == g.Dx[u]+1 {
			visit[v] = 1
			if g.Dy[v] == g.dis && g.My[v] != -1 {
				continue
			}

			if g.My[v] == -1 || g.dfs(g.My[v], visit) {
				g.My[v] = u
				g.Mx[u] = v
				return true
			}
		}
	}
	return false
}
