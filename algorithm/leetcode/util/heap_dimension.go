package util

type Priority2DQueue [][]int

func (q Priority2DQueue) Len() int {
	return len(q)
}

// 返回true的话，就交换两个元素
func (q Priority2DQueue) Less(i, j int) bool {
	return q[i][0] < q[j][0]
}

func (q Priority2DQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *Priority2DQueue) Push(x interface{}) {
	if nil == x {
		return
	}

	*q = append(*q, x.([]int))
}

func (q *Priority2DQueue) Pop() interface{} {
	n := q.Len()
	r := (*q)[n-1]
	*q = (*q)[:n-1]
	return r
}
