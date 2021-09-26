package leetcode

// 1728. 猫和老鼠
// 能否找到一条猫获胜的路径
// 思路  深度遍历 + 剪枝 + 回溯
func canMouseWin(grid []string, catJump int, mouseJump int) bool {
	// if len(grid) == 0 {
	// 	return false
	// }

	// m,n := len(grid), len(grid[0])

	// matrix := make([][]byte, m)
	// for i := 0; i < m; i++ {
	// 	matrix[i] = make([]byte, n)
	// }
	// way := make(map[string]bool)  // 记录遍历的轨迹
	// // 猫，老鼠，食物的位置
	// cat,mouse,food := -1,-1,-1
	// for i, v := range grid {
	// 	cs := []byte(v)
	// 	for j,c := range cs {
	// 		matrix[i][j] = c
	// 		pos := i * n + j
	// 		if c == 'M' {
	// 			mouse = pos
	// 		}else if c == 'F' {
	// 			food = pos
	// 		}else if c == 'C' {
	// 			cat = pos
	// 		}
	// 	}
	// }
	return false

}