package leetcode

// N:2661 找得第一个是某行或某列都填满的元素下标 (前提是每个元素都不相同)
func FindFirstCompleteIndex(arr []int, mat [][]int) int {
	if len(arr) == 0 || len(mat) == 0 {
		return -1
	}
	// 初始化元素所在的行列
	m, n := len(mat), len(mat[0])
	mapIndex := make(map[int][]int)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			mapIndex[mat[i][j]] = []int{i, j}
		}
	}

	// 取出元素所在的行列，并统计当前行或列已经填充的值
	// cx 统计每行被填充的列 cy统计每列被填充的行
	cx, cy := make([]int, m), make([]int, n)
	for i := 0; i < len(arr); i++ {
		if v, ok := mapIndex[arr[i]]; ok {
			cx[v[0]] += 1
			cy[v[1]] += 1
			if cx[v[0]] == n || cy[v[1]] == m {
				return i
			}
		}
	}

	return -1
}
