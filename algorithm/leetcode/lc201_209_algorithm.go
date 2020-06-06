package leetcode

import (
	"container/list"
	"math"
)

// N: 202 快乐数
func isHappy(n int) bool {
	if n == 1 {
		return true
	}

	// 记录已经访问过的值 如果再次被访问 那么说明有了循环，就不用再检测了
	old := make(map[int]bool)
	for {
		if n == 1 {
			return true
		}
		if _, ok := old[n]; ok {
			return false
		}
		old[n] = true
		newN := 0
		for {
			if n == 0 {
				break
			}
			newN += int(math.Pow(float64(n%10), 2))
			n = n / 10
		}
		n = newN
	}
}

// 207. 课程表
// 在n步内能走完全部课程，即课程交叉无环，可以进行拓扑排序
// 图结构使用邻接表表示法

type graphNode struct {
	id    int // 节点的唯一标识
	nodes []int
	exist map[int]bool
}

func canFinish(numCourses int, prerequisites [][]int) bool {

	if numCourses <= 0 {
		return true
	}

	ingress := make([]int, numCourses)

	adj := make([]graphNode, numCourses)

	for i := range adj {
		adj[i] = graphNode{
			id:    i,
			nodes: make([]int, 0),
			exist: make(map[int]bool, numCourses),
		}
	}

	for _, v := range prerequisites {

		left := v[0]
		right := v[1]
		ingress[right]++
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
			n++
		}
	}

	for {
		if l.Len() <= 0 {
			break
		}

		p := l.Front()
		l.Remove(p)

		for _, v := range adj[p.Value.(int)].nodes {
			ingress[v]--
			if ingress[v] == 0 {
				l.PushBack(v)
				n++
			}
		}
	}

	return n == numCourses
}
