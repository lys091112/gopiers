package leetcode

// 二叉搜索树集合

/**
 * go中切片的引用传递 以及反转
 */
// 315 计算右侧小于当前元素的个数
func countSmaller(nums []int) []int {
	if len(nums) <= 0 {
		return nil
	}

	if len(nums) == 1 {
		return []int{0}
	}

	length := len(nums)
	result := make([]int, length)

	root := &TreeNode{Val: nums[length-1], Count: 0}

	for i := length - 2; i >= 0; i -- {
		node := &TreeNode{Val: nums[i], Count: 0}
		insertNode(root, node, result, length-i-1)
	}

	// 切片反转
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

// Count 记录左子树中比该值小的元素的个数
func insertNode(parent *TreeNode, node *TreeNode, result []int, index int) {

	if parent.Val >= node.Val {
		parent.Count += 1
		if parent.Left == nil {
			parent.Left = node
		} else {
			insertNode(parent.Left, node, result, index)
		}

	} else if parent.Val < node.Val {
		// 碰到符合的元素，就加上比该元素小的值，以及他自身 一层层遍历
		result[index] += parent.Count + 1
		if parent.Right == nil {
			parent.Right = node
		} else {
			insertNode(parent.Right, node, result, index)
		}
	}
}

func countSmaller2(nums []int) []int {
	if len(nums) <= 0 {
		return nil
	}

	if len(nums) == 1 {
		return []int{0}
	}

	keys := make(map[int]int, len(nums))
	pairNums := make([]Pair, len(nums))
	for i, v := range nums {
		keys[i] = 0
		pairNums[i] = Pair{Key: v, Value: i}
	}
	mergeAndSort(keys, pairNums, 0, len(nums)-1)

	result := make([]int, len(nums))
	for i := range nums {
		result[i] = keys[i]
	}

	return result
}

// 先分治 后合并
func mergeAndSort(keys map[int]int, nums []Pair, low int, high int) {
	if low >= high {
		return
	}
	mid := (low + high) / 2
	mergeAndSort(keys, nums, low, mid)
	mergeAndSort(keys, nums, mid+1, high)

	mergeCallSamller(keys, nums, low, mid, high)
}

func mergeCallSamller(keys map[int]int, nums []Pair, low int, mid int, high int) {
	t := make([]Pair, high-low+1)
	i, j := low, mid+1
	k := 0

	for ; i <= mid && j <= high; {
		if nums[i].Key.(int) > nums[j].Key.(int) {
			t[k] = nums[j]
			j += 1
		} else {
			keys[nums[i].Value.(int)] += j - mid - 1
			t[k] = nums[i]
			i += 1
		}
		k += 1
	}

	for {
		if i > mid {
			break
		}
		t[k] = nums[i]
		keys[nums[i].Value.(int)] += j - mid - 1
		i += 1
		k += 1

	}

	for {
		if j > high {
			break
		}
		t[k] = nums[j]
		j += 1
		k += 1
	}
	for idx, v := range t {
		nums[idx+low] = v
	}
}

// 327. 区间和的个数
func countRangeSum(nums []int, lower int, upper int) int {

	return 0
}
