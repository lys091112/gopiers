package util

// Abs 获取绝对值
func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// Min 获取小值
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max 获取最大值
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Divisor Stein 带实现
func Divisor(a, b int) int {
	if a < b {
		a, b = b, a
	}

	for {
		if a%b == 0 {
			return b
		}
		a, b = b, a%b
	}
}
