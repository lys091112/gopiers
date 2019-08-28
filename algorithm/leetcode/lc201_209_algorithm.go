package leetcode

import "container/list"

// 207. 课程表
// 在n步内能走完全部课程，即课程交叉无环，可以进行拓扑排序
// 图结构使用邻接表表示法

type GraphNode struct {
	id    int // 节点的唯一标识
	nodes []int
	exist map[int]bool
}

func canFinish(numCourses int, prerequisites [][]int) bool {

	if numCourses <= 0 {
		return true
	}

	ingress := make([]int, numCourses)

	adj := make([]GraphNode, numCourses)

	for i := range adj {
		adj[i] = GraphNode{
			id:    i,
			nodes: make([]int, 0),
			exist: make(map[int]bool, numCourses),
		}
	}

	for _, v := range prerequisites {

		left := v[0]
		right := v[1]
		ingress[right] += 1
		if _, ok := adj[left].exist[right]; !ok {
			adj[left].exist[right] = true
			adj[left].nodes = append(adj[left].nodes, right)
		}
	}

	n := 0
	l := list.New()
	for i := range ingress {
		if ingress[i] == 0 {
			l.PushFront(i)
			n += 1
		}
	}

	for {
		if l.Len() <= 0 {
			break
		}

		p := l.Front()
		l.Remove(p)

		for _, v := range adj[p.Value.(int)].nodes {
			ingress[v] -= 1
			if ingress[v] == 0 {
				l.PushBack(v)
				n += 1
			}
		}
	}

	return n == numCourses
}
