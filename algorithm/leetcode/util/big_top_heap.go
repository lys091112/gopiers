package util

// 大顶堆

type BigTopHeap []int

func (q BigTopHeap) Len() int {
	return len(q)
}

// 返回true的话，就交换两个元素
func (q BigTopHeap) Less(i, j int) bool {
	return q[i] > q[j]
}

func (q BigTopHeap) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *BigTopHeap) Push(x interface{}) {
	if nil == x {
		return
	}

	*q = append(*q, x.(int))
}

func (q *BigTopHeap) Pop() interface{} {
	n := q.Len()
	r := (*q)[n-1]
	*q = (*q)[:n-1]
	return r
}
