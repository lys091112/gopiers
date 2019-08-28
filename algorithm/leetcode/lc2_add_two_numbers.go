package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */



// 主要是学习指针的使用
// 通过遍历指针以及移动指针来获取数值
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var result *ListNode
	var end *ListNode
	p1 := l1
	p2 := l2

	n := 0
	for {
		if p1 == nil && p2 == nil && n == 0 {
			break
		}
		r1 := fetchValue(p1)
		r2 := fetchValue(p2)

		t := r1 + r2 + n
		if t >= 10 {
			n = 1
			t = t % 10
		} else {
			n = 0
		}

		if result == nil {
			result = &ListNode{Val: t, Next: nil}
			end = result
		} else {
			newNode := &ListNode{Val: t, Next: nil}
			end.Next = newNode
			end = newNode
		}

		if p1 != nil {
			p1 = p1.Next
		}
		if p2 != nil {
			p2 = p2.Next
		}

	}

	// 考虑遍历完成后，仍有进位的情况
	if n != 0 {
		newNode := &ListNode{Val: n, Next: nil}
		end.Next = newNode
		end = newNode
	}

	return result

}

func fetchValue(p *ListNode) int {
	if p == nil {
		return 0
	}
	return p.Val
}
