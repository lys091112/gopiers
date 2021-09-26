package weekly

import (
	"fmt"
	"math"
	"sort"
)

func countMatches(items [][]string, ruleKey string, ruleValue string) int {

	if len(items) == 0 {
		return 0
	}

	index := -1
	if "type" == ruleKey {
		index = 0
	}else if"color" == ruleKey {
		index = 1
	} else if "name" == ruleKey {
		index = 2
	}

	count := 0
	for _, v := range items {
		if ruleValue == v[index] {
			count++
		}

	}

	return count 
}

//5691. 通过最少操作次数使数组的和相等
// 计算的是每个可以变换的量合并后的数组，按倒序相加和大于等于两个数组的原始差值的个数
func minOperations(nums1 []int, nums2 []int) int {

	count1 := sumArray(nums1)
	count2 := sumArray(nums2)
	if count1 == count2 {
		return 0
	}
	maxN,minN := nums1,nums2
	
	sub := count1 - count2
	if count1 < count2 {
		sub = count2 - count1
		maxN,minN = nums2,nums1
	}

	// 无论如何也无法相等的情况
	if len(maxN) > 6 * len(minN) {
		return -1
	}

	subArray := make([]int,0)
	for _, v := range maxN {
		if 0 < v -1 {
			subArray = append(subArray, v - 1)
		}
	}
	for _, v := range minN {
		if 6 - v > 0 {
			subArray = append(subArray, 6-v)
		}
	}

	s2 := sort.IntSlice(subArray)
	sort.Sort(sort.Reverse(s2))
	

	count := 0 
	for _, v := range s2 {
		if v == 0 {
			continue
		}
		count++
		sub -= v
		if sub <= 0 {
			break
		}
	}

	return count
}

func sumArray(nums []int) int {
	count := 0
	for _,v := range nums {
		count += v
	}
	return count
}

// 5690. 最接近目标价格的甜点成本
// TODO 0-1背包算法解决
var ans int
func closestCost(baseCosts []int, toppingCosts []int, target int) int {


	// chioseCost := make([]int,2*len(toppingCosts))
	// l := len(toppingCosts)
	// for i, v := range toppingCosts {
	// 	chioseCost[i] = v
	// 	chioseCost[l+i] = 2*v
	// }

	for _, v := range baseCosts {

		fmt.Printf("v=%d,\n",v)
		findNear(v,target,toppingCosts,0)
		if ans == target {
			break
		} 
	}

	return ans
}


func findNear(base, target int,chioseCost []int,index int) {
	m := math.Abs(float64(base-target))
	n := math.Abs(float64(ans-target))

	if m < n  || (m == n && base < ans) {
		fmt.Printf("chioce index=%d,ans=%d\n",index,base)
		ans = base
	} 

	if base == target || index == len(chioseCost)  {
		return 
	}


	// 选择index
	// 不选择index
	findNear(base,target,chioseCost,index+1)
	findNear(base+chioseCost[index],target,chioseCost,index+1)
	findNear(base+2*chioseCost[index],target,chioseCost,index+1)
}


// 5692. 车队 II
func getCollisionTimes(cars [][]int) []float64 {

	return nil
}