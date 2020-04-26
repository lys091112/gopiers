package leetcode

// 200. 岛屿的个数 并查集的方式
func numIslands(grid [][]byte) int {
	if len(grid) <= 0 {
		return 0
	}

	row := len(grid)
	col := len(grid[0])

	ns := make([]int, row*col)

	// 初始化 1 初始化为当前下标， 0初始化为统一的父亲 -1 对于0的联通不做处理
	for i := range grid {
		for j := range grid[i] {
			idx := i*col + j
			if grid[i][j] == '1' {
				ns[idx] = idx
			} else {
				ns[idx] = -1
			}
		}
	}

	// 从左至右遍历当前i,j和[i][j+1] [i+1][j] 之间的联通关系
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] != '1' {
				continue
			}
			if j < col-1 && grid[i][j+1] == '1' {
				union(ns, i*col+j, i*col+j+1)
			}

			if i < row-1 && grid[i+1][j] == '1' {
				union(ns, i*col+j, col*(i+1)+j)
			}
		}
	}

	// 获取集合 排除掉父亲为-1的0集合 剩余的数量即为岛屿数量
	return len(collect(ns))

}

// 根据父节点获取集合
func collect(ns []int) map[int][]int {

	res := make(map[int][]int)

	for i, v := range ns {
		if v == -1 {
			continue
		}
		p := findParent(ns, v)
		if s, ok := res[p]; ok {
			res[p] = append(s, i)
		} else {
			m := make([]int, 1, 1)
			m[0] = i
			res[p] = m
		}
	}

	return res
}

// 合并节点
func union(ns []int, idx1 int, idx2 int) {
	f1 := findParent(ns, idx1)
	f2 := findParent(ns, idx2)

	if f1 == f2 {
		return
	}

	if f1 < f2 {
		ns[f2] = f1
	} else {
		ns[f1] = f2
	}

}

func findParent(ns []int, idx int) int {
	if idx != ns[idx] {
		ns[idx] = findParent(ns, ns[idx])
	}

	return ns[idx]
}

func numIslandsWithDfs(grid [][]byte) int {
	if len(grid) <= 0 {
		return 0
	}

	// 从[0][0] 开始， 分别计算[i][j]的子节点[i+1][j] [i][j+1] [i][j-1]
	// 如果自己的为1，那么说明联通，可以深度遍历其子节点，如果为0，则终止
	// 记录访问过的节点
	row := len(grid)
	col := len(grid[0])

	size := 0
	wayRecord := make(map[int]int, row*col)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '0' {
				continue
			}
			if _, ok := wayRecord[i*col+j]; ok {
				continue
			}
			wayRecord[i*col+j] = 1
			scan(i, j, grid, col, row, wayRecord)
			size++
		}
	}

	return size
}

func scan(i int, j int, grid [][]byte, col int, row int, wayRecord map[int]int) {
	if j < col-1 && !exist(wayRecord, i*col+j+1) && grid[i][j+1] == '1' {
		wayRecord[i*col+j+1] = 1
		scan(i, j+1, grid, col, row, wayRecord)
	}

	if j > 0 && !exist(wayRecord, i*col+j-1) && grid[i][j-1] == '1' {
		wayRecord[i*col+j+1] = 1
		scan(i, j-1, grid, col, row, wayRecord)
	}

	if i < row-1 && !exist(wayRecord, (i+1)*col+j) && grid[i+1][j] == '1' {
		wayRecord[(i+1)*col+j] = 1
		scan(i+1, j, grid, col, row, wayRecord)
	}
}

func exist(wayRecord map[int]int, key int) bool {
	_, ok := wayRecord[key]
	return ok
}
