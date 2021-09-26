package weekly

import (
	"fmt"
	"math"
	"sort"
)

// 5661. 替换隐藏数字得到的最晚时间
func maximumTime(time string) string {
	if  len(time) <= 0 {
		return ""
	}

	cs := []byte(time)

	one, two := cs[0],cs[1]

	hour:= getMaxTimeHour(one,two)
	min := getMaxMin(cs[3],cs[4])

	return hour + ":" + min
}

func getMaxMin(one,two byte) string {
	if one == '?' {
		if two == '?' {
			return "59"
		}
		return "5" + string(two)
	}
	if two == '?' {
		return string(one) + "9"
	}

	return string(one) + string(two)
}

func getMaxTimeHour(one,two byte) string{
	if one == '?' {
		if two == '?' {
			return "23"
		}
		if two - 48 <= 3 {
			return "2" + string(two)
		}
		return "1" + string(two)
	}

	if two == '?' {
		if one - 48 <= 1 {
			return string(one) + "9"
		}
		return string(one) + "3"
	}

	return string(one)+string(two)
}

// 5662. 满足三条件之一需改变的最少字符数
func minCharacters(a string, b string) int {
	if a == "" || b == "" {
		return 0
	}


	maxCa,minCa := getMaxChar(a)
	maxCb,minCb := getMaxChar(b)
	la , lb := len(a), len(b)

	minTimes := math.MaxInt32
	 
	// fc
	if la == lb {
		minTimes = fcTimes(a,b)
	}

	fa := fabTimes(b,maxCa,minCa)
	fb := fabTimes(a,maxCb,minCb)
	if fa <= fb && minTimes > fa {
		minTimes = fa

	}else if fa >= fb && minTimes > fb {
		minTimes = fb

	}
	return minTimes
}

func getMaxChar(a string) (byte,byte) {
	ca := []byte(a)	
	maxC := ca[0] 
	minC := ca[0]

	for i := 1;i < len(a); i ++ {
		if maxC < ca[i] {
			maxC = ca[i]
		}
		if minC > ca[i] {
			minC = ca[i]
		}
	}
	return maxC,minC
}

func fabTimes(c string, maxB byte,minB byte) int{
	cb := []byte(c)

	times := 0
	minTimes := 0
	for i :=0 ; i < len(c); i ++ {
		// c 中所有的字母大于MaxB
		if cb[i] <= maxB {				
			times++
		}

		// c 中所有的字母小于minB
		if cb[i] >= minB {				
			minTimes++
		}
	}

	if times> minTimes {
		return minTimes 
	}
	
	return times 
}

func fcTimes(a, b string) int {
	m := make(map[byte]int)
	times := 0
	for i := 0; i < len(a); i++ {
		m[a[i]] ++ 
		if m[a[i]] > times {
			times = m[a[i]]
		}
	}
	for i := 0; i < len(b); i++ {
		m[b[i]] ++ 
		if m[b[i]] > times {
			times = m[b[i]]
		}
	}

	return len(a) + len(b) - times
}

// m[i][j] = matrix[i][j] xor m[i-1][j] xor m[i][j-1] xor m[i-1][j-1]
// 5663. 找出第 K 大的异或坐标值
func kthLargestValue(matrix [][]int, k int) int {
	if len(matrix) == 0 {
		return 0
	}
	
	m := len(matrix[0])

	tmp := make(map[int]int)
	tmp[0] = matrix[0][0]

	for i := range matrix {
		for j := range matrix[i] {				
			idx := i * m + j
			if i > 0 {
				if j == 0 {
					tmp[idx] = matrix[i][j] ^ tmp[(i-1)* m + j]
				}else {
					tmp[idx] = matrix[i][j] ^ tmp[i* m + j - 1] ^ tmp[(i-1)* m + j] ^ tmp[(i-1)* m + j-1]
				}

			}else {
				if j == 0 {
					tmp[idx] = matrix[i][j]
				}else {
					tmp[idx] = matrix[i][j]^tmp[i * m  + j-1]
				}
			}
		}
	}

	// 取 tmp 中的topN 元素 
	values := make([]int,len(tmp))
	for k, v := range tmp {
		values[k] = v
	}
	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	fmt.Printf("%v", values)


	return values[k-1]
}

// N: 1739. 放置盒子
// 题解： https://leetcode-cn.com/problems/building-boxes/solution/fang-zhi-he-zi-zhong-gui-zhong-ju-de-si-wfjiu/
func minimumBoxes(n int) int {

	level, nodes := 1,0

	for {
		t := nodes + level * (level + 1) / 2
		if t >= n {
			break
		}
		nodes = t
		level++
	}

	// 底部面积
	area := level * (level + 1) / 2

	// 记录剩下的node
	// 第一层每增加1，可以使当前层增加nowLeve+1个格子
	nowLevel := 0
	for nodes < n {
		area++
		nodes += nowLevel + 1
		nowLevel++
	}

	return area
}

// 正方形数列的推导： https://www.zhihu.com/question/408552181  n * (n+1) * (n + 2) / 6
// 底面最少，能盛放的总数量是个等差数列，1，3，6，10 即 n*(n+1) / 2
// 所以需要计算一共可以摆放多少层，sum(x*(x+1) / 2) <= n,求x的最大值，即可得出等差数列可以摆放的最大数量
// 然后在补足剩余的盒子
func minimumBoxesErFen(n int) int {
	low,high := 1, n

	for low < high {
		mid := (low + high) / 2
		if mid * (mid + 1) * (mid + 2) / 6 > n {
			high = mid
		}else {
			low = mid + 1
		}
	}
	low--


	area := low * (low+1) / 2

	// 已经放置的盒子总数
	cell := low * (low + 1) * (low + 2) / 6

	remain := n - cell
	low, high = 1, remain
	for low < high {
		mid := (low + high) / 2
		if mid * (mid + 1) * (mid + 2) / 6 > remain {
			high = mid
		}else {
			low = mid + 1
		}
	}
	return area + low
}