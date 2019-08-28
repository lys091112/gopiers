package leetcode

import "testing"

func TestNumIslands(t *testing.T) {
	grid := [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}
	resultOfInt(1, numIslands(grid), t)

	grid = [][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}
	resultOfInt(3, numIslands_with_dfs(grid), t)

	grid = [][]byte{{'1', '1', '0', '0', '0'}}
	resultOfInt(1, numIslands_with_dfs(grid), t)

	grid = [][]byte{}
	resultOfInt(0, numIslands_with_dfs(grid), t)

	grid = [][]byte{{'1'}}
	resultOfInt(1, numIslands_with_dfs(grid), t)
	grid = [][]byte{{'0'}}
	resultOfInt(0, numIslands_with_dfs(grid), t)
}
