package leetcode

import (
	"container/list"
	"math"
	"sort"
)

// 16 三数之和 转化为求剩余两数之和距离target距离最近的节点
func threeSumClosest(nums []int, target int) int {

	sort.IntSlice(nums).Sort()

	min := math.MaxInt32
	for i := range nums {
		j := i + 1
		k := len(nums) - 1

		for {
			if j >= k {
				break
			}

			threeSum := nums[i] + nums[j] + nums[k]
			if math.Abs(float64(min-target)) > math.Abs(float64(target-threeSum)) {
				min = threeSum
			}

			if target > threeSum {
				j = j + 1
			} else if target < threeSum {
				k = k - 1
			} else {
				break
			}
		}
	}
	return min

}

// 19 删除链表的倒数第N个节点
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }

 * 想一次遍历，那么需要记录倒数第n个节点的位置，那么需要一个追踪游标来记录什么时候开始里节点有n个位置
 * end 用来记录位置，保持当end走到最尾部节点是，target就是倒数第n个节点
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	if head == nil {
		return nil
	}

	if n <= 0 {
		return head
	}

	temp := head
	var preTarget *ListNode
	var target *ListNode
	end := 0

Loop:
	for {
		if end == n {
			target = head
		} else if end > n {
			preTarget = target
			target = target.Next
		}

		if temp == nil {
			break Loop
		}

		temp = temp.Next
		end = end + 1

	}

	if target == head {
		head = head.Next
	} else {
		preTarget.Next = target.Next
	}

	return head
}

// 20. 有效的括号
// 栈的使用
// TIP 并不一定到 len(s)/2 处的括号一定与前括号匹配，也有可能在后面有单独的匹配，
// 测试用例来帮助你完善思想 "(([]){})"
func isValid(s string) bool {
	if s == "" {
		return true
	}
	if len(s)%2 == 1 {
		return false
	}

	ll := list.New()
	for _, v := range s {
		if ll.Len() <= 0 {
			ll.PushBack(v)
		} else {
			end := ll.Back()
			if v == ')' && end.Value == '(' || (v == ']' && end.Value == '[') || (v == '}' && end.Value == '{') {
				ll.Remove(end)
			} else {
				ll.PushBack(v)
			}
		}
	}

	return ll.Len() == 0
}
