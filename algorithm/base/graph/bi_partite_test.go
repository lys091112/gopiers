package graph

import "testing"

func TestMatch(t *testing.T) {
	graph := [][]int{{0, 1, 2}, {1}, {0, 2, 3}, {2}}

	biGraph := NewBipartite(4, 4, graph)

	num := biGraph.Match()
	t.Log("num=", num)
	t.Logf("value=%v", *biGraph)

	biGraph2 := NewBiHopcorftpartite(4, 4, graph)
	t.Logf("=====%p",biGraph2)
	num = biGraph2.Match()
	t.Log("num=", num)
	t.Logf("value=%v", *biGraph2)
}
