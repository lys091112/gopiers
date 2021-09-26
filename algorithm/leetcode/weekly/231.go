package weekly

import (
	"fmt"
	"math"
)

// N:5697. 检查二进制字符串字段
func checkOnesSegment(s string) bool {
	if len(s) == 0 {
		return false
	}

	end := false
	for _,v := range s {
		if v == '1' && end {
			return false
		}else if v == '0' {
			end = true
		}
	}
	return true
}

// N:5698. 构成特定和需要添加的最少元素 
func minElements(nums []int, limit int, goal int) int {

	sum := 0
	for _,v := range nums {
		sum += v
	}

	// 需要的变化量
	change := goal - sum
	if change == 0 {
		return 0
	}
	if(change < 0) {
		change = -change
	}


	// 贪心算法 ，多少个值之和等于change
	 return  (change + limit - 1) / limit
}

// N:5699. 从第一个节点出发到最后一个节点的受限路径数
// 求最短路径，并获取符合条件的值
var	distance = make(map[int]int) // 记录到n的距离

func countRestrictedPaths(n int, edges [][]int) int {

	weightEdge := make(map[string]int) // 记录边的权重
	record := make(map[string]bool) // 记录边的是否访问过

	distance[n] = 0
	table := make([][]int,n+1)
	for i := 1 ; i < n ; i ++ {
		distance[i] = math.MaxInt32
		table[i] = make([]int,0)
	}

	for _, v := range edges {
		table[v[0]] = append(table[v[0]],v[1])
		table[v[1]] = append(table[v[1]],v[0])
		weightEdge[fmt.Sprintf("%d_%d",v[0],v[1])] = v[2]
		weightEdge[fmt.Sprintf("%d_%d",v[1],v[0])] = v[2]
	}
	fmt.Printf("%v\n",table)

	dfsPath(n,table,record,weightEdge)
	fmt.Println("\n%m", distance)

	return 0
}

func edgeKey(u1,u2 int) string{
	return fmt.Sprintf("%d_%d",u1,u2);
}

func dfsPath(e int,table [][]int, record map[string]bool,weightEdge map[string]int) {
	
	next := make([]int,0)
	for _, v := range table[e] {
		key := edgeKey(e,v)

		// 如果新的距离比原来的较小，则更新
		hasChange := false
		if distance[e] + weightEdge[key] < distance[v] {
			fmt.Printf("change(%d,%d) %d + %d < %d",e,v,distance[e], weightEdge[key] ,distance[v])
			distance[v] = distance[e] + weightEdge[key]
			fmt.Printf(" = %d\n",distance[v])
			hasChange = true
		}
		// 如果 v 访问过 ，并且本次未修改， 则跳过,
		if _,ok := record[key];!ok || hasChange {
			next = append(next,v)
		}

		// 记录下一个节点入队
		record[key] = true
	}

	// 遍历edge之后的队列是否为空，如果不为空，则继续遍历
	for _, v := range next {
		dfsPath(v,table,record,weightEdge)
	}
	
}

// N:5700. 使所有区间的异或结果为零