package weekly

import (
	"container/heap"
	"fmt"
)

// N:5709. 最大升序子数组和
func maxAscendingSum(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	end :=  0
	res,tmp := 0, 0

	for {
		if end >= len(nums) {
			if res < tmp {
				res = tmp
			}
			break
		}
		if end > 0 && nums[end] <= nums[end-1] {
			// 不能持续上升，修改序列和
			if res < tmp {
				res = tmp
			}
			tmp = 0
		}

		tmp += nums[end]
		end++
	}
	return res
}

// N: 5710. 积压订单中的订单总数
// 并没有要求订单只能买之前的sell订单,所以应该整理出全部的订单后在进行统一处理
func getNumberOfBacklogOrders(orders [][]int) int {

	sell := make(map[int]int,0)
	buy := make(map[int]int,0)

	for _, v := range orders {
		if v[2] == 0 {
			buy[v[0]] += v[1]
		}else {
			sell[v[0]] += v[1]
		}
	}

	buyOrder := V2PriorityQueue{}
	sellOrder := V2PriorityQueue{}
	for k,v := range buy {
		heap.Push(&buyOrder,[]int{k,v})
	}
	for k, v := range sell {
		heap.Push(&sellOrder,[]int{k,v})
	}

	fmt.Printf("sellOrder= %v,buyOrder=%v \n",sellOrder,buyOrder)
	res := 0
	for {
		bF := heap.Pop(&buyOrder)
		if nil == bF {
			break
		}

		b :=bF.([]int)
		for {
			if b[1] <= 0 {
				break
			}
			sF := heap.Pop(&sellOrder)
			if nil == sF {
				break
			}
			s := sF.([]int)
			if s[0] > b[0]{
				heap.Push(&sellOrder,s)
				break
			}else {
				if s[1] <= b[1] {
					b[1] -= s[1]
				}else {
					s[1] -= b[1]
					b[1] = 0
					heap.Push(&sellOrder,s)
				}
			}
		} 
		res += b[1]
		fmt.Printf("sellOrder = %v, buyOrder = %v, res = %d\n",sellOrder, buyOrder, res)
	}

	for {
		sF := heap.Pop(&sellOrder)
		if nil== sF {
			break
		}
		res += sF.([]int)[1]
		fmt.Println("res = ", res)
	}
	return res % (1000000000 + 7)
	
}

type V2PriorityQueue [][]int

func (q V2PriorityQueue) Len() int {
	return len(q)
}

func (q V2PriorityQueue) Less(i, j int) bool {
	return q[i][0] < q[j][0]
}

func (q V2PriorityQueue) Swap(i, j int) {
	if i == j || i < 0 || j < 0 {
		return 
	}
	q[i], q[j] = q[j], q[i]
}

func (q *V2PriorityQueue) Push(x interface{}) {
	if nil == x {
		return
	}

	*q = append(*q, x.([]int))
}

func (q *V2PriorityQueue) Pop() interface{} {
	n := q.Len()
	if n == 0 {
		return nil
	}
	r := (*q)[n-1]
	*q = (*q)[:n-1]
	return r
}

func (q *V2PriorityQueue) Front() interface{} {
	if 0 == q.Len() {
		return nil
	}

	return (*q)[0]
}