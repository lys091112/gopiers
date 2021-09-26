package util

// 创建空二维数组
func InitArray(row int, col int, initValue... int) [][]int {
	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
		if len(initValue) == 1 {
			for j := 0;j < col; j ++ {
				dp[i][j] = initValue[0]
			}
		}
	}
	return dp
}


// 翻转数组切片
func Reverse(nums []int) []int {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}

	return nums
}
