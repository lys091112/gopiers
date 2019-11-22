package util

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// TODO Stein 带实现
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
