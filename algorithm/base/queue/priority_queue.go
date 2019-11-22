package queue

/**
 * 通过使用heap下的函数组成优先队列
 */

type PriorityQueue []int

func (q PriorityQueue) Len() int {
	return len(q)
}

func (q PriorityQueue) Less(i, j int) bool {
	return q[i] > q[j]
}

func (q PriorityQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *PriorityQueue) Push(x interface{}) {
	if nil == x {
		return
	}

	*q = append(*q, x.(int))
}

func (q *PriorityQueue) Pop() interface{} {
	n := q.Len()
	r := (*q)[n-1]
	*q = (*q)[:n-1]
	return r
}
