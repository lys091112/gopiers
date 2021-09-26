package leetcode

import (
	"container/list"
	"fmt"
	"sort"

	"github.com/lys091112/gopiers/algorithm/base"
)

// N:925
// 每次都一个叠词，然后看匹配的数量是否大于该叠词长度
func isLongPressedName(name string, typed string) bool {
	if len(name) <= 0 {
		return true
	}

	prev := byte(0)

	targetIdx := 0
	count := 0
	targetCount := 0

	// 比对匹配数量 主要是要确定初始的赋值
	for i := 0;i <= len(name); i++ {
		if prev == 0 {
			prev = name[i]
			count ++
		}else if i < len(name) && name[i] == prev {
			count ++
		}else {
			for {
				if targetIdx == len(typed) || typed[targetIdx] != prev  {
					break
				}
				targetIdx++
				targetCount++
			}

			// 比较的核心
			if targetCount < count {
				return false
			}

			if(i < len(name)) {
				prev = name[i]
				targetCount = 0
				count = 1
			}
		}
	}


	return targetIdx == len(typed) 
}

// 双指针移动
func isLongPressedName02(name string, typed string) bool {
	if len(name) == 0 {
		return true
	}

	i , j := 0,0
	for j < len(typed) {
		if i < len(name) && name[i] == typed[j] {
			i++
			j++
		}else if j > 0 && typed[j-1] == typed[j] {				
			j++
		}else {
			return false
		}
	}
	return i == len(name)
}

// N: 933 窗口次数

const LENGTH_933 = 3001

// RecentCounter 窗口次数
// 循环数组
type RecentCounter struct {
	recentData  []int
	recentStart int
	recentEnd   int
}

// Constructor933  构造
func Constructor933() RecentCounter {

	return RecentCounter{recentData: make([]int, LENGTH_933+1),
		recentStart: 0, recentEnd: 0}
}

// Ping 返回3000毫秒前的数据
// 正确解法：使用FIFO队列，移除不合法的数据，然后统计剩下的数量
func (rc *RecentCounter) Ping(t int) int {
	i := rc.recentStart
	for {
		if i == rc.recentEnd {
			break
		}
		if t-rc.recentData[i] <= 3000 {
			break
		}

		i++
		if i > LENGTH_933 {
			i = 0
		}
	}

	// 0   1 ...  3001
	// end start
	// 输入3002是, start 先从1移动到2，然后 end从0 移动到 1
	// end 代表的是下一个要填充的位置，如果end == start 代表没有元素
	// 注意边界值，必须要start先腾出位置后，end再去覆盖
	rc.recentStart = i
	rc.recentData[rc.recentEnd] = t
	rc.recentEnd++
	if rc.recentEnd > LENGTH_933 {
		rc.recentEnd = 0
	}

	if rc.recentStart > rc.recentEnd {
		return rc.recentEnd + LENGTH_933 - rc.recentStart + 1
	}

	return rc.recentEnd - rc.recentStart
}

// N:946
// 判断三种情况
// 	1. 栈顶等于当前值，remove(stack.top) -> next one
//	2. 栈顶的索引大于当前值的索引，即该值已经被放入到栈内，无法被取出，直接返回
//	3. 遍历pushed栈，继续放入还未放入的元素
func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) <= 0 && len(popped) <= 0 {
		return true
	}
	m := make(map[int]int, len(pushed))
	for i, v := range pushed {
		m[v] = i
	}

	topIndex := 0
	stack := list.New()
	for i := 0; i < len(popped); i++ {
		top := stack.Front()
		v := popped[i]

		if top != nil && v == top.Value.(int) {
			stack.Remove(top)
		} else if top != nil && m[v] < m[top.Value.(int)] {
			return false
		} else {
			for {
				if topIndex < len(pushed) {
					if v != pushed[topIndex] {
						stack.PushFront(pushed[topIndex])
						topIndex++
					} else {
						topIndex++
						break
					}
				} else {
					break
				}
			}
		}
	}
	return true
}

// N:952 按公因数计算最大组件大小
// 对于较大数据集，如果两两比对公约数，可能会超时
// 考虑埃氏筛法，进行筛选
// TODO
func largestComponentSize(A []int) int {
	if len(A) <= 0 {
		return 0
	}

	find := base.UnionFind{}
	find.MakeSet(len(A))

	tmax := 0
	mv := make(map[int]int, len(A))
	for i, v := range A {
		if tmax < v {
			tmax = v
		}
		mv[v] = i
	}

	mp := make(map[int]bool, tmax)

	for i := 2; i < tmax; i++ {
		if b, ok := mp[i]; ok && b {
			continue
		}
		root := -1
		if v, ok := mv[i]; ok && v >= 0 {
			root = v
		}
		for j := i + i; j <= tmax; j += i {
			if v, ok := mv[j]; ok {
				if root == -1 {
					root = v
				} else {
					find.Union(root, v)
					fmt.Printf("compair, %d, (%d,%d)\n", i, A[root], j)
				}
			}
			mp[j] = true
		}
	}

	max := 0
	m := find.Collections()
	fmt.Println(m)
	for _, v := range m {
		if len(v) > max {
			max = len(v)
		}
	}
	return max
}

// 954. 二倍数对数组
// 按照一定的顺序，当前值一定有对应的值匹配，否则返回失败
// 注意在等比缩小时，奇偶数的问题
func canReorderDoubled(A []int) bool {
	if len(A) <= 0 {
		return false
	}

	m := make(map[int]int, len(A))

	for _, v := range A {
		m[v] += 1
	}

	sort.Ints(A)
	for _, d := range A {
		if m[d] <= 0 {
			continue
		}

		var n int
		if d < 0 {
			if (-d)%2 != 0 {
				return false
			}
			n = d / 2
		} else {
			n = d * 2
		}

		if v, ok := m[n]; ok && v > 0 {
			m[d] -= 1
			m[n] -= 1
		} else {
			return false
		}
	}

	return true
}
