package leetcode

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"

	"github.com/lys091112/gopiers/algorithm/leetcode/util"
)

// N: 1879
// tags: 状压DP
// description:
// 基于状压DP， 状态递推公式： dp[i][s] = min(dp[i][s], dp[i-1][s ^ (1 << (j-1))] + nums1[i-1] ^ nums2[j-1]) 其中 i表示当前只考虑nums1的前i个数据， s的取值为（0 ~ 2^n-1)
// 数据的校验规则：
//  1. 由于只考虑nums1的前i个数，所以s中的1的个数一定和 i 相等
//  2. numbs[i] 和 nums2[j] 配对，需要保证nums2的第j位一定为1 即 (s>>(j-1))&1 = 1
//
// cost:
//  1. 时间复杂度 O(n^2 * 2^n)
//  2. 空间复杂度 O(n^2 * 2^n)
func minimumXORSum(nums1 []int, nums2 []int) int {
	if len(nums1) == 0 {
		return 0
	}
	n := len(nums1)
	mask := 1 << n
	dp := util.InitArray(n+10, mask, math.MaxInt)
	dp[0][0] = 0

	for i := 1; i <= n; i++ {
		for s := 0; s < mask; s++ {
			if getCnt(s, n) != i {
				continue
			}
			for j := 1; j <= n; j++ {
				// 判断s的第j位是否为1
				if 1&(s>>(j-1)) != 1 {
					continue
				}
				// TIP: 这里有运算优先级的问题 + 的优先级等同于 ^，如果不把^阔起来，会导致+号先执行
				dp[i][s] = util.Min(dp[i][s], dp[i-1][s^(1<<(j-1))]+(nums1[i-1]^nums2[j-1]))
				fmt.Printf("Now data is %d , %s, %d,%d\n", i, strconv.FormatInt(int64(s), 2), j, dp[i][s])
			}
		}
	}

	return dp[n][mask-1]
}

// 获取s表示的二进制中，1的个数
func getCnt(s, n int) int {
	ans := 0
	for i := 0; i < n; i++ {
		ans += (s >> i) & 1
		// if (s>>i)&1 == 1 {
		// 	ans += 1
		// }
	}

	return ans
}

type calc func() int

// 退火算法的初始条件
type saConfig struct {
	initTemp  float64
	finalTemp float64
	alpha     float64
	N         int
}

func shuffle(nums []int) {
	for i := 0; i < len(nums); i++ {
		j := i + rand.Intn(len(nums)-i)
		swap(nums, i, j)
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]

}

func sa(saConfig *saConfig, nums2 []int, calc calc) {
	shuffle(nums2)
	n := len(nums2)
	for i := 0; i < saConfig.N; i++ {

		for t := saConfig.initTemp; t > saConfig.finalTemp; t *= saConfig.alpha {
			a, b := rand.Intn(n), rand.Intn(n)
			prev := calc()
			swap(nums2, a, b)
			curr := calc()

			// 接收的条件：1 效果变好 即 prev > curr 2. 效果变化，但是有概率能接受
			// 反过来讲：效果不好且在概率外就不接受
			// 在这里 如果效果不好，即curr-prev > 0 ,则该结果一定大于1,即会重置数据，不接受该结果，如果效果好，则会有一定概率不接受该结果
			if math.Exp(float64(curr-prev)/t) > rand.Float64() {
				swap(nums2, a, b)
			}
		}

	}
}

// N: 1879
// tags: 退火算法
// description:
//
//	https://blog.csdn.net/weixin_62457573/article/details/135489346
func minimumXORSumV2(nums1 []int, nums2 []int) int {
	res := math.MaxInt
	saConfig := saConfig{
		initTemp:  1e5,
		finalTemp: 1e-5,
		alpha:     0.9,
		N:         30,
	}
	n2 := nums2[:]
	sa(&saConfig, n2, func() int {
		ans := 0
		for i := range nums1 {
			ans += nums1[i] ^ n2[i]
		}
		if ans < res {
			res = ans
		}

		return ans
	})

	return res
}
