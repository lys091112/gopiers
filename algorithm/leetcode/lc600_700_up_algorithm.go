package leetcode

import (
	"container/heap"
	"container/list"

	"github.com/lys091112/gopiers/algorithm/leetcode/util"
)

/**
 * N :675 高尔夫砍树
 *
 * 输入二维数组， 限制条件： 树木要从低到高砍
 */
func cutOffTree(forest [][]int) int {
	if len(forest) <= 0 {
		return 0
	}

	// 根据树高进行排序
	trees := make([][]int, 0)
	for i := range forest {
		for j := range forest[i] {
			if forest[i][j] == 0 {
				continue
			}
			trees = append(trees, []int{forest[i][j], i, j})
		}
	}
	util.Int2DSlice(trees).Int2DSort()

	distance := 0
	// 计算下个点到下下个点之间的距离 dist() 计算两个坐标间的最短距离（考虑A* 或者其他方法)
	// 从（0，0） 出发 到下个点
	x1, y1 := 0, 0
	for i := range trees {
		td := dist(forest, x1, y1, trees[i][1], trees[i][2])
		if td == -1 {
			return -1
		}
		distance += td
		x1 = trees[i][1]
		y1 = trees[i][2]
	}

	return distance
}

// a*
func dist(forest [][]int, x1, y1, x2, y2 int) int {
	r, c := len(forest), len(forest[0])

	var q util.Priority2DQueue = make([][]int, 0)
	seen := make(map[int]int)
	q.Push([]int{0, 0, x1, y1})

	seen[x1*c+y1] = 0

	orient := [][]int{{-1, 0, 1, 0}, {0, -1, 0, 1}}
	for {
		if q.Len() <= 0 {
			break
		}
		v := heap.Pop(&q).([]int)
		if v[2] == x2 && v[3] == y2 {
			return v[1]
		}

		// 计算
		for i := 0; i < 4; i ++ {
			nr := v[2] + orient[0][i]
			nc := v[3] + orient[1][i]

			if nr >= 0 && nr < r && nc >= 0 && nc < c && forest[nr][nc] > 0 {
				ncost := v[1] + 1 + util.Abs(nr-x2) + util.Abs(nc-y2)
				if cost, ok := seen[nr*c+nc]; (ok && ncost < cost) || !ok {
					heap.Push(&q, []int{ncost, v[1] + 1, nr, nc})
					seen[nr*c+nc] = ncost
				}
			}
		}

	}

	return -1
}

// bfs
func BFS(forest [][]int, x1, y1, x2, y2 int) int {
	r, c := len(forest), len(forest[0])
	seen := make([][]bool, r)
	for i := 0; i < r; i ++ {
		seen[i] = make([]bool, c)
	}
	q := list.New()
	q.PushBack([]int{x1, y1, 0})
	seen[x1][y1] = true

	orient := [][]int{{-1, 0, 1, 0}, {0, -1, 0, 1}}
	for {
		if q.Len() <= 0 {
			break
		}

		ele := q.Front()
		n := ele.Value.([]int)
		if n[0] == x2 && n[1] == y2 {
			return n[2]
		}
		q.Remove(ele)

		for i := 0; i < 4; i ++ {
			nr := n[0] + orient[0][i]
			nc := n[1] + orient[1][i]

			if nr >= 0 && nr < r && nc >= 0 && nc < c && !seen[nr][nc] && forest[nr][nc] > 0 {
				q.PushBack([]int{nr, nc, n[2] + 1})
				seen[nr][nc] = true
			}
		}
	}
	return -1
}

/**
 * N: 689  三个无重叠子数组的最大和
 * 思想： 动态规划
 * 状态转移方程：
 *    dp[i][n]=max(dp[i-1][n], dp[i-k][n-1]+sumRange(i-k+1, i));
 */
func maxSumOfThreeSubarrays(nums []int, k int) []int {

	if len(nums) < 3*k {
		return []int{}
	}

	prefixSum := make([]int, len(nums))
	prefixSum[0] = nums[0]

	// 前缀和
	for i := 1; i < len(nums); i ++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
	}

	// 动态规划 dp 记录
	dp := initArray(len(nums), 3)
	for n := 0; n < 3; n++ {
		for i := 0; i < len(nums); i ++ {
			prevL := 0
			if i-1 >= 0 {
				prevL = dp[i-1][n]
			}
			prevK := 0
			tailSum := 0
			if i-k >= 0 {
				if n > 0 {
					prevK = dp[i-k][n-1]
				}
			}
			if i-k >= 0 {
				tailSum = prefixSum[i] - prefixSum[i-k]
			} else {
				tailSum = prefixSum[i]
			}
			dp[i][n] = maxInt(prevL, prevK+tailSum)
		}
	}

	//
	result := make([]int, 0)
	curr := len(nums) - 1
	n := 2

	for {
		if n < 0 {
			break
		}
		v := dp[curr][n]
		// 因为要找最小序
		for {
			if curr > 0 && dp[curr-1][n] == v {
				curr -= 1
			} else {
				break
			}

		}

		result = append(result, curr-k+1)
		n -= 1
		curr -= k
	}

	return Reverse(result)
}

func maxInt(left, right int) int {
	if left > right {
		return left
	}
	return right
}
