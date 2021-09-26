package weekly

// 5654. 盒子中小球的最大数量
func countBalls(lowLimit int, highLimit int) int {
	max := 0
	battles := make(map[int]int)

	for i := lowLimit; i <= highLimit; i ++ {
		num := getNumber(i)
		battles[num]++
		if battles[num] > max {
			max = battles[num]
		} 
	}

	return max
}

func getNumber(i int) int {
	if i < 10 {
		return i
	}

	a := i % 10
	return a + getNumber(i / 10)
}

// N: 5665. 从相邻元素对还原数组
func restoreArray(adjacentPairs [][]int) []int {
	if len(adjacentPairs) == 0 {
		return []int{}
	}

	if len(adjacentPairs) == 1 {
		return adjacentPairs[0]
	}

	res := make([]int,0)
	record := make(map[int][]int)

	for i, v := range adjacentPairs {
		left,right := v[0],v[1]

		if _,ok := record[left]; !ok {
			record[left] = make([]int,0)
		}
		record[left] = append(record[left],i)

		if _,ok := record[right]; !ok {
			record[right] = make([]int,0)
		}
		record[right] = append(record[right],i)
	}

	// 找一个长度为1的做为起始点
	start := -1
	for k, v := range record {
		if len(v) == 1{
			start = k
			break
		}
	}

	way := make(map[int]bool)
	preV := start
	for {
		idx := record[start]
		res = append(res,start)
		preV = start

		var v []int
		if !way[idx[0]]  {
			// 访问
			v = adjacentPairs[idx[0]]
			way[idx[0]] = true
		}else {
			if len(idx) <= 1 || way[idx[1]] {
				break
			}
			way[idx[1]] = true
			v = adjacentPairs[idx[1]]
		}
		if v[0] == preV {
			start = v[1]
		}else {
			start = v[0]
		}

	}

	return res
}

// N: 5667. 你能在你最喜欢的那天吃到你最喜欢的糖果吗？
func canEat(candiesCount []int, queries [][]int) []bool {

	if len(queries) <= 0 {
		return []bool{}
	}

	preNums := make([]int,len(candiesCount))
	preNums[0] = 0

	// 前缀和
	for i, v := range candiesCount {
		if i == 0 {
		preNums[i] = v

		}else {
		preNums[i] = preNums[i-1] + v
		}
	}

	res := make([]bool,len(queries))

	for i, v := range queries {
		min,max := v[1] + 1 , (v[1] + 1) * v[2]

		if v[0] == 0 {
			res[i] = preNums[0] > min
		}else {
			// 不够吃 或者 吃不了
				res[i] = !(preNums[v[0]] < min || preNums[v[0]-1]+1 > max)
			// if preNums[v[0]] < min || preNums[v[0]-1] > max {
			// 	res[i] = false
			// }else {
			// 	// 判断长度 
			// 	// TODO 某个数组 能在几天内消灭完,贪心检查
			// 读题 和天无关 第二天不一定最多只能吃v[2]个
			// 	can := preNums[v[0]]
			// 	for j := 0; j <= v[1]; j ++ {
			// 		can -= v[2]
			// 		if can < 0 {
			// 			break
			// 		}
			// 	}
			// 	res[i] = can < candiesCount[v[0]]
			// }
		}

	}

	return res
}

// 5666. 回文串分割 IV
// TODO 动态规划的状态表达式需要在斟酌
// count[i,3] 表示字符串0..i 需要包含 3个回文字符串
// 动态规划 count[0..i,3] = m[0..i,1](回文) +  count[i+1..n,2]

func checkPartitioning(s string) bool {

	return false
}
