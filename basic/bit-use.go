package basic

func findN(n int) int {
	n |= n >> 1;
	n |= n >> 2;
	n |= n >> 4;
	n |= n >> 8
	n |= n >> 16 // 整型一般是 32 位, 将右侧的都变为1
	return (n + 1) >> 1
}
