package leetcode

import "testing"

func TestCalculateMinimumHP(t *testing.T) {
	dungeon := [][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}
	// dungeon := [][]int{{2, 0, -1}}
	res := calculateMinimumHP(dungeon)
	t.Log(res)

}
