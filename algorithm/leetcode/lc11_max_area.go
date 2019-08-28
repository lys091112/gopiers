package leetcode

import "fmt"

// 思想： 面积的公式为长*宽
// 宽= high-low  长= 高度最低的那个木板值
// 因此算法的思想为，总是向着面积可能变大的放下方向移动，即，那个木板短，那么他的位置就要变
// 因为如果变化木板比较长的那个位置，那么新的面积，木板长度只会变小或者不变（取决去两个木板的最小值），但是宽地因为移动而变小
// 因此面积一定减少, 所以需要移动木板比较短的那个位置

func maxArea(height []int) int {

	if len(height) < 2 {
		return 0
	}

	//
	low, high := 0, len(height)-1
	left, right := 0, len(height)-1
	max := 0

	for {
		// 临界条件
		if low >= high {
			break
		}

		t := (high - low) * If(height[low] > height[high], height[high], height[low]).(int)
		if t > max {
			max = t
			left = low
			right = high
		}

		if height[low] > height[high] {
			high = high - 1
		} else {
			low = low + 1
		}
	}

	fmt.Printf("left=%d, right=%d", left, right)

	return max
}
