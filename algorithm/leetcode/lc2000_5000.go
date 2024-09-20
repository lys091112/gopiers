package leetcode

import (
	"container/list"
	"fmt"
	"math"
	"sort"
	"strconv"

	"github.com/lys091112/gopiers/algorithm/leetcode/util"
)

// N:  1561 你可以获得的最大硬币数目
// 排序 过滤调最低的一组，然后去剩下一组中两两的小值
func maxCoins(piles []int) int {
	if len(piles) == 0 {
		return 0
	}
	if len(piles)%3 != 0 {
		return -1
	}

	// sort
	sort.Sort(sort.IntSlice(piles))

	num := 0
	s := len(piles) / 3
	for i := s; i < len(piles); i += 2 {
		num += piles[i]
	}
	return num
}

// N: 2024 考试的最大困扰度
func maxConsecutiveAnswers(answerKey string, k int) int {
	if answerKey == "" {
		return 0
	}

	return util.Max(getCnt(answerKey,'T', k), getCnt(answerKey,'F', k))
}

// getCnt 获取ij 区间内 ==c的数量不超过k的最大长度
func getCnt(answerKey string,c byte, k int) int {
	ans,cnt :=0, 0
	 i, j :=0, 0
	for {
		if  j >= len(answerKey)  {
			break
		}
		if answerKey[j] == c {
			cnt++
		}
		j++
		
		for cnt >= k {
			if answerKey[i] == c {
				cnt--
			}
			i++
		}
		
		ans = util.Max(ans, j-i+1)
	}
	return ans
}


// N: 最大网络秩
// O(N^2)  考虑一种O(M+N)复杂度的算法
func maximalNetworkRank(n int, roads [][]int) int {
	if n == 0 {
		return 0
	}
	// 记录出度  记录边
	out := make([]int, n, 0)
	hasRoad := util.InitArray(n,n)

	// 出度和 减去共有边 
	for i := 0; i < len(roads); i ++ {
		out[roads[i][0]]++
		out[roads[i][1]]++
		
		hasRoad[roads[i][0]][roads[i][1]] = 1
		hasRoad[roads[i][1]][roads[i][0]] = 1
	}

	max := 0
	for i := 0 ;i < n-1 ; i ++ {
		for j := 1; i < n; j ++ {
			if i == j {
				continue
			}
			if max < out[i] + out[j] {
				max = out[i] + out[j]
			}
			if hasRoad[out[i]][out[j]] == 1 {
				max = max - 1
			}
		}
	}

	return max

}

// dp[i][0] 代表左侧等于0 的个数 dp[i][1] 代表右侧等于1的个数
// dp[i][0] = (dp[i-1][0], dp[i-1][0] + 1 {if s[i] == '0'})
// dp[i][1] = (dp[i+1][1], dp[i+1][1] + 1 {if s[i+1] == '1'})
// 主要是要区分分割的临界值   1） 两个非空字符串 因此要求一定要有分割  2） 动态规划 一定要有明确的初始值
func maxScore(s string) int {

	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = []int{0, 0}
	}

	if s[0] == '0' {
		dp[0][0] = 1
	}
	dp[len(s)-1][1] = 0

	for i := 1; i < len(s)-1; i++ {
		if s[i] == '0' {
			dp[i][0] = dp[i-1][0] + 1
		} else {
			dp[i][0] = dp[i-1][0]
		}
	}

	for i := len(s) - 2; i >= 0; i-- {
		if s[i+1] == '1' {
			dp[i][1] = dp[i+1][1] + 1
		} else {
			dp[i][1] = dp[i+1][1]
		}
	}

	max := 0
	for _, v := range dp {
		fmt.Printf("%v\n", v)
		if max < v[0]+v[1] {
			max = v[0] + v[1]
		}
	}
	return max
}

// N: 5394 对角线遍历 II
// 中序遍历 + BFS
// left = (row+1,col) left = (row,col+1)
// 由于同一个节点可能被多次访问，因此需要记录
func findDiagonalOrder(nums [][]int) []int {
	if len(nums) <= 1 {
		return nums[0]
	}

	result := make([]int, 0)
	result = append(result, nums[0][0])

	stack := list.New()
	stack.PushBack([]int{0, 0})
	row := len(nums)

	way := make(map[string]bool, 0)
	way["0_0"] = true

	for {
		if stack.Len() <= 0 {
			break
		}
		top := stack.Front()
		stack.Remove(top)
		tv := top.Value.([]int)

		// 找左节点
		r, c := tv[0]+1, tv[1]
		key := ""
		if r < row && c < len(nums[r]) {
			key = strconv.Itoa(r) + "_" + strconv.Itoa(c)
			if ok := way[key]; !ok {
				way[key] = true
				result = append(result, nums[r][c])
				stack.PushBack([]int{r, c})
			}
		}

		// 找右节点
		r, c = tv[0], tv[1]+1
		if r < row && c < len(nums[r]) {
			key = strconv.Itoa(r) + "_" + strconv.Itoa(c)
			if ok := way[key]; !ok {
				way[key] = true
				result = append(result, nums[r][c])
				stack.PushBack([]int{r, c})
			}
		}
	}
	return result
}

// N: 5649 解码异或后的数组
// v = first ^ second 
//  first ^ second ^ first = second
func decode(encoded []int, first int) []int {
	if len(encoded) <= 0 {
		return []int{}
	}
	res := make([]int, 1)
	res[0] = first
	tmp := first
	for _, v := range encoded {
		tmp = v^tmp 
		res = append(res, tmp)
	}
	return res 
}

// N:5652
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func swapNodes(head *ListNode, k int) *ListNode {
	 count := 0
	 
	 countN := head
	 for countN != nil {
		 count++
		 countN = countN.Next
	 }

	 if count < 2 * (k-1) {
		k = count - k + 1 
	 }
	 fisrtIndex := k - 1
	 secondIndex := count - k 
	 var firstN ,secondN *ListNode
	 var idx int = 0
	 countN = head
	 for countN != nil {
		if idx == fisrtIndex {
			 firstN = countN 
		}
		if (idx == secondIndex) {
			secondN = countN 
			break
		} 
		idx++
		countN = countN.Next
	 }

	 tmpV := firstN.Val
	 firstN.Val = secondN.Val
	 secondN.Val = tmpV
	 return head
}


// N: 5639. 完成所有工作的最短时间
func minimumTimeRequired(jobs []int, k int) int {
	if(len(jobs) <= 0) {
		return 0
	}
	if len(jobs) == 1 {
		return jobs[0]
	}

	stack := make([]int,k)
	sL := sort.IntSlice(jobs)
	sort.Sort(sort.Reverse(sL))

	for i, v := range sL {
		if i < k {
			stack[i] = v
		}else {
			// 提取最小任务，然后新增
			idx := findMinIdx(stack)
			stack[idx] += v
		}
	}
	return  findMax(stack)
}

func findMax(nums []int) (max int){
	max = -1
	for _, v := range nums {
		if max < v {
			max = v
		}
	}
	return 
}

func findMinIdx(nums []int) (idx int) {
	idx = 0
	value := math.MaxInt32
	for i, v := range nums {
		if value > v {
			value = v
			idx = i
		}
	}
	return 
}