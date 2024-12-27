package leetcode

import (
	"container/heap"
	"fmt"

	"github.com/lys091112/gopiers/algorithm/util"
)

/**
 * N:502 IPO
 *  首先将数据按照起始资本进行分类, 然后在map中遍历并删除
 *
 *  ERROR way: 用map记录某个数据是否被访问过， 然后每次都要重新遍历全部的Capital,如果被访问过则忽略，这种方式
 *  的问题是会造成执行超时，
 */
func findMaximizedCapital(k int, W int, Profits []int, Capital []int) int {
	if k <= 0 {
		return 0
	}
	fmt.Printf("profits=%v, Capital=%v\n", Profits, Capital)

	m := make(map[int][]int, len(Capital))
	for i, v := range Capital {
		if _, ok := m[v]; !ok {
			m[v] = make([]int, 0)
		}
		m[v] = append(m[v], Profits[i])
	}

	// 初始化堆元素时，一定要默认长度为0， 不然会引入冗余数据，影响结果
	var h util.BigTopHeap = make([]int, 0)
	heap.Init(&h)

	// 停止条件为没有项目或者当前的k已经执行完
	for {
		if k <= 0 {
			break
		}
		// 遍历起步资金，对所有小于当前资本的项目建立大顶堆，收益最大的先被执行
		for key, value := range m {
			if key <= W {
				for _, v := range value {
					heap.Push(&h, v)
					fmt.Printf("input heap key=%d, value=%d\n", key, v)
				}
				delete(m, key)
			}
		}
		if h.Len() <= 0 {
			break
		}
		v := heap.Pop(&h)
		W += v.(int)
		k--
		fmt.Printf("chioce: key=%d, value=%d,W=%d\n", k, v.(int), W)
	}
	return W

}

/**
 * N:540
 * 还有一种o(n) 的位运算方法：(可以用来应对无序数据）
 * for(int i : nums) num ^= i;
 *
 * 本题的一个限制条件是 输入是有序的数组
 */
func singleNonDuplicate(nums []int) int {

	if len(nums) <= 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	left := 0
	right := len(nums) - 1

	for {
		if left >= right {
			break
		}
		mid := (left + right) / 2

		if isOdd(mid) {
			if nums[mid] == nums[mid+1] {
				right = mid - 1
			} else if nums[mid] == nums[mid-1] {
				left = mid + 1
			} else {
				left, right = mid, mid
			}

		} else {
			if nums[mid] == nums[mid+1] {
				left = mid + 2
			} else if nums[mid] == nums[mid-1] {
				right = mid - 2
			} else {
				left, right = mid, mid
			}
		}

	}

	return nums[right]

}

func isOdd(n int) bool {
	return n%2 == 1
}

// N: 560 和为K的子数组
func subarraySum(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	// <前缀和 , 在当前下标下，该前缀和存在的次数>
	sfMap := map[int]int{}
	sfMap[0] = 1
	res, t := 0, 0

	// 在遍历中就可以进行计算，因为我们只计算比
	//当前下标小的值，而这些值已经被计算过了
	// 查询当前前缀和-k（即已经计算过的前缀和）是否存在，以及存在的次数
	for _, v := range nums {
		t += v
		// 存在，那么增加统计量
		if value, ok := sfMap[t-k]; ok {
			res += value
		}
		sfMap[t]++
	}
	fmt.Printf("m=%v\n", sfMap)
	return res

}
